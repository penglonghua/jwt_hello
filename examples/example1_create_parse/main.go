package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// For HMAC signing method, the key can be any []byte. It is recommended to generate
// a key using crypto/rand or something equivalent. You need the same key for signing
// and validating.

// 同样的 key 用于签名和校验
var hmacSampleSecret []byte

func Create() (string, error) {

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",                                                   // 这个是自定义数据
		"nbf": time.Date(2022, 3, 15, 12, 40, 0, 0, time.Local).Unix(), // 时间戳描述 not before
		//"nbf": time.Now().Unix(),

		// 当前时间 应该大于，等于 nbf 才有效，否则无效.
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString(hmacSampleSecret) // 签名

	//fmt.Println(tokenString, err)
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
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}

func main() {
	tokenString, err := Create()
	if err != nil {
		return
	}

	fmt.Println("token:", tokenString)

	Parse(tokenString)

}
