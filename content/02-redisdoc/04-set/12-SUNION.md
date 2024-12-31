---
author: 黄健宏
title: SUNION
date: 2024-10-28 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2412
tags:
  - Redis
  - 集合
  - SUNION
---
# SUNION key [key …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N)， `N` 是所有给定集合的成员数量之和。

返回一个集合的全部成员，该集合是所有给定集合的并集。

不存在的 `key` 被视为空集。

## 返回值

并集成员的列表。

## 代码示例

```shell
redis> SMEMBERS songs
1) "Billie Jean"

redis> SMEMBERS my_songs
1) "Believe Me"

redis> SUNION songs my_songs
1) "Billie Jean"
2) "Believe Me"
```