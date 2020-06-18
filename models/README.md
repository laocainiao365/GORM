# 建立数据模型

引入包
```golang
import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)
```
建立和数据库对应的结构体字段
```golang
type User struct {
  Id int `json:"id"`
  Username string `json:"username"`
  Age int	`json:"age"`
  Sex string `json:"sex"`
}
```
指定所要操作的表
```golang
func (User) TableName() string {
  return "user" //要操作user表
}
```
声明全局 DB
```golang
var DB *gorm.DB
```
建立和数据库的连接及配置：
```golang
/**
以下参数的介绍：
mysql => 操作那种数据库
root => 数据库用户名
12345678 => 数据库密码
beegon => 那个数据库
**/
func init() {
  DB,_ = gorm.Open("mysql", "root:12345678@/beego?charset=utf8&parseTime=True&loc=Local")
}
```
