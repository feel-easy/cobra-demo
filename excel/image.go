package excel

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func ImageDemo() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// 关闭工作簿
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 插入图片
	if err := f.AddPicture("Sheet1", "A2", "image.png", ""); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的缩放比例
	if err := f.AddPicture("Sheet1", "D2", "image.jpg",
		`{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		fmt.Println(err)
	}
	// 在工作表中插入图片，并设置图片的打印属性
	if err := f.AddPicture("Sheet1", "H2", "image.gif", `{
			"x_offset": 15,
			"y_offset": 10,
			"print_obj": true,
			"lock_aspect_ratio": false,
			"locked": false
	}`); err != nil {
		fmt.Println(err)
	}
	// 保存工作簿
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
}
