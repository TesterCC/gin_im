package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
