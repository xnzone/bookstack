---
author: 黄健宏
title: LLEN
date: 2024-03-07 15:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2309
tags:
  - Redis
  - 列表
  - LLEN
---

# LLEN key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

返回列表 `key` 的长度。

如果 `key` 不存在，则 `key` 被解释为一个空列表，返回 `0` .

如果 `key` 不是列表类型，返回一个错误。

## 返回值

列表 `key` 的长度。

## 代码示例

```shell
# 空列表

redis> LLEN job
(integer) 0

# 非空列表

redis> LPUSH job "cook food"
(integer) 1

redis> LPUSH job "have lunch"
(integer) 2

redis> LLEN job
(integer) 2
```