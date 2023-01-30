package handlers

// WIP HANDLER

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// )

// func UploadFile(w http.ResponseWriter, r *http.Request) {
// 	file, fileHeader, err := r.FormFile("file")

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	defer file.Close()

// 	err = os.MkdirAll(storage_folder, os.ModePerm)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	dst, err := os.Create(fmt.Sprintf(storage_folder+"/%s", filepath.Base(fileHeader.Filename)))

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer dst.Close()

// 	_, err = io.Copy(dst, file)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprintf(w, "uploaded")
// }
