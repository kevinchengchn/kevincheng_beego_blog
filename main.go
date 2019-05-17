package main

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kevinchengchn/kevincheng_beego_blog/models"
	_ "github.com/kevinchengchn/kevincheng_beego_blog/routers"
)

func init() {
	models.Init()
	beego.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	beego.Run()
}
