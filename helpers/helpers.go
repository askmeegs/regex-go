package helpers

import (
	"fmt"
	"regexp"
)

func regexFactory(reg string, s string) bool {
	r, err := regexp.Compile(reg)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return r.MatchString(s)
}

// ares words - starts with "ar" ends with "es"
func Ares(s string) bool {
	return regexFactory(`\A(ar)+\w*(es)+`, s)
}

// Is valid U.S. phone number
// (ddd)-ddd-dddd
func PhoneNumber(s string) bool {
	return regexFactory(`\A\(\d{3}\)-\d{3}-\d{4}\z`, s)
}

// Is valid email address
// [a-z,.,-]@[fqdn]
func Email(s string) bool {
	return regexFactory(`\A[\w\d\.]{3,}@[\w\d]+\.[\w]{2,}\z`, s)
}

// A list of colors of the rainbow?
// blue, orange, blue, red.
func RainbowList(s string) bool {
	return regexFactory(`\A[\w\d\.]{3,}@[\w\d]+\.[\w]{2,}\z`, s)

}
