package controllers

import (
	"rental-property-api/services"

	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

func (p *PropertyController) Get() {
	limit:=p.GetString("limit")
	properties:=services.GetAllProperties(limit)
	p.Data["json"] = map[string]interface{}{
		"Result":map[string]interface{}{
			"Count":len(properties),
			"Items":properties,
		},
	}
	p.ServeJSON()
}

func (p *PropertyController) GetByID(){
	id:=p.Ctx.Input.Param(":id")
	property, found := services.GetPropertyByID(id)
	if !found{
		p.Ctx.ResponseWriter.WriteHeader(404)
		p.Data["json"]=map[string]string{
			"Error":"Property not found",
		}
		p.ServeJSON()
		return
	}
	p.Data["json"]=property
	p.ServeJSON()
}