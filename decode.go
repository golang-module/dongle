package dongle

// defines a decode struct.
// 定义 decode 结构体
type decode struct {
	dongle
}

// returns a new decode instance.
// 初始化 decode 结构体
func newDecode() decode {
	return decode{}
}

// FromString decodes from string.
// 对字符串进行解密
func (d decode) FromString(s string) decode {
	d.src = string2bytes(s)
	return d
}

// FromBytes decodes from byte slice.
// 对字节切片进行解密
func (d decode) FromBytes(b []byte) decode {
	d.src = b
	return d
}

// String implements the interface Stringer for decode struct.
// 实现 Stringer 接口
func (d decode) String() string {
	return d.ToString()
}

// ToString outputs as string.
// 输出字符串
func (d decode) ToString() string {
	return bytes2string(d.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (d decode) ToBytes() []byte {
	if len(d.dst) == 0 {
		return []byte("")
	}
	return d.dst
}
