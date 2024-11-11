package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/husainof/weather-telegram-bot/internal/bot/config"
)

// Send any text message to the bot after the bot has been started
var msgs []string
var timeout int
var url = "https://api.telegram.org/bot/"

func main() {
	cfg := config.GetConfig()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		// bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(cfg.Bot.Token, opts...)
	if err != nil {
		log.Fatal(err)
	}

	// b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handler)

	b.Start(ctx)

	var upd models.Update

	b.ProcessUpdate(ctx, &upd)
	b.SendMessage(ctx, &bot.SendMessageParams{})
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	if update.Message == nil {
		return
	}

	m, errSend := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
	if errSend != nil {
		fmt.Printf("error sending message: %v\n", errSend)
		return
	}

	time.Sleep(time.Second * 2)

	_, errEdit := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:    m.Chat.ID,
		MessageID: m.ID,
		Text:      "New Message!",
	})
	if errEdit != nil {
		fmt.Printf("error edit message: %v\n", errEdit)
		return
	}

	var upd models.Update

	b.ProcessUpdate(ctx, &upd)

}
