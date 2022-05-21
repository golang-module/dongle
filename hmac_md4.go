package dongle

import (
	"crypto/hmac"
	"hash"

	"golang.org/x/crypto/md4"
)

// ByHmacMd4 encrypts by hmac with md4.
// 通过 hmac-md4 加密
func (e encrypt) ByHmacMd4(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(md4.New, string2bytes(v))
	case []byte:
		mac = hmac.New(md4.New, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
