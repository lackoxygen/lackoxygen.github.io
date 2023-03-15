---
layout: post
title: springBoot注解路由
subtitle: 路由用法
tags: [java]
comments: true
---


# ⭐ springBoot输出hello world

## 创建HelloController

```java
package com.lackoxygen.spring.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HelloController {
    @GetMapping(value = "hello")
    public String hello() {
        return "hello world";
    }
}
```

## 请求方式注解

- org.springframework.web.bind.annotation.*

| **注解名**            | 限制请求方法         |
|--------------------|----------------|
| **RequestMapping** | 默认不限制或通过method |
| **GetMapping**     | GET            |
| **PostMapping**    | POST           |
| **PutMapping**     | PUT            |
| **DeleteMapping**  | DELETE         |
| **PatchMapping**   | PATCH          |

## RestController

#### org.springframework.web.bind.annotation

- **相当于@Controller 和 @ResponseBody注解的结合**

## 请求方法

#### org.springframework.web.bind.annotation.RequestMethod

- **RequestMethod.GET**
- **RequestMethod.POST**
- **RequestMethod.PUT**
- **RequestMethod.DELETE**
- **RequestMethod.OPTIONS**
- **RequestMethod.HEAD**
- **RequestMethod.PATCH**
- **RequestMethod.TRACE**

## 返回

**返回String类型需要注明produces 为
application/json;charset=utf8否则可能出现乱码，返回对象无需注明，返回对象会执行spring默认的json解析器**

```java
package com.lackoxygen.spring.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.http.MediaType;

@RestController
public class HelloController {
    @GetMapping(value = "hello", produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
    public String hello() {
        return "hello world";
    }
}
```

## 启动访问http://127.0.0.1:8080/hello

- 默认端口8080