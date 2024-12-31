---
author: 黄健宏
title: SDIFF
date: 2024-10-28 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2414
tags:
  - Redis
  - 集合
  - SDIFF
---


# SDIFF key [key …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N)， `N` 是所有给定集合的成员数量之和。

返回一个集合的全部成员，该集合是所有给定集合之间的差集。

不存在的 `key` 被视为空集。

## 返回值

一个包含差集成员的列表。

## 代码示例

```shell
redis> SMEMBERS peter's_movies
1) "bet man"
2) "start war"
3) "2012"

redis> SMEMBERS joe's_movies
1) "hi, lady"
2) "Fast Five"
3) "2012"

redis> SDIFF peter's_movies joe's_movies
1) "bet man"
2) "start war"
```
