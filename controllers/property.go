package controllers

import (
	"rental-property-api/models"
	"rental-property-api/services"
	"strconv"
	"strings"
	"github.com/beego/beego/v2/server/web"
)

type PropertyController struct {
	web.Controller
}

func (p *PropertyController) Get() {
	filter := p.getFilter()
	properties := services.GetFilteredProperties(filter)
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
func (p *PropertyController) getFilter() models.PropertyFilter {
	filter := models.PropertyFilter{}
	// min_price
	minPrice := p.GetString("min_price")
	if minPrice != "" {
		value, err := strconv.ParseFloat(minPrice,64)
		if err == nil {
			filter.MinPrice=value
		}

	}
	//max_price
	maxPrice := p.GetString("max_price")
	if maxPrice != "" {
		value, err := strconv.ParseFloat(maxPrice,64)
		if err == nil {
			filter.MaxPrice=value
		}
	}
	//min_star_rating
	minStar:=p.GetString("min_star_rating")
	if minStar!=""{
		value, err:=strconv.Atoi(minStar)
		if err == nil{
			filter.MinStarRating = value
		}
	}
	//min_review_score
	minReviewScore:=p.GetString("min_review_score")
	if minReviewScore!=""{
		value, err:=strconv.ParseFloat(minReviewScore, 64)
		if err == nil{
			filter.MinReviewScore = value
		}
	}

	//min_reviews
	minReviews:=p.GetString("min_reviews")
	if minReviews!=""{
		value, err := strconv.Atoi(minReviews)
		if err == nil{
			filter.MinReviews=value
		}
	}

	// published
	published := p.GetString("published")
	if published != ""{
		value, err := strconv.ParseBool(published)
		if err==nil{
			filter.Published = &value
		}
	}
	//property_type
	propertyType := p.GetString("property_type")
	if propertyType!=""{
		filter.PropertyType=propertyType
	}
	//feed
	feed:=p.GetString("feed")
	if feed!=""{
		value, err := strconv.Atoi(feed)
		if err == nil{
			filter.Feed=value
		}
	}
	// min_bedroom
	minBedroom:=p.GetString("min_bedroom")
	if minBedroom!=""{
		value, err := strconv.Atoi(minBedroom)
		if err == nil{
			filter.MinBedroom = value
		}
	}

	//amenities
	amenities := p.GetString("amenities")
	if amenities != ""{
		filter.Amenities = strings.Split(amenities, ",")
	}

	//limit
	limit := p.GetString("limit")
	if limit != "" {
		value, err := strconv.Atoi(limit)
		if err == nil {
			filter.Limit=value
		}
	}
	return filter
}