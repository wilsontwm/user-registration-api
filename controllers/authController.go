package controllers

import (
	"encoding/json"
	"github.com/wilsontwm/user-registration"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	_ "user-registration-api/models"
	"user-registration-api/utils"
)

type LoginInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=16"`
}

type SignupInput struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=16"`
}

type GetActivationInput struct {
	Email string `validate:"required,email"`
}

type ActivateAccountInput struct {
	ActivationCode string `validate:"required"`
}

type ForgetPasswordInput struct {
	Email string `validate:"required,email"`
}

type ResetPasswordInput struct {
	ResetPasswordCode string `validate:"required"`
	Password          string `validate:"required,min=8,max=16"`
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	input := LoginInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	// Validate the input
	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, utils.GetValidationErrorMessage(err))
		return
	}

	// Login in the user
	in := &userreg.User{}
	in.Email = input.Email
	in.Password = input.Password

	// Login the account
	user, err := userreg.Login(in)

	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, resp, user, "")
}

// Sign up of the user
var Signup = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	input := SignupInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	// Validate the input
	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, utils.GetValidationErrorMessage(err))
		return
	}

	// Save the data into database
	in := &userreg.User{}
	in.Name = input.Name
	in.Email = input.Email
	in.Password = input.Password

	// Signup the account
	user, err := userreg.Signup(in)

	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, resp, user, "")
}

// Get the user activation code by email
var GetActivation = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	input := GetActivationInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	// Validate the input
	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, utils.GetValidationErrorMessage(err))
		return
	}

	// Retrieve the activation code from email
	in := &userreg.User{}
	in.Email = input.Email
	user, err := userreg.GetActivationCode(in)

	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, resp, user, "")
}

// Activate the user account
var ActivateAccount = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	input := ActivateAccountInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	// Validate the input
	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, utils.GetValidationErrorMessage(err))
		return
	}

	// Activate the account
	in := &userreg.User{}
	in.ActivationCode = &input.ActivationCode
	user, err := userreg.ActivateAccount(in)

	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, resp, user, "")
}

// Forget user password
var ForgetPassword = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	input := ForgetPasswordInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	// Validate the input
	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, utils.GetValidationErrorMessage(err))
		return
	}

	// Set the reset password code for the user
	in := &userreg.User{}
	in.Email = input.Email
	user, err := userreg.ForgetPassword(in)

	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, resp, user, "")
}

// Reset user password
var ResetPassword = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})

	input := ResetPasswordInput{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		utils.Fail(w, http.StatusInternalServerError, resp, err.Error())
		return
	}

	// Validate the input
	validate = validator.New()
	err = validate.Struct(input)
	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, utils.GetValidationErrorMessage(err))
		return
	}

	// Reset the password
	in := &userreg.User{}
	in.ResetPasswordCode = &input.ResetPasswordCode
	in.Password = input.Password
	user, err := userreg.ResetPassword(in)

	if err != nil {
		utils.Fail(w, http.StatusBadRequest, resp, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, resp, user, "")
}

// Home (after authentication)
var Home = func(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]interface{})
	data := map[string]interface{}{
		"userID": r.Context().Value("userID"),
		"email":  r.Context().Value("email"),
		"name":   r.Context().Value("name"),
	}
	utils.Success(w, http.StatusOK, resp, data, "User is authenticated.")
}
