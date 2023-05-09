---
layout: post
title: 提取doc文档的内容
subtitle: 使用不同的方法提取word文档
tags: [ python,php ]
comments: true
---
### 提取word文档里面的内容

### 通过`antiword`解析doc内容

**安装`antiword`**

- mac

```
brew install antiword
```

- centos

```
wget https://src.fedoraproject.org/lookaside/pkgs/antiword/antiword-0.37.tar.gz/f868e2a269edcbc06bf77e89a55898d1/antiword-0.37.tar.gz

tar -zxvf antiword-0.37.tar.gz

cd antiword-0.37

make && make install
```

**在cli使用**

```
antiword my_test.doc
```


**在代码使用**

```
import subprocess

result = subprocess.run(['antiword', 'my_test.doc'], capture_output=True)
output = result.stdout.decode('utf-8')
print(output)
```