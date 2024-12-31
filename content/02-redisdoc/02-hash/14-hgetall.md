---
author: 黄健宏
title: HGETALL
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2214
tags:
  - Redis
  - 哈希表
  - HGETALL
---

# HGETALL

**HGETALL key**

返回哈希表 `key` 中，所有的域和值。

在返回值里，紧跟每个域名(field name)之后是域的值(value)，所以返回值的长度是哈希表大小的两倍。

**可用版本：**

>= 2.0.0

**时间复杂度：**

O(N)， `N` 为哈希表的大小。

**返回值：**

以列表形式返回哈希表的域和域的值。

若 `key` 不存在，返回空列表。

```shell
redis> HSET people jack "Jack Sparrow"
(integer) 1

redis> HSET people gump "Forrest Gump"
(integer) 1

redis> HGETALL people
1) "jack"          # 域
2) "Jack Sparrow"  # 值
3) "gump"
4) "Forrest Gump"
```