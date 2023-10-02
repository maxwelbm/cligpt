package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/cobra"
)

var gptcli = &cobra.Command{
	Use:   "cligpt",
	Short: "Chat with cligpt in console.",
	Run: func(cmd *cobra.Command, args []string) {
		var questionParam string

		if len(args) > 1 {
			questionParam = strings.Join(os.Args[1:], " ")
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("ask: ")

			scanner.Scan()

			question := scanner.Text()
			questionParam = validatedQuestion(question)
		}

		// DEBUG: question
		// log.Println("question: ", questionParam)

		response(client, ctx, questionParam)
	},
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
