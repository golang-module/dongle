# dongle  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/dongle.svg)](https://github.com/golang-module/dongle/releases)
[![Go Build](https://github.com/golang-module/dongle/actions/workflows/test.yml/badge.svg)](https://github.com/golang-module/dongle/actions)
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

### 用法示例
#### 编码&解码

##### Hex 编码、解码
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

##### Base16 编码、解码

```go
// 对字符串进行 base16 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase16().ToString() // 68656c6c6f20776f726c64
// 对字符串进行 base16 解码，输出字符串
dongle.Decode.FromString("68656c6c6f20776f726c64").ByBase16().ToString() // hello world

// 对字节切片进行 base16 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase16().ToBytes() // []byte("68656c6c6f20776f726c64")
// 对字节切片进行 base16 解码，输出字节切片
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByBase16().ToBytes() // []byte("hello world")
```

##### Base32 编码、解码

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

##### Base58 编码、解码

```go
// 对字符进行 base58 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase58().ToString() // StV1DL6CwTryKyV
// 对字符串进行 base58 解码，输出字符串
dongle.Decode.FromString("StV1DL6CwTryKyV").ByBase58().ToString() // hello world

// 对字节切片进行 base58 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase58().ToBytes() // []byte("StV1DL6CwTryKyV")
// 对字节切片进行 base58 解码，输出字节切片
dongle.Decode.FromBytes([]byte("StV1DL6CwTryKyV")).ByBase58().ToBytes() // []byte("hello world")
```

##### Base62 编码、解码

```go
// 对字符进行 base62 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase62().ToString() // AAwf93rvy4aWQVw
// 对字符串进行 base62 解码，输出字符串
dongle.Decode.FromString("AAwf93rvy4aWQVw").ByBase62().ToString() // hello world

// 对字节切片进行 base62 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase62().ToBytes() // []byte("AAwf93rvy4aWQVw")
// 对字节切片进行 base62 解码，输出字节切片
dongle.Decode.FromBytes([]byte("AAwf93rvy4aWQVw")).ByBase62().ToBytes() // []byte("hello world")
```

##### Base64 编码、解码

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

##### Base64URL 编码、解码
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

##### Base85 编码、解码

```go
// 对字符串进行 base85 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase85().ToString() // BOu!rD]j7BEbo7
// 对字符串进行 base85 解码，输出字符串
dongle.Decode.FromString("BOu!rD]j7BEbo7").ByBase85().ToString() // hello world

// 对字节切片进行 base85 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase85().ToBytes() // []byte("BOu!rD]j7BEbo7")
// 对字节切片进行 base85 解码，输出字节切片
dongle.Decode.FromBytes([]byte("BOu!rD]j7BEbo7")).ByBase85().ToBytes() // []byte("hello world")
```

##### Base91 编码、解码

```go
// 对字符串进行 base91 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase91().ToString() // TPwJh>Io2Tv!lE
// 对字符串进行 base91 解码，输出字符串
dongle.Decode.FromString("TPwJh>Io2Tv!lE").ByBase91().ToString() // hello world

// 对字节切片进行 base91 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase91().ToBytes() // []byte("TPwJh>Io2Tv!lE")
// 对字节切片进行 base91 解码，输出字节切片
dongle.Decode.FromBytes([]byte("BOu!rD]j7BEbo7")).ByBase91().ToBytes() // []byte("hello world")
```

##### Base100 编码、解码

```go
// 对字符串进行 base100 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase100().ToString() // 👟👜👣👣👦🐗👮👦👩👣👛
// 对字符串进行 base100 解码，输出字符串
dongle.Decode.FromString("👟👜👣👣👦🐗👮👦👩👣👛").ByBase100().ToString() // hello world

// 对字节切片进行 base100 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase100().ToBytes() // []byte("👟👜👣👣👦🐗👮👦👩👣👛")
// 对字节切片进行 base100 解码，输出字节切片
dongle.Decode.FromBytes([]byte("👟👜👣👣👦🐗👮👦👩👣👛")).ByBase100().ToBytes() // []byte("hello world")
```

##### SafeURL 编码、解码
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

#### 加密&解密

##### Md4 加密

```go
// 对字符串进行 md4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd4().ToHexString() // aa010fbc1d14c795d86ef98c95479d17
// 对字符串进行 md4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd4().ToBase64String() // qgEPvB0Ux5XYbvmMlUedFw==

// 对字节切片进行 md4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToHexBytes() // []byte("aa010fbc1d14c795d86ef98c95479d17")
// 对字节切片进行 md4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase64Bytes() // []byte("qgEPvB0Ux5XYbvmMlUedFw==")

```

##### Hmac-md4 加密
```go
// 对字符串进行 hmac-md4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToHexString() // 7a9df5247cbf76a8bc17c9c4f5a75b6b
// 对字符串进行 hmac-md4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToBase64String() // ep31JHy/dqi8F8nE9adbaw==

// 对字节切片进行 hmac-md4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToHexBytes() // []byte("7a9df5247cbf76a8bc17c9c4f5a75b6b")
// 对字节切片进行 hmac-md4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase64Bytes() // []byte("ep31JHy/dqi8F8nE9adbaw==")
```

##### Md5 加密
```go
// 对字符串进行 md5 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
// 对字符串进行 md5 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd5().ToBase64String() // XrY7u+Ae7tCTyyK7j1rNww==

// 对字节切片进行 md5 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToHexBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// 对字节切片进行 md5 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase64Bytes() // []byte("XrY7u+Ae7tCTyyK7j1rNww==")
```

##### Hmac-md5 加密
```go
// 对字符串进行 hmac-md5 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToHexString() // 4790626a275f776956386e5a3ea7b726
// 对字符串进行 hmac-md5 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToBase64String() // R5Biaidfd2lWOG5aPqe3Jg==

// 对字节切片进行 hmac-md5 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToHexBytes() // []byte("4790626a275f776956386e5a3ea7b726")
// 对字节切片进行 hmac-md5 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase64Bytes() // []byte("R5Biaidfd2lWOG5aPqe3Jg==")
```

##### Sha1 加密
```go
// 对字符串进行 sha1 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha1().ToHexString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// 对字符串进行 sha1 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha1().ToBase64String() // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// 对字节切片进行 sha1 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToHexBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// 对字节切片进行 sha1 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase64Bytes() // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

##### Hmac-sha1 加密
```go
// 对字符串进行 hmac-sha1 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToHexString() // 91c103ef93ba7420902b0d1bf0903251c94b4a62
// 对字符串进行 hmac-sha1 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToBase64String() // kcED75O6dCCQKw0b8JAyUclLSmI=

// 对字节切片进行 hmac-sha1 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToHexBytes() // []byte("91c103ef93ba7420902b0d1bf0903251c94b4a62")
// 对字节切片进行 hmac-sha1 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase64Bytes() // []byte("kcED75O6dCCQKw0b8JAyUclLSmI=")
```

##### Sha224 加密
```go
// 对字符串进行 sha224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha224().ToHexString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// 对字符串进行 sha224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha224().ToBase64String() // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// 对字节切片进行 sha224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToHexBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// 对字节切片进行 sha224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase64Bytes() // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

##### Hmac-sha224 加密
```go
// 对字符串进行 hmac-sha224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToHexString() // e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec
// 对字符串进行 hmac-sha224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToBase64String() // 4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A====

// 对字节切片进行 hmac-sha224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToHexBytes() // []byte("e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec")
// 对字节切片进行 hmac-sha224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase64Bytes() // []byte("4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A==")
```

##### Sha256 加密
```go
// 对字符串进行 sha256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha256().ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字符串进行 sha256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha256().ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字节切片进行 sha256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字节切片进行 sha256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Hmac-sha256 加密
```go
// 对字符串进行 hmac-sha256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字符串进行 hmac-sha256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字节切片进行 hmac-sha256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字节切片进行 hmac-sha256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Sha384 加密
```go
// 对字符串进行 sha384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha384().ToHexString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// 对字符串进行 sha384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySha384().ToBase64String() // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// 对字节切片进行 sha384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToHexBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// 对字节切片进行 sha384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase64Bytes() // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

##### Hmac-sha384 加密
```go
// 对字符串进行 hmac-sha384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToHexString() // 421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8
// 对字符串进行 hmac-sha384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToBase64String() // Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o

// 对字节切片进行 hmac-sha384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToHexBytes() // []byte("421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8")
// 对字节切片进行 hmac-sha384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase64Bytes() // []byte("Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o")
```

##### Sha512 加密
```go
// 对字符串进行 sha512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToHexBytes() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// 对字符串进行 sha512 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToBase64String() // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// 对字节切片进行 sha512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToHexBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// 对字节切片进行 sha512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase64Bytes() // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")
```

##### Hmac-sha512 加密
```go
// 对字符串进行 hmac-sha512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToHexString() // d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1
// 对字符串进行 hmac-sha512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToBase64String() // 2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==

// 对字节切片进行 hmac-sha512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToHexBytes() // []byte("d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1")
// 对字节切片进行 hmac-sha512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase64Bytes() // []byte("2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==")
```

##### Rc4 加密
```go
// 对字符串进行 rc4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToHexString() // eba154b4cb5a9038dbbf9d
// 对字符串进行 rc4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToBase64String() // 66FUtMtakDjbv50=

// 对字节切片进行 rc4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToHexBytes() // []byte("eba154b4cb5a9038dbbf9d")
// 对字节切片进行 rc4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToBase64Bytes() // []byte("66FUtMtakDjbv50=")
```

##### Rsa 加密、解密

```go
pkcs1PublicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCtdjE3fOKpAmc6eIi1I/GElIJW
hTLBZb/SpSC+Pl7bONDyt/sG7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuA
bOKi7TBgojusBLsrddCrGn08gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU
0qcFQKXvfJaBTUX6gwIDAQAB
-----END PUBLIC KEY-----`

pkcs8PublicKey := `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`

pkcs1PrivateKey := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCtdjE3fOKpAmc6eIi1I/GElIJWhTLBZb/SpSC+Pl7bONDyt/sG
7FzxMHLoxnlJnYDHmYiu0iy0s/EHhNL+1bPnBiuAbOKi7TBgojusBLsrddCrGn08
gZLRZ3ZP/oAS6+wf2/vfh2jCCKCVHbBZLHQIN3MU0qcFQKXvfJaBTUX6gwIDAQAB
AoGAFwAfEo56t5JcAcLNzccQVVYj2jkbO820G8hNiSxYA5WLD0QaAxcAU/Lqqbb3
ii1aUB0ppJS13NgnU6nnGGdZzUYBG1Hai6EkVyCGrI4amQ93AaVdKncL8gJ4RZAm
YzPPUwSMEESsu24pS1NF1G1Y8C+28b/Wr0oqOsCvL6PhsMECQQDwsPJJoWRx7ZJw
E1K5KLT0cXKyrIpyXY3I6tyA5imCzOzccf3d1vDgB0L9sdSO7bG3ceSwpAeiWEbg
5jGZemPzAkEAuH6U4pEI4AMbWnatpK55Rc235NDgmT3VyIuRaKC02YXAZ+jznFep
XMd4DTli4R9r3j2YVhUpyDVbdQpFH98DMQJAQpOvcU6DSkA80WOG7lCkPTlkUKgJ
Y7kdDwZoF/+SW+vzWMbvQf3CgzV/Ak2+TgrRrbyDVZkJw45HjM4fyiRgoQJBALH/
/qlxgPyQQs3O/s2KQBsm1auAE5IF5MLuVUZ69sF/mBko2hEXSqHnGV645TuKU0pC
Zz12ga9WO3z6gaK0SaECQQDah1pKt9ViBBy4USXK3OWXEloHuTwmyr9AbLqqI5tQ
2eNuH0NkuJYQmnXmHLbKOELoYocldEBXmkzPXSN+X9kV
-----END RSA PRIVATE KEY-----`

pkcs8PrivateKey := `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKrNk1r1Wtx7DJTr
AOhXtj2QAepfVUrQHdFvoY2ZB7jMsR9x7txVNoutzhUZMqXfm0AMbVxEeq1obhL9
a22mIZkGHEnLgyk5dvp4g+JUuyfaUv6smjld1tKveDKPEQ5BD3uKG3DiUN3nAyjh
sg67DUu0x7McLWi62UzrH78EHQFJAgMBAAECgYAeo3nHWzPNURVUsUMcan96U5bE
YA2AugxfQVMNf2HvOGidZ2adh3udWrQY/MglERNcTd5gKriG2rDEH0liBecIrNKs
BL4lV+qHEGRUcnDDdtUBdGInEU8lve5keDgmX+/huXSRJ+3tYA5u9j+32RquVczv
Idtb5XnBLUl61k0osQJBAON5+eJjtw6xpn+pveU92BSHvaJYVyrLHwUjR07aNKb7
GlGVM3MGf1FCa8WQUo9uUzYxGLtg5Qf3sqwOrwPd5UsCQQDAOF/zWqGuY3HfV/1w
giXiWp8rc+S8tanMj5M37QQbYW5YLjUmJImoklVahv3qlgLZdEN5ZSueM5jfoSFt
Nts7AkBKoRDvSiGbi4MBbTHkzLZgfewkH/FxE7S4nctePk553fXTgCyh9ya8BRuQ
dHnxnpNkOxVPHEnnpEcVFbgrf5gjAkB7KmRI4VTiEfRgINhTJAG0VU7SH/N7+4cu
fPzfA+7ywG5c8Fa79wOB0SoB1KeUjcSLo5Ssj2fwea1F9dAeU90LAkBJQFofveaD
a3YlN4EQZOcCvJKmg7xwWuGxFVTZDVVEws7UCQbEOEEXZrNd9x0IF5kpPLR+rxua
RPgUNaDGIh5o
-----END PRIVATE KEY-----`

// 使用公钥对字符串进行 rsa 加密
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs1PublicKey)
// 使用 pkcs1 格式密钥对经过 hex 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs1PrivateKey, dongle.PKCS1).ToString() // hello world
// 使用 pkcs1 格式密钥对经过 base64 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs1PrivateKey, dongle.PKCS1).ToString() // hello world

// 使用公钥对字节切片进行 rsa 加密
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa([]byte(pkcs1PublicKey))
// 使用 pkcs1 格式密钥对经过 hex 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa([]byte(pkcs1PrivateKey), dongle.PKCS1).ToByte() // []bytes("hello world)
// 使用 pkcs1 格式密钥对经过 base64 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa([]byte(pkcs1PrivateKey), dongle.PKCS1).ToByte() // []bytes("hello world)

// 使用公钥对字符串进行 rsa 加密，输出经过 hex 编码的字符串
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs8PublicKey)
// 使用 pkcs8 格式密钥对经过 hex 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToString() // hello world
// 使用 pkcs8 格式密钥对经过 base64 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToString() // hello world

// 使用公钥对字节切片进行 rsa 加密，输出经过 hex 编码的字节切片
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs8PublicKey)
// 使用 pkcs8 格式密钥对经过 hex 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromHexBytes(cipherText).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToByte() // []bytes("hello world)
// 使用 pkcs8 格式密钥对经过 base64 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs8PrivateKey, dongle.PKCS8).ToByte() // []bytes("hello world)
```

##### Aes 加密、解密

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、CFB、OFB、CTR、ECB、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("0123456789abcdef") // key 长度必须是 16、24 或 32 字节
cipher.SetIV("0123456789abcdef") // iv 长度必须是 16 字节，ECB 模式不需要设置 iv

// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToHexString() // c1e9b4529aac9793010f4677f6358efe
// 对经过 hex 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByAes(cipher).ToString() // hello world

// 对字符串进行 aes 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase64String() // wem0Upqsl5MBD0Z39jWO/g==
// 对经过 base64 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromBase64String("wem0Upqsl5MBD0Z39jWO/g==").ByAes(cipher).ToString() // hello world

// 对字节切片进行 aes 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToHexBytes() // []byte("c1e9b4529aac9793010f4677f6358efe")
// 对经过 hex 编码的字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromHexBytes([]byte("c1e9b4529aac9793010f4677f6358efe")).ByAes(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 aes 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase64Bytes() // []byte("ZdgjvfHFgaHe0cukLgPKUg==")
// 对经过 base64 编码的字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("wem0Upqsl5MBD0Z39jWO/g==")).ByAes(cipher).ToBytes() // []byte("hello world")
```

##### Des 加密、解密

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("12345678") // key 长度必须是 8 字节
cipher.SetIV("123456788") // iv 长度必须是 8 字节

// 对字符串进行 des 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// 对经过 hex 编码的字符串进行 des 解密，输出字符串
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(cipher).ToString() // hello world

// 对字符串进行 des 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// 对经过 base64 编码的字符串进行 des 解密，输出字符串
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").ByDes(cipher).ToString() // hello world

// 对字节切片进行 des 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// 对经过 hex 编码的字节切片进行 des 解密，输出字符串
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).ByDes(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 des 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// 对经过 base64 编码的字节切片进行 des 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).ByDes(cipher).ToBytes() // []byte("hello world")
```

##### 3Des 加密、解密

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC, CFB, CTR, ECB, OFB
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("123456781234567812345678") // key 长度必须是 24
cipher.SetIV("123456788") // iv 长度必须是 8

// 对字符串进行 3des 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// 对经过 hex 编码的字符串进行 3des 解密，输出字符串
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").By3Des(cipher).ToString() // hello world

// 对字符串进行 3des 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// 对经过 base64 编码的字符串进行 3des 解密，输出字符串
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").By3Des(cipher).ToString() // hello world

// 对字节切片进行 3des 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// 对经过 hex 编码的字节切片进行 3des 解密，输出字符串
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).By3Des(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 3des 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// 对经过 base64 编码的字节切片进行 3des 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).By3Des(cipher).ToBytes() // []byte("hello world")
```

##### Sm3 加密
```go
// 对字符串进行 sm3 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySm3().ToHexString() // 44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88
// 对字符串进行 sm3 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").BySm3().ToBase64String() // RPAGHmn6b9/CkMSUZUoF3AwFPaflxSuE75Op1n0//4g=

// 对字节切片进行 sm3 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySm3().ToHexBytes() // []byte("44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88")
// 对字节切片进行 sm3 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySm3().ToBase64Bytes() // []byte("RPAGHmn6b9/CkMSUZUoF3AwFPaflxSuE75Op1n0//4g=")
```

### 错误处理

> 如果有多个错误发生，只返回第一个错误，前一个错误排除后才返回下一个错误

```go
e := dongle.Encrypy.FromString("hello world").ByRsa("xxxx")
if e.Error != nil {
// 错误处理...
log.Fatal(e.Error)
}
fmt.Println(e.ToString())
// 输出
invalid public key, please make sure the public key is valid
```

### 功能清单

- [x] Hex 编码、解码
- [x] Base16 编码、解码
- [x] Base32 编码、解码
- [x] Base58 编码、解码
- [x] Base62 编码、解码
- [x] Base64 编码、解码
- [x] Base64URL 编码、解码
- [x] SafeURL 编码、解码
- [x] Base85 编码、解码
- [x] Base91 编码、解码
- [x] Base100 编码、解码
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
- [ ] Tea 加密、解密
- [ ] Xtea 加密、解密
- [x] Aes-ECB/CBC/CTR/CFB/OFB-NoPadding/ZeroPadding/PKCS5Padding/PKCS7Padding 加密、解密
- [x] Dde-ECB/CBC/CTR/CFB/OFB-NoPadding/ZeroPadding/PKCS5Padding/PKCS7Padding 加密、解密
- [x] 3des-ECB/CBC/CTR/CFB/OFB-NoPadding/ZeroPadding/PKCS5Padding/PKCS7Padding 加密、解密
- [x] Rsa-PKCS1Pem/PKCS8Pem 加密、解密
- [ ] Ecc 加密、解密
- [ ] Sm2 加密、解密
- [x] Sm3 加密
- [ ] Sm4 加密、解密
- [ ] Sm7 加密、解密
- [ ] Sm9 加密、解密
- [ ] Rsa 签名、验签
- [ ] Dsa 签名、验签

### 参考项目
* [javascript/crypto-js](https://github.com/brix/crypto-js)
* [nodejs/crypto](https://nodejs.org/api/crypto.html)
* [java/jasypt](https://github.com/jasypt/jasypt)
* [python/pycryptodome](https://github.com/Legrandin/pycryptodome)

### 在线网站
* [www.ssleye.com](https://www.ssleye.com/ssltool)
* [www.sojson.com](https://www.sojson.com/encrypt.html)
* [tool.chacuo.net](https://tool.chacuo.net/cryptaes)
* [www.oktools.net](https://oktools.net/aes)

### 赞助
`Dongle` 是一个非商业开源项目, 如果你想支持 `Dongle`, 你可以为开发者 [购买一杯咖啡](https://www.gouguoyin.com/zanzhu.html)

### 致谢

`Dongle`已获取免费的 JetBrains 开源许可证，在此表示感谢

<a href="https://www.jetbrains.com"><img src="https://github-oss.oss-cn-beijing.aliyuncs.com/jetbrains.png" height="100" alt="JetBrains"/></a>