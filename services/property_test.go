package services
import (
	"testing"
	"rental-property-api/models"
)


func TestTransformProperty(t *testing.T) {
	source := models.SourceProperty{
		ID:"TEST-001",
		Feed:11,
		Country:"Japan",
		City:"Tokyo",
		PropertyName:"Test Hotel",
		USDPrice:100,
		Images:[]string{
			"1.jpg",
			"2.jpg",
		},
		LonLat:models.LonLat{
			Coordinates:[]float64{
				139.6,
				35.6,
			},
		},
	}

	result := TransformProperty(source)

	if result.ID != "TEST-001" {
		t.Errorf(
			"expected ID TEST-001 got %s",
			result.ID,
		)
	}
	if result.GeoInfo.Lat != 35.6 {
		t.Errorf(
			"latitude mapping failed",
		)
	}

	if result.GeoInfo.Lon != 139.6 {
		t.Errorf(
			"longitude mapping failed",
		)
	}
	if result.Property.Image.Count != 2 {
		t.Errorf(
			"image count failed",
		)
	}
}

func TestGetPropertyByIDFound(t *testing.T){
	Properties=[]models.SourceProperty{
		{
			ID:"ABC-001",
		},
	}
	result,found:=GetPropertyByID(
		"ABC-001",
	)
	if !found {
		t.Error(
			"expected property found",
		)
	}
	if result.ID!="ABC-001" {
		t.Error(
			"wrong property",
		)
	}
}