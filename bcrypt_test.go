package dongle

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcrypt_ToString(t *testing.T) {
	s1 := Sign.FromString("").ByBcrypt(10)
	v1 := Verify.FromString("").ByBcrypt(s1.ToString())
	assert.Equal(t, false, v1)

	s2 := Sign.FromString("hello world").ByBcrypt(10)
	fmt.Println("bcrypt sign：", s2)
	v2 := Verify.FromString("hello world").ByBcrypt(s2.ToString())
	assert.Equal(t, true, v2)

	s3 := Sign.FromString("hello world").ByBcrypt()
	fmt.Println("bcrypt sign：", s3)
	v3 := Verify.FromString("hello world").ByBcrypt(s2.ToString())
	assert.Equal(t, true, v3)
}

func TestBcrypt_ToBytes(t *testing.T) {
	s1 := Sign.FromBytes([]byte("")).ByBcrypt(1)
	v1 := Verify.FromBytes([]byte("")).ByBcrypt(s1.ToBytes())
	assert.Equal(t, false, v1)

	s2 := Sign.FromBytes([]byte("hello world")).ByBcrypt(10)
	v2 := Verify.FromBytes([]byte("hello world")).ByBcrypt(s2.ToBytes())
	assert.Equal(t, true, v2)
}

func TestBcrypt_BcryptRounds_Error(t *testing.T) {
	s := Sign.FromString("hello world").ByBcrypt(1)
	fmt.Println(s.Error)
	assert.Equal(t, invalidBcryptRoundsError(), s.Error)
}
