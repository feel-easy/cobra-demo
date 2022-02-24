/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cobra-demo/excel"

	"github.com/spf13/cobra"
)

// excelCmd represents the excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "A brief description of your command",
	Long:  ``,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	f := excelize.NewFile()
	// 	// 创建一个工作表
	// 	index := f.NewSheet("Sheet2")
	// 	// 设置单元格的值
	// 	f.SetCellValue("Sheet2", "A2", "Hello world.")
	// 	f.SetCellValue("Sheet1", "B2", 100)
	// 	// 设置工作簿的默认工作表
	// 	f.SetActiveSheet(index)
	// 	// 根据指定路径保存文件
	// 	if err := f.SaveAs("Book1.xlsx"); err != nil {
	// 		fmt.Println(err)
	// 	}
	// },
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		excel.CreateExcel()
	},
}
var imgCmd = &cobra.Command{
	Use:   "img",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		excel.ImageDemo()
	},
}

var readCmd = &cobra.Command{
	Use:   "img",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		excel.ReadExcel()
	},
}

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		excel.TableExcel()
	},
}

func init() {
	excelCmd.AddCommand(createCmd)
	excelCmd.AddCommand(imgCmd)
	excelCmd.AddCommand(readCmd)
	excelCmd.AddCommand(tableCmd)
	rootCmd.AddCommand(excelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// excelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// excelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
