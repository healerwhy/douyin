package token

import "github.com/golang-jwt/jwt/v4"

type ParseToken struct{}

func (*ParseToken) ParseToken(tokenString string) (*Claims, error) {
	// 解码token
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(AccessSecret), nil
	})
	// 如果不为空，解码成功
	if tokenClaims != nil {
		// 断言转换
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
