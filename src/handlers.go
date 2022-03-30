package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) Status {
	status := Default

	if update.Message == nil {
		return status
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "help":
		msg.Text = "I understand: \n/noteList\n/status\n/noteAdd"
	case "noteList":
		msg.Text = ReadFileLineByLine("./Notes.txt")
	case "status":
		msg.Text = "I'm ok."
	case "noteAdd":
		msg.Text = update.Message.From.UserName
		status = AddingNote
	default:
		msg.Text = "I understand: \n/noteList\n/status\n/noteAdd"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return status
}
func HandleAddingNote(bot *tgbotapi.BotAPI, update tgbotapi.Update) Status {
	status := Default

	if update.Message == nil {
		return status
	}

	AppendToFile("./Notes.txt", update.Message.Text)

	return status
}
func ReadFileLineByLine(filename string) string {
	result := ""
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return result
}
func AppendToFile(file, text string) {
	f, err := os.OpenFile(file,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + text); err != nil {
		log.Println(err)
	}
}
