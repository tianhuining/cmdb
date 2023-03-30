package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	ver bool
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo-api 后端",
	Long:  "demo-api 后端API",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println("1.0")
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print demo-api version")
}