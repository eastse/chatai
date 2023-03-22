package chat

import (
	"fmt"
	"os"
	"path"
	"reflect"

	"chatai/app"
	"chatai/chat/chatgpt"
	"chatai/common/config"

	"github.com/mattn/go-runewidth"
)

var (
	instance ChatEntity
)

type ChatEntity interface {
	ID() int64
	GetConfig() *config.ChatConfig
	SendMessage(msg string, ctrlC chan bool)
	GetModels() []string
	Close()
}

func Load(cfg *config.ChatConfig) {
	if instance != nil {
		instance.Close()
	}
	switch cfg.Developer {
	case config.DeveloperOpenAI:
		instance = chatgpt.New(cfg)
	}
}

func SendMessage(msg string, ctrl_c chan bool) {
	instance.SendMessage(msg, ctrl_c)
}

func GetConfig() config.ChatConfig {
	return *instance.GetConfig()
}

func GetModels() []string {
	return instance.GetModels()
}

func Delete(cfg *config.ChatConfig) {
	delete(app.Chats, cfg.ID)
	os.Remove(cfg.HistoryFile)
	os.Remove(path.Join(app.HOME_DIR, fmt.Sprintf(".%v", cfg.ID)))
}

func ShowChatInfo() {
	cfg := instance.GetConfig()
	content := "\n"

	typ := reflect.TypeOf(cfg).Elem()
	val := reflect.Indirect(reflect.ValueOf(cfg))

	for i := 0; i < val.NumField(); i++ {
		t := typ.Field(i).Name
		if t == "Type" {
			continue
		}
		v := val.Field(i)
		content = fmt.Sprintf("%v\033[36m   • %v:\033[m %v\n", content, t, v)
	}
	content += "\n"
	fmt.Print(content)
}

func ShowChatInfo2() {
	cfg := instance.GetConfig()
	content := "\n"

	typ := reflect.TypeOf(cfg).Elem()
	val := reflect.Indirect(reflect.ValueOf(cfg))

	keys := []string{}
	vals := []string{}
	maxWidth := 0
	for i := 0; i < val.NumField(); i++ {
		t := typ.Field(i).Name
		if t == "Type" {
			continue
		}
		v := val.Field(i)
		keys = append(keys, t)
		vals = append(vals, fmt.Sprint(v))
		if w := runewidth.StringWidth(t); w > maxWidth {
			maxWidth = w
		}
	}

	spaceText := "                                     "
	for i, k := range keys {
		w := runewidth.StringWidth(k)
		space := ""
		if count := maxWidth - w; count > 0 {
			space = spaceText[:count]
		}
		content = fmt.Sprintf("%v\033[36m   • %v:%v\033[m %v\n", content, k, space, vals[i])
	}

	content += "\n"
	fmt.Print(content)
}
