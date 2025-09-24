---
author: 黄健宏
title: ECHO
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21802
tags:
  - Redis
  - 调试
  - ECHO
---

# ECHO message

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

打印一个特定的信息 `message` ，测试时使用。

## 返回值

`message` 自身。

## 代码示例

```bash
redis> ECHO "Hello Moto"
"Hello Moto"

redis> ECHO "Goodbye Moto"
"Goodbye Moto"
```