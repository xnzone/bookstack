---
author: 黄健宏
title: BGREWRITEAOF
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21303
tags:
  - Redis
  - 持久化
  - BGREWRITEAOF
---

# BGREWRITEAOF

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(N)， `N` 为要追加到 AOF 文件中的数据数量。

执行一个 [AOF文件](http://redis.io/topics/persistence#append-only-file) 重写操作。重写会创建一个当前 AOF 文件的体积优化版本。

即使 [BGREWRITEAOF](#bgrewriteaof) 执行失败，也不会有任何数据丢失，因为旧的 AOF 文件在 [BGREWRITEAOF](#bgrewriteaof) 成功之前不会被修改。

重写操作只会在没有其他持久化工作在后台执行时被触发，也就是说：

- 如果 Redis 的子进程正在执行快照的保存工作，那么 AOF 重写的操作会被预定(scheduled)，等到保存工作完成之后再执行 AOF 重写。在这种情况下， [BGREWRITEAOF](#bgrewriteaof) 的返回值仍然是 `OK` ，但还会加上一条额外的信息，说明 [BGREWRITEAOF](#bgrewriteaof) 要等到保存操作完成之后才能执行。在 Redis 2.6 或以上的版本，可以使用 [INFO [section]](../client_and_server/info.html#info) 命令查看 [BGREWRITEAOF](#bgrewriteaof) 是否被预定。
    
- 如果已经有别的 AOF 文件重写在执行，那么 [BGREWRITEAOF](#bgrewriteaof) 返回一个错误，并且这个新的 [BGREWRITEAOF](#bgrewriteaof) 请求也不会被预定到下次执行。
    

从 Redis 2.4 开始， AOF 重写由 Redis 自行触发， [BGREWRITEAOF](#bgrewriteaof) 仅仅用于手动触发重写操作。

请移步 [持久化文档(英文)](http://redis.io/topics/persistence) 查看更多相关细节。

## 返回值

反馈信息。

## 代码示例

```bash
redis> BGREWRITEAOF
Background append only file rewriting started
```