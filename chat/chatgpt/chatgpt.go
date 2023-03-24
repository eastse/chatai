package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"

	"chatai/app"
	"chatai/common/config"
	"chatai/common/history"
	"chatai/common/renderer"
	"chatai/common/request"
	"chatai/common/utils"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var API_KEY string

type chatGPT struct {
	isFirst    bool
	config     *config.ChatConfig
	record     []*ChatMessage
	topic      *ChatMessage
	answer     *ChatMessage
	isCtrl_c   bool
	history    *history.History
	reviewFile string
}

func init() {
	API_KEY = os.Getenv("CHATGPT_API_KEY")
	if API_KEY == "" {
		fmt.Println("Please set your CHATGPT_API_KEY in your environment variable")
		os.Exit(0)
	}
}

func New(config *config.ChatConfig) *chatGPT {
	c := &chatGPT{
		isFirst:    true,
		config:     config,
		record:     []*ChatMessage{},
		topic:      &ChatMessage{Role: ChatMessageRoleUser},
		answer:     &ChatMessage{Role: ChatMessageRoleAssistant},
		history:    history.New(config.HistoryFile),
		reviewFile: path.Join(app.HOME_DIR, fmt.Sprintf(".%v", config.ID)),
	}
	c.loadReview()
	return c
}

func (c *chatGPT) ID() int64 {
	return c.config.ID
}

func (c *chatGPT) GetConfig() *config.ChatConfig {
	return c.config
}

func (c *chatGPT) SendMessage(message string, ctrl_c chan bool) {
	app.ShowSpinner()
	c.isFirst = true
	c.answer.Content = ""
	c.topic.Content = message
	c.isCtrl_c = false

	msg := []*ChatMessage{}
	if c.config.Prompt != "" {
		msg = append(msg, &ChatMessage{ChatMessageRoleSystem, c.config.Prompt})
	}
	msg = append(msg, c.record...)
	msg = append(msg, &ChatMessage{ChatMessageRoleUser, message})

	chatRequest := ChatRequest{
		Model:       GPT3Dot5Turbo,
		Temperature: c.config.Temperature,
		Messages:    msg,
		Stream:      true,
	}

	jsonData, err := json.Marshal(&chatRequest)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	defer c.updateRecord()

	ctx, cancel := context.WithCancel(context.Background())
	req := request.PostStream(ctx, fmt.Sprintf("%s%s", BASE_URL, CHAT_URL), bytes.NewBuffer(jsonData))
	req.Request.Header.Set("Content-Type", "application/json")
	req.Request.Header.Set("Accept", "text/event-stream")
	req.Request.Header.Set("Cache-Control", "no-cache")
	req.Request.Header.Set("Connection", "keep-alive")
	req.Request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))
	req.DoStream()
	defer cancel()

	for {
		select {
		case <-ctrl_c:
			app.StopSpinner()
			if req.Response != nil {
				req.Response.Body.Close()
			}
			c.isCtrl_c = true
			return
		case data, ok := <-req.C:
			if ok {
				c.process(data)
			} else {
				return
			}
		}
	}
}

func (c *chatGPT) Close() {
	c.history.Close()
}

func (c *chatGPT) ClearReview() {
	c.record = c.record[:0]
}

func (c *chatGPT) process(data []byte) {
	resp := &ChatResponse{}
	if err := json.Unmarshal(data, resp); err != nil {
		return
	}

	choice := resp.Choices
	if len(choice) > 0 {
		msg := choice[0]
		content := msg.Delta.Content
		if content != "" {
			c.answer.Content += content
			renderer.RenderStream(content, c.isFirst)
			if c.isFirst {
				c.isFirst = false
			}
		}
	}
}

func (c *chatGPT) updateRecord() {
	if c.isCtrl_c {
		return
	}

	c.writeHistoryData()

	if c.config.Review == 0 {
		return
	}

	if len(c.record) < int(c.config.Review)*2 {
		c.record = append(c.record, c.topic, c.answer)
		c.topic = &ChatMessage{Role: ChatMessageRoleUser}
		c.answer = &ChatMessage{Role: ChatMessageRoleAssistant}
	} else {
		c.record = append(c.record, c.topic, c.answer)
		c.topic = c.record[0]
		c.answer = c.record[1]
		c.record = c.record[2:]
	}

	c.writeReviewData()

}

func (c *chatGPT) writeHistoryData() {
	str := fmt.Sprintf("\n### %v:\n%v\n### %v:\n%v\n***\n",
		cases.Title(language.English).String(ChatMessageRoleUser), c.topic.Content,
		cases.Title(language.English).String(ChatMessageRoleAssistant), c.answer.Content)
	c.history.Write(str)
}

func (c *chatGPT) loadReview() {
	if c.config.Review == 0 || !utils.PathExists(c.reviewFile) {
		return
	}

	jsonData, err := os.ReadFile(c.reviewFile)
	if err != nil {
		fmt.Printf("read review data err: %v\n", err)
		return
	}

	err = json.Unmarshal(jsonData, &c.record)
	if err != nil {
		fmt.Printf("unmarshal review data err: %v\n", err)
	}
}

func (c *chatGPT) writeReviewData() {
	data, err := json.Marshal(c.record)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	os.WriteFile(c.reviewFile, data, 0644)
}
