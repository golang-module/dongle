# dongle  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/dongle.svg)](https://github.com/golang-module/dongle/releases)
[![Go Build](https://github.com/golang-module/dongle/actions/workflows/bulid.yml/badge.svg)](https://github.com/golang-module/dongle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/dongle)](https://goreportcard.com/report/github.com/golang-module/dongle)
[![codecov](https://codecov.io/gh/golang-module/openssl/branch/main/graph/badge.svg)](https://codecov.io/gh/golang-module/dongle)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/dongle)
![License](https://img.shields.io/github/license/golang-module/dongle)

简体中文 | [English](README.md)

一个轻量级、语义化、对开发者友好的 golang 编码解码、加密解密库

`Dongle` 已被 [awesome-go](https://github.com/avelino/awesome-go#security "awesome-go") 收录, 如果您觉得不错，请给个 star 吧

[github.com/golang-module/dongle](https://github.com/golang-module/dongle "github.com/golang-module/dongle")

[gitee.com/golang-module/dongle](https://gitee.com/golang-module/dongle "gitee.com/golang-module/dongle")

### 安装使用
```go
// 使用 github 库
go get -u github.com/golang-module/dongle

import (
    "github.com/golang-module/dongle"
)

// 使用 gitee 库
go get -u gitee.com/golang-module/dongle

import (
    "gitee.com/golang-module/dongle"
)
```

### 编码&解码
#### Hex 编码、解码
```go
// 对字符串进行 hex 编码，输出字符串
dongle.Encode.FromString("hello world").ByHex().ToString() // 68656c6c6f20776f726c64
// 对字符串进行 hex 解码，输出字符串
dongle.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToString() // hello world

// 对字节切片进行 hex 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")
// 对字节切片进行 hex 解码，输出字节切片
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToBytes() // []byte("hello world")
```

#### Base32 编码、解码
```go
// 对字符进行 base32 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// 对字符串进行 base32 解码，输出字符串
dongle.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToString() // hello world

// 对字节切片进行 base32 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")
// 对字节切片进行 base32 解码，输出字节切片
dongle.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToBytes() // []byte("hello world")
```

#### Base64 编码、解码
```go
// 对字符串进行 base64 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase64().ToString() // aGVsbG8gd29ybGQ=
// 对字符串进行 base64 解码，输出字符串
dongle.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToString() // hello world

// 对字节切片进行 base64 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")
// 对字节切片进行 base64 解码，输出字节切片
dongle.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToBytes() // []byte("hello world")
```

#### Base64URL 编码、解码
```go
// 对 url 字符进行 base64 编码，输出字符串
dongle.Encode.FromString("www.gouguoyin.cn").ByBase64URL().ToString() // d3d3LmdvdWd1b3lpbi5jbg==
// 对 url 字符进行 base64 解码，输出字符串
dongle.Decode.FromString("d3d3LmdvdWd1b3lpbi5jbg==").ByBase64URL().ToString() // www.gouguoyin.cn

// 对 url 字节切片进行 base64 编码，输出字节切片
dongle.Encode.FromBytes([]byte("www.gouguoyin.cn")).ByBase64URL().ToBytes() // []byte("d3d3LmdvdWd1b3lpbi5jbg==")
// 对 url 字符进行 base64 解码，输出字节切片
dongle.Decode.FromBytes([]byte("d3d3LmdvdWd1b3lpbi5jbg==")).ByBase64URL().ToBytes() // []byte("www.gouguoyin.cn")
```

#### SafeURL 编码、解码
```go
// 对 url 字符进行转义编码，输出字符串
dongle.Encode.FromString("www.gouguoyin.cn?sex=男&age=18").BySafeURL().ToString() // www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18
// 对 url 字符进行转义解码，输出字符串
dongle.Decode.FromString("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18").BySafeURL().ToString() // www.gouguoyin.cn?sex=男&age=18

// 对 url 字节切片进行转义编码，输出字节切片
dongle.Encode.FromBytes([]byte("www.gouguoyin.cn?sex=男&age=18")).BySafeURL().ToBytes() // []byte("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18")
// 对 url 字符进行转义解码，输出字节切片
dongle.Decode.FromBytes([]byte("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18")).BySafeURL().ToBytes() // []byte("www.gouguoyin.cn?sex=男&age=18")
```

### 加密&解密
#### Md4 加密
```go
// 对字符串进行 md4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd4().ToHexString() // aa010fbc1d14c795d86ef98c95479d17
// 对字符串进行 md4 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd4().ToBase32String() // VIAQ7PA5CTDZLWDO7GGJKR45C4======
// 对字符串进行 md4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd4().ToBase64String() // qgEPvB0Ux5XYbvmMlUedFw==

// 对字节切片进行 md4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToHexBytes() // []byte("aa010fbc1d14c795d86ef98c95479d17")
// 对字节切片进行 md4 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase32Bytes() // []byte("VIAQ7PA5CTDZLWDO7GGJKR45C4======")
// 对字节切片进行 md4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase64Bytes() // []byte("qgEPvB0Ux5XYbvmMlUedFw==")

// 对文件进行 md4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToHexString() // 1240c5c0fb26b585999357915c56b511
// 对文件进行 md4 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToBase32String() // CJAMLQH3E22YLGMTK6IVYVVVCE======
// 对文件进行 md4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToBase64String() // EkDFwPsmtYWZk1eRXFa1EQ==

// 对文件进行 md4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToHexBytes() // []byte("1240c5c0fb26b585999357915c56b511")
// 对文件进行 md4 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToBase32Bytes() // []byte("CJAMLQH3E22YLGMTK6IVYVVVCE======")
// 对文件进行 md4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToBase64Bytes() // []byte("EkDFwPsmtYWZk1eRXFa1EQ==")
```

#### Hmac-md4 加密
```go
// 对字符串进行 hmac-md4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToHexString() // 7a9df5247cbf76a8bc17c9c4f5a75b6b
// 对字符串进行 hmac-md4 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToBase32String() // PKO7KJD4X53KRPAXZHCPLJ23NM======
// 对字符串进行 hmac-md4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToBase64String() // ep31JHy/dqi8F8nE9adbaw==

// 对字节切片进行 hmac-md4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToHexBytes() // []byte("7a9df5247cbf76a8bc17c9c4f5a75b6b")
// 对字节切片进行 hmac-md4 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase32Bytes() // []byte("PKO7KJD4X53KRPAXZHCPLJ23NM======")
// 对字节切片进行 hmac-md4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase64Bytes() // []byte("ep31JHy/dqi8F8nE9adbaw==")
```

#### Md5 加密
```go
// 对字符串进行 md5 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
// 对字符串进行 md5 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToBase32String() // L23DXO7AD3XNBE6LEK5Y6WWNYM======
// 对字符串进行 md5 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToBase64String() // XrY7u+Ae7tCTyyK7j1rNww==

// 对字节切片进行 md5 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToHexBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// 对字节切片进行 md5 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase32Bytes() // []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")
// 对字节切片进行 md5 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase64Bytes() // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// 对文件进行 md5 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToHexString() // 014f03f9025ea81a8a0e9734be540c53
// 对文件进行 md5 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBase32String() // AFHQH6ICL2UBVCQOS42L4VAMKM======
// 对文件进行 md5 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBase64String() // AU8D+QJeqBqKDpc0vlQMUw==

// 对文件进行 md5 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToHexBytes() // []byte("014f03f9025ea81a8a0e9734be540c53")
// 对文件进行 md5 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToBase32Bytes() // []byte("AFHQH6ICL2UBVCQOS42L4VAMKM======")
// 对文件进行 md5 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToBase64Bytes() // []byte("AU8D+QJeqBqKDpc0vlQMUw==")
```

#### Hmac-md5 加密
```go
// 对字符串进行 hmac-md5 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToHexString() // 4790626a275f776956386e5a3ea7b726
// 对字符串进行 hmac-md5 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToBase32String() // I6IGE2RHL53WSVRYNZND5J5XEY======
// 对字符串进行 hmac-md5 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToBase64String() // R5Biaidfd2lWOG5aPqe3Jg==

// 对字节切片进行 hmac-md5 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToHexBytes() // []byte("4790626a275f776956386e5a3ea7b726")
// 对字节切片进行 hmac-md5 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase32Bytes() // []byte("I6IGE2RHL53WSVRYNZND5J5XEY======")
// 对字节切片进行 hmac-md5 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase64Bytes() // []byte("R5Biaidfd2lWOG5aPqe3Jg==")
```

#### Sha1 加密
```go
// 对字符串进行 sha1 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha1().ToHexString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// 对字符串进行 sha1 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").BySha1().ToBase32String() // FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN
// 对字符串进行 sha1 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha1().ToBase64String() // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// 对字节切片进行 sha1 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToHexBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// 对字节切片进行 sha1 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase32Bytes() // []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")
// 对字节切片进行 sha1 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase64Bytes() // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

#### Hmac-sha1 加密
```go
// 对字符串进行 hmac-sha1 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToHexString() // 91c103ef93ba7420902b0d1bf0903251c94b4a62
// 对字符串进行 hmac-sha1 加密，输出经过 base22 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToBase32String() // SHAQH34TXJ2CBEBLBUN7BEBSKHEUWSTC
// 对字符串进行 hmac-sha1 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToBase64String() // kcED75O6dCCQKw0b8JAyUclLSmI=

// 对字节切片进行 hmac-sha1 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToHexBytes() // []byte("91c103ef93ba7420902b0d1bf0903251c94b4a62")
// 对字节切片进行 hmac-sha1 加密，输出经过 base22 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase32Bytes() // []byte("SHAQH34TXJ2CBEBLBUN7BEBSKHEUWSTC")
// 对字节切片进行 hmac-sha1 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase64Bytes() // []byte("kcED75O6dCCQKw0b8JAyUclLSmI=")
```

#### Sha224 加密
```go
// 对字符串进行 sha224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha224().ToHexString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// 对字符串进行 sha224 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").BySha224().ToBase32String() // F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===
// 对字符串进行 sha224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha224().ToBase64String() // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// 对字节切片进行 sha224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToHexBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// 对字节切片进行 sha224 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase32Bytes() // []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")
// 对字节切片进行 sha224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase64Bytes() // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

#### Hmac-sha224 加密
```go
// 对字符串进行 hmac-sha224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToHexString() // e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec
// 对字符串进行 hmac-sha224 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToBase32String() // 4FNZ4WT6ZSY7C7OIDXAHZEE2REMTNXBUFHOA3FAKZS6OY===
// 对字符串进行 hmac-sha224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToBase64String() // 4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A====

// 对字节切片进行 hmac-sha224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToHexBytes() // []byte("e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec")
// 对字节切片进行 hmac-sha224 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase32Bytes() // []byte("4FNZ4WT6ZSY7C7OIDXAHZEE2REMTNXBUFHOA3FAKZS6OY===")
// 对字节切片进行 hmac-sha224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase64Bytes() // []byte("4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A==")
```

#### Sha256 加密
```go
// 对字符串进行 sha256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha256().ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字符串进行 sha256 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").BySha256().ToBase32String() // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// 对字符串进行 sha256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha256().ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字节切片进行 sha256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字节切片进行 sha256 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase32Bytes() // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// 对字节切片进行 sha256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

#### Hmac-sha256 加密
```go
// 对字符串进行 hmac-sha256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字符串进行 hmac-sha256 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToBase32String() // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// 对字符串进行 hmac-sha256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字节切片进行 hmac-sha256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字节切片进行 hmac-sha256 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase32Bytes() // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// 对字节切片进行 hmac-sha256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

#### Sha384 加密
```go
// 对字符串进行 sha384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha384().ToHexString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// 对字符串进行 sha384 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").BySha384().ToBase32String() // 7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===
// 对字符串进行 sha384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha384().ToBase64String() // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// 对字节切片进行 sha384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToHexBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// 对字节切片进行 sha384 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase32Bytes() // []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")
// 对字节切片进行 sha384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase64Bytes() // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

#### Hmac-sha384 加密
```go
// 对字符串进行 hmac-sha384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToHexString() // 421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8
// 对字符串进行 hmac-sha384 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToBase32String() // IIP4VJ2AEFVDDO6ND6DPEIJOBRUKUSYVNKHLYKXFLM7HLRHOAUE6UAZFUBLQVZZZABVWDWI5QF76Q===
// 对字符串进行 hmac-sha384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToBase64String() // Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o

// 对字节切片进行 hmac-sha384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToHexBytes() // []byte("421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8")
// 对字节切片进行 hmac-sha384 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase32Bytes() // []byte("IIP4VJ2AEFVDDO6ND6DPEIJOBRUKUSYVNKHLYKXFLM7HLRHOAUE6UAZFUBLQVZZZABVWDWI5QF76Q===")
// 对字节切片进行 hmac-sha384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase64Bytes() // []byte("Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o")
```

#### Sha512 加密
```go
// 对字符串进行 sha512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToHexBytes() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// 对字符串进行 sha512 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToBase32String() // GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=
// 对字符串进行 sha512 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToBase64String() // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// 对字节切片进行 sha512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToHexBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// 对字节切片进行 sha512 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase32Bytes() // []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")
// 对字节切片进行 sha512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase64Bytes() // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")
```

#### Hmac-sha512 加密
```go
// 对字符串进行 hmac-sha512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToHexString() // d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1
// 对字符串进行 hmac-sha512 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToBase32String() // 3FY3PEF3YKSKZAIGFO77VRUTZHBDJOXBO3EPV5PDATN5WFJQGKUCN4JDKOLEWSSPXB5L5TJNYI3WHCTDBS5NKSTLSSY7N325LYUDLUI=
// 对字符串进行 hmac-sha512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToBase64String() // 2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==

// 对字节切片进行 hmac-sha512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToHexBytes() // []byte("d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1")
// 对字节切片进行 hmac-sha512 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase32Bytes() // []byte("3FY3PEF3YKSKZAIGFO77VRUTZHBDJOXBO3EPV5PDATN5WFJQGKUCN4JDKOLEWSSPXB5L5TJNYI3WHCTDBS5NKSTLSSY7N325LYUDLUI=")
// 对字节切片进行 hmac-sha512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase64Bytes() // []byte("2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==")
```

#### Rc4 加密
```go
// 对字符串进行 rc4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToHexString() // eba154b4cb5a9038dbbf9d
// 对字符串进行 rc4 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToBase32String() // 5OQVJNGLLKIDRW57TU======
// 对字符串进行 rc4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToBase64String() // 66FUtMtakDjbv50=

// 对字节切片进行 rc4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToHexBytes() // []byte("eba154b4cb5a9038dbbf9d")
// 对字节切片进行 rc4 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToBase32Bytes() // []byte("5OQVJNGLLKIDRW57TU======")
// 对字节切片进行 rc4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToBase64Bytes() // []byte("66FUtMtakDjbv50=")
```

#### Aes 加密、解密
```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("1234567887654321") // key 长度必须是 16、24 或 32
cipher.SetIV("1234567887654321") // iv 长度必须是 16、24 或 32

// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToHexString() // 65d823bdf1c581a1ded1cba42e03ca52
// 对经过 hex 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromHexString("65d823bdf1c581a1ded1cba42e03ca52").ByAes(cipher).ToString() // hello world

// 对字符串进行 aes 加密，输出经过 base32 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase32String() // MXMCHPPRYWA2DXWRZOSC4A6KKI======
// 对经过 base32 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromBase32String("MXMCHPPRYWA2DXWRZOSC4A6KKI======").ByAes(cipher).ToString() // hello world

// 对字符串进行 aes 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase64String() // ZdgjvfHFgaHe0cukLgPKUg==
// 对经过 base64 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromBase64String("ZdgjvfHFgaHe0cukLgPKUg==").ByAes(cipher).ToString() // hello world

// 对字节切片进行 aes 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToHexBytes() // []byte("65d823bdf1c581a1ded1cba42e03ca52")
// 对经过 hex 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromHexBytes([]byte("65d823bdf1c581a1ded1cba42e03ca52")).ByAes(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 aes 加密，输出经过 base32 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase32Bytes() // []byte("MXMCHPPRYWA2DXWRZOSC4A6KKI======")
// 对经过 base32 编码的字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromBase32Bytes([]byte("MXMCHPPRYWA2DXWRZOSC4A6KKI======")).ByAes(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 aes 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase64Bytes() // []byte("ZdgjvfHFgaHe0cukLgPKUg==")
// 对经过 base64 编码的字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("ZdgjvfHFgaHe0cukLgPKUg==")).ByAes(cipher).ToBytes() // []byte("hello world")
```

### 错误处理
> 如果有多个错误发生，只返回第一个错误，前一个错误排除后才返回下一个错误

```go
e := dongle.Encrypy.FromFile("./demo.txt").ByMd5()
if e.Error != nil {
    // 错误处理...
    log.Fatal(e.Error)
}
fmt.Println(e.ToString())
// 输出
invalid file "./demo.txt", please make sure the file exists
```

### 功能清单
- [x] Hex 编码、解码
- [x] Base32 编码、解码
- [ ] Base58 编码、解码
- [x] Base64 编码、解码
- [x] Base64URL 编码、解码
- [x] SafeURL 编码、解码
- [x] Md4 加密
- [x] Hmac-md4 加密
- [x] Md5 加密
- [x] Hmac-md5 加密
- [x] Sha1 加密
- [x] Hmac-sha1 加密
- [x] Sha224 加密
- [x] Hmac-sha224 加密
- [x] Sha256 加密
- [x] Hmac-sha256 加密
- [x] Sha384 加密
- [x] Hmac-sha384 加密
- [x] Sha512 加密
- [x] Hmac-sha512 加密
- [ ] Rc2 加密
- [x] Rc4 加密
- [ ] Rc5 加密
- [ ] Rc6 加密
- [ ] Sm2 加密
- [ ] Sm3 加密
- [ ] Sm4 加密
- [x] AES-CBC-NoPadding 加密、解密
- [x] AES-CBC-ZeroPadding 加密、解密
- [x] AES-CBC-PKCS5Padding 加密、解密
- [x] AES-CBC-PKCS7Padding 加密、解密
- [x] AES-CTR-NoPadding 加密、解密
- [x] AES-CTR-ZeroPadding 加密、解密
- [x] AES-CTR-PKCS5Padding 加密、解密
- [x] AES-CTR-PKCS7Padding 加密、解密
- [x] AES-CFB-NoPadding 加密、解密
- [x] AES-CFB-ZeroPadding 加密、解密
- [x] AES-CFB-PKCS5Padding 加密、解密
- [x] AES-CFB-PKCS7Padding 加密、解密
- [x] AES-OFB-NoPadding 加密、解密
- [x] AES-OFB-ZeroPadding 加密、解密
- [x] AES-OFB-PKCS5Padding 加密、解密
- [x] AES-OFB-PKCS7Padding 加密、解密
- [ ] AES-GCM-NoPadding 加密、解密
- [ ] AES-GCM-ZeroPadding 加密、解密
- [ ] AES-GCM-PKCS5Padding 加密、解密
- [ ] AES-GCM-PKCS7Padding 加密、解密
- [ ] DES 加密、解密
- [ ] 3DES 加密、解密
- [ ] RSA 加密、解密
- [ ] DSA 加密、解密

### 参考项目
* [brix/crypto-js](https://github.com/brix/crypto-js)
* [nodejs/crypto](https://nodejs.org/api/crypto.html)
* [paramiko/paramiko](https://github.com/paramiko/paramiko)
* [cossacklabs/themis](https://github.com/cossacklabs/themis)

### 赞助
`Dongle` 是一个非商业开源项目, 如果你想支持 `Dongle`, 你可以为开发者 [购买一杯咖啡](https://www.gouguoyin.cn/zanzhu.html)

### 致谢
`Dongle`已获取免费的 JetBrains 开源许可证，在此表示感谢

<a href="https://www.jetbrains.com"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" height="100" alt="JetBrains"/></a>