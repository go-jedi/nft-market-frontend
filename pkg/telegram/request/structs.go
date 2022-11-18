package requestProject

type UserCheckAuth struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckAuthResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
	Status  int    `json:"status"`
}

type UserRegister struct {
	TeleId   int64  `json:"tele_id"`
	TeleName string `json:"tele_name"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserCheckIsTerms struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsTermsResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
	Status  int    `json:"status"`
}

type UserCheckIsForm struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsFormResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
	Status  int    `json:"status"`
}

type UserCheckIsStatusBot struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsStatusBotResponse struct {
	Message string `json:"message"`
	Result  bool   `json:"result"`
	Status  int    `json:"status"`
	IsAdmin bool   `json:"is_admin"`
}
