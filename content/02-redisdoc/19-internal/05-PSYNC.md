---
author: 黄健宏
title: PSYNC
date: 2024-12-29 10:32:21
image: https://s2.loli.net/2025/09/24/hzeyjtaJWSTmg32.png
cover: false
weight: 21905
tags:
  - Redis
  - 内部命令
  - PSYNC
---

# PSYNC master_run_id offset

> 可用版本： >= 2.8.0
> 
> 时间复杂度： 不明确

用于复制功能(replication)的内部命令。

更多信息请参考 [复制（Replication）](../topic/replication.html#replication-topic) 文档。

## 返回值

序列化数据。

## 代码示例

```bash
127.0.0.1:6379> PSYNC ? -1
"REDIS0006\xfe\x00\x00\x02kk\x02vv\x00\x03msg\x05hello\xff\xc3\x96P\x12h\bK\xef"
```