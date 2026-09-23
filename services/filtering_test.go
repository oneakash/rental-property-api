package services

import (
	"testing"
	"rental-property-api/models"
)

func TestFilterAND(t *testing.T){
	Properties = []models.SourceProperty{
		{
			ID:"1",
			Feed:11,
			Published:false,
			USDPrice:100,
		},
		{
			ID:"2",
			Feed:11,
			Published:true,
			USDPrice:100,
		},
		{
			ID:"3",
			Feed:12,
			Published:false,
			USDPrice:100,
		},

	}

	filter := models.PropertyFilter{
		Feed:11,
		Published:&[]bool{false}[0],
	}

	result := GetFilteredProperties(filter)

	if len(result)!=1 {
		t.Errorf(
			"expected 1 result got %d",
			len(result),
		)
	}
	if result[0].ID!="1" {
		t.Errorf(
			"wrong property returned",
		)
	}
}

func TestAmenitiesOR(t *testing.T){
	Properties=[]models.SourceProperty{
		{
			ID:"1",
			Amenities:[]string{
				"Internet",
			},
		},
		{
			ID:"2",
			Amenities:[]string{
				"Gym",
			},
		},
	}

	filter:=models.PropertyFilter{
		Amenities:[]string{
			"Internet",
			"Parking",
		},
	}
	result:=GetFilteredProperties(filter)
	if len(result)!=1 {
		t.Errorf(
			"amenity OR filtering failed",
		)
	}

}

func TestEmptyResult(t *testing.T){
	Properties=[]models.SourceProperty{}
	filter:=models.PropertyFilter{
		Feed:11,
	}
	result:=GetFilteredProperties(filter)

	if result==nil {
		t.Errorf(
			"expected empty slice, got nil",
		)

	}
	if len(result)!=0 {
		t.Errorf(
			"expected 0 results",
		)
	}
}