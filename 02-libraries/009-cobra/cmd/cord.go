/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"getme/util"

	"github.com/spf13/cobra"
)

// cordCmd represents the cord command
var cordCmd = &cobra.Command{
	Use:   "cord",
	Short: "Shortcut to open CORD dev env",
	Long: "Shortcut to open CORD dev env",
	Aliases: []string{"cord-tools", "cordtools"},
	Run: func(cmd *cobra.Command, args []string) {
		util.ExecuteCMD(
			"code",
			"/Users/gabrielcalvo/go/src/github.com/cord-tools", 
			"Shortcut called to open CORD (Dev env)",
			"Error trying to open CORD",
		)

		util.ExecuteCMD(
			"open",
			"-a Docker",
			"Starting Docker",
			"Error trying to start Docker",
		)
	},
}

func init() {
	rootCmd.AddCommand(cordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
