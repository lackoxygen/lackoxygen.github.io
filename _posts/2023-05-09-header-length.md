---
layout: post
title: data header length
subtitle: content-length
tags: [linux]
comments: true
---

### 使用字节表示数据长度

一个byte等于8bit，最大长度(2<<7) - 1, 取值范围(0-255)

多个字节的大小，2 << (8 * 字节数) - 1

### 通过最大长度计算需要的字节数

```
function bytes_num(int $count)
{
    return ceil(bcdiv(bcdiv(log($count + 1), log(2), 10), 8, 10));
}
```

**一个字节表示长度(0-255)**

- pow(2, 8) - 1

```
$length = 255

$byte = $length & 0xFF
```

**二个字节表示长度(0-65535)**

- pow(2, 8 * 2) - 1

```
$length = 65535; 

$highByte = ($length >> 8) & 0xFF;
$lowByte = $length & 0xFF;

$bytes = [$highByte, $lowByte];

```


**三个字节表示长度(0 - 16777215)**

- pow(2, 8 * 3) - 1

```
$length = 16777215; 

$highByte = ($length >> (8 * 2)) & 0xFF;
$middleByte = ($length >> (8 * 1)) & 0xFF;
$lowByte = $length & 0xFF;

$bytes = [$highByte, $middleByte, $lowByte];
```

**4个字节表示长度(0 - 4294967295)**

- pow(2, 8 * 4) - 1

```
$length = 4294967295; 

$highByte = ($length >> (8 * 3)) & 0xFF;
$middleByte1 = ($length >> (8 * 2)) & 0xFF;
$middleByte2 = ($length >> (8 * 1)) & 0xFF;
$lowByte = $length & 0xFF;

$bytes = [$highByte, $middleByte1, $middleByte2, $lowByte];
```




