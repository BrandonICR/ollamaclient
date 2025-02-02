package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/BrandonICR/ollamaclient"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(dir, "_example") {
		dir += "/_example"
	}
	f, err := os.OpenFile(dir+"/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	model, domain := loadArgs()

	client := ollamaclient.NewClient(&http.Client{}, domain)

	generate := ollamaclient.NewGenerateRequest(model, `You are a dream psychologist, you are 
		going to tell me the meaning of each dream that I'm goin to share with you in other messages.`).
		WithStream(false).Build()
	genRes, err := client.Generate(ctx, generate)
	if err != nil {
		return fmt.Errorf("generate: %w", err)
	}
	log.Println("generate - prompt:", generate.Prompt, " - response:", ollamaclient.CutThink(genRes.Response))

	message := ollamaclient.Message{
		Role:    "user",
		Content: "I dreamed that I threw myself off a building.",
	}
	messages := []ollamaclient.Message{message}
	chat := ollamaclient.NewChatRequest(model, messages).WithStream(false).Build()
	chatRes, err := client.Chat(ctx, chat)
	if err != nil {
		return fmt.Errorf("chat message 1: %w", err)
	}
	log.Println("chat_1 - content:", message.Content, " - response:", ollamaclient.CutThink(chatRes.Message.Content))

	message = ollamaclient.Message{
		Role:    "user",
		Content: "But in the midst of the fall I regretted it.",
	}
	messages = append(messages, chatRes.Message, message)
	chat = ollamaclient.NewChatRequest(model, messages).WithStream(false).Build()
	chatRes, err = client.Chat(ctx, chat)
	if err != nil {
		return fmt.Errorf("chat message 2: %w", err)
	}
	log.Println("chat_2 - content:", message.Content, " - response:", ollamaclient.CutThink(chatRes.Message.Content))

	return nil
}

func loadArgs() (model, domain string) {
	args := os.Args
	model = "deepseek-r1:14b"
	domain = "http://localhost:11434"
	if len(args) == 0 {
		return
	}
	for _, arg := range args {
		if strings.HasPrefix(arg, "model=") {
			model, _ = strings.CutPrefix(arg, "model=")
		}
		if strings.HasPrefix(arg, "domain=") {
			domain, _ = strings.CutPrefix(arg, "domain=")
		}
	}
	return
}
