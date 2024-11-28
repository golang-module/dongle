# dongle  #

[![Dongle Release](https://img.shields.io/github/release/dromara/dongle.svg)](https://github.com/dromara/dongle/releases)
[![Go Build](https://github.com/dromara/dongle/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/dongle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/dongle)](https://goreportcard.com/report/github.com/dromara/dongle)
[![codecov](https://codecov.io/gh/dromara/dongle/branch/main/graph/badge.svg)](https://codecov.io/gh/dromara/dongle)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/dongle)
![License](https://img.shields.io/github/license/dromara/dongle)

English | [简体中文](README.cn.md)

### Introduction

A simple, semantic and developer-friendly `golang` crypto package, has been included by [awesome-go](https://github.com/avelino/awesome-go#security "awesome-go")

### Repository

[github.com/dromara/dongle](https://github.com/dromara/dongle "github.com/dromara/dongle")

[gitee.com/dromara/dongle](https://gitee.com/dromara/dongle "gitee.com/dromara/dongle")

[gitcode.com/dromara/dongle](https://gitcode.com/dromara/dongle "gitcode.com/dromara/dongle")

### Installation

```go
// By github
go get -u github.com/dromara/dongle
import "github.com/dromara/dongle"

// By gitee
go get -u gitee.com/dromara/dongle
import "gitee.com/dromara/dongle"

// By gitcode
go get -u gitcode.com/dromara/dongle
import "gitcode.com/dromara/dongle"
```

`Dongle` was donated to the [dromara](https://dromara.org/ "dromara") organization, the repository url has changed, if the previous repository used was `golang-module/dongle`, please replace the original repository with the new repository in `go.mod`, for example

```go
go mod edit -replace github.com/golang-module/dongle=github.com/dromara/dongle
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

##### Encode and decode by base16

```go
// Encode by base16 from string and output string
dongle.Encode.FromString("hello world").ByBase16().ToString() // 68656c6c6f20776f726c64
// Decode by base16 from string and output string
dongle.Decode.FromString("68656c6c6f20776f726c64").ByBase16().ToString() // hello world

// Encode by base16 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase16().ToBytes() // []byte("68656c6c6f20776f726c64")
// Decode by base16 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("68656c6c6f20776f726c64")).ByBase16().ToBytes() // []byte("hello world")
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

##### Encode and decode by base45

```go
// Encode by base45 from string and output string
dongle.Encode.FromString("hello world").ByBase45().ToString() // +8D VD82EK4F.KEA2
// Decode by base45 from string and output string
dongle.Decode.FromString("+8D VD82EK4F.KEA2").ByBase45().ToString() // hello world

// Encode by base45 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase45().ToBytes() // []byte("+8D VD82EK4F.KEA2")
// Decode by base45 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("+8D VD82EK4F.KEA2")).ByBase45().ToBytes() // []byte("hello world")
```

##### Encode and decode by base58

```go
// Encode by base58 from string and output string
dongle.Encode.FromString("hello world").ByBase58().ToString() // StV1DL6CwTryKyV
// Decode by base58 from string and output string
dongle.Decode.FromString("StV1DL6CwTryKyV").ByBase58().ToString() // hello world

// Encode by base58 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase58().ToBytes() // []byte("StV1DL6CwTryKyV")
// Decode by base58 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("StV1DL6CwTryKyV")).ByBase58().ToBytes() // []byte("hello world")
```

##### Encode and decode by base62

>Customizable alphabet, the alphabet must be 62 bytes, the default alphabet is 0123456789ABCDEFGHIJKNMNOPQLRSTUVWXYZabcdefghijklmnopqrstuvwxyz

```go
// Encode by base62 from string and output string
dongle.Encode.FromString("hello world").ByBase62().ToString() // AAwf93rvy4aWQVw
// Decode by base62 from string and output string
dongle.Decode.FromString("AAwf93rvy4aWQVw").ByBase62().ToString() // hello world

// Encode by base62 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase62().ToBytes() // []byte("AAwf93rvy4aWQVw")
// Decode by base62 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("AAwf93rvy4aWQVw")).ByBase62().ToBytes() // []byte("hello world")

// Encode by base62 with custom alphabet from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase62("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789").ToBytes() // []byte("KK6pJD158Ekgaf6")
// Decode by base62 with custom alphabet from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("KK6pJD158Ekgaf6")).ByBase62("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789").ToBytes() // []byte("hello world")
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

##### Encode and decode by base85

```go
// Encode by base85 from string and output string
dongle.Encode.FromString("hello world").ByBase85().ToString() // BOu!rD]j7BEbo7
// Decode by base85 from string and output string
dongle.Decode.FromString("BOu!rD]j7BEbo7").ByBase85().ToString() // hello world

// Encode by base85 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase85().ToBytes() // []byte("BOu!rD]j7BEbo7")
// Decode by base85 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("BOu!rD]j7BEbo7")).ByBase85().ToBytes() // []byte("hello world")
```

##### Encode and decode by base91

```go
// Encode by base91 from string and output string
dongle.Encode.FromString("hello world").ByBase91().ToString() // TPwJh>Io2Tv!lE
// Decode by base91 from string and output string
dongle.Decode.FromString("TPwJh>Io2Tv!lE").ByBase91().ToString() // hello world

// Encode by base91 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase91().ToBytes() // []byte("TPwJh>Io2Tv!lE")
// Decode by base91 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("TPwJh>Io2Tv!lE")).ByBase91().ToBytes() // []byte("hello world")
```

##### Encode and decode by base100

```go
// Encode by base100 from string and output string
dongle.Encode.FromString("hello world").ByBase100().ToString() // 👟👜👣👣👦🐗👮👦👩👣👛
// Decode by base100 from string and output string
dongle.Decode.FromString("TPwJh>Io2Tv!lE").ByBase100().ToString() // hello world

// Encode by base100 from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("hello world")).ByBase100().ToBytes() // []byte("👟👜👣👣👦🐗👮👦👩👣👛")
// Decode by base100 from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("👟👜👣👣👦🐗👮👦👩👣👛")).ByBase100().ToBytes() // []byte("hello world")
```

##### Encode and decode by safeURL

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

##### Encode and decode by morse

> The default value of separator is `/`

```go
// Encode by morse from string and output string
dongle.Encode.FromString("dongle").ByMorse().ToString() // -../---/-./--./.-../.
// Decode by morse from string and output string
dongle.Decode.FromString("-../---/-./--./.-../.").ByMorse().ToString() // dongle

// Encode by morse from byte slice and output byte slice
dongle.Encode.FromBytes([]byte("dongle")).ByMorse("|").ToBytes() // []byte("-..|---|-.|--.|.-..|.")
// Decode by morse from byte slice and output byte slice
dongle.Decode.FromBytes([]byte("-..|---|-.|--.|.-..|.")).ByMorse("|").ToBytes() // []byte("dongle")
```

#### Encrypt and decrypt

##### Encrypt by md2

```go
// Encrypt by md2 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd2().ToHexString() // d9cce882ee690a5c1ce70beff3a78c77
// Encrypt by md2 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd2().ToBase64String() // 2czogu5pClwc5wvv86eMdw==

// Encrypt by md2 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd2().ToHexBytes() // []byte("d9cce882ee690a5c1ce70beff3a78c77")
// Encrypt by md2 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd2().ToBase64Bytes() // []byte("2czogu5pClwc5wvv86eMdw==")

```

##### Encrypt by md4

```go
// Encrypt by md4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToHexString() // aa010fbc1d14c795d86ef98c95479d17
// Encrypt by md4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd4().ToBase64String() // qgEPvB0Ux5XYbvmMlUedFw==

// Encrypt by md4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToHexBytes() // []byte("aa010fbc1d14c795d86ef98c95479d17")
// Encrypt by md4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd4().ToBase64Bytes() // []byte("qgEPvB0Ux5XYbvmMlUedFw==")
```

##### Encrypt by md5

```go
// Encrypt by md5 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
// Encrypt by md5 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByMd5().ToBase64String() // XrY7u+Ae7tCTyyK7j1rNww==

// Encrypt by md5 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToHexBytes() // []byte("5eb63bbbe01eeed093cb22bb8f5acdc3")
// Encrypt by md5 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd5().ToBase64Bytes() // []byte("XrY7u+Ae7tCTyyK7j1rNww==")
```

##### Encrypt by sha1

```go
// Encrypt by sha1 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha1().ToHexString() // 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
// Encrypt by sha1 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha1().ToBase64String() // Kq5sNclPz7QV2+lfQIuc6R7oRu0=

// Encrypt by sha1 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToHexBytes() // []byte("2aae6c35c94fcfb415dbe95f408b9ce91ee846ed")
// Encrypt by sha1 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha1().ToBase64Bytes() // []byte("Kq5sNclPz7QV2+lfQIuc6R7oRu0=")
```

##### Encrypt by sha3

> include sha3-224, sha3-256, sha3-384, sha3-512

```go
// Encrypt by sha3-224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").Sha3(224).ToHexString() // dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5
// Encrypt by sha3-224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").Sha3(224).ToBase64String() // 37fxjHfpKLtW+ustonKRvXkLwQRc3kXzIQu2xQ==
// Encrypt by sha3-224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(224).ToHexBytes() // []byte("dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5")
// Encrypt by sha3-224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(224).ToBase64Bytes() // []byte("37fxjHfpKLtW+ustonKRvXkLwQRc3kXzIQu2xQ==")

// Encrypt by sha3-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").Sha3(256).ToHexString() // 644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938
// Encrypt by sha3-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").Sha3(256).ToBase64String() // ZEvMflZDcwQJmarInnYi88px+6HZcv2Uoxw7+/JOOTg=
// Encrypt by sha3-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(256).ToHexBytes() // []byte("644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938")
// Encrypt by sha3-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(256).ToBase64Bytes() // []byte("ZEvMflZDcwQJmarInnYi88px+6HZcv2Uoxw7+/JOOTg=")

// Encrypt by sha3-384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").Sha3(384).ToHexString() // 83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b
// Encrypt by sha3-384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").Sha3(384).ToBase64String() // g7/yjd4bG/WBAHHGZDwI5bBb24Nu/9cLQD6o6gpjTcSZfrEFOqNZP1kPnGNjDdkL
// Encrypt by sha3-384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(384).ToHexBytes() // []byte("83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b")
// Encrypt by sha3-384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(384).ToBase64Bytes() // []byte("g7/yjd4bG/WBAHHGZDwI5bBb24Nu/9cLQD6o6gpjTcSZfrEFOqNZP1kPnGNjDdkL")

// Encrypt by sha3-512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").Sha3(512).ToHexString() // 840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a
// Encrypt by sha3-512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").Sha3(512).ToBase64String() // hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg==
// Encrypt by sha3-512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(512).ToHexBytes() // []byte("840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a")
// Encrypt by sha3-512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(512).ToBase64Bytes() // []byte("hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg==")

```

##### Encrypt by sha224

```go
// Encrypt by sha224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha224().ToHexString() // 2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b
// Encrypt by sha224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha224().ToBase64String() // LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==

// Encrypt by sha224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToHexBytes() // []byte("2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b")
// Encrypt by sha224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha224().ToBase64Bytes() // []byte("LwVHf8JLtPrv2GUXFW2v3s7EW4rTzyUipWNYKw==")
```

##### Encrypt by sha256

```go
// Encrypt by sha256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha256().ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by sha256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha256().ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by sha256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by sha256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha256().ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by sha384

```go
// Encrypt by sha384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha384().ToHexString() // fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
// Encrypt by sha384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha384().ToBase64String() // /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9

// Encrypt by sha384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToHexBytes() // []byte("fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd")
// Encrypt by sha384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha384().ToBase64Bytes() // []byte("/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9")
```

##### Encrypt by sha512

> include sha512, sha512-224, sha512-256

```go
// Encrypt by sha512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha512().ToHexBytes() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// Encrypt by sha512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha512().ToBase64String() // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==
// Encrypt by sha512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToHexBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// Encrypt by sha512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase64Bytes() // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")

// Encrypt by sha512-224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha512(224).ToHexBytes() // 22e0d52336f64a998085078b05a6e37b26f8120f43bf4db4c43a64ee
// Encrypt by sha512-224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha512(224).ToBase64String() // IuDVIzb2SpmAhQeLBabjeyb4Eg9Dv020xDpk7g==
// Encrypt by sha512-224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(224).ToHexBytes() // []byte("22e0d52336f64a998085078b05a6e37b26f8120f43bf4db4c43a64ee")
// Encrypt by sha512-224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(224).ToBase64Bytes() // []byte("IuDVIzb2SpmAhQeLBabjeyb4Eg9Dv020xDpk7g==")

// Encrypt by sha512-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySha512(256).ToHexBytes() // 0ac561fac838104e3f2e4ad107b4bee3e938bf15f2b15f009ccccd61a913f017
// Encrypt by sha512-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySha512(256).ToBase64String() // CsVh+sg4EE4/LkrRB7S+4+k4vxXysV8AnMzNYakT8Bc=
// Encrypt by sha512-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(256).ToHexBytes() // []byte("0ac561fac838104e3f2e4ad107b4bee3e938bf15f2b15f009ccccd61a913f017")
// Encrypt by sha512-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(256).ToBase64Bytes() // []byte("CsVh+sg4EE4/LkrRB7S+4+k4vxXysV8AnMzNYakT8Bc=")
```

##### Encrypt by shake128

```go
// Encrypt by shake128-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByShake128(256).ToHexString() // 3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b8
// Encrypt by shake128-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByShake128(256).ToBase64String() // OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLg=
// Encrypt by shake128-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(256).ToHexBytes() // []byte("3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b8")
// Encrypt by shake128-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(256).ToBase64Bytes() // []byte("OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLg=")

// Encrypt by shake128-512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByShake128(512).ToHexString() // 3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b899072665674f26cc494a4bcf027c58267e8ee2da60e942759de86d2670bba1aa
// Encrypt by shake128-512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByShake128(512).ToBase64String() // OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLiZByZlZ08mzElKS88CfFgmfo7i2mDpQnWd6G0mcLuhqg==
// Encrypt by shake128-512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(512).ToHexBytes() // []byte("3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b899072665674f26cc494a4bcf027c58267e8ee2da60e942759de86d2670bba1aa")
// Encrypt by shake128-512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(512).ToBase64Bytes() // []byte("OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLiZByZlZ08mzElKS88CfFgmfo7i2mDpQnWd6G0mcLuhqg==")
```

##### Encrypt by shake256

```go
// Encrypt by shake256-384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByShake256(384).ToHexString() // 369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df2
// Encrypt by shake256-384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByShake256(384).ToBase64String() // Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3y
// Encrypt by shake256-384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(384).ToHexBytes() // []byte("369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df2")
// Encrypt by shake256-384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(384).ToBase64Bytes() // []byte("Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3y")

// Encrypt by shake256-512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByShake256(512).ToHexString() // 369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df27853a8735271f5cc2d9e889544357116
// Encrypt by shake256-512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByShake256(512).ToBase64String() // Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3yeFOoc1Jx9cwtnoiVRDVxFg==
// Encrypt by shake256-512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(512).ToHexBytes() // []byte("369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df27853a8735271f5cc2d9e889544357116")
// Encrypt by shake256-512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(512).ToBase64Bytes() // []byte("Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3yeFOoc1Jx9cwtnoiVRDVxFg==")
```

##### Encrypt by ripemd160

```go
// Encrypt by ripemd160 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByRipemd160().ToHexString() // 98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f
// Encrypt by ripemd160 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByRipemd160().ToBase64String() // mMYVeEzLX+WTb7wMvp39tAjZLw8=

// Encrypt by ripemd160 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRipemd160().ToHexBytes() // []byte("98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f")
// Encrypt by ripemd160 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRipemd160().ToBase64Bytes() // []byte("mMYVeEzLX+WTb7wMvp39tAjZLw8=")
```

##### Encrypt by blake2b
> include blake2b-256, blake2b-384, blake2b-512

```go
// Encrypt by blake2b-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByBlake2b(256).ToHexBytes() // 256c83b297114d201b30179f3f0ef0cace9783622da5974326b436178aeef610
// Encrypt by blake2b-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByBlake2b(256).ToBase64String() // JWyDspcRTSAbMBefPw7wys6Xg2ItpZdDJrQ2F4ru9hA=
// Encrypt by blake2b-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(256).ToHexBytes() // []byte("256c83b297114d201b30179f3f0ef0cace9783622da5974326b436178aeef610")
// Encrypt by blake2b-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(256).ToBase64Bytes() // []byte("JWyDspcRTSAbMBefPw7wys6Xg2ItpZdDJrQ2F4ru9hA=")

// Encrypt by blake2b-384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByBlake2b(384).ToHexBytes() // 8c653f8c9c9aa2177fb6f8cf5bb914828faa032d7b486c8150663d3f6524b086784f8e62693171ac51fc80b7d2cbb12b
// Encrypt by blake2b-384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByBlake2b(384).ToBase64String() // jGU/jJyaohd/tvjPW7kUgo+qAy17SGyBUGY9P2UksIZ4T45iaTFxrFH8gLfSy7Er
// Encrypt by blake2b-384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(384).ToHexBytes() // []byte("8c653f8c9c9aa2177fb6f8cf5bb914828faa032d7b486c8150663d3f6524b086784f8e62693171ac51fc80b7d2cbb12b")
// Encrypt by blake2b-384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(384).ToBase64Bytes() // []byte("jGU/jJyaohd/tvjPW7kUgo+qAy17SGyBUGY9P2UksIZ4T45iaTFxrFH8gLfSy7Er")

// Encrypt by blake2b-512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByBlake2b(512).ToHexBytes() // 021ced8799296ceca557832ab941a50b4a11f83478cf141f51f933f653ab9fbcc05a037cddbed06e309bf334942c4e58cdf1a46e237911ccd7fcf9787cbc7fd0
// Encrypt by blake2b-512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByBlake2b(512).ToBase64String() // Ahzth5kpbOylV4MquUGlC0oR+DR4zxQfUfkz9lOrn7zAWgN83b7QbjCb8zSULE5YzfGkbiN5EczX/Pl4fLx/0A==
// Encrypt by blake2b-512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(512).ToHexBytes() // []byte("021ced8799296ceca557832ab941a50b4a11f83478cf141f51f933f653ab9fbcc05a037cddbed06e309bf334942c4e58cdf1a46e237911ccd7fcf9787cbc7fd0")
// Encrypt by blake2b-512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(512).ToBase64Bytes() // []byte("Ahzth5kpbOylV4MquUGlC0oR+DR4zxQfUfkz9lOrn7zAWgN83b7QbjCb8zSULE5YzfGkbiN5EczX/Pl4fLx/0A==")
```

##### Encrypt by blake2s
> Currently, only blake2s-256 is supported

```go
// Encrypt by blake2s-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByBlake2s(256).ToHexString() // 9aec6806794561107e594b1f6a8a6b0c92a0cba9acf5e5e93cca06f781813b0b
// Encrypt by blake2s-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByBlake2s(256).ToBase64String() // muxoBnlFYRB+WUsfaoprDJKgy6ms9eXpPMoG94GBOws=

// Encrypt by blake2s-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2s(256).ToHexBytes() // []byte("9aec6806794561107e594b1f6a8a6b0c92a0cba9acf5e5e93cca06f781813b0b")
// Encrypt by blake2s-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2s(256).ToBase64Bytes() // []byte("muxoBnlFYRB+WUsfaoprDJKgy6ms9eXpPMoG94GBOws=")
```

##### Encrypt by hmac-md2

```go
// Encrypt by hmac-md2 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd2([]byte("dongle")).ToHexString() // 88ed6ef9ab699d03a702f2a6fb1c0673
// Encrypt by hmac-md2 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd2([]byte("dongle")).ToBase64String() // iO1u+atpnQOnAvKm+xwGcw==

// Encrypt by hmac-md2 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd2([]byte("dongle")).ToHexBytes() // []byte("88ed6ef9ab699d03a702f2a6fb1c0673")
// Encrypt by hmac-md2 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd2([]byte("dongle")).ToBase64Bytes() // []byte("iO1u+atpnQOnAvKm+xwGcw==")
```

##### Encrypt by hmac-md4

```go
// Encrypt by hmac-md4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4([]byte("dongle")).ToHexString() // 7a9df5247cbf76a8bc17c9c4f5a75b6b
// Encrypt by hmac-md4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd4([]byte("dongle")).ToBase64String() // ep31JHy/dqi8F8nE9adbaw==

// Encrypt by hmac-md4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToHexBytes() // []byte("7a9df5247cbf76a8bc17c9c4f5a75b6b")
// Encrypt by hmac-md4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase64Bytes() // []byte("ep31JHy/dqi8F8nE9adbaw==")
```

##### Encrypt by hmac-md5

```go
// Encrypt by hmac-md5 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5([]byte("dongle")).ToHexString() // 4790626a275f776956386e5a3ea7b726
// Encrypt by hmac-md5 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacMd5([]byte("dongle")).ToBase64String() // R5Biaidfd2lWOG5aPqe3Jg==

// Encrypt by hmac-md5 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToHexBytes() // []byte("4790626a275f776956386e5a3ea7b726")
// Encrypt by hmac-md5 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase64Bytes() // []byte("R5Biaidfd2lWOG5aPqe3Jg==")
```

##### Encrypt by hmac-sha1

```go
// Encrypt by hmac-sha1 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1([]byte("dongle")).ToHexString() // 91c103ef93ba7420902b0d1bf0903251c94b4a62
// Encrypt by hmac-sha1 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha1([]byte("dongle")).ToBase64String() // kcED75O6dCCQKw0b8JAyUclLSmI=

// Encrypt by hmac-sha1 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToHexBytes() // []byte("91c103ef93ba7420902b0d1bf0903251c94b4a62")
// Encrypt by hmac-sha1 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase64Bytes() // []byte("kcED75O6dCCQKw0b8JAyUclLSmI=")
```

##### Encrypt by hmac-sha3

> include hmac-sha3-224, hmac-sha3-256, hmac-sha3-384, hmac-sha3-512

```go
// Encrypt by hmac-sha3-224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 224).ToHexString() // fb8f061d9d1dddd2f5d3b9064a5e98e3e4b6df27ea93ce67627583ce
// Encrypt by hmac-sha3-224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 224).ToBase64String() // +48GHZ0d3dL107kGSl6Y4+S23yfqk85nYnWDzg==
// Encrypt by hmac-sha3-224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 224).ToHexBytes() // []byte("fb8f061d9d1dddd2f5d3b9064a5e98e3e4b6df27ea93ce67627583ce")
// Encrypt by hmac-sha3-224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 224).ToBase64Bytes() // []byte("+48GHZ0d3dL107kGSl6Y4+S23yfqk85nYnWDzg==")

// Encrypt by hmac-sha3-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 256).ToHexString() // 8193367fde28cf5c460adb449a04b3dd9c184f488bdccbabf0526c54f90c4460
// Encrypt by hmac-sha3-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 256).ToBase64String() // gZM2f94oz1xGCttEmgSz3ZwYT0iL3Mur8FJsVPkMRGA=
// Encrypt by hmac-sha3-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 256).ToHexBytes() // []byte("8193367fde28cf5c460adb449a04b3dd9c184f488bdccbabf0526c54f90c4460")
// Encrypt by hmac-sha3-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 256).ToBase64Bytes() // []byte("gZM2f94oz1xGCttEmgSz3ZwYT0iL3Mur8FJsVPkMRGA=")

// Encrypt by hmac-sha3-384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 384).ToHexString() // 3f76f5cda69cada3ee6b33f8458cd498b063075db263dd8b33f2a3992a8804f9569a7c86ffa2b8f0748babeb7a6fc0e7
// Encrypt by hmac-sha3-384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 384).ToBase64String() // P3b1zaacraPuazP4RYzUmLBjB12yY92LM/KjmSqIBPlWmnyG/6K48HSLq+t6b8Dn
// Encrypt by hmac-sha3-384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 384).ToHexBytes() // []byte("3f76f5cda69cada3ee6b33f8458cd498b063075db263dd8b33f2a3992a8804f9569a7c86ffa2b8f0748babeb7a6fc0e7")
// Encrypt by hmac-sha3-384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 384).ToBase64Bytes() // []byte("P3b1zaacraPuazP4RYzUmLBjB12yY92LM/KjmSqIBPlWmnyG/6K48HSLq+t6b8Dn")

// Encrypt by hmac-sha3-512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 512).ToHexString() // a99653d0407d659eccdeed43bb7cccd2e2b05a2c34fd3467c4198cf2ad26a466738513e88839fb55e64eb49df65bc52ed0fec2775bd9e086edd4fb4024add4a2
// Encrypt by hmac-sha3-512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 512).ToBase64String() // qZZT0EB9ZZ7M3u1Du3zM0uKwWiw0/TRnxBmM8q0mpGZzhRPoiDn7VeZOtJ32W8Uu0P7Cd1vZ4Ibt1PtAJK3Uog==
// Encrypt by hmac-sha3-512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 512).ToHexBytes() // []byte("a99653d0407d659eccdeed43bb7cccd2e2b05a2c34fd3467c4198cf2ad26a466738513e88839fb55e64eb49df65bc52ed0fec2775bd9e086edd4fb4024add4a2")
// Encrypt by hmac-sha3-512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 512).ToBase64Bytes() // []byte("qZZT0EB9ZZ7M3u1Du3zM0uKwWiw0/TRnxBmM8q0mpGZzhRPoiDn7VeZOtJ32W8Uu0P7Cd1vZ4Ibt1PtAJK3Uog==")
```

##### Encrypt by hmac-sha224

```go
// Encrypt by hmac-sha224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224([]byte("dongle")).ToHexString() // e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec
// Encrypt by hmac-sha224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha224([]byte("dongle")).ToBase64String() // 4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A====

// Encrypt by hmac-sha224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToHexBytes() // []byte("e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec")
// Encrypt by hmac-sha224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase64Bytes() // []byte("4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A==")
```

##### Encrypt by hmac-sha256

```go
// Encrypt by hmac-sha256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256([]byte("dongle")).ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// Encrypt by hmac-sha256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha256([]byte("dongle")).ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// Encrypt by hmac-sha256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// Encrypt by hmac-sha256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Encrypt by hmac-sha384

```go
// Encrypt by hmac-sha384 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384([]byte("dongle")).ToHexString() // 421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8
// Encrypt by hmac-sha384 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha384([]byte("dongle")).ToBase64String() // Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o

// Encrypt by hmac-sha384 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToHexBytes() // []byte("421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8")
// Encrypt by hmac-sha384 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase64Bytes() // []byte("Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o")
```

##### Encrypt by hmac-sha512

> include hmac-sha512, hmac-sha512-224, hmac-sha512-256

```go
// Encrypt by hmac-sha512 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle")).ToHexString() // d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1
// Encrypt by hmac-sha512 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle")).ToBase64String() // 2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==
// Encrypt by hmac-sha512 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToHexBytes() // []byte("d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1")
// Encrypt by hmac-sha512 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase64Bytes() // []byte("2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==")

// Encrypt by hmac-sha512-224 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 224).ToHexString() // c4145bcc385c29f0e5683cd5450be9deb522d556de3b046a7ffa1eb3
// Encrypt by hmac-sha512-224 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 224).ToBase64String() // xBRbzDhcKfDlaDzVRQvp3rUi1VbeOwRqf/oesw==
// Encrypt by hmac-sha512-224 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 224).ToHexBytes() // []byte("c4145bcc385c29f0e5683cd5450be9deb522d556de3b046a7ffa1eb3")
// Encrypt by hmac-sha512-224 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 224).ToBase64Bytes() // []byte("xBRbzDhcKfDlaDzVRQvp3rUi1VbeOwRqf/oesw==")

// Encrypt by hmac-sha512-256 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 256).ToHexString() // e99fae71bcb43651ae10e952989eadf897faccb43966ee5122bb1b1d82f7a7c2
// Encrypt by hmac-sha512-256 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 256).ToBase64String() // 6Z+ucby0NlGuEOlSmJ6t+Jf6zLQ5Zu5RIrsbHYL3p8I=
// Encrypt by hmac-sha512-256 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 256).ToHexBytes() // []byte("e99fae71bcb43651ae10e952989eadf897faccb43966ee5122bb1b1d82f7a7c2")
// Encrypt by hmac-sha512-256 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 256).ToBase64Bytes() // []byte("6Z+ucby0NlGuEOlSmJ6t+Jf6zLQ5Zu5RIrsbHYL3p8I=")
```

##### Encrypt by hmac-ripemd160

```go
// Encrypt by hmac-ripemd160 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacRipemd160([]byte("dongle")).ToHexString() // 3691ad040e80c43dc6e8ffe9bc6ef3d5bd8786b8
// Encrypt by hmac-ripemd160 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacRipemd160([]byte("dongle")).ToBase64String() // NpGtBA6AxD3G6P/pvG7z1b2Hhrg=

// Encrypt by hmac-ripemd160 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacRipemd160([]byte("dongle")).ToHexBytes() // []byte("3691ad040e80c43dc6e8ffe9bc6ef3d5bd8786b8")
// Encrypt by hmac-ripemd160 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacRipemd160([]byte("dongle")).ToBase64Bytes() // []byte("NpGtBA6AxD3G6P/pvG7z1b2Hhrg=")
```

##### Encrypt by hmac-sm3

```go
// Encrypt by hmac-sm3 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByHmacSm3().ToHexString() // 8c733aae1d553c466a08c3e9e5daac3e99ae220181c7c1bc8c2564961de751b3
// Encrypt by hmac-sm3 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByHmacSm3().ToBase64String() // jHM6rh1VPEZqCMPp5dqsPpmuIgGBx8G8jCVklh3nUbM=

// Encrypt by hmac-sm3 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSm3().ToHexBytes() // []byte("8c733aae1d553c466a08c3e9e5daac3e99ae220181c7c1bc8c2564961de751b3")
// Encrypt by hmac-sm3 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSm3().ToBase64Bytes() // []byte("jHM6rh1VPEZqCMPp5dqsPpmuIgGBx8G8jCVklh3nUbM=")
```

##### Encrypt and decrypt by aes

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、CFB、OFB、CTR、ECB
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("0123456789abcdef") // key must be 16, 24 or 32 bytes
// cipher.WithKey([]byte("0123456789abcdef")) 
cipher.SetIV("0123456789abcdef") // iv must be 16 bytes (ECB mode doesn't require setting iv)
// cipher.WithIV([]byte("0123456789abcdef"))

// Encrypt by aes from string and output raw string
rawString := dongle.Encrypt.FromString("hello world").ByAes(cipher).ToRawString()
// Decrypt by aes from raw string and output string
dongle.Decrypt.FromRawString(rawString).ByAes(cipher).ToString() // hello world

// Encrypt by aes from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToHexString() // c1e9b4529aac9793010f4677f6358efe
// Decrypt by aes from string with hex encoding and output string
dongle.Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByAes(cipher).ToString() // hello world

// Encrypt by aes from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase64String() // wem0Upqsl5MBD0Z39jWO/g==
// Decrypt by aes from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("wem0Upqsl5MBD0Z39jWO/g==").ByAes(cipher).ToString() // hello world

// Encrypt by aes from byte slice and output raw byte slice
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToRawBytes()
// Decrypt by aes from raw byte slice and output byte slice
dongle.Decrypt.FromRawBytes(rawBytes)).ByAes(cipher).ToBytes() // []byte("hello world")

// Encrypt by aes from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToHexBytes() // []byte("c1e9b4529aac9793010f4677f6358efe")
// Decrypt by aes from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("c1e9b4529aac9793010f4677f6358efe")).ByAes(cipher).ToBytes() // []byte("hello world")

// Encrypt by aes from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase64Bytes() // []byte("wem0Upqsl5MBD0Z39jWO/g==")
// Decrypt by aes from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("wem0Upqsl5MBD0Z39jWO/g==")).ByAes(cipher).ToBytes() // []byte("hello world")

```

##### Encrypt and decrypt by blowfish

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、CFB、OFB、CTR、ECB
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("0123456789abcdef")  // key must from 1 to 56 bytes
// cipher.WithKey([]byte("0123456789abcdef"))
cipher.SetIV("12345678"))           // iv must be 8 bytes
// cipher.WithIV([]byte("12345678"))          

// Encrypt by blowfish from string and output raw string
rawString := dongle.Encrypt.FromString("hello world").ByBlowfish(cipher).ToRawString()
// Decrypt by blowfish from raw string and output string
dongle.Decrypt.FromRawString(rawString).ByBlowfish(cipher).ToString() // hello world

// Encrypt by blowfish from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByBlowfish(cipher).ToHexString() // c1e9b4529aac9793010f4677f6358efe
// Decrypt by blowfish from string with hex encoding and output string
dongle.Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByBlowfish(cipher).ToString() // hello world

// Encrypt by blowfish from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByBlowfish(cipher).ToBase64String() // ksNyTXILWZgtIaq5p7ufQA==
// Decrypt by blowfish from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("ksNyTXILWZgtIaq5p7ufQA==").ByBlowfish(cipher).ToString() // hello world

// Encrypt by blowfish from byte slice and output raw byte slice
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).ByBlowfish(cipher).ToRawBytes()
// Decrypt by blowfish from raw byte slice and output byte slice
dongle.Decrypt.FromRawBytes(rawBytes).ByBlowfish(cipher).ToBytes() // []byte("hello world")

// Encrypt by blowfish from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlowfish(cipher).ToHexBytes() // []byte("c1e9b4529aac9793010f4677f6358efe")
// Decrypt by blowfish from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("c1e9b4529aac9793010f4677f6358efe")).ByBlowfish(cipher).ToBytes() // []byte("hello world")

// Encrypt by blowfish from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlowfish(cipher).ToBase64Bytes() // []byte("ksNyTXILWZgtIaq5p7ufQA==")
// Decrypt by blowfish from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("ksNyTXILWZgtIaq5p7ufQA==")).ByBlowfish(cipher).ToBytes() // []byte("hello world")
```

##### Encrypt and decrypt by des

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("12345678") // key must be 8 bytes
// cipher.WithKey([]byte("12345678")) 
cipher.SetIV("12345678") // iv must be 8 bytes
// cipher.WithIV([]byte("12345678")) 

// Encrypt by des from string and output raw string
rawString := dongle.Encrypt.FromString("hello world").ByDes(cipher).ToRawString()
// Decrypt by des from raw string and output string
dongle.Decrypt.FromRawString(rawString).ByDes(cipher).ToString() // hello world

// Encrypt by des from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// Decrypt by des from string with hex encoding and output string
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(cipher).ToString() // hello world

// Encrypt by des from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// Decrypt by des from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").ByDes(cipher).ToString() // hello world

// Encrypt by des from byte slice and output raw byte slice
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToRawBytes()
// Decrypt by des from raw byte slice and output byte slice
dongle.Decrypt.FromRawBytes(rawBytes).ByDes(cipher).ToBytes() // []byte("hello world")

// Encrypt by des from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// Decrypt by des from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).ByDes(cipher).ToBytes() // []byte("hello world")

// Encrypt by des from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// Decrypt by des from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).ByDes(cipher).ToBytes() // []byte("hello world")
```

##### Encrypt and decrypt by 3des

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("123456781234567812345678") // key must be 24 bytes
// cipher.WithKey([]byte("123456781234567812345678"))
cipher.SetIV("12345678") // v must be 8 bytes
// cipher.WithIV([]byte("12345678")) 

// Encrypt by 3des from string and output raw string
rawString := dongle.Encrypt.FromString("hello world").By3Des(cipher).ToRawString()
// Decrypt by 3des from raw string and output string
dongle.Decrypt.FromRawString(rawString).By3Des(cipher).ToString() // hello world

// Encrypt by 3des from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// Decrypt by des from string with hex encoding and output string
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// Decrypt by 3des from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").By3Des(cipher).ToString() // hello world

// Encrypt by 3des from byte slice and output raw byte slice
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToRawBytes()
// Decrypt by 3des from raw byte slice and output byte slice
dongle.Decrypt.FromRawBytes(rawBytes).By3Des(cipher).ToBytes() // []byte("hello world")

// Encrypt by 3des from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToHexBytes() // []byte("0b2a92e81fb49ce1a43266aacaea7b81")
// Decrypt by 3des from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("0b2a92e81fb49ce1a43266aacaea7b81")).By3Des(cipher).ToBytes() // []byte("hello world")

// Encrypt by 3des from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// Decrypt by 3des from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).By3Des(cipher).ToBytes() // []byte("hello world")
```

##### Encrypt and decrypt by rsa

```go
pkcs1PublicKey := []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`)

pkcs1PrivateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

pkcs8PublicKey := []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`)

pkcs8PrivateKey := []byte(`-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`)
```

###### 1.RSA encrypt by public key and decrypt by private key

```go
// Use pkcs1 format public key to encrypt by rsa from string
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs1PublicKey)
// Use pkcs1 format private key to decrypt by rsa from string with hex encoding
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs1PrivateKey).ToString() // hello world
// Use pkcs1 format private key to decrypt by rsa from string with base64 encoding
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs1PrivateKey).ToString() // hello world

// Use pkcs1 format public key to encrypt by rsa from byte
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs1PublicKey)
// Use pkcs1 format private key to decrypt by rsa from byte with hex encoding
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs1PrivateKey).ToByte() // []bytes("hello world)
// Use pkcs1 format private key to decrypt by rsa from byte with base64 encoding
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs1PrivateKey).ToByte() // []bytes("hello world)

// Use pkcs8 format public key to encrypt by rsa from string
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs8PublicKey)
// Use pkcs8 format private key to decrypt by rsa from string with hex encoding
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs8PrivateKey).ToString() // hello world
// Use pkcs8 format private key to decrypt by rsa from string with base64 encoding
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs8PrivateKey).ToString() // hello world

// Use pkcs8 format public key to encrypt by rsa from byte
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs8PublicKey)
// Use pkcs8 format private key to decrypt by rsa from byte with hex encoding
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs8PrivateKey).ToByte() // []bytes("hello world)
// Use pkcs8 format private key to decrypt by rsa from byte with base64 encoding
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs8PrivateKey).ToByte() // []bytes("hello world)
```

###### 2.RSA encrypt by private key and decrypt by public key

```go
// Use pkcs1 format private key to encrypt by rsa from string
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs1PrivateKey)
// Use pkcs1 format public key to decrypt by rsa from string with hex encoding
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs1PublicKey).ToString() // hello world
// Use pkcs1 format public key to decrypt by rsa from string with base64 encoding
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs1PublicKey).ToString() // hello world

// Use pkcs1 format private key to encrypt by rsa from byte
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs1PrivateKey)
// Use pkcs1 format public key to decrypt by rsa from byte with hex encoding
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa([]byte(pkcs1PublicKey)).ToByte() // []bytes("hello world)
// Use pkcs1 format public key to decrypt by rsa from byte with base64 encoding
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs1PublicKey).ToByte() // []bytes("hello world)

// Use pkcs8 format private key to encrypt by rsa from string
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs8PrivateKey)
// Use pkcs8 format  public key to decrypt by rsa from string with hex encoding
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs8PublicKey).ToString() // hello world
// Use pkcs8 format  public key to decrypt by rsa from string with base64 encoding
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs8PublicKey).ToString() // hello world

// Use pkcs8 format private key to encrypt by rsa from byte
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs8PrivateKey)
// Use pkcs8 format public key to decrypt by rsa from byte with hex encoding
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs8PublicKey).ToByte() // []bytes("hello world)
// Use pkcs8 format public key to decrypt by rsa from byte with base64 encoding
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs8PublicKey).ToByte() // []bytes("hello world)
```

##### Encrypt and decrypt by tea

> If plaintext more than 8 bytes, block encryption with empty string padding. Rounds must be even, the default value is 64

```go
// Encrypt by tea from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByTea([]byte("0123456789abcdef")).ToHexString() // bfa8d956247c0dcecfe2b0ef44b08aab
// Decrypt by tea from string with hex encoding and output string
dongle.Decrypt.FromHexString("bfa8d956247c0dcecfe2b0ef44b08aab").ByTea([]byte("0123456789abcdef")).ToString() // hello world

// Encrypt by tea from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByTea([]byte("0123456789abcdef")).ToBase64String() // v6jZViR8Dc7P4rDvRLCKqw==
// Decrypt by tea from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("v6jZViR8Dc7P4rDvRLCKqw==").ByTea([]byte("0123456789abcdef")).ToString() // hello world

// Encrypt by tea from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByTea([]byte("0123456789abcdef"), 32).ToHexBytes() // []byte("bfa8d956247c0dcecfe2b0ef44b08aab")
// Decrypt by tea from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromHexBytes([]byte("bfa8d956247c0dcecfe2b0ef44b08aab")).ByTea("4539658173a59fc2", 64).ToBytes() // []byte("hello world")

// Encrypt by tea from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByTea([]byte("0123456789abcdef"), 32).ToBase64Bytes() // []byte("v6jZViR8Dc7P4rDvRLCKqw==")
// Decrypt by tea from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes(()byte("v6jZViR8Dc7P4rDvRLCKqw==")).ByTea([]byte("0123456789abcdef"), 64).ToBytes() // []byte("hello world")
```

##### Encrypt and decrypt by rc4

```go
// Encrypt by rc4 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").ByRc4([]byte("dongle")).ToHexString() // eba154b4cb5a9038dbbf9d
// Decrypt by rc4 from string with hex encoding and output string
dongle.Encrypt.FromHexString("eba154b4cb5a9038dbbf9d").ByRc4([]byte("dongle")).ToString() // hello world

// Encrypt by rc4 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").ByRc4([]byte("dongle")).ToBase64String() // 66FUtMtakDjbv50=
// Decrypt by rc4 from string with base64 encoding and output string
dongle.Decrypt.FromBase64String("66FUtMtakDjbv50=").ByRc4([]byte("dongle")).ToString() // hello world

// Encrypt by rc4 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4([]byte("dongle")).ToHexBytes() // []byte("eba154b4cb5a9038dbbf9d")
// Decrypt by rc4 from byte slice with hex encoding and output byte slice
dongle.Decrypt.FromBase64Bytes([]byte("hello world")).ByRc4([]byte("dongle")).ToBytes() // []byte("eba154b4cb5a9038dbbf9d")

// Encrypt by rc4 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4([]byte("dongle")).ToBase64Bytes() // []byte("66FUtMtakDjbv50=")
// Decrypt by rc4 from byte slice with base64 encoding and output byte slice
dongle.Decrypt.FromBase64Bytes([]byte("66FUtMtakDjbv50=")).ByRc4([]byte("dongle")).ToBytes() // []byte("hello world")
```

##### Encrypt by sm3

```go
// Encrypt by sm3 from string and output string with hex encoding
dongle.Encrypt.FromString("hello world").BySm3().ToHexString() // 44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88
// Encrypt by sm3 from string and output string with base64 encoding
dongle.Encrypt.FromString("hello world").BySm3().ToBase64String() // RPAGHmn6b9/CkMSUZUoF3AwFPaflxSuE75Op1n0//4g=

// Encrypt by sm3 from byte slice and output byte slice with hex encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySm3().ToHexBytes() // []byte("44f0061e69fa6fdfc290c494654a05dc0c053da7e5c52b84ef93a9d67d3fff88")
// Encrypt by sm3 from byte slice and output byte slice with base64 encoding
dongle.Encrypt.FromBytes([]byte("hello world")).BySm3().ToBase64Bytes() // []byte("RPAGHmn6b9/CkMSUZUoF3AwFPaflxSuE75Op1n0//4g=")
```

#### Sign and verify

##### Bcrypt sign and verify

> Rounds allowed range is 4 to 31, the default value is 10

```go
// Sign by bcrypt from string and output string
sign := dongle.Sign.FromString("hello world").ByBcrypt()
// Verify by bcrypt from signature string without encoding and message string, output bool
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByBcrypt().ToBool() // true

// Sign by bcrypt from byte slice and output byte slice
sign := dongle.Sign.FromBytes([]byte("hello world")).ByBcrypt(10)
// Verify by bcrypt from signature byte slice without encoding and message byte slice, output bool
dongle.Verify.FromRawBytes([]byte(sign.ToRawBytes()), []byte("hello world")).ByBcrypt(10).ToBool() // true
```

##### Ed25519 sign and verify

```go
var publicKey, privateKey []byte

// generate raw public key and private key
publicKey, privateKey, _ = ed25519.GenerateKey(nil)

// get public key with hex encoding
hexPublicKey := dongle.Encode.FromBytes(publicKey).ByHex().ToBytes()
// get private key with hex encoding
hexPrivateKey := dongle.Encode.FromBytes(privateKey).ByHex().ToBytes()
// get public key with base64 encoding
base64PublicKey := dongle.Encode.FromBytes(publicKey).ByBase64().ToBytes()
// get private key with base64 encoding
base64PrivateKey := dongle.Encode.FromBytes(privateKey).ByBase64().ToBytes()

// Sign by ed25519 use raw private key
sign := dongle.Sign.FromString("hello world").ByEd25519(privateKey, dongle.Raw)
// Verify by ed25519 from raw signature string and message string use raw public key, output bool
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByEd25519(publicKey, dongle.Raw).ToBool() // true
// Verify by ed25519 from hex signature string and message string use raw public key, output bool
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByEd25519(publicKey, dongle.Raw).ToBool() // true
// Verify by ed25519 from base64 signature string and message string use raw public key, output bool
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByEd25519(publicKey, dongle.Raw).ToBool() // true

// Sign by ed25519 use hex private key
sign := dongle.Sign.FromString("hello world").ByEd25519(hexPrivateKey, dongle.Hex)
// Verify by ed25519 from raw signature string and message string use hex public key, output bool
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByEd25519(hexPublicKey, dongle.Hex).ToBool() // true
// Verify by ed25519 from hex signature string and message string use hex public key, output bool
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByEd25519(hexPublicKey, dongle.Hex).ToBool() // true
// Verify by ed25519 from base64 signature string and message string use hex public key, output bool
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByEd25519(hexPublicKey, dongle.Hex).ToBool() // true

// Sign by ed25519 use base64 private key
sign := dongle.Sign.FromString("hello world").ByEd25519(base64PrivateKey, dongle.Base64)
// Verify by ed25519 from raw signature string and message string use base64 public key, output bool
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByEd25519(base64PublicKey, dongle.Base64).ToBool() // true
// Verify by ed25519 from hex signature string and message string use base64 public key, output bool
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByEd25519(base64PublicKey, dongle.Base64).ToBool() // true
// Verify by ed25519 from base64 signature string and message string use base64 public key, output bool
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByEd25519(base64PublicKey, dongle.Base64).ToBool() // true

```

##### Rsa sign and verify

> Hash algorithm only supports MD5, SHA1, SHA224, SHA256, SHA384, SHA512, RIPEMD160

```go
pkcs1PublicKey := []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`)

pkcs1PrivateKey := []byte(`-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`)

pkcs8PublicKey := []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`)

pkcs8PrivateKey := []byte(`-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`)

// Sign by rsa from string
sign := dongle.Sign.FromString("hello world").ByRsa(pkcs1PrivateKey, dongle.MD5)
// Verify by rsa from raw signature string without encoding and message string, output bool
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByRsa(pkcs1PublicKey, dongle.MD5).ToBool() // true
// Verify by rsa from signature string with hex encoding and message string, output bool
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByRsa(pkcs1PublicKey, dongle.MD5).ToBool() // true
// Verify by rsa from signature string with base64 encoding and message string, output bool
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByRsa(pkcs1PublicKey, dongle.MD5).ToBool() // true

// Sign by rsa from byte slice
sign := dongle.Sign.FromBytes([]byte("hello world")).ByRsa([]byte(pkcs8PrivateKey), dongle.SHA512).ToRawBytes()
// Verify by rsa from raw signature byte slice without encoding and message byte slice, output bool
dongle.Verify.FromRawBytes(sign.ToRawBytes(), []byte("hello world")).ByRsa(pkcs1PublicKey, dongle.SHA512).ToBool() // true
// Verify by rsa from raw signature byte slice with hex encoding and message byte slice, output bool
dongle.Verify.FromHexBytes(sign.ToHexBytes(), []byte("hello world")).ByRsa(pkcs1PublicKey, dongle.SHA512).ToBool() // true
// Verify by rsa from raw signature byte slice with base64 encoding and message byte slice, output bool
dongle.Verify.FromBase64Bytes(sign.ToBase64Bytes(), []byte("hello world")).ByRsa(pkcs1PublicKey, dongle.SHA512).ToBool() // true
```

#### OpenSSL

##### RSA key

```go
// Generate PKCS#1 format RSA key pair
publicKey, privateKey := openssl.RSA.GenKeyPair(opensssl.PKCS1, 1024)
// Generate PKCS#8 format RSA key pair
publicKey, privateKey := openssl.RSA.GenKeyPair(opensssl.PKCS8, 2048)

// Verify RSA key pair matches
openssl.RSA.VerifyKeyPair(publicKey, privateKey) // true
openssl.RSA.VerifyKeyPair(publicKey, []byte("xxx")) // false

// Whether is RSA public key
openssl.RSA.IsPublicKey(publicKey) // true
openssl.RSA.IsPublicKey(privateKey) // false

// Whether is RSA private key
openssl.RSA.IsPrivateKey(privateKey) // true
openssl.RSA.IsPrivateKey(publicKey) // false

// Parse RSA public key
pub, err := openssl.RSA.ParsePublicKey(publicKey)
// Parse RSA private key
pri, err := openssl.RSA.ParsePrivateKey(privateKey)

// Format PKCS#1 RSA public key, add header, tail and newline character
openssl.RSA.FormatPublicKey(openssl.PKCS1, publicKey1)
// Format PKCS#8 RSA public key, add header, tail and newline character
openssl.RSA.FormatPublicKey(openssl.PKCS8, publicKey8)
// Format PKCS#1 RSA private key, add header, tail and newline character
openssl.RSA.FormatPrivateKey(openssl.PKCS1, privateKey1)
// Format PKCS#8 RSA private key, add header, tail and newline character
openssl.RSA.FormatPrivateKey(openssl.PKCS8, privateKey8)

// Compress RSA key, remove head, tail and newline character
keyBytes, err := openssl.RSA.CompressKey(key)

// Export RSA public key from RSA private key
publicKey, err := openssl.RSA.ExportPublicKey(privateKey)

```

### Error handling

> If more than one error occurs, only the first error is returned

```go
e := dongle.Encrypt.FromString("hello world").ByRsa([]byte("xxxx"))
if e.Error != nil {
    // Error handle...
    log.Fatal(e.Error)
}
fmt.Println(e.ToString())
// Output
rsa: invalid public key, please make sure the public key is valid
```

### References

* [openssl/openssl](https://github.com/openssl/openssl)
* [javascript/crypto-js](https://github.com/brix/crypto-js)
* [nodejs/crypto](https://nodejs.org/api/crypto.html)
* [java/jasypt](https://github.com/jasypt/jasypt)

### Sponsors

`Dongle` is a non-commercial open source project. If you want to support `Dongle`, you can [buy a cup of coffee](https://opencollective.com/go-carbon) for developer.

### Thanks

`Dongle` had been being developed with GoLand under the free JetBrains Open Source license, I would like to express my thanks here.

<a href="https://www.jetbrains.com"><img src="https://github-oss.oss-cn-beijing.aliyuncs.com/jetbrains.png" height="100" alt="JetBrains"/></a>