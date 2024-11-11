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
const systemPrompt string = "You are a software Developer assistant intended to gain insights about programming concepts particularly but not limited to functions in various programming languages and library. Your response is meant to be rendered on a terminal so markdown styled response is not needed. In cases where a manpage styled response is appropiate, generate response as such."

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
	ctx := context.Background()
	// get directly from enviromental variables
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("gemini_api_key")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash") // is gemini flash the best model in this case?
	model.SystemInstruction = genai.NewUserContent(genai.Text(systemPrompt))
	// change to os args
	resp, err := model.GenerateContent(ctx, genai.Text("chrome runtimeOnmessage with runtime sendmesssage"))
	if err != nil {
		log.Fatal(err)
	}
	printResponse(resp)
}
