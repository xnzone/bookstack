---
author: 黄健宏
title: SETRANGE
date: 2024-03-10 17:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 2109 
tags: ["Redis", "字符串", "SETRANGE"]
---

# SETRANGE key offset value

> 可用版本： >= 2.2.0
> 
> 时间复杂度：对于长度较短的字符串，命令的平摊复杂度O(1)；对于长度较大的字符串，命令的复杂度为 O(M) ，其中 M 为 `value` 的长度。

从偏移量 `offset` 开始， 用 `value` 参数覆写(overwrite)键 `key` 储存的字符串值。

不存在的键 `key` 当作空白字符串处理。

`SETRANGE` 命令会确保字符串足够长以便将 `value` 设置到指定的偏移量上， 如果键 `key` 原来储存的字符串长度比偏移量小(比如字符串只有 `5` 个字符长，但你设置的 `offset` 是 `10` )， 那么原字符和偏移量之间的空白将用零字节(zerobytes, `"\x00"` )进行填充。

因为 Redis 字符串的大小被限制在 512 兆(megabytes)以内， 所以用户能够使用的最大偏移量为 2^29-1(536870911) ， 如果你需要使用比这更大的空间， 请使用多个 `key` 。

Warning

当生成一个很长的字符串时， Redis 需要分配内存空间， 该操作有时候可能会造成服务器阻塞(block)。 在2010年出产的Macbook Pro上， 设置偏移量为 536870911(512MB 内存分配)将耗费约 300 毫秒， 设置偏移量为 134217728(128MB 内存分配)将耗费约 80 毫秒， 设置偏移量 33554432(32MB 内存分配)将耗费约 30 毫秒， 设置偏移量为 8388608(8MB 内存分配)将耗费约 8 毫秒。


## 返回值

`SETRANGE` 命令会返回被修改之后， 字符串值的长度。

## 代码示例


对非空字符串执行 `SETRANGE` 命令：

```shell
redis> SET greeting "hello world"
OK

redis> SETRANGE greeting 6 "Redis"
(integer) 11

redis> GET greeting
"hello Redis"
```

对空字符串/不存在的键执行 `SETRANGE` 命令：

```shell
redis> EXISTS empty_string
(integer) 0

redis> SETRANGE empty_string 5 "Redis!"   # 对不存在的 key 使用 SETRANGE
(integer) 11

redis> GET empty_string                   # 空白处被"\x00"填充
"\x00\x00\x00\x00\x00Redis!"
```