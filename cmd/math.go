/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cobra-demo/math"
	"fmt"

	"github.com/spf13/cobra"
)

// mathCmd represents the math command
var mathCmd = &cobra.Command{
	Use:   "math",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("math called")
	},
}

func init() {
	mathCmd.AddCommand(math.AddCmd)
	mathCmd.AddCommand(math.DivideCmd)
	mathCmd.AddCommand(math.MinusCmd)
	mathCmd.AddCommand(math.MultiplyCmd)

	rootCmd.AddCommand(mathCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mathCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mathCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
