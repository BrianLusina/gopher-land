package wallet

import "fmt"

type SecurityCode struct {
	code int
}

func newSecurityCode(code int) *SecurityCode {
	return &SecurityCode{code: code}
}

func (sc *SecurityCode) checkCode(incomingCode int) error {
	if sc.code != incomingCode {
		return fmt.Errorf("Security code is invalid")
	}
	return nil
}
