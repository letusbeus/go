package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePath := os.Args[1]

	// Получаем базовое имя файла
	baseName := filepath.Base(filePath)

	// Получаем расширение файла
	fileExt := filepath.Ext(baseName)

	// Имя файла без расширения
	fileName := strings.TrimSuffix(baseName, fileExt)

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
