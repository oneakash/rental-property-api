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