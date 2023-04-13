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

**`strconv.ParseComplex`把字符串复数解析成complex128或complex64**

```
parseComplex, err := strconv.ParseComplex("3.14-2i", 64)
if err != nil {
    fmt.Println(err)
}
fmt.Println(parseComplex)
```
bitSize 64: 转换成complex64

bitSize 128: 转换成complex128

---
