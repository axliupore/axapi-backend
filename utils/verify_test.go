package utils

import (
	"fmt"
	"testing"
)

func TestIsAnyBlank(t *testing.T) {
	if IsAnyBlank("12", 1, "dad") {
		fmt.Println("error")
	}
}
