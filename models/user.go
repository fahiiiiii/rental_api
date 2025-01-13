package controllers

import (
    "yourproject/initialize"
    "yourproject/models"
    "github.com/beego/beego/v2/server/web"
)

type UserController struct {
    web.Controller
}

func (c *UserController) CreateUser() {
    user := models.User{
        Name: "John Doe",
        Email: "john@example.com",
    }
    
    db := initialize.GetDB()
    result := db.Create(&user)
    if result.Error != nil {
        c.Data["json"] = map[string]interface{}{
            "error": result.Error.Error(),
        }
    } else {
        c.Data["json"] = user
    }
    c.ServeJSON()
}