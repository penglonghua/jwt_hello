package main

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var (
	mySigningKey = []byte("AllYourBase")
)

func main() {

	type MyCustomClaims struct {
		Foo                string `json:"foo"`
		jwt.StandardClaims        // 嵌入标准的
	}

	// Create the Claims
	claims := MyCustomClaims{
		"bar",
		jwt.StandardClaims{
			ExpiresAt: 15000,  // 时间戳
			Issuer:    "test", // 发行者
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	// 肯定这个又是 非法的，因为过期时间 太端了
	// 测试一下

	Parse(ss)

}

// parsing and validating a token using the HMAC signing method
func Parse(tokenString string) {

	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	// 这样写有 灵活性.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)                       //  Token is expired
		fmt.Println("----------", token.Valid) //  false
	}

}
