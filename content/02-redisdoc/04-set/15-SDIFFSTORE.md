---
author: 黄健宏
title: SDIFFSTORE 
date: 2024-10-28 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20415
tags:
  - Redis
  - 集合
  - SDIFFSTORE 
---
# SDIFFSTORE destination key [key …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N)， `N` 是所有给定集合的成员数量之和。

这个命令的作用和 [SDIFF key [key …]](../14-sdiff) 类似，但它将结果保存到 `destination` 集合，而不是简单地返回结果集。

如果 `destination` 集合已经存在，则将其覆盖。

`destination` 可以是 `key` 本身。

## 返回值

结果集中的元素数量。

## 代码示例

```shell
redis> SMEMBERS joe's_movies
1) "hi, lady"
2) "Fast Five"
3) "2012"

redis> SMEMBERS peter's_movies
1) "bet man"
2) "start war"
3) "2012"

redis> SDIFFSTORE joe_diff_peter joe's_movies peter's_movies
(integer) 2

redis> SMEMBERS joe_diff_peter
1) "hi, lady"
2) "Fast Five"
```
