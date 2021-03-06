package core

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
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

// GenerateToken export
func GenerateToken(jsonbyte []byte) string {
	xidstr := xid.New().String()

	var expiresAt int64
	var key string
	expiresAt = time.Now().AddDate(0, 30, 0).Unix()
	//expiresAt = time.Now().Add(time.Millisecond * 1000 * 5).Unix()
	//expiresAt = time.Now().Add(time.Millisecond * 1000 * 60 * 30).Unix()
	key = cfg.Jwt.AccessKey

	payload := PayLoad{
		Data: jsonbyte,
		Xid:  xidstr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "dlog",
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(cfg.Jwt.Alg), &payload)
	tokenstr, _ := token.SignedString([]byte(key))

	return tokenstr
}

// VaildToken export
func VaildToken(tokenString string) string {
	if len(tokenString) == 0 {
		return ""
	}

	decodeToken, err := _decodeToken(tokenString, cfg.Jwt.AccessKey)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(decodeToken.Data)
}

func _decodeToken(tokenString string, secret string) (*PayLoad, *JwtException) {
	var payload PayLoad
	token, err := jwt.ParseWithClaims(tokenString, &payload, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &JwtException{Code: INVAILD}
		}
		return []byte(secret), nil
	})

	var exception *JwtException
	if err != nil {
		parseE, _ := err.(*jwt.ValidationError)
		if parseE.Errors == EXPIRED {
			exception = &JwtException{Code: EXPIRED}
		} else {
			exception = &JwtException{Code: INVAILD}
		}
	} else if !token.Valid {
		exception = &JwtException{Code: INVAILD}
	}

	return &payload, exception
}
