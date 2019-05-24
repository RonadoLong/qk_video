package model

// Account Account.
type Account struct {
	ID        int64  `json:"id" bind:"require"`
	UserID    string `json:"userID"`
	Username  string `json:"username" bind:"require"`
	Phone     string `json:"phone" bind:"require"`
	Password  string `json:"password" bind:"require"`
	Avatar    string `json:"avatar" bind:"require"`
	Recommend string `json:"recommend" bind:"require"`
	VipLevel  int8   `json:"vipLevel" bind:"require"`
}

type RegisterReq struct {
	Phone     string `json:"phone" bind:"require"`
	Password  string `json:"password" bind:"require"`
	Code      string `json:"code" bind:"require"`
	Recommend string `json:"recommend" bind:"require"`
}

type LoginReq struct {
	Phone     string `json:"phone" bind:"require"`
	Password  string `json:"password" bind:"require"`
	VerifyVal string `json:"verifyVal" bind:"require"`
	CodeID    string `json:"codeID"  bind:"require"`
}

type AccountResp struct {
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	Level    int8   `json:"level"`
	Username string `json:"username"`
}
