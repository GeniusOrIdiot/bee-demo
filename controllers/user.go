package controllers

import (
	"bee-demo/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

// CMS API
type UserController struct {
	beego.Controller
}

func (u *UserController) URLMapping() {
	u.Mapping("add", u.Add)
	u.Mapping("find", u.Find)
}

// @Title get
// @Description Logs out current logged in user session
// @Param	id	path	string	true	"User id"
// @Success 200 {object} models.User
// @router /get/:id [GET]
func (u *UserController) Find() {
	id, _ := u.GetInt(":id")
	user, err := models.Get(id)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = user
	}

	u.ServeJSON()
}

// @Title Add user
// @Description Register a user.
// @Param	body	body	models.User	true	"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [POST]
func (u *UserController) Add() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	_ = models.Create(user)
	return
}
