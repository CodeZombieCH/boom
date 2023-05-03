package book

import (
	"canonical/assessment/client"
	"canonical/assessment/cmd/cli/utils"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new book",
	Run: func(cmd *cobra.Command, args []string) {
		parsedBook, err := parseBook(cmd.Flags())
		if err != nil {
			utils.PrintErrorJSON(fmt.Sprintf("failed to parse book: %v", err))
			os.Exit(1)
		}

		apiClient := utils.CreateApiClient()
		book, err := apiClient.CreateBook(parsedBook)
		if err != nil {
			utils.PrintErrorJSON(fmt.Sprintf("failed to create book: %v", err))
			os.Exit(1)
		}

		utils.PrintJSON(book)
	},
}

const (
	BookCmdFlagTitle           string = "title"
	BookCmdFlagAuthor          string = "author"
	BookCmdFlagPublicationDate string = "publication-date"
	BookCmdFlagEdition         string = "edition"
	BookCmdFlagDescription     string = "description"
	BookCmdFlagGenre           string = "genre"
)

func init() {
	createCmd.Flags().StringP(BookCmdFlagTitle, "t", "", "title")
	createCmd.MarkFlagRequired(BookCmdFlagTitle)

	createCmd.Flags().StringP(BookCmdFlagAuthor, "a", "", "Author")
	createCmd.Flags().StringP(BookCmdFlagPublicationDate, "p", "", "Publication Date")
	createCmd.Flags().StringP(BookCmdFlagEdition, "e", "", "Edition")
	createCmd.Flags().StringP(BookCmdFlagDescription, "d", "", "Description")
	createCmd.Flags().StringP(BookCmdFlagGenre, "g", "", "Genre")
}

func parseBook(flags *pflag.FlagSet) (*client.Book, error) {
	var book = &client.Book{}

	titleRaw, _ := flags.GetString(BookCmdFlagTitle)
	book.Title = titleRaw

	if author, _ := flags.GetString(BookCmdFlagAuthor); len(author) > 0 {
		book.Title = titleRaw
	}

	if publicationDateRaw, _ := flags.GetString(BookCmdFlagPublicationDate); len(publicationDateRaw) > 0 {
		date, err := time.Parse("2006-01-02", publicationDateRaw)
		if err != nil {
			return nil, err
		}
		book.PublicationDate = &date
	}

	if edition, _ := flags.GetString(BookCmdFlagEdition); len(edition) > 0 {
		book.Edition = &edition
	}

	if description, _ := flags.GetString(BookCmdFlagDescription); len(description) > 0 {
		book.Description = &description
	}

	if genre, _ := flags.GetString(BookCmdFlagGenre); len(genre) > 0 {
		book.Genre = &genre
	}

	return book, nil
}
