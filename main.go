package main

import (
	_ "PagingBybeego/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"PagingBybeego/models"
)

func main() {
	orm.RegisterDataBase("default","mysql",beego.AppConfig.String("username")+":"+beego.AppConfig.String("password")+ "@/paging?charset=utf8&loc=Asia%2FShanghai",30)
	orm.RegisterModel(new(models.Person))
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

