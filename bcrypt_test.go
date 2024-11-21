package dongle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcrypt_String(t *testing.T) {
	s1 := Sign.FromString("").ByBcrypt(10)
	v1 := Verify.FromRawString(s1.ToRawString(), "").ByBcrypt()
	assert.Equal(t, false, v1.ToBool())

	s2 := Sign.FromString("hello world").ByBcrypt(10)
	v2 := Verify.FromRawString(s2.ToRawString(), "hello world").ByBcrypt()
	assert.Equal(t, true, v2.ToBool())

	s3 := Sign.FromString("hello world").ByBcrypt()
	v3 := Verify.FromRawString(s3.ToRawString(), "hello world").ByBcrypt()
	assert.Equal(t, true, v3.ToBool())

	s4 := Sign.FromString("hello world").ByBcrypt()
	v4 := Verify.FromRawString(fmt.Sprintf("%s", s4), "hello world").ByBcrypt()
	assert.Equal(t, true, v4.ToBool())
}

func TestBcrypt_Bytes(t *testing.T) {
	s1 := Sign.FromBytes([]byte("")).ByBcrypt(1)
	v1 := Verify.FromRawBytes(s1.ToRawBytes(), []byte("")).ByBcrypt()
	assert.Equal(t, false, v1.ToBool())

	s2 := Sign.FromBytes([]byte("hello world")).ByBcrypt(10)
	v2 := Verify.FromRawBytes(s2.ToRawBytes(), []byte("hello world")).ByBcrypt()
	assert.Equal(t, true, v2.ToBool())
}

func TestBcrypt_Rounds_Error(t *testing.T) {
	s := Sign.FromString("hello world").ByBcrypt(1)
	assert.Equal(t, NewBcryptError().RoundsError(), s.Error)
}
