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

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../Sprint_6_final/index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка при парсинге формы", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Не удалось прочитать файл", http.StatusInternalServerError)
		return
	}
	inputText := string(fileData)

	convertedText, err := service.AutoConvert(inputText)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка при конвертации: %v", err), http.StatusInternalServerError)
		return
	}

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

	_, err = outputFile.WriteString(inputText + "\n" + convertedText)
	if err != nil {
		http.Error(w, "Не удалось записать в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(inputText + "<br>" + convertedText))

}
