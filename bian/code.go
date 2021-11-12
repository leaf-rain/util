package bian

import "errors"

const (
	ErrWaiting = 429
	ErrBan     = 418
)

func CodeCheck(code int) error {
	switch code {
	case ErrWaiting:
	case ErrBan:
	default:

	}
	return errors.New("TODO： 待完善code校验")
}
