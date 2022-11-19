package requestProject

type UserCheckAuth struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckAuthResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserRegister struct {
	TeleId   int64  `json:"tele_id"`
	TeleName string `json:"tele_name"`
}

type UserRegisterResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserGetUserLanguage struct {
	TeleId int64 `json:"tele_id"`
}

type UserLanguage struct {
	Id     int    `json:"id"`
	TeleId int64  `json:"tele_id"`
	Lang   string `json:"lang"`
}

type UserGetUserLanguageResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Result  []UserLanguage `json:"result"`
}

type UserUpdateLanguage struct {
	TeleId int64  `json:"tele_id"`
	Lang   string `json:"lang"`
}

type UserUpdateLanguageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserGetUserCurrency struct {
	TeleId int64 `json:"tele_id"`
}

type UserCurrency struct {
	Id       int    `json:"id"`
	TeleId   int64  `json:"tele_id"`
	Currency string `json:"currency"`
}

type UserGetUserCurrencyResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Result  []UserCurrency `json:"result"`
}

type UserUpdateCurrency struct {
	TeleId   int64  `json:"tele_id"`
	Currency string `json:"currency"`
}

type UserUpdateCurrencyResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserCheckIsTerms struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsTermsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserAgreeTerms struct {
	TeleId int64 `json:"tele_id"`
}

type UserAgreeTermsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserGetUserProfile struct {
	TeleId int64 `json:"tele_id"`
}

type UserProfile struct {
	Id           int   `json:"id"`
	TeleId       int64 `json:"tele_id"`
	Balance      int64 `json:"balance"`
	Conclusion   int64 `json:"conclusion"`
	Verification bool  `json:"verification"`
}

type UserGetUserProfileResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Result  []UserProfile `json:"result"`
}

type Payment struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UserGetAllPaymentsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Result  []Payment `json:"result"`
}
