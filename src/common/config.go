package common

import "time"

type ApplicationContext struct {
	Server       Server       `yaml:"server"`
	DBConfigs    []DBConfig   `yaml:"dbconfigs"`
	ExcelService ExcelService `yaml:"excelService"`
}

type Server struct {
	Name    string `yaml:"name"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Version string `yaml:"version"`
}

type ExcelService struct {
	ExcelFilePath string `yaml:"excelFilePath"`
	MapperXMl     string `yaml:"mapperXML"`
	SheetName     string `yaml:"sheetName"`
	HasTitle      bool   `yaml:"hasTitle"`
	MaxProcessNum int    `yaml:"maxProcessNum"`
	Create        bool   `yaml:"create"`
	Truncate      bool   `yaml:"truncate"`
	Limit         int    `yaml:"limit"`
	DBSource      string `yaml:"dbSourceRef"`
}

type DBConfig struct {
	Name   string `yaml:"name"`
	Config Config `yaml:"config"`
}

type Config struct {
	Mode            bool          `yaml:"mode"`
	Switch          bool          `yaml:"switch"`
	Driver          string        `yaml:"driver"`
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	UserName        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	DataBaseName    string        `yaml:"databasename"`
	ConnMaxLifetime time.Duration `yaml:"lifetime"`
	MaxOpenNum      int           `yaml:"max-open-num"`
	MaxIdleNum      int           `yaml:"max-idle-num"`
}
