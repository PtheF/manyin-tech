package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtBuilder struct {
	expire          int64
	claimsGenerator func() jwt.MapClaims
	claims          jwt.MapClaims
	issuer          string
	jwtStr          string
	stamp           string
}

type JwtAuthError struct {
	IsExpired bool
}

func (e *JwtAuthError) Error() string {
	if e.IsExpired {
		return "JwtExpired"
	}

	return ""
}

func NewMapClaimsJwt(mapClaimsGetter func() jwt.MapClaims) *JwtBuilder {
	return &JwtBuilder{claimsGenerator: mapClaimsGetter}
}

func (builder *JwtBuilder) Exp(exp int64) *JwtBuilder {
	builder.expire = exp
	return builder
}

func (builder *JwtBuilder) Iss(iss string) *JwtBuilder {
	builder.issuer = iss
	return builder
}

func (builder *JwtBuilder) Stamp(stamp string) *JwtBuilder {
	builder.stamp = stamp
	return builder
}

func (builder *JwtBuilder) Sign(secret string) (string, error) {

	builder.claims = builder.claimsGenerator()
	if builder.expire != 0 {
		builder.claims["exp"] = time.Now().Unix() + builder.expire
	}

	builder.claims["expSec"] = builder.expire

	if builder.issuer != "" {
		builder.claims["iss"] = builder.issuer
	}

	if builder.stamp != "" {
		builder.claims["stamp"] = builder.stamp
	}

	logs.Info("map claims: %v", builder.claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, builder.claims)
	if tokenStr, err := token.SignedString([]byte(secret)); err != nil {
		logs.Error("jwt sign error: %v", err)
		return "", err
	} else {
		return tokenStr, nil
	}
}

func RefreshJwtWithExpired(jwtStr string) *JwtBuilder {
	return &JwtBuilder{jwtStr: jwtStr}
}

func (builder *JwtBuilder) Refresh(secret string) (string, error) {
	token, err := jwt.Parse(builder.jwtStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		jwtErr := err.(*jwt.ValidationError)

		switch jwtErr.Errors {
		case 16:
			logs.Info("origin jwt expired, start refresh")
			claims := token.Claims.(jwt.MapClaims)
			builder.claims = claims

			logs.Info("origin jwt claims: %v", claims)
			logs.Info("origin jwt expired sec: %v", builder.claims["expSec"])

			if builder.expire != 0 {
				builder.claims["exp"] = time.Now().Unix() + builder.expire
				builder.claims["expSec"] = builder.expire
			} else {
				expSec := builder.claims["expSec"]
				if expSec != 0 {
					logs.Info("default expired: %v", expSec)

					// interface{} = 20
					// 这个20其实被 json 解析成了 float64，所以需要先转化为 float64，再转成 int64
					builder.claims["exp"] = time.Now().Unix() + int64(expSec.(float64))
				}
			}

			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, builder.claims)
			if tokenStr, err := newToken.SignedString([]byte(secret)); err != nil {
				return "", err
			} else {
				return tokenStr, nil
			}
		default:
			logs.Error("other error")
			return "", jwtErr
		}
	} else {
		logs.Info("origin jwt not expired...")
		return "", nil
	}
}

func JwtAuth(jwtStr string, secret string, claims jwt.Claims, jwtType interface{}) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(jwtStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		jwtErr := err.(*jwt.ValidationError)

		switch jwtErr.Errors {
		case 16:
			logs.Info("jwt expired: %v", jwtStr)

			return token, &JwtAuthError{IsExpired: true}

		default:
			return token, err
		}
	}

	return token, nil
}
