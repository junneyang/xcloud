package main

import (
	"fmt"
	"os"

	"github.com/tealeg/xlsx"
)

func GetDeptInfo() (deptInfo map[string]string) {
	deptInfo = make(map[string]string)

	DeptInfo_File := os.Getenv("DeptInfo_File")
	DeptInfo_Sheet := os.Getenv("DeptInfo_Sheet")

	deptInfo_File, err := xlsx.OpenFile(DeptInfo_File)
	if err != nil {
		fmt.Println("DeptInfo_File 打开失败, 请确定 文件路径/文件名 是否正确...")
		return
	}
	deptInfo_Sheet := deptInfo_File.Sheet[DeptInfo_Sheet]
	for _, row := range deptInfo_Sheet.Rows {
		deptInfo[row.Cells[1].Value] = row.Cells[0].Value
	}
	return deptInfo
}
