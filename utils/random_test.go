package utils

import (
	"fmt"
	"testing"
)

func TestRandomString(t *testing.T) {
	fmt.Println(RandomString(20))
}

func TestRandomNumber(t *testing.T) {
	fmt.Println(RandomNumber(6))
}
