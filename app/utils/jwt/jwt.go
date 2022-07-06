package jwt

import (
	"Walker/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

// 指定加密密钥
//var jwtSecret=[]byte(global.ConfigValue.GetString("Jwt.JwtSecret"))
var jwtSecret = []byte("111")

//Claim是一些实体（通常指的用户）的状态和额外的元数据

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token

func GenerateToken(UserId uint) (string, error) {
	//设置token有效时间
	claims := Claims{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(time.Duration(global.ConfigValue.GetInt("Jwt.ExpireTime")) * time.Second).Unix(),
			// 指定token发行人
			Issuer: "tangfeipeng",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）

func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	sign, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//判断err具体原因
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token格式有误")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token未激活")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 刷新token
				sign.Valid = true
				return RefreshToken(sign)
			} else {
				return nil, errors.New("无法解析token")
			}
		}
	}
	return AnalysisToken(sign)
}

//解析token
func AnalysisToken(sign *jwt.Token) (*Claims, error) {
	if claims, ok := sign.Claims.(*Claims); ok == false || claims == nil {
		return nil, errors.New("无法解析token")
	} else {
		return claims, nil
	}
}
func VerifyToken(token string) error {
	if claims, err := ParseToken(token); err == nil {
		//判断是否过了刷新有效期
		if time.Now().Unix()-(claims.ExpiresAt+global.ConfigValue.GetInt64("Jwt.RefreshExpireTime")) > 0 {
			//刷新时间也过期了
			return errors.New("token已过期")
		} else {
			return nil
		}
	} else {
		return err
	}
}
func RefreshToken(sign *jwt.Token) (*Claims, error) {
	if claims, err := AnalysisToken(sign); err != nil {
		return nil, err
	} else {
		//token刷新
		if _, err := GenerateToken(claims.UserId); err != nil {
			return nil, err
		} else {
			return claims, nil
		}
	}
}
