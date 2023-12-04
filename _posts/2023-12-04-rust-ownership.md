---
layout: post
title: Rust Ownership
subtitle: Ownership
tags: [ rust ]
comments: true
---

### Rust使用权

**类型是否实现copy特征的对应表**

| 类型                                          | Copy trait |
|---------------------------------------------|------------|
| i8, i16, i32, i64, u8, u16, u32, u64, usize | 是          |
| f32，f32                                     | 是          |
| char                                        | 是          |
| bool                                        | 是          |
| Tuple                                       | 是          |
| [T; N] T impl Copy                          | 是          |
| [T; N] T not impl Copy                      | 否          |
| fn                                          | 是          |
| String                                      | 否          |
| Vec T impl Copy                             | 是          |
| Vec T not impl Copy                         | 否          |
| Box<T>                                      | 否          |
| Rc<T>                                       | 否          |
| Rc<T>                                       | 否          |

**实现了copy trait类型，变量转移给另一个变量，不会发生所有权转移**

```
let a:usize = 1024;

let b:usize = a;    //usize具有copy特征，变量a不转移所有权 接下来还可以继续使用

assert_eq!(a, b);
```

**未实现copy trait类型，变量转移给另一个变量时会发生所有权转移**

```
let a:String = String::from("1024");

let b:String = a;    //String不具有copy特征，a转移所有权给b, a moved

assert_eq!(a, b);   //move occurs because `a` has type `String`, which does not implement the `Copy` trait

//How to solve: use clone, let b:String = a.clone()
```

**引用**

```
let a: String = String::from("1024");

let b: &String = &a;    //引用（References）允许你借用变量的值而不获取其所有权。因此，b 不会转移 a 的所有权，而是创建了一个指向 a 的不可变引用

assert_eq!(a, *b);
```