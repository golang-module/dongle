package dongle

// sign defines a sign struct.
// 定义 sign 结构体
type sign struct {
	dongle
}

// returns a new sign instance.
// 初始化 sign 结构体
func newSign() sign {
	return sign{}
}

// FromString signs from string.
// 对字符串进行签名
func (s sign) FromString(message string) sign {
	s.src = string2bytes(message)
	return s
}

// FromBytes signs from byte slice.
// 对字节切片进行签名
func (s sign) FromBytes(message []byte) sign {
	s.src = message
	return s
}

// String implements the interface Stringer for sign struct.
// 实现 Stringer 接口
func (s sign) String() string {
	return s.ToString()
}

// ToString outputs as string.
// 输出字符串
func (s sign) ToString() string {
	return bytes2string(s.dst)
}

// ToBytes outputs as byte slice.
// 输出字节切片
func (s sign) ToBytes() []byte {
	if len(s.dst) == 0 {
		return []byte("")
	}
	return s.dst
}
