package main

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/aelpxy/broccoli/handlers"
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

			r := http.NewServeMux()
			r.Handle("/", http.FileServer(http.Dir(storage_folder)))
			r.HandleFunc("/upload", handlers.UploadFile(storage_folder))

			log.Printf("Listening on http://0.0.0.0:%v \n", application_port)
			log.Printf("Storing files on %v \n", storage_folder)

			if err := http.ListenAndServe(":"+application_port, r); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().StringVarP(&application_port, "port", "p", "8080", "Port to expose webserver on.")
	cmd.Flags().StringVarP(&storage_folder, "folder", "f", "uploads", "Folder to store uploaded data on.")

	rootCmd := &cobra.Command{Use: "broccoli"}
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
