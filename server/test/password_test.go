package test

import (
	"fmt"
	"server/utils"
	"testing"
)

func TestPassword(t *testing.T) {
	for i := 0; i < 10; i++ {
		password, err := utils.GeneratePassword("123455")
		fmt.Println(password)
		fmt.Println(err)
		fmt.Println("---------------")
		fmt.Println(utils.VerifyPassword(password, "22"))
	}
}
