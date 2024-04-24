package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// cyymallUser business-level http error codes.
// the cyymallUserNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	cyymallUserNO       = 62
	cyymallUserName     = "cyymallUser"
	cyymallUserBaseCode = errcode.HCode(cyymallUserNO)

	ErrCreateCyymallUser         = errcode.NewError(cyymallUserBaseCode+1, "failed to create "+cyymallUserName)
	ErrDeleteByIDCyymallUser     = errcode.NewError(cyymallUserBaseCode+2, "failed to delete "+cyymallUserName)
	ErrDeleteByIDsCyymallUser    = errcode.NewError(cyymallUserBaseCode+3, "failed to delete by batch ids "+cyymallUserName)
	ErrUpdateByIDCyymallUser     = errcode.NewError(cyymallUserBaseCode+4, "failed to update "+cyymallUserName)
	ErrGetByIDCyymallUser        = errcode.NewError(cyymallUserBaseCode+5, "failed to get "+cyymallUserName+" details")
	ErrGetByConditionCyymallUser = errcode.NewError(cyymallUserBaseCode+6, "failed to get "+cyymallUserName+" details by conditions")
	ErrListByIDsCyymallUser      = errcode.NewError(cyymallUserBaseCode+7, "failed to list by batch ids "+cyymallUserName)
	ErrListByLastIDCyymallUser   = errcode.NewError(cyymallUserBaseCode+8, "failed to list by last id "+cyymallUserName)
	ErrListCyymallUser           = errcode.NewError(cyymallUserBaseCode+9, "failed to list of "+cyymallUserName)
	// error codes are globally unique, adding 1 to the previous error code
)
