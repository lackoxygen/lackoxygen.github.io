---
layout: post
title: 创建springBoot项目
subtitle: 使用mvn搭建springBoot项目
tags: [java]
comments: true
---

# ⭐ IDEA创建springBoot项目

## pom.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.lackoxygen.spring</groupId>
    <artifactId>spring</artifactId>
    <version>1.0-SNAPSHOT</version>

    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.5.0</version>
        #标记springboot父项目版本号2.5.0
    </parent>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
            #添加starter-web依赖 提供spring mvc支持
            <version>${project.parent.version}</version>
        </dependency>
    </dependencies>
</project>
```

## 创建启动类

**Application.java**

```java
package com.lackoxygen.spring;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
```

**@SpringBootApplication启动入口，由@EnableAutoConfiguration 和 @ComponentScan等组成**

## 启动

```
mvn spring-boot:run
```

- springBoot默认内置tomcat

## 修改依赖容器

```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.lackoxygen.spring</groupId>
    <artifactId>spring</artifactId>
    <version>1.0-SNAPSHOT</version>

    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.5.0</version>
        #标记springboot父项目版本号2.5.0
    </parent>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
            #添加starter-web依赖 提供spring mvc支持
            <version>${project.parent.version}</version>
            <exclusions>
                <exclusion>#排除tomcat容器
                    <groupId>org.springframework.boot</groupId>
                    <artifactId>spring-boot-starter-tomcat</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            #设置jetty支持
            <artifactId>spring-boot-starter-jetty</artifactId>
        </dependency>
    </dependencies>
</project>
```
