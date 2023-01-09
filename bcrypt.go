package dongle

import (
	"golang.org/x/crypto/bcrypt"
)

// ByBcrypt signs by bcrypt.
// 通过 bcrypt 签名
func (s Signer) ByBcrypt(rounds ...int) Signer {
	if len(s.src) == 0 || s.Error != nil {
		return s
	}
	if len(rounds) == 0 {
		// 10 is the standard number of rounds in bcrypt.
		rounds = []int{10}
	}
	dst, err := bcrypt.GenerateFromPassword(s.src, rounds[0])
	if rounds[0] < 4 || rounds[0] > 31 || err != nil {
		s.Error = invalidBcryptRoundsError()
		return s
	}
	s.dst = dst
	return s
}

// ByBcrypt verify by bcrypt.
// 通过 bcrypt 验签
func (v Verifier) ByBcrypt() Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	v.Error = bcrypt.CompareHashAndPassword(interface2bytes(v.sign), v.src)
	return v
}
