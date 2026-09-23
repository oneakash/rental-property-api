package services

import (
	"encoding/json"
	"log"
	"os"
	"rental-property-api/models"
	"strconv"
)
var Properties []models.SourceProperty
func LoadProperties() error{
	file, err := os.ReadFile(
		"data/rental_properties.json")
	if err != nil{
		return err
	}
	err = json.Unmarshal(
		file,
		&Properties,
	)
	if err != nil{
		return err
	}
	log.Println(
		"Loaded properties:",
		len(Properties),
	)
	return nil
}
func TransformProperty(source models.SourceProperty,) models.PropertyResponse{
	response := models.PropertyResponse{
		ID: source.ID,
		Feed: source.Feed,
		Published: source.Published,
		GeoInfo: models.GeoInfo{
			Breadcrumbs: parseCategories(source.Categories),
			City: source.City,
			Country: source.Country,
			CountryCode: source.CountryCode,
			Name: source.Display,
			LocationID: source.LocationID,
			State: source.State,
			StateAbbr: source.StateAbbr,
			Lat: getLatitude(source),
			Lon: getLongitude(source),
		},
		Property: models.PropertyInfo{
			Amenities: source.Amenities,
			Name: source.PropertyName,
			Slug: source.PropertySlug,
			PropertyType: source.PropertyType,
			Price: source.USDPrice,
			ReviewScore: source.ReviewScore,
			StarRating: source.StarRating,
			Counts: models.CountInfo{
				Bathroom: source.BathroomCount,
				Bedroom: source.BedroomCount,
				Reviews: source.NumberOfReview,
				Occupancy: source.Occupancy,
			},
			Image: models.ImageInfo{
				Count: len(source.Images),
				Images: source.Images,
			},
		},
	}
	return response
}
func getLatitude(source models.SourceProperty,)float64{
	if len(source.LonLat.Coordinates)>1{
		return source.LonLat.Coordinates[1]
	}
	return 0
}
func getLongitude(source models.SourceProperty,)float64{
	if len(source.LonLat.Coordinates)>0{
		return source.LonLat.Coordinates[0]
	}
	return 0
}
func parseCategories(categoryString string,)[]models.Breadcrumb{
	var breadcrumbs []models.Breadcrumb
	err:=json.Unmarshal(
		[]byte(categoryString),
		&breadcrumbs,
	)
	if err!=nil{
		return []models.Breadcrumb{}
	}
	return breadcrumbs
}

func GetAllProperties(limit string) []models.PropertyResponse{
	responses:=make([]models.PropertyResponse, 0)
	count:=0
	limitValue:=0
	if limit!=""{
		value, err := strconv.Atoi(limit)
		if err == nil{
			limitValue = value
		}
	}
	for _,property:=range Properties{
		response:=TransformProperty(property)
		responses = append(responses, response)
		count++
		if limitValue > 0 && count >= limitValue{
			break
		}
	}
	return responses
}

func GetPropertyByID(id string)(models.PropertyResponse, bool){
	for _, property:=range Properties{
		if property.ID==id{
			return TransformProperty(property), true
		}
	}
	return models.PropertyResponse{}, false
}