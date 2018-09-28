package main

import (
	_ "Npro/routers"
	"github.com/astaxie/beego"
	"Npro/models"
)


func init(){
	models.InitDataBase()
}
func main() {
	beego.Run()
}
