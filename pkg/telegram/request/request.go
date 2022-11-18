package requestProject

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseTest string = "http://localhost:8080/api-v1"
var baseProd string = "http://localhost:8080/api-v1"
var baseUrl string = ""

var contentType string = "application/json"

func checkNeedBase(isTesting string, finalPath string) {
	if isTesting == "true" {
		baseUrl = fmt.Sprintf("%s/%s", baseTest, finalPath)
	} else {
		baseUrl = fmt.Sprintf("%s/%s", baseProd, finalPath)
	}
}

func CheckAuth(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/checkAuth")
	userCheckAuth := UserCheckAuth{
		TeleId: teleId,
	}
	userCheckAuthJson, err := json.Marshal(userCheckAuth)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckAuthJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCheckAuthResponse UserCheckAuthResponse
		err = json.Unmarshal([]byte(body), &userCheckAuthResponse)
		if err != nil {
			return false, err
		}
		if userCheckAuthResponse.Status == 200 && userCheckAuthResponse.Result {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, nil
}

func RegisterUser(teleId int64, teleName string) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/registration")
	userRegister := UserRegister{
		TeleId:   teleId,
		TeleName: teleName,
	}
	userRegisterJson, err := json.Marshal(userRegister)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userRegisterJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userRegisterResponse UserRegisterResponse
		err = json.Unmarshal([]byte(body), &userRegisterResponse)
		if err != nil {
			return false, err
		}
		if userRegisterResponse.Status == 200 {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, nil
}

func CheckIsTerms(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/checkIsTerms")
	userCheckIsTerms := UserCheckIsTerms{
		TeleId: teleId,
	}
	userCheckIsTermsJson, err := json.Marshal(userCheckIsTerms)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckIsTermsJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCheckIsTermsResponse UserCheckIsTermsResponse
		err = json.Unmarshal([]byte(body), &userCheckIsTermsResponse)
		if err != nil {
			return false, err
		}
		if userCheckIsTermsResponse.Status == 200 && userCheckIsTermsResponse.Result {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, nil
}

func CheckIsForm(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/checkIsForm")
	userCheckIsForm := UserCheckIsForm{
		TeleId: teleId,
	}
	userCheckIsFormJson, err := json.Marshal(userCheckIsForm)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckIsFormJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCheckIsFormResponse UserCheckIsFormResponse
		err = json.Unmarshal([]byte(body), &userCheckIsFormResponse)
		if err != nil {
			return false, err
		}
		if userCheckIsFormResponse.Status == 200 && userCheckIsFormResponse.Result {
			return true, nil
		} else {
			return false, nil
		}
	}

	return false, nil
}

func CheckIsStatusBot(teleId int64) (bool, bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/checkIsStatus")
	userCheckIsStatusBot := UserCheckIsStatusBot{
		TeleId: teleId,
	}
	userCheckIsStatusBotJson, err := json.Marshal(userCheckIsStatusBot)
	if err != nil {
		return false, false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckIsStatusBotJson))
	if err != nil {
		return false, false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, false, err
		}
		var userCheckIsStatusBotResponse UserCheckIsStatusBotResponse
		err = json.Unmarshal([]byte(body), &userCheckIsStatusBotResponse)
		if err != nil {
			return false, false, err
		}
		if userCheckIsStatusBotResponse.Status == 200 && userCheckIsStatusBotResponse.Result {
			return userCheckIsStatusBotResponse.Result, userCheckIsStatusBotResponse.IsAdmin, nil
		} else {
			return userCheckIsStatusBotResponse.Result, userCheckIsStatusBotResponse.IsAdmin, nil
		}
	}
	return false, false, nil
}
