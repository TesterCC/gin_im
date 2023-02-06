package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"gin_im/define"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/smtp"
)

type UserClaims struct {
	//Id    string `json:"_id"`   // 对应： uc.Id, str not ObjectiveID
	Id    primitive.ObjectID `json:"_id"`   // 对应： uc.Id, str not ObjectiveID
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GetMd5 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// key
var myKey = []byte("im")

// GenerateToken
// 生成 token
func GenerateToken(_id, email string) (string, error) {

	objectiID, err := primitive.ObjectIDFromHex(_id)

	UserClaim := &UserClaims{
		//Id:               _id,
		Id:               objectiID,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// AnalyseToken
// 解析 token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// GetUUID
// 生成唯一码
func GetUUID(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// SendCode
// 发送验证码  Attention:未知原因腾讯系邮箱会拦截，收件邮箱最好不要是腾讯系的
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Admin <grtest00@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "[Notice]验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "grtest00@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}