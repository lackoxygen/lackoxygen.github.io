---
layout: post
title: php-ext创建一个类
subtitle: zend class
tags: [php]
comments: true
---

### ⭐如何在PHP扩展里面定义自己的类

**使用php扩展实现一个getter setter 类**
```
<?php

class Conn
{
    protected string $host;
    protected string $pass;
    protected int $port;

    public function __construct(string $host, string $pass, int $port)
    {
        $this->host = $host;
        $this->pass = $pass;
        $this->port = $port;
    }

    /**
     * @param string $host
     */
    public function setHost(string $host): void
    {
        $this->host = $host;
    }

    /**
     * @param string $pass
     */
    public function setPass(string $pass): void
    {
        $this->pass = $pass;
    }

    /**
     * @param int $port
     */
    public function setPort(int $port): void
    {
        $this->port = $port;
    }

    /**
     * @return string
     */
    public function getHost(): string
    {
        return $this->host;
    }

    /**
     * @return string
     */
    public function getPass(): string
    {
        return $this->pass;
    }

    /**
     * @return int
     */
    public function getPort(): int
    {
        return $this->port;
    }
}
```

**注册扩展模块入口**
```
zend_module_entry example_module_entry = {
        STANDARD_MODULE_HEADER,
        "example",                    /* Extension name */
        example_functions,            /* zend_function_entry */
        PHP_MINIT(example),           /* PHP_MINIT - Module initialization (priority 1)*/
        PHP_MSHUTDOWN(example),       /* PHP_MSHUTDOWN - Module shutdown (priority 5)*/
        PHP_RINIT(example),           /* PHP_RINIT - Request initialization (priority 2)*/
        PHP_RSHUTDOWN(example),       /* PHP_RSHUTDOWN - Request shutdown (priority 4)*/
        PHP_MINFO(example),           /* PHP_MINFO - Module info (priority 3)*/
        PHP_EXAMPLE_VERSION,          /* Version */
        STANDARD_MODULE_PROPERTIES
};
```

**全局变量定义**
```
zend_class_entry *zend_conn_ce;
```

**类方法定义**
```
static const zend_function_entry conn_functions[] = {
        ZEND_ME(conn, __construct, arginfo_conn___construct, ZEND_ACC_PUBLIC)
        ZEND_ME(conn, setHost, arginfo_conn_setHost, ZEND_ACC_PUBLIC)
        ZEND_ME(conn, setPass, arginfo_conn_setPass, ZEND_ACC_PUBLIC)
        ZEND_ME(conn, setPort, arginfo_conn_setPort, ZEND_ACC_PUBLIC)
        ZEND_ME(conn, getHost, NULL, ZEND_ACC_PUBLIC)
        ZEND_ME(conn, getPass, NULL, ZEND_ACC_PUBLIC)
        ZEND_ME(conn, getPort, NULL, ZEND_ACC_PUBLIC)
        ZEND_FE_END
};
```

**方法参数注册**
```
ZEND_BEGIN_ARG_INFO_EX(arginfo_conn___construct, 0, 0, 0)
                ZEND_ARG_INFO(0, host)
                ZEND_ARG_INFO(0, pass)
                ZEND_ARG_INFO(0, port)
ZEND_END_ARG_INFO()

ZEND_BEGIN_ARG_INFO_EX(arginfo_conn_setHost, 0, 0, 0)
                ZEND_ARG_INFO(0, host)
ZEND_END_ARG_INFO();

ZEND_BEGIN_ARG_INFO_EX(arginfo_conn_setPass, 0, 0, 0)
                ZEND_ARG_INFO(0, pass)
ZEND_END_ARG_INFO();

ZEND_BEGIN_ARG_INFO_EX(arginfo_conn_setPort, 0, 0, 0)
                ZEND_ARG_INFO(0, port)
ZEND_END_ARG_INFO();
```
{: .box-note}
**getter无参数，无需定义，传入NULL占位即可**


**方法实现**
```
ZEND_METHOD(conn, __construct)
{
    zend_string *host = NULL;
    zend_string *pass = NULL;
    zend_long port = 0;

    zval *object, tmp;

    object = ZEND_THIS;

    int argc = ZEND_NUM_ARGS();  //获取参数个数

    //解析失败抛出异常
    if (zend_parse_parameters_ex(ZEND_PARSE_PARAMS_QUIET, argc, "SSl", &host, &pass, &port) == FAILURE) {
        zend_class_entry *ce;

        if (Z_TYPE(EX(This)) == IS_OBJECT) {
            ce = Z_OBJCE(EX(This));
        } else if (Z_CE(EX(This))) {
            ce = Z_CE(EX(This));
        } else {
            ce = object;
        }
        zend_throw_error(NULL, "Wrong parameters for %s([string $host [, string $pass [, int $port]]])", ZSTR_VAL(ce->name));
        return;
    }

    if (host){
        ZVAL_STR(&tmp, host);
        zend_update_property_ex(zend_conn_ce, object, ZSTR_KNOWN(ZEND_STR_HOST), &tmp);
    }

    if (pass){
        ZVAL_STR(&tmp, pass);
        zend_update_property_ex(zend_conn_ce, object, ZSTR_KNOWN(ZEND_STR_PASS), &tmp);
    }

    if (port){
        ZVAL_LONG(&tmp, port);
        zend_update_property_ex(zend_conn_ce, object, ZSTR_KNOWN(ZEND_STR_PORT), &tmp);
    }
}


ZEND_METHOD(conn, setHost)
{
    zend_string *host;

    zval tmp;

    if (zend_parse_parameters(ZEND_NUM_ARGS(), "S", &host) == FAILURE) {
        RETURN_FALSE;
    }

    ZVAL_STR(&tmp, host);

    zend_update_property_ex(zend_conn_ce, getThis(), ZSTR_KNOWN(ZEND_STR_HOST), &tmp);

    RETURN_TRUE;
}

ZEND_METHOD(conn, setPass)
{
    zend_string *pass;

    zval tmp;

    if (zend_parse_parameters(ZEND_NUM_ARGS(), "S", &pass) == FAILURE) {
        RETURN_FALSE;
    }

    ZVAL_STR(&tmp, pass);

    zend_update_property_ex(zend_conn_ce, getThis(), ZSTR_KNOWN(ZEND_STR_HOST), &tmp);

    RETURN_TRUE;
}

ZEND_METHOD(conn, setPort)
{
    zend_long port;

    zval tmp;

    if (zend_parse_parameters(ZEND_NUM_ARGS(), "l", &port) == FAILURE) {
        RETURN_FALSE;
    }

    ZVAL_LONG(&tmp, port);

    zend_update_property_ex(zend_conn_ce, getThis(), ZSTR_KNOWN(ZEND_STR_PORT), &tmp);

    RETURN_TRUE;
}

ZEND_METHOD(conn, getHost)
{
    zval *prop, rv;

    if (zend_parse_parameters_none() == FAILURE) {
		return;
	}

    prop = zend_read_property_ex(zend_conn_ce, getThis(), ZSTR_KNOWN(ZEND_STR_HOST), 0, &rv);

    ZVAL_DEREF(prop);

    ZVAL_COPY(return_value, prop);
}

ZEND_METHOD(conn, getPass)
{
    zval *prop, rv;

    if (zend_parse_parameters_none() == FAILURE) {
        return;
    }

    prop = zend_read_property_ex(zend_conn_ce, getThis(), ZSTR_KNOWN(ZEND_STR_PASS), 0, &rv);

    ZVAL_DEREF(prop);

    ZVAL_COPY(return_value, prop);
}

ZEND_METHOD(conn, getPort)
{
    zval *prop, rv;

    if (zend_parse_parameters_none() == FAILURE) {
        return;
    }

    prop = zend_read_property_ex(zend_conn_ce, getThis(), ZSTR_KNOWN(ZEND_STR_PORT), 0, &rv);

    ZVAL_DEREF(prop);

    ZVAL_COPY(return_value, prop);
}
```

{: .box-note}
**zend_parse_parameters_ex(flags, num_args, format, ...)用来解析方法参数**
- flags: 
- - ZEND_PARSE_PARAMS_QUIET 
- - ZEND_PARSE_PARAMS_THROW

- num_args: 参数个数
- format：参数类型

- -  b: Boolean
- -  d: Double
- -  f: Float
- -  h: Resource handle
- -  l: Integer
- -  L: Long integer
- -  O: Object
- -  s: String
- -  a: Array
- -  A: Associative array
- -  z: Callable function


**在PHP_MINIT方法内注册对象**
```
PHP_MINIT_FUNCTION (example) {
    zend_class_entry conn_ce;
    INIT_CLASS_ENTRY(conn_ce, "Lackoxygen\\Conn", conn_functions);  //初始化结构 如不需要命名空间只需填写类名
    zend_conn_ce = zend_register_internal_class_ex(&conn_ce, NULL);//注册类
    zend_declare_property_string(zend_conn_ce, "host", sizeof("host")-1, "", ZEND_ACC_PROTECTED);//注册属性protected string host
    zend_declare_property_string(zend_conn_ce, "pass", sizeof("pass")-1, "", ZEND_ACC_PROTECTED);//注册属性protected string pass
    zend_declare_property_long(zend_conn_ce, "port", sizeof("port")-1, NULL, ZEND_ACC_PROTECTED);//注册属性protected int port
    return SUCCESS;
}
```

**安装并测试**

```
phpize
./configure
make && make install
echo extension=example >> php.ini

$conn = new \Lackoxygen\Conn('127.0.0.1', '123456', 6379);

$conn->setHost('192.168.0.1');
$conn->setPass('admin');
$conn->setPort(3306);

echo $conn->getHost() . "\n";
echo $conn->getPass() . "\n";
echo $conn->getPort() . "\n";
```


