package dongle

// ByBase16 encodes by base16.
// 通过 base16 编码
func (e encode) ByBase16() encode {
	return e.ByHex()
}

// ByBase16 decodes by base16.
// 通过 base16 解码
func (d decode) ByBase16() decode {
	if d.ByHex().Error != nil {
		d.Error = invalidCiphertextError("base16")
		return d
	}
	return d.ByHex()
}
