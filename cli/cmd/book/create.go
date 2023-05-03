package book

import (
	"canonical/assessment/client"
	"encoding/json"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new book",
	Run: func(cmd *cobra.Command, args []string) {
		parsedBook, err := parseBook(cmd.Flags())

		apiClient := client.NewApiClient("http://localhost:8080/api")

		book, err := apiClient.CreateBook(parsedBook)
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

	createCmd.Flags().StringP("author", "a", "", "Author")
	createCmd.Flags().StringP("publication-date", "p", "", "Publication Date")
	createCmd.Flags().StringP("edition", "e", "", "Edition")
	createCmd.Flags().StringP("description", "d", "", "Description")
	createCmd.Flags().StringP("genre", "g", "", "Genre")
}

func parseBook(flags *pflag.FlagSet) (*client.Book, error) {
	title, _ := flags.GetString("title")
	author, _ := flags.GetString("author")
	publicationDateRaw, _ := flags.GetString("publication-date")
	edition, _ := flags.GetString("edition")
	description, _ := flags.GetString("description")
	genre, _ := flags.GetString("genre")

	publicationDate, err := time.Parse("2006-01-02", publicationDateRaw)
	if err != nil {
		return nil, err
	}

	return &client.Book{
		Title:           title,
		Author:          author,
		PublicationDate: publicationDate,
		Edition:         edition,
		Description:     description,
		Genre:           genre,
	}, nil
}
