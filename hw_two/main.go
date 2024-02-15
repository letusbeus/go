package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	result, err := CountLetters("Hello, world!")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func CountLetters(s string) (string, error) {
	var (
		ltrs    = make(map[rune]int)
		ordered []rune // для сохранения порядка встречи букв в строке
	)

	// Преобразуем строку в нижний регистр сразу, чтобы не делать это в цикле
	s = strings.ToLower(s)

	// Проходим по строке и учитываем только буквы
	for _, ltr := range s {
		if unicode.IsLetter(ltr) {
			if _, exists := ltrs[ltr]; !exists {
				ordered = append(ordered, ltr) // добавляем букву в порядке появления
			}
			ltrs[ltr]++
		}
	}

	if len(ltrs) == 0 {
		return "", fmt.Errorf("provided string does not contain letters")
	}

	// Вычисляем общее количество букв
	totalLetters := 0
	for _, count := range ltrs {
		totalLetters += count
	}

	sb := strings.Builder{}
	for _, ltr := range ordered { // используем сохраненный порядок букв
		count := ltrs[ltr]
		percent := float64(count) / float64(totalLetters)                 // вычисляем процент
		str := fmt.Sprintf("%s - %d %.1f\n", string(ltr), count, percent) // форматируем строку с округлением до одной десятой
		sb.WriteString(str)
	}

	return sb.String(), nil
}
