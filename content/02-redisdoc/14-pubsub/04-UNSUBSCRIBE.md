---
author: 黄健宏
title: UNSUBSCRIBE
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21404
tags:
  - Redis
  - 发布与订阅
  - UNSUBSCRIBE
---

# UNSUBSCRIBE [channel [channel …]]

> 可用版本： >= 2.0.0
> 
> 时间复杂度： O(N) ， `N` 是客户端已订阅的频道的数量。

指示客户端退订给定的频道。

如果没有频道被指定，也即是，一个无参数的 `UNSUBSCRIBE` 调用被执行，那么客户端使用 `SUBSCRIBE` 命令订阅的所有频道都会被退订。在这种情况下，命令会返回一个信息，告知客户端所有被退订的频道。

## 返回值

这个命令在不同的客户端中有不同的表现。