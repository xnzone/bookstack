---
author: 黄健宏
title: GEOHASH
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20706
tags:
  - Redis
  - 地理位置
  - GEOHASH
---

# GEOHASH key member [member …]

> 可用版本： >= 3.2.0
> 
> 时间复杂度： 寻找每个位置元素的复杂度为 O(log(N)) ， 其中 N 为给定键包含的位置元素数量。

返回一个或多个位置元素的 [Geohash](https://en.wikipedia.org/wiki/Geohash) 表示。

## 返回值

一个数组， 数组的每个项都是一个 geohash 。 命令返回的 geohash 的位置与用户给定的位置元素的位置一一对应。

## 代码示例

```bash
redis> GEOADD Sicily 13.361389 38.115556 "Palermo" 15.087269 37.502669 "Catania"
(integer) 2

redis> GEOHASH Sicily Palermo Catania
1) "sqc8b49rny0"
2) "sqdtr74hyu0"
```
