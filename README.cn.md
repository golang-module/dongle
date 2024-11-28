# dongle  #

[![Dongle Release](https://img.shields.io/github/release/dromara/dongle.svg)](https://github.com/dromara/dongle/releases)
[![Go Build](https://github.com/dromara/dongle/actions/workflows/test.yml/badge.svg)](https://github.com/dromara/dongle/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/dromara/dongle)](https://goreportcard.com/report/github.com/dromara/dongle)
[![codecov](https://codecov.io/gh/dromara/dongle/branch/main/graph/badge.svg)](https://codecov.io/gh/dromara/dongle)
[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/dromara/dongle)
![License](https://img.shields.io/github/license/dromara/dongle)

简体中文 | [English](README.md)

### 项目简介
一个轻量级、语义化、对开发者友好的 golang 编码解码、加密解密和签名验签库，已被 [awesome-go](https://github.com/avelino/awesome-go#security "awesome-go") 收录

### 仓库地址

[github.com/dromara/dongle](https://github.com/dromara/dongle "github.com/dromara/dongle")

[gitee.com/dromara/dongle](https://gitee.com/dromara/dongle "gitee.com/dromara/dongle")

[gitcode.com/dromara/dongle](https://gitcode.com/dromara/dongle "gitcode.com/dromara/dongle")

### 安装使用

```go
// 使用 github 库
go get -u github.com/dromara/dongle
import "github.com/dromara/dongle"

// 使用 gitee 库
go get -u gitee.com/dromara/dongle
import "gitee.com/dromara/dongle"

// 使用 gitcode 库
go get -u gitcode.com/dromara/dongle
import "gitcode.com/dromara/dongle"
```

`Dongle` 已经捐赠给了 [dromara](https://dromara.org/ "dromara") 组织，仓库地址发生了改变，如果之前用的仓库地址是 `golang-module/dongle`，请在 `go.mod` 里将原地址更换为新地址，如

```go
go mod edit -replace github.com/golang-module/dongle=github.com/dromara/dongle
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

##### Base45 编码、解码

```go
// 对字符进行 base45 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase45().ToString() // +8D VD82EK4F.KEA2
// 对字符串进行 base45 解码，输出字符串
dongle.Decode.FromString("+8D VD82EK4F.KEA2").ByBase45().ToString() // hello world

// 对字节切片进行 base45 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase45().ToBytes() // []byte("+8D VD82EK4F.KEA2")
// 对字节切片进行 base45 解码，输出字节切片
dongle.Decode.FromBytes([]byte("+8D VD82EK4F.KEA2")).ByBase45().ToBytes() // []byte("hello world")
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

> 可自定义编码表，编码表必须是62位，默认编码表是 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz

```go
// 对字符进行 base62 编码，输出字符串
dongle.Encode.FromString("hello world").ByBase62().ToString() // AAwf93rvy4aWQVw
// 对字符串进行 base62 解码，输出字符串
dongle.Decode.FromString("AAwf93rvy4aWQVw").ByBase62().ToString() // hello world

// 对字节切片进行 base62 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase62().ToBytes() // []byte("AAwf93rvy4aWQVw")
// 对字节切片进行 base62 解码，输出字节切片
dongle.Decode.FromBytes([]byte("AAwf93rvy4aWQVw")).ByBase62().ToBytes() // []byte("hello world")

// 自定义编码表对字节切片进行 base62 编码，输出字节切片
dongle.Encode.FromBytes([]byte("hello world")).ByBase62("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789").ToBytes() // []byte("KK6pJD158Ekgaf6")
// 自定义编码表对字节切片进行 base62 解码，输出字节切片
dongle.Decode.FromBytes([]byte("KK6pJD158Ekgaf6")).ByBase62("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789").ToBytes() // []byte("hello world")
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

##### Morse 编码、解码

> 默认分隔符是 `/`

```go
// 对字符串进行 morse 编码，输出字符串
dongle.Encode.FromString("dongle").ByMorse().ToString() // -../---/-./--./.-../.
// 对字符串进行 morse 解码，输出字符串
dongle.Decode.FromString("-../---/-./--./.-../.").ByMorse().ToString() // dongle

// 对字节切片进行 morse 编码，输出字节切片
dongle.Encode.FromBytes([]byte("dongle")).ByMorse("|").ToBytes() // []byte("-..|---|-.|--.|.-..|.")
// 对字节切片进行 morse 解码，输出字节切片
dongle.Decode.FromBytes([]byte("-..|---|-.|--.|.-..|.")).ByMorse("|").ToBytes() // []byte("dongle")
```

#### 加密&解密

##### Md2 加密

```go
// 对字符串进行 md2 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd2().ToHexString() // d9cce882ee690a5c1ce70beff3a78c77
// 对字符串进行 md2 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByMd2().ToBase64String() // 2czogu5pClwc5wvv86eMdw==

// 对字节切片进行 md2 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd2().ToHexBytes() // []byte("d9cce882ee690a5c1ce70beff3a78c77")
// 对字节切片进行 md2 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByMd2().ToBase64Bytes() // []byte("2czogu5pClwc5wvv86eMdw==")

```

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

##### Sha3 加密

> 包含 sha3-224, sha3-256, sha3-384, sha3-512

```go
// 对字符串进行 sha3-224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(224).ToHexString() // dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5
// 对字符串进行 sha3-224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(224).ToBase64String() // 37fxjHfpKLtW+ustonKRvXkLwQRc3kXzIQu2xQ==
// 对字节切片进行 sha3-224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(224).ToHexBytes() // []byte("dfb7f18c77e928bb56faeb2da27291bd790bc1045cde45f3210bb6c5")
// 对字节切片进行 sha3-224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(224).ToBase64Bytes() // []byte("37fxjHfpKLtW+ustonKRvXkLwQRc3kXzIQu2xQ==")

// 对字符串进行 sha3-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(256).ToHexString() // 644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938
// 对字符串进行 sha3-256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(256).ToBase64String() // ZEvMflZDcwQJmarInnYi88px+6HZcv2Uoxw7+/JOOTg=
// 对字节切片进行 sha3-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(256).ToHexBytes() // []byte("644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938")
// 对字节切片进行 sha3-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(256).ToBase64Bytes() // []byte("ZEvMflZDcwQJmarInnYi88px+6HZcv2Uoxw7+/JOOTg=")

// 对字符串进行 sha3-384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(384).ToHexString() // 83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b
// 对字符串进行 sha3-384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(384).ToBase64String() // g7/yjd4bG/WBAHHGZDwI5bBb24Nu/9cLQD6o6gpjTcSZfrEFOqNZP1kPnGNjDdkL
// 对字节切片进行 sha3-384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(384).ToHexBytes() // []byte("83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa3593f590f9c63630dd90b")
// 对字节切片进行 sha3-384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(384).ToBase64Bytes() // []byte("g7/yjd4bG/WBAHHGZDwI5bBb24Nu/9cLQD6o6gpjTcSZfrEFOqNZP1kPnGNjDdkL")

// 对字符串进行 sha3-512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(512).ToHexString() // 840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a
// 对字符串进行 sha3-512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").Sha3(512).ToBase64String() // hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg==
// 对字节切片进行 sha3-512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(512).ToHexBytes() // []byte("840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a")
// 对字节切片进行 sha3-512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).Sha3(512).ToBase64Bytes() // []byte("hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg==")

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

##### Sha512 加密

> 包含 sha512, sha512-224, sha512-256

```go
// 对字符串进行 sha512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToHexBytes() // 309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f
// 对字符串进行 sha512 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512().ToBase64String() // MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==
// 对字节切片进行 sha512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToHexBytes() // []byte("309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f")
// 对字节切片进行 sha512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512().ToBase64Bytes() // []byte("MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==")

// 对字符串进行 sha512-224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512(224).ToHexBytes() // 22e0d52336f64a998085078b05a6e37b26f8120f43bf4db4c43a64ee
// 对字符串进行 sha512-224 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512(224).ToBase64String() // IuDVIzb2SpmAhQeLBabjeyb4Eg9Dv020xDpk7g==
// 对字节切片进行 sha512-224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(224).ToHexBytes() // []byte("22e0d52336f64a998085078b05a6e37b26f8120f43bf4db4c43a64ee")
// 对字节切片进行 sha512-224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(224).ToBase64Bytes() // []byte("IuDVIzb2SpmAhQeLBabjeyb4Eg9Dv020xDpk7g==")

// 对字符串进行 sha512-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512(256).ToHexBytes() // 0ac561fac838104e3f2e4ad107b4bee3e938bf15f2b15f009ccccd61a913f017
// 对字符串进行 sha512-256 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").BySha512(256).ToBase64String() // CsVh+sg4EE4/LkrRB7S+4+k4vxXysV8AnMzNYakT8Bc=
// 对字节切片进行 sha512-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(256).ToHexBytes() // []byte("0ac561fac838104e3f2e4ad107b4bee3e938bf15f2b15f009ccccd61a913f017")
// 对字节切片进行 sha512-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).BySha512(256).ToBase64Bytes() // []byte("CsVh+sg4EE4/LkrRB7S+4+k4vxXysV8AnMzNYakT8Bc=")
```

##### Shake128 加密

```go
// 对字符串进行 shake128-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake128(256).ToHexString() // 3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b8
// 对字符串进行 shake128-256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake128(256).ToBase64String() // OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLg=
// 对字节切片进行 shake128-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(256).ToHexBytes() // []byte("3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b8")
// 对字节切片进行 shake128-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(256).ToBase64Bytes() // []byte("OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLg=")

// 对字符串进行 shake128-512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake128(512).ToHexString() // 3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b899072665674f26cc494a4bcf027c58267e8ee2da60e942759de86d2670bba1aa
// 对字符串进行 shake128-512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake128(512).ToBase64String() // OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLiZByZlZ08mzElKS88CfFgmfo7i2mDpQnWd6G0mcLuhqg==
// 对字节切片进行 shake128-512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(512).ToHexBytes() // []byte("3a9159f071e4dd1c8c4f968607c30942e120d8156b8b1e72e0d376e8871cb8b899072665674f26cc494a4bcf027c58267e8ee2da60e942759de86d2670bba1aa")
// 对字节切片进行 shake128-512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake128(512).ToBase64Bytes() // []byte("OpFZ8HHk3RyMT5aGB8MJQuEg2BVrix5y4NN26IccuLiZByZlZ08mzElKS88CfFgmfo7i2mDpQnWd6G0mcLuhqg==")
```

##### Shake256 加密

```go
// 对字符串进行 shake256-384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake256(384).ToHexString() // 369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df2
// 对字符串进行 shake256-384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake256(384).ToBase64String() // Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3y
// 对字节切片进行 shake256-384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(384).ToHexBytes() // []byte("369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df2")
// 对字节切片进行 shake256-384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(384).ToBase64Bytes() // []byte("Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3y")

// 对字符串进行 shake256-512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake256(512).ToHexString() // 369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df27853a8735271f5cc2d9e889544357116
// 对字符串进行 shake256-512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByShake256(512).ToBase64String() // Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3yeFOoc1Jx9cwtnoiVRDVxFg==
// 对字节切片进行 shake256-512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(512).ToHexBytes() // []byte("369771bb2cb9d2b04c1d54cca487e372d9f187f73f7ba3f65b95c8ee7798c527f4f3c2d55c2d46a29f2e945d469c3df27853a8735271f5cc2d9e889544357116")
// 对字节切片进行 shake256-512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByShake256(512).ToBase64Bytes() // []byte("Npdxuyy50rBMHVTMpIfjctnxh/c/e6P2W5XI7neYxSf088LVXC1Gop8ulF1GnD3yeFOoc1Jx9cwtnoiVRDVxFg==")
```

##### Ripemd160 加密

```go
// 对字符串进行 ripemd160 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByRipemd160().ToHexString() // 98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f
// 对字符串进行 ripemd160 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByRipemd160().ToBase64String() // mMYVeEzLX+WTb7wMvp39tAjZLw8=

// 对字节切片进行 ripemd160 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRipemd160().ToHexBytes() // []byte("98c615784ccb5fe5936fbc0cbe9dfdb408d92f0f")
// 对字节切片进行 ripemd160 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRipemd160().ToBase64Bytes() // []byte("mMYVeEzLX+WTb7wMvp39tAjZLw8=")
```

##### Blake2b 加密
> 包含 blake2b-256, blake2b-384, blake2b-512

```go
// 对字符串进行 blake2b-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2b(256).ToHexBytes() // 256c83b297114d201b30179f3f0ef0cace9783622da5974326b436178aeef610
// 对字符串进行 blake2b-256 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2b(256).ToBase64String() // JWyDspcRTSAbMBefPw7wys6Xg2ItpZdDJrQ2F4ru9hA=
// 对字节切片进行 blake2b-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(256).ToHexBytes() // []byte("256c83b297114d201b30179f3f0ef0cace9783622da5974326b436178aeef610")
// 对字节切片进行 blake2b-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(256).ToBase64Bytes() // []byte("JWyDspcRTSAbMBefPw7wys6Xg2ItpZdDJrQ2F4ru9hA=")

// 对字符串进行 blake2b-384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2b(384).ToHexBytes() // 8c653f8c9c9aa2177fb6f8cf5bb914828faa032d7b486c8150663d3f6524b086784f8e62693171ac51fc80b7d2cbb12b
// 对字符串进行 blake2b-384 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2b(384).ToBase64String() // jGU/jJyaohd/tvjPW7kUgo+qAy17SGyBUGY9P2UksIZ4T45iaTFxrFH8gLfSy7Er
// 对字节切片进行 blake2b-384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(384).ToHexBytes() // []byte("8c653f8c9c9aa2177fb6f8cf5bb914828faa032d7b486c8150663d3f6524b086784f8e62693171ac51fc80b7d2cbb12b")
// 对字节切片进行 blake2b-384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(384).ToBase64Bytes() // []byte("jGU/jJyaohd/tvjPW7kUgo+qAy17SGyBUGY9P2UksIZ4T45iaTFxrFH8gLfSy7Er")

// 对字符串进行 blake2b-512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2b(512).ToHexBytes() // 021ced8799296ceca557832ab941a50b4a11f83478cf141f51f933f653ab9fbcc05a037cddbed06e309bf334942c4e58cdf1a46e237911ccd7fcf9787cbc7fd0
// 对字符串进行 blake2b-512 加密，输出经过 base6 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2b(512).ToBase64String() // Ahzth5kpbOylV4MquUGlC0oR+DR4zxQfUfkz9lOrn7zAWgN83b7QbjCb8zSULE5YzfGkbiN5EczX/Pl4fLx/0A==
// 对字节切片进行 blake2b-512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(512).ToHexBytes() // []byte("021ced8799296ceca557832ab941a50b4a11f83478cf141f51f933f653ab9fbcc05a037cddbed06e309bf334942c4e58cdf1a46e237911ccd7fcf9787cbc7fd0")
// 对字节切片进行 blake2b-512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2b(512).ToBase64Bytes() // []byte("Ahzth5kpbOylV4MquUGlC0oR+DR4zxQfUfkz9lOrn7zAWgN83b7QbjCb8zSULE5YzfGkbiN5EczX/Pl4fLx/0A==")
```

##### Blake2s 加密
> 目前仅支持 blake2s-256

```go
// 对字符串进行 blake2s-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2s(256).ToHexString() // 9aec6806794561107e594b1f6a8a6b0c92a0cba9acf5e5e93cca06f781813b0b
// 对字符串进行 blake2s-256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlake2s(256).ToBase64String() // muxoBnlFYRB+WUsfaoprDJKgy6ms9eXpPMoG94GBOws=

// 对字节切片进行 blake2s-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2s(256).ToHexBytes() // []byte("9aec6806794561107e594b1f6a8a6b0c92a0cba9acf5e5e93cca06f781813b0b")
// 对字节切片进行 blake2s-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlake2s(256).ToBase64Bytes() // []byte("muxoBnlFYRB+WUsfaoprDJKgy6ms9eXpPMoG94GBOws=")
```

##### Hmac-md2 加密

```go
// 对字符串进行 hmac-md2 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd2([]byte("dongle")).ToHexString() // 88ed6ef9ab699d03a702f2a6fb1c0673
// 对字符串进行 hmac-md2 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd2([]byte("dongle")).ToBase64String() // iO1u+atpnQOnAvKm+xwGcw==

// 对字节切片进行 hmac-md2 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd2([]byte("dongle")).ToHexBytes() // []byte("88ed6ef9ab699d03a702f2a6fb1c0673")
// 对字节切片进行 hmac-md2 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd2([]byte("dongle")).ToBase64Bytes() // []byte("iO1u+atpnQOnAvKm+xwGcw==")
```

##### Hmac-md4 加密

```go
// 对字符串进行 hmac-md4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4([]byte("dongle")).ToHexString() // 7a9df5247cbf76a8bc17c9c4f5a75b6b
// 对字符串进行 hmac-md4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd4([]byte("dongle")).ToBase64String() // ep31JHy/dqi8F8nE9adbaw==

// 对字节切片进行 hmac-md4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToHexBytes() // []byte("7a9df5247cbf76a8bc17c9c4f5a75b6b")
// 对字节切片进行 hmac-md4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd4([]byte("dongle")).ToBase64Bytes() // []byte("ep31JHy/dqi8F8nE9adbaw==")
```

##### Hmac-md5 加密

```go
// 对字符串进行 hmac-md5 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5([]byte("dongle")).ToHexString() // 4790626a275f776956386e5a3ea7b726
// 对字符串进行 hmac-md5 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacMd5([]byte("dongle")).ToBase64String() // R5Biaidfd2lWOG5aPqe3Jg==

// 对字节切片进行 hmac-md5 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToHexBytes() // []byte("4790626a275f776956386e5a3ea7b726")
// 对字节切片进行 hmac-md5 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacMd5([]byte("dongle")).ToBase64Bytes() // []byte("R5Biaidfd2lWOG5aPqe3Jg==")
```

##### Hmac-sha1 加密

```go
// 对字符串进行 hmac-sha1 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1([]byte("dongle")).ToHexString() // 91c103ef93ba7420902b0d1bf0903251c94b4a62
// 对字符串进行 hmac-sha1 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha1([]byte("dongle")).ToBase64String() // kcED75O6dCCQKw0b8JAyUclLSmI=

// 对字节切片进行 hmac-sha1 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToHexBytes() // []byte("91c103ef93ba7420902b0d1bf0903251c94b4a62")
// 对字节切片进行 hmac-sha1 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha1([]byte("dongle")).ToBase64Bytes() // []byte("kcED75O6dCCQKw0b8JAyUclLSmI=")
```

##### Hmac-sha3 加密

> 包含 hmac-sha3-224, hmac-sha3-256, hmac-sha3-384, hmac-sha3-512

```go
// 对字符串进行 hmac-sha3-224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 224).ToHexString() // fb8f061d9d1dddd2f5d3b9064a5e98e3e4b6df27ea93ce67627583ce
// 对字符串进行 hmac-sha3-224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 224).ToBase64String() // +48GHZ0d3dL107kGSl6Y4+S23yfqk85nYnWDzg==
// 对字节切片进行 hmac-sha3-224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 224).ToHexBytes() // []byte("fb8f061d9d1dddd2f5d3b9064a5e98e3e4b6df27ea93ce67627583ce")
// 对字节切片进行 hmac-sha3-224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 224).ToBase64Bytes() // []byte("+48GHZ0d3dL107kGSl6Y4+S23yfqk85nYnWDzg==")

// 对字符串进行 hmac-sha3-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 256).ToHexString() // 8193367fde28cf5c460adb449a04b3dd9c184f488bdccbabf0526c54f90c4460
// 对字符串进行 hmac-sha3-256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 256).ToBase64String() // gZM2f94oz1xGCttEmgSz3ZwYT0iL3Mur8FJsVPkMRGA=
// 对字节切片进行 hmac-sha3-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 256).ToHexBytes() // []byte("8193367fde28cf5c460adb449a04b3dd9c184f488bdccbabf0526c54f90c4460")
// 对字节切片进行 hmac-sha3-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 256).ToBase64Bytes() // []byte("gZM2f94oz1xGCttEmgSz3ZwYT0iL3Mur8FJsVPkMRGA=")

// 对字符串进行 hmac-sha3-384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 384).ToHexString() // 3f76f5cda69cada3ee6b33f8458cd498b063075db263dd8b33f2a3992a8804f9569a7c86ffa2b8f0748babeb7a6fc0e7
// 对字符串进行 hmac-sha3-384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 384).ToBase64String() // P3b1zaacraPuazP4RYzUmLBjB12yY92LM/KjmSqIBPlWmnyG/6K48HSLq+t6b8Dn
// 对字节切片进行 hmac-sha3-384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 384).ToHexBytes() // []byte("3f76f5cda69cada3ee6b33f8458cd498b063075db263dd8b33f2a3992a8804f9569a7c86ffa2b8f0748babeb7a6fc0e7")
// 对字节切片进行 hmac-sha3-384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 384).ToBase64Bytes() // []byte("P3b1zaacraPuazP4RYzUmLBjB12yY92LM/KjmSqIBPlWmnyG/6K48HSLq+t6b8Dn")

// 对字符串进行 hmac-sha3-512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 512).ToHexString() // a99653d0407d659eccdeed43bb7cccd2e2b05a2c34fd3467c4198cf2ad26a466738513e88839fb55e64eb49df65bc52ed0fec2775bd9e086edd4fb4024add4a2
// 对字符串进行 hmac-sha3-512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha3([]byte("dongle"), 512).ToBase64String() // qZZT0EB9ZZ7M3u1Du3zM0uKwWiw0/TRnxBmM8q0mpGZzhRPoiDn7VeZOtJ32W8Uu0P7Cd1vZ4Ibt1PtAJK3Uog==
// 对字节切片进行 hmac-sha3-512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 512).ToHexBytes() // []byte("a99653d0407d659eccdeed43bb7cccd2e2b05a2c34fd3467c4198cf2ad26a466738513e88839fb55e64eb49df65bc52ed0fec2775bd9e086edd4fb4024add4a2")
// 对字节切片进行 hmac-sha3-512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha3([]byte("dongle"), 512).ToBase64Bytes() // []byte("qZZT0EB9ZZ7M3u1Du3zM0uKwWiw0/TRnxBmM8q0mpGZzhRPoiDn7VeZOtJ32W8Uu0P7Cd1vZ4Ibt1PtAJK3Uog==")
```

##### Hmac-sha224 加密

```go
// 对字符串进行 hmac-sha224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224([]byte("dongle")).ToHexString() // e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec
// 对字符串进行 hmac-sha224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha224([]byte("dongle")).ToBase64String() // 4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A====

// 对字节切片进行 hmac-sha224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToHexBytes() // []byte("e15b9e5a7eccb1f17dc81dc07c909a891936dc3429dc0d940accbcec")
// 对字节切片进行 hmac-sha224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha224([]byte("dongle")).ToBase64Bytes() // []byte("4VueWn7MsfF9yB3AfJCaiRk23DQp3A2UCsy87A==")
```

##### Hmac-sha256 加密

```go
// 对字符串进行 hmac-sha256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256([]byte("dongle")).ToHexString() // b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
// 对字符串进行 hmac-sha256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha256([]byte("dongle")).ToBase64String() // uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=

// 对字节切片进行 hmac-sha256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToHexBytes() // []byte("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
// 对字节切片进行 hmac-sha256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha256([]byte("dongle")).ToBase64Bytes() // []byte("uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=")
```

##### Hmac-sha384 加密

```go
// 对字符串进行 hmac-sha384 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384([]byte("dongle")).ToHexString() // 421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8
// 对字符串进行 hmac-sha384 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha384([]byte("dongle")).ToBase64String() // Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o
// 对字节切片进行 hmac-sha384 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToHexBytes() // []byte("421fcaa740216a31bbcd1f86f2212e0c68aa4b156a8ebc2ae55b3e75c4ee0509ea0325a0570ae739006b61d91d817fe8")
// 对字节切片进行 hmac-sha384 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha384([]byte("dongle")).ToBase64Bytes() // []byte("Qh/Kp0AhajG7zR+G8iEuDGiqSxVqjrwq5Vs+dcTuBQnqAyWgVwrnOQBrYdkdgX/o")
```

##### Hmac-sha512 加密

> 包含 hmac-sha512, hmac-sha512-224, hmac-sha512-256

```go
// 对字符串进行 hmac-sha512 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle")).ToHexString() // d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1
// 对字符串进行 hmac-sha512 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle")).ToBase64String() // 2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==
// 对字节切片进行 hmac-sha512 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToHexBytes() // []byte("d971b790bbc2a4ac81062bbffac693c9c234bae176c8faf5e304dbdb153032a826f12353964b4a4fb87abecd2dc237638a630cbad54a6b94b1f6ef5d5e2835d1")
// 对字节切片进行 hmac-sha512 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle")).ToBase64Bytes() // []byte("2XG3kLvCpKyBBiu/+saTycI0uuF2yPr14wTb2xUwMqgm8SNTlktKT7h6vs0twjdjimMMutVKa5Sx9u9dXig10Q==")

// 对字符串进行 hmac-sha512-224 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 224).ToHexString() // c4145bcc385c29f0e5683cd5450be9deb522d556de3b046a7ffa1eb3
// 对字符串进行 hmac-sha512-224 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 224).ToBase64String() // xBRbzDhcKfDlaDzVRQvp3rUi1VbeOwRqf/oesw==
// 对字节切片进行 hmac-sha512-224 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 224).ToHexBytes() // []byte("c4145bcc385c29f0e5683cd5450be9deb522d556de3b046a7ffa1eb3")
// 对字节切片进行 hmac-sha512-224 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 224).ToBase64Bytes() // []byte("xBRbzDhcKfDlaDzVRQvp3rUi1VbeOwRqf/oesw==")

// 对字符串进行 hmac-sha512-256 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 256).ToHexString() // e99fae71bcb43651ae10e952989eadf897faccb43966ee5122bb1b1d82f7a7c2
// 对字符串进行 hmac-sha512-256 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSha512([]byte("dongle"), 256).ToBase64String() // 6Z+ucby0NlGuEOlSmJ6t+Jf6zLQ5Zu5RIrsbHYL3p8I=
// 对字节切片进行 hmac-sha512-256 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 256).ToHexBytes() // []byte("e99fae71bcb43651ae10e952989eadf897faccb43966ee5122bb1b1d82f7a7c2")
// 对字节切片进行 hmac-sha512-256 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSha512([]byte("dongle"), 256).ToBase64Bytes() // []byte("6Z+ucby0NlGuEOlSmJ6t+Jf6zLQ5Zu5RIrsbHYL3p8I=")
```

##### Hmac-ripemd160 加密

```go
// 对字符串进行 hmac-ripemd160 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacRipemd160([]byte("dongle")).ToHexString() // 3691ad040e80c43dc6e8ffe9bc6ef3d5bd8786b8
// 对字符串进行 hmac-ripemd160 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacRipemd160([]byte("dongle")).ToBase64String() // NpGtBA6AxD3G6P/pvG7z1b2Hhrg=

// 对字节切片进行 hmac-ripemd160 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacRipemd160([]byte("dongle")).ToHexBytes() // []byte("3691ad040e80c43dc6e8ffe9bc6ef3d5bd8786b8")
// 对字节切片进行 hmac-ripemd160 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacRipemd160([]byte("dongle")).ToBase64Bytes() // []byte("NpGtBA6AxD3G6P/pvG7z1b2Hhrg=")
```

##### Hmac-sm3 加密

```go
// 对字符串进行 hmac-sm3 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSm3([]byte("dongle")).ToHexString() // 8c733aae1d553c466a08c3e9e5daac3e99ae220181c7c1bc8c2564961de751b3
// 对字符串进行 hmac-sm3 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByHmacSm3([]byte("dongle")).ToBase64String() // jHM6rh1VPEZqCMPp5dqsPpmuIgGBx8G8jCVklh3nUbM=

// 对字节切片进行 hmac-sm3 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSm3([]byte("dongle")).ToHexBytes() // []byte("8c733aae1d553c466a08c3e9e5daac3e99ae220181c7c1bc8c2564961de751b3")
// 对字节切片进行 hmac-sm3 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByHmacSm3([]byte("dongle")).ToBase64Bytes() // []byte("jHM6rh1VPEZqCMPp5dqsPpmuIgGBx8G8jCVklh3nUbM=")
```

##### Aes 加密、解密

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、CFB、OFB、CTR、ECB
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("0123456789abcdef") // key 长度必须是 16、24 或 32 字节
// cipher.WithKey([]byte("0123456789abcdef")) 
cipher.SetIV("0123456789abcdef") // iv 长度必须是 16 字节，ECB 模式不需要设置 iv
// cipher.WithIV([]byte("0123456789abcdef")) 

// 对字符串进行 aes 加密，输出未经编码的原始字符串
rawString := dongle.Encrypt.FromString("hello world").ByAes(cipher).ToRawString()
// 对未经编码的原始字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromRawString(rawString).ByAes(cipher).ToString() // hello world

// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToHexString() // c1e9b4529aac9793010f4677f6358efe
// 对经过 hex 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByAes(cipher).ToString() // hello world

// 对字符串进行 aes 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByAes(cipher).ToBase64String() // wem0Upqsl5MBD0Z39jWO/g==
// 对经过 base64 编码的字符串进行 aes 解密，输出字符串
dongle.Decrypt.FromBase64String("wem0Upqsl5MBD0Z39jWO/g==").ByAes(cipher).ToString() // hello world

// 对字节切片进行 aes 加密，输出未经编码的原始字节切片
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToRawBytes()
// 对未经编码的原始字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromRawBytes(rawBytes).ByAes(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 aes 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToHexBytes() // []byte("c1e9b4529aac9793010f4677f6358efe")
// 对经过 hex 编码的字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromHexBytes([]byte("c1e9b4529aac9793010f4677f6358efe")).ByAes(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 aes 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByAes(cipher).ToBase64Bytes() // []byte("ZdgjvfHFgaHe0cukLgPKUg==")
// 对经过 base64 编码的字节切片进行 aes 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("wem0Upqsl5MBD0Z39jWO/g==")).ByAes(cipher).ToBytes() // []byte("hello world")
```

##### Blowfish 加密、解密

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、CFB、OFB、CTR、ECB
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("0123456789abcdef") // key 长度必须是 1-56 字节
// cipher.WithKey([]byte("0123456789abcdef")) 
cipher.SetIV("12345678"))         // iv 长度必须是 8 字节，ECB 模式不需要设置 iv
// cipher.WithIV([]byte("12345678"))          

// 对字符串进行 blowfish 加密，输出未经编码的原始字符串
rawString := dongle.Encrypt.FromString("hello world").ByBlowfish(cipher).ToRawString()
// 对未经编码的原始字符串进行 blowfish 解密，输出字符串
dongle.Decrypt.FromRawString(rawString).ByBlowfish(cipher).ToString() // hello world

// 对字符串进行 blowfish 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlowfish(cipher).ToHexString() // c1e9b4529aac9793010f4677f6358efe
// 对经过 hex 编码的字符串进行 blowfish 解密，输出字符串
dongle.Decrypt.FromHexString("c1e9b4529aac9793010f4677f6358efe").ByBlowfish(cipher).ToString() // hello world

// 对字符串进行 blowfish 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByBlowfish(cipher).ToBase64String() // ksNyTXILWZgtIaq5p7ufQA==
// 对经过 base64 编码的字符串进行 blowfish 解密，输出字符串
dongle.Decrypt.FromBase64String("ksNyTXILWZgtIaq5p7ufQA==").ByBlowfish(cipher).ToString() // hello world

// 对字节切片进行 blowfish 加密，输出未经编码的原始字节切片
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).ByBlowfish(cipher).ToRawBytes()
// 对未经编码的原始字节切片进行 blowfish 解密，输出字节切片
dongle.Decrypt.FromRawBytes(rawBytes).ByBlowfish(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 blowfish 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlowfish(cipher).ToHexBytes() // []byte("c1e9b4529aac9793010f4677f6358efe")
// 对经过 hex 编码的字节切片进行 blowfish 解密，输出字节切片
dongle.Decrypt.FromHexBytes([]byte("c1e9b4529aac9793010f4677f6358efe")).ByBlowfish(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 blowfish 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByBlowfish(cipher).ToBase64Bytes() // []byte("ksNyTXILWZgtIaq5p7ufQA==")
// 对经过 base64 编码的字节切片进行 blowfish 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("ksNyTXILWZgtIaq5p7ufQA==")).ByBlowfish(cipher).ToBytes() // []byte("hello world")
```

##### Des 加密、解密

```go
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC、ECB、CFB、OFB、CTR
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("12345678") // key 长度必须是 8 字节
// cipher.WithKey([]byte("12345678")) 
cipher.SetIV("12345678") // iv 长度必须是 8 字节
// cipher.WithIV([]byte("12345678")) 

// 对字符串进行 des 加密，输出未经编码的原始字符串
rawString := dongle.Encrypt.FromString("hello world").ByDes(cipher).ToRawString()
// 对未经编码的原始字符串进行 des 解密，输出字符串
dongle.Decrypt.FromRawString(rawString).ByDes(cipher).ToString() // hello world

// 对字符串进行 des 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
// 对经过 hex 编码的字符串进行 des 解密，输出字符串
dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(cipher).ToString() // hello world

// 对字符串进行 des 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByDes(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// 对经过 base64 编码的字符串进行 des 解密，输出字符串
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").ByDes(cipher).ToString() // hello world

// 对字节切片进行 des 加密，输出未经编码的原始字节切片
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).ByDes(cipher).ToRawBytes()
// 对未经编码的原始字节切片进行 des 解密，输出字节切片
dongle.Decrypt.FromRawBytes(rawBytes).ByDes(cipher).ToBytes() // []byte("hello world")

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
cipher := dongle.NewCipher()
cipher.SetMode(dongle.CBC) // CBC, CFB, CTR, ECB, OFB
cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
cipher.SetKey("123456781234567812345678") // key 长度必须是 24
// cipher.WithKey([]byte("123456781234567812345678"))
cipher.SetIV("12345678") // iv 长度必须是 8
// cipher.WithIV([]byte("12345678")) 

// 对字符串进行 3des 加密，输出未经编码的原始字符串
rawString := dongle.Encrypt.FromString("hello world").By3Des(cipher).ToRawString()
// 对未经编码的原始字符串进行 3des 解密，输出字符串
dongle.Decrypt.FromRawString(rawString).By3Des(cipher).ToString() // hello world

// 对字符串进行 3des 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToHexString() // 92c3724d720b59982d21aab9a7bb9f40
// 对经过 hex 编码的字符串进行 3des 解密，输出字符串
dongle.Decrypt.FromHexString("92c3724d720b59982d21aab9a7bb9f40").By3Des(cipher).ToString() // hello world

// 对字符串进行 3des 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").By3Des(cipher).ToBase64String() // CyqS6B+0nOGkMmaqyup7gQ==
// 对经过 base64 编码的字符串进行 3des 解密，输出字符串
dongle.Decrypt.FromBase64String("CyqS6B+0nOGkMmaqyup7gQ==").By3Des(cipher).ToString() // hello world

// 对字节切片进行 3des 加密，输出未经编码的原始字节切片
rawBytes := dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToRawBytes()
// 对未经编码的原始字节切片进行 3des 解密，输出字节切片
dongle.Decrypt.FromRawBytes(rawBytes).By3Des(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 3des 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToHexBytes() // []byte("92c3724d720b59982d21aab9a7bb9f40")
// 对经过 hex 编码的字节切片进行 3des 解密，输出字符串
dongle.Decrypt.FromHexBytes([]byte("92c3724d720b59982d21aab9a7bb9f40")).By3Des(cipher).ToBytes() // []byte("hello world")

// 对字节切片进行 3des 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).By3Des(cipher).ToBase64Bytes() // []byte("CyqS6B+0nOGkMmaqyup7gQ==")
// 对经过 base64 编码的字节切片进行 3des 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("CyqS6B+0nOGkMmaqyup7gQ==")).By3Des(cipher).ToBytes() // []byte("hello world")
```

##### Rsa 加密、解密

```go
pkcs1PublicKey :=  []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`)

pkcs1PrivateKey :=  []byte(`-----BEGIN RSA PRIVATE KEY-----
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

pkcs8PublicKey :=  []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`)

pkcs8PrivateKey :=  []byte(`-----BEGIN PRIVATE KEY-----
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

###### 1、RSA 公钥加密、私钥解密

```go
// 使用 pkcs1 格式公钥对字符串进行 rsa 加密
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs1PublicKey)
// 使用 pkcs1 格式私钥对经过 hex 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs1PrivateKey).ToString() // hello world
// 使用 pkcs1 格式私钥对经过 base64 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs1PrivateKey).ToString() // hello world

// 使用 pkcs1 格式公钥对字节切片进行 rsa 加密
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs1PublicKey)
// 使用 pkcs1 格式私钥对经过 hex 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs1PrivateKey).ToByte() // []bytes("hello world)
// 使用 pkcs1 格式私钥对经过 base64 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs1PrivateKey).ToByte() // []bytes("hello world)

// 使用 pkcs8 格式公钥对字符串进行 rsa 加密，输出经过 hex 编码的字符串
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs8PublicKey)
// 使用 pkcs8 格式私钥对经过 hex 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs8PrivateKey).ToString() // hello world
// 使用 pkcs8 格式私钥对经过 base64 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs8PrivateKey).ToString() // hello world

// 使用 pkcs8 格式公钥对字节切片进行 rsa 加密，输出经过 hex 编码的字节切片
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs8PublicKey)
// 使用 pkcs8 格式私钥对经过 hex 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs8PrivateKey).ToByte() // []bytes("hello world)
// 使用 pkcs8 格式私钥对经过 base64 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs8PrivateKey).ToByte() // []bytes("hello world)
```

###### 2、RSA 私钥加密、公钥解密

```go
// 使用 pkcs1 格式私钥对字符串进行 rsa 加密
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs1PrivateKey)
// 使用 pkcs1 格式公钥对经过 hex 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs1PublicKey).ToString() // hello world
// 使用 pkcs1 格式公钥对经过 base64 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs1PublicKey).ToString() // hello world

// 使用 pkcs1 格式私钥对字节切片进行 rsa 加密
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs1PrivateKey)
// 使用 pkcs1 格式公钥对经过 hex 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs1PublicKey).ToByte() // []bytes("hello world)
// 使用 pkcs1 格式公钥对经过 base64 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs1PublicKey).ToByte() // []bytes("hello world)

// 使用 pkcs8 格式私钥对字符串进行 rsa 加密，输出经过 hex 编码的字符串
cipherText := dongle.Encrypt.FromString("hello world").ByRsa(pkcs8PrivateKey)
// 使用 pkcs8 格式公钥对经过 hex 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromHexString(cipherText.ToHexString()).ByRsa(pkcs8PublicKey).ToString() // hello world
// 使用 pkcs8 格式公钥对经过 base64 编码的字符串进行 rsa 解密，输出字符串
dongle.Decrypt.FromBase64String(cipherText.ToBase64String()).ByRsa(pkcs8PublicKey).ToString() // hello world

// 使用 pkcs8 格式私钥对字节切片进行 rsa 加密，输出经过 hex 编码的字节切片
cipherText := dongle.Encrypt.FromBytes([]byte("hello world")).ByRsa(pkcs8PrivateKey)
// 使用 pkcs8 格式公钥对经过 hex 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromHexBytes(cipherText.ToHexBytes()).ByRsa(pkcs8PublicKey).ToByte() // []bytes("hello world)
// 使用 pkcs8 格式公钥对经过 base64 编码的字节切片进行 rsa 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(cipherText.ToBase64Bytes()).ByRsa(pkcs8PublicKey).ToByte() // []bytes("hello world)
```

##### Tea 加密、解密

> 加密明文如果超过 8 字节，进行空字符串填充分组解密；迭代轮次(rounds) 必须是偶数，默认是 64

```go
// 对字符串进行 tea 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByTea([]byte("0123456789abcdef")).ToHexString() // bfa8d956247c0dcecfe2b0ef44b08aab
// 对经过 hex 编码的字符串进行 tea 解密，输出字符串
dongle.Decrypt.FromHexString("bfa8d956247c0dcecfe2b0ef44b08aab").ByTea([]byte("0123456789abcdef")).ToString() // hello world

// 对字符串进行 tea 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByTea([]byte("0123456789abcdef")).ToBase64String() // v6jZViR8Dc7P4rDvRLCKqw==
// 对经过 base64 编码的字符串进行 tea 解密，输出字符串
dongle.Decrypt.FromBase64String("v6jZViR8Dc7P4rDvRLCKqw==").ByTea([]byte("0123456789abcdef")).ToString() // hello world

// 对字节切片进行 tea 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByTea([]byte("0123456789abcdef"), 32).ToHexBytes() // []byte("bfa8d956247c0dcecfe2b0ef44b08aab")
// 对经过 hex 编码的字节切片进行 tea 解密，输出字节切片
dongle.Decrypt.FromHexBytes([]byte("bfa8d956247c0dcecfe2b0ef44b08aab")).ByTea([]byte("0123456789abcdef"), 64).ToBytes() // []byte("hello world")

// 对字节切片进行 tea 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByTea([]byte("0123456789abcdef"), 32).ToBase64Bytes() // []byte("v6jZViR8Dc7P4rDvRLCKqw==")
// 对经过 base64 编码的字节切片进行 tea 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes(()byte("v6jZViR8Dc7P4rDvRLCKqw==")).ByTea([]byte("0123456789abcdef"), 64).ToBytes() // []byte("hello world")
```

##### Rc4 加密、解密

```go
// 对字符串进行 rc4 加密，输出经过 hex 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4([]byte("dongle")).ToHexString() // eba154b4cb5a9038dbbf9d
// 对经过 hex 编码的字符串进行 rc4 解密，输出字符串
dongle.Encrypt.FromHexString("eba154b4cb5a9038dbbf9d").ByRc4([]byte("dongle")).ToString() // hello world

// 对字符串进行 rc4 加密，输出经过 base64 编码的字符串
dongle.Encrypt.FromString("hello world").ByRc4([]byte("dongle")).ToBase64String() // 66FUtMtakDjbv50=
// 对经过 base64 编码的字符串进行 rc4 解密，输出字符串
dongle.Decrypt.FromBase64String("66FUtMtakDjbv50=").ByRc4([]byte("dongle")).ToString() // hello world

// 对字节切片进行 rc4 加密，输出经过 hex 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4([]byte("dongle")).ToHexBytes() // []byte("eba154b4cb5a9038dbbf9d")
// 对经过 hex 编码的字节切片进行 rc4 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes([]byte("hello world")).ByRc4([]byte("dongle")).ToBytes() // []byte("eba154b4cb5a9038dbbf9d")

// 对字节切片进行 rc4 加密，输出经过 base64 编码的字节切片
dongle.Encrypt.FromBytes([]byte("hello world")).ByRc4([]byte("dongle")).ToBase64Bytes() // []byte("66FUtMtakDjbv50=")
// 对经过 base64 编码的字节切片进行 rc4 解密，输出字节切片
dongle.Decrypt.FromBase64Bytes([]byte("66FUtMtakDjbv50=")).ByRc4([]byte("dongle")).ToBytes() // []byte("hello world")
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

#### 签名&验签

##### Bcrypt 签名、验签

> 迭代轮次(rounds) 取值范围 4-31，默认是 10

```go
// 对字符串进行 bcrypt 签名
sign := dongle.Sign.FromString("hello world").ByBcrypt()
// 对未经编码的原始签名字符串进行 bcrypt 验签
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByBcrypt().ToBool() // true
// 对经过 hex 编码的签名字符串进行 bcrypt 验签
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByBcrypt().ToBool() // true
// 对经过 base64 编码的签名字符串进行 bcrypt 验签
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByBcrypt().ToBool() // true

// 对字节切片进行 bcrypt 签名
sign := dongle.Sign.FromBytes([]byte("hello world")).ByBcrypt(10)
// 对未经编码的原始签名字节切片进行 bcrypt 验签
dongle.Verify.FromRawBytes([]byte(sign.ToRawBytes()), []byte("hello world")).ByBcrypt(10).ToBool() // true
// 对经过 hex 编码的签名字节切片进行 bcrypt 验签
dongle.Verify.FromHexBytes([]byte(sign.ToHexBytes()), []byte("hello world")).ByBcrypt(10).ToBool() // true
// 对经过 base64 编码的签名字节切片进行 bcrypt 验签
dongle.Verify.FromBase64Bytes([]byte(sign.ToBase64Bytes()), []byte("hello world")).ByBcrypt(10).ToBool() // true
```

##### Ed25519 签名、验签

```go
var publicKey, privateKey []byte

// 生成未经编码的原始公钥、私钥
publicKey, privateKey, _ = ed25519.GenerateKey(nil)

// 获取经过 hex 编码的公钥
hexPublicKey := dongle.Encode.FromBytes(publicKey).ByHex().ToBytes()
// 获取经过 hex 编码的私钥
hexPrivateKey := dongle.Encode.FromBytes(privateKey).ByHex().ToBytes()
// 获取经过 base64 编码的公钥
base64PublicKey := dongle.Encode.FromBytes(publicKey).ByBase64().ToBytes()
// 获取经过 base64 编码的私钥
base64PrivateKey := dongle.Encode.FromBytes(privateKey).ByBase64().ToBytes()

// 通过未经编码的原始私钥对字符串进行 ed25519 签名
sign := dongle.Sign.FromString("hello world").ByEd25519(privateKey, dongle.Raw)
// 通过未经编码的原始公钥对未经编码的原始签名字符串进行 ed25519 验签
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByEd25519(publicKey, dongle.Raw).ToBool() // true
// 通过未经编码的原始公钥对经过 hex 编码的签名字符串进行 ed25519 验签
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByEd25519(publicKey, dongle.Raw).ToBool() // true
// 通过未经编码的原始公钥对经过 base64 编码的签名字符串进行 ed25519 验签
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByEd25519(publicKey, dongle.Raw).ToBool() // true

// 通过经过 hex 编码的私钥对字符串进行 ed25519 签名
sign := dongle.Sign.FromString("hello world").ByEd25519(hexPrivateKey, dongle.Hex)
// 通过经过 hex 编码的公钥对未经编码的原始签名字符串进行 ed25519 验签
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByEd25519(hexPublicKey, dongle.Hex).ToBool() // true
// 通过经过 hex 编码的公钥对经过 hex 编码的签名字符串进行 ed25519 验签
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByEd25519(hexPublicKey, dongle.Hex).ToBool() // true
// 通过经过 hex 编码的公钥对经过 base64 编码的签名字符串进行 ed25519 验签
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByEd25519(hexPublicKey, dongle.Hex).ToBool() // true

// 通过经过 base64 编码的私钥对字符串进行 ed25519 签名
sign := dongle.Sign.FromString("hello world").ByEd25519(base64PrivateKey, dongle.Base64)
// 通过经过 base64 编码的公钥对未经编码的原始签名字符串进行 ed25519 验签
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByEd25519(base64PublicKey, dongle.Base64).ToBool() // true
// 通过经过 base64 编码的公钥对经过 hex 编码的签名字符串进行 ed25519 验签
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByEd25519(base64PublicKey, dongle.Base64).ToBool() // true
// 通过经过 base64 编码的公钥对经过 base64 编码的签名字符串进行 ed25519 验签
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByEd25519(base64PublicKey, dongle.Base64).ToBool() // true

```

##### Rsa 签名、验签

> hash 算法仅支持 MD5, SHA1, SHA224, SHA256, SHA384, SHA512, RIPEMD160

```go
pkcs1PublicKey := []byte(`-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAK12MTd84qkCZzp4iLUj8YSUglaFMsFlv9KlIL4+Xts40PK3+wbsXPEw
cujGeUmdgMeZiK7SLLSz8QeE0v7Vs+cGK4Bs4qLtMGCiO6wEuyt10KsafTyBktFn
dk/+gBLr7B/b+9+HaMIIoJUdsFksdAg3cxTSpwVApe98loFNRfqDAgMBAAE=
-----END RSA PUBLIC KEY-----`)

pkcs1PrivateKey :=  []byte(`-----BEGIN RSA PRIVATE KEY-----
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

pkcs8PublicKey :=  []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCqzZNa9VrcewyU6wDoV7Y9kAHq
X1VK0B3Rb6GNmQe4zLEfce7cVTaLrc4VGTKl35tADG1cRHqtaG4S/WttpiGZBhxJ
y4MpOXb6eIPiVLsn2lL+rJo5XdbSr3gyjxEOQQ97ihtw4lDd5wMo4bIOuw1LtMez
HC1outlM6x+/BB0BSQIDAQAB
-----END PUBLIC KEY-----`)

pkcs8PrivateKey :=  []byte(`-----BEGIN PRIVATE KEY-----
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

// 对字符串进行 rsa 签名
sign := dongle.Sign.FromString("hello world").ByRsa(pkcs1PrivateKey, dongle.MD5)
// 对未经编码的原始签名字符串进行 rsa 验签
dongle.Verify.FromRawString(sign.ToRawString(), "hello world").ByRsa(pkcs1PublicKey, dongle.MD5).ToBool() // true
// 对经过 hex 编码的签名字符串进行 rsa 验签
dongle.Verify.FromHexString(sign.ToHexString(), "hello world").ByRsa(pkcs1PublicKey, dongle.MD5).ToBool() // true
// 对经过 base64 编码的s签名字符串进行 rsa 验签
dongle.Verify.FromBase64String(sign.ToBase64String(), "hello world").ByRsa(pkcs1PublicKey, dongle.MD5).ToBool() // true

// 对字节切片进行 rsa 签名
sign := dongle.Sign.FromBytes([]byte("hello world")).ByRsa(pkcs8PrivateKey, dongle.SHA512)
// 对未经编码的原始签名字节切片进行 rsa 验签
dongle.Verify.FromRawBytes(sign.ToRawBytes(), []byte("hello world")).ByRsa(pkcs1PublicKey, dongle.SHA512).ToBool() // true
// 对经过 hex 编码的签名字节切片进行 rsa 验签
dongle.Verify.FromHexBytes(sign.ToHexBytes(), []byte("hello world")).ByRsa(pkcs1PublicKey, dongle.SHA512).ToBool() // true
// 对经过 base64 编码的签名字节切片进行 rsa 验签
dongle.Verify.FromBase64Bytes(sign.ToBase64Bytes(), []byte("hello world")).ByRsa(pkcs1PublicKey, dongle.SHA512).ToBool() // true
```

#### OpenSSL

##### RSA 密钥

```go
// 生成 PKCS#1 格式 RSA 密钥对
publicKey, privateKey := openssl.RSA.GenKeyPair(opensssl.PKCS1, 1024)
// 生成 PKCS#8 格式 RSA 密钥对
publicKey, privateKey := openssl.RSA.GenKeyPair(opensssl.PKCS8, 2048)

// 验证 RSA 密钥对是否匹配
openssl.RSA.VerifyKeyPair(publicKey, privateKey) // true
openssl.RSA.VerifyKeyPair(publicKey, []byte("xxx")) // false

// 验证是否是 RSA 公钥
openssl.RSA.IsPublicKey(publicKey) // true
openssl.RSA.IsPublicKey(privateKey) // false

// 验证是否是 RSA 私钥
openssl.RSA.IsPrivateKey(privateKey) // true
openssl.RSA.IsPrivateKey(publicKey) // false

// 解析 RSA 公钥
pub, err := openssl.RSA.ParsePublicKey(publicKey)
// 解析 RSA 私钥
pri, err := openssl.RSA.ParsePrivateKey(privateKey)

// 格式化 PKCS#1 格式 RSA 公钥，添加头尾和换行符
openssl.RSA.FormatPublicKey(openssl.PKCS1, publicKey1)
// 格式化 PKCS#8 格式RSA 公钥，添加头尾和换行符
openssl.RSA.FormatPublicKey(openssl.PKCS8, publicKey8)
// 格式化 PKCS#1 格式 RSA 私钥，添加头尾和换行符
openssl.RSA.FormatPrivateKey(openssl.PKCS1, privateKey1)
// 格式化 PKCS#8 格式 RSA 私钥，添加头尾和换行符
openssl.RSA.FormatPrivateKey(openssl.PKCS8, privateKey8)

// 压缩 RSA 密钥，去掉头尾和换行符
keyBytes, err := openssl.RSA.CompressKey(key)

// 从 RSA 私钥里导出公钥
publicKey, err := openssl.RSA.ExportPublicKey(privateKey)

```

### 错误处理

> 如果有多个错误发生，只返回第一个错误，前一个错误排除后才返回下一个错误

```go
e := dongle.Encrypt.FromString("hello world").ByRsa([]byte("xxxx"))
if e.Error != nil {
    // 错误处理...
    log.Fatal(e.Error)
}
fmt.Println(e.ToString())
// 输出
rsa: invalid public key, please make sure the public key is valid
```

### 参考项目

* [openssl/openssl](https://github.com/openssl/openssl)
* [javascript/crypto-js](https://github.com/brix/crypto-js)
* [nodejs/crypto](https://nodejs.org/api/crypto.html)
* [java/jasypt](https://github.com/jasypt/jasypt)

### 赞助

`Dongle` 是一个非商业开源项目, 如果你想支持 `Dongle`, 你可以为开发者 [购买一杯咖啡](https://www.gouguoyin.com/zanzhu.html)

### 致谢

`Dongle`已获取免费的 JetBrains 开源许可证，在此表示感谢

<a href="https://www.jetbrains.com"><img src="https://github-oss.oss-cn-beijing.aliyuncs.com/jetbrains.png" height="100" alt="JetBrains"/></a>