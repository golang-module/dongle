---
name: "\U0001F41B Bug report"
about: Report a reproducible bug or regression.
labels: Bug
---
Hello,

I encountered an issue with the following code:
```go
dongle.Encrypt.FromString("hello world").ByMd5().ToString()
```

golang version: **such as 1.16**

dongle version: **such as 1.0.0**

I expected to get:

```
5eb63bbbe01eeed093cb22bb8f5acdc3
```
<!--
    Always give your expectations. Each use has their owns.
    You may want daylight saving time to be taken into account,
    someone else want it to be ignored. You may want timezone
    to be used in comparisons, someone else may not, etc.
-->

But I actually get:

```
014f03f9025ea81a8a0e9734be540c53
```

Thanks!
