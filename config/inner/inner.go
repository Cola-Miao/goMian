package inner

import "time"

const (
	JWTExpiresTime = time.Hour * 24 * 3
	JWTFlushTime   = time.Hour * 24
)

var (
	JWTSecret = []byte("todo_mv")
)
