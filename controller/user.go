package controllers

import (
	"demo03/models"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"strconv"
)

type AboutController struct {
	beego.Controller
}


type Jsonstr struct {
	Url   string
	Title string
}

func (c *AboutController) Get() {
	user := []models.User{}
	status := models.DB.Find(&user)

	beego.Info(status)

	c.Data["json"] = user
	c.ServeJSON()
}

//查询
func (this *AboutController) Find(){
	user := models.User{Id: 2}
	models.DB.Find(&user)
	this.Data["json"] = user
	this.ServeJSON()
}

//增加
func (this *AboutController) Add(){
	username := this.GetString("username")
	age := this.GetString("age")
	sex := this.GetString("sex")
	a, _ := strconv.Atoi(age)
	beego.Info(username)
	beego.Info(a)
	beego.Info(sex)

	user := models.User{
		Username: username,
		Age: a,
		Sex: sex,
	}

	var status *gorm.DB = models.DB.Create(&user)

	beego.Info(status.Error)
	if status.Error != nil{
		this.Ctx.WriteString("增加失败！")
		return
	}
	this.Ctx.WriteString("添加成功")
}

// 更新
func (this *AboutController) Update(){
	user := models.User{Id: 5}
	models.DB.Find(&user)
	beego.Info(user)
	user.Username = "赵强强22"
	user.Age = 34

	var status *gorm.DB = models.DB.Save(&user)

	if status.Error != nil{
		beego.Info(status.Error)
		return
	}


	this.Ctx.WriteString("update success!")
}

//删除
func (this *AboutController) Delete(){
	id,_ := this.GetInt("id")

	beego.Info(id)

	user := models.User{Id: id}
	var status *gorm.DB = models.DB.Delete(&user)

	if status.Error != nil{
		beego.Info(status.Error)
		return
	}

	this.Ctx.WriteString("delete success!")
}



