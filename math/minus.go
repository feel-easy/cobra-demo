package math

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var MinusCmd = &cobra.Command{
	Use:   "minus",
	Short: "Minus subcommand minus all passed args.",
	Run: func(cmd *cobra.Command, args []string) {
		values := ConvertArgsToFloat64Slice(args, ErrorHandling(parseHandling))
		result := calc(values, MINUS)
		fmt.Printf("%s = %.2f\n", strings.Join(args, "-"), result)
	},
}

func init() {
	// cmd.MainCmd.AddCommand(MinusCmd)
}
