# ChatAI

## Introduction:
The ChatGPT API CLI application is a command-line interface based tool designed to provide easy and quick access to the ChatGPT API. With a simple and intuitive command structure, this tool allows users to easily create, edit and delete chats, switch between different chats, and customize various settings as per their requirements.


## Installation

Then, set the environment variable `CHATGPT_API_KEY` in your system with your API key.


## Command Structure:
The following is a list of commands available in the ChatGPT API CLI application.

#### Usage:
- `help` This command displays the help information containing all the available commands.
- `chat`: This command is used to switch between different chats.
   - `chat info`: This sub-command provides the current chat's details.
   - `chat new`: This sub-command creates a new chat.
   - `chat edit`: This sub-command allows the user to edit the current chat.
   - `chat delete`: This sub-command deletes the selected chat.
- `setting`: This command is used to customize various settings as per requirements.
   - `setting multi-line`: This sub-command sets the default input mode to single-line or multi-line.
   - `setting quick`: This sub-command sets the quick answer chat.
   - `setting input-prompt`: This sub-command sets the input prompt prefix.
   - `setting single-prompt-color`: This sub-command sets the single-line prompt color.
   - `setting multi-prompt-color`: This sub-command sets the multi-line prompt color.
   - `setting single-text-color`: This sub-command sets the single-line input text color.
   - `setting multi-text-color`: This sub-command sets the multi-line input text color.
- `stop-cmd`: This command stops the execution of commands in the application. Users can press Ctrl + C to exit the application.
- `multi-line`: This command switches the input mode between single-line and multi-line.


## Readline Shortcut

`Meta`+`B` means press `Esc` and `n` separately.
Users can change that in terminal simulator(i.e. iTerm2) to `Alt`+`B`
Notice: `Meta`+`B` is equals with `Alt`+`B` in windows.

* Shortcut in normal mode

| Shortcut           | Comment                           |
| ------------------ | --------------------------------- |
| `Ctrl`+`A`         | Beginning of line                 |
| `Ctrl`+`B` / `←`   | Backward one character            |
| `Meta`+`B`         | Backward one word                 |
| `Ctrl`+`C`         | Send io.EOF                       |
| `Ctrl`+`D`         | Delete one character              |
| `Meta`+`D`         | Delete one word                   |
| `Ctrl`+`E`         | End of line                       |
| `Ctrl`+`F` / `→`   | Forward one character             |
| `Meta`+`F`         | Forward one word                  |
| `Ctrl`+`G`         | Cancel                            |
| `Ctrl`+`H`         | Delete previous character         |
| `Ctrl`+`I` / `Tab` | Command line completion           |
| `Ctrl`+`J`         | Line feed                         |
| `Ctrl`+`K`         | Cut text to the end of line       |
| `Ctrl`+`L`         | Clear screen                      |
| `Ctrl`+`M`         | Same as Enter key                 |
| `Ctrl`+`N` / `↓`   | Next line (in history)            |
| `Ctrl`+`P` / `↑`   | Prev line (in history)            |
| `Ctrl`+`R`         | Search backwards in history       |
| `Ctrl`+`S`         | Search forwards in history        |
| `Ctrl`+`T`         | Transpose characters              |
| `Meta`+`T`         | Transpose words (TODO)            |
| `Ctrl`+`U`         | Cut text to the beginning of line |
| `Ctrl`+`W`         | Cut previous word                 |
| `Backspace`        | Delete previous character         |
| `Meta`+`Backspace` | Cut previous word                 |
| `Enter`            | Line feed                         |


* Shortcut in Search Mode (`Ctrl`+`S` or `Ctrl`+`r` to enter this mode)

| Shortcut                | Comment                                 |
| ----------------------- | --------------------------------------- |
| `Ctrl`+`S`              | Search forwards in history              |
| `Ctrl`+`R`              | Search backwards in history             |
| `Ctrl`+`C` / `Ctrl`+`G` | Exit Search Mode and revert the history |
| `Backspace`             | Delete previous character               |
| Other                   | Exit Search Mode                        |

* Shortcut in Complete Select Mode (double `Tab` to enter this mode)

| Shortcut                | Comment                                  |
| ----------------------- | ---------------------------------------- |
| `Ctrl`+`F`              | Move Forward                             |
| `Ctrl`+`B`              | Move Backward                            |
| `Ctrl`+`N`              | Move to next line                        |
| `Ctrl`+`P`              | Move to previous line                    |
| `Ctrl`+`A`              | Move to the first candicate in current line |
| `Ctrl`+`E`              | Move to the last candicate in current line |
| `Tab` / `Enter`         | Use the word on cursor to complete       |
| `Ctrl`+`C` / `Ctrl`+`G` | Exit Complete Select Mode                |
| Other                   | Exit Complete Select Mode                |
