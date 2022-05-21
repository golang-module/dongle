package dongle

import (
	"crypto/hmac"
	"crypto/sha512"
	"hash"
)

// ByHmacSha384 encrypts by hmac with sha384.
// 通过 hmac-sha384 加密
func (e encrypt) ByHmacSha384(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha512.New384, string2bytes(v))
	case []byte:
		mac = hmac.New(sha512.New384, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
