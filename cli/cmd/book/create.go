package book

import (
	"canonical/assessment/client"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("book create called")

		title, _ := cmd.Flags().GetString("title")

		c := client.NewApiClient("http://localhost:8080/api")
		book, err := c.CreateBook(client.Book{Title: title})
		if err != nil {
			fmt.Printf("failed to create book: %v/n", err)
		}
		json, err := json.MarshalIndent(book, "", "  ")
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
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringP("title", "t", "", "title")
	createCmd.MarkFlagRequired("title")
}
