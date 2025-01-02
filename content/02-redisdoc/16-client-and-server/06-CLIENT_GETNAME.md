---
author: 黄健宏
title: CLIENT GETNAME
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21606
tags:
  - Redis
  - 客户端与服务器
  - CLIENT GETNAME
---

# CLIENT GETNAME

> 可用版本： >= 2.6.9
> 
> 时间复杂度： O(1)

返回 `CLIENT SETNAME` 命令为连接设置的名字。

因为新创建的连接默认是没有名字的， 对于没有名字的连接， `CLIENT GETNAME` 返回空白回复。

## 返回值

如果连接没有设置名字，那么返回空白回复； 如果有设置名字，那么返回名字。

## 代码示例

```bash
# 新连接默认没有名字

redis 127.0.0.1:6379> CLIENT GETNAME
(nil)

# 设置名字

redis 127.0.0.1:6379> CLIENT SETNAME hello-world-connection
OK

# 返回名字

redis 127.0.0.1:6379> CLIENT GETNAME
"hello-world-connection"
```