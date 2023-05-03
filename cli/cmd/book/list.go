package book

import (
	"canonical/assessment/client"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all store books",
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
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
