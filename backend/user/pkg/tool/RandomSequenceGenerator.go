package tool

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	// 可用的字符和数字
	CHARACTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	// 序列长度
	SEQUENCE_LENGTH = 6
)

// 生成一个包含数字和字符的六位随机序列
func generateRandomSequence() (string, error) {
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

// 测试生成指定数量的随机序列并计算重复率
func testRandomSequences(count int) {
	sequences := make(map[string]struct{})
	duplicates := 0

	for i := 0; i < count; i++ {
		sequence, err := generateRandomSequence()
		if err != nil {
			fmt.Println("Error generating random sequence:", err)
			return
		}

		fmt.Println(sequence)
		if _, exists := sequences[sequence]; exists {
			duplicates++
		} else {
			sequences[sequence] = struct{}{}
		}
	}

	duplicateRate := float64(duplicates) / float64(count) * 100
	fmt.Printf("生成的序列总数: %d\n", count)
	fmt.Printf("重复的序列数量: %d\n", duplicates)
	fmt.Printf("重复率: %.2f%%\n", duplicateRate)
}
