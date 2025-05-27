package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5encrypto对目标字符串取Hash salt：加盐字段，iteration：hash迭代轮数。
func Md5encrypto(password string, salt string, iteration int) string {
	h := md5.New()
	h.Write([]byte(salt))
	h.Write([]byte(password))
	var res []byte
	res = h.Sum(nil)
	for i := 0; i < iteration-1; i++ {
		h.Reset()
		h.Write(res)
		res = h.Sum(nil)
	}
	return hex.EncodeToString(res)
}
