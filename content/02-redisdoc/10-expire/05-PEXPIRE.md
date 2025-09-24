---
author: 黄健宏
title: PEXPIRE
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21005
tags:
  - Redis
  - 自动过期
  - PEXPIRE
---

# PEXPIRE key milliseconds

> 可用版本： >= 2.6.0
> 
> 时间复杂度： O(1)

这个命令和 `EXPIRE` 命令的作用类似，但是它以毫秒为单位设置 `key` 的生存时间，而不像 `EXPIRE` 命令那样，以秒为单位。

## 返回值

设置成功，返回 `1` `key` 不存在或设置失败，返回 `0`

## 代码示例

```bash
redis> SET mykey "Hello"
OK

redis> PEXPIRE mykey 1500
(integer) 1

redis> TTL mykey    # TTL 的返回值以秒为单位
(integer) 2

redis> PTTL mykey   # PTTL 可以给出准确的毫秒数
(integer) 1499
```
