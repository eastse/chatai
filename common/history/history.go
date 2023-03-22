package history

import (
	"fmt"
	"os"

	"chatai/app"
	"chatai/common/utils"
)

type History struct {
	file *os.File
}

func New(file string) *History {
	h := &History{}
	if !utils.PathExists(app.HISTORY_DIR) {
		err := os.MkdirAll(app.HISTORY_DIR, 0755)
		if err != nil {
			fmt.Printf("create history directory err: %v\n", err)
		}
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	h.file = f

	return h
}

func (c *History) Write(text string) {
	_, err := c.file.WriteString(text)
	if err != nil {
		fmt.Printf("write history err: %v\n", err)
	}
}

func (c *History) Close() {
	c.file.Close()
}
