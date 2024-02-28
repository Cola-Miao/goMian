package inner

import "time"

const (
	JWTExpiresTime = time.Hour * 24 * 3
	JWTFlushTime   = time.Hour * 24

	RedisTimeout = time.Second * 10

	CookieExpiresTime = 60 * 60 * 24 * 3

	InterviewBufferTime  = time.Hour * 24
	InterviewExpiresTime = time.Hour * 8

	ErrorGroupLimit = 8
)

var (
	JWTSecret = []byte("todo_mv")
)
