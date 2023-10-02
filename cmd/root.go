/*
copyright © 2023 Maxwel Mazur maxwelbm@hotmail.com
*/
package cmd

import (
	"context"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
)

var (
	ctx    = context.Background()
	apiKey = os.Getenv("API_KEY")
	client = gpt3.NewClient(apiKey)
)

func Execute() {
	if apiKey == "" {
		log.Println("Environment variable not exported, export to use 'cligpt' with this example: 'export API_KEY=<GITHUB_TOKEN>'")
		os.Exit(1)
	}

	err := gptcli.Execute()
	if err != nil {
		os.Exit(1)
	}
}
