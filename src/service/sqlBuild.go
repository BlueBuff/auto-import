package service

import (
	"github.com/go-errors/errors"
	"bytes"
	"strings"
	"fmt"
	"hdg.com/auto-demo/src/model"
)

type DynamicSQLBuilder interface {
	Build() ([]string, error)
	CreateTable() (string, error)
}

type SQLType int

const (
	QUERY  SQLType = iota
	DELETE
	INSERT
	UPDATE
)

type DynamicSQLBuilderImpl struct {
	DynamicSQLBuilder
	typ            SQLType
	TableName      string
	Columns        []string
	ColumnTypes    []string
	Values         [][]string
	ColumnTypeInfo map[string]model.Column
	Size           int
}

func NewDynamicSQLBuilderImpl(typ SQLType, tableName string, columns []string, values [][]string, size int, columnTypeInfos map[string]model.Column) DynamicSQLBuilder {
	builder := DynamicSQLBuilderImpl{
		typ:            typ,
		TableName:      tableName,
		Columns:        columns,
		Values:         values,
		Size:           size,
		ColumnTypeInfo: columnTypeInfos,
	}
	return &builder
}

func (dsb *DynamicSQLBuilderImpl) Build() ([]string, error) {
	if len(dsb.Columns) == 0 {
		return nil, errors.New("字段长度不能为0")
	}
	var sql []string
	switch dsb.typ {
	case INSERT:
		sql = dsb.queryBuild()
	default:
		return nil, errors.New("暂时不支持当前类型")
	}
	return sql, nil
}

//insert into tb_bid (bid,bidName,password,allowIps,validateIpFlag,status) value (5,'测试','1adwq','127.0.0.1',1,1)
func (dsb *DynamicSQLBuilderImpl) queryBuild() []string {
	var buf bytes.Buffer
	buf.WriteString("insert into ")
	buf.WriteString(dsb.TableName)
	buf.WriteString(" (")
	buf.WriteString(strings.Join(dsb.Columns, ","))
	buf.WriteString(") ")
	buf.WriteString("values ")
	valueContainer := make([]string, 0)
	var valueBuf bytes.Buffer
	for i, size := 0, len(dsb.Values); i < size; i++ {
		value := dsb.Values[i]
		valueBuf.WriteString("(")
		valueBuf.WriteString(strings.Join(value, ","))
		valueBuf.WriteString(")")
		valueContainer = append(valueContainer, valueBuf.String())
		valueBuf.Reset()
	}
	sqlContainer := make([]string, 0)
	size := dsb.Size
	for i := 0; i < len(valueContainer); i += size {
		inSlice := valueContainer[i : i+size]
		newInSlice := make([]string, 0)
		for _, val := range inSlice {
			if val != "" {
				newInSlice = append(newInSlice, val)
			}
		}
		if len(newInSlice) == 0 {
			continue
		}
		sqlContainer = append(sqlContainer, fmt.Sprintf("%s %s", buf.String(), strings.Join(newInSlice, ",")))
	}
	return sqlContainer
}

func (dsb *DynamicSQLBuilderImpl) CreateTable() (string, error) {
	var buf bytes.Buffer
	buf.WriteString("create table ")
	buf.WriteString(dsb.TableName)
	buf.WriteString("(")
	buf.WriteString("id int primary key auto_increment,")
	var columnDDL = make([]string, 0)
	for columnName, column := range dsb.ColumnTypeInfo {
		if strings.EqualFold("string", column.ColumnType) {
			columnDDL = append(columnDDL, fmt.Sprintf("%s varchar(%d)",columnName, column.ColumnSize))
		}
		if strings.EqualFold("int", column.ColumnType) {
			columnDDL = append(columnDDL,fmt.Sprintf("%s int",columnName))
		}
	}
	buf.WriteString(strings.Join(columnDDL, ","))
	buf.WriteString(")")
	return buf.String(), nil
}
