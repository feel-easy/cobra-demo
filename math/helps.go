package math

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	parseHandling int
)

type OpType int

const (
	ADD      OpType = 1
	MINUS    OpType = 2
	MULTIPLY OpType = 3
	DIVIDE   OpType = 4
)

type ErrorHandling int

const (
	ContinueOnParseError  ErrorHandling = 1 // 解析错误尝试继续处理
	ExitOnParseError      ErrorHandling = 2 // 解析错误程序停止
	PanicOnParseError     ErrorHandling = 3 // 解析错误 panic
	ReturnOnDividedByZero ErrorHandling = 4 // 除0返回
	PanicOnDividedByZero  ErrorHandling = 5 // 除0 painc
)

func calc(values []float64, opType OpType) float64 {
	var result float64
	if len(values) == 0 {
		return result
	}

	result = values[0]
	for i := 1; i < len(values); i++ {
		switch opType {
		case ADD:
			result += values[i]
		case MINUS:
			result -= values[i]
		case MULTIPLY:
			result *= values[i]
		case DIVIDE:
			if values[i] == 0 {
				switch ErrorHandling(dividedByZeroHanding) {
				case ReturnOnDividedByZero:
					return result
				case PanicOnDividedByZero:
					panic(errors.New("divided by 0"))
				}
			}
			result /= values[i]
		}
	}

	return result
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v", cmd.Name(), args, err)
	os.Exit(1)
}

func ConvertArgsToFloat64Slice(args []string, errorHandling ErrorHandling) []float64 {
	result := make([]float64, 0, len(args))
	for _, arg := range args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			switch errorHandling {
			case ExitOnParseError:
				fmt.Fprintf(os.Stderr, "invalid number: %s\n", arg)
				os.Exit(1)
			case PanicOnParseError:
				panic(err)
			}
		} else {
			result = append(result, value)
		}
	}

	return result
}
