package test

import (
	"fmt"
	"gin_im/define"
	"os"
	"testing"
)

var MailPassword = os.Getenv("MailPassword")

// go test -v get_env_var_test.go
func TestGetMailPassword(t *testing.T) {
	fmt.Println("MailPassword --> ", MailPassword)

	emailPwd := define.MailPassword
	fmt.Println("Get define.MailPassword value --> ", emailPwd)  // need root or Administrator
}
