---
author: mixu
title: 附录
date: 2024-10-10 10:23:32
image: https://s2.loli.net/2025/09/24/lcAnB5wXs3GbCRm.jpg
cover: false
weight: 406
tags: ["distributed system", "分布式系统"]
---

如果你已经读到了这里，谢谢你。

如果你喜欢这本书，请在[Github](https://github.com/mixu/)（或[Twitter](https://twitter.com/mikitotakada)）上关注我。我喜欢看到我产生了一些积极的影响。“创造的价值超过你捕获的价值”，诸如此类的话。

非常感谢以下人士的帮助：logpath, alexras, globalcitizen, graue, frankshearar, roryokane, jpfuentes2, eeror, cmeiklejohn, stevenproctor eos2102 和 steveloughran！当然，任何剩余的错误和遗漏都是我的错！

值得注意的是，我关于最终一致性的章节相当偏向伯克利；我想改变这一点。我还跳过了时间的一个突出用例：一致的快照。还有一些主题我应该扩展：即，对安全性和活性属性的明确讨论以及对一致性哈希的更详细讨论。然而，我要去[Strange Loop 2013](https://thestrangeloop.com/)了，所以随它去吧。

如果这本书有第六章，它可能会关于如何利用和处理大量数据的方法。看来最常见的“大数据”计算类型是将[一个大型数据集通过一个简单的程序](https://en.wikipedia.org/wiki/SPMD)。我不确定后续章节会是什么（也许是高性能计算，鉴于目前的重点是可行性），但我可能几年后会知道。

## 关于分布式系统的书籍

#### Distributed Algorithms (Lynch)

这可能是关于分布式算法最常被推荐的书籍。我也推荐它，但有一个警告。它非常全面，但是为研究生读者编写的，所以你会花很多时间阅读关于同步系统和共享内存算法的内容，然后才能接触到对从业者来说最有趣的部分。

#### Introduction to Reliable and Secure Distributed Programming (Cachin, Guerraoui & Rodrigues)

对于从业者来说，这是一本有趣的书。它篇幅短，充满了实际的算法实现。

#### Replication: Theory and Practice

如果你对复制感兴趣，这本书太棒了。关于复制的章节在很大程度上是基于这本书的有趣部分以及更多近期阅读材料的综合。

#### Distributed Systems: An Algorithmic Approach (Ghosh)

#### Introduction to Distributed Algorithms (Tel)

#### Transactional Information Systems: Theory, Algorithms, and the Practice of Concurrency Control and Recovery (Weikum & Vossen)

这本书是关于传统的事务信息系统的，例如本地RDBMS。最后有两章是关于分布式事务的，但书的重点是事务处理。

#### Transaction Processing: Concepts and Techniques by Gray and Reuter

一本经典之作。我发现Weikum & Vossen的内容更更新。

## Seminal papers

Each year, the [Edsger W. Dijkstra Prize in Distributed Computing](https://en.wikipedia.org/wiki/Dijkstra_Prize) is given to outstanding papers on the principles of distributed computing. Check out the link for the full list, which includes classics such as:

- "[Time, Clocks and Ordering of Events in a Distributed System](http://research.microsoft.com/users/lamport/pubs/time-clocks.pdf)" - Leslie Lamport
- "[Impossibility of Distributed Consensus With One Faulty Process](http://theory.lcs.mit.edu/tds/papers/Lynch/jacm85.pdf)" - Fisher, Lynch, Patterson
- "[Unreliable failure detectors and reliable distributed systems](https://scholar.google.com/scholar?q=Unreliable+Failure+Detectors+for+Reliable+Distributed+Systems)" - Chandra and Toueg

Microsoft Academic Search has a list of [top publications in distributed & parallel computing ordered by number of citations](http://libra.msra.cn/RankList?entitytype=1&topDomainID=2&subDomainID=16&last=0&start=1&end=100) - this may be an interesting list to skim for more classics.

Here are some additional lists of recommended papers:

- [Nancy Lynch's recommended reading list](http://courses.csail.mit.edu/6.852/08/handouts/handout3.pdf) from her course on Distributed systems.
- [NoSQL Summer paper list](http://nosqlsummer.org/papers) - a curated list of papers related to this buzzword.
- [A Quora question on seminal papers in distributed systems](https://www.quora.com/What-are-the-seminal-papers-in-distributed-systems-Why).

### Systems

- [The Google File System](https://research.google.com/archive/gfs.html) - Ghemawat, Gobioff and Leung
- [MapReduce: Simplified Data Processing on Large Clusters](https://research.google.com/archive/mapreduce.html) - Dean and Ghemawat
- [Dynamo: Amazon’s Highly Available Key-value Store](https://scholar.google.com/scholar?q=Dynamo%3A+Amazon%27s+Highly+Available+Key-value+Store) - DeCandia et al.
- [Bigtable: A Distributed Storage System for Structured Data](https://research.google.com/archive/bigtable.html) - Chang et al.
- [The Chubby Lock Service for Loosely-Coupled Distributed Systems](https://research.google.com/archive/chubby.html) - Burrows
- [ZooKeeper: Wait-free coordination for Internet-scale systems](http://labs.yahoo.com/publication/zookeeper-wait-free-coordination-for-internet-scale-systems/) - Hunt, Konar, Junqueira, Reed, 2010