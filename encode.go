package dongle

import (
	"io/ioutil"
)

// encode defines a encode struct.
// 定义 encode 结构体
type encode struct {
	dongle
}

// returns a new encode instance.
// 初始化 encode 结构体
func newEncode() encode {
	return encode{}
}

// FromString encodes from string.
// 对字符串进行编码
func (e encode) FromString(s string) encode {
	e.src = string2bytes(s)
	return e
}

// FromBytes encodes from byte slice.
// 对字节切片进行编码
func (e encode) FromBytes(b []byte) encode {
	e.src = b
	return e
}

// FromFile encodes from file.
// 对文件进行编码
func (e encode) FromFile(f interface{}) encode {
	filename := ""
	switch v := f.(type) {
	case string:
		filename = v
	case []byte:
		filename = bytes2string(v)
	}
	if len(filename) == 0 {
		return e
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		e.Error = invalidFileError(filename)
		return e
	}
	e.src = bytes
	return e
}

// String implements the interface Stringer for encode struct.
// 实现 Stringer 接口
func (e encode) String() string {
	return e.ToString()
}

// ToString outputs as string.
// 输出字符串
func (e encode) ToString() string {
	return bytes2string(e.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (e encode) ToBytes() []byte {
	if len(e.dst) == 0 {
		return []byte("")
	}
	return e.dst
}
