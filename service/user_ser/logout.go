package user_ser

import (
	"server/service/redis_ser"
	"server/utils/jwts"
	"time"
)

func (UserService) Logout(claims *jwts.CustomClaims, token string) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)
}
