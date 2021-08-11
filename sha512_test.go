package dongle

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt_BySha512_FromString_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		encodeMode string // 编码模式
		expected   string // 期望值
	}{
		{"", "hex", ""},
		{"", "base32", ""},
		{"", "base64", ""},
		{"hello world", "xxx", ""},
		{"hello world", "hex", "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"},
		{"hello world", "base32", "GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y="},
		{"hello world", "base64", "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha512()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha512_FromString_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{"", "hex", []byte{}},
		{"", "base32", []byte{}},
		{"", "base64", []byte{}},
		{"hello world", "xxx", []byte("")},
		{"hello world", "hex", []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")},
		{"hello world", "base32", []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")},
		{"hello world", "base64", []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")},
	}

	for index, test := range tests {
		e := Encrypt.FromString(test.input).BySha512()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha512_FromBytes_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   string // 期望值
	}{
		{[]byte(""), "hex", ""},
		{[]byte(""), "base32", ""},
		{[]byte(""), "base64", ""},
		{[]byte("hello world"), "xxx", ""},
		{[]byte("hello world"), "hex", "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"},
		{[]byte("hello world"), "base32", "GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y="},
		{[]byte("hello world"), "base64", "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw=="},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha512()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToString(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestEncrypt_BySha512_FromBytes_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      []byte // 输入值
		encodeMode string // 编码模式
		expected   []byte // 期望值
	}{
		{[]byte(""), "hex", []byte{}},
		{[]byte(""), "base32", []byte{}},
		{[]byte(""), "base64", []byte{}},
		{[]byte("hello world"), "xxx", []byte("")},
		{[]byte("hello world"), "hex", []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")},
		{[]byte("hello world"), "base32", []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")},
		{[]byte("hello world"), "base64", []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")},
	}

	for index, test := range tests {
		e := Encrypt.FromBytes(test.input).BySha512()
		assert.Nil(e.Error)
		assert.Equal(test.expected, e.ToBytes(test.encodeMode), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_BySha512_FromFile(t *testing.T) {
	file := "./demo.txt"
	e := Encrypt.FromFile(file).BySha512()
	assert.Equal(t, invalidFileError(file), e.Error, "Should catch an exception")
}
