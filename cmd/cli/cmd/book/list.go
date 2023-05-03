package book

import (
	"canonical/assessment/cmd/cli/utils"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all books",
	Run: func(cmd *cobra.Command, args []string) {
		apiClient := utils.CreateApiClient()

		books, err := apiClient.ListBooks()
		if err != nil {
			utils.PrintErrorJSON(fmt.Sprintf("failed to list books: %v/n", err))
		}

		json, err := json.MarshalIndent(books, "", "  ")
		if err != nil {
			utils.PrintErrorJSON(fmt.Sprintf("failed to parse response: %v/n", err))
		}

		utils.PrintJSON(string(json))
	},
}

func init() {

}
