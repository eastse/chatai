package request

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"chatai/app"
)

var headerData = []byte("data: ")

type ChatRequest struct {
	Request  *http.Request
	Response *http.Response
	C        chan []byte
}

func Post(ctx context.Context, url string, data *bytes.Buffer) *ChatRequest {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, data)
	if err != nil {
		fmt.Printf("create request err: %v", err)
		os.Exit(0)
	}

	return &ChatRequest{
		Request: request,
		C:       nil,
	}
}

func PostStream(ctx context.Context, url string, data *bytes.Buffer) *ChatRequest {
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, data)
	if err != nil {
		fmt.Printf("create request err: %v", err)
		os.Exit(0)
	}

	return &ChatRequest{
		Request: request,
		C:       make(chan []byte, 100),
	}
}

func (c *ChatRequest) DoStream() {
	go func() {
		defer close(c.C)

		resp, err := http.DefaultClient.Do(c.Request)
		app.StopSpinner()

		if err != nil {
			if !errors.Is(err, context.Canceled) {
				fmt.Println(err)
			}
			return
		}
		c.Response = resp

		defer resp.Body.Close()

		if resp.StatusCode >= 300 {
			printResponse(resp)
			return
		}

		// 读取响应体
		reader := bufio.NewReader(resp.Body)
		for {
			line, e := reader.ReadBytes('\n')
			if e != nil {
				return
			}

			if bytes.Equal(line, []byte{'\n'}) || len(line) == 0 {
				continue
			}

			line = bytes.TrimPrefix(line, headerData)

			c.C <- line
		}
	}()
}

func (c *ChatRequest) Do() (body []byte, err error) {
	defer close(c.C)

	body = []byte{}
	resp, err := http.DefaultClient.Do(c.Request)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		err = errors.New(resp.Status)
		printResponse(resp)
		return
	}

	body, err = io.ReadAll(resp.Body)

	return
}

func printResponse(resp *http.Response) {
	var dump []byte
	dump, err := httputil.DumpResponse(resp, true)
	log.Printf("\n%s\n", string(dump))
	if err != nil {
		fmt.Printf("print response err: %v\n", err)
	}
}
