---
author: 黄健宏
title: MGET
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20118
tags: ["Redis", "字符串", "MGET"]
---

# MGET key [key …]

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(N) ，其中 N 为给定键的数量。

返回给定的一个或多个字符串键的值。

如果给定的字符串键里面， 有某个键不存在， 那么这个键的值将以特殊值 `nil` 表示。

## 返回值

`MGET` 命令将返回一个列表， 列表中包含了所有给定键的值。

## 代码示例

```shell
redis> SET redis redis.com
OK

redis> SET mongodb mongodb.org
OK

redis> MGET redis mongodb
1) "redis.com"
2) "mongodb.org"

redis> MGET redis mongodb mysql     # 不存在的 mysql 返回 nil
1) "redis.com"
2) "mongodb.org"
3) (nil)
```