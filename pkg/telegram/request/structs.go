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
	IsPremium    bool  `json:"is_premium"`
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

type Collection struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Count         int    `json:"count"`
	CollectionUid string `json:"collection_uid"`
}

type UserGetAllCollectionsResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Result  []Collection `json:"result"`
}

type UserGetAllTokensCollection struct {
	UidCollection string `json:"uid_collection"`
}

type TokensGetByCollection struct {
	Id             int     `json:"id"`
	NameCollection string  `json:"name_collection"`
	Count          int     `json:"count"`
	NameToken      string  `json:"name_token"`
	PriceToken     float64 `json:"price_token"`
	TokenUid       string  `json:"token_uid"`
}

type UserGetAllTokensCollectionResponse struct {
	Status  int                     `json:"status"`
	Message string                  `json:"message"`
	Result  []TokensGetByCollection `json:"result"`
}

type UserGetToken struct {
	TokenUid string `json:"token_uid"`
}

type Token struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	Price          float64 `json:"price"`
	Author         string  `json:"author"`
	Blockchain     string  `json:"blockchain"`
	UidCollection  string  `json:"uid_collection"`
	NameCollection string  `json:"name_collection"`
	TokenUid       string  `json:"token_uid"`
}

type UserGetTokenResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  []Token `json:"result"`
}

type UserCheckIsAdmin struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsAdminResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserCreateReferral struct {
	TeleId        int64  `json:"tele_id"`
	TeleName      string `json:"tele_name"`
	AdminReferral int64  `json:"admin_referral"`
}

type UserCreateReferralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserGetUserReferral struct {
	TeleId int64 `json:"tele_id"`
}

type Referral struct {
	Id            int64  `json:"id"`
	TeleId        int64  `json:"tele_id"`
	TeleName      string `json:"tele_name"`
	Created       string `json:"created"`
	AdminReferral int64  `json:"admin_referral"`
}

type UserGetUserReferralResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Result  []Referral `json:"result"`
}

type AdminUserProfileGet struct {
	TeleId int64 `json:"tele_id"`
}

type AdminUserProfile struct {
	Id           int64  `json:"id"`
	TeleId       int64  `json:"tele_id"`
	TeleName     string `json:"tele_name"`
	Balance      int64  `json:"balance"`
	IsPremium    bool   `json:"is_premium"`
	Verification bool   `json:"verification"`
	Conclusion   int64  `json:"conclusion"`
}

type AdminUserProfileGetResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Result  []AdminUserProfile `json:"result"`
}

type UserAdminUpdatePremium struct {
	TeleId int64 `json:"tele_id"`
}

type UserAdminUpdatePremiumResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
