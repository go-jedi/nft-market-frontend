package requestProject

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var baseTest string = "http://localhost:8000/api-v1"
var baseProd string = "http://195.211.99.62:8000/api-v1"
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

func RegisterUser(teleId int64, teleName string, teleIdAdmin int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/registration")
	userRegister := UserRegister{
		TeleId:      teleId,
		TeleName:    teleName,
		TeleIdAdmin: teleIdAdmin,
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

func GetUserMinPrice(teleId int64) ([]UserMinPrice, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/getUserMinPrice")
	userGetUserMinPrice := UserGetUserMinPrice{
		TeleId: teleId,
	}
	userGetUserMinPriceJson, err := json.Marshal(userGetUserMinPrice)
	if err != nil {
		return []UserMinPrice{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetUserMinPriceJson))
	if err != nil {
		return []UserMinPrice{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []UserMinPrice{}, err
		}
		var userGetUserMinPriceResponse UserGetUserMinPriceResponse
		err = json.Unmarshal([]byte(body), &userGetUserMinPriceResponse)
		if err != nil {
			return []UserMinPrice{}, err
		}
		if userGetUserMinPriceResponse.Status == 200 && len(userGetUserMinPriceResponse.Result) > 0 {
			return userGetUserMinPriceResponse.Result, nil
		}
	}

	return []UserMinPrice{}, nil
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

func GetAllCollections() ([]Collection, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "collection/getAll")
	response, err := http.Get(baseUrl)
	if err != nil {
		return []Collection{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []Collection{}, err
		}
		var userGetAllCollectionsResponse UserGetAllCollectionsResponse
		err = json.Unmarshal([]byte(body), &userGetAllCollectionsResponse)
		if err != nil {
			return []Collection{}, err
		}
		if userGetAllCollectionsResponse.Status == 200 && len(userGetAllCollectionsResponse.Result) > 0 {
			return userGetAllCollectionsResponse.Result, nil
		}
	}

	return []Collection{}, nil
}

func GetAllTokensCollection(collectionUid string) ([]TokensGetByCollection, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "collection/getAllTokensCollection")
	userGetAllTokensCollection := UserGetAllTokensCollection{
		UidCollection: collectionUid,
	}
	userGetAllTokensCollectionJson, err := json.Marshal(userGetAllTokensCollection)
	if err != nil {
		return []TokensGetByCollection{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetAllTokensCollectionJson))
	if err != nil {
		return []TokensGetByCollection{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []TokensGetByCollection{}, err
		}
		var userGetAllTokensCollectionResponse UserGetAllTokensCollectionResponse
		err = json.Unmarshal([]byte(body), &userGetAllTokensCollectionResponse)
		if err != nil {
			return []TokensGetByCollection{}, err
		}
		if userGetAllTokensCollectionResponse.Status == 200 && len(userGetAllTokensCollectionResponse.Result) > 0 {
			return userGetAllTokensCollectionResponse.Result, nil
		}
	}

	return []TokensGetByCollection{}, nil
}

func GetToken(tokenUid string) ([]Token, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "collection/getToken")
	userGetToken := UserGetToken{
		TokenUid: tokenUid,
	}
	userGetTokenJson, err := json.Marshal(userGetToken)
	if err != nil {
		return []Token{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetTokenJson))
	if err != nil {
		return []Token{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []Token{}, err
		}
		var userGetTokenResponse UserGetTokenResponse
		err = json.Unmarshal([]byte(body), &userGetTokenResponse)
		if err != nil {
			return []Token{}, err
		}
		if userGetTokenResponse.Status == 200 && len(userGetTokenResponse.Result) > 0 {
			return userGetTokenResponse.Result, nil
		}
	}

	return []Token{}, nil
}

func CheckIsAdmin(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/checkIsAdmin")
	userCheckIsAdmin := UserCheckIsAdmin{
		TeleId: teleId,
	}
	userCheckIsAdminJson, err := json.Marshal(userCheckIsAdmin)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckIsAdminJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCheckIsAdminResponse UserCheckIsAdminResponse
		err = json.Unmarshal([]byte(body), &userCheckIsAdminResponse)
		if err != nil {
			return false, err
		}
		if userCheckIsAdminResponse.Status == 200 && userCheckIsAdminResponse.Result {
			return userCheckIsAdminResponse.Result, nil
		}
	}

	return false, nil
}

func CreateReferral(teleId int64, teleName string, adminReferral int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/createReferral")
	userCreateReferral := UserCreateReferral{
		TeleId:        teleId,
		TeleName:      teleName,
		AdminReferral: adminReferral,
	}
	userCreateReferralJson, err := json.Marshal(userCreateReferral)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCreateReferralJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCreateReferralResponse UserCreateReferralResponse
		err = json.Unmarshal([]byte(body), &userCreateReferralResponse)
		if err != nil {
			return false, err
		}
		if userCreateReferralResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func CheckUserReferral(teleId int64) ([]Count, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/checkUserReferral")
	userCheckUserReferral := UserCheckUserReferral{
		TeleId: teleId,
	}
	userCheckUserReferralJson, err := json.Marshal(userCheckUserReferral)
	if err != nil {
		return []Count{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckUserReferralJson))
	if err != nil {
		return []Count{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []Count{}, err
		}
		var userCheckUserReferralResponse UserCheckUserReferralResponse
		err = json.Unmarshal([]byte(body), &userCheckUserReferralResponse)
		if err != nil {
			return []Count{}, err
		}
		if userCheckUserReferralResponse.Status == 200 && len(userCheckUserReferralResponse.Result) > 0 {
			return userCheckUserReferralResponse.Result, nil
		}
	}

	return []Count{}, nil
}

func GetUserReferral(teleId int64, teleIdUser int64) ([]Referral, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/getUserReferral")
	userGetUserReferral := UserGetUserReferral{
		TeleId:     teleId,
		TeleIdUser: teleIdUser,
	}
	userGetUserReferralJson, err := json.Marshal(userGetUserReferral)
	if err != nil {
		return []Referral{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetUserReferralJson))
	if err != nil {
		return []Referral{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []Referral{}, err
		}
		var userGetUserReferralResponse UserGetUserReferralResponse
		err = json.Unmarshal([]byte(body), &userGetUserReferralResponse)
		if err != nil {
			return []Referral{}, err
		}
		if userGetUserReferralResponse.Status == 200 && len(userGetUserReferralResponse.Result) > 0 {
			return userGetUserReferralResponse.Result, nil
		}
	}

	return []Referral{}, nil
}

func GetUsersReferral(teleId int64, limit int) ([]Referral, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/getUsersReferral")
	userGetUsersReferral := UserGetUsersReferral{
		TeleId: teleId,
		Limit:  limit,
	}
	userGetUsersReferralJson, err := json.Marshal(userGetUsersReferral)
	if err != nil {
		return []Referral{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetUsersReferralJson))
	if err != nil {
		return []Referral{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []Referral{}, err
		}
		var userGetUsersReferralResponse UserGetUsersReferralResponse
		err = json.Unmarshal([]byte(body), &userGetUsersReferralResponse)
		if err != nil {
			return []Referral{}, err
		}
		if userGetUsersReferralResponse.Status == 200 && len(userGetUsersReferralResponse.Result) > 0 {
			return userGetUsersReferralResponse.Result, nil
		}
	}

	return []Referral{}, nil
}

func AdminGetUserProfile(teleId int64) ([]AdminUserProfile, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/getUserProfile")
	AdminUserProfileGet := AdminUserProfileGet{
		TeleId: teleId,
	}
	AdminUserProfileGetJson, err := json.Marshal(AdminUserProfileGet)
	if err != nil {
		return []AdminUserProfile{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(AdminUserProfileGetJson))
	if err != nil {
		return []AdminUserProfile{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []AdminUserProfile{}, err
		}
		fmt.Println("body -->", string(body))
		var adminUserProfileGetResponse AdminUserProfileGetResponse
		err = json.Unmarshal([]byte(body), &adminUserProfileGetResponse)
		if err != nil {
			return []AdminUserProfile{}, err
		}
		if adminUserProfileGetResponse.Status == 200 && len(adminUserProfileGetResponse.Result) > 0 {
			return adminUserProfileGetResponse.Result, nil
		}
	}

	return []AdminUserProfile{}, nil
}

func AdminCheckIsPremium(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/checkIsPremium")
	userAdminCheckIsPremium := UserAdminCheckIsPremium{
		TeleId: teleId,
	}
	userAdminCheckIsPremiumJson, err := json.Marshal(userAdminCheckIsPremium)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminCheckIsPremiumJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminCheckIsPremiumResponse UserAdminCheckIsPremiumResponse
		err = json.Unmarshal([]byte(body), &userAdminCheckIsPremiumResponse)
		if err != nil {
			return false, err
		}
		if userAdminCheckIsPremiumResponse.Status == 200 && userAdminCheckIsPremiumResponse.Result {
			return true, nil
		}
	}

	return false, nil
}

func AdminUpdatePremium(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/updatePremium")
	userAdminUpdatePremium := UserAdminUpdatePremium{
		TeleId: teleId,
	}
	userAdminUpdatePremiumJson, err := json.Marshal(userAdminUpdatePremium)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminUpdatePremiumJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminUpdatePremiumResponse UserAdminUpdatePremiumResponse
		err = json.Unmarshal([]byte(body), &userAdminUpdatePremiumResponse)
		if err != nil {
			return false, err
		}
		if userAdminUpdatePremiumResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func AdminCheckIsVerified(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/checkIsVerified")
	userAdminCheckIsVerified := UserAdminCheckIsVerified{
		TeleId: teleId,
	}
	userAdminCheckIsVerifiedJson, err := json.Marshal(userAdminCheckIsVerified)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminCheckIsVerifiedJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminCheckIsVerifiedResponse UserAdminCheckIsVerifiedResponse
		err = json.Unmarshal([]byte(body), &userAdminCheckIsVerifiedResponse)
		if err != nil {
			return false, err
		}
		if userAdminCheckIsVerifiedResponse.Status == 200 && userAdminCheckIsVerifiedResponse.Result {
			return true, nil
		}
	}

	return false, nil
}

func AdminUpdateVerification(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/updateVerification")
	userAdminUpdateVerification := UserAdminUpdateVerification{
		TeleId: teleId,
	}
	userAdminUpdateVerificationJson, err := json.Marshal(userAdminUpdateVerification)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminUpdateVerificationJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminUpdateVerificationResponse UserAdminUpdateVerificationResponse
		err = json.Unmarshal([]byte(body), &userAdminUpdateVerificationResponse)
		if err != nil {
			return false, err
		}
		if userAdminUpdateVerificationResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func AdminUpdateMinimPrice(teleId int64, minPrice float64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/adminUpdateMinimPrice")
	adminUpdateMinPrice := AdminUpdateMinPrice{
		TeleId:   teleId,
		MinPrice: minPrice,
	}
	adminUpdateMinPriceJson, err := json.Marshal(adminUpdateMinPrice)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(adminUpdateMinPriceJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var adminUpdateMinPriceResponse AdminUpdateMinPriceResponse
		err = json.Unmarshal([]byte(body), &adminUpdateMinPriceResponse)
		if err != nil {
			return false, err
		}
		if adminUpdateMinPriceResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func AdminAddBalance(teleId int64, needPrice float64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/adminAddBalance")
	userAdminAdBalance := UserAdminAdBalance{
		TeleId:    teleId,
		NeedPrice: needPrice,
	}
	userAdminAdBalanceJson, err := json.Marshal(userAdminAdBalance)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminAdBalanceJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminAdBalanceResponse UserAdminAdBalanceResponse
		err = json.Unmarshal([]byte(body), &userAdminAdBalanceResponse)
		if err != nil {
			return false, err
		}
		if userAdminAdBalanceResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func AdminChangeMinUser(teleId int64, minPrice float64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/adminChangeMinUser")
	userAdminChangeMinUser := UserAdminChangeMinUser{
		TeleId:   teleId,
		MinPrice: minPrice,
	}
	userAdminChangeMinUserJson, err := json.Marshal(userAdminChangeMinUser)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminChangeMinUserJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminChangeMinUserResponse UserAdminChangeMinUserResponse
		err = json.Unmarshal([]byte(body), &userAdminChangeMinUserResponse)
		if err != nil {
			return false, err
		}
		if userAdminChangeMinUserResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func AdminChangeBalance(teleId int64, needPrice float64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/adminChangeBalance")
	userAdminChangeBalance := UserAdminChangeBalance{
		TeleId:    teleId,
		NeedPrice: needPrice,
	}
	userAdminChangeBalanceJson, err := json.Marshal(userAdminChangeBalance)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userAdminChangeBalanceJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userAdminChangeBalanceResponse UserAdminChangeBalanceResponse
		err = json.Unmarshal([]byte(body), &userAdminChangeBalanceResponse)
		if err != nil {
			return false, err
		}
		if userAdminChangeBalanceResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func CheckIsBlockUser(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/checkIsBlockUser")
	userCheckIsBlockUser := UserCheckIsBlockUser{
		TeleId: teleId,
	}
	userCheckIsBlockUserJson, err := json.Marshal(userCheckIsBlockUser)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCheckIsBlockUserJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCheckIsBlockUserResponse UserCheckIsBlockUserResponse
		err = json.Unmarshal([]byte(body), &userCheckIsBlockUserResponse)
		if err != nil {
			return false, err
		}
		if userCheckIsBlockUserResponse.Status == 200 && userCheckIsBlockUserResponse.Result {
			return true, nil
		}
	}

	return false, nil
}

func BlockUser(teleId int64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "admin/blockUser")
	userBlockUser := UserBlockUser{
		TeleId: teleId,
	}
	userBlockUserJson, err := json.Marshal(userBlockUser)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userBlockUserJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userBlockUserResponse UserBlockUserResponse
		err = json.Unmarshal([]byte(body), &userBlockUserResponse)
		if err != nil {
			return false, err
		}
		if userBlockUserResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func GetAdminByUser(teleId int64) ([]AdminByUser, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/getAdminByUser")
	userGetAdminByUser := UserGetAdminByUser{
		TeleId: teleId,
	}
	userGetAdminByUserJson, err := json.Marshal(userGetAdminByUser)
	if err != nil {
		return []AdminByUser{}, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userGetAdminByUserJson))
	if err != nil {
		return []AdminByUser{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []AdminByUser{}, err
		}
		var userGetAdminByUserResponse UserGetAdminByUserResponse
		err = json.Unmarshal([]byte(body), &userGetAdminByUserResponse)
		if err != nil {
			return []AdminByUser{}, err
		}
		if userGetAdminByUserResponse.Status == 200 && len(userGetAdminByUserResponse.Result) > 0 {
			return userGetAdminByUserResponse.Result, nil
		}
	}

	return []AdminByUser{}, nil
}

func CreateDepot(mammothId int64, mammothUsername string, workerId int64, workerUsername string, amount float64) (bool, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "depot/createDepot")
	userCreateDepot := UserCreateDepot{
		MammothId:       mammothId,
		MammothUsername: mammothUsername,
		WorkerId:        workerId,
		WorkerUsername:  workerUsername,
		Amount:          amount,
	}
	userCreateDepotJson, err := json.Marshal(userCreateDepot)
	if err != nil {
		return false, err
	}
	response, err := http.Post(baseUrl, contentType, bytes.NewBuffer(userCreateDepotJson))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, err
		}
		var userCreateDepotResponse UserCreateDepotResponse
		err = json.Unmarshal([]byte(body), &userCreateDepotResponse)
		if err != nil {
			return false, err
		}
		if userCreateDepotResponse.Status == 200 {
			return true, nil
		}
	}

	return false, nil
}

func GetAllExchangeRates() ([]ExchangeRatesGet, error) {
	checkNeedBase(os.Getenv("IS_TESTING"), "user/exchangeRates")
	response, err := http.Get(baseUrl)
	if err != nil {
		return []ExchangeRatesGet{}, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return []ExchangeRatesGet{}, err
		}
		var getAllExchangeRatesResponse GetAllExchangeRatesResponse
		err = json.Unmarshal([]byte(body), &getAllExchangeRatesResponse)
		if err != nil {
			return []ExchangeRatesGet{}, err
		}
		if getAllExchangeRatesResponse.Status == 200 && len(getAllExchangeRatesResponse.Result) > 0 {
			return getAllExchangeRatesResponse.Result, nil
		}
	}

	return []ExchangeRatesGet{}, nil
}
