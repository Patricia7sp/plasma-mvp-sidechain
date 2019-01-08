package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Reserve errors 100 ~ 199
const (
	DefaultCodespace sdk.CodespaceType = "msgs"

	CodeInvalidAddress     sdk.CodeType = 1
	CodeInvalidOIndex      sdk.CodeType = 2
	CodeInvalidAmount      sdk.CodeType = 3
	CodeInvalidTransaction sdk.CodeType = 4
)

func ErrInvalidTransaction(codespace sdk.CodespaceType, msg string, args ...interface{}) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidTransaction, msg, args)
}

func ErrInvalidAddress(codespace sdk.CodespaceType, msg string, args ...interface{}) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidAddress, msg, args)
}

func ErrInvalidOIndex(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidOIndex, msg)
}

func ErrInvalidAmount(codespace sdk.CodespaceType, msg string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidAmount, msg)
}
