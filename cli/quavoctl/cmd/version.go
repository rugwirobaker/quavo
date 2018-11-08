package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//VERSION  is the current command version
const VERSION = "v0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number of quavoctl.",
	Long:  `All software has versions. This is quavoclt's.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("quavo command-line interface %s\n", VERSION)
	},
}
