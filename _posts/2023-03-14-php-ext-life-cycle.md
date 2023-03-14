---
layout: post
title: php扩展生命周期
subtitle: php扩展运行时生命周期方法
tags: [php]
comments: true
---

#### 1.模块初始化
```
PHP_MINIT
```

#### 2.请求初始化
```
PHP_RINIT
```

#### 3.模块信息 php --ri
```
PHP_MINFO_FUNCTION
```

#### 4.关闭
```
PHP_RSHUTDOWN
```

#### 5.释放
```
PHP_MSHUTDOWN
```

#### example
```
zend_module_entry example_module_entry = {
	STANDARD_MODULE_HEADER,
	"example",					/* 扩展名称 */
	ext_functions,				/* zend_function_entry */
    PHP_MINIT(example),			/* PHP_MINIT - Module initialization */
    PHP_MSHUTDOWN(example),		/* PHP_MSHUTDOWN - Module shutdown */
    PHP_RINIT(example),			/* PHP_RINIT - Request initialization */
    PHP_RSHUTDOWN(example),		/* PHP_RSHUTDOWN - Request shutdown */
    PHP_MINFO(example),			/* PHP_MINFO - Module info */
	PHP_EXAMPLE_VERSION,		/* Version */
	STANDARD_MODULE_PROPERTIES
};