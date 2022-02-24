/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cobra-demo/demo"

	"github.com/spf13/cobra"
)

// demoCmd represents the demo command
var demoCmd = &cobra.Command{
	Use:   "demo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	demo.ArrayDemo()
	// },
}

var arrayCmd = &cobra.Command{
	Use:   "array",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		demo.MapDemo()
	},
}

var mapCmd = &cobra.Command{
	Use:   "map",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		demo.ArrayDemo()
	},
}

var classCmd = &cobra.Command{
	Use:   "class",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		demo.ClassDemo()
	},
}

func init() {
	demoCmd.AddCommand(mapCmd)
	demoCmd.AddCommand(arrayCmd)
	demoCmd.AddCommand(classCmd)
	rootCmd.AddCommand(demoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
