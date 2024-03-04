package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestEmail(t *testing.T) {
	if htmlTemplate, err := os.ReadFile("template.html"); err == nil {
		fmt.Println(string(htmlTemplate))
	}
}

func TestRegular(t *testing.T) {
	fmt.Println(verifyEmail("12345@126.com")) //true
	fmt.Println(verifyEmail("12345126.com"))  //false
}

func TestCode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Println(generateCode())
	}
}
