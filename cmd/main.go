package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"os"
	"strings"
	"time"
)

func main() {
	var sk string
	flag.StringVar(&sk, "s", "", "sk")
	flag.Parse()
	fmt.Println("正在初始化...")
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
			Model:       gogpt.GPT3TextDavinci003,
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