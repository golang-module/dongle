package dongle

import (
	"crypto/hmac"
	"crypto/sha1"
	"hash"
)

// ByHmacSha1 encrypts by hmac with sha1.
// 通过 hmac-sha1 加密
func (e encrypt) ByHmacSha1(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(sha1.New, string2bytes(v))
	case []byte:
		mac = hmac.New(sha1.New, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
