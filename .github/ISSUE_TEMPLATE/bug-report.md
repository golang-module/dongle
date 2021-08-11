---
name: "\U0001F41B Bug report"
about: Report a reproducible bug or regression.
labels: Bug
---
Hello,

I encountered an issue with the following code:
```go
carbon.Parse("2020-08-05 13:14:15").AddDays(3).ToDateTimeString()
```

golang version: **such as 1.16**

openssl version: **such as 1.0.0**

I expected to get:

```
2020-08-08 13:14:15
```
<!--
    Always give your expectations. Each use has their owns.
    You may want daylight saving time to be taken into account,
    someone else want it to be ignored. You may want timezone
    to be used in comparisons, someone else may not, etc.
-->

But I actually get:

```
2020-08-07 13:14:15
```

Thanks!