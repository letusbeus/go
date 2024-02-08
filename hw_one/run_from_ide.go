package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Введите полный путь до файла:")

	var filePath string
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		log.Fatal("Ошибка при чтении ввода: ", err)
	}

	// Получаем базовое имя файла
	baseName := filepath.Base(filePath)

	// Получаем расширение файла
	fileExt := filepath.Ext(baseName)

	// Имя файла без расширения
	fileName := strings.TrimSuffix(baseName, fileExt)

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
