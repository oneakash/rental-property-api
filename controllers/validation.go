package controllers

import (
	"fmt"
	"strconv"
)

func newParamError(name string, message string)error{
	return fmt.Errorf("%s: %s", name, message)
}

//validate float
func parseFloatParam(value string, name string,)(float64, error){
	if value ==""{
		return 0, nil
	}
	result, err:=strconv.ParseFloat(value, 64)
	if err!=nil{
		return 0, fmt.Errorf(
			"invalid %s: must be a valid number", name,
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
		return 0, fmt.Errorf(
			"invalid %s: must be a valid integer", name,
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
		fmt.Errorf(
			"invalid %s: must be true or false", name,
		)
	}
	return &result, nil
}