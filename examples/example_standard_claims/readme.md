# StandardClaims

> https://pkg.go.dev/github.com/golang-jwt/jwt#example-NewWithClaims-StandardClaims


Example (atypical) using the StandardClaims type by itself to parse a token. 
The StandardClaims type is designed to be embedded into your custom types to provide standard validation features. 

You can use it alone, but there's no way to retrieve other fields after parsing. 
See the CustomClaimsType example for intended usage.

单独使用 StandardClaims 类型来解析令牌的示例（非典型）。
StandardClaims 类型旨在嵌入到您的自定义类型中以提供标准验证功能。 
你可以单独使用它，但是解析后无法检索其他字段。 有关预期用途，请参阅 CustomClaimsType 示例。


> 一般还是要使用 CustomClaimsType.