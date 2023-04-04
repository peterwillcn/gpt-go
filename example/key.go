package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hanyuancheung/gpt-go"
)

func main() {
	fileName := "key.csv"
	fs, _ := os.Open(fileName)
	defer fs.Close()
	r := csv.NewReader(fs)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row[0])
		apiKey := row[0]

		ctx := context.Background()
		client := gpt.NewClient(apiKey)
		questionParam := "Hello"
		GetResponse(client, ctx, questionParam)
	}
}

// GetResponse get response from gpt3
func GetResponse(client gpt.Client, ctx context.Context, question string) {
	err := client.CompletionStreamWithEngine(ctx, &gpt.CompletionRequest{
		Model: gpt.TextDavinci003Engine,
		Prompt: []string{
			question,
		},
		MaxTokens:   3000,
		Temperature: 0,
	}, func(resp *gpt.CompletionResponse) {
		fmt.Print(resp.Choices[0].Text)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\n")
}
