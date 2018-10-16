package dao

import (
	"testing"
	"hdg.com/auto-demo/src/common"
	"fmt"
)

func TestDynamicDaoImpl_Insert(t *testing.T) {
	db, err := common.GetDefaultDB()
	if err != nil {
		panic(err)
	}
	dao := NewDynamicDaoImpl(db)
	sql := `insert into tb_bid (bid,bidName,password,allowIps,validateIpFlag,status) value (5,'测试','1adwq','127.0.0.1',1,1)`
	if err := dao.Insert(sql); err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}

func TestDynamicDaoImpl_DeleteTable(t *testing.T) {
	db, err := common.GetDefaultDB()
	if err != nil {
		panic(err)
	}
	dao := NewDynamicDaoImpl(db)
	if err := dao.DeleteTable("tb_bid"); err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}

func TestDynamicDaoImpl_CreateTable(t *testing.T) {
	db, err := common.GetDefaultDB()
	if err != nil {
		panic(err)
	}
	dao := NewDynamicDaoImpl(db)
	sql := `create table tb_dept(Id int primary key auto_increment, Name varchar(18),description varchar(100))`
	if err := dao.Insert(sql); err != nil {
		fmt.Println("失败")
	} else {
		fmt.Println("成功")
	}
}

func TestDynamicDaoImpl_HasTable(t *testing.T) {
	db, err := common.GetDefaultDB()
	if err != nil {
		panic(err)
	}
	dao := NewDynamicDaoImpl(db)
	fmt.Println(dao.HasTable("tb_dept"))
}

func TestDynamicDaoImpl_TruncateTable(t *testing.T) {
	db, err := common.GetDefaultDB()
	if err != nil {
		panic(err)
	}
	dao := NewDynamicDaoImpl(db)
	if err:=dao.TruncateTable("t_order");err!=nil {
		fmt.Println("清理成功")
	}
}
