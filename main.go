package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/upload", uploadImageHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static/")))

	http.Handle("/", r)
	log.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

func uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to read uploaded file", http.StatusBadRequest)
		return
	}

	defer file.Close()

	newFile, err := os.Create("./uploads/" + header.Filename)
	if err != nil {
		http.Error(w, "Unable to create a file on local storage", http.StatusInternalServerError)
		return
	}

	defer newFile.Close()

	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Unable to copy the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Image upload sukes")

}
