---
author: 黄健宏
title: PSETEX
date: 2024-03-07 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 104 
tags: ["Redis", "字符串", "PSETEX"]
---
# PSETEX key milliseconds value

> 可用版本： >= 2.6.0
> 
> 时间复杂度： O(1)

这个命令和 `SETEX` 命令相似， 但它以毫秒为单位设置 `key` 的生存时间， 而不是像 `SETEX` 命令那样以秒为单位进行设置。

## 返回值

命令在设置成功时返回 `OK` 。

## 代码示例

```shell
redis> PSETEX mykey 1000 "Hello"
OK

redis> PTTL mykey
(integer) 999

redis> GET mykey
"Hello"
```