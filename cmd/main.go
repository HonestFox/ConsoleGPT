package main

import (
	"bufio"
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"os"
	"strings"
	"time"
)

func main() {
	//if time.Date(2023, 3, 30, 0, 0, 0, 0, time.Local).Before(time.Now()) {
	//	fmt.Println("过期了!!!")
	//	time.Sleep(10 * time.Second)
	//}

	fmt.Println("正在初始化...")

	//var sk = "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"

	var sk string
	fmt.Println("请输入sk:")
	for {
		reader := bufio.NewReader(os.Stdin)
		sk, _ = reader.ReadString('\n')
		sk = strings.TrimSpace(sk)
		if len(sk) > 0 {
			break
		}
	}
	if len(sk) == 0 {
		panic("s参数不能为空")
	}

	c := gogpt.NewClient(sk)
	ctx := context.Background()
	fmt.Println("初始化完成")

	ts := time.Now().GoString()

	for {
		var input string
		fmt.Println("请输入:")
		reader := bufio.NewReader(os.Stdin)
		for {
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if len(input) > 0 {
				break
			}
		}
		req := gogpt.CompletionRequest{
			Model:       "gpt-3.5-turbo",
			MaxTokens:   3500,
			Prompt:      input,
			Temperature: 0.9,
			User:        ts,
		}
		resp, err := c.CreateCompletion(ctx, req)
		if err != nil {
			fmt.Println("异常:", err.Error())
			continue
		}
		fmt.Println(resp.Choices[0].Text)
		fmt.Println("-----------")
	}
}
