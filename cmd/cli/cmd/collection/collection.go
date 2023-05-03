package collection

import (
	"github.com/spf13/cobra"
)

var CollectionCmd = &cobra.Command{
	Use:   "collection",
	Short: "Manages collections of books",
}

func init() {
	CollectionCmd.AddCommand(createCmd)
	CollectionCmd.AddCommand(listCmd)
}
