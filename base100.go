package dongle

// reference: https://github.com/stek29/base100
// verify: https://ctf.bugku.com/tool/base100

// ByBase100 encodes by base100.
// 通过 base100 编码
func (e encode) ByBase100() encode {
	if len(e.src) == 0 {
		return e
	}
	buf := make([]byte, len(e.src)*4)
	for i, b := range e.src {
		buf[i*4+0] = 0xf0
		buf[i*4+1] = 0x9f
		buf[i*4+2] = byte((uint16(b)+55)/64 + 0x8f)
		buf[i*4+3] = (b+55)%64 + 0x80
	}
	e.dst = buf
	return e
}

// ByBase100 decodes by base100.
// 通过 base100 解码
func (d decode) ByBase100() decode {
	if len(d.src) == 0 {
		return d
	}
	if len(d.src)%4 != 0 {
		d.Error = invalidDecodingError("base100")
		return d
	}
	buf := make([]byte, len(d.src)/4)
	for i := 0; i != len(d.src); i += 4 {
		if d.src[i+0] != 0xf0 || d.src[i+1] != 0x9f {
			d.Error = invalidDecodingError("base100")
			return d
		}
		buf[i/4] = (d.src[i+2]-0x8f)*64 + d.src[i+3] - 0x80 - 55
	}
	d.dst = buf
	return d
}
