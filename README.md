# GORM
在beego项目中使用GORM,快速实现增、删、查、改数据库数据，这是个简易教程。

### 安装
- 先go get "github.com/jinzhu/gorm" and "github.com/go-sql-driver/mysql"
- 或者在项目中引入，等启动项目时等待安装。

### 使用
1. 在models目录下新建模型文件如： user.go并引入"github.com/jinzhu/gorm"、"github.com/go-sql-driver/mysql"
```golang
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)
```
2. 建立模型，如：
```golang
type User struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Age int	`json:"age"`
	Sex string `json:"sex"`
}
```
3. 声明要操作的表,如：
```golang
func (User) TableName() string {
	return "user" //声明操作数据库中 - user表
}
```
4. 声明package models包下的与数据库连接，如：
```golang
var DB *gorm.DB //声明全局变量，以便在其他go文件中和controller中使用
func init() {
	DB,_ = gorm.Open("mysql", "root:12345678@/beego?charset=utf8&parseTime=True&loc=Local")
}
```
5. 在入口文件中，执行完run()方法后，销毁gorm，如：
```golang
//main.go 文件内
func main() {
	beego.Run()
	defer models.DB.Close()//销毁
}
```
6. 其他文件中使用，具体使用操作，查看其他文件...
