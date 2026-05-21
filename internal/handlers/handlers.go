package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}

// Обработчик для создания новой заметки.
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Открываем файл по тегу из HTML-файла, с именем "myFile":
	//*Вместо _ далее была переменная handler
	file, _, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	//Служебные данные:
	//fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename)
	//fmt.Fprintf(w, "handler.Header %#v\n", handler.Header)
	//fmt.Fprintln(w, r.MultipartForm)
	//fmt.Fprintln(w, handler)

	//Получение полезного внутреннего содержимого файла ("io.ReadAll(file)")
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	converted, err := service.TextType(string(data))
	if err != nil {
		log.Println(err)
	}

	//Введена переменная, хранящая имя файла, поскольку из оригинального файла из формы не достать его
	//расширение: fmt.Fprintf(w, "handler.Filename %v\n", handler.Filename) даёт имя без расширения,
	//из-за чего использование filepath.Ext() бессмысленно.
	newFileNameString := time.Now().UTC().String() + ".txt"

	fileName := filepath.Join("../", newFileNameString)
	//fmt.Fprintln(w, fileName)

	result, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// отложенное закрытие файла
	defer result.Close()
	_, err = result.WriteString(converted)

	fmt.Fprintln(w, converted)
}
