---
author: 黄健宏
title: MIGRATE
date: 2024-12-29 10:32:21
image: /covers/02-redisdoc.jpg
cover: false
weight: 21901
tags:
  - Redis
  - 内部命令
  - MIGRATE
---

# MIGRATE host port key destination-db timeout [COPY] [REPLACE]

> 可用版本： >= 2.6.0
> 
> 时间复杂度： 这个命令在源实例上实际执行 `DUMP` 命令和 `DEL` 命令，在目标实例执行 `RESTORE` 命令，查看以上命令的文档可以看到详细的复杂度说明。 `key` 数据在两个实例之间传输的复杂度为 O(N) 。

将 `key` 原子性地从当前实例传送到目标实例的指定数据库上，一旦传送成功， `key` 保证会出现在目标实例上，而当前实例上的 `key` 会被删除。

这个命令是一个原子操作，它在执行的时候会阻塞进行迁移的两个实例，直到以下任意结果发生：迁移成功，迁移失败，等待超时。

命令的内部实现是这样的：它在当前实例对给定 `key` 执行 `DUMP` 命令 ，将它序列化，然后传送到目标实例，目标实例再使用 `RESTORE` 对数据进行反序列化，并将反序列化所得的数据添加到数据库中；当前实例就像目标实例的客户端那样，只要看到 `RESTORE` 命令返回 `OK` ，它就会调用 `DEL` 删除自己数据库上的 `key` 。

`timeout` 参数以毫秒为格式，指定当前实例和目标实例进行沟通的**最大间隔时间**。这说明操作并不一定要在 `timeout` 毫秒内完成，只是说数据传送的时间不能超过这个 `timeout` 数。

`MIGRATE` 命令需要在给定的时间规定内完成 IO 操作。如果在传送数据时发生 IO 错误，或者达到了超时时间，那么命令会停止执行，并返回一个特殊的错误： `IOERR` 。

当 `IOERR` 出现时，有以下两种可能：

- `key` 可能存在于两个实例
    
- `key` 可能只存在于当前实例
    

唯一不可能发生的情况就是丢失 `key` ，因此，如果一个客户端执行 `MIGRATE` 命令，并且不幸遇上 `IOERR` 错误，那么这个客户端唯一要做的就是检查自己数据库上的 `key` 是否已经被正确地删除。

如果有其他错误发生，那么 `MIGRATE` 保证 `key` 只会出现在当前实例中。（当然，目标实例的给定数据库上可能有和 `key` 同名的键，不过这和 `MIGRATE` 命令没有关系）。

## 可选项

- `COPY` ：不移除源实例上的 `key` 。
    
- `REPLACE` ：替换目标实例上已存在的 `key` 。
    

## 返回值

迁移成功时返回 `OK` ，否则返回相应的错误。

## 代码示例

先启动两个 Redis 实例，一个使用默认的 6379 端口，一个使用 7777 端口。

$ ./redis-server &
[1] 3557

...

$ ./redis-server --port 7777 &
[2] 3560

...

然后用客户端连上 6379 端口的实例，设置一个键，然后将它迁移到 7777 端口的实例上：

$ ./redis-cli

redis 127.0.0.1:6379> flushdb
OK

redis 127.0.0.1:6379> SET greeting "Hello from 6379 instance"
OK

redis 127.0.0.1:6379> MIGRATE 127.0.0.1 7777 greeting 0 1000
OK

redis 127.0.0.1:6379> EXISTS greeting                           # 迁移成功后 key 被删除
(integer) 0

使用另一个客户端，查看 7777 端口上的实例：

$ ./redis-cli -p 7777

redis 127.0.0.1:7777> GET greeting
"Hello from 6379 instance"