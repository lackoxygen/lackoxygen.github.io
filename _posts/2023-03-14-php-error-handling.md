---
layout: post
title: PHP错误处理
subtitle: PHP错误处理
tags: [php]
comments: true
---

php 异常处理
```
        $err = null;
        # 设置异常回调
        set_error_handler(function (
            int $code,
            string $msg,
            string $file,
            int $line,
            array $context
        ) use (&$err) {
            $err = $msg;
        });

        try {
            # 删除一个不存在的文件，触发异常处理
            unlink(__DIR__ . '/foo.txt');
        } finally {
            # 清除异常处理
            restore_error_handler();
        }

        if ($err !== null){
            //next
        }
```