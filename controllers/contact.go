package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"mennr.tech/api/helper"
	"mennr.tech/api/services"
)

func GetContactUs(w http.ResponseWriter, r *http.Request) {
	var response services.JSONResponse

	response.Error = false
	response.Message = "no data to show"

	err := helper.EncodeJson(w, http.StatusOK, response)
	if err != nil {
		helper.ErrorResponse(w, err, http.StatusInternalServerError)
	}

}

func PostContactUs(w http.ResponseWriter, r *http.Request) {
	data, err := helper.DecodeJson[services.Contact](w, r)
	if err != nil {
		fmt.Println(err)
		helper.ErrorResponse(w, err)
		return
	}

	err = data.HandleContactData()

	if err != nil {
		err = errors.New("500 internal server error")
		helper.ErrorResponse(w, err, http.StatusInternalServerError)
		return

	}
	var response services.JSONResponse
	response.Error = false
	response.Message = "successfully received the details. Details we got."
	response.Data = data
	helper.EncodeJson(w, http.StatusAccepted, response)
}
