package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// configure the generative ai system first
const systemInstructionPrompt string = "You are a software Developer assistant intended to gain insights about programming concepts particularly but not limited to functions in various programming languages and library. Your response is meant to be rendered on a terminal and piped into a CLI that renders maarkdown(markdown styled response are greatly appreciated as they will be properly parsed). In cases where a manpage styled response is appropiate, generate response as such."
const modelName = "gemini-2.5-flash-preview-04-17"

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Print(part)
			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("user prompt must be wrapped as a single string")
	}
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("gemini_api_key")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel(modelName)
	model.SystemInstruction = genai.NewUserContent(genai.Text(systemInstructionPrompt))
	// change to os args
	resp, err := model.GenerateContent(ctx, genai.Text(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	printResponse(resp)
}
