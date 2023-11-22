package jwt

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"log"
)

// DecodedToken 解码token信息
type DecodedToken struct {
	UserId   string `json:"userId"`   //用户id
	Username string `json:"username"` //用户名
	Iat      int    `json:"iat"`      //签发时间
	Iss      string `json:"iss"`      //签发者
}

// GenerateToken 生成JWT
func GenerateToken(claims *jwt.Token, secretKey string) (token string, err error) {
	// 将密钥转换为字节数组
	secretString := secretKey
	secret := []byte(secretString)
	// 使用密钥签名并获得完整的编码后的字符串令牌
	token, err = claims.SignedString(secret)
	return token, err
}

// VerifyToken 验证jwt，返回解码后的token信息
func VerifyToken(token string, secretKey string) *DecodedToken {
	// 将密钥转换为字节数组
	secretString := secretKey
	secret := []byte(secretString)

	// 解析JWT并验证签名
	decoded, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})

	if err != nil {
		return nil
	}
	if !decoded.Valid {
		return nil
	}
	// 将解码后的JWT转换成结构体
	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	jsonErr := json.Unmarshal(jsonString, &decodedToken)
	if jsonErr != nil {
		log.Printf("JWT解析错误：%v", jsonErr)
	}
	return &decodedToken
}
