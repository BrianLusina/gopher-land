package base7

import "strconv"

func ConvertToBase7(num int) string {
	if num < 0 {
		return "-" + ConvertToBase7(-num)
	}
	if num < 7 {
		return strconv.Itoa(num)
	}
	return ConvertToBase7(num/7) + strconv.Itoa(num%7)
}
