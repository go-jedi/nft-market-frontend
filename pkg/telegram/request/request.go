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

func GetUserLanguage(teleId int64) ([]UserLanguage, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/getUserLanguage")
	userCheckIsLanguage := UserGetUserLanguage{
		TeleId: teleId,
	}
	userCheckIsLanguageJson, err := json.Marshal(userCheckIsLanguage)
	if err != nil {
		return []UserLanguage{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckIsLanguageJson))
	if err != nil {
		return []UserLanguage{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []UserLanguage{}, err
		}
		var userCheckIsLanguageResponse UserGetUserLanguageResponse
		err = json.Unmarshal([]byte(body), &userCheckIsLanguageResponse)
		if err != nil {
			return []UserLanguage{}, err
		}
		if userCheckIsLanguageResponse.Status == 200 && len(userCheckIsLanguageResponse.Result) > 0 {
			return userCheckIsLanguageResponse.Result, nil
		}
	}

	return []UserLanguage{}, nil
}

func UpdateLanguage(teleId int64, lang string) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/updateLanguage")
	userUpdateLanguage := UserUpdateLanguage{
		TeleId: teleId,
		Lang:   lang,
	}
	userUpdateLanguageJson, err := json.Marshal(userUpdateLanguage)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userUpdateLanguageJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userUpdateLanguageResponse UserUpdateLanguageResponse
		err = json.Unmarshal([]byte(body), &userUpdateLanguageResponse)
		if err != nil {
			return false, err
		}
		if userUpdateLanguageResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func GetUserCurrency(teleId int64) ([]UserCurrency, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/getUserCurrency")
	userGetUserCurrency := UserGetUserCurrency{
		TeleId: teleId,
	}
	userGetUserCurrencyJson, err := json.Marshal(userGetUserCurrency)
	if err != nil {
		return []UserCurrency{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetUserCurrencyJson))
	if err != nil {
		return []UserCurrency{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []UserCurrency{}, err
		}
		var userGetUserCurrencyResponse UserGetUserCurrencyResponse
		err = json.Unmarshal([]byte(body), &userGetUserCurrencyResponse)
		if err != nil {
			return []UserCurrency{}, err
		}
		if userGetUserCurrencyResponse.Status == 200 && len(userGetUserCurrencyResponse.Result) > 0 {
			return userGetUserCurrencyResponse.Result, nil
		}
	}

	return []UserCurrency{}, nil
}

func UpdateCurrency(teleId int64, currency string) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/updateCurrency")
	userUpdateCurrency := UserUpdateCurrency{
		TeleId:   teleId,
		Currency: currency,
	}
	userUpdateCurrencyJson, err := json.Marshal(userUpdateCurrency)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userUpdateCurrencyJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userUpdateCurrencyResponse UserUpdateCurrencyResponse
		err = json.Unmarshal([]byte(body), &userUpdateCurrencyResponse)
		if err != nil {
			return false, err
		}
		if userUpdateCurrencyResponse.Status == 200 {
			return true, nil
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

func AgreeTerms(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/agreeTerms")
	userAgreeTerms := UserAgreeTerms{
		TeleId: teleId,
	}
	userAgreeTermsJson, err := json.Marshal(userAgreeTerms)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAgreeTermsJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAgreeTermsResponse UserAgreeTermsResponse
		err = json.Unmarshal([]byte(body), &userAgreeTermsResponse)
		if err != nil {
			return false, err
		}
		if userAgreeTermsResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func GetUserProfile(teleId int64) ([]UserProfile, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/getUserProfile")
	userGetUserProfile := UserGetUserProfile{
		TeleId: teleId,
	}
	userGetUserProfileJson, err := json.Marshal(userGetUserProfile)
	if err != nil {
		return []UserProfile{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetUserProfileJson))
	if err != nil {
		return []UserProfile{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []UserProfile{}, err
		}
		var userGetUserProfileResponse UserGetUserProfileResponse
		err = json.Unmarshal([]byte(body), &userGetUserProfileResponse)
		if err != nil {
			return []UserProfile{}, err
		}
		if userGetUserProfileResponse.Status == 200 && len(userGetUserProfileResponse.Result) > 0 {
			return userGetUserProfileResponse.Result, nil
		}
	}

	return []UserProfile{}, nil
}

func GetAllPayments() ([]Payment, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "payment/getAll")
	response, err := http.Get(baseUrl)
	if err != nil {
		return []Payment{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []Payment{}, err
		}
		var userGetAllPaymentsResponse UserGetAllPaymentsResponse
		err = json.Unmarshal([]byte(body), &userGetAllPaymentsResponse)
		if err != nil {
			return []Payment{}, err
		}
		if userGetAllPaymentsResponse.Status == 200 && len(userGetAllPaymentsResponse.Result) > 0 {
			return userGetAllPaymentsResponse.Result, nil
		}
	}

	return []Payment{}, nil
}
