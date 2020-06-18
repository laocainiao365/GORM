# 使用GORM

### 在控制器中使用GORM

导入要是操作的`models`包和`gorm`包
```golang
import (
    "demo03/models" //导入models
    "github.com/astaxie/beego"
    "github.com/jinzhu/gorm"//导入gorm
)
```
使用GORM操作数据库

基础增删查改查看-[user](https://github.com/laocainiao365/GORM/blob/master/controller/user.go)