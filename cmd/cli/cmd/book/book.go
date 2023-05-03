package book

import (
	"github.com/spf13/cobra"
)

var BookCmd = &cobra.Command{
	Use:   "book",
	Short: "manages books",
}

func init() {
	BookCmd.AddCommand(createCmd)
	BookCmd.AddCommand(listCmd)
}
