---
author: 黄健宏
title: QUIT
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21602
tags:
  - Redis
  - 客户端与服务器
  - QUIT
---

# QUIT

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(1)

请求服务器关闭与当前客户端的连接。

一旦所有等待中的回复(如果有的话)顺利写入到客户端，连接就会被关闭。

## 返回值

总是返回 `OK` (但是不会被打印显示，因为当时 Redis-cli 已经退出)。

## 代码示例

```bash
$ redis

redis> QUIT

$
```