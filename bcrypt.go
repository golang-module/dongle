package dongle

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptError struct {
}

func NewBcryptError() BcryptError {
	return BcryptError{}
}

func (e BcryptError) RoundsError() error {
	return fmt.Errorf("bcrypt: invalid rounds, the rounds at least 4 and at most 31")
}

// ByBcrypt signs by bcrypt.
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
		s.Error = NewBcryptError().RoundsError()
		return s
	}
	s.dst = dst
	return s
}

// ByBcrypt verify by bcrypt.
func (v Verifier) ByBcrypt() Verifier {
	if len(v.src) == 0 || v.Error != nil {
		return v
	}
	v.Error = bcrypt.CompareHashAndPassword(v.sign, v.src)
	return v
}
