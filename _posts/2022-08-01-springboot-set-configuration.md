---
layout: post
title: springBoot预配置
subtitle: springBoot不同方式修改配置
tags: [java]
comments: true
---

# ⭐ YAML文件配置

## 创建配置文件

**在src/main/resources目录下创建application.yml**

```yaml
server:
  servlet:
    context-path: /api  #路由前缀
  port: 8080  #访问端口

spring:
  profiles:
    active: develop  #运行环境
```

## 自定义运行环境

**在application.yml相同目录创建application-develop.yml和application.prod.yml**

## 切换环境

**设置application.yml**

```yaml
spring:
  profiles:
    active: develop  #运行环境
...
```

## 自定义配置

```yaml
define:
  host: 127.0.0.1
  port: 8080
  options:
    user-agent: foo
```

------

# 💡 动态配置

**@SpringBootConfiguration注解的类会被认为是配置类**

## 改变启动端口号

```java
package com.lackoxygen.spring.startup;

import com.lackoxygen.spring.bean.User;
import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.boot.web.servlet.server.ConfigurableServletWebServerFactory;
import org.springframework.boot.SpringBootConfiguration;
import org.springframework.context.annotation.Bean;

@SpringBootConfiguration
public class Web implements WebServerFactoryCustomizer<ConfigurableServletWebServerFactory> {
    @Override
    public void customize(ConfigurableServletWebServerFactory factory) {
        factory.setPort(9501);
    }
}
```

## 配置bean

**创建Foo.java**

```java
package com.lackoxygen.spring.some;

public class Foo {
    private String keyword;

    public void setKeyword(String keyword) {
        this.keyword = keyword;
    }

    public String getKeyword() {
        return keyword;
    }

    @Override
    public String toString() {
        return "Foo{" +
                "keyword='" + keyword + '\'' +
                '}';
    }
}
```

**bean注入**

```java
package com.lackoxygen.spring.startup;

import com.lackoxygen.spring.some.Foo;
import org.springframework.boot.web.server.WebServerFactoryCustomizer;
import org.springframework.boot.web.servlet.server.ConfigurableServletWebServerFactory;
import org.springframework.boot.SpringBootConfiguration;
import org.springframework.context.annotation.Bean;

@SpringBootConfiguration
public class Web implements WebServerFactoryCustomizer<ConfigurableServletWebServerFactory> {
    @Bean
    public Foo foo() {
        return new Foo();
    }
}
```

## 替换JSON转换器

**添加依赖**

```xml

<dependency>
    <groupId>com.alibaba</groupId>
    <artifactId>fastjson</artifactId>
    <version>2.0.12</version>
</dependency>
```

**创建MVC.java**

```java
package com.lackoxygen.spring.startup;

import com.alibaba.fastjson.serializer.SerializerFeature;
import com.alibaba.fastjson.support.config.FastJsonConfig;
import com.alibaba.fastjson.support.spring.FastJsonHttpMessageConverter;
import org.springframework.boot.SpringBootConfiguration;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageConverter;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurationSupport;

import java.util.ArrayList;
import java.util.List;

@SpringBootConfiguration
public class MVC extends WebMvcConfigurationSupport {
    @Override
    protected void configureMessageConverters(List<HttpMessageConverter<?>> converters) {
        super.configureMessageConverters(converters);

        FastJsonHttpMessageConverter fastJsonHttpMessageConverter = new FastJsonHttpMessageConverter();

        FastJsonConfig fastJsonConfig = new FastJsonConfig();

        fastJsonConfig.setSerializerFeatures(
                SerializerFeature.PrettyFormat
        );

        List<MediaType> mediaTypes = new ArrayList<>();

        mediaTypes.add(MediaType.APPLICATION_JSON_UTF8);

        fastJsonHttpMessageConverter.setSupportedMediaTypes(mediaTypes);

        fastJsonHttpMessageConverter.setFastJsonConfig(fastJsonConfig);

        converters.add(fastJsonHttpMessageConverter);
    }
}
```