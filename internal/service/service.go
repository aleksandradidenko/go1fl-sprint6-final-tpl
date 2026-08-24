package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)

	if data == "" {
		return "", fmt.Errorf("empty data")
	}

	if strings.ContainsAny(data, "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ") {
		return morse.ToMorse(data), nil
	}

	if strings.Trim(data, ".- ") == "" {
		return morse.ToText(data), nil
	}

	return "", fmt.Errorf("unknown data format")
}
