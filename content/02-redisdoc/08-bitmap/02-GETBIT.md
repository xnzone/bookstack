---
author: 黄健宏
title: GETBIT
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2802
tags:
  - Redis
  - 位图
  - GETBIT
---

# GETBIT key offset

> 可用版本： >= 2.2.0
> 
> 时间复杂度： O(1)

对 `key` 所储存的字符串值，获取指定偏移量上的位(bit)。

当 `offset` 比字符串值的长度大，或者 `key` 不存在时，返回 `0` 。

## 返回值

字符串值指定偏移量上的位(bit)。

## 代码示例

```bash
# 对不存在的 key 或者不存在的 offset 进行 GETBIT， 返回 0

redis> EXISTS bit
(integer) 0

redis> GETBIT bit 10086
(integer) 0

# 对已存在的 offset 进行 GETBIT

redis> SETBIT bit 10086 1
(integer) 0

redis> GETBIT bit 10086
(integer) 1
```
