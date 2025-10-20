// can be used as 
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/memory"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	chatMem := memory.NewConversationBuffer()
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	llm, err := googleai.New(
		ctx,
		googleai.WithAPIKey(os.Getenv("GEMINI_API_KEY")),
		googleai.WithDefaultModel("gemini-2.5-flash"),
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Printf("You: ")
		input, _ := reader.ReadString('\n')

		if input == "quit" {
			break
		}

		messages, _ := chatMem.ChatHistory.Messages(ctx)

		var convo string
		for _, message := range messages {
			convo += message.GetContent() + "\n"
		}

		fullPrompt := convo + "Human: " + input + "\nAI: "

		response, err := llms.GenerateFromSinglePrompt(ctx, llm, fullPrompt)
		if err != nil {
			log.Fatal(err)
			continue
		}
		
		chatMem.ChatHistory.AddUserMessage(ctx, input)
		chatMem.ChatHistory.AddAIMessage(ctx, response)

		fmt.Println("Conversation History: " + convo)
		fmt.Printf("AI: %s\n", response)
	}
}
