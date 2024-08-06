package token

import "time"

// token maker; 方便切换不同 token 生成方式
type Maker interface {
	// 为指定 username 生成 token
	CreateToken(username string, duration time.Duration) (string, error)

	VerifyToken(token string) (*Payload, error)
}
