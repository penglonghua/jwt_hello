package main

import (
	"bytes"
	"crypto/rsa"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/golang-jwt/jwt"
)

// https://pkg.go.dev/github.com/golang-jwt/jwt#example-package-GetTokenViaHTTP

// 通过 http 的方式获取  token

var (
	verifyKey *rsa.PublicKey

	serverPort int
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Define some custom types were going to use within our tokens
type CustomerInfo struct {
	Name string
	Kind string
}

type CustomClaimsExample struct {
	*jwt.StandardClaims        // 标准的
	TokenType           string // type
	CustomerInfo               // 用户信息
}

func main() {

	// 通过 http请求来 获取肯定是发送 post请求
	res, err := http.PostForm(fmt.Sprintf("http://localhost:%v/authenticate", serverPort), url.Values{
		"user": {"test"},
		"pass": {"known"},
	})
	if err != nil {
		fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code", res.StatusCode)
	}

	// Read the token out of the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	tokenString := strings.TrimSpace(buf.String())

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return verifyKey, nil
	})
	fatal(err)

	//claims := token.Claims.(*CustomClaimsExample) // 强制类型转换 不推荐这么写， 一旦转错， panic异常了.
	//fmt.Println(claims.CustomerInfo.Name)

	if claim, ok := token.Claims.(*CustomClaimsExample); !ok {
		fmt.Println(err) // maybe error
	} else {
		fmt.Println(claim.CustomerInfo.Name, claim.CustomerInfo.Kind)
	}

}
