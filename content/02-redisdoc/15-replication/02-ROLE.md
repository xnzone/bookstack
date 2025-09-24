---
author: 黄健宏
title: ROLE
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21502
tags:
  - Redis
  - 复制
  - ROLE
---

# ROLE

> 可用版本： >= 2.8.12
> 
> 时间复杂度： O(1)

返回实例在复制中担任的角色， 这个角色可以是 `master` 、 `slave` 或者 `sentinel` 。 除了角色之外， 命令还会返回与该角色相关的其他信息， 其中：

- 主服务器将返回属下从服务器的 IP 地址和端口。
    
- 从服务器将返回自己正在复制的主服务器的 IP 地址、端口、连接状态以及复制偏移量。
    
- Sentinel 将返回自己正在监视的主服务器列表。
    

## 返回值

`ROLE` 命令将返回一个数组。

## 代码示例

```bash
### 主服务器

1) "master"
2) (integer) 3129659
3) 1) 1) "127.0.0.1"
      2) "9001"
      3) "3129242"
   2) 1) "127.0.0.1"
      2) "9002"
      3) "3129543"

### 从服务器

1) "slave"
2) "127.0.0.1"
3) (integer) 9000
4) "connected"
5) (integer) 3167038

### Sentinel

1) "sentinel"
2) 1) "resque-master"
   2) "html-fragments-master"
   3) "stats-master"
   4) "metadata-master"
```