package whitelist

import "errors"

var (
	ErrIpAdd    = errors.New("白名单IP添加失败")
	ErrIpRemove = errors.New("白名单IP移除失败")
)
