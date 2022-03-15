
# What the heck is a JWT?

> https://jwt.io/introduction

## What is JSON Web Token?



JSON Web Token (JWT) is an open standard ([RFC 7519](https://tools.ietf.org/html/rfc7519)) 
that defines a compact and self-contained way for securely transmitting information between parties as a JSON object.

This information can be verified and trusted because it is digitally signed.


JWTs can be signed using a secret (with the **HMAC** algorithm) or a public/private key pair using **RSA** or **ECDSA**.


其实这个地方 也涉及到签名.

虽然 JWT 可以加密以在各方之间提供保密性，但我们将专注于签名令牌。 
签名的令牌可以验证其中包含的声明的完整性，而加密的令牌会向其他方隐藏这些声明。 
当使用公钥/私钥对对令牌进行签名时，签名还证明只有持有私钥的一方才是签署它的一方。




## When should you use JSON Web Tokens?

* Authorization.

Single Sign 单点登陆.

Single Sign On is a feature that widely uses JWT nowadays, because of its small overhead and its ability to be easily used across different domains.
单点登录是当今广泛使用 JWT 的一项功能，因为它的开销很小并且能够在不同的域中轻松使用。




* Information Exchange: 

您还可以验证内容没有被篡改。

JSON Web Tokens are a good way of securely transmitting information between parties.

JSON Web 令牌是在各方之间安全传输信息的好方法。

Because JWTs can be signed.

for example, using public/private key pairs—you can be sure the senders are who they say they are.
Additionally, as the signature is calculated using the **header** and the **payload**, you can also verify that the content hasn't been tampered with.

***

这个地方 刚刚 提到了 header 和 payload, 马上就必须讲 jwt 的 structure ?

## What is the JSON Web Token structure?

In its compact form, JSON Web Tokens consist of three parts separated by dots (.), which are:

紧凑格式.

* Header
* Payload
* Signature

Therefore, a JWT typically looks like the following.
```text

xxxxx.yyyyy.zzzzz

```

### Header

The header _typically_ consists of two parts:
the type of the token, which is JWT, and the signing algorithm being used, such as HMAC SHA256 or RSA.

For example:

```json5
{
  "alg": "HS256",  // 签名算法 比如 HMAC SHA256 or RSA.
  "typ": "JWT"     // the type of the token , 通常就是 JWT
}
```

Then, this JSON is `Base64Url` encoded to form the first part of the JWT.


### Payload

The second part of the token is the payload, which contains the **claims**.

这个地方 翻译为 声明 要好一点， k8里面也是该关键字.


Claims are statements about an entity (typically, the user) and additional data.

> 通常来说就是 数据本身，一般是关于用户的数据.


There are **three** types of claims: **registered**, **public**, and **private** claims.


* Registered claims:
  These are a set of predefined claims which are not mandatory but recommended, to provide a set of useful, interoperable claims.
  这些是一组预定义的声明，它们不是强制性的，而是推荐的，以提供一组有用的、可互操作的声明。

> 其实就是 预定义的声明，推荐使用的.


Some of them are: **iss** (issuer), **exp** (expiration time), **sub** (subject), aud (audience), and `others`.

有哪些那?




> Notice that the claim names are only three characters long as JWT is meant to be compact.

3个字母.

* Public claims:
  These can be defined at will by those using JWTs.
  But to avoid collisions(冲突) they should be defined in the [IANA JSON Web Token Registry](https://www.iana.org/assignments/jwt/jwt.xhtml) or be defined as a URI that contains a collision resistant namespace.
  

* Private claims:
  These are the `custom` claims created to share information between parties that agree on using them and are neither registered or public claims.

An example payload could be:

```json5
{
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true
}
```


> Do note that for signed tokens this information, though protected against tampering, is readable by anyone. Do not put secret information in the payload or header elements of a JWT unless it is encrypted.

请注意，对于已签名的令牌，此信息虽然受到保护以防篡改，但任何人都可以读取。 除非已加密，否则请勿将机密信息放入 JWT 的有效负载或标头元素中。


> 加密之后，再加密.


###  Signature


To create the signature part you have to take the encoded header, the encoded payload, a secret, the algorithm specified in the header, 
and sign that.


For example if you want to use the HMAC SHA256 algorithm, the signature will be created in the following way:


```text
HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)
```


The signature is used to verify the message wasn't changed along the way,
in the case of tokens signed with a private key, it can also verify that the sender of the JWT is who it says it is.

### Putting all together


![](https://cdn.auth0.com/content/jwt/encoded-jwt3.png)


If you want to play with JWT and put these concepts into practice, you can use [jwt.io Debugger](https://jwt.io/#debugger-io) to decode, verify, and generate JWTs.

***

## How do JSON Web Tokens work?



In general, you should not keep tokens longer than required.


Whenever the user wants to access a protected route or resource, the user agent should send the JWT, 
typically in the Authorization header using the Bearer schema. 
The content of the header should look like the following:

```text
Authorization: Bearer <token>
```


## Why should we use JSON Web Tokens?
