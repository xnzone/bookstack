---
author: 黄健宏
title: DEBUG OBJECT
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21806
tags:
  - Redis
  - 调试
  - DEBUG OBJECT
---

# DEBUG OBJECT key

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

`DEBUG OBJECT` 是一个调试命令，它不应被客户端所使用。

查看 `OBJECT` 命令获取更多信息。

## 返回值

当 `key` 存在时，返回有关信息。 当 `key` 不存在时，返回一个错误。

## 代码示例

```bash
redis> DEBUG OBJECT my_pc
Value at:0xb6838d20 refcount:1 encoding:raw serializedlength:9 lru:283790 lru_seconds_idle:150

redis> DEBUG OBJECT your_mac
(error) ERR no such key
```