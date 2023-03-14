---
layout: post
title: springBoot读取配置
subtitle: 通过不同方式读取加载配置
tags: [java]
comments: true
---

# ⭐ 读取配置

## 通过属性注解读取

**org.springframework.beans.factory.annotation.Value**

**创建ConfigController.php**

```java
package com.lackoxygen.spring.controller;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class ConfigController {
    @Value("${spring.profiles.active}")
    private String active;  //develop
}
```

## 通过env读取

**org.springframework.core.env.Environment;**

```java
package com.lackoxygen.spring.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.env.Environment;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.io.Serializable;
import java.util.HashMap;

@RestController
public class ConfigController {
    @Autowired
    protected Environment environment;

    @GetMapping(path = "kv")
    public void kv() {
        System.out.println(this.environment.getProperty("spring.profiles.active"));
    }
}
```

## 通过@ConfigurationProperties注入

**创建WebConfigure.java**

```java
package com.lackoxygen.spring.configure;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@ConfigurationProperties("server")
@Component
public class WebConfigure {
    private Integer port;

    public void setPort(Integer port) {
        this.port = port;
    }

    public Integer getPort() {
        return port;
    }
}
```