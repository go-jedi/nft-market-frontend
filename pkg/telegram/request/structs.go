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
	TeleId      int64  `json:"tele_id"`
	TeleName    string `json:"tele_name"`
	TeleIdAdmin int64  `json:"tele_id_admin"`
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
	Id           int     `json:"id"`
	TeleId       int64   `json:"tele_id"`
	Balance      float64 `json:"balance"`
	IsPremium    bool    `json:"is_premium"`
	Conclusion   float64 `json:"conclusion"`
	Verification bool    `json:"verification"`
}

type UserGetUserProfileResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Result  []UserProfile `json:"result"`
}

type UserGetUserMinPrice struct {
	TeleId int64 `json:"tele_id"`
}

type UserMinPrice struct {
	MinimPrice float64 `json:"minim_price"`
}

type UserGetUserMinPriceResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Result  []UserMinPrice `json:"result"`
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

type UserCheckUserReferral struct {
	TeleId int64 `json:"tele_id"`
}

type Count struct {
	Count int64 `json:"count"`
}

type UserCheckUserReferralResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Result  []Count `json:"result"`
}

type Referral struct {
	Id            int64  `json:"id"`
	TeleId        int64  `json:"tele_id"`
	TeleName      string `json:"tele_name"`
	Created       string `json:"created"`
	AdminReferral int64  `json:"admin_referral"`
}

type UserGetUserReferral struct {
	TeleId     int64 `json:"tele_id"`
	TeleIdUser int64 `json:"tele_id_user"`
}

type UserGetUserReferralResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Result  []Referral `json:"result"`
}

type UserGetUsersReferral struct {
	TeleId int64 `json:"tele_id"`
	Limit  int   `json:"limit"`
}

type UserGetUsersReferralResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Result  []Referral `json:"result"`
}

type AdminUserProfileGet struct {
	TeleId int64 `json:"tele_id"`
}

type AdminUserProfile struct {
	Id           int64   `json:"id"`
	TeleId       int64   `json:"tele_id"`
	TeleName     string  `json:"tele_name"`
	Balance      float64 `json:"balance"`
	IsPremium    bool    `json:"is_premium"`
	Verification bool    `json:"verification"`
	Conclusion   float64 `json:"conclusion"`
}

type AdminUserProfileGetResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Result  []AdminUserProfile `json:"result"`
}

type UserAdminCheckIsPremium struct {
	TeleId int64 `json:"tele_id"`
}

type UserAdminCheckIsPremiumResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserAdminUpdatePremium struct {
	TeleId int64 `json:"tele_id"`
}

type UserAdminUpdatePremiumResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserAdminCheckIsVerified struct {
	TeleId int64 `json:"tele_id"`
}

type UserAdminCheckIsVerifiedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserAdminUpdateVerification struct {
	TeleId int64 `json:"tele_id"`
}

type UserAdminUpdateVerificationResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type AdminUpdateMinPrice struct {
	TeleId   int64   `json:"tele_id"`
	MinPrice float64 `json:"min_price"`
}

type AdminUpdateMinPriceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserAdminAdBalance struct {
	TeleId    int64   `json:"tele_id"`
	NeedPrice float64 `json:"need_price"`
}

type UserAdminAdBalanceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserAdminChangeMinUser struct {
	TeleId   int64   `json:"tele_id"`
	MinPrice float64 `json:"min_price"`
}

type UserAdminChangeMinUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserAdminChangeBalance struct {
	TeleId    int64   `json:"tele_id"`
	NeedPrice float64 `json:"need_price"`
}

type UserAdminChangeBalanceResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserCheckIsBlockUser struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsBlockUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserBlockUser struct {
	TeleId int64 `json:"tele_id"`
}

type UserBlockUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserGetAdminByUser struct {
	TeleId int64 `json:"tele_id"`
}

type AdminByUser struct {
	TeleId   int64  `json:"tele_id"`
	TeleName string `json:"tele_name"`
}

type UserGetAdminByUserResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Result  []AdminByUser `json:"result"`
}

type UserCreateDepot struct {
	MammothId       int64   `json:"mammoth_id"`
	MammothUsername string  `json:"mammoth_username"`
	WorkerId        int64   `json:"worker_id"`
	WorkerUsername  string  `json:"worker_username"`
	Amount          float64 `json:"amount"`
	IsShowName      bool    `json:"is_show_name"`
}

type UserCreateDepotResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ExchangeRatesGet struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type GetAllExchangeRatesResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Result  []ExchangeRatesGet `json:"result"`
}

type UserCheckIsVisibleName struct {
	TeleId int64 `json:"tele_id"`
}

type UserCheckIsVisibleNameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserChangeVisibleName struct {
	TeleId int64 `json:"tele_id"`
}

type UserChangeVisibleNameResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserBalance struct {
	Balance float64 `json:"balance"`
}

type UserGetUserBalance struct {
	TeleId int64 `json:"tele_id"`
}

type UserGetUserBalanceResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Result  []UserBalance `json:"result"`
}

type UserBuyUserToken struct {
	TeleId     int64   `json:"tele_id"`
	TokenUid   string  `json:"token_uid"`
	TokenPrice float64 `json:"token_price"`
}

type UserBuyUserTokenResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserNft struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Author   string  `json:"author"`
	TokenUid string  `json:"token_uid"`
}

type UserGetNft struct {
	NftBuy  []UserNft `json:"nft_buy"`
	NftSell []UserNft `json:"nft_sell"`
}

type UserGetUserNft struct {
	TeleId int64 `json:"tele_id"`
}

type UserGetUserNftResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Result  UserGetNft `json:"result"`
}

type UserSellUserToken struct {
	TeleId     int64   `json:"tele_id"`
	TokenUid   string  `json:"token_uid"`
	TokenPrice float64 `json:"token_price"`
}

type UserSellUserTokenResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type UserAdminBuyTokenUser struct {
	TeleId          int64   `json:"tele_id"`
	TokenUid        string  `json:"token_uid"`
	PriceUser       float64 `json:"price_user"`
	UidPaymentEvent string  `json:"uid_payment_event"`
}

type UserAdminBuyTokenUserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UserCheckUserToken struct {
	TeleId   int64  `json:"tele_id"`
	TokenUid string `json:"token_uid"`
}

type UserCheckUserTokenResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}

type UserGetUserPaymentEvent struct {
	EventUid string `json:"event_uid"`
}

type PaymentEvent struct {
	TeleId    int64   `json:"tele_id"`
	Uid       string  `json:"uid"`
	NameToken string  `json:"name_token"`
	Price     float64 `json:"price"`
}

type UserGetUserPaymentEventResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Result  []PaymentEvent `json:"result"`
}

type UserCreateWithDrawEvent struct {
	TeleId int64   `json:"tele_id"`
	Price  float64 `json:"price"`
}

type UserCreateWithDrawEventResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

type UserGetWithDrawEvent struct {
	WithDrawEventUid string `json:"with_draw_event_uid"`
}

type WithDrawEventGet struct {
	TeleId     int64   `json:"tele_id"`
	Uid        string  `json:"uid"`
	Price      float64 `json:"price"`
	IsFinished bool    `json:"is_finished"`
}

type UserGetWithDrawEventResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Result  []WithDrawEventGet `json:"result"`
}

type UserAdminWithdrawApprove struct {
	TeleId           int64  `json:"tele_id"`
	WithDrawEventUid string `json:"with_draw_event_uid"`
}

type UserAdminWithdrawApproveResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  bool   `json:"result"`
}
