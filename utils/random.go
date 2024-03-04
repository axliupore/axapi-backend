package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"unsafe"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

const (
	// 6 bits to represent a letter index
	letterIdBits = 6
	// All 1-bits as many as letterIdBits
	letterIdMask = 1<<letterIdBits - 1
	letterIdMax  = 63 / letterIdBits
)

func RandomString(n int) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdMax letters!
	for i, cache, remain := n-1, src.Int63(), letterIdMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdMax
		}
		if idx := int(cache & letterIdMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func RandomNumber(n int) string {
	// 使用当前时间的纳秒部分作为种子，以确保每次运行都有不同的种子
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// 根据参数 n 生成相应位数的随机整数，并格式化为字符串
	formatString := fmt.Sprintf("%%0%dd", n)
	num := fmt.Sprintf(formatString, random.Intn(int(math.Pow10(n))))
	return num
}
