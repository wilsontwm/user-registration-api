package utils

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
	"strings"
)

// Return success response
func Success(w http.ResponseWriter, status int, resp map[string]interface{}, data interface{}, message string) {
	resp["success"] = true
	resp["data"] = data
	resp["message"] = message
	Respond(w, status, resp)
}

// Return fail response
func Fail(w http.ResponseWriter, status int, resp map[string]interface{}, message string) {
	resp["error"] = message
	Respond(w, status, resp)
}

// Return json response
func Respond(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// Build the error message
func GetValidationErrorMessage(err error) string {
	var errors []string
	for _, errz := range err.(validator.ValidationErrors) {
		// Build the custom errors here
		switch tag := errz.ActualTag(); tag {
		case "required":
			errors = append(errors, errz.StructField()+" is required.")
		case "email":
			errors = append(errors, errz.StructField()+" is an invalid email address.")
		case "min":
			if errz.Type().Kind() == reflect.String {
				errors = append(errors, errz.StructField()+" must be more than or equal to "+errz.Param()+" character(s).")
			} else {
				errors = append(errors, errz.StructField()+" must be larger than "+errz.Param()+".")
			}
		case "max":
			if errz.Type().Kind() == reflect.String {
				errors = append(errors, errz.StructField()+" must be lesser than or equal to "+errz.Param()+" character(s).")
			} else {
				errors = append(errors, errz.StructField()+" must be smaller than "+errz.Param()+".")
			}
		default:
			errors = append(errors, errz.StructField()+" is invalid.")
		}
	}

	return strings.Join(errors, " ")
}
