---
author: 黄健宏
title: HLEN
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 406
tags:
  - Redis
  - 哈希表
  - HLEN
---

# HLEN

**HLEN key**

返回哈希表 `key` 中域的数量。

**时间复杂度：**

O(1)

**返回值：**

哈希表中域的数量。

当 `key` 不存在时，返回 `0` 。

{{< highlight shell >}}
redis> HSET db redis redis.com
(integer) 1

redis> HSET db mysql mysql.com
(integer) 1

redis> HLEN db
(integer) 2

redis> HSET db mongodb mongodb.org
(integer) 1

redis> HLEN db
(integer) 3
{{< /highlight >}}