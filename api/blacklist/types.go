package blacklist

type AddIPToBlacklistRequest struct {
	IP string `json:"ip"`
}

type Response struct {
	Message string `json:"msg"`
}

type IpResponse struct {
	Data []string `json:"data"`
}
