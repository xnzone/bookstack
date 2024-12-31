---
author: 黄健宏
title: SISMEMBER
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2402
tags:
  - Redis
  - 集合
  - SISMEMBER
---

# SISMEMBER key member

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(1)

判断 `member` 元素是否集合 `key` 的成员。

## 返回值

如果 `member` 元素是集合的成员，返回 `1` 。 如果 `member` 元素不是集合的成员，或 `key` 不存在，返回 `0` 。

## 代码示例

```shell
redis> SMEMBERS joe's_movies
1) "hi, lady"
2) "Fast Five"
3) "2012"

redis> SISMEMBER joe's_movies "bet man"
(integer) 0

redis> SISMEMBER joe's_movies "Fast Five"
(integer) 1
```