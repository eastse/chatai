package cmd

import (
	"fmt"
	"sort"
	"strconv"

	"chatai/app"
	"chatai/common/config"

	"github.com/AlecAivazis/survey/v2"
)

func setDefaultMultiline() {
	str := ""
	prompt := &survey.Select{
		Message: "Always enable multiline text input mode on startup:",
		Options: []string{"true", "false"},
	}

	if err := survey.AskOne(prompt, &str); err != nil {
		return
	}

	selected, _ := strconv.ParseBool(str)
	if selected != app.Setting.MultilineInput {
		app.Setting.MultilineInput = selected
		app.UpdateConfig()
	}
	fmt.Println("\033[32m Successfully!")
}

func selectQuickChat() {
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
		if v.ID == app.Setting.QuickChatID {
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

	err := survey.AskOne(prompt, &selected, survey.WithPageSize(10))
	if err != nil {
		return
	}

	chatCfg := chatCfgs[selected]
	if chatCfg.ID != app.Setting.QuickChatID {
		app.Setting.QuickChatID = chatCfg.ID
		app.UpdateConfig()
	}
	fmt.Println("\033[32m Successfully!")
}

func setInputPrompt() {
	str := ""
	prompt := &survey.Input{
		Message: "Prompt:",
		Help:    "Please enter the characters displayed when prompted:",
	}

	if err := survey.AskOne(prompt, &str, survey.WithValidator(survey.MaxLength(20))); err != nil {
		return
	}

	if str != app.Setting.InputPrompt {
		app.Setting.InputPrompt = str
		app.UpdateConfig()
		loadPromptStyle()
	}
	fmt.Println("\033[32m Successfully!")
}

func setTargetColor(target *string) {
	selected, err := selectColor()
	if err != nil {
		return
	}

	if selected != *target {
		*target = selected
		app.UpdateConfig()
		loadPromptStyle()
	}

	fmt.Println("\033[32m Successfully!")
}

func selectColor() (color string, err error) {
	prompt := &survey.Select{
		Message: "Choose a color:",
		Options: []string{"Red", "Green", "Yellow", "Blue", "Purple", "Cyan"},
	}
	err = survey.AskOne(prompt, &color)

	return
}
