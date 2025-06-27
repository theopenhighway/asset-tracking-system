package apperror

import (
	"asset-tracking-system/constant"
	"errors"
)

var (
	ErrorNotFound            = errors.New(constant.ResponseMsgErrorNotFound)
	ErrorInternalServer      = errors.New(constant.ResponseMsgErrorInternalServer)
	ErrorInvalidRequest      = errors.New(constant.ResponseMsgErrorInvalidRequest)
	ErrorWrongPwdEmail       = errors.New(constant.ResponseMsgErrorWrongPwdEmail)
	ErrorAuthentication      = errors.New(constant.ResponseMsgErrorAuthentication)
	ErrorUnauthAccess        = errors.New(constant.ResponseMsgErrorUnauthAccess)
	ErrorEmptyPwdEmail       = errors.New(constant.ResponseMsgErrorEmptyPwdEmail)
	ErrorWalletNotFound      = errors.New(constant.ResponseMsgErrorWalletNotFound)
	ErrorTrfMinimum          = errors.New(constant.ResponseMsgErrorTrfMinimum)
	ErrorTrfMaxmimum         = errors.New(constant.ResponseMsgErrorTrfMaxmimum)
	ErrorAccountExist        = errors.New(constant.ResponseMsgErrorAccountExist)
	ErrorInvalidResetToken   = errors.New(constant.ResponseMsgErrorInvalidResetToken)
	ErrorInsufficientBalance = errors.New(constant.ResponseMsgErrorInsufficientBalance)
	ErrorTrfOwnWallet        = errors.New(constant.ResponseMsgErrorTrfOwnWallet)
	ErrorGameBoxNotExist     = errors.New(constant.ResponseMsgErrorGameBoxNotExist)
	ErrorInsufficientGame    = errors.New(constant.ResponseMsgErrorInsufficientGame)
	ErrorInvalidEmail        = errors.New(constant.ResponseMsgErrorInvalidEmail)
	ErrorInvalidPwd          = errors.New(constant.ResponseMsgErrorInvalidPwd)
	ErrorEmailNotFound       = errors.New(constant.ResponseMsgErrorEmailNotFound)
	ErrorPwdLength           = errors.New(constant.ResponseMsgErrorPwdLength)
	ErrorInvalidSourceFunds  = errors.New(constant.ResponseMsgErrorInvalidSourceFund)
)
