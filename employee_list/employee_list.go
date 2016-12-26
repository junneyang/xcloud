package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {
	PrintLogo()
	// 1. GET BASIC INFO
	IDCard_File := os.Getenv("IDCard_File")
	IDCard_Sheet := os.Getenv("IDCard_Sheet")
	TEMPLATE_File := os.Getenv("TEMPLATE_File")
	TEMPLATE_Sheet := os.Getenv("TEMPLATE_Sheet")
	Position_No := os.Getenv("Position_No")

	idCard_File, err := xlsx.OpenFile(IDCard_File)
	if err != nil {
		fmt.Println("IDCard_File 打开失败, 请确定 文件路径/文件名 是否正确...")
		return
	}
	idCard_Sheet := idCard_File.Sheet[IDCard_Sheet]

	template_File, err := xlsx.OpenFile(TEMPLATE_File)
	if err != nil {
		fmt.Println("TEMPLATE_File 打开失败, 请确定 文件路径/文件名 是否正确...")
		return
	}
	template_Sheet := template_File.Sheet[TEMPLATE_Sheet]
	deptInfo := GetDeptInfo()
	idInfo := GetIDInfo()

	// 2. GENEGRATE TABLE SHEET
	for i, row := range idCard_Sheet.Rows {
		if i == 0 || i == 1 {
			continue
		}
		row_new := template_Sheet.AddRow()
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue(row.Cells[1].Value)
		row_new.AddCell()
		row_new.AddCell()
		tmp_dept_no, _ := strconv.ParseFloat(deptInfo[row.Cells[12].Value], 64)
		row_new.AddCell().SetFloat(tmp_dept_no)
		tmp_position_no, _ := strconv.ParseFloat(Position_No, 64)
		row_new.AddCell().SetFloat(tmp_position_no)
		if row.Cells[10].Value == "长期工" {
			row_new.AddCell().SetValue("储备岗")
			row_new.AddCell().SetValue("长期工")
		} else if row.Cells[10].Value == "小时工" {
			row_new.AddCell().SetValue("辅助储备岗")
			row_new.AddCell().SetValue("业务外包工")
		}
		start_time := time.Now().Add(time.Hour * 24)
		row_new.AddCell().SetDate(start_time)
		channel := row.Cells[11].Value
		if channel == "内部推荐" || channel == "回聘" || channel == "现场应聘" || channel == "网络应聘" || channel == "人才市场" {
			row_new.AddCell().SetValue(channel)
		} else {
			row_new.AddCell().SetValue("劳务机构")
		}
		if channel == "内部推荐" {
			row_new.AddCell().SetValue("Y")
			employee_no := row.Cells[14].Value
			employee_name := row.Cells[13].Value
			if strings.HasPrefix(employee_no, "900") {
				row_new.AddCell().SetValue("Y")
			} else {
				row_new.AddCell().SetValue("")
			}
			row_new.AddCell().SetValue(employee_no)
			row_new.AddCell().SetValue(employee_name)
		} else {
			row_new.AddCell().SetValue("N")
			row_new.AddCell().SetValue("")
			row_new.AddCell().SetValue("")
			row_new.AddCell().SetValue("")
		}
		if channel == "回聘" {
			row_new.AddCell().SetValue("Y")
		} else {
			row_new.AddCell().SetValue("N")
		}
		row_new.AddCell().SetValue("N")
		row_new.AddCell().SetDate(start_time)
		row_new.AddCell().SetDate(start_time)
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("武汉")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")

		addr := row.Cells[5].Value
		addr6 := string([]rune(addr)[0:6])
		row_new.AddCell().SetValue(addr6)
		row_new.AddCell().SetValue(addr6)
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("身份证")
		idno := row.Cells[0].Value
		row_new.AddCell().SetValue(idno)
		row_new.AddCell().SetValue(addr)

		start_end_day := row.Cells[6].Value
		start_day := strings.Split(start_end_day, "-")[0]
		end_day := strings.Split(start_end_day, "-")[1]
		tmp_start_day, _ := time.Parse("2006.01.02", start_day)
		tmp_end_day, _ := time.Parse("2006.01.02", end_day)
		row_new.AddCell().SetDate(tmp_start_day)
		row_new.AddCell().SetDate(tmp_end_day)
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("")
		row_new.AddCell().SetValue("中国")

		gender := row.Cells[2].Value
		row_new.AddCell().SetValue(gender)
		row_new.AddCell().SetValue("")
		nationality := row.Cells[3].Value
		row_new.AddCell().SetValue(nationality + "族")

		for i := 0; i < 26; i++ {
			row_new.AddCell().SetValue("")
		}

		if channel == "内部推荐" || channel == "现场应聘" || channel == "人才市场" {
			row_new.AddCell().SetValue("")
		} else if channel == "回聘" {
			row_new.AddCell().SetValue(row.Cells[16].Value)
		} else if channel == "网络应聘" {
			row_new.AddCell().SetValue(row.Cells[15].Value)
		} else {
			row_new.AddCell().SetValue(row.Cells[11].Value)
		}

		if row.Cells[12].Value == "CFC" && row.Cells[10].Value == "长期工" {
			row_new.Cells[2].SetValue("CN34")
		}

		// 3. MRAK
		v, _ := strconv.Atoi(row.Cells[4].Value)
		t, _ := time.Parse("2006-01-02", "1900-01-00")
		birth := t.AddDate(0, 0, v)
		if time.Now().Sub(birth.AddDate(18, 0, 0)).Seconds() < 0 {
			style := template_Sheet.Rows[1].Cells[2].GetStyle()
			row_new.Cells[1].SetStyle(style)
		}

		if tmp_end_day.Sub(time.Now()).Seconds() < 0 {
			style := template_Sheet.Rows[1].Cells[3].GetStyle()
			row_new.Cells[1].SetStyle(style)
		} else if tmp_end_day.Sub(time.Now().AddDate(0, 1, 0)).Seconds() < 0 {
			style := template_Sheet.Rows[1].Cells[4].GetStyle()
			row_new.Cells[1].SetStyle(style)
		}

		if _, ok := idInfo[idno]; ok {
			//			style := xlsx.NewStyle()
			//			fill := xlsx.NewFill("solid", "#00FF0000", "#FF000000")
			//			style.Fill = *fill
			//			style.ApplyFill = true
			//			row_new.Cells[1].SetStyle(style)
			style := template_Sheet.Rows[1].Cells[1].GetStyle()
			row_new.Cells[1].SetStyle(style)
		}
	}

	// 4. SAVE THE SHEET FILE
	//	now := time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")
	now := time.Now().Add(time.Hour * 24).Format("20060102")
	fileName := "入职人员信息表-" + now + ".xlsx"
	err = template_File.Save(fileName)
	if err != nil {
		fmt.Println("TEMPLATE_File 保存失败...", err)
		return
	} else {
		log := fmt.Sprint("处理完毕, 目标文件已经保存到-->>>>> ", fileName, "\n")
		fmt.Println(log)
	}

}
