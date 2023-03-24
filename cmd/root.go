package cmd

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unicode"

	"chatai/app"
	"chatai/chat"
	"chatai/common/config"

	"github.com/chzyer/readline"
	"github.com/common-nighthawk/go-figure"
)

var (
	reader            *readline.Instance
	isWork            = false
	multiline         = false
	isInput           = true
	isStopCmd         = false
	promptChar        = ""
	singlePromptColor = ""
	singleTextColor   = ""
	multiPromptColor  = ""
	multiTextColor    = ""
)

func Execute() {
	multiline = app.Setting.MultilineInput
	clearScreen()
	loadPromptStyle()
	chat.Load(app.GetChatConfig(app.Setting.UsingChatID))
	chat.ShowChatInfo()

	var ctrl_c = make(chan bool)

	go run(ctrl_c)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, syscall.SIGINT)
	for {
		<-interrupt
		if isWork {
			ctrl_c <- true
		} else {
			return
		}
	}
}

func run(ctrl_c chan bool) {
	reader, _ = readline.NewEx(&readline.Config{
		Prompt:          getPrompt(),
		AutoComplete:    completer,
		InterruptPrompt: "\033[31m ^C\033[33m",
		EOFPrompt:       "exit",
		FuncFilterInputRune: func(key rune) (rune, bool) {
			switch key {
			case readline.CharCtrlZ: // stop cmd
				if isStopCmd {
					isStopCmd = false
					println("\033[31m ^Z\n\033[0mReceived commands resumed\033[0m\n")
				}
				return key, false
			case 4: // ctrl+d
				if multiline {
					isInput = false
					reader.WriteStdin([]byte{'\n'})
					return key, false
				} else {
					return key, true
				}
			}
			return key, true
		},
	})

	defer reader.Close()
	defer os.Exit(0)

	cmds := make([]string, 0)
	for {
		isInput = true
		line, err := reader.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		isWork = true

		if multiline {
			cmds = append(cmds, line)
			if isInput {
				reader.SetPrompt("» ")
				continue
			}
			line = strings.Join(cmds, "\n")
			line = strings.TrimSpace(line)
			cmds = cmds[:0]
			reader.SetPrompt(getPrompt())
		} else {
			line = strings.TrimSpace(line)
		}

		if isStopCmd {
			chat.SendMessage(line, ctrl_c)
			isWork = false
			continue
		}

		switch {
		case line == "chat":
			toggleChat()

		case strings.HasPrefix(line, "chat "):
			switch line[5:] {
			case "new":
				newChat()
			case "edit":
				editChat()
			case "info":
				chat.ShowChatInfo()
			case "clear":
				clearReview()
			case "delete":
				deleteChat()
			default:
				fmt.Printf("%vinvalid cmd!\n", config.ColorRed)
			}

		case line == "multi-line":
			toggleMultiline()

		case line == "stop-cmd":
			isStopCmd = true
			println("\033[0mStop receiving commands Press \033[31mCtrl+Z\033[0m to resume\n")

		case line == "help":
			io.WriteString(reader.Stderr(), "\033[33m"+helpinfo)

		case line == "setting":
			fmt.Printf("app.Setting: %v\n", app.Setting)

		case strings.HasPrefix(line, "setting "):
			switch line[8:] {
			case "quick":
				selectQuickChat()
			case "multi-line":
				setDefaultMultiline()
			case "input-prompt":
				setInputPrompt()
			case "single-prompt-color":
				setTargetColor(&app.Setting.SinglePromptColor)
			case "multi-prompt-color":
				setTargetColor(&app.Setting.MultiPromptColor)
			case "single-text-color":
				setTargetColor(&app.Setting.SingleTextColor)
			case "multi-text-color":
				setTargetColor(&app.Setting.MultiTextColor)
			default:
				fmt.Printf("%vinvalid cmd!\n", config.ColorRed)
			}

		case line == "exit":
			fmt.Println("Goodbye!")
			return

		case line == "":
		default:
			chat.SendMessage(line, ctrl_c)
		}

		isWork = false
	}
}

func loadPromptStyle() {
	promptChar = app.Setting.InputPrompt
	singlePromptColor = config.Colors[app.Setting.SinglePromptColor]
	singleTextColor = config.Colors[app.Setting.SingleTextColor]
	multiPromptColor = config.Colors[app.Setting.MultiPromptColor]
	multiTextColor = config.Colors[app.Setting.MultiTextColor]
	if reader != nil {
		reader.SetPrompt(getPrompt())
	}
}

func getPrompt() (prompt string) {
	if multiline {
		prompt = fmt.Sprintf("%v%v%v", multiPromptColor, promptChar, multiTextColor)
	} else {
		prompt = fmt.Sprintf("%v%v%v", singlePromptColor, promptChar, singleTextColor)
	}
	return
}

func clearScreen() {
	os.Stdin.Write([]byte("\033[2J\033[0;0H"))

	var colors = []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
	myFigure := figure.NewColorFigure("Chat AI", "small", "blue", true)
	str := myFigure.String()

	rand.Seed(time.Now().UnixNano())
	banner := ""
	for _, v := range str {
		if !unicode.IsSpace(v) {
			banner = banner + fmt.Sprintf("%v%v", colors[rand.Intn(len(colors)-1)], string(v))
		} else {
			banner = banner + string(v)
		}
	}
	print(banner)
}

func toggleMultiline() {
	multiline = !multiline
	if multiline {
		reader.SetPrompt(getPrompt())
		println("\033[0mMulti-line input is enabled press \033[31mCtrl+D\033[0m to send\n")
	} else {
		reader.SetPrompt(getPrompt())
		println("\033[0mMulti-line input is disabled press \033[31mEnter\033[0m to send\n")
	}
}
