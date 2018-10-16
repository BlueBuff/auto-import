package model

import "encoding/xml"

/**
<?xml version="1.0" encoding="UTF-8" ?>
<tables>
    <table database="test" tableName="t_order" create="true" primaryKey="auto">
        <describe>订单表</describe>
        <columns>
            <!--<column columnName="id" columnType="int" primaryKey="true">主键</column>-->
            <column columnName="orderId" columnType="string">订单号</column>
            <column columnName="count" columnType="int">票数</column>
            <column columnName="showDate" columnType="string">放映日期</column>
            <!--......-->
        </columns>
    </table>
</tables>
 */
type ExcelModel struct {
	//如果有类型为xml.Name的XMLName字段，则解析时会保存元素名到该字段
	XMLName xml.Name `xml:"tables"`
	Table   []Table  `xml:"table"`
}

/**
表元数据
 */
type Table struct {
	XMLName    xml.Name `xml:"table" json:"-"`
	Id         string   `xml:"id,attr" json:"id"`
	DataBase   string   `xml:"database,attr" json:"dataBase"`
	TableName  string   `xml:"tableName,attr" json:"tableName"`
	PrimaryKey string   `xml:"primaryKey,attr" json:"primaryKey"`
	Describe   string   `xml:"describe" json:"describe"`
	Columns    Columns  `xml:"columns" json:"columns"`
}

type Columns struct {
	Column  []Column `xml:"column" json:"column"`
	XMLName xml.Name `xml:"columns"`
}

type Column struct {
	XMLName    xml.Name `xml:"column" json:"-"`
	ColumnName string   `xml:"columnName" json:"columnName"`
	ColumnType string   `xml:"columnType" json:"columnType"`
	TableHead  string   `xml:"columnDesc" json:"columnDesc"`
	ColumnSize int      `xml:"columnSize" json:"columnSize"`
}
