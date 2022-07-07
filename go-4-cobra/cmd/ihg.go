/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"getme/util"

	"github.com/spf13/cobra"
)

// ihgCmd represents the ihg command
var ihgCmd = &cobra.Command{
	Use:   "ihg",
	Short: "Shortcut to open IHG",
	Long: "Shortcut to open IHG",
	Run: func(cmd *cobra.Command, args []string) {
		util.ExecuteCMD(
			"code",
			"/Users/gabrielcalvo/dev/automation/analytics-help-tool", 
			"Shortcut called to open IHG Migration Tool",
			"Error trying to open IHG Migration Tool",
		)
	},
}

func init() {
	rootCmd.AddCommand(ihgCmd)
}
