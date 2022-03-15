package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func main() {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	//jwt.StandardClaims{
	//	Audience:  "",  // 观众
	//	ExpiresAt: 0,   // 过期时间戳
	//	Id:        "",  // id
	//	IssuedAt:  0,   // 发行时间
	//	Issuer:    "",  // 发行者
	//	NotBefore: 0,   // 在什么时间前 无效
	//	Subject:   "",  // 主题
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}
