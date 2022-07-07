/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"getme/util"

	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "shc",
	Short: "Shortcut to open SHC",
	Long: "Shortcut to open SHC",
	Run: func(cmd *cobra.Command, args []string) {
		util.ExecuteCMD(
			"code",
			"/Users/gabrielcalvo/dev/clients/shc-automation", 
			"Shortcut called to open SHC",
			"Error trying to open SHC",
		)
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
