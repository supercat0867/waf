package blacklist

import "errors"

var (
	ErrIpAdd    = errors.New("err Add IP To BlackList")
	ErrIpRemove = errors.New("err Remove IP To BlackList")
)
