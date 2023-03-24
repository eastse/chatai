package cmd

import "github.com/chzyer/readline"

var completer = readline.NewPrefixCompleter(
	readline.PcItem("help"),
	readline.PcItem("chat",
		readline.PcItem("info"),
		readline.PcItem("new"),
		readline.PcItem("edit"),
		readline.PcItem("clear"),
		readline.PcItem("delete"),
	),
	readline.PcItem("setting",
		readline.PcItem("multi-line"),
		readline.PcItem("quick"),
		readline.PcItem("input-prompt"),
		readline.PcItem("single-prompt-color"),
		readline.PcItem("multi-prompt-color"),
		readline.PcItem("single-text-color"),
		readline.PcItem("multi-text-color"),
	),
	readline.PcItem("stop-cmd"),
	readline.PcItem("multi-line"),
)

var helpinfo = `
commands:

    help                    Show help information
    chat                    Switch chat
    ├── info                Show current chat information
    ├── new                 Create a new chat
    ├── edit                Edit the current chat
    ├── clear               Clear the current chat review context
    ├── delete              Delete the selected chat
    setting
    ├── multi-line          Set default single-line or multi-line input mode
    ├── quick               Set quick answer chat
    ├── input-prompt        Set input prompt prefix
    ├── single-prompt-color Set single-line prompt color
    ├── multi-prompt-color  Set multi-line prompt color
    ├── single-text-color   Set single-line input text color
    ├── multi-text-color    Set multi-line input text color
    stop-cmd                Stop receiving commands. Press Ctrl+C to exit
    multi-line              Switch single-line and multi-line input mode

Shortcut:

  Meta + B  means press  Esc  and  n  separately. Users can change that in
  terminal simulator(i.e. iTerm2) to  Alt + B Notice:  Meta + B  is equals
  with  Alt + B  in windows.

  • Shortcut in normal mode

        SHORTCUT     |            COMMENT
  -------------------+---------------------------------
    Ctrl + A         | Beginning of line
    Ctrl + B  /  ←   | Backward one character
    Meta + B         | Backward one word
    Ctrl + C         | Send io.EOF
    Ctrl + D         | Delete one character
    Meta + D         | Delete one word
    Ctrl + E         | End of line
    Ctrl + F  /  →   | Forward one character
    Meta + F         | Forward one word
    Ctrl + G         | Cancel
    Ctrl + H         | Delete previous character
    Ctrl + I  /  Tab | Command line completion
    Ctrl + J         | Line feed
    Ctrl + K         | Cut text to the end of line
    Ctrl + L         | Clear screen
    Ctrl + M         | Same as Enter key
    Ctrl + N  /  ↓   | Next line (in history)
    Ctrl + P  /  ↑   | Prev line (in history)
    Ctrl + R         | Search backwards in history
    Ctrl + S         | Search forwards in history
    Ctrl + T         | Transpose characters
    Meta + T         | Transpose words (TODO)
    Ctrl + U         | Cut text to the beginning of line
    Ctrl + W         | Cut previous word
    Backspace        | Delete previous character
    Meta + Backspace | Cut previous word
    Enter            | Line feed

  • Shortcut in Search Mode ( Ctrl + S  or  Ctrl + r  to enter this mode)

          SHORTCUT        |            COMMENT
  ------------------------+---------------------------------
    Ctrl + S              | Search forwards in history
    Ctrl + R              | Search backwards in history
    Ctrl + C  /  Ctrl + G | Exit Search Mode and revert the history
    Backspace             | Delete previous character
    Other                 | Exit Search Mode

  • Shortcut in Complete Select Mode (double  Tab  to enter this mode)

          SHORTCUT        |            COMMENT
  ------------------------+---------------------------------
    Ctrl + F              | Move Forward
    Ctrl + B              | Move Backward
    Ctrl + N              | Move to next line
    Ctrl + P              | Move to previous line
    Ctrl + A              | Move to the first candicate in current line
    Ctrl + E              | Move to the last candicate in current line
    Tab  /  Enter         | Use the word on cursor to complete
    Ctrl + C  /  Ctrl + G | Exit Complete Select Mode
    Other                 | Exit Complete Select Mode

`
