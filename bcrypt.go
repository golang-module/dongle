package dongle

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var (
	invalidBcryptRoundsError = func() error {
		return fmt.Errorf("bcrypt: invalid bcrypt round, the round is outside allowed range (4,31)")
	}
)

// ByBcrypt signs by bcrypt.
// 通过 bcrypt 签名
func (s sign) ByBcrypt(rounds int) sign {
	if len(s.src) == 0 {
		return s
	}
	hash, err := bcrypt.GenerateFromPassword(s.src, rounds)
	if rounds < 4 || rounds > 31 || err != nil {
		s.Error = invalidBcryptRoundsError()
		return s
	}
	s.dst = hash
	return s
}

// ByBcrypt verify by bcrypt.
// 通过 bcrypt 验签
func (v verify) ByBcrypt(sign interface{}) bool {
	var hashedPassword []byte
	switch v := sign.(type) {
	case string:
		hashedPassword = string2bytes(v)
	case []byte:
		hashedPassword = v
	}
	err := bcrypt.CompareHashAndPassword(hashedPassword, v.src)
	if err != nil {
		return false
	}
	return true
}
