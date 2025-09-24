---
author: 黄健宏
title: HMGET
date: 2024-03-07 15:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20211
tags:
  - Redis
  - 哈希表
  - HMGET
---

# HMGET

**HMGET key field [field …]**

返回哈希表 `key` 中，一个或多个给定域的值。

如果给定的域不存在于哈希表，那么返回一个 `nil` 值。

因为不存在的 `key` 被当作一个空哈希表来处理，所以对一个不存在的 `key` 进行 `HMGET` 操作将返回一个只带有 `nil` 值的表。

**可用版本：**

>= 2.0.0

**时间复杂度：**

O(N)， `N` 为给定域的数量。

**返回值：**

一个包含多个给定域的关联值的表，表值的排列顺序和给定域参数的请求顺序一样。

```shell
redis> HMSET pet dog "doudou" cat "nounou"    # 一次设置多个域
OK

redis> HMGET pet dog cat fake_pet             # 返回值的顺序和传入参数的顺序一样
1) "doudou"
2) "nounou"
3) (nil)                                      # 不存在的域返回nil值
```