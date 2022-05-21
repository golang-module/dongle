package dongle

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

// ByHmacSha224 encrypts by hmac with sha224.
// 通过 hmac-sha224 加密
func (e encrypt) ByHmacSha224(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha256.New224, string2bytes(v))
	case []byte:
		mac = hmac.New(sha256.New224, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
