---
author: 黄健宏
title: HGET
date: 2024-03-11 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 20203
tags:
  - Redis
  - 哈希表
  - HGET
---

# HGET hash field

> 可用版本： >= 2.0.0
> 
> 时间复杂度： O(1)

返回哈希表中给定域的值。

## 返回值

`HGET` 命令在默认情况下返回给定域的值。

如果给定域不存在于哈希表中， 又或者给定的哈希表并不存在， 那么命令返回 `nil` 。

## 代码示例

域存在的情况：

```shell
redis> HSET homepage redis redis.com
(integer) 1

redis> HGET homepage redis
"redis.com"
```

域不存在的情况：

```shell
redis> HGET site mysql
(nil)
```