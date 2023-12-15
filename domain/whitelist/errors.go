package whitelist

import "errors"

var (
	ErrIpAdd    = errors.New("err Add IP To WhiteList")
	ErrIpRemove = errors.New("err Remove IP To WhiteList")
)
