package main

import (
	"fmt"
	"log"

	"baliance.com/gooxml/document"
)

func main() {
	// 打开模板文件
	doc, err := document.Open("report_tmpl.docx")
	if err != nil {
		log.Fatalf("无法打开模板: %s", err)
	}

	// 遍历所有表格
	tbl := doc.Tables()[0]

	for _, row := range tbl.Rows() {
		data := row.Cells()[1].Paragraphs()[0].Runs()[0].Text()
		var text string
		switch data {
		case "EventType":
			text = "森林防火"
		case "EventTime":
			text = "2023年10月28日"
		case "HighTime":
			text = "2023年10月28日, 09:00"
		case "Address":
			text = "浙江省杭州市"
		case "OvgTime":
			text = "1.2h"
		case "CountTime":
			text = "222.3"
		case "CaseCount":
			text = "3"
		default:
			fmt.Printf("未知数据: %s\n", data)
		}
		row.Cells()[1].Paragraphs()[0].Runs()[0].ClearContent()
		row.Cells()[1].Paragraphs()[0].Runs()[0].AddText(text)
	}

	// 保存为新文件
	err = doc.SaveToFile("output.docx")
	if err != nil {
		log.Fatalf("保存文件失败: %s", err)
	}

	fmt.Println("生成成功：output.docx")
}
