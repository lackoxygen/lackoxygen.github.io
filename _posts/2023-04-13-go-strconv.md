---
layout: post
title: go strconv
subtitle: go字符串类型转换
tags: [go]
comments: true
---

**`strconv.ParseBool`字符串转换布尔值**

```
fmt.Println(strconv.ParseBool("1")) //true nil
fmt.Println(strconv.ParseBool("t")) //true nil
fmt.Println(strconv.ParseBool("T")) //true nil
fmt.Println(strconv.ParseBool("true")) //true nil
fmt.Println(strconv.ParseBool("TRUE")) //true nil 
fmt.Println(strconv.ParseBool("True")) //true nil

fmt.Println(strconv.ParseBool("0")) //false nil
fmt.Println(strconv.ParseBool("f")) //false nil
fmt.Println(strconv.ParseBool("F")) //false nil
fmt.Println(strconv.ParseBool("false")) //false nil
fmt.Println(strconv.ParseBool("FALSE")) //false nil
fmt.Println(strconv.ParseBool("False")) //false nil

fmt.Println(strconv.ParseBool("enable")) //false error
```

字符串("1" | "t" | "T" | "true" | "TRUE" | "True")返回true,无错误

字符串("0" | "f" | "F" | "false" | "FALSE" | "False")返回false,无错误

其他字符串 返回false，有错误

---

**`strconv.FormatBool`把布尔值转换成布尔值字符串**

```
fmt.Printf("`true`  %s\n", strconv.FormatBool(true))
fmt.Printf("`false`  %s\n", strconv.FormatBool(false))
```

---

**`strconv.AppendBool`把布尔值转成字符放入字节切片**

```
var myByte []byte

myByte = strconv.AppendBool(myByte, true)

myByte = strconv.AppendBool(myByte, false)

fmt.Println(myByte)
```
---

**`strconv.ParseComplex`字符串复数解析成complex128或complex64**

```
parseComplex, err := strconv.ParseComplex("3.14-2i", 64)
if err != nil {
    fmt.Println(err)
}
fmt.Println(parseComplex)
```

`bitSize`可用范围：(64 -> 转换成complex64), (128 -> 转换成complex128)

---

**`strconv.ParseUint`字符串转uint64**


```
var err error

var ui uint64

bitSize := 8

s := strconv.Itoa(int(uint(1) << bitSize))

ui, err = strconv.ParseUint(s, 10, bitSize)

fmt.Println(err, ui)
```

1 << 8 超过uint8的最大值，解析时溢出抛出错误:strconv.ParseUint: parsing "256": value out of range 255

`base`: 字符串的进制，填写0会自动推断，默认是10进制

`bitSize`: bit位数 (8 -> uint8) (16 -> uint16) (32 -> uint32) (64 -> uint64)

---

**`strconv.ParseInt`字符串转int64**

```
parseInt, err := strconv.ParseInt("010101", 2, 64)
if err != nil {
    fmt.Println(err)
}
fmt.Println(parseInt)
```

用法跟`strconv.ParseUint`一致

---

**`strconv.Atoi`字符串转int**

```
parseInt, err := strconv.Atoi("11111")
if err != nil {
    fmt.Println(err)
}
fmt.Println(parseInt)
```

十进制字符串转换int格式

---

**`strconv.FormatComplex`complex 类型的值转换为字符串**

```
var cx complex128

cx = 3.14 - 2i

fmt.Println(strconv.FormatComplex(cx, 'e', 2, 128))
```

**`strconv.FormatFloat`格式化浮点数**

fmt:
'b' (-ddddp±ddd, a binary exponent),
'e' (-d.dddde±dd, a decimal exponent),
'E' (-d.ddddE±dd, a decimal exponent),
'f' (-ddd.dddd, no exponent),
'g' ('e' for large exponents, 'f' otherwise),
'G' ('E' for large exponents, 'f' otherwise),
'x' (-0xd.ddddp±ddd, a hexadecimal fraction and binary exponent), or
'X' (-0Xd.ddddP±ddd, a hexadecimal fraction and binary exponent).

prec: 保留的精度

bitSize: float32  | float64


```
f := 3.1415926

fmt.Println(strconv.FormatFloat(f, 'f', 2, 64))
```
保留2位小数，小数会4舍五入

---

**`strconv.AppendFloat`把浮点数放入切片**

fmt: 跟FormatFloat一样
prec: 保留的精度
bitSize: float32  | float64

```
f := 3.1415926

var myByte []byte

fmt.Println(strconv.AppendFloat(myByte, f, 'f', 2, 32))
```

**`strconv.Quote`将一个字符串转换为带双引号的 Go 语言字符串字面量表示**

```
fmt.Println(strconv.Quote("Hello, 世界\\n"))
```

**`strconv.AppendQuote`字符串转换为带双引号的字符串追加到切片**

```
var myByte []byte

fmt.Println(strconv.AppendQuote(myByte, "println\n"))
```
---

**`strconv.AppendQuoteToASCII`非ASCII 字符和特殊字符都已转换为转义序列**

```
s := "in 中国"

fmt.Println(strconv.QuoteToASCII(s))
```

'中国'非ASCII 字符 会转换\u4e2d\u56fd形式表示

---

**`strconv.AppendQuoteToASCII`非ASCII 字符和特殊字符都已转换为转义序列后放入切片**

```
s := "in 中国"

var myByte []byte

fmt.Println(strconv.AppendQuoteToASCII(myByte, s))
```
---

**`strconv.QuoteToGraphic`将一个字符串转换为带双引号的 Go 语言字符串字面量表示**

```
s := "in 中国\n"

fmt.Println(s)  //in 中国

fmt.Println(strconv.QuoteToGraphic(s)) //in 中国 \n
```
---

**`strconv.AppendQuoteToGraphic`把带双引号的字符串追加到切片**

```
s := "in 中国\n"

var myByte []byte

myByte = strconv.AppendQuoteToGraphic(myByte, s)

fmt.Println(myByte)
```

""也会写入到[]byte

---

**`strconv.QuoteRune`返回一个带单引号的表示输入字符**

```
fmt.Println(strconv.QuoteRune('我'))
```

---

**`strconv.AppendQuoteRune` Rune类型的值转换为带单引号放入bytes**

```
r := '\n'

var myByte []byte

myByte = strconv.AppendQuoteRune(myByte, r)

fmt.Println(myByte)

```

---

**`strconv.QuoteRuneToASCII`**






