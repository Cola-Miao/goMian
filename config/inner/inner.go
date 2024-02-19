package inner

import "time"

const (
	JWTExpiresTime = time.Hour * 24 * 3
	JWTFlushTime   = time.Hour * 24

	RedisTimeout = time.Second * 10
)

var (
	JWTSecret = []byte("todo_mv")
)
