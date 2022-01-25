package math

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var MultiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "Multiply subcommand multiply all passed args.",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calc(values, MULTIPLY)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "*"), result)
	},
}

func init() {
	// cmd.MainCmd.AddCommand(multiplyCmd)
}
