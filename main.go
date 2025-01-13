package main

import (
    "rental_api/conf"
    _ "rental_api/routers"
      "rental_api/initialize"
    "github.com/beego/beego/v2/server/web"
)

func init() {
    // Initialize database connection
    initialize.InitializeDB()
}

func main() {
    conf.InitDB()
    web.Run()
}


// func main() {
// 	if beego.BConfig.RunMode == "dev" {
// 		beego.BConfig.WebConfig.DirectoryIndex = true
// 		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
// 	}
// 	beego.Run()
// }
