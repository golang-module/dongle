package dongle

import (
	"golang.org/x/crypto/bcrypt"
)

// ByBcrypt signs by bcrypt.
// 通过 bcrypt 签名
func (s sign) ByBcrypt(rounds ...int) sign {
	if len(s.src) == 0 {
		return s
	}
	if len(rounds) == 0 {
		// 10 is the standard number of rounds in bcrypt.
		rounds = []int{10}
	}
	hash, err := bcrypt.GenerateFromPassword(s.src, rounds[0])
	if rounds[0] < 4 || rounds[0] > 31 || err != nil {
		s.Error = invalidBcryptRoundsError()
		return s
	}
	s.dst = hash
	return s
}

// ByBcrypt verify by bcrypt.
// 通过 bcrypt 验签
func (v verify) ByBcrypt(sign interface{}) bool {
	err := bcrypt.CompareHashAndPassword(interface2bytes(sign), v.src)
	if err == nil {
		return true
	}
	return false
}
