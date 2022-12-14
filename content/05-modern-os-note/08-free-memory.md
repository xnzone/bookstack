---
author: xnzone 
title: "空闲内存管理"
date: 2022-09-10 10:23:32
image: /covers/os.png
cover: false
weight: 12
tags: ["os", "操作系统"]
---

## 空闲内存管理需要解决什么问题

空闲内存管理是一个广泛的话题，这个内容不仅仅用于操作系统，其他的语言设计也有类似的设计，比如C++的STL的内存分配管理，Go语言的内存管理。它是一个程序设计过程中经常会讨论的问题

我们给定一个大的内存空间给你，你每次程序都需要分配内存，分配到最后，这一块大内存的碎片就会越来越多，这也是我们所说的外部碎片，空闲内存管理主要是解决外部碎片的问题，让外部碎片尽可能小，内存利用率尽可能高

## 底层机制
空闲内存管理可以用位图和链表两种数据结构来实现，一般来说用链表会比较多，这里也只讨论链表管理空闲内存。空闲内存管理一般需要提供以下操作：

- 分割
- 合并
- 追踪已分配的内存空间
- 嵌入空闲列表
- 让堆增长

当然这些操作不一定都需要提供，但是最基本的分割、合并和追踪三个一定要提供。

## 基本策略
空闲内存管理的基本策略有以下几种

- 最优匹配
- 最差匹配
- 首次匹配
- 下次匹配

除了基本的策略外，还有一些技术和算法来改进内存分配，常用的有

- 分离空闲列表。C++的STL使用此方法
- 伙伴系统
- 使用树优化查询效率

下面会简单介绍这几个策略

### 最优匹配
最优匹配（best fit）策略非常简单：首先遍历整个空闲列表，找到和请求大小一样或更大的空闲块，然后返回这组候选者中最小的一块。这就是所谓的最优匹配（也可以称为最小匹配）。只需要遍历一次空闲列表，就足以找到正确的块并返回。

最优匹配背后的想法很简单：选择最接近用户请求大小的块，从而尽量避免空间浪费。然而，这有代价。简单的实现在遍历查找正确的空闲块时，要付出较高的性能代价。

简单来说就是找到最小一个刚好满足分配大小的内存进行分配

### 最差匹配
最差匹配（worst fit）方法与最优匹配相反，它尝试找最大的空闲块，分割并满足用户需求后，将剩余的块（很大）加入空闲列表。最差匹配尝试在空闲列表中保留较大的块，而不是像最优匹配那样可能剩下很多难以利用的小块。但是，最差匹配同样需要遍历整个空闲列表。更糟糕的是，大多数研究表明它的表现非常差，导致过量的碎片，同时还有很高的开销。

简单来说就是找到一个最大一个满足分配大小的内存进行分配。很显然不可取

### 首次匹配
首次匹配（first fit）策略就是找到第一个足够大的块，将请求的空间返回给用户。同样，剩余的空闲空间留给后续请求。

首次匹配有速度优势（不需要遍历所有空闲块），但有时会让空闲列表开头的部分有很多小块。因此，分配程序如何管理空闲列表的顺序就变得很重要。一种方式是基于地址排序（address-based ordering）。通过保持空闲块按内存地址有序，合并操作会很容易，从而减少了内存碎片

简单来说就是找到第一个比满足分配大小的内存进行分配

### 下次匹配
不同于首次匹配每次都从列表的开始查找，下次匹配（next fit）算法多维护一个指针，指向上一次查找结束的位置。其想法是将对空闲空间的查找操作扩散到整个列表中去，避免对列表开头频繁的分割。这种策略的性能与首次匹配很接近，同样避免了遍历查找。

简单来说就是从上一次分配的位置开始，找到第一个满足分配大小的内存进行分配

### 分离空闲列表
基本想法很简单：如果某个应用程序经常申请一种（或几种）大小的内存空间，那就用一个独立的列表，只管理这样大小的对象。其他大小的请求都交给更通用的内存分配程序

简单来说就是维护固定长度大小内存的链表，每次分配的时候，从刚好满足的里面挑一个空闲内存进行分配

这种方法的好处显而易见。通过拿出一部分内存专门满足某种大小的请求，碎片就不再是问题了。而且，由于没有复杂的列表查找过程，这种特定大小的内存分配和释放都很快。

具体来说，在内核启动时，它为可能频繁请求的内核对象创建一些对象缓存（object cache），如锁和文件系统 inode 等。这些的对象缓存每个分离了特定大小的空闲列表，因此能够很快地响应内存请求和释放。如果某个缓存中的空闲空间快耗尽时，它就向通用内存分配程序申请一些内存厚块（slab）（总量是页大小和对象大小的公倍数）。相反，如果给定厚块中对象的引用计数变为 0，通用的内存分配程序可以从专门的分配程序中回收这些空间，这通常发生在虚拟内存系统需要更多的空间的时候。这个思想到现在很多系统设计也在采用，就算是平时的业务代码中都可能涉及到的思想

### 伙伴系统
因为合并对分配程序很关键，所以人们设计了一些方法，让合并变得简单。在这种系统中，空闲空间首先从概念上被看成大小为 2N 的大空间。当有一个内存分配请求时，空闲空间被递归地一分为二，直到刚好可以满足请求的大小（再一分为二就无法满足）。这时，请求的块被返回给用户

伙伴系统的漂亮之处在于块被释放时。如果将这个 8KB 的块归还给空闲列表，分配程序会检查“伙伴”8KB 是否空闲。如果是，就合二为一，变成 16KB 的块。然后会检查这个 16KB 块的伙伴是否空闲，如果是，就合并这两块。这个递归合并过程继续上溯，直到合并整个内存区域，或者某一个块的伙伴还没有被释放。

伙伴系统运转良好的原因，在于很容易确定某个块的伙伴。


## 小结
这里主要讨论了空闲内存管理的一些基本操作，这一部分的很多思想都会被用于业务代码以及底层系统的设计，所以这部分内容后续也需要更深入了解