package services
import (
	"rental-property-api/models"
)

func GetFilteredProperties(
	filter models.PropertyFilter,
) []models.PropertyResponse {
	responses := make([]models.PropertyResponse, 0)
	count := 0
	for _, property := range Properties {
		// Minimum price
		if filter.MinPrice > 0 {
			if property.USDPrice < filter.MinPrice {
				continue
			}
		}
		// Maximum price
		if filter.MaxPrice > 0 {
			if property.USDPrice > filter.MaxPrice {
				continue
			}
		}
		// Feed filter
		if filter.Feed > 0 {
			if property.Feed != filter.Feed {
				continue
			}
		}
		// Property type
		if filter.PropertyType != "" {
			if property.PropertyType != filter.PropertyType {
				continue
			}
		}
		// Minimum bedroom
		if filter.MinBedroom > 0 {
			if property.BedroomCount < filter.MinBedroom {
				continue
			}
		}

		// Minimum star rating
		if filter.MinStarRating > 0 {
			if property.StarRating < filter.MinStarRating {
				continue
			}
		}

		// Minimum review score
		if filter.MinReviewScore > 0 {
			if property.ReviewScore < filter.MinReviewScore {
				continue
			}
		}

		// Minimum reviews
		if filter.MinReviews > 0 {
			if property.NumberOfReview < filter.MinReviews {
				continue
			}
		}

		// Published filter
		if filter.Published != nil {
			if property.Published != *filter.Published {
				continue
			}
		}

		//amenities OR filter
		if len(filter.Amenities)>0{
			if !hasAnyAmenity(property.Amenities, filter.Amenities,){
				continue
			}
		}

		// If passed all filters
		response := TransformProperty(property)
		responses = append(responses, response)
		count++

		// Apply limit
		if filter.Limit > 0 && count >= filter.Limit {
			break
		}
	}
	return responses
}

func hasAnyAmenity(
	propertyAmenities []string,
	requiredAmenities []string,
) bool {
	for _, required:=range requiredAmenities{
		for _,available:= range propertyAmenities{
			if required==available{
				return true
			}
		}
	}
	return false
}