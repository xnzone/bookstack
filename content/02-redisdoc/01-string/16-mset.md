---
author: 黄健宏
title: MSET
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20116
tags: ["Redis", "字符串", "MSET"]
---

# MSET key value [key value …]

> 可用版本： >= 1.0.1
> 
> 时间复杂度： O(N)，其中 N 为被设置的键数量。

同时为多个键设置值。

如果某个给定键已经存在， 那么 `MSET` 将使用新值去覆盖旧值， 如果这不是你所希望的效果， 请考虑使用 `MSETNX` 命令， 这个命令只会在所有给定键都不存在的情况下进行设置。

`MSET` 是一个原子性(atomic)操作， 所有给定键都会在同一时间内被设置， 不会出现某些键被设置了但是另一些键没有被设置的情况。

## 返回值

`MSET` 命令总是返回 `OK` 。

## 代码示例

同时对多个键进行设置：

```shell
redis> MSET date "2012.3.30" time "11:00 a.m." weather "sunny"
OK

redis> MGET date time weather
1) "2012.3.30"
2) "11:00 a.m."
3) "sunny"
```

覆盖已有的值：

```shell
redis> MGET k1 k2
1) "hello"
2) "world"

redis> MSET k1 "good" k2 "bye"
OK

redis> MGET k1 k2
1) "good"
2) "bye"
```