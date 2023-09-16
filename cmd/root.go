/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var (
	ctx     = context.Background()
	apiKey  = os.Getenv("API_KEY")
	client  = gpt3.NewClient(apiKey)
	rootCmd = &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command, args []string) {
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit {
				fmt.Print("Digite sua pergunta (quit para sair):")

				if !scanner.Scan() {
					break
				}

				question := scanner.Text()
				questionParam := validatedQuestion(question)
				switch questionParam {
				case "quit":
					quit = true
				case "":
					continue

				default:
					response(client, ctx, questionParam)
				}
			}
		},
	}
)

func Execute() {
	if apiKey == "" {
		log.Println("Environment variable not exported, export to use 'cligpt' with this example: 'export API_KEY=<GITHUB_TOKEN>'")
		os.Exit(1)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func response(client gpt3.Client, ctx context.Context, quesiton string) {
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			quesiton,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(13)
	}
	fmt.Printf("\n")
}

func validatedQuestion(question string) string {
	quest := strings.Trim(question, " ")
	keywords := []string{"", "loop", "break", "continue", "clear", "cls", "exit", "block"}
	for _, x := range keywords {
		if quest == x {
			return ""
		}
	}
	return quest
}
