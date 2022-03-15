
# payload

> https://datatracker.ietf.org/doc/html/rfc7519#section-4.1

## Registered Claim Names

保留使用的.


### "iss" (Issuer) Claim

发行人.
The "iss" (issuer) claim identifies the principal that issued the JWT.
The processing of this claim is generally application specific.
此声明的处理通常是特定于应用程序的。

区分大小写的.
可选的.





### "sub" (Subject) Claim


The subject value MUST either be scoped to be locally unique in the context of the issuer or be globally unique.

主题值必须被限定为在发行人的上下文中是本地唯一的，或者是全球唯一的。


此声明的处理通常是特定于应用程序的。
大小写敏感， 可选的.


### "aud" (Audience) Claim

观众.

既然前面有 发现人，那么就有 接收人, 接收人就是 观众.

这个地方 暂时不用去管这个.


### "exp" (Expiration Time) Claim

“exp”（过期时间）声明标识过期时间 或者在此之后，JWT 不得被接受处理。

数字

可选的.

这个地方 也没有说单位的.

这个地方是一个什么数字?

**代表这个JWT的过期时间；**


过期时间戳. 秒数.


### "nbf" (Not Before) Claim


The "nbf" (not before) claim identifies the time before which the JWT
MUST NOT be accepted for processing.


The processing of the "nbf"
claim requires that the current date/time MUST be after or equal to
the not-before date/time listed in the "nbf" claim.

也是可选的.

声明了一个 nbf 时间戳.
比如 1个小时后的时间戳，那么结果是什么?

token 为非法的.

**当前时间 应该大于，等于 nbf 才有效，否则无效.**

> 反过来说： 定义在什么时间之前，该jwt都是**不可用的.**


用在什么情况下那?


时间。

**是一个时间， 代表这个 jwt 这个 生效的开始时间，意味着在这个时间之前验证 jwt 是会失败的。**


### "iat" (Issued At) Claim

The "iat" (issued at) claim identifies the time at which the JWT was
issued.


This claim can be used to determine the age of the JWT.

jwt的时间.

签发时间.



### "jti" (JWT ID) Claim


The "jti" (JWT ID) claim provides a unique identifier for the JWT.


The "jti" claim can be used
to prevent the JWT from being replayed.


可以使用“jti”声明 以防止 JWT 被重放。


## 4.2 Public Claim Names


唯一即可:


## 4.3 Private Claim Names

即不是 已注册的，也不是 共有的.


私有声明，有可能会冲突，注意使用

Private Claim Names are subject to collision and should
be used with caution.

