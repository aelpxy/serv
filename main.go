package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	
	"github.com/spf13/cobra"
)

var (
	application_port string
	storage_folder   string
)

func main() {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Serve the application",
		Long:  "Serve the application",
		Run: func(cmd *cobra.Command, args []string) {
			
			if application_port == "" {
				log.Println("Port is missing")
			}

			log.Println("Starting...")

			g := http.NewServeMux()
			g.Handle("/", http.FileServer(http.Dir(storage_folder)))
			g.HandleFunc("/upload", uploadHandler)

			log.Printf("Listening on http://0.0.0.0:%v \n", application_port)
			log.Printf("Storing files on %v \n", storage_folder)

			if err := http.ListenAndServe(":"+application_port, g); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&application_port, "port", "p", "8080", "Port to expose webserver on.")
	cmd.Flags().StringVarP(&storage_folder, "folder", "f", "content", "Folder to store uploaded data on.")

	rootCmd := &cobra.Command{Use: "broccoli"}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = os.MkdirAll(storage_folder, os.ModePerm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dst, err := os.Create(fmt.Sprintf(storage_folder+"/%s", filepath.Base(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "success")
}
