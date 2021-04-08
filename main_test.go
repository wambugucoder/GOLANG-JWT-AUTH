package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang_auth/services"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	app := Setup()

	t.Run("Hello Api", func(t *testing.T) {
		res, err := app.Test(httptest.NewRequest("GET", "/api/v1", nil))

		assert.Equal(t, "200 OK", res.Status, "Status Code")
		assert.Equal(t, nil, err, "No Error")
	})

	t.Run("Register Api", func(t *testing.T) {
		registerDetails := services.RegistrationInput{
			Username: "abcdefgh",
			Email:    "abc@gmail.com",
			Password: "123456",
		}
		//CONVERT STRUCT TO BYTES
		marshal, _ := json.Marshal(registerDetails)

		//CONVERT BYTES TO IO READER
		ioData := bytes.NewReader(marshal)

		req := httptest.NewRequest("POST", "/api/v1/auth/register", ioData)

		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req)

		assert.Equal(t, "200 OK", res.Status, "Status Code")
		assert.Equal(t, nil, err, "No Error")
	})

}
