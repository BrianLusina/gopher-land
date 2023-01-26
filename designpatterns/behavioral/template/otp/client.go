package otp

import "fmt"

func main() {
	smsOTP := &Sms{}
	o := OTP{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println("")
	emailOTP := &Email{}
	o = OTP{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
