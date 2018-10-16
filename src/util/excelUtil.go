package util

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
)

func Parse(filePath, sheetName string, limit int, hasHead bool) ([]string, [][]string, error) {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("文件读取失败;err=%s", err)
	}
	headList := make([]string, 0)
	excelDataRowsList := make([][]string, 0)
	for _, sheet := range xlFile.Sheets {
		if !strings.EqualFold(sheet.Name, sheetName) {
			continue
		}
		for i, row := range sheet.Rows {
			if limit > 0 && i > limit {
				break
			}
			exceldataColumnList := make([]string, 0)
			for _, cell := range row.Cells {
				str:=cell.String()
				if i == 0 && hasHead {
					headList = append(headList, str)
				} else {
					exceldataColumnList = append(exceldataColumnList, str)
				}
				//switch cell.Type() {
				//case xlsx.CellTypeString:
				//	str:= cell.String()
				//	fmt.Printf("str %d [%s]\t", j, str)
				//	if i == 0 && hasHead {
				//		headList = append(headList,str)
				//	} else {
				//		exceldataColumnList = append(exceldataColumnList, str)
				//	}
				//case xlsx.CellTypeStringFormula:
				//	fmt.Printf("formula %d %s\t", j, cell.Formula())
				//case xlsx.CellTypeNumeric:
				//	x, _ := cell.Int64()
				//	str:=cell.String()
				//	fmt.Printf("int %d %d %s\t", j, x,str)
				//	if i == 0 && hasHead {
				//		headList = append(headList,fmt.Sprintf("%d",x))
				//	} else {
				//		exceldataColumnList = append(exceldataColumnList, fmt.Sprintf("%d",x))
				//	}
				//case xlsx.CellTypeBool:
				//	b:=cell.Bool()
				//	fmt.Printf("bool %d %v\t", j, b)
				//	if i == 0 && hasHead {
				//		headList = append(headList,fmt.Sprintf("%v",b))
				//	} else {
				//		exceldataColumnList = append(exceldataColumnList, fmt.Sprintf("%v",b))
				//	}
				//case xlsx.CellTypeDate:
				//	t, _ := cell.GetTime(false)
				//	fmt.Printf("date %d %v\t", j, t)
				//	if i == 0 && hasHead {
				//		headList = append(headList,fmt.Sprintf("%v",t))
				//	} else {
				//		exceldataColumnList = append(exceldataColumnList, fmt.Sprintf("%v",t))
				//	}
				//default:
				//	return nil,nil,errors.New("excel数据类型不支持")
				//}
			}
			if len(exceldataColumnList) == 0 {
				continue
			}
			excelDataRowsList = append(excelDataRowsList, exceldataColumnList)
		}
	}
	return headList, excelDataRowsList, nil
}
