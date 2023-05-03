package cmd

import (
	"os"

	"canonical/assessment/cmd/cli/cmd/book"
	"canonical/assessment/cmd/cli/cmd/collection"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boom",
	Short: "boom - a simple book management system",
	Long: "" +
		"          _ ._  _ , _ ._\n" +
		"        (_ ' ( `  )_  .__)\n" +
		"      ( (  (    )   `)  ) _)\n" +
		"     (__ (_   (_ . _) _) ,__)\n" +
		"         `~~`\\ ' . /`~~`\n" +
		"              ;   ;\n" +
		"              /   \\\n" +
		"_____________/_ __ \\_____________\n" +
		"         *** b o o m ***\n" +
		"‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾‾\n" +
		`(from *Boo*k *M*anagement System) is a simple book management system.
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(collection.CollectionCmd)
	rootCmd.AddCommand(book.BookCmd)
}
