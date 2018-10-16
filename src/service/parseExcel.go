package service

import (
	"encoding/xml"
	"hdg.com/auto-demo/src/model"
	"io/ioutil"
	"fmt"
	"sync"
)

var ExcelParseService ExcelModelService
var once sync.Once

type ExcelModelService interface {
	Parse(filePath string) (interface{}, error)
}

type OrderExcelService struct {
	ExcelModelService
}

func NewOrderExcelModelService() ExcelModelService{
	if ExcelParseService == nil {
		once.Do(func() {
			ExcelParseService = new(OrderExcelService)
		})
	}
	return ExcelParseService
}

func (*OrderExcelService) Parse(filePath string) (interface{}, error) {
	if data, err :=ioutil.ReadFile(filePath);err!=nil {
		return nil,fmt.Errorf("读取%s文件失败,err:%s",filePath,err.Error())
	}else{
		bs := model.ExcelModel{}
		//把xml数据解析成bs对象
		if err:=xml.Unmarshal([]byte(data), &bs);err!=nil {
			return nil,err
		}
		return &bs,nil
	}
}