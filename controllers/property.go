package controllers

import (
	"rental-property-api/services"

	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

func (p *PropertyController) Get() {
	properties:=services.GetAllProperties()
	p.Data["json"] = map[string]interface{}{
		"Result":map[string]interface{}{
			"Count":len(properties),
			"Items":properties,
		},
	}
	p.ServeJSON()
}

func (p *PropertyController) GetByID(){
	p.Data["json"] = map[string]string{
		"message":"Get property by ID working",
	}
	p.ServeJSON()
}