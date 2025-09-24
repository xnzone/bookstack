---
author: 黄健宏
title: SCARD
date: 2024-03-07 15:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20407
tags:
  - Redis
  - 集合
  - SCARD
---

# SCARD key

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(1)

返回集合 `key` 的基数(集合中元素的数量)。

## 返回值

集合的基数。 当 `key` 不存在时，返回 `0` 。

## 代码示例

```shell
redis> SADD tool pc printer phone
(integer) 3

redis> SCARD tool   # 非空集合
(integer) 3

redis> DEL tool
(integer) 1

redis> SCARD tool   # 空集合
(integer) 0
```