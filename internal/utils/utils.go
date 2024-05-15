package utils

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"time"
)

// 生成随机字符串
func GenerateRandomString(n int) (string, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// 生成随机整数
func GenerateRandomInt(max int64) (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}

// 获取当前时间戳
func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}
