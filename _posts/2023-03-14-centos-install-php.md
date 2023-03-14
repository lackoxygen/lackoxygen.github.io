---
layout: post
title: CentOS环境安装PHP
subtitle: CentOS环境安装PHP
tags: [php]
comments: true
---

**下载PHP源码文件**

```
wget https://www.php.net/distributions/php-7.4.33.tar.gz
```

**安装依赖**

```
yum -y install gcc gcc-c++ glibc \
libjpeg libjpeg-devel libpng libpng-devel \
freetype freetype-devel zlibzlib-devel glibc glibc-devel \
glib2 glib2-devel libxml2-devel curl curl-devel openssl openssl-devel
```

**添加运行用户**

```
groupadd www-data
 
useradd -g www -M -s /sbin/nologin www
```

**编译前配置**

- 不同版本参数有差异, 使用configure -h 查看配置选项

```
 ./configure --prefix=/usr/local/php/7.4 --enable-fpm --exec-prefix=/usr/local/php/7.4 --with-config-file-scan-dir=/usr/local/php/7.4/etc/conf.d \
 --with-config-file-path=/usr/local/php/7.4/etc --with-curl --enable-mysqlnd --with-mysqli=mysqlnd --with-pdo-mysql=mysqlnd \
 --with-gettext --with-iconv-dir --with-libdir=lib64 \
 --with-mysqli --with-fpm-user=www-data --with-fpm-group=www-data \
 --with-pdo-mysql --with-pdo-sqlite --with-pear --enable-ftp \
 --with-zlib --enable-bcmath --enable-inline-optimization \
 --enable-mbregex --enable-mbstring --enable-opcache --enable-pcntl --enable-sockets \
 --enable-sysvsem --enable-xml
```

**编译并安装**

```
make && make install
```

**复制配置**
```
# 拷贝配置文件
cp php.ini-development /usr/local/php/7.4/etc/php.ini

cp /usr/local/php/7.4/etc/php-fpm.conf.default /usr/local/php/7.4/etc/php-fpm.conf

cp /usr/local/php/7.4/etc/php-fpm.d/www.conf.default /usr/local/php/7.4/etc/php-fpm.d/www.conf

cp sapi/fpm/init.d.php-fpm /etc/init.d/php-fpm-7.4 && chmod +x /etc/init.d/php-fpm-7.4

# 启动fpm
service php-fpm-7.4 start

# 软链
ln -s /usr/local/php/7.4/bin/* /usr/bin

php -v
```

**常见问题**

{: .box-warning}
> No package 'oniguruma' found

```
# 安装系统对应版本的rpm包
yum -y install https://repo.huaweicloud.com/centos/7.9.2009/cloud/x86_64/openstack-queens/Packages/o/oniguruma-6.7.0-1.el7.x86_64.rpm
yum -y install https://repo.huaweicloud.com/centos/7.9.2009/cloud/x86_64/openstack-queens/Packages/o/oniguruma-devel-6.7.0-1.el7.x86_64.rpm
```

---

{: .box-warning}
> configure: WARNING: unrecognized options

删除提示对应参数选项

---


