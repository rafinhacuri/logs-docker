package main

import (
	"fmt"
	"os"
	"strconv"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var telegramBot *telegram.BotAPI

func initTelegram() error {
	tokenTelegram := os.Getenv("TELEGRAM_TOKEN")

	var err error
	telegramBot, err = telegram.NewBotAPI(tokenTelegram)
	if err != nil {
		return err
	}

	return nil
}

func sendTelegramMessage(sistema, mensagem string) error {
	chatEnv := os.Getenv("TELEGRAM_CHAT")

	telegramChatID, err := strconv.ParseInt(chatEnv, 10, 64)
	if err != nil {
		return err
	}

	mensagemFormatada := fmt.Sprintf("Erro no container *%s*:\n```\n%s\n```", sistema, mensagem)

	msg := telegram.NewMessage(telegramChatID, mensagemFormatada)
	msg.ParseMode = "Markdown"

	if _, err := telegramBot.Send(msg); err != nil {
		return err
	}

	return nil
}
