package blacklist

import "errors"

var (
	ErrIpAdd    = errors.New("黑名单IP添加失败")
	ErrIpRemove = errors.New("黑名单IP移除失败")
)
