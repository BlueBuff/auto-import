package service

import (
	"hdg.com/auto-demo/src/common"
	"fmt"
	"hdg.com/auto-demo/src/model"
	"hdg.com/auto-demo/src/util"
	"strings"
	"strconv"
	"hdg.com/auto-demo/src/dao"
)

type ServerService interface {
	Dispatch()
}

type ServerServiceImpl struct {
	ServerService
	dao dao.DynamicDao
}

func NewServerService() ServerService {
	if db, err := common.GetDefaultDB(); err != nil {
		panic(err)
	} else {
		svc := ServerServiceImpl{
			dao: dao.NewDynamicDaoImpl(db),
		}
		return &svc
	}
}

func (svc *ServerServiceImpl) Dispatch() {
	excelModel := NewOrderExcelModelService()
	count := 1
	fmt.Printf("(%d)正在初始化元数据数据...\n", count)
	if xmlData, err := excelModel.Parse(common.ConfigurationContext.ExcelService.MapperXMl); err != nil {
		fmt.Println("解析错误:", err)
		return
	} else {
		em := xmlData.(*model.ExcelModel)
		excelServiceConfig := common.ConfigurationContext.ExcelService
		count++
		fmt.Printf("(%d)正在读取excel文件...\n", count)
		heads, excelData, err := util.Parse(excelServiceConfig.ExcelFilePath, excelServiceConfig.SheetName, excelServiceConfig.Limit, excelServiceConfig.HasTitle)
		if err != nil {
			fmt.Println("读取excel文件失败")
			return
		}
		count++
		fmt.Printf("(%d)读取成功,数据一共有%d条.\n", count, len(excelData))
		columnMap := make(map[string]string)
		columns := make([]string, 0)
		columnType := make([]string, 0)
		columnInfoMap := make(map[string]model.Column)
		var tableName string
		for _, table := range em.Table {
			if !strings.EqualFold(table.Id, excelServiceConfig.SheetName) {
				continue
			}
			tableName = table.TableName
			for _, column := range table.Columns.Column {
				columnMap[column.TableHead] = column.ColumnType
				columns = append(columns, column.ColumnName)
				columnType = append(columnType, column.ColumnType)
				columnInfoMap[column.ColumnName] = column
			}
		}
		valuesList := make([][]string, 0)
		for _, excelValue := range excelData {
			values := make([]string, 0)
			for cos, value := range excelValue {
				if typ, ok := columnMap[heads[cos]]; !ok {
					continue
				} else {
					switch strings.ToLower(typ) {
					case "string":
						val := fmt.Sprintf("'%v'", value)
						values = append(values, val)
					case "int":
						iVal, err := strconv.Atoi(value)
						if err != nil {
							fmt.Sprintf("数据格式错误,value:%s 不是int类型,", value)
						}
						values = append(values, fmt.Sprintf("%d", iVal))
					default:
						fmt.Println("暂不支持当前类型")
						return
					}
				}
			}
			valuesList = append(valuesList, values)
		}
		count++
		fmt.Printf("(%d)动态创建sql中...\n", count)
		sqlBuilder := NewDynamicSQLBuilderImpl(INSERT, tableName, columns, valuesList, excelServiceConfig.MaxProcessNum, columnInfoMap)
		//如果表不存在 并且允许创建
		existsTable := svc.dao.HasTable(tableName)
		if !existsTable {
			if excelServiceConfig.Create {
				count++
				fmt.Printf("(%d)%s 表不存在，创建...\n", count, tableName)
				if sql, err := sqlBuilder.CreateTable(); err != nil {
					fmt.Println("生成sql失败,", err)
					return
				} else {
					count++
					fmt.Sprintf("(%d) create table sql==>\n%s\n", count, sql)
					count++
					if err := svc.dao.CreateTable(sql); err != nil {
						fmt.Printf("(%d)创建%s表失败!\n",count, tableName)
					} else {
						fmt.Printf("(%d)创建%s表成功\n",count, tableName)
					}
				}
			} else {
				count++
				fmt.Printf("(%d)您必须先创建表，或者配置成自动创建表。配置applicationContext.yaml -> create: true \n",count)
				return
			}
		} else {
			if excelServiceConfig.Truncate {
				count++
				fmt.Printf("(%d)清理上次保存的数据...\n", count)
				if err := svc.dao.TruncateTable(tableName); err != nil {
					fmt.Println("清理表数据失败")
				}
			}
		}
		if sqlList, err := sqlBuilder.Build(); err != nil {
			fmt.Println(err)
		} else {
			count++
			fmt.Printf("(%d)sql数量:%d 执行sql...\n", count, len(sqlList))
			for _, sql := range sqlList {
				if err := svc.dao.Insert(sql); err != nil {
					fmt.Printf("插入数据出现错误！err:%v", err)
				}
			}
		}
	}
}
