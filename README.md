# jwt_hello


go.dev 上面有的:



* [markbates/goth](github.com/markbates/goth)
* [dgrijalva/jwt-go](https://pkg.go.dev/github.com/dgrijalva/jwt-go) 
* 新仓库地址是 [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
 


## What the heck is a JWT?

JWT.io has [a great introduction](https://jwt.io/introduction) to JSON Web Tokens.

In short, it's a signed JSON object that does something useful (for example, authentication). 
It's commonly used for `Bearer` tokens in `Oauth 2`. 
A token is made of three parts, separated by .'s. 
The first two parts are JSON objects, that have been base64url encoded. The last part is the signature, encoded the same way.

最后一部分是签名的.

The first part is called the header.
It contains the necessary information for verifying the last part, the signature.


The part in the middle is the interesting bit.
It's called the Claims and contains the actual stuff you care about.



Refer to [RFC 7519](https://datatracker.ietf.org/doc/html/rfc7519) for information about reserved keys and the proper way to add your own.


## What's in the box?


Current supported signing algorithms are HMAC SHA, RSA, RSA-PSS, and ECDSA, though hooks are present for adding your own.


