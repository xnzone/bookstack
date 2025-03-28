---
author: xnzone
title: "基本概念"
date: 2024-09-10 10:23:32
image: /covers/distr-fun-zh.jpg
cover: false
weight: 402
tags:
  - distributed
  - system
  - CAP
  - 分布式系统
---

在本章中，我们将上下探索抽象层次，查看一些不可能的结果（CAP和FLP），然后为了性能再回到较低的层次。

如果你有过编程经验，抽象层次的概念可能对你来说并不陌生。你总是会在某个抽象层次上工作，通过某个API与较低层次的层进行接口，并可能为用户提供某个更高层次的API或用户界面。七层[OSI计算机网络模型](https://en.wikipedia.org/wiki/OSI_model)就是一个很好的例子。

我认为，分布式编程在很大程度上是处理分布的后果（显而易见！）。也就是说，存在一种紧张关系：现实中有许多节点，而我们希望系统“像单一系统一样工作”。这意味着要找到一个良好的抽象，平衡可行性、可理解性和性能。

当我们说X比Y更抽象时，我们是什么意思？首先，X并没有引入任何新的或根本不同于Y的东西。实际上，X可能会去掉Y的一些方面，或者以更易于管理的方式呈现它们。其次，假设X从Y中去掉的东西对当前问题并不重要，那么X在某种意义上比Y更容易理解。

正如[Nietzsche](http://oregonstate.edu/instruct/phl201/modules/Philosophers/Nietzsche/Truth_and_Lie_in_an_Extra-Moral_Sense.htm)写道：

> 每个概念都是通过将不平等的事物等同而产生的。没有一片叶子是完全相同的，而“叶子”这个概念是通过对这些个体差异的任意抽象而形成的，通过遗忘这些区别；现在它引发了这样一个想法：在自然界中可能存在一些除了叶子之外的东西，这些东西会被称为“叶子”——某种原始形式，所有的叶子都是在这种形式的基础上编织、标记、复制、着色、卷曲和涂抹的，但都是由不熟练的手完成的，因此没有一个副本能够成为原始形式的正确、可靠和忠实的图像。

抽象从根本上来说是虚假的。每种情况都是独特的，每个节点也是如此。但抽象使得世界变得可管理：简化的问题陈述——摆脱现实——在分析上更易于处理，只要我们没有忽视任何本质的东西，解决方案就会广泛适用。

确实，如果我们保留的东西是本质的，那么我们可以推导出的结果将是广泛适用的。这就是为什么不可能结果如此重要：它们以问题的最简单形式展示，并证明在某些约束或假设下无法解决。

所有的抽象都忽略了一些东西，以便将现实中独特的事物等同起来。关键是去掉所有不必要的东西。你怎么知道什么是本质的？嗯，你可能无法事先知道。

每次我们从系统的规范中排除某个方面时，我们就有可能引入错误和/或性能问题的来源。这就是为什么有时我们需要朝相反的方向走，选择性地重新引入一些真实硬件和现实世界问题的某些方面。重新引入一些特定的硬件特性（例如，物理顺序性）或其他物理特性，可能足以使系统的性能达到足够的水平。

考虑到这一点，我们可以保留多少现实，同时仍然能够处理一些仍然可以被识别为分布式系统的东西？系统模型是我们认为重要特征的规范；在指定了一个模型后，我们可以查看一些不可能的结果和挑战。

## 一个系统模型

分布式系统的一个关键特性是分布。更具体地说，分布式系统中的程序：

- 在独立节点上并发运行...
- 通过可能引入非确定性和消息丢失的网络连接...
- 并且没有共享内存或共享时钟。

这带来了许多影响：

- 每个节点并发执行程序
- 知识是局部的：节点只能快速访问其本地状态，关于全局状态的任何信息都有可能过时
- 节点可以独立地故障和恢复
- 消息可能会延迟或丢失（与节点故障无关；很难区分网络故障和节点故障）
- 并且节点之间的时钟不同步（本地时间戳与全局实际时间顺序不对应，且无法轻易观察）

系统模型列举了与特定系统设计相关的许多假设。

系统模型

关于分布式系统实现环境和设施的一组假设

系统模型在对环境和设施的假设上有所不同。这些假设包括：

- 节点具有什么能力以及它们可能如何失败
- 通信链路如何操作以及它们可能如何失败
- 整个系统的属性，例如对时间和顺序的假设

一个稳健的系统模型是做出最弱假设的模型：为这样的系统编写的任何算法对不同环境非常宽容，因为它做出的假设非常少且非常弱。

另一方面，我们可以通过做出强假设来创建一个易于推理的系统模型。例如，假设节点不会失败意味着我们的算法不需要处理节点故障。然而，这样的系统模型是不现实的，因此很难在实践中应用。

让我们更详细地看看节点、链路以及时间和顺序的属性。

### 我们系统模型中的节点

节点作为计算和存储的主机。它们具有：

- 执行程序的能力
- 将数据存储到易失性内存中的能力（在故障时可能会丢失）和存储到稳定状态中的能力（在故障后可以读取）
- 一个时钟（可能被假设为准确，也可能不被假设）

节点执行确定性算法：局部计算、计算后的局部状态以及发送的消息都是由接收到的消息和接收消息时的局部状态唯一决定的。

有许多可能的故障模型描述节点可能发生故障的方式。在实践中，大多数系统假设崩溃恢复故障模型：即节点只能通过崩溃来失败，并且可以（可能）在某个后续时刻恢复。

另一种选择是假设节点可以以任何任意方式失效。这被称为[拜占庭容错](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance)。拜占庭故障在现实商业系统中很少被处理，因为对任意故障具有弹性的算法运行成本更高且实现更复杂。我在这里不讨论它们。

### 我们系统模型中的通信链路

通信链路将各个节点连接在一起，并允许消息在任一方向发送。许多讨论分布式算法的书籍假设每对节点之间都有单独的链路，这些链路为消息提供FIFO（先进先出）顺序，只能传递已发送的消息，并且已发送的消息可能会丢失。

一些算法假设网络是可靠的：消息永远不会丢失，也不会无限期延迟。这在某些现实环境中可能是合理的假设，但一般来说，考虑网络是不可靠的，并且可能会出现消息丢失和延迟是更可取的。

当网络故障而节点本身仍然正常运行时，就会发生网络分区。在这种情况下，消息可能会丢失或延迟，直到网络分区被修复。被分区的节点可能会被某些客户端访问，因此必须与崩溃的节点区别对待。下图说明了节点故障与网络分区的区别：

![](https://s2.loli.net/2024/11/18/SinhOBNoEQwMscm.png)

对通信链路做出进一步假设是很少见的。我们可以假设链路仅在一个方向上工作，或者我们可以为不同的链路引入不同的通信成本（例如，由于物理距离造成的延迟）。然而，除了长距离链路（广域网延迟）外，这些在商业环境中很少被关注，因此我在这里不讨论它们；更详细的成本和拓扑模型允许在复杂性增加的情况下进行更好的优化。

### 时间/顺序假设

物理分布的一个后果是，每个节点以独特的方式体验世界。这是不可避免的，因为信息只能以光速传播。如果节点之间的距离不同，那么从一个节点发送到其他节点的任何消息将在不同的时间到达，并且在其他节点处可能以不同的顺序到达。

时间假设是捕捉我们在多大程度上考虑这一现实的假设的方便简写。主要有两种替代方案：

**同步系统模型**

进程以锁步方式执行；消息传输延迟有已知的上限；每个进程都有一个准确的时钟。

**异步系统模型**

没有时间假设——例如，进程以独立的速率执行；消息传输延迟没有上限；有用的时钟不存在。

同步系统模型对时间和顺序施加了许多约束。它基本上假设节点有相同的体验：发送的消息总是在特定的最大传输延迟内被接收，并且进程以锁步方式执行。这是方便的，因为它允许你作为系统设计者对时间和顺序做出假设，而异步系统模型则不允许。

异步性是一种非假设：它只是假设你不能依赖时间（或“时间传感器”）。

在同步系统模型中解决问题更容易，因为对执行速度、最大消息传输延迟和时钟准确性的假设都有助于解决问题，因为你可以根据这些假设进行推理，并通过假设不发生不便的故障场景来排除它们。

当然，假设同步系统模型并不特别现实。现实世界的网络会遭遇故障，消息延迟没有硬性界限。现实世界的系统充其量是部分同步的：它们可能偶尔正常工作并提供一些上限，但也会有消息无限期延迟和时钟不同步的情况。我在这里不会真正讨论同步系统的算法，但你可能会在许多其他入门书籍中遇到它们，因为它们在分析上更简单（但不现实）。

### 共性问题

在接下来的文本中，我们将改变系统模型的参数。接下来，我们将讨论如何通过讨论两个不可能的结果（FLP和CAP）来看待以下两个系统属性的变化：

- 网络分区是否包含在故障模型中，以及
- 同步与异步的时间假设

当然，为了进行讨论，我们还需要引入一个待解决的问题。我将讨论的问题是[共识问题](https://en.wikipedia.org/wiki/Consensus_%28computer_science%29)。

如果多个计算机（或节点）就某个值达成一致，则称它们达成共识。更正式地说：

1. 一致性：每个正确的进程必须就相同的值达成一致。
2. 完整性：每个正确的进程最多决定一个值，如果它决定了某个值，则该值必须由某个进程提议。
3. 终止性：所有进程最终达成决策。
4. 有效性：如果所有正确的进程提议相同的值V，则所有正确的进程决定V。

共识问题是许多商业分布式系统的核心。毕竟，我们希望在不必处理分布后果（例如节点之间的分歧/偏差）的情况下，获得分布式系统的可靠性和性能，而解决共识问题使得解决一些相关的、更高级的问题（如原子广播和原子提交）成为可能。

### 两个不可能的结果

第一个不可能的结果，被称为FLP不可能性结果，是一个特别与设计分布式算法的人相关的不可能性结果。第二个结果——CAP定理——是一个相关的结果，更与从业者相关；那些需要在不同系统设计之间进行选择的人，但不直接关心算法设计。

## FLP不可能性结果

我将简要总结[FLP不可能性结果](https://en.wikipedia.org/wiki/Consensus_%28computer_science%29#Solvability_results_for_some_agreement_problems)，尽管它在学术界被认为是[更重要的](https://en.wikipedia.org/wiki/Dijkstra_Prize)。FLP不可能性结果（以作者Fischer、Lynch和Patterson命名）在异步系统模型下考察共识问题（技术上讲，是协议问题，这是一种非常弱的共识问题形式）。假设节点只能通过崩溃来失败；网络是可靠的，并且异步系统模型的典型时间假设成立：例如，消息延迟没有上限。

在这些假设下，FLP结果声明：“在一个受故障影响的异步系统中，不存在解决共识问题的（确定性）算法，即使消息永远不会丢失，最多只有一个进程可能失败，并且它只能通过崩溃（停止执行）来失败。”

这个结果意味着，在一个非常最小的系统模型下，没有办法以不会被无限期延迟的方式解决共识问题。论证是，如果这样的算法存在，那么就可以设计一个执行该算法的过程，在其中它将保持未决状态（“双值”）任意长的时间，通过延迟消息传递——这在异步系统模型中是允许的。因此，这样的算法不可能存在。

这个不可能性结果很重要，因为它强调了假设异步系统模型会导致权衡：解决共识问题的算法必须在消息传递的保证不成立时放弃安全性或活性。

这一见解对设计算法的人特别相关，因为它对我们知道在异步系统模型中可解决的问题施加了严格的约束。CAP定理是一个相关的定理，更与从业者相关：它做出稍微不同的假设（网络故障而非节点故障），并对选择系统设计的从业者有更明确的影响。

## CAP定理

CAP定理最初是计算机科学家Eric Brewer提出的一个猜想。它是一个流行且相当有用的思考方式，用于理解系统设计所做的保证之间的权衡。它甚至有一个由[Gilbert](http://www.comp.nus.edu.sg/~gilbert/biblio.html)和[Lynch](https://zh.wikipedia.org/wiki/Nancy_Lynch)提供的[正式证明](https://www.google.com/search?q=Brewer%27s+conjecture+and+the+feasibility+of+consistent%2C+available%2C+partition-tolerant+web+services)，而且，尽管[某个讨论网站](http://news.ycombinator.com/)认为Nathan Marz对此进行了反驳，但实际上并没有。

该定理指出，在以下三个属性中：

- 一致性：所有节点在同一时间看到相同的数据。
- 可用性：节点故障不会阻止存活的节点继续操作。
- 分区容忍性：系统在网络和/或节点故障导致消息丢失的情况下继续操作。

只有两个属性可以同时满足。我们甚至可以将其绘制成一个漂亮的图表，从三个属性中选择两个，给我们三种对应于不同交集的系统类型：

![](https://s2.loli.net/2024/11/18/6MOxmyJjlqF7I8U.png)

请注意，该定理指出中间部分（拥有所有三个属性）是不可实现的。然后我们得到三种不同的系统类型：

- CA（一致性 + 可用性）。示例包括全严格的仲裁协议，如两阶段提交。
- CP（一致性 + 分区容忍性）。示例包括在少数分区不可用的情况下的多数仲裁协议，如Paxos。
- AP（可用性 + 分区容忍性）。示例包括使用冲突解决的协议，如Dynamo。

CA和CP系统设计都提供相同的一致性模型：强一致性。唯一的区别在于CA系统不能容忍任何节点故障；CP系统可以在非拜占庭故障模型中容忍多达`f`个故障，给定`2f+1`个节点（换句话说，只要多数`f+1`保持正常，它可以容忍少数`f`个节点的故障）。原因很简单：

- CA系统不区分节点故障和网络故障，因此必须停止在所有地方接受写入，以避免引入分歧（多个副本）。它无法判断远程节点是否宕机，还是仅仅网络连接出现故障：因此，唯一安全的做法是停止接受写入。
- CP系统通过对分区两侧施加不对称行为来防止分歧（例如，保持单副本一致性）。它只保留多数分区，并要求少数分区变得不可用（例如，停止接受写入），这保留了一定程度的可用性（多数分区）并确保单副本一致性。

我将在讨论Paxos的复制章节中更详细地讨论这一点。重要的是，CP系统将网络分区纳入其故障模型，并使用像Paxos、Raft或视图戳记复制这样的算法区分多数分区和少数分区。CA系统不具备分区意识，历史上更为常见：它们通常使用两阶段提交算法，并在传统的分布式关系数据库中很常见。

假设发生分区，该定理简化为可用性和一致性之间的二元选择。

![](https://s2.loli.net/2024/11/18/FEpMHdjZKLvIY6m.png)

我认为可以从CAP定理中得出四个结论：

首先，许多早期分布式关系数据库系统的设计并没有考虑分区容忍性（例如，它们是CA设计）。分区容忍性是现代系统的重要属性，因为如果系统是地理分布的（许多大型系统都是），网络分区变得更加可能。

第二，在网络分区期间，强一致性和高可用性之间存在紧张关系。CAP定理说明了强保证与分布式计算之间的权衡。

在某种意义上，承诺一个由独立节点通过不可预测的网络连接组成的分布式系统“以一种与非分布式系统无异的方式运行”是相当疯狂的。

强一致性保证要求我们在分区期间放弃可用性。这是因为在继续接受分区两侧的写入时，无法防止两个无法相互通信的副本之间的分歧。

我们如何解决这个问题？通过加强假设（假设没有分区）或通过削弱保证。一致性可以与可用性（以及离线可访问性和低延迟的相关能力）进行权衡。如果“一致性”被定义为低于“所有节点在同一时间看到相同数据”的某种东西，那么我们可以同时拥有可用性和某种（较弱的）一致性保证。

第三，在正常操作中，强一致性和性能之间存在紧张关系。

强一致性/单副本一致性要求节点在每个操作上进行通信和达成一致。这导致在正常操作期间的高延迟。

如果你能接受其他于经典一致性模型的一致性模型，即允许副本滞后或分歧的一致性模型，那么你可以在正常操作期间减少延迟，并在分区情况下保持可用性。

当涉及的消息和节点较少时，操作可以更快完成。但实现这一点的唯一方法是放宽保证：让一些节点的联系频率降低，这意味着节点可能包含旧数据。

这也使得异常的发生成为可能。你不再保证获得最新的值。根据所做的保证类型，你可能会读取到比预期更旧的值，甚至丢失一些更新。

第四——间接地——如果我们不想在网络分区期间放弃可用性，那么我们需要探索是否有其他一致性模型可以满足我们的目的。

例如，即使用户数据被地理复制到多个数据中心，并且这两个数据中心之间的链接暂时失效，在许多情况下，我们仍然希望允许用户使用网站/服务。这意味着稍后需要调和两个不同的数据集，这既是技术挑战也是商业风险。但通常，技术挑战和商业风险都是可管理的，因此提供高可用性是更可取的。

一致性和可用性并不是真正的二元选择，除非你将自己限制在强一致性上。但强一致性只是一个一致性模型：在该模型中，你必须放弃可用性以防止多个数据副本同时处于活动状态。正如[Brewer本人指出的](http://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed)，对“2/3”的解释是误导性的。

如果你从这次讨论中只记住一个观点，那就是：“一致性”并不是一个单一、明确的属性。请记住：

> [ACID](https://en.wikipedia.org/wiki/ACID)一致性 !=  
> [CAP](https://en.wikipedia.org/wiki/CAP_theorem)一致性 !=  
> [Oatmeal](https://en.wikipedia.org/wiki/Oatmeal)一致性

相反，一致性模型是系统对使用它的程序所提供的任何保证。

一致性模型

程序员与系统之间的契约，其中系统保证如果程序员遵循某些特定规则，则对数据存储的操作结果将是可预测的。

CAP中的“C”是“强一致性”，但“一致性”并不是“强一致性”的同义词。

让我们看看一些替代的一致性模型。

## 强一致性与其他一致性模型

一致性模型可以分为两种类型：强一致性模型和弱一致性模型：

- 强一致性模型（能够维护单一副本）
    - 线性一致性
    - 顺序一致性
- 弱一致性模型（不是强一致性）
    - 客户端中心一致性模型
    - 因果一致性：可用的最强模型
    - 最终一致性模型

强一致性模型保证更新的表观顺序和可见性等同于非复制系统。另一方面，弱一致性模型则不做这样的保证。

请注意，这绝不是一个详尽的列表。同样，一致性模型只是程序员与系统之间的任意契约，因此它们几乎可以是任何东西。

### 强一致性模型

强一致性模型可以进一步分为两种相似但略有不同的一致性模型：

- _线性一致性_：在线性一致性下，所有操作**看起来**都是以与全局实时操作顺序一致的顺序原子性执行的。（Herlihy & Wing, 1991）
- _顺序一致性_：在顺序一致性下，所有操作**看起来**都是以某种顺序原子性执行的，该顺序与各个节点看到的顺序一致，并且在所有节点上相等。（Lamport, 1979）

关键区别在于，线性一致性要求操作生效的顺序等于实际的实时操作顺序。顺序一致性允许操作重新排序，只要在每个节点上观察到的顺序保持一致。区分这两者的唯一方法是观察进入系统的所有输入和时间；从与节点交互的客户端的角度来看，这两者是等价的。

这种区别似乎无关紧要，但值得注意的是，顺序一致性不具备组合性。

强一致性模型允许你作为程序员用一个分布式节点集群替换单个服务器，而不会遇到任何问题。

所有其他一致性模型都有异常（与保证强一致性的系统相比），因为它们的行为与非复制系统可区分。但通常这些异常是可以接受的，要么是因为我们不在乎偶尔出现的问题，要么是因为我们编写了处理不一致性的代码。

请注意，实际上并没有任何普遍适用的弱一致性模型分类，因为“不是强一致性模型”（例如，“在某种方式上与非复制系统可区分”）几乎可以是任何东西。

### 客户端中心一致性模型

_客户端中心一致性模型_是涉及客户端或会话概念的一种一致性模型。例如，客户端中心一致性模型可能保证客户端永远不会看到数据项的旧版本。这通常通过在客户端库中构建额外的缓存来实现，以便如果客户端移动到包含旧数据的副本节点，则客户端库返回其缓存的值，而不是来自副本的旧值。

如果客户端所在的副本节点不包含最新版本，客户端仍然可能会看到旧版本的数据，但他们永远不会看到旧版本的值重新出现的异常情况（例如，因为他们连接到不同的副本）。请注意，有许多种类的客户端中心一致性模型。

### 最终一致性

_最终一致性_模型表示，如果你停止更改值，那么在某个未定义的时间后，所有副本将达成一致，拥有相同的值。暗示在此之前，副本之间的结果在某种未定义的方式上是不一致的。由于它是[平凡可满足的](http://www.bailis.org/blog/safety-and-liveness-eventual-consistency-is-not-safe/)（仅具有活性属性），因此没有补充信息时是无用的。

说某事仅仅是最终一致性就像说“人最终会死”。这是一种非常弱的约束，我们可能希望对两件事有更具体的描述：

首先，“最终”是多长时间？有一个严格的下限，或者至少对系统通常需要多长时间才能收敛到相同值的想法是有用的。

其次，副本如何达成一致？一个始终返回“42”的系统是最终一致的：所有副本都同意相同的值。它只是不收敛到有用的值，因为它只是不断返回相同的固定值。相反，我们希望对方法有更好的了解。例如，一种决定的方法是让具有最大时间戳的值始终获胜。

因此，当供应商说“最终一致性”时，他们的意思是某种更精确的术语，例如“最终最后写入者获胜，并在此期间读取最新观察到的值”的一致性。“如何？”很重要，因为不好的方法可能导致写入丢失——例如，如果一个节点的时钟设置不正确并且使用了时间戳。

我将在关于弱一致性模型的复制方法章节中更详细地探讨这两个问题。

---

## 推荐阅读

- [Brewer's Conjecture and the Feasibility of Consistent, Available, Partition-Tolerant Web Services](http://lpd.epfl.ch/sgilbert/pubs/BrewersConjecture-SigAct.pdf) - Gilbert & Lynch, 2002
- [Impossibility of distributed consensus with one faulty process](https://scholar.google.com/scholar?q=Impossibility+of+distributed+consensus+with+one+faulty+process) - Fischer, Lynch and Patterson, 1985
- [Perspectives on the CAP Theorem](https://scholar.google.com/scholar?q=Perspectives+on+the+CAP+Theorem) - Gilbert & Lynch, 2012
- [CAP Twelve Years Later: How the "Rules" Have Changed](http://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed) - Brewer, 2012
- [Uniform consensus is harder than consensus](https://scholar.google.com/scholar?q=Uniform+consensus+is+harder+than+consensus) - Charron-Bost & Schiper, 2000
- [Replicated Data Consistency Explained Through Baseball](http://pages.cs.wisc.edu/~remzi/Classes/739/Papers/Bart/ConsistencyAndBaseballReport.pdf) - Terry, 2011
- [Life Beyond Distributed Transactions: an Apostate's Opinion](https://scholar.google.com/scholar?q=Life+Beyond+Distributed+Transactions%3A+an+Apostate%27s+Opinion) - Helland, 2007
- [If you have too much data, then 'good enough' is good enough](http://dl.acm.org/citation.cfm?id=1953140) - Helland, 2011
- [Building on Quicksand](https://scholar.google.com/scholar?q=Building+on+Quicksand) - Helland & Campbell, 2009