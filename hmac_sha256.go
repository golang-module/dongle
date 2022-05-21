package dongle

import (
	"crypto/hmac"
	"crypto/sha256"
	"hash"
)

// ByHmacSha256 encrypts by hmac with sha256.
// 通过 hmac-sha256 加密
func (e encrypt) ByHmacSha256(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha256.New, string2bytes(v))
	case []byte:
		mac = hmac.New(sha256.New, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
