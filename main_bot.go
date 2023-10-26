package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6945696492:AAG6MfMhvbGkZz-k-tOsYj6fMt5YDT-lyZE")
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			// Обработка команды
			command := update.Message.Command()
			switch command {
			//case "sorry":
			//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сколько извинений он тебе пообещал?))")
			//	_, err := bot.Send(msg)
			//	if err != nil {
			//		log.Println(err)
			//	}
			case "ascii_art":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пришли мне любую фотографию и свершиться чудо")
				bot.Send(msg)

			default:
				// Обработка неизвестных команд
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			}
		} else {

			if update.Message.Photo != nil && len(update.Message.Photo) > 0 {
				// Вызов функции для обработки фотографии
				pathToSave := "ascii_art.jpg"

				output, err := os.Create(pathToSave)
				if err != nil {
					log.Panic(err)
				}
				defer output.Close()

				photo := update.Message.Photo
				fileID := photo[len(photo)-1].FileID
				fileConfig := tgbotapi.FileConfig{FileID: fileID}
				file, err := bot.GetFile(fileConfig)
				if err != nil {
					log.Panic(err)
				}
				fileURL := file.Link(bot.Token)
				resp, err := http.Get(fileURL)
				if err != nil {
					log.Panic(err)
				}
				defer resp.Body.Close()

				_, err = io.Copy(output, resp.Body)
				if err != nil {
					log.Panic(err)
				}
				text := ascii_image_converter()
				ascii_image_converter_to_jpg()

				pho := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("ascii_art-ascii-art.png"))
				if _, err = bot.Send(pho); err != nil {
					log.Fatalln(err)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
				bot.Send(msg)

			} else {
				// В случае неизвестного сообщения
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестное сообщение")
				_, err := bot.Send(msg)
				if err != nil {
					log.Println(err)
				}
			}

		}
	}

}
