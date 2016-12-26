package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

func GetIDInfo() (idInfo map[string]string) {
	idInfo = make(map[string]string)

	Dimission_ID_File := os.Getenv("Dimission_ID_File")
	Dimission_ID_Sheet := os.Getenv("Dimission_ID_Sheet")

	dimission_ID_File, err := xlsx.OpenFile(Dimission_ID_File)
	if err != nil {
		fmt.Println("Dimission_ID_File 打开失败, 请确定 文件路径/文件名 是否正确...")
		return
	}
	dimission_ID_Sheet := dimission_ID_File.Sheet[Dimission_ID_Sheet]
	for _, row := range dimission_ID_Sheet.Rows {
		idInfo[row.Cells[1].Value] = row.Cells[0].Value
	}
	return idInfo
}
