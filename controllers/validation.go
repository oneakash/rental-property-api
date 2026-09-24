package controllers

import (
	apierrors "rental-property-api/errors"
	"strconv"
)

func newParamError(name string, message string)error{
	return &apierrors.APIError{
		StatusCode: 400,
		Field: name,
		Message: message,
	}
}

//validate float
func parseFloatParam(value string, name string,)(float64, error){
	if value ==""{
		return 0, nil
	}
	result, err:=strconv.ParseFloat(value, 64)
	if err!=nil{
		return 0, newParamError(
			name,
			"must be a valid number",
		)
	}
	return result, nil
}

//validate int
func parseIntParam(
	value string, name string,
) (int, error) {
	if value==""{
		return 0, nil
	}
	result, err:=strconv.Atoi(value)
	if err!=nil{
		return 0, newParamError(
			name,
			"must be a valid integer",
		)
	}
	return result, nil
}

func parseBoolParam(value string, name string,)(*bool, error){
	if value==""{
		return nil, nil
	}
	result, err:=strconv.ParseBool(value)
	if err!=nil{
		return nil,
		newParamError(
			name,
			"must be true or false",
		)
	}
	return &result, nil
}