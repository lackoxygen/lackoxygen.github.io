---
layout: post
title: webFlux 快速入门
subtitle: webFlux基本用法
tags: [java]
comments: true
---

# ⭐ webFlux快速入门

## 添加依赖

```xml

<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-webflux</artifactId>
</dependency>
```

## MVC注解使用

**创建FluxController.java**

```java
package com.lackoxygen.spring.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Mono;

@RestController
public class FluxController {
    @GetMapping(path = "flux")
    public Mono<String> Hello() {
        return Mono.just("hello mono");
    }
}
```


