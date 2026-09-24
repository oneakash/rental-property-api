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

// @Title Get All Properties
// @Description Returns rental properties with optional filtering
// @Param limit query int false "Maximum number of results"
// @Param min_price query float64 false "Minimum price"
// @Param max_price query float64 false "Maximum price"
// @Param feed query int false "Property feed"
// @Param published query bool false "Published status"
// @Param property_type query string false "Property type"
// @Param amenities query string false "Amenities separated by comma"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @router / [get]
func (p *PropertyController) Get() {
	filter, err := p.getFilter()
	if err!=nil{
		p.Ctx.ResponseWriter.WriteHeader(400)
		p.Data["json"]=map[string]string{
			"Error":err.Error(),
		}
		p.ServeJSON()
		return
	}
	properties := services.GetFilteredProperties(filter)
	p.Data["json"] = map[string]interface{}{
		"Result":map[string]interface{}{
			"Count":len(properties),
			"Items":properties,
		},
	}
	p.ServeJSON()
}

// @Title Get Property By ID
// @Description Returns a single rental property
// @Param id path string true "Property ID"
// @Success 200 {object} models.PropertyResponse
// @Failure 404 {object} map[string]string
// @router /:id [get]
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
func (p *PropertyController) getFilter() (models.PropertyFilter, error, ){
	filter := models.PropertyFilter{}
	// min_price
	minPrice, err := parseFloatParam(
		p.GetString("min_price"),
		"min_price",
	)
	if err!=nil{
		return  filter, err
	}
	if minPrice<0{
		return filter, newParamError("min_price", "must be greater than or equal to 0")
	}
	filter.MinPrice = minPrice
	//max_price
	maxPrice, err:=parseFloatParam(
		p.GetString("max_price"),
		"max_price",
	)
	if err != nil{
		return filter, err
	}
	if maxPrice<0{
		return filter, newParamError("max_price", "must be greater than or equal to 0")
	}
	filter.MaxPrice = maxPrice

	//validate price relationship
	if minPrice>0 && maxPrice>0 && minPrice>maxPrice{
		return filter, newParamError(
			"min_price",
			"cannot be greater than max_price",
		)
	}
	//min_star_rating
	minStarRating, err:=parseIntParam(
		p.GetString("min_star_rating"),
		"min_star_rating",
	)
	if err!=nil{
		return filter, err
	}
	if minStarRating < 0 || minStarRating > 5 {
		return filter, newParamError(
			"min_star_rating",
			"must be between 1 and 5",
		)
	}
	filter.MinStarRating = minStarRating
	//min_review_score
	minReviewScore, err := parseFloatParam(
		p.GetString("min_review_score"),
		"min_review_score",
	)
	if err!=nil{
		return filter, err
	}
	if minReviewScore <0 || minReviewScore>10{
		return filter, newParamError("min_review_score", "must be between 0 and 10")
	}
	filter.MinReviewScore=minReviewScore

	//min_reviews
	minReviews, err:=parseIntParam(
		p.GetString("min_reviews"),
		"min_reviews",
	)
	if err!=nil{
		return filter, err
	}
	if minReviews <0{
		return filter, newParamError("min_reviews", "must be greater than or equal to 0")
	}
	filter.MinReviews=minReviews

	// published
	published, err:=parseBoolParam(
		p.GetString("published"),
		"published",
	)
	if err!=nil{
		return filter,err
	}
	filter.Published=published
	//property_type
	propertyType := p.GetString("property_type")
	if propertyType!=""{
		switch propertyType{
			case "Hotel",
				"House",
				"Apartment",
				"Villa",
				"Resort",
				"Hostel":
				filter.PropertyType=propertyType
		default:
			return filter, newParamError("property_type", "must be one of Hotel, House, Apartment, Villa, Resort, Hostel")
		}
		
	}
	//feed
	feed, err := parseIntParam(
		p.GetString("feed"),
		"feed",
	)
	if err!=nil{
		return filter, err
	}
	if feed!=0 && feed!=11 && feed!=12 && feed!=22 &&
		feed!=24{
		return filter, newParamError("feed","must be one of 11, 12, 22, or 24")
	}
	filter.Feed = feed
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
	return filter, nil
}