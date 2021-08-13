# dongle  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/dongle.svg)](https://github.com/golang-module/dongle/releases)
[![Go Build](https://github.com/golang-module/dongle/actions/workflows/bulid.yml/badge.svg)](https://github.com/golang-module/dongle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/dongle)](https://goreportcard.com/report/github.com/golang-module/dongle)
[![codecov](https://codecov.io/gh/golang-module/openssl/branch/main/graph/badge.svg)](https://codecov.io/gh/golang-module/dongle)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/dongle)
![License](https://img.shields.io/github/license/golang-module/dongle)

简体中文 | [English](README.md)

一个轻量级、语义化、对开发者友好的 golang 编码解码、加密解密库

如果您觉得不错，请给个 star 吧

[github.com/golang-module/dongle](https://github.com/golang-module/dongle "github.com/golang-module/dongle")

[gitee.com/go-package/dongle](https://gitee.com/go-package/dongle "gitee.com/go-package/dongle")

#### 安装使用
```go
// 使用 github 库
go get -u github.com/golang-module/dongle

import (
    "github.com/golang-module/dongle"
)

// 使用 gitee 库
go get -u gitee.com/go-package/dongle

import (
    "gitee.com/go-package/dongle"
)
```

#### 编码&加密

##### Base32 编码
```go
// 对字符串进行 base32 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// 对字符串进行 base32 编码，输出字节切片
dongle.Encode.FromString("hello world").ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")

// 对字节切片进行 base32 编码，输出字符串
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// 对字节切片进行 base32 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")
```

##### Base64 编码
```go
// 对字符串进行 base64 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase64().ToString() // aGVsbG8gd29ybGQ=
// 对字符串进行 base64 编码，输出字节切片
dongle.Encode.FromString("hello world").ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")

// 对字节切片进行 base64 编码，输出字符串
dongle.Encode.FromBytes([]byte("hello world")).ByBase64().ToString() // aGVsbG8gd29ybGQ=
// 对字节切片进行 base64 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")
```

##### Hex 编码
#####
```go
// 对字符串进行 hex 编码，输出字符串
dongle.Encode.FromString("hello world").ByHex().ToString() // 68656c6c6f20776f726c64=
// 对字符串进行 hex 编码，输出字节切片
dongle.Encode.FromString("hello world").ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")

// 对字节切片进行 hex 编码，输出字符串
dongle.Encode.FromBytes([]byte("hello world")).ByHex().ToString() // 68656c6c6f20776f726c64
// 对字节切片进行 hex 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")
```

##### Md5 加密
```go
// 对字符串进行 md5 加密，输出 hex 编码字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
dongle.Encrypt.FromString("hello world").ByMd5().ToString("hex") // 5eb63bbbe01eeed093cb22bb8f5acdc3
// 对字符串进行 md5 加密，输出 base32 编码字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToString("base32") // L23DXO7AD3XNBE6LEK5Y6WWNYM======
// 对字符串进行 md5 加密，输出 base64 编码字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToString("base64") // XrY7u+Ae7tCTyyK7j1rNww==

// 对字符串进行 md5 加密，输出 hex 编码字节切片
dongle.Encrypt.FromString("hello world").ByMd5().ToBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
dongle.Encrypt.FromString("hello world").ByMd5().ToBytes("hex") // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// 对字符串进行 md5 加密，输出 base32 编码字节切片
dongle.Encrypt.FromString("hello world").ByMd5().ToBytes("base32") // []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")
// 对字符串进行 md5 加密，输出 base64 编码字节切片
dongle.Encrypt.FromString("hello world").ByMd5().ToBytes("base64") // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// 对字节切片进行 md5 加密，输出 hex 编码字符串
dongle.Encrypt.FromBytes("hello world").ByMd5().ToString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
dongle.Encrypt.FromBytes("hello world").ByMd5().ToString("hex") // 5eb63bbbe01eeed093cb22bb8f5acdc3
// 对字节切片进行 md5 加密，输出 base32 编码字符串
dongle.Encrypt.FromBytes("hello world").ByMd5().ToString("base32") // L23DXO7AD3XNBE6LEK5Y6WWNYM======
// 对字节切片进行 md5 加密，输出 base64 编码字符串
dongle.Encrypt.FromBytes("hello world").ByMd5().ToString("base64") // XrY7u+Ae7tCTyyK7j1rNww==

// 对字节切片进行 md5 加密，输出 hex 编码字节切片
dongle.Encrypt.FromBytes("hello world").ByMd5().ToBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
dongle.Encrypt.FromBytes("hello world").ByMd5().ToBytes("hex") // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// 对字节切片进行 md5 加密，输出 base32 编码字节切片
dongle.Encrypt.FromBytes("hello world").ByMd5().ToBytes("base32") // []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")
// 对字节切片进行 md5 加密，输出 base64 编码字节切片
dongle.Encrypt.FromBytes("hello world").ByMd5().ToBytes("base64") // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// 对文件进行 md5 加密，输出 hex 编码字符串
dongle.Encrypt.FromFile("./LICENSE")).ByMd5().ToString() // 014f03f9025ea81a8a0e9734be540c53
dongle.Encrypt.FromFile("./LICENSE")).ByMd5().ToString("hex") // 014f03f9025ea81a8a0e9734be540c53
// 对文件进行 md5 加密，输出 base32 编码字符串
dongle.Encrypt.FromFile("./LICENSE")).ByMd5().ToString("base32") // AFHQH6ICL2UBVCQOS42L4VAMKM======
// 对文件进行 md5 加密，输出 base64 编码字符串
dongle.Encrypt.FromFile("./LICENSE")).ByMd5().ToString("base64") // AU8D+QJeqBqKDpc0vlQMUw==

// 对文件进行 md5 加密，输出 hex 编码字节切片
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes() // []byte("014f03f9025ea81a8a0e9734be540c53")
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes("hex") // []byte("014f03f9025ea81a8a0e9734be540c53")
// 对文件进行 md5 加密，输出 base32 编码字节切片
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes("base32") // []byte("AFHQH6ICL2UBVCQOS42L4VAMKM======")
// 对文件进行 md5 加密，输出 base64 编码字节切片
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBytes("base64") // []byte("AU8D+QJeqBqKDpc0vlQMUw==")
```

##### Sha1 加密
```go
// 对字符串进行 sha1 加密，输出 hex 编码字符串
dongle.Encrypt.FromString("hello world").BySha1().ToString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
dongle.Encrypt.FromString("hello world").BySha1().ToString("hex") // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// 对字符串进行 sha1 加密，输出 base32 编码字符串
dongle.Encrypt.FromString("hello world").BySha1().ToString("base32") // FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN
// 对字符串进行 sha1 加密，输出 base64 编码字符串
dongle.Encrypt.FromString("hello world").BySha1().ToString("base64") // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// 对字符串进行 sha1 加密，输出 hex 编码字节切片
dongle.Encrypt.FromString("hello world").BySha1().ToBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
dongle.Encrypt.FromString("hello world").BySha1().ToBytes("hex") // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// 对字符串进行 sha1 加密，输出 base32 编码字节切片
dongle.Encrypt.FromString("hello world").BySha1().ToBytes("base32") // []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")
// 对字符串进行 sha1 加密，输出 base64 编码字节切片
dongle.Encrypt.FromString("hello world").BySha1().ToBytes("base64") // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")

// 对字节切片进行 sha1 加密，输出 hex 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha1().ToString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
dongle.Encrypt.FromBytes("hello world").BySha1().ToString("hex") // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// 对字节切片进行 sha1 加密，输出 base32 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha1().ToString("base32") // FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN
// 对字节切片进行 sha1 加密，输出 base64 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha1().ToString("base64") // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// 对字节切片进行 sha1 加密，输出 hex 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha1().ToBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
dongle.Encrypt.FromBytes("hello world").BySha1().ToBytes("hex") // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// 对字节切片进行 sha1 加密，输出 base32 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha1().ToBytes("base32") // []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")
// 对字节切片进行 sha1 加密，输出 base64 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha1().ToBytes("base64") // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

##### Sha224 加密
```go
// 对字符串进行 sha224 加密，输出 hex 编码字符串
dongle.Encrypt.FromString("hello world").BySha224().ToString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
dongle.Encrypt.FromString("hello world").BySha224().ToString("hex") // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// 对字符串进行 sha224 加密，输出 base32 编码字符串
dongle.Encrypt.FromString("hello world").BySha224().ToString("base32") // F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===
// 对字符串进行 sha224 加密，输出 base64 编码字符串
dongle.Encrypt.FromString("hello world").BySha224().ToString("base64") // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// 对字符串进行 sha224 加密，输出 hex 编码字节切片
dongle.Encrypt.FromString("hello world").BySha224().ToBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
dongle.Encrypt.FromString("hello world").BySha224().ToBytes("hex") // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// 对字符串进行 sha224 加密，输出 base32 编码字节切片
dongle.Encrypt.FromString("hello world").BySha224().ToBytes("base32") // []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")
// 对字符串进行 sha224 加密，输出 base64 编码字节切片
dongle.Encrypt.FromString("hello world").BySha224().ToBytes("base64") // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")

// 对字节切片进行 sha224 加密，输出 hex 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha224().ToString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
dongle.Encrypt.FromBytes("hello world").BySha224().ToString("hex") // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// 对字节切片进行 sha224 加密，输出 base32 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha224().ToString("base32") // F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===
// 对字节切片进行 sha224 加密，输出 base64 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha224().ToString("base64") // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// 对字节切片进行 sha224 加密，输出 hex 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha224().ToBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
dongle.Encrypt.FromBytes("hello world").BySha224().ToBytes("hex") // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// 对字节切片进行 sha224 加密，输出 base32 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha224().ToBytes("base32") // []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")
// 对字节切片进行 sha224 加密，输出 base64 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha224().ToBytes("base64") // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

##### Sha256 加密
```go
// 对字符串进行 sha256 加密，输出 hex 编码字符串
dongle.Encrypt.FromString("hello world").BySha256().ToString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
dongle.Encrypt.FromString("hello world").BySha256().ToString("hex") // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字符串进行 sha256 加密，输出 base32 编码字符串
dongle.Encrypt.FromString("hello world").BySha256().ToString("base32") // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// 对字符串进行 sha256 加密，输出 base64 编码字符串
dongle.Encrypt.FromString("hello world").BySha256().ToString("base64") // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字符串进行 sha256 加密，输出 hex 编码字节切片
dongle.Encrypt.FromString("hello world").BySha256().ToBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
dongle.Encrypt.FromString("hello world").BySha256().ToBytes("hex") // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字符串进行 sha256 加密，输出 base32 编码字节切片
dongle.Encrypt.FromString("hello world").BySha256().ToBytes("base32") // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// 对字符串进行 sha256 加密，输出 base64 编码字节切片
dongle.Encrypt.FromString("hello world").BySha256().ToBytes("base64") // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")

// 对字节切片进行 sha256 加密，输出 hex 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha256().ToString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
dongle.Encrypt.FromBytes("hello world").BySha256().ToString("hex") // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字节切片进行 sha256 加密，输出 base32 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha256().ToString("base32") // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// 对字节切片进行 sha256 加密，输出 base64 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha256().ToString("base64") // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字节切片进行 sha256 加密，输出 hex 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha256().ToBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
dongle.Encrypt.FromBytes("hello world").BySha256().ToBytes("hex") // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字节切片进行 sha256 加密，输出 base32 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha256().ToBytes("base32") // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// 对字节切片进行 sha256 加密，输出 base64 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha256().ToBytes("base64") // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Sha384 加密
```go
// 对字符串进行 sha384 加密，输出 hex 编码字符串
dongle.Encrypt.FromString("hello world").BySha384().ToString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
dongle.Encrypt.FromString("hello world").BySha384().ToString("hex") // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// 对字符串进行 sha384 加密，输出 base32 编码字符串
dongle.Encrypt.FromString("hello world").BySha384().ToString("base32") // 7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===
// 对字符串进行 sha384 加密，输出 base64 编码字符串
dongle.Encrypt.FromString("hello world").BySha384().ToString("base64") // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// 对字符串进行 sha384 加密，输出 hex 编码字节切片
dongle.Encrypt.FromString("hello world").BySha384().ToBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
dongle.Encrypt.FromString("hello world").BySha384().ToBytes("hex") // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// 对字符串进行 sha384 加密，输出 base32 编码字节切片
dongle.Encrypt.FromString("hello world").BySha384().ToBytes("base32") // []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")
// 对字符串进行 sha384 加密，输出 base64 编码字节切片
dongle.Encrypt.FromString("hello world").BySha384().ToBytes("base64") // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")

// 对字节切片进行 sha384 加密，输出 hex 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha384().ToString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
dongle.Encrypt.FromBytes("hello world").BySha384().ToString("hex") // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// 对字节切片进行 sha384 加密，输出 base32 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha384().ToString("base32") // 7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===
// 对字节切片进行 sha384 加密，输出 base64 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha384().ToString("base64") // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// 对字节切片进行 sha384 加密，输出 hex 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha384().ToBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
dongle.Encrypt.FromBytes("hello world").BySha384().ToBytes("hex") // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// 对字节切片进行 sha384 加密，输出 base32 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha384().ToBytes("base32") // []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")
// 对字节切片进行 sha384 加密，输出 base64 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha384().ToBytes("base64") // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

##### Sha512 加密
```go
// 对字符串进行 sha512 加密，输出 hex 编码字符串
dongle.Encrypt.FromString("hello world").BySha512().ToString() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
dongle.Encrypt.FromString("hello world").BySha512().ToString("hex") // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// 对字符串进行 sha512 加密，输出 base32 编码字符串
dongle.Encrypt.FromString("hello world").BySha512().ToString("base32") // GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=
// 对字符串进行 sha512 加密，输出 base64 编码字符串
dongle.Encrypt.FromString("hello world").BySha512().ToString("base64") // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// 对字符串进行 sha512 加密，输出 hex 编码字节切片
dongle.Encrypt.FromString("hello world").BySha512().ToBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
dongle.Encrypt.FromString("hello world").BySha512().ToBytes("hex") // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// 对字符串进行 sha512 加密，输出 base32 编码字节切片
dongle.Encrypt.FromString("hello world").BySha512().ToBytes("base32") // []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")
// 对字符串进行 sha512 加密，输出 base64 编码字节切片
dongle.Encrypt.FromString("hello world").BySha512().ToBytes("base64") // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")

// 对字节切片进行 sha512 加密，输出 hex 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha512().ToString() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
dongle.Encrypt.FromBytes("hello world").BySha512().ToString("hex") // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// 对字节切片进行 sha512 加密，输出 base32 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha512().ToString("base32") // GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=
// 对字节切片进行 sha512 加密，输出 base64 编码字符串
dongle.Encrypt.FromBytes("hello world").BySha512().ToString("base64") // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// 对字节切片进行 sha512 加密，输出 hex 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha512().ToBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
dongle.Encrypt.FromBytes("hello world").BySha512().ToBytes("hex") // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// 对字节切片进行 sha512 加密，输出 base32 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha512().ToBytes("base32") // []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")
// 对字节切片进行 sha512 加密，输出 base64 编码字节切片
dongle.Encrypt.FromBytes("hello world").BySha512().ToBytes("base64") // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")
```

#### 解码&解密
##### Base32 解码
```go
// 对字符串进行 base64 解码，输出字符串
dongle.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToString() // hello world
// 对字符串进行 base64 解码，输出字节切片
dongle.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToBytes() // []byte("hello world")

// 对字节切片进行 base64 解码，输出字符串
dongle.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToString() // hello world
// 对字节切片进行 base64 解码，输出字节切片
dongle.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToBytes() // []byte("hello world")
```

##### Base64 解码
```go
// 对字符串进行 base64 解码，输出字符串
dongle.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToString() // hello world
// 对字符串进行 base64 解码，输出字节切片
dongle.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToBytes() // []byte("hello world")

// 对字节切片进行 base64 解码，输出字符串
dongle.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToString() // hello world
// 对字节切片进行 base64 解码，输出字节切片
dongle.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToBytes() // []byte("hello world")
```

##### Hex 解码
```go
// 对字符串进行 hex 解码，输出字符串
dongle.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToString() // hello world
// 对字符串进行 hex 解码，输出字节切片
dongle.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToBytes() // []byte("hello world")

// 对字节切片进行 hex 解码，输出字符串
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToString() // hello world
// 对字节切片进行 hex 解码，输出字节切片
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToBytes() // []byte("hello world")
```

#### 错误处理
> 如果有多个错误发生，只返回第一个错误，前一个错误排除后才返回下一个错误

```go
e := dongle.Encrypy.FromFile("./demo.txt").ByMd5()
if e.Error != nil {
    // 错误处理...
    log.Fatal(c.Error)
}
fmt.Println(c.ToString())
// 输出
invalid file "./demo.txt", please make sure the file exists
```

#### 待做清单
 - [ ] AES 加密解密
 - [ ] DES 加密解密
 - [ ] 3DES 加密解密
 - [ ] RSA 加密解密
 - [ ] RC2 加密解密
 - [ ] RC4 加密解密
 - [ ] RC5 加密解密
 - [ ] RC6 加密解密
 - [ ] SM2 加密解密
 - [ ] SM3 加密解密
 - [ ] SM4 加密解密