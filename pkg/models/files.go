package models

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"math/rand"
	"time"
)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func generateRandomName(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}


func GenerateUniqueFilename(path, originalFilename string) (string, error) {
	// Получаем расширение файла
	ext := filepath.Ext(originalFilename)

	for i := 0; i < 10; i++ {
		// Генерируем случайную часть имени
		randomPart := generateRandomName(8)

		// Собираем новое имя файла с оригинальным расширением
		newFilename := randomPart + ext
		filePath := filepath.Join(path, newFilename)

		// Проверяем существование файла
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return newFilename, nil
		}
	}
	return "", errors.New("не удалось сгенерировать уникальное имя файла")
}

func SaveFile(src multipart.File, dstPath string) error {
	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}

func DeleteFileIfExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	}
	return nil
}
