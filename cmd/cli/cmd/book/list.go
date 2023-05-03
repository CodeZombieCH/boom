package book

import (
	"canonical/assessment/client"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all books",
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := client.NewApiClient("http://localhost:8080/api")

		books, err := apiClient.ListBooks()
		if err != nil {
			fmt.Printf("failed to list books: %v/n", err)
		}

		json, err := json.MarshalIndent(books, "", "  ")
		if err != nil {
			fmt.Printf("failed to parse response: %v/n", err)
		}

		fmt.Println(string(json))
	},
}

func init() {

}
