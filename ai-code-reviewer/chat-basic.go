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
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

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

		response, err := llms.GenerateFromSinglePrompt(ctx, llm, input)
		if err != nil {
			log.Fatal(err)
			continue
		}

		fmt.Printf("AI: %s\n", response)
	}
}
