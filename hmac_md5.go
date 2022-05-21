package dongle

import (
	"crypto/hmac"
	"crypto/md5"
	"hash"
)

// ByHmacMd5 encrypts by hmac with md5.
// 通过 hmac-md5 加密
func (e encrypt) ByHmacMd5(key interface{}) encrypt {
	if len(e.src) == 0 {
		return e
	}
	var mac hash.Hash
	switch v := key.(type) {
	case string:
		mac = hmac.New(md5.New, string2bytes(v))
	case []byte:
		mac = hmac.New(md5.New, v)
	}
	mac.Write(e.src)
	e.dst = mac.Sum(nil)
	return e
}
