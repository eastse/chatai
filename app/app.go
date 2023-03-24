package app

import (
	"fmt"
	"os"
	"path"
	"time"

	"chatai/common/config"
	"chatai/common/snowflake"
	"chatai/common/utils"

	"github.com/briandowns/spinner"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

const (
	home_name string = "/.chatai"
)

var (
	Setting     config.SettingConfig
	Chats       map[int64]*config.ChatConfig
	sp          *spinner.Spinner
	HOME_DIR    string = ""
	HISTORY_DIR string = "/history"
	CFG_FILE    string = ".config"
)

func init() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(0)
	}

	HOME_DIR = path.Join(userHome, home_name)
	HISTORY_DIR = path.Join(HOME_DIR, HISTORY_DIR)
	CFG_FILE = path.Join(HOME_DIR, CFG_FILE)

	initConfig()

	sp = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
}

func initConfig() {
	viper.SetConfigFile(CFG_FILE)
	viper.SetConfigType("yaml")

	if !utils.PathExists(CFG_FILE) {
		// create directory
		err := os.MkdirAll(path.Dir(CFG_FILE), 0755)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}

		// create cfgfile
		file, err := os.Create(CFG_FILE)
		if err != nil {
			fmt.Println(err)
			return
		}
		file.Close()

		addInitConfig()
		return
	} else {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("read config err: %v", err)
			os.Exit(0)
		}
	}

	if err := viper.UnmarshalKey("settings", &Setting); err != nil {
		fmt.Printf("config unmarshal err: %v\n", err)
		os.Exit(0)
	}

	if err := viper.UnmarshalKey("chats", &Chats); err != nil {
		fmt.Printf("config unmarshal err: %v\n", err)
		os.Exit(0)
	}
}

func addInitConfig() {
	chatID := snowflake.GetId().Int64()
	Setting = config.SettingConfig{
		UsingChatID:       chatID,
		QuickChatID:       chatID,
		InputPrompt:       "\ue0c0  ",
		SinglePromptColor: "Green",
		MultiPromptColor:  "Cyan",
		SingleTextColor:   "Yellow",
		MultiTextColor:    "Yellow",
	}
	Chats = map[int64]*config.ChatConfig{
		chatID: {
			ID:          chatID,
			Title:       "default",
			Temperature: 1,
			Model:       openai.GPT3Dot5Turbo,
			Developer:   config.DeveloperOpenAI,
			HistoryFile: path.Join(HISTORY_DIR, "default.md"),
		},
	}
	UpdateConfig()
}

func UpdateConfig() {
	viper.Set("settings", Setting)
	viper.Set("chats", Chats)
	if err := viper.WriteConfig(); err != nil {
		fmt.Printf("write config err: %v", err)
		os.Exit(0)
	}
}

func GetChatConfig(id int64) *config.ChatConfig {
	return Chats[id]
}

func ShowSpinner() {
	sp.Start()
}

func StopSpinner() {
	if sp.Active() {
		sp.Stop()
	}
}
