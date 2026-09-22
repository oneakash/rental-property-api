package services

import (
	"encoding/json"
	"log"
	"os"
	"rental-property-api/models"
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
			City: source.City,
			Country: source.Country,
			CountryCode: source.CountryCode,
			Name: source.Display,
			LocationID: source.LocationID,
			State: source.State,
			StateAbbr: source.StateAbbr,
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