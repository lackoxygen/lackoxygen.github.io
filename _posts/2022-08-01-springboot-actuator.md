---
layout: post
title: springBoot actuator
subtitle: 项目添加actuator
tags: [java]
comments: true
---

# ⭐ Actuator监控

## 添加依赖

```xml

<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-actuator</artifactId>
</dependency>
```

## 配置management

```yaml
management:
  endpoint:
    web:
      exposure:
        include: '*' #health,info,httptrace
```

## 配置WebMVC

```java
package com.lackoxygen.spring.startup;

import org.springframework.boot.SpringBootConfiguration;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurationSupport;

@SpringBootConfiguration
@EnableWebMvc
public class MVC extends WebMvcConfigurationSupport {

}
```

## 访问

```shell
curl http://127.0.0.1:8080/actuator/health

{"status":"UP"}
```
