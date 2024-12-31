---
author: 黄健宏
title: SUNIONSTORE
date: 2024-10-28 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20413
tags:
  - Redis
  - 集合
  - SUNIONSTORE
---

# SUNIONSTORE destination key [key …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度: O(N)， `N` 是所有给定集合的成员数量之和。

这个命令类似于 [SUNION key [key …]](../12-sunion) 命令，但它将结果保存到 `destination` 集合，而不是简单地返回结果集。

如果 `destination` 已经存在，则将其覆盖。

`destination` 可以是 `key` 本身。

## 返回值

结果集中的元素数量。

## 代码示例

```shell
redis> SMEMBERS NoSQL
1) "MongoDB"
2) "Redis"

redis> SMEMBERS SQL
1) "sqlite"
2) "MySQL"

redis> SUNIONSTORE db NoSQL SQL
(integer) 4

redis> SMEMBERS db
1) "MySQL"
2) "sqlite"
3) "MongoDB"
4) "Redis"
```
