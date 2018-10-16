## 自动将excel文件导入到数据库

### 1、修改配置文件

>./resources/applicationContext.yaml

```
server:
  name: auto-demo
  host: '127.0.0.1'
  port: 8088
  version: 1.0.1
excelService:
  excelFilePath: 'resources/订单数据.xlsx'  #excel文件路径
  mapperXML: 'resources/mapper.xml'        #映射文件路径
  sheetName: 'Sheet1' #excel sheet名字
  hasTitle: true      #是否包含标题
  maxProcessNum: 100  #每次执行多少条sql,测试可以少一点，也可以调整到 5000
  limit: -1           #当limit为-1时，导入全部,测试用可以少一点，我测试用5
  create: true        #当表不存在是，是否自动创建
  truncate: true      #重试是否清理数据
  dbSourceRef: local  #引用哪一个数据源，与dbconfigs[0].name匹配
dbconfigs:
   # 数据源配置信息
   - name: local
     config:
      switch: true     #mysql开关
      mode: false       #是否打印sql debug日志
      driver: mysql
      host: 127.0.0.1  #数据库地址
      port: 3306       #数据库端口号
      username: root   #数据库用户名
      password: 123456 #数据库密码
      databasename: db_dingzuo #数据库名字
      lifetime: 10
      max-open-num: 5
      max-idle-num: 5
```

### 2、配置映射文件

> ./resources/mapper.xml

```
<?xml version="1.0" encoding="UTF-8" ?>
<tables>
    <table id="sheet1" database="test" tableName="t_order" primaryKey="auto">
        <describe>订单表</describe>
        <columns>
            <column>
                <columnName>orderId</columnName><!--数据库字段名字-->
                <columnType>string</columnType><!--数据类型-->
                <columnSize>20</columnSize><!--字段长度-->
                <columnDesc>订单号</columnDesc><!--与excel表头字段要一致-->
            </column>
            <column>
                <columnName>count</columnName>
                <columnType>int</columnType>
                <columnSize>10</columnSize>
                <columnDesc>票数</columnDesc>
            </column>
            <column>
                <columnName>showDate</columnName>
                <columnType>string</columnType>
                <columnSize>20</columnSize>
                <columnDesc>上映日期</columnDesc>
            </column>
            <!--......-->
        </columns>
    </table>
</tables>
```

### 3、执行

#### linux 和 mac 环境
> ./auto

#### windows
>start ./main.exe 或者双击运行