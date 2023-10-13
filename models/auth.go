package models

type TokenResponse struct {
	Token string `json:"access_token"`
	Type  string `json:"token_type"`
	Time  int    `json:"expires_in"`
}

type CallBackLoginResponse struct {
	AccToken     string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	Time         int    `json:"expires_in"`
	TokenRefresh string `json:"refresh_token"`
}
