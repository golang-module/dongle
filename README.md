# dongle  #
[![Carbon Release](https://img.shields.io/github/release/golang-module/dongle.svg)](https://github.com/golang-module/dongle/releases)
[![Go Build](https://github.com/golang-module/dongle/actions/workflows/bulid.yml/badge.svg)](https://github.com/golang-module/dongle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-module/dongle)](https://goreportcard.com/report/github.com/golang-module/dongle)
[![codecov](https://codecov.io/gh/golang-module/dongle/branch/main/graph/badge.svg)](https://codecov.io/gh/golang-module/dongle)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/golang-module/dongle)
![License](https://img.shields.io/github/license/golang-module/dongle)

English | [简体中文](README.cn.md)

### Introduction
A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption

`Dongle` has been included by [awesome-go](https://github.com/avelino/awesome-go#security "awesome-go"), if you think
it is helpful, please give me a star

[github.com/golang-module/dongle](https://github.com/golang-module/dongle "github.com/golang-module/dongle")

[gitee.com/golang-module/dongle](https://gitee.com/golang-module/dongle "gitee.com/golang-module/dongle")

### Installation
```go
// By github
go get -u github.com/golang-module/dongle

import (
    "github.com/golang-module/dongle"
)

// By gitee
go get -u gitee.com/golang-module/dongle

import (
"gitee.com/golang-module/dongle"
)
```

### Usage and example

#### Encode and decode

##### Encode and decode by hex

```go
// Encode by hex from string and output string
dongle.Encode.FromString("hello world").ByHex().ToString() // 68656c6c6f20776f726c64
// Decode by hex from string and output string
dongle.Decode.FromString("68656c6c6f20776f726c64").ByHex().ToString() // hello world

// Encode by hex from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByHex().ToBytes() // []byte("68656c6c6f20776f726c64")
// Decode by hex from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByHex().ToBytes() // []byte("hello world")
```

##### Encode and decode by base32

```go
// Encode by base32 from string and output string
dongle.Encode.FromString("hello world").ByBase32().ToString() // NBSWY3DPEB3W64TMMQ======
// Decode by base32 from string and output string
dongle.Decode.FromString("NBSWY3DPEB3W64TMMQ======").ByBase32().ToString() // hello world

// Encode by base32 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("NBSWY3DPEB3W64TMMQ======")
// Decode by base32 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("NBSWY3DPEB3W64TMMQ======")).ByBase32().ToBytes() // []byte("hello world")
```

##### Encode and decode by base58

```go
// Encode by base58 from string and output string
dongle.Encode.FromString("hello world").ByBase32().ToString() // StV1DL6CwTryKyV
// Decode by base58 from string and output string
dongle.Decode.FromString("StV1DL6CwTryKyV").ByBase32().ToString() // hello world

// Encode by base58 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase32().ToBytes() // []byte("StV1DL6CwTryKyV")
// Decode by base58 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("StV1DL6CwTryKyV")).ByBase32().ToBytes() // []byte("hello world")
```

##### Encode and decode by base64

```go
// Encode by base64 from string and output string
dongle.Encode.FromString("hello world").ByBase64().ToString() // aGVsbG8gd29ybGQ=
// Decode by base64 from string and output string
dongle.Decode.FromString("aGVsbG8gd29ybGQ=").ByBase64().ToString() // hello world

// Encode by base64 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase64().ToBytes() // []byte("aGVsbG8gd29ybGQ=")
// Decode by base64 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("aGVsbG8gd29ybGQ=")).ByBase64().ToBytes() // []byte("hello world")

```

##### Encode and decode by base64URL
```go
// Encode by base64 from url string and output string
dongle.Encode.FromString("www.gouguoyin.cn").ByBase64URL().ToString() // d3d3LmdvdWd1b3lpbi5jbg==
// Decode by base64 from string and output url string
dongle.Decode.FromString("d3d3LmdvdWd1b3lpbi5jbg==").ByBase64URL().ToString() // www.gouguoyin.cn

// Encode by base64 from url byte slice and output byte slice
dongle.Encode.FromBytes([]byte("www.gouguoyin.cn")).ByBase64URL().ToBytes() // []byte("d3d3LmdvdWd1b3lpbi5jbg==")
// Decode by base64 from byte slice and output url byte slice
dongle.Decode.FromBytes([]byte("d3d3LmdvdWd1b3lpbi5jbg==")).ByBase64URL().ToBytes() // []byte("www.gouguoyin.cn")
```

##### Encode and decode by SafeURL
```go
// Encode by escape from url string and output string
dongle.Encode.FromString("www.gouguoyin.cn?sex=男&age=18").BySafeURL().ToString() // www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18
// Decode by escape from string and output url string
dongle.Decode.FromString("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18").BySafeURL().ToString() // www.gouguoyin.cn?sex=男&age=18

// Encode by escape from url byte slice and output byte slice
dongle.Encode.FromBytes([]byte("www.gouguoyin.cn?sex=男&age=18")).BySafeURL().ToBytes() // []byte("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18")
// Decode by escape from byte slice and output url byte slice
dongle.Decode.FromBytes([]byte("www.gouguoyin.cn%3Fsex%3D%E7%94%B7%26age%3D18")).BySafeURL().ToBytes() // []byte("www.gouguoyin.cn?sex=男&age=18")
```

#### Encrypt and decrypt

##### Encrypt by md4

```go
// Encrypt by md4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToHexString() // aa010fbc1d14c795d86ef98c95479d17
// Encrypt by md4 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToBase32String() // VIAQ7PA5CTDZLWDO7GGJKR45C4======
// Encrypt by md4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToBase64String() // qgEPvB0Ux5XYbvmMlUedFw==

// Encrypt by md4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToHexBytes() // []byte("aa010fbc1d14c795d86ef98c95479d17")
// Encrypt by md4 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase32Bytes() // []byte("VIAQ7PA5CTDZLWDO7GGJKR45C4======")
// Encrypt by md4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase64Bytes() // []byte("qgEPvB0Ux5XYbvmMlUedFw==")

// Encrypt by md4 from file and output string with hex encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToHexString() // 1240c5c0fb26b585999357915c56b511
// Encrypt by md4 from file and output string with base32 encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToBase32String() // CJAMLQH3E22YLGMTK6IVYVVVCE======
// Encrypt by md4 from file and output string with base64 encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd4().ToBase64String() // EkDFwPsmtYWZk1eRXFa1EQ==

// Encrypt by md4 from file and output byte slice with hex encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToHexBytes() // []byte("1240c5c0fb26b585999357915c56b511")
// Encrypt by md4 from file and output byte slice with base32 encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToBase32Bytes() // []byte("CJAMLQH3E22YLGMTK6IVYVVVCE======")
// Encrypt by md4 from file and output byte slice with base64 encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd4().ToBase64Bytes() // []byte("EkDFwPsmtYWZk1eRXFa1EQ==")
```

##### Encrypt by hmac-md4
```go
// Encrypt by hmac-md4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToHexString() // 7a9df5247cbf76a8bc17c9c4f5a75b6b
// Encrypt by hmac-md4 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToBase32String() // PKO7KJD4X53KRPAXZHCPLJ23NM======
// Encrypt by hmac-md4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4("dongle").ToBase64String() // ep31JHy/dqi8F8nE9adbaw==

// Encrypt by hmac-md4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToHexBytes() // []byte("7a9df5247cbf76a8bc17c9c4f5a75b6b")
// Encrypt by hmac-md4 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase32Bytes() // []byte("PKO7KJD4X53KRPAXZHCPLJ23NM======")
// Encrypt by hmac-md4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase64Bytes() // []byte("ep31JHy/dqi8F8nE9adbaw==")
```

##### Encrypt by md5
```go
// Encrypt by md5 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
// Encrypt by md5 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToBase32String() // L23DXO7AD3XNBE6LEK5Y6WWNYM======
// Encrypt by md5 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToBase64String() // XrY7u+Ae7tCTyyK7j1rNww==

// Encrypt by md5 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToHexBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// Encrypt by md5 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase32Bytes() // []byte("L23DXO7AD3XNBE6LEK5Y6WWNYM======")
// Encrypt by md5 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase64Bytes() // []byte("XrY7u+Ae7tCTyyK7j1rNww==")

// Encrypt by md5 from file and output string with hex encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToHexString() // 014f03f9025ea81a8a0e9734be540c53
// Encrypt by md5 from file and output string with base32 encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBase32String() // AFHQH6ICL2UBVCQOS42L4VAMKM======
// Encrypt by md5 from file and output string with base64 encoding
dongle.Encrypt.FromFile("./LICENSE").ByMd5().ToBase64String() // AU8D+QJeqBqKDpc0vlQMUw==

// Encrypt by md5 from file and output byte slice with hex encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToHexBytes() // []byte("014f03f9025ea81a8a0e9734be540c53")
// Encrypt by md5 from file and output byte slice with base32 encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToBase32Bytes() // []byte("AFHQH6ICL2UBVCQOS42L4VAMKM======")
// Encrypt by md5 from file and output byte slice with base64 encoding
dongle.Encrypt.FromFile([]byte("./LICENSE")).ByMd5().ToBase64Bytes() // []byte("AU8D+QJeqBqKDpc0vlQMUw==")
```

##### Encrypt by hmac-md5
```go
// Encrypt by hmac-md5 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToHexString() // 4790626a275f776956386e5a3ea7b726
// Encrypt by hmac-md5 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToBase32String() // I6IGE2RHL53WSVRYNZND5J5XEY======
// Encrypt by hmac-md5 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5("dongle").ToBase64String() // R5Biaidfd2lWOG5aPqe3Jg==

// Encrypt by hmac-md5 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToHexBytes() // []byte("4790626a275f776956386e5a3ea7b726")
// Encrypt by hmac-md5 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase32Bytes() // []byte("I6IGE2RHL53WSVRYNZND5J5XEY======")
// Encrypt by hmac-md5 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase64Bytes() // []byte("R5Biaidfd2lWOG5aPqe3Jg==")
```

##### Encrypt by sha1
```go
// Encrypt by sha1 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha1().ToHexString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// Encrypt by sha1 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").BySha1().ToBase32String() // FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN
// Encrypt by sha1 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha1().ToBase64String() // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// Encrypt by sha1 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToHexBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// Encrypt by sha1 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase32Bytes() // []byte("FKXGYNOJJ7H3IFO35FPUBC445EPOQRXN")
// Encrypt by sha1 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase64Bytes() // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

##### Encrypt by hmac-sha1
```go
// Encrypt by hmac-sha1 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToHexString() // 91c103ef93ba7420902b0d1bf0903251c94b4a62
// Encrypt by hmac-sha1 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToBase32String() // SHAQH34TXJ2CBEBLBUN7BEBSKHEUWSTC
// Encrypt by hmac-sha1 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1("dongle").ToBase64String() // kcED75O6dCCQKw0b8JAyUclLSmI=

// Encrypt by hmac-sha1 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToHexBytes() // []byte("91c103ef93ba7420902b0d1bf0903251c94b4a62")
// Encrypt by hmac-sha1 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase32Bytes() // []byte("SHAQH34TXJ2CBEBLBUN7BEBSKHEUWSTC")
// Encrypt by hmac-sha1 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase64Bytes() // []byte("kcED75O6dCCQKw0b8JAyUclLSmI=")
```

##### Encrypt by sha224

```go
// Encrypt by sha224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha224().ToHexString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// Encrypt by sha224 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").BySha224().ToBase32String() // F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===
// Encrypt by sha224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha224().ToBase64String() // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// Encrypt by sha224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToHexBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// Encrypt by sha224 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase32Bytes() // []byte("F4CUO76CJO2PV36YMULRK3NP33HMIW4K2PHSKIVFMNMCW===")
// Encrypt by sha224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase64Bytes() // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

##### Encrypt by hmac-sha224
```go
// Encrypt by hmac-sha224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToHexString() // e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec
// Encrypt by hmac-sha224 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToBase32String() // 4FNZ4WT6ZSY7C7OIDXAHZEE2REMTNXBUFHOA3FAKZS6OY===
// Encrypt by hmac-sha224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224("dongle").ToBase64String() // 4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A====

// Encrypt by hmac-sha224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToHexBytes() // []byte("e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec")
// Encrypt by hmac-sha224 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase32Bytes() // []byte("4FNZ4WT6ZSY7C7OIDXAHZEE2REMTNXBUFHOA3FAKZS6OY===")
// Encrypt by hmac-sha224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase64Bytes() // []byte("4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A==")
```

##### Encrypt by sha256

```go
// Encrypt by sha256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha256().ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by sha256 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").BySha256().ToBase32String() // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// Encrypt by sha256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha256().ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by sha256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by sha256 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase32Bytes() // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// Encrypt by sha256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by hmac-sha256
```go
// Encrypt by hmac-sha256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by hmac-sha256 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToBase32String() // XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====
// Encrypt by hmac-sha256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256("dongle").ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by hmac-sha256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by hmac-sha256 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase32Bytes() // []byte("XFGSPOMTJU7ARJJOKLL5U7NL7LCIJ37DPJJYB3UQRD32ZYXPZXUQ====")
// Encrypt by hmac-sha256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by sha384

```go
// Encrypt by sha384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha384().ToHexString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// Encrypt by sha384 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").BySha384().ToBase32String() // 7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===
// Encrypt by sha384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha384().ToBase64String() // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// Encrypt by sha384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToHexBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// Encrypt by sha384 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase32Bytes() // []byte("7W6Y45NGP4U7OANE4BADQXROEOMGGA7KCARZEENPSB74XOBVPCZ6IF6LOHHGI3X5BAM53DAIRXQ32===")
// Encrypt by sha384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase64Bytes() // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

##### Encrypt by hmac-sha384
```go
// Encrypt by hmac-sha384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToHexString() // 421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8
// Encrypt by hmac-sha384 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToBase32String() // IIP4VJ2AEFVDDO6ND6DPEIJOBRUKUSYVNKHLYKXFLM7HLRHOAUE6UAZFUBLQVZZZABVWDWI5QF76Q===
// Encrypt by hmac-sha384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384("dongle").ToBase64String() // Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o

// Encrypt by hmac-sha384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToHexBytes() // []byte("421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8")
// Encrypt by hmac-sha384 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase32Bytes() // []byte("IIP4VJ2AEFVDDO6ND6DPEIJOBRUKUSYVNKHLYKXFLM7HLRHOAUE6UAZFUBLQVZZZABVWDWI5QF76Q===")
// Encrypt by hmac-sha384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase64Bytes() // []byte("Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o")
```

##### Encrypt by sha512
```go
// Encrypt by sha512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha512().ToHexBytes() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// Encrypt by sha512 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").BySha512().ToBase32String() // GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=
// Encrypt by sha512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha512().ToBase64String() // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==

// Encrypt by sha512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToHexBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// Encrypt by sha512 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase32Bytes() // []byte("GCPMYSE4CLLOWTGEB5IMSAXSWTIO257OKENHY6U3ZU6KQ3KM3BXZRHOTLPC76SMWODNDIJK3IWYM7WBQ5APWAXOPPXCVILUTV2ONO3Y=")
// Encrypt by sha512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase64Bytes() // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")
```

##### Encrypt by hmac-sha512
```go
// Encrypt by hmac-sha512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToHexString() // d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1
// Encrypt by hmac-sha512 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToBase32String() // 3FY3PEF3YKSKZAIGFO77VRUTZHBDJOXBO3EPV5PDATN5WFJQGKUCN4JDKOLEWSSPXB5L5TJNYI3WHCTDBS5NKSTLSSY7N325LYUDLUI=
// Encrypt by hmac-sha512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512("dongle").ToBase64String() // 2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==

// Encrypt by hmac-sha512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToHexBytes() // []byte("d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1")
// Encrypt by hmac-sha512 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase32Bytes() // []byte("3FY3PEF3YKSKZAIGFO77VRUTZHBDJOXBO3EPV5PDATN5WFJQGKUCN4JDKOLEWSSPXB5L5TJNYI3WHCTDBS5NKSTLSSY7N325LYUDLUI=")
// Encrypt by hmac-sha512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase64Bytes() // []byte("2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==")
```

##### Encrypt by rc4
```go
// Encrypt by rc4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToHexString() // eba154b4cb5a9038dbbf9d
// Encrypt by rc4 from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToBase32String() // 5OQVJNGLLKIDRW57TU======
// Encrypt by rc4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByRc4("dongle").ToBase64String() // 66FUtMtakDjbv50=

// Encrypt by rc4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToHexBytes() // []byte("eba154b4cb5a9038dbbf9d")
// Encrypt by rc4 from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToBase32Bytes() // []byte("5OQVJNGLLKIDRW57TU======")
// Encrypt by rc4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4("dongle").ToBase64Bytes() // []byte("66FUtMtakDjbv50=")
```

##### Encrypt and decrypt by aes
```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC, ECB, CFB, OFB, CTR, GCM
cipher.SetPadding(dongle.PKCS7) // No, Zero, PKCS5, PKCS7
cipher.SetKey("1234567887654321") // key must be 16, 24 or 32 bytes
cipher.SetIV("1234567887654321") // iv must be 16, 24 or 32 bytes

// Encrypt by aes from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToHexString() // 65d823bdf1c581a1ded1cba42e03ca52
// Decrypt by aes from string with hex encoding and output string
dongle.Decrypt.FromHexString("65d823bdf1c581a1ded1cba42e03ca52").ByAes(cipher).ToString() // hello world

// Encrypt by aes from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase32String() // MXMCHPPRYWA2DXWRZOSC4A6KKI======
// Decrypt by aes from string with base32 encoding and output string
dongle.Decrypt.FromBase32String("MXMCHPPRYWA2DXWRZOSC4A6KKI======").ByAes(cipher).ToString() // hello world

// Encrypt by aes from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase64String() // ZdgjvfHFgaHe0cukLgPKUg==
// Decrypt by aes from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("ZdgjvfHFgaHe0cukLgPKUg==").ByAes(cipher).ToString() // hello world

// Encrypt by aes from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToHexBytes() // []byte("65d823bdf1c581a1ded1cba42e03ca52")
// Decrypt by aes from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("65d823bdf1c581a1ded1cba42e03ca52")).ByAes(cipher).ToBytes() // []byte("hello world")

// Encrypt by aes from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase32Bytes() // []byte("MXMCHPPRYWA2DXWRZOSC4A6KKI======")
// Decrypt by aes from byte slice with base32 encoding and output byte slice
dongle.Decrypt.FromBase32Bytes([]byte("MXMCHPPRYWA2DXWRZOSC4A6KKI======")).ByAes(cipher).ToBytes() // []byte("hello world")

// Encrypt by aes from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase64Bytes() // []byte("ZdgjvfHFgaHe0cukLgPKUg==")
// Decrypt by aes from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("ZdgjvfHFgaHe0cukLgPKUg==")).ByAes(cipher).ToBytes() // []byte("hello world")

```

##### Encrypt and decrypt by des

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR、GCM
cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
cipher.SetKey("12345678")       // key must be 8 bytes
cipher.SetIV("123456788") // iv must be 8 bytes

// Encrypt by des from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// Decrypt by des from string with hex encoding and output string
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(cipher).ToString() // hello world

// Encrypt by des from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToBase32String() // BMVJF2A7WSOODJBSM2VMV2T3QE======
// Decrypt by des from string with base32 encoding and output string
dongle.Decrypt.FromBase32String("BMVJF2A7WSOODJBSM2VMV2T3QE======").ByDes(cipher).ToString() // hello world

// Encrypt by des from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// Decrypt by des from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").ByDes(cipher).ToString() // hello world

// Encrypt by des from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// Decrypt by des from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).ByDes(cipher).ToBytes() // []byte("hello world")

// Encrypt by des from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToBase32Bytes() // []byte("BMVJF2A7WSOODJBSM2VMV2T3QE======")
// Decrypt by des from byte slice with base32 encoding and output byte slice
dongle.Decrypt.FromBase32Bytes([]byte("BMVJF2A7WSOODJBSM2VMV2T3QE======")).ByDes(cipher).ToBytes() // []byte("hello world")

// Encrypt by des from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// Decrypt by des from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).ByDes(cipher).ToBytes() // []byte("hello world")
```

##### Encrypt and decrypt by 3des

```go
cipher := NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR、GCM
cipher.SetPadding(dongle.PKCS7)           // No、Zero、PKCS5、PKCS7
cipher.SetKey("123456781234567812345678") // key must be 24 bytes
cipher.SetIV("123456788") // iv must be 8 bytes

// Encrypt by 3des from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// Decrypt by des from string with hex encoding and output string
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from string and output string with base32 encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToBase32String() // BMVJF2A7WSOODJBSM2VMV2T3QE======
// Decrypt by 3des from string with base32 encoding and output string
dongle.Decrypt.FromBase32String("BMVJF2A7WSOODJBSM2VMV2T3QE======").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// Decrypt by 3des from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// Decrypt by 3des from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).By3Des(cipher).ToBytes() // []byte("hello world")

// Encrypt by 3des from byte slice and output byte slice with base32 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToBase32Bytes() // []byte("BMVJF2A7WSOODJBSM2VMV2T3QE======")
// Decrypt by des from byte slice with base32 encoding and output byte slice
dongle.Decrypt.FromBase32Bytes([]byte("BMVJF2A7WSOODJBSM2VMV2T3QE======")).By3Des(cipher).ToBytes() // []byte("hello world")

// Encrypt by 3des from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// Decrypt by 3des from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).By3Des(cipher).ToBytes() // []byte("hello world")
```

### Error handling

> If more than one error occurs, only the first error is returned

```go
e := dongle.Encrypy.FromFile("./demo.txt").ByMd5()
if e.Error != nil {
// Error handle...
log.Fatal(e.Error)
}
fmt.Println(e.ToHexString())
// Output
invalid file "./demo.txt", please make sure the file exists
```

### Feature list
- [x] Encoding and decoding by Hex
- [x] Encoding and decoding by Base32
- [x] Encoding and decoding by Base58
- [x] Encoding and decoding by Base64
- [x] Encoding and decoding by Base64URL
- [x] Encoding and decoding by SafeURL
- [x] Encryption by Md4
- [x] Encryption by Hmac-md4
- [x] Encryption by Md5
- [x] Encryption by Hmac-md5
- [x] Encryption by Sha1
- [x] Encryption by Hmac-sha1
- [x] Encryption by Sha224
- [x] Encryption by Hmac-sha224
- [x] Encryption by Sha256
- [x] Encryption by Hmac-sha256
- [x] Encryption by Sha384
- [x] Encryption by Hmac-sha384
- [x] Encryption by Sha512
- [x] Encryption by Hmac-sha512
- [ ] Encryption by Rc2
- [x] Encryption by Rc4
- [ ] Encryption by Rc5
- [ ] Encryption by Rc6
- [ ] Encryption and decryption by Sm2
- [ ] Encryption and decryption by Sm3
- [ ] Encryption and decryption by Sm4
- [ ] Encryption and decryption by RSA
- [ ] Encryption and decryption by DSA
- [x] Encryption and decryption by AES-CBC-NoPadding
- [x] Encryption and decryption by AES-CBC-ZeroPadding
- [x] Encryption and decryption by AES-CBC-PKCS5Padding
- [x] Encryption and decryption by AES-CBC-PKCS7Padding
- [x] Encryption and decryption by AES-CTR-NoPadding
- [x] Encryption and decryption by AES-CTR-ZeroPadding
- [x] Encryption and decryption by AES-CTR-PKCS5Padding
- [x] Encryption and decryption by AES-CTR-PKCS7Padding
- [x] Encryption and decryption by AES-CFB-NoPadding
- [x] Encryption and decryption by AES-CFB-ZeroPadding
- [x] Encryption and decryption by AES-CFB-PKCS5Padding
- [x] Encryption and decryption by AES-CFB-PKCS7Padding
- [x] Encryption and decryption by AES-OFB-NoPadding
- [x] Encryption and decryption by AES-OFB-ZeroPadding
- [x] Encryption and decryption by AES-OFB-PKCS5Padding
- [x] Encryption and decryption by AES-OFB-PKCS7Padding
- [x] Encryption and decryption by DES-CBC-NoPadding
- [x] Encryption and decryption by DES-CBC-ZeroPadding
- [x] Encryption and decryption by DES-CBC-PKCS5Padding
- [x] Encryption and decryption by DES-CBC-PKCS7Padding
- [x] Encryption and decryption by DES-CTR-NoPadding
- [x] Encryption and decryption by DES-CTR-ZeroPadding
- [x] Encryption and decryption by DES-CTR-PKCS5Padding
- [x] Encryption and decryption by DES-CTR-PKCS7Padding
- [x] Encryption and decryption by DES-CFB-NoPadding
- [x] Encryption and decryption by DES-CFB-ZeroPadding
- [x] Encryption and decryption by DES-CFB-PKCS5Padding
- [x] Encryption and decryption by DES-CFB-PKCS7Padding
- [x] Encryption and decryption by DES-OFB-NoPadding
- [x] Encryption and decryption by DES-OFB-ZeroPadding
- [x] Encryption and decryption by DES-OFB-PKCS5Padding
- [x] Encryption and decryption by DES-OFB-PKCS7Padding
- [x] Encryption and decryption by 3DES-CBC-NoPadding
- [x] Encryption and decryption by 3DES-CBC-ZeroPadding
- [x] Encryption and decryption by 3DES-CBC-PKCS5Padding
- [x] Encryption and decryption by 3DES-CBC-PKCS7Padding
- [x] Encryption and decryption by 3DES-CTR-NoPadding
- [x] Encryption and decryption by 3DES-CTR-ZeroPadding
- [x] Encryption and decryption by 3DES-CTR-PKCS5Padding
- [x] Encryption and decryption by 3DES-CTR-PKCS7Padding
- [x] Encryption and decryption by 3DES-CFB-NoPadding
- [x] Encryption and decryption by 3DES-CFB-ZeroPadding
- [x] Encryption and decryption by 3DES-CFB-PKCS5Padding
- [x] Encryption and decryption by 3DES-CFB-PKCS7Padding
- [x] Encryption and decryption by 3DES-OFB-NoPadding
- [x] Encryption and decryption by 3DES-OFB-ZeroPadding
- [x] Encryption and decryption by 3DES-OFB-PKCS5Padding
- [x] Encryption and decryption by 3DES-OFB-PKCS7Padding

### References
* [brix/crypto-js](https://github.com/brix/crypto-js)
* [nodejs/crypto](https://nodejs.org/api/crypto.html)
* [paramiko/paramiko](https://github.com/paramiko/paramiko)
* [cossacklabs/themis](https://github.com/cossacklabs/themis)

### Sponsors
`Dongle` is a non-commercial open source project. If you want to support `Dongle`, you can [buy a cup of coffee](https://www.gouguoyin.cn/zanzhu.html) for developer.

### Thanks
`Dongle` had been being developed with GoLand under the free JetBrains Open Source license, I would like to express my
thanks here.

<a href="https://www.jetbrains.com"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" height="100" alt="JetBrains"/></a>