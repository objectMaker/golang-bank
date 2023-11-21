package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	fmt.Println(time.Now().UnixNano())
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 生成min-max的随机整数
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandomName(nameLengthParams ...int) string {
	randomString := "abcdefghijklmnopqrstuvwxyz"
	//循环nameLength次
	var nameLength int
	if len(nameLengthParams) == 0 {
		nameLength = 6
	} else {
		nameLength = nameLengthParams[0]
	}
	var name string
	for i := 0; i <= nameLength; i++ {
		name += string(randomString[RandomInt(0, int64(len(randomString)))])
	}
	return name
}
