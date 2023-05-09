---
layout: post
title: python识别图片里面的文字
subtitle: ocr图片识别
tags: [python]
comments: true
---

### 使用`pytesseract`库识别图片里面的文字

**安装软件依赖**

- mac

```
brew install tesseract

brew install tesseract-lang # 中文语言包

```

- centos

```
yum install epel-release

yum install tesseract

yum install tesseract-langpack-chi_sim  # 中文语言包
```

### 安装`pytesseract`
```
pip3 install pytesseract
```

### 安装`pillow`用来打开图片
```
pip3 install pillow
```

### 图片识别

- `lang`语言

```
image = Image.open(image_path)
text = pytesseract.image_to_string(image, lang='chi_sim')

print(text)
```


