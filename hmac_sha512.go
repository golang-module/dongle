package dongle

import (
	"crypto/hmac"
	"crypto/sha512"
	"hash"
)

// ByHmacSha512 encrypts by hmac with sha512.
// 通过 hmac-sha512 加密
func (e encrypt) ByHmacSha512(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha512.New, string2bytes(v))
	case []byte:
		mac = hmac.New(sha512.New, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
