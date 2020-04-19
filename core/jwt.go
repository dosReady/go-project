package core

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
)

// JwtException export
type JwtException struct {
	Code uint32
}

const (
	// INVAILD export
	INVAILD uint32 = 20
	// EXPIRED export
	EXPIRED uint32 = 16
	// PASS export
	PASS uint32 = 0
)

func (je JwtException) Error() string {
	switch {
	case je.Code == INVAILD:
		return fmt.Sprintln("[JWT] 유효하지않은 토큰입니다.")
	case je.Code == EXPIRED:
		return fmt.Sprintln("[JWT] 만료된 토큰입니다.")
	default:
		return fmt.Sprintln("[JWT] 알수없는 오류입니다.")
	}
}

// PayLoad export
type PayLoad struct {
	Data []byte
	Xid  string
	jwt.StandardClaims
}

var cfg = GetConfig()

// GenerateAccessToken export
func GenerateAccessToken(obj interface{}) (string, string) {
	xidstr := xid.New().String()
	jsonobj := EncodingJSON(obj)

	payload := PayLoad{
		Data: jsonobj,
		Xid:  xidstr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Millisecond * 1000).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.Jwt.Alg), &payload)
	tokenstr, _ := token.SignedString([]byte(cfg.Jwt.AccessKey))

	return tokenstr, xidstr
}

// GenerateRefreshToken export
func GenerateRefreshToken(xidval string) string {
	payload := struct {
		Xid string
		jwt.StandardClaims
	}{
		Xid: xidval,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 30, 0).Unix(),
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.Jwt.Alg), payload)
	tokenstr, _ := token.SignedString([]byte(cfg.Jwt.RefreshKey))
	return tokenstr
}

// VaildAccessToken export
func VaildAccessToken(tokenString string) (*PayLoad, uint32) {
	decodeAccess, err := _decodeToken(tokenString, cfg.Jwt.AccessKey)
	return decodeAccess, err
}

// VaildRefreshToken export
func VaildRefreshToken(tokenString string) (*PayLoad, uint32) {
	decdoeRefresh, err := _decodeToken(tokenString, cfg.Jwt.RefreshKey)
	return decdoeRefresh, err
}

func _decodeToken(tokenString string, secret string) (*PayLoad, uint32) {
	var payload PayLoad
	token, err := jwt.ParseWithClaims(tokenString, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &JwtException{Code: INVAILD}
		}
		return []byte(secret), nil
	})

	var exception JwtException
	if err != nil {
		parseE, _ := err.(*jwt.ValidationError)
		if parseE.Errors == EXPIRED {
			exception = JwtException{Code: EXPIRED}
		} else {
			exception = JwtException{Code: INVAILD}
		}
	} else if !token.Valid {
		exception = JwtException{Code: INVAILD}
	}

	return &payload, exception.Code
}
