---
author: 黄健宏
title: BRPOPLPUSH
date: 2024-03-07 15:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 20316
tags:
  - Redis
  - 列表
  - BRPOPLPUSH
---

# BRPOPLPUSH source destination timeout

> 可用版本： >= 2.2.0
> 
> 时间复杂度： O(1)

[BRPOPLPUSH](../../02-redisdoc/03-list/16-brpoplpush/) 是 [RPOPLPUSH source destination](../../02-redisdoc/03-list/07-rpoplpush/) 的阻塞版本，当给定列表 `source` 不为空时， [BRPOPLPUSH](../../02-redisdoc/03-list/16-brpoplpush/) 的表现和 [RPOPLPUSH source destination](../../02-redisdoc/03-list/07-rpoplpush/) 一样。

当列表 `source` 为空时， [BRPOPLPUSH](../../02-redisdoc/03-list/16-brpoplpush/) 命令将阻塞连接，直到等待超时，或有另一个客户端对 `source` 执行 [LPUSH key value [value …]](../../02-redisdoc/03-list/01-lpush/) 或 [RPUSH key value [value …]](../../02-redisdoc/03-list/03-rpush/) 命令为止。

超时参数 `timeout` 接受一个以秒为单位的数字作为值。超时参数设为 `0` 表示阻塞时间可以无限期延长(block indefinitely) 。

更多相关信息，请参考 [RPOPLPUSH source destination](../../02-redisdoc/03-list/07-rpoplpush/) 命令。

## 返回值

假如在指定时间内没有任何元素被弹出，则返回一个 `nil` 和等待时长。 反之，返回一个含有两个元素的列表，第一个元素是被弹出元素的值，第二个元素是等待时长。

## 代码示例[

```shell
# 非空列表

redis> BRPOPLPUSH msg reciver 500
"hello moto"                        # 弹出元素的值
(3.38s)                             # 等待时长

redis> LLEN reciver
(integer) 1

redis> LRANGE reciver 0 0
1) "hello moto"

# 空列表

redis> BRPOPLPUSH msg reciver 1
(nil)
(1.34s)
```

## 模式：安全队列

参考 [RPOPLPUSH source destination](../../02-redisdoc/03-list/07-rpoplpush/) 命令的《安全队列》一节。

## 模式：循环列表

参考 [RPOPLPUSH source destination](../../02-redisdoc/03-list/16-brpoplpush/) 命令的《循环列表》一节。