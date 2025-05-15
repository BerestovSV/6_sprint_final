package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler возвращает HTML-страницу index.html по корневому эндпоинту "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Using IndexHandler")
	http.ServeFile(w, r, "../index.html")
}

// UploadHandler обрабатывает загрузку файла и конвертацию содержимого
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Парсинг формы
	err := r.ParseMultipartForm(10 << 20) // максимум 10MB
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
		return
	}

	// 2. Получение файла
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// 3. Чтение содержимого файла
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Не удалось прочитать файл", http.StatusInternalServerError)
		return
	}
	inputText := string(fileData)

	// 4. Конвертация через AutoConvert
	convertedText, err := service.AutoConvert(inputText)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при конвертации: %v", err), http.StatusInternalServerError)
		return
	}

	// 5. Создание локального файла
	timestamp := time.Now().UTC().Format("20060102_150405")
	ext := filepath.Ext(handler.Filename)
	if ext == "" {
		ext = ".txt"
	}
	outputFileName := fmt.Sprintf("converted_%s%s", timestamp, ext)

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		http.Error(w, "Не удалось создать файл для записи", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	// 6. Запись результата
	_, err = outputFile.WriteString(convertedText)
	if err != nil {
		http.Error(w, "Не удалось записать в файл", http.StatusInternalServerError)
		return
	}

	// 7. Ответ пользователю
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(convertedText))
}
