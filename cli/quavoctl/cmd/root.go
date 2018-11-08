package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//Global Flags
var (
	Key   string
	Value string
)

var (
	port    = ":4000"
	message string
)

func init() {

}

var rootCmd = &cobra.Command{
	Use:   "quavoctl [COMMANDS]",
	Short: "quavoctl is a command line interface to interact with quavo deamon",
	Long: `A simple command line based client to interact with a quavo key-value store.
		   Complete documentation is available at http://github/rugwirobaker/quavo`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

//Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
