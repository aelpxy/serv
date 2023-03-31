package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func UploadFile(storageFolder string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		errCh := make(chan error)
		go func() {
			err = os.MkdirAll(storageFolder, 0755)
			if err != nil {
				errCh <- err
				return
			}

			dstPath := filepath.Join(storageFolder, fileHeader.Filename)
			dst, err := os.Create(dstPath)
			if err != nil {
				errCh <- err
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				errCh <- err
				return
			}
			errCh <- nil
		}()

		if err = <-errCh; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully\n")
	}
}
