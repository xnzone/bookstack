---
author: xnzone 
title: 实验操作
date: 2021-09-10
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1222
tags: ["xv6", "os", "mmu"]
---

## Exercise 1
>Exercise 1. In the file kern/pmap.c, you must implement code for the following functions (probably in the order given).

```c
boot_alloc()    
mem_init() (only up to the call to check_page_free_list(1))
page_init()   
page_alloc()   
page_free()  
```

>check_page_free_list() and check_page_alloc() test your physical page allocator. You should boot JOS and see whether check_page_alloc() reports success. Fix your code so that it passes. You may find it helpful to add your own assert()s to verify that your assumptions are correct.

练习1主要是为了熟悉内存分配的操作，页面初始化，页面内存分配以及页面内存释放等

`boot_alloc`主要是启动分配内存，目的是如果n大于0的话，就分配一个新的内存，否则就是返回下一个内存给它，可以看到代码里面的注释`ROUNDUP`就是向上取`n`的倍数，所以`boot_alloc`里面要填充的代码就很简单了。为了方便调试，可以把日志打印出来
```c
cprintf("boot_alloc memory at %x\n", nextfree);
cprintf("Next memory at %x\n", ROUNDUP((char *) (nextfree+n), PGSIZE));
if (n != 0) {
    char *next = nextfree;
    nextfree = ROUNDUP((char *) (nextfree + n), PGSIZE);
    return next;
} else {
    return nextfree;
}
```

使用`make qemu-nox`可以看到输出
```text
boot_alloc memory at f0114000
Next memory at f0115000
```

在`mem_init`中可以看到调用过`boot_alloc`是`boot_alloc(PGSIZE)`,`PGSIZE`是4096，是相吻合的。但是`boot_alloc`最开始的地址是`f0114000`就不是很清楚了

`mem_init`是初始化内存，其中有一个要求是需要初始化`pages`变量，这个变量是保存`struct PageInfo`的结构体的，只需要开辟一块内存空间即可。所以还是使用`boot_alloc`函数用来分配，就很简单了
```c
pages = (struct PageInfo*)boot_alloc(sizeof(struct PageInfo) * npages);
memset(pages, 0, npages * sizeof(struct PageInfo));
```

需要注意的是需要给结构体分配n个页面的大小，即`sizeof(struct PageInfo) * npages`，注意`sizeof`的使用

`page_init`就比较复杂了，不过根据注释来，也可以实现
```c
void
page_init(void)
{
    size_t i;
	// 1 Mark physical page 0 as in use
	pages[0].pp_ref = 1;
	// 2. [PGSIZE, npages_basemem * PGSIZE] is free
	for(i = 1; i < npages_basemem; i++) {
		pages[i].pp_ref = 0;
		pages[i].pp_link = page_free_list;
		page_free_list = &pages[i];
	}
	// 3. IO hole([IOPHYSMEM, EXTPHYSMEM]) never be allocated
	for(i = (IOPHYSMEM >> PGSHIFT); i < (EXTPHYSMEM >> PGSHIFT); i++) {
		pages[i].pp_ref = 1;
	}
	// 4. extended memory, some is free, some is used
	size_t first_free_page = (PADDR(boot_alloc(0))) >> PGSHIFT;
	for(i = (EXTPHYSMEM >> PGSHIFT); i < first_free_page; i++) {
		pages[i].pp_ref = 1;
	}
	for(i = first_free_page; i < npages; i++) {
		pages[i].pp_ref = 0;
		pages[i].pp_link = page_free_list;
		page_free_list = &pages[i];
	}
}
```

>Where is the kernel in physical memory?  Which pages are already in use for page tables and other data structures?

答案就很明显了，IO之后都是内核，最开始在`mem_init`初始化的`kern_pgdir`占用的内存是已经使用的

`page_alloc`是通过上面的页表进行分配内存，主要是从`page_free_list`中取一片内存划分出去，同时如果需要初始化的话，用`page2kva`函数进行初始化
```c
struct PageInfo *
page_alloc(int alloc_flags)
{
	// Fill this function in
	if(page_free_list == NULL) {
		return NULL;
	}
	struct PageInfo *ret = page_free_list;
	page_free_list = page_free_list->pp_link;
	if(alloc_flags & ALLOC_ZERO) {
		memset(page2kva(ret), 0, PGSIZE);
	}
	return ret;	
}
```

`page_release`是释放内存，就是把链表归还到`page_free_list`中，但是要注意`pp_ref ！= 0` 或者`pp_link != NULL`的情况，提示直接给出`panic`，这里应该不要给`panic`，因为会不通过这样的话，就很简单了。把释放链表的下一个节点改成`page_free_list`，然后更新`page_free_list`的指向就可以了
```c
void
page_free(struct PageInfo *pp)
{
	// Fill this function in
	// Hint: You may want to panic if pp->pp_ref is nonzero or
	// pp->pp_link is not NULL.
	pp->pp_link = page_free_list;
	page_free_list = pp;
}
```

完成之后，运行`make qemo-nox`，如果看见下面的输出，就说明没有问题
```bash
check_page_free_list() succeeded!
check_page_alloc() succeeded!
```

## Exercise 2
>Exercise 2. Look at chapters 5 and 6 of the Intel 80386 Reference Manual, if you haven't done so already. Read the sections about page translation and page-based protection closely (5.2 and 6.4). We recommend that you also skim the sections about segmentation; while JOS uses the paging hardware for virtual memory and protection, segment translation and segment-based protection cannot be disabled on the x86, so you will need a basic understanding of it.

练习2的目的是看80386的手册，熟悉段和页转换

## Exercise 3
>Exercise 3. While GDB can only access QEMU's memory by virtual address, it's often useful to be able to inspect physical memory while setting up virtual memory. Review the QEMU monitor commands from the lab tools guide, especially the xp command, which lets you inspect physical memory. To access the QEMU monitor, press Ctrl-a c in the terminal (the same binding returns to the serial console).

>Use the xp command in the QEMU monitor and the x command in GDB to inspect memory at corresponding physical and virtual addresses and make sure you see the same data.

>Our patched version of QEMU provides an info pg command that may also prove useful: it shows a compact but detailed representation of the current page tables, including all mapped memory ranges, permissions, and flags. Stock QEMU also provides an info mem command that shows an overview of which ranges of virtual addresses are mapped and with what permissions.

练习3的目的是使用GDB和QEMU来查看真实的物理内存地址。qemu的命令是`xp`,GDB的命令是`x`查看对应的地址(实际我也不会用)

qemu的`info pg`命令也可以，gdb的`info mem`命令(不会用，放过了)

## Question 1
>Assuming that the following JOS kernel code is correct, what type should variable x have, uintptr_t or physaddr_t?
```c
mystery_t x;
char* value = return_a_pointer();
*value = 10;
x = (mystery_t) value;
```

特别注意此处 `uintptr_t` 和 `physaddr_t`  是 32-bit 整型，不能直接解引用，除非将其转换为指针类型，但请注意将物理地址转换为指针型再去寻址没用意义，因为 MMU 会将其看作是虚拟地址，从而会出错。解引用针对的都虚拟地址。

因此，变量 x 的类型应该是 uintptr_t 。

## Exercise 4
>Exercise 4. In the file kern/pmap.c, you must implement code for the following functions.
```c
pgdir_walk()
boot_map_region()
page_lookup()
page_remove()
page_insert()
```

>check_page(), called from mem_init(), tests your page table management routines. You should make sure it reports success before proceeding.

`pgdir_walk`是给定页目录指针`pgdir`，返回指向线性地址页表入口(PTE)的指针。需要走两级页表目录。

提示：
1. `page2pa`是把`struct PageInfo`转换成物理地址;`page2kva`是把`struct PageInfo`转成线性地址，内部调用`page2pa`和`KADDR`，`KADDR`是把物理地址转成虚拟地址；
2. x86的MMU检查页表权限位，因此在页表目录中也保持权限位是有必要的，翻译过来就是分配的地址，最后要加上权限位
3. `inc/mmu.h`中定义的宏是非常有用的，需要配合使用完成这个代码

有了上面的提示和分析，就知道`pgdir_walk`是干嘛的了，核心还是`page2kva`转换成虚拟地址，里面还有额外的分配要求，补上就可以了
```c
pte_t *
pgdir_walk(pde_t *pgdir, const void *va, int create)
{
    pde_t *pde = &pgdir[PDX(va)];
    pte_t *ret;
    if(*pde & PTE_P) { // 存在的话，不需要重新分配
        ret = KADDR(PTE_ADDR(*pde));
    } else {
        if(!create) return NULL; // 不创建的话，返回
        struct PageInfo *pp = page_alloc(ALLOC_ZERO);
        struct PageInfo* pp = page_alloc(ALLOC_ZERO);
        if (pp == NULL) {
            return NULL;
        }
        pp->pp_ref++;
        ret = (pte_t*)page2kva(pp);
        *pde = PADDR(ret) | PTE_P | PTE_W | PTE_U;
    }
    return &ret[PTX(va)];
}
```

`boot_map_region`是用来映射地址的，把`[va, va+size]`的虚拟地址映射到`[pa, pa + size]的物理地址。只是设置`UTOP`上面的静态映射关系
```c
static void
boot_map_region(pde_t *pgdir, uintptr_t va, size_t size, physaddr_t pa, int perm)
{
	// Fill this function in
	int i;
	for(i = 0; i < size / PGSIZE; i++, va += PGSIZE, pa += PGSIZE) {
		pte_t *pte = pgdir_walk(pgdir, (void *)va, 1);
		if(!pte) panic("boot_map_region err, out of memory");
		*pte = pa | perm | PTE_P;
	}
}
```

`page_lookup`返回在虚拟地址映射的页面。还有提示使用`pgdir_walk`和`pa2page`
```c
struct PageInfo *
page_lookup(pde_t *pgdir, void *va, pte_t **pte_store)
{
	// Fill this function in
	pte_t *pte = pgdir_walk(pgdir, va, 0);
	if(!pte || !(*pte & PTE_P)) {
		return NULL;
	}
	if(pte_store) {
		*pte_store = pte;
	}
	return pa2page(PTE_ADDR(*pte));
}
```

`page_remove`移除虚拟地址中没有映射的页面
```c
void
page_remove(pde_t *pgdir, void *va)
{
	// Fill this function in
	pte_t *pte;
	struct PageInfo *pp = page_lookup(pgdir, va, &pte);
	if(!pp || !(*pte & PTE_P)) {
		return;
	}

	page_decref(pp);
	*pte = 0;
	tlb_invalidate(pgdir, va);
}
```

`page_insert`把物理地址映射到虚拟地址，还有一堆注释，注意阅读使用
```c
int
page_insert(pde_t *pgdir, struct PageInfo *pp, void *va, int perm)
{
	// Fill this function in
	pte_t *pte = pgdir_walk(pgdir, va, 1);
	if(!pte) {
		return -E_NO_MEM;
	}
	pp->pp_ref++;
	if(*pte & PTE_P) {
		page_remove(pgdir, va);
	}
	*pte = page2pa(pp) | perm |PTE_P;
	return 0;
}
```

运行出现`check_page() succeeded!`就是成功

## Exercise 5
>Exercise 5. Fill in the missing code in mem_init() after the call to check_page().
>Your code should now pass the check_kern_pgdir() and check_page_installed_pgdir() checks.

练习5是填写`mem_init`在`check_page`之后的代码

```c
// UPAGES
boot_map_region(kern_pgdir, UPAGES, PTSIZE, PADDR(pages), PTE_U);
// KSTACKTOP
boot_map_region(kern_pgdir, KSTACKTOP - KSTKSIZE, KSTKSIZE, PADDR(bootstack), PTE_W);
// KERNBASE
boot_map_region(kern_pgdir, KERNBASE, 0xffffffff - KERNBASE, 0, PTE_W);
```

运行出现结果
```bash
check_kern_pgdir() succeeded!
check_page_free_list() succeeded!
check_page_installed_pgdir() succeeded!
```

## Question
> What entries (rows) in the page directory have been filled in at this point? What addresses do they map and where do they point? In other words, fill out this table as much as possible:

|Entry|Base Virtual Address|Points to (logically):|
|---|---|---|
|1023|?|Page table for top 4MB of phys memory|
|1022|?|?|
|.|?|?|
|.|?|?|
|.|?|?|
|2|0x00800000|?|
|1|0x00400000|?|
|0|0x00000000|[see next question]|


最终的结果是
|Entry|Base Virtual Address|Points to (logically):|
|---|---|---|
|1023|0xffc00000|Page table for top 4MB of phys memory|
|...|...|...|
|960|0xf0000000|Page table for [0,4)MB of phys memory|
|959|0xefc00000|Kernel Stack and Invalid Memory|
|...|...|...|
|957|0xef400000|UVPT, User read-only virtual page table|
|956|0xef000000|UPAGES, Read-only copies of the Page structures|
|...|...|...|


>We have placed the kernel and user environment in the same address space. Why will user programs not be able to read or write the kernel's memory? What specific mechanisms protect the kernel memory?

Page-Directory Entry 和 Page-Table Entry 的U/S位(User/supervisor)，如果没有将其置1，那么用户将没有访问权限

>What is the maximum amount of physical memory that this operating system can support? Why?

2G, 因为`UPAGES`是4MB,而`sizeof(struct PageInfo) = 8Byte`， 所以能够使用`4MB/8B = 512K`页， 一页的大小是4KB， 所以最多使用`512K * 4KB = 2G`的物理内存

>How much space overhead is there for managing memory, if we actually had the maximum amount of physical memory? How is this overhead broken down?

4MB的 PageInfos to manage memory plus 2MB for Page Table plus 4KB for Page Directory if we have 2GB physical memory. Total:6MB+4KB

4MB的`PageInfo`， 2MB的`Page Table`， 4KB的`Page Directory` 所以总共是`6MB + 4KB`

>Revisit the page table setup in kern/entry.S and kern/entrypgdir.c. Immediately after we turn on paging, EIP is still a low number (a little over 1MB). At what point do we transition to running at an EIP above KERNBASE? What makes it possible for us to continue executing at a low EIP between when we enable paging and when we begin running at an EIP above KERNBASE? Why is this transition necessary?

在执行`jmp *%eax`完成之后,之前的实验也说清楚了

## 挑战
挑战的代码就不写了