---
layout: post
title: vue3 Component inside <Transition> renders
subtitle: vue3 Component inside <Transition> renders non-element root node that cannot be animated
tags: [ vue ]
comments: true
---

### Vue3离开当前页面，在点回来出现白屏问题，控制台提示`Component inside <Transition> renders non-element root node that cannot be animated`
#### 解决办法：把组件放在div里面
```
<template>
<div>
    <组件/>
</div>
</template>
```