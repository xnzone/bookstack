---
author: 黄健宏
title: SAVE
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21301
tags:
  - Redis
  - 持久化
  - SAVE
---

# SAVE

> 可用版本： >= 1.0.0
> 
> 时间复杂度： O(N)， `N` 为要保存到数据库中的 key 的数量。

[SAVE](#save) 命令执行一个同步保存操作，将当前 Redis 实例的所有数据快照(snapshot)以 RDB 文件的形式保存到硬盘。

一般来说，在生产环境很少执行 [SAVE](#save) 操作，因为它会阻塞所有客户端，保存数据库的任务通常由 [BGSAVE](bgsave.html#bgsave) 命令异步地执行。然而，如果负责保存数据的后台子进程不幸出现问题时， [SAVE](#save) 可以作为保存数据的最后手段来使用。

请参考文档： [Redis 的持久化运作方式(英文)](http://redis.io/topics/persistence) 以获取更多消息。

## 返回值

保存成功时返回 `OK` 。

## 代码示例

```bash
redis> SAVE
OK
```
