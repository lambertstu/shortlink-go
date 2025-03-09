package tool

import (
	"crypto/rand"
	"math/big"
)

const (
	// 可用的字符和数字
	CHARACTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// 序列长度
	SEQUENCE_LENGTH = 6
)

// 生成一个包含数字和字符的六位随机序列
func GenerateRandomSequence() (string, error) {
	var sequence string
	for i := 0; i < SEQUENCE_LENGTH; i++ {
		// 从字符集中随机选择一个字符
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(CHARACTERS))))
		if err != nil {
			return "", err
		}
		sequence += string(CHARACTERS[index.Int64()])
	}
	return sequence, nil
}
