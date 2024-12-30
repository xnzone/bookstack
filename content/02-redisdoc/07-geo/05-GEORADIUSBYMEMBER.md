---
author: 黄健宏
title: GEORADIUSBYMEMBER
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 705
tags:
  - Redis
  - 地理位置
  - GEORADIUSBYMEMBER
---

# GEORADIUSBYMEMBER key member radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [ASC|DESC] [COUNT count]

> 可用版本： >= 3.2.0
> 
> 时间复杂度： O(log(N)+M)， 其中 N 为指定范围之内的元素数量， 而 M 则是被返回的元素数量。

这个命令和 `GEORADIUS` 命令一样， 都可以找出位于指定范围内的元素， 但是 `GEORADIUSBYMEMBER` 的中心点是由给定的位置元素决定的， 而不是像 `GEORADIUS` 那样， 使用输入的经度和纬度来决定中心点。

关于 `GEORADIUSBYMEMBER` 命令的更多信息， 请参考 `GEORADIUS` 命令的文档。

## 返回值

一个数组， 数组中的每个项表示一个范围之内的位置元素。

## 代码示例

```bash
redis> GEOADD Sicily 13.583333 37.316667 "Agrigento"
(integer) 1

redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2

redis> GEORADIUSBYMEMBER Sicily Agrigento 100 km
1) "Agrigento"
2) "Palermo"
```
