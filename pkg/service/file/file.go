package file

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"math/rand"
	"time"
)

type FileServece interface {
	AddFile(path, name string, file multipart.File) (string, error)
	DeleteFileIfExists(path string) error
}

type filesService struct {
	Letters string
}

func NewFilesService() *filesService {
    return &filesService{
        Letters: "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
    }
}

func (s *filesService) AddFile(path, name string, file multipart.File) (string, error) {
	fileName, err := s.generateUniqueFilename(path, name)
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(path, fileName)

	err = s.saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	return fileName, nil
}


func (s *filesService) generateUniqueFilename(path, originalFilename string) (string, error) {
	// Получаем расширение файла
	ext := filepath.Ext(originalFilename)

	for i := 0; i < 10; i++ {
		// Генерируем случайную часть имени
		randomPart := s.generateRandomName(10)

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

func (s *filesService) generateRandomName(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = s.Letters[rand.Intn(len(s.Letters))]
	}
	return string(b)
}


func (s *filesService) saveFile(file multipart.File, filePath string) error {
	path, err := os.Create(filePath)
	if err != nil {
		return errors.New("не удалось записать файл")
	}
	defer path.Close()

	if _, err := io.Copy(path, file); err != nil {
		return errors.New("не удалось сохранить файл")
	}
	return nil
}

func (s *filesService) DeleteFileIfExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	} 
	return nil
}