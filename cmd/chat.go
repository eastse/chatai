package cmd

import (
	"fmt"
	"math"
	"path"
	"sort"
	"strconv"

	"chatai/app"
	"chatai/chat"
	"chatai/common/config"
	"chatai/common/snowflake"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sashabaranov/go-openai"
)

func toggleChat() {
	chatCfgs := []*config.ChatConfig{}
	for _, v := range app.Chats {
		chatCfgs = append(chatCfgs, v)
	}

	sort.Slice(chatCfgs, func(i, j int) bool {
		return chatCfgs[i].Sort > chatCfgs[j].Sort
	})

	using := 0
	opts := []string{}
	for i, v := range chatCfgs {
		opts = append(opts, v.Title)
		if v.ID == app.Setting.UsingChatID {
			using = i
		}
	}

	selected := 0
	prompt := &survey.Select{
		Message: "Select the chat to switch to:",
		Options: opts,
		Default: opts[using],
		Description: func(value string, index int) string {
			if index == using {
				return "Using"
			}
			return ""
		},
	}
	if err := survey.AskOne(prompt, &selected, survey.WithPageSize(10)); err != nil {
		return
	}

	chatCfg := chatCfgs[selected]
	if chatCfg.ID != app.Setting.UsingChatID {
		app.Setting.UsingChatID = chatCfg.ID
		chat.Load(chatCfg)
		app.UpdateConfig()
		clearScreen()
		chat.ShowChatInfo()
	}
}

func validatorFloat(min, max float64) survey.Validator {
	return func(val interface{}) error {
		if v, err := strconv.ParseFloat(val.(string), 64); err == nil {
			if v < min || v > max {
				return fmt.Errorf("invalid argument please input %v ~ %v flaot", min, max)
			}
		} else {
			return fmt.Errorf("invalid argument please input %v ~ %v flaot", min, max)
		}
		return nil
	}
}

func validatorInt(min, max int64) survey.Validator {
	return func(val interface{}) error {
		if v, err := strconv.ParseInt(val.(string), 10, 64); err == nil {
			if v < min || v > max {
				return fmt.Errorf("invalid argument please input %v ~ %v integer", min, max)
			}
		} else {
			return fmt.Errorf("invalid argument please input %v ~ %v integer", min, max)
		}
		return nil
	}
}

func newChat() {
	var question = []*survey.Question{
		{
			Name: "title",
			Prompt: &survey.Input{
				Message: "Title: ",
				Help:    "The title of the chat, used to distinguish",
			},
			Validate: survey.ComposeValidators(survey.Required, survey.MaxLength(60)),
		},
		{
			Name: "prompt",
			Prompt: &survey.Input{
				Message: "Prompt: ",
				Default: "",
				Help:    "The prompt(s) to generate completions",
			},
		},
		{
			Name: "model",
			Prompt: &survey.Select{
				Message: "Model: ",
				Default: openai.GPT3Dot5Turbo,
				Options: chat.GetModels(),
			},
		},
		{
			Name: "temperature",
			Prompt: &survey.Input{
				Message: "Temperature: ",
				Default: "1",
				Help:    "Defaults 1 What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.We generally recommend altering this or top_p but not both.",
			},
			Validate: validatorFloat(0, 2),
		},
		{
			Name: "review",
			Prompt: &survey.Input{
				Message: "Review: ",
				Default: "1",
				Help:    "The number of sent context information questions and answers together is 1, if the number is too large, too many tokens will be consumed",
			},
			Validate: validatorInt(0, math.MaxInt64),
		},
		{
			Name: "sort",
			Prompt: &survey.Input{
				Message: "Sort: ",
				Default: "0",
				Help:    "The heavier the sorting weight, the more priority it receives",
			},
			Validate: validatorInt(0, math.MaxInt64),
		},
	}

	chatID := snowflake.GetId().Int64()
	info := &config.ChatConfig{
		ID:          chatID,
		Developer:   config.DeveloperOpenAI,
		HistoryFile: path.Join(app.HISTORY_DIR, fmt.Sprintf("%v.md", chatID)),
	}
	err := survey.Ask(question, info)
	if err != nil {
		fmt.Printf("Invalid argument err:%v", err.Error())
		return
	}
	app.Chats[chatID] = info
	app.Setting.UsingChatID = chatID
	app.UpdateConfig()
	clearScreen()
	chat.Load(info)
	chat.ShowChatInfo()
}

func editChat() {
	info := chat.GetConfig()
	var question = []*survey.Question{
		{
			Name: "title",
			Prompt: &survey.Input{
				Message: "Title: ",
				Default: info.Title,
				Help:    "The title of the chat, used to distinguish",
			},
		},
		{
			Name: "prompt",
			Prompt: &survey.Input{
				Message: "Prompt: ",
				Default: info.Prompt,
				Help:    "The prompt(s) to generate completions",
			},
		},
		{
			Name: "model",
			Prompt: &survey.Select{
				Message: "Model: ",
				Default: info.Model,
				Options: chat.GetModels(),
				Description: func(value string, index int) string {
					if value == info.Model {
						return "Using"
					}
					return ""
				},
			},
		},
		{
			Name: "temperature",
			Prompt: &survey.Input{
				Message: "Temperature: ",
				Default: strconv.FormatFloat(float64(info.Temperature), 'f', -1, 32),
				Help:    "Defaults 1 What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.We generally recommend altering this or top_p but not both.",
			},
			Validate: validatorFloat(0, 2),
		},
		{
			Name: "review",
			Prompt: &survey.Input{
				Message: "Review: ",
				Default: strconv.Itoa(int(info.Review)),
				Help:    "The number of sent context information questions and answers together is 1, if the number is too large, too many tokens will be consumed",
			},
			Validate: validatorInt(0, math.MaxInt64),
		},
		{
			Name: "sort",
			Prompt: &survey.Input{
				Message: "Sort: ",
				Default: strconv.Itoa(int(info.Sort)),
				Help:    "The heavier the sorting weight, the more priority it receives.",
			},
			Validate: validatorInt(0, math.MaxInt64),
		},
	}

	err := survey.Ask(question, &info)
	if err != nil {
		fmt.Printf("Invalid argument err:%v", err.Error())
		return
	}
	app.Chats[info.ID] = &info
	app.UpdateConfig()
	fmt.Println("\033[32m Successfully!")
	chat.Load(&info)
	chat.ShowChatInfo()
}

func deleteChat() {
	chatCfgs := []*config.ChatConfig{}
	for _, v := range app.Chats {
		if v.ID == app.Setting.UsingChatID {
			continue
		}
		chatCfgs = append(chatCfgs, v)
	}

	if len(chatCfgs) == 0 {
		fmt.Println("\033[31mno conversations to delete!")
		return
	}

	sort.Slice(chatCfgs, func(i, j int) bool {
		return chatCfgs[i].Sort > chatCfgs[j].Sort
	})

	opts := []string{}
	for _, v := range chatCfgs {
		opts = append(opts, v.Title)
	}

	selected := []int{}
	prompt := &survey.MultiSelect{
		Message: "Select the chat to delete:",
		Options: opts,
		Help:    "Select the chat to delete. if you want to delete the current chat, please switch to other chat first",
	}
	survey.AskOne(prompt, &selected, survey.WithPageSize(10), survey.WithRemoveSelectAll())

	if len(selected) == 0 {
		return
	}

	confirm := false
	confirmPrompt := &survey.Confirm{
		Message: "You sure you want to delete it?",
	}
	err := survey.AskOne(confirmPrompt, &confirm)
	if err != nil {
		return
	}

	if confirm {
		for _, v := range selected {
			cfg := chatCfgs[v]
			chat.Delete(cfg)
		}
		app.UpdateConfig()
		fmt.Println("\033[32m Successfully!")
	}
}
