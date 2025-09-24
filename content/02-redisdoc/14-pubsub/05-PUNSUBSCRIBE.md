---
author: 黄健宏
title: PUNSUBSCRIBE
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21405
tags:
  - Redis
  - 发布与订阅
  - PUNSUBSCRIBE
---

# PUNSUBSCRIBE [pattern [pattern …]]

> 可用版本： >= 2.0.0
> 
> 时间复杂度： O(N+M) ，其中 `N` 是客户端已订阅的模式的数量， `M` 则是系统中所有客户端订阅的模式的数量。

指示客户端退订所有给定模式。

如果没有模式被指定，也即是，一个无参数的 `PUNSUBSCRIBE` 调用被执行，那么客户端使用 [PSUBSCRIBE pattern [pattern …]](psubscribe.html#psubscribe) 命令订阅的所有模式都会被退订。在这种情况下，命令会返回一个信息，告知客户端所有被退订的模式。

## 返回值

这个命令在不同的客户端中有不同的表现。