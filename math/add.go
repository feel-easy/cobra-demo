package math

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var ()

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add subcommand add all passed args.",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calc(values, ADD)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "+"), result)
	},
}

func init() {
	// cmd.MainCmd.AddCommand(addCmd)
	// math.MathCmd.AddCommand(addCmd)
	AddCmd.Flags().BoolP("hehe", "t", false, "这就是一个flag")
}
