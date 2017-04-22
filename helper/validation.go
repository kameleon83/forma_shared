package helper

import (
	"fmt"
	"regexp"
)

func ValNumber(s string) bool {
	return reg(`(?m)^[0-9]+$`, s)
}

func ValEmail(s string) bool {
	return reg(`^([a-z0-9_\.-]+\@[\da-z\.-]+\.[a-z\.]{2,6})$`, s)
}
func ValSmtp(s string) bool {
	return reg(`^(smtp|mail).[a-zA-Z-]+.(fr|com|net):(587)$`, s)
}

func reg(reg, str string) bool {
	r, _ := regexp.Compile(reg)
	if r.MatchString(str) {
		fmt.Println(str)
		return true
	}
	fmt.Println(str)

	return false
}
