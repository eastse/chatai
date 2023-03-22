package main

import (
	"flag"
	"fmt"
	"strings"

	"chatai/app"
	"chatai/chat"
	"chatai/cmd"
)

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		arg := strings.Join(flag.Args(), " ")
		if cfg, ok := app.Chats[app.Setting.QuickChatID]; ok {
			chat.Load(cfg)
			chat.SendMessage(arg, nil)
		} else {
			fmt.Println("Please set up the conversation for the quick question and answer first!")
		}
	} else {
		cmd.Execute()
	}
}
