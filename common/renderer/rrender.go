package renderer

import (
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/charmbracelet/glamour"
	"github.com/chzyer/readline"
	"github.com/mattn/go-runewidth"
)

var (
	winWidth int                   // Window Broadband Total Characters
	r        *glamour.TermRenderer // markdown renderer
	content  string                // Receive data
	lastData bytes.Buffer          // The last rendering result
	nowData  bytes.Buffer          // current rendering result
)

func init() {
	winWidth = readline.GetScreenWidth()

	r, _ = glamour.NewTermRenderer(
		glamour.WithStyles(DarkStyleConfig),
		glamour.WithPreservedNewLines(),
		glamour.WithWordWrap(0),
	)
}

func RenderOncef(text string) {
	data, err := r.Render(text)
	if err != nil {
		fmt.Print(data)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}

func RenderStream(text string, isFirst bool) {
	if isFirst {
		content = ""
		winWidth = readline.GetScreenWidth()
		lastData.Reset()
		nowData.Reset()
	}
	content = content + text

	ansiData, _ := r.Render(content)
	nowData.Reset()
	nowData.WriteString(ansiData)

	var err error
	var nowRow, lastRow string
	isEOF := false
	isEqualLine := true
	backLine := 0
	for {
		lastRow, err = lastData.ReadString('\n')
		if err == io.EOF {
			isEOF = true
		}

		if isEqualLine {
			nowRow, _ = nowData.ReadString('\n')
		} else {
			line, _ := ComputeStrLine(lastRow)
			backLine += line
		}

		if isEqualLine && lastRow != nowRow {
			isEqualLine = false
			lLine, _ := ComputeStrLine(lastRow)

			if lLine >= 2 {
				backLine = lLine - 1
			}
		}

		if isEOF {
			break
		}
	}

	if backLine > 0 {
		for i := 0; i < backLine; i++ {
			fmt.Print("\033[2K\033[1A")
		}
		fmt.Printf("\033[2K\r%v", nowRow)
	} else {
		fmt.Printf("\033[2K\r%v", nowRow)
	}

	if nowData.Len() > 0 {
		for {
			line, err := nowData.ReadString('\n')
			fmt.Print(line)
			if err == io.EOF {
				break
			}
		}
	}

	lastData.Reset()
	lastData.WriteString(ansiData)
}

// 计算字符串在屏幕的实际占用列, 实际显示行数
func ComputeStrLine(data string) (line int, isNextCur bool) {
	str := RemoveANSI(data)
	count := 0
	for _, v := range str {
		width := runewidth.RuneWidth(v)
		count += width
		if count == winWidth {
			count = 0
			line++
		} else if count > winWidth {
			count = width
			line++
		}
	}

	if line == 0 && count == 0 {
		// 换行符
		return 1, true
	}

	if count == 0 {
		isNextCur = true
	} else {
		line++
	}

	if line == 0 {
		line = 1
	}

	return
}

func RemoveANSI(str string) string {
	re := regexp.MustCompile("\x1b\\[[0-9;]*[a-zA-Z]")
	return re.ReplaceAllString(str, "")
}

func RenderTextOnce(content string) string {
	s, err := r.Render(content)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return s
}
