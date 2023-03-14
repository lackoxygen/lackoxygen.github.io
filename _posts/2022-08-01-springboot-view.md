---
layout: post
title: springBoot集成模版引擎
subtitle: springBoot集成模版示例
tags: [java]
comments: true
---

# ⭐ 集成模板引擎

## 官方集成如下几种模板引擎

- FreeMarker
- Velocity
- Thymeleaf
- Groovy
- Mustache
- JSP

## 添加FreeMarker依赖

```xml

<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-freemarker</artifactId>
</dependency>
```

## 配置模板

```yaml
spring:
  freemarker: #模板配置
    template-loader-path: classpath:/templates  #模板加载路径
    suffix: .html #模板后缀
```

## 创建文件夹

- **resources/static**
- **resources/templates**

## 在templates下创建index.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>hello world</title>
</head>
<body>
...
</body>
</html>
```

## 加载模板
**创建ViewController**

```java
package com.lackoxygen.spring.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class ViewController {
    @GetMapping(value = "index.html")
    public String html() {
        return "index";
    }
}
```

**使用模板不能使用@RestController**