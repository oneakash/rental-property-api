package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

func (p *PropertyController) Get() {
	p.Data["json"] = map[string]string{
		"message": "API working",
	}
	p.ServeJSON()
}

func (p *PropertyController) GetByID(){
	p.Data["json"] = map[string]string{
		"message":"Get property by ID working",
	}
	p.ServeJSON()
}