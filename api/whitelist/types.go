package whitelist

type AddIPToWhitelistRequest struct {
	IP string `json:"ip"`
}

type Response struct {
	Message string `json:"msg"`
}
