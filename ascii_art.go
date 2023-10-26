package main

import (
	"log"
	"os"
	"os/exec"
)

func ascii_image_converter() string {
	scriptName := "ascii-image-converter"

	// Аргументы, которые вы хотите передать вашему скрипту
	args := []string{"ascii_art.jpg", "-d", "25,20", "-b", "--save-txt", "."}

	// Создаем команду для запуска скрипта
	cmd := exec.Command(scriptName, args...)

	// Устанавливаем стандартные потоки ввода/вывода/ошибок
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Запускаем скрипт
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	// Открываем файл с текстом
	file, err := os.ReadFile("ascii_art-ascii-art.txt")
	if err != nil {
		log.Panic(err)
	}

	text := string(file)

	return text

}

func ascii_image_converter_to_jpg() {
	scriptName := "ascii-image-converter"

	// Аргументы, которые вы хотите передать вашему скрипту
	args := []string{"ascii_art.jpg", "-C", "-d", "60,30", "--save-img", "."}

	// Создаем команду для запуска скрипта
	cmd := exec.Command(scriptName, args...)

	// Устанавливаем стандартные потоки ввода/вывода/ошибок
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Запускаем скрипт
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
