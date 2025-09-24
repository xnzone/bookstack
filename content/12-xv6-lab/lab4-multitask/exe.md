---
author: xnzone 
title: 实验操作
date: 2021-09-10
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1242
tags: ["xv6", "os", "multitask"]
---
## Exercise 1
>Exercise 1. Implement mmio_map_region in kern/pmap.c. To see how this is used, look at the beginning of lapic_init in kern/lapic.c. You'll have to do the next exercise, too, before the tests for mmio_map_region will run.

补全`mmio_map_region`，主要是利用`boot_map_region`进行补全代码。同时要注意页面对齐的问题
```c
void *
mmio_map_region(physaddr_t pa, size_t size)
{
    static uintptr_t base = MMIOBASE;
	size = ROUNDUP(pa + size, PGSIZE);
	pa = ROUNDDOWN(pa, PGSIZE);
	size -= pa;
	if(base + size >= MMIOLIM) {
		panic("not enough memory");
	}
	boot_map_region(kern_pgdir, base, size, pa, PTE_PCD | PTE_PWT | PTE_W);
	base += size;
	return (void*)(base - size);
}
```

## Exercise 2
>Exercise 2. Read boot_aps() and mp_main() in kern/init.c, and the assembly code in kern/mpentry.S. Make sure you understand the control flow transfer during the bootstrap of APs. Then modify your implementation of page_init() in kern/pmap.c to avoid adding the page at MPENTRY_PADDR to the free list, so that we can safely copy and run AP bootstrap code at that physical address. Your code should pass the updated check_page_free_list() test (but might fail the updated check_kern_pgdir() test, which we will fix soon).

主要是在代码中把`MPENTRY_PADDR`的页面过滤掉
```c
void
page_init(void)
{
    for(i = 1; i < npages_basemem; i++) {
		if(i == MPENTRY_PADDR >> PGSHIFT) {
			pages[i].pp_ref = 1;
			continue;
		}
		pages[i].pp_ref = 0;
		pages[i].pp_link = page_free_list;
		page_free_list = &pages[i];
	}
}
```

运行`make qemu-nox`，如果出现下面内容，则说明通过
```bash
check_page_free_list() succeeded!
check_page_alloc() succeeded!
check_page() succeeded!
```

## Question
>1.Compare kern/mpentry.S side by side with boot/boot.S. Bearing in mind that kern/mpentry.S is compiled and linked to run above KERNBASE just like everything else in the kernel, what is the purpose of macro MPBOOTPHYS? Why is it necessary in kern/mpentry.S but not in boot/boot.S? In other words, what could go wrong if it were omitted in kern/mpentry.S?

>Hint: recall the differences between the link address and the load address that we have discussed in Lab 1.

`MPBOOTPHYS`宏的目的是计算绝对地址，因为在`boot_aps()`中`memmove()`将`kern/mpentry.S`的代码移动到了`MPENTRY_PADDR`，如果不使用`MPBOOTPHYS`宏，就会寻址到`0xf0000000`之上的地址，而实模式是只能寻址1M。

## Exercist 3
>Exercise 3. Modify mem_init_mp() (in kern/pmap.c) to map per-CPU stacks starting at KSTACKTOP, as shown in inc/memlayout.h. The size of each stack is KSTKSIZE bytes plus KSTKGAP bytes of unmapped guard pages. Your code should pass the new check in check_kern_pgdir().

通过`inc/memlayout.h`里面的内容，去初始化`mem_int_mp`，就是对每个CPU进行内存初始化
```c
static void
mem_init_mp(void)
{
	int i;
	uintptr_t kstack = KSTACKTOP - KSTKSIZE;
	for(i = 0; i < NCPU; i++) {
		boot_map_region(kern_pgdir, kstack, KSTKSIZE, PADDR(percpu_kstacks[i]), PTE_W);
		kstack -= KSTKSIZE + KSTKGAP;
	}
}
```

运行`make qemu-nox`，如果出现以下内容，则说明代码没问题
```bash
check_kern_pgdir() succeeded!
check_page_free_list() succeeded!
check_page_installed_pgdir() succeeded!
```

## Exercise 4
>Exercise 4. The code in trap_init_percpu() (kern/trap.c) initializes the TSS and TSS descriptor for the BSP. It worked in Lab 3, but is incorrect when running on other CPUs. Change the code so that it can work on all CPUs. (Note: your new code should not use the global ts variable any more.)


主要是修改`trap_init_percpu`的代码，其实代码主要都给你写好了，只是要针对当前的cpu去修改代码
```c
void
trap_init_percpu(void)
{
	int cid = thiscpu->cpu_id;

	thiscpu->cpu_ts.ts_esp0 = KSTACKTOP - cid * (KSTKSIZE + KSTKGAP);
	thiscpu->cpu_ts.ts_ss0 = GD_KD;
	thiscpu->cpu_ts.ts_iomb = sizeof(struct Taskstate);

	gdt[(GD_TSS0 >> 3) + cid] = SEG16(STS_T32A, (uint32_t) (&thiscpu->cpu_ts),sizeof(struct Taskstate) - 1, 0);
	gdt[(GD_TSS0 >> 3) + cid].sd_s = 0;

	ltr(GD_TSS0 + 8 * cid);
	lidt(&idt_pd);
}
```

运行`make qemu-nox CPUS=4`，可以看到下面的结果，就是成功
```bash
SMP: CPU 0 found 4 CPU(s)
enabled interrupts: 1 2
SMP: CPU 1 starting
SMP: CPU 2 starting
SMP: CPU 3 starting
```

## Exercise 5
>Exercise 5. Apply the big kernel lock as described above, by calling lock_kernel() and unlock_kernel() at the proper locations.

这个非常简单。就是按照文档里面说的添加就好了
```c
// i386_init
lock_kernel();
boot_aps()

// mp_main
lock_kernel();
sched_yield();

// trap
lock_kernel();
assert(curenv);


// env_run
lcr3(PADDR(e->env_pgdir));
unlock_kernel();
env_pop_tf(&e->env_tf);	
```

## Question
>2.It seems that using the big kernel lock guarantees that only one CPU can run the kernel code at a time. Why do we still need separate kernel stacks for each CPU? Describe a scenario in which using a shared kernel stack will go wrong, even with the protection of the big kernel lock.

当中断发生的时候，在检查锁之前会压入这些信息：
```c
uint32_t tf_err;
uintptr_t tf_eip;
uint16_t tf_cs;
uint16_t tf_padding3;
uint32_t tf_eflags;
```

所以如果不区分CPU使用栈的话，地址会比较混乱，导致无法起到保护作用


## Exercise 6
>Exercise 6. Implement round-robin scheduling in sched_yield() as described above. Don't forget to modify syscall() to dispatch sys_yield().

>Make sure to invoke sched_yield() in mp_main.

>Modify kern/init.c to create three (or more!) environments that all run the program user/yield.c.

>Run make qemu. You should see the environments switch back and forth between each other five times before terminating, like below.

>Test also with several CPUS: make qemu CPUS=2.
```bash
...
Hello, I am environment 00001000.
Hello, I am environment 00001001.
Hello, I am environment 00001002.
Back in environment 00001000, iteration 0.
Back in environment 00001001, iteration 0.
Back in environment 00001002, iteration 0.
Back in environment 00001000, iteration 1.
Back in environment 00001001, iteration 1.
Back in environment 00001002, iteration 1.
...
```

>After the yield programs exit, there will be no runnable environment in the system, the scheduler should invoke the JOS kernel monitor. If any of this does not happen, then fix your code before proceeding.

就是根据文档里面实现`sched_yield`函数，轮询算法还是比较好实现的
```c
void
sched_yield(void)
{
	struct Env *idle;

	idle = curenv;
	size_t idx = idle == NULL ? -1 : ENVX(idle->env_id);

	for(size_t i = 0; i < NENV; i++) {
		idx = (idx + 1 == NENV) ? 0 : idx + 1;
		if(envs[idx].env_status == ENV_RUNNABLE) {
			env_run(&envs[idx]);
			return;
		}
	}
	if(idle && idle->env_status == ENV_RUNNING && idle->env_cpunum == cpunum()) {
		env_run(idle);
		return;
	}

	// sched_halt never returns
	sched_halt();
}
```

同时需要修改`syscall`来增加一个系统调用，就是
```c
case SYS_yield:
    sys_yield();
    ret = 0;
    break;
```

最后，正如练习6里面的内容，需要把`fork`相关的测试注释掉，然后自己写几个测试用例，主要修改是在`i386_init()`
```c
/* 需要注释掉的代码
#if defined(TEST)
	// Don't touch -- used by grading script!
	ENV_CREATE(TEST, ENV_TYPE_USER);
#else
	// Touch all you want.
	ENV_CREATE(user_primes, ENV_TYPE_USER);
#endif // TEST*
*/
// 添加的代码
    ENV_CREATE(user_yield, ENV_TYPE_USER);
    ENV_CREATE(user_yield, ENV_TYPE_USER);
    ENV_CREATE(user_yield, ENV_TYPE_USER);
```

最终运行`make qemu-nox CPUS=2`可以看到运行结果如上面所示就是成功

## Question
>3.In your implementation of env_run() you should have called lcr3(). Before and after the call to lcr3(), your code makes references (at least it should) to the variable e, the argument to env_run. Upon loading the %cr3 register, the addressing context used by the MMU is instantly changed. But a virtual address (namely e) has meaning relative to a given address context--the address context specifies the physical address to which the virtual address maps. Why can the pointer e be dereferenced both before and after the addressing switch?

因为所有环境虚拟内存的内核部分都是相同的

>4.Whenever the kernel switches from one environment to another, it must ensure the old environment's registers are saved so they can be restored properly later. Why? Where does this happen?

在`trap.c`里，`curenv->env_tf = *tf`保存了当前的trap结构


## Exercise 7
>Exercise 7. Implement the system calls described above in kern/syscall.c and make sure syscall() calls them. You will need to use various functions in kern/pmap.c and kern/env.c, particularly envid2env(). For now, whenever you call envid2env(), pass 1 in the checkperm parameter. Be sure you check for any invalid system call arguments, returning -E_INVAL in that case. Test your JOS kernel with user/dumbfork and make sure it works before proceeding.

实现五个函数，并且在`syscall`添加相关的调用
```c
int32_t
syscall(uint32_t syscallno, uint32_t a1, uint32_t a2, uint32_t a3, uint32_t a4, uint32_t a5)
{
    int ret = 0;

	switch (syscallno) {
	// 以下case是新增的，其他就不需要了
	case SYS_exofork:
		ret = sys_exofork();
		break;
	case SYS_env_set_status:
		ret = sys_env_set_status(a1, a2);
		break;
	case SYS_page_alloc:
		ret = sys_page_alloc(a1, (void*)a2, a3);
		break;
	case SYS_page_map:
		ret = sys_page_map(a1, (void*)a2, a3, (void*)a4, a5);
		break;
	case SYS_page_unmap:
		ret = sys_page_unmap(a1, (void*)a2);
		break;
	default:
		return -E_INVAL;
	}
	return ret;
}
```

```c
static envid_t
sys_exofork(void)
{
	struct Env *child_env;
	int result = env_alloc(&child_env, curenv->env_id);
	if(result != 0) return result;
	child_env->env_status = ENV_NOT_RUNNABLE;
	child_env->env_tf = curenv->env_tf;
	child_env->env_tf.tf_regs.reg_eax = 0; // 子环境的返回值
	return child_env->env_id; // 返回子环境的id
}

static int
sys_env_set_status(envid_t envid, int status)
{
	struct Env *e;
	if(envid2env(envid, &e, 1) != 0) {
		return -E_BAD_ENV;
	}
	if(status != ENV_RUNNABLE && status != ENV_NOT_RUNNABLE) {
		return -E_INVAL;
	}
	e->env_status = status;
	return 0;
}

static int
sys_page_alloc(envid_t envid, void *va, int perm)
{
	struct Env *e;
	if(envid2env(envid, &e, 1) != 0) {
		return -E_BAD_ENV;
	}
	if((uintptr_t)va > UTOP || PGOFF(va) != 0 || perm < (PTE_U | PTE_P) || (perm & ~PTE_SYSCALL) != 0) {
		return -E_INVAL;
	}
	struct PageInfo *pp = page_alloc(ALLOC_ZERO);
	if(pp == NULL || page_insert(e->env_pgdir, pp, va, perm) != 0) {
		return -E_NO_MEM;
	}
	return 0;
}

static int
sys_page_map(envid_t srcenvid, void *srcva,
	     envid_t dstenvid, void *dstva, int perm)
{
	struct Env *srcenv, *dstenv;
	if(envid2env(srcenvid, &srcenv, 1) != 0 || envid2env(dstenvid, &dstenv, 1) != 0) {
		return -E_BAD_ENV;
	}
	if((uintptr_t)srcva > UTOP || PGOFF(srcva) != 0 ||(uintptr_t)dstva > UTOP || PGOFF(dstva) || perm < (PTE_U | PTE_P) || (perm & ~PTE_SYSCALL) != 0) {
		return -E_INVAL;
	}

	pte_t *src_pte;
	struct PageInfo *pp = page_lookup(srcenv->env_pgdir, srcva, &src_pte);
	if((perm & PTE_W) && (*src_pte & PTE_W) == 0) {
		return -E_INVAL;
	}
	if(page_insert(dstenv->env_pgdir, pp, dstva, perm) != 0) {
		return -E_NO_MEM;
	}
	return 0;
}

static int
sys_page_unmap(envid_t envid, void *va)
{
	struct Env *e;
	if(envid2env(envid, &e, 1) != 0) {
		return -E_BAD_ENV;
	}
	if((uintptr_t)va >= UTOP || PGOFF(va) != 0) {
		return -E_INVAL;
	}
	page_remove(e->env_pgdir, va);
	return 0;
}
```

同时还有要注释`mp_main`中关于自旋的部分
```c
void
mp_main(void)
{
	// Remove this after you finish Exercise 6
	// for (;;);
}
```

运行`make grade`，可以看到如下信息就是成功
```c
dumbfork: OK (9.5s) 
    (Old jos.out.dumbfork failure log removed)
Part A score: 5/5
```

## Exercise 8
>Exercise 8. Implement the sys_env_set_pgfault_upcall system call. Be sure to enable permission checking when looking up the environment ID of the target environment, since this is a "dangerous" system call.

经历过Exercise 7，那么这个就很简单了,就是把函数保存在环境里面
```c
static int
sys_env_set_pgfault_upcall(envid_t envid, void *func)
{
	struct Env *e;
	if(envid2env(envid, &e, 1) != 0) {
		return -E_BAD_ENV;
	}
	e->env_pgfault_upcall = func;
	return 0;
}
```

同时需要修改系统调用的代码
```c
int32_t
syscall(uint32_t syscallno, uint32_t a1, uint32_t a2, uint32_t a3, uint32_t a4, uint32_t a5)
{
    int ret = 0;

	switch (syscallno) {
	// 以下case是新增的，其他就不需要了
	case SYS_env_set_pgfault_upcall:
		ret = sys_env_set_pgfault_upcall(a1, (void*)a2);
		break;
	default:
		return -E_INVAL;
	}
	return ret;
}
```

## Exercise 9
>Exercise 9. Implement the code in page_fault_handler in kern/trap.c required to dispatch page faults to the user-mode handler. Be sure to take appropriate precautions when writing into the exception stack. (What happens if the user environment runs out of space on the exception stack?)

针对`page_fault_handler`需要处理在`UXSTACKTOP`上面的错误处理。比较简单
```c
void
page_fault_handler(struct Trapframe *tf)
{
	uint32_t fault_va;

	// Read processor's CR2 register to find the faulting address
	fault_va = rcr2();

	// Handle kernel-mode page faults.

	// LAB 3: Your code here.
	if((tf->tf_cs & 0x3) == 0) {
		panic("page fault in kernel");
	}

	// LAB 4: Your code here.
	if(curenv->env_pgfault_upcall != 0) {
		struct UTrapframe *utf;
		size_t sz = sizeof(struct UTrapframe);
		if(tf->tf_esp >= UXSTACKTOP - PGSIZE && tf->tf_esp < UXSTACKTOP) {
			utf = (struct UTrapframe*)(tf->tf_esp - 4 - sz);
		} else {
			utf = (struct UTrapframe*)(UXSTACKTOP - sz);
		}
		user_mem_assert(curenv, (void*)utf, sz, PTE_W);
		utf->utf_eflags = tf->tf_eflags;
        utf->utf_eip = tf->tf_eip;
        utf->utf_err = tf->tf_err;
        utf->utf_esp = tf->tf_esp;
        utf->utf_fault_va = fault_va;
        utf->utf_regs = tf->tf_regs;

		curenv->env_tf.tf_eip = (uintptr_t)curenv->env_pgfault_upcall;
		curenv->env_tf.tf_esp = (uintptr_t)utf;
		env_run(curenv);
	}

	// Destroy the environment that caused the fault.
	cprintf("[%08x] user fault va %08x ip %08x\n",
		curenv->env_id, fault_va, tf->tf_eip);
	print_trapframe(tf);
	env_destroy(curenv);
}
```

## Exercise 10
>Exercise 10. Implement the _pgfault_upcall routine in lib/pfentry.S. The interesting part is returning to the original point in the user code that caused the page fault. You'll return directly there, without going back through the kernel. The hard part is simultaneously switching stacks and re-loading the EIP.

这个地方真不会，其实如果按照里面的汇编语言注视，应该也没问题。但是懒得看了，直接复制github上面的解决方案
```armasm
.text
.globl _pgfault_upcall
_pgfault_upcall:
	// Call the C page fault handler.
	pushl %esp			// function argument: pointer to UTF
	movl _pgfault_handler, %eax
	call *%eax
	addl $4, %esp			// pop function argument
	
	// LAB 4: Your code here.
	movl 0x28(%esp), %edx # trap-time eip
	subl $0x4, 0x30(%esp) # we have to use subl now because we can't use after popfl
	movl 0x30(%esp), %eax # trap-time esp-4
	movl %edx, (%eax)
	addl $0x8, %esp

	// LAB 4: Your code here.
	popal

	// LAB 4: Your code here.
	addl $0x4, %esp #eip
	popfl

	// LAB 4: Your code here.
	popl %esp

	// LAB 4: Your code here.
	ret
```

## Exercise 11
>Exercise 11. Finish set_pgfault_handler() in lib/pgfault.c.

总结来说就是分配内存，然后直接运行就可以了
```c
void
set_pgfault_handler(void (*handler)(struct UTrapframe *utf))
{
	int r;

	if (_pgfault_handler == 0) {
		// LAB 4: Your code here.
		envid_t eid = sys_getenvid();
		r = sys_page_alloc(eid, (void*)(UXSTACKTOP - PGSIZE), PTE_W|PTE_U|PTE_P);
		if (r < 0) {
			panic("set_pgfault_handler: %e", r);
		}
		r = sys_env_set_pgfault_upcall(eid, (void*)_pgfault_handler);
		if(r < 0) {
			panic("set_pgfault_handler: %e", r);	
		}
	}

	// Save handler pointer for assembly to call.
	_pgfault_handler = handler;
}
```

然后运行`make grade`，出现以下结果，则说明上面的都是正确的
```c
faultread: OK (8.4s) 
faultwrite: OK (8.7s) 
faultdie: OK (9.0s) 
faultregs: OK (8.4s) 
faultalloc: OK (8.4s) 
faultallocbad: OK (8.2s) 
faultnostack: OK (8.7s) 
faultbadhandler: OK (8.4s) 
faultevilhandler: OK (8.6s) 
```

期间还出现了一个小插曲，之前在lab3中写的关于`user_mem_check`函数有点问题，导致`faultdie`的结果失败，其实是因为内存检测的边缘条件问题，已经在原来的内容里面修改了


## Exercise 12
>Exercise 12. Implement fork, duppage and pgfault in lib/fork.c.

>Test your code with the forktree program. It should produce the following messages, with interspersed 'new env', 'free env', and 'exiting gracefully' messages. The messages may not appear in this order, and the environment IDs may be different.
```text
	1000: I am ''
	1001: I am '0'
	2000: I am '00'
	2001: I am '000'
	1002: I am '1'
	3000: I am '11'
	3001: I am '10'
	4000: I am '100'
	1003: I am '01'
	5000: I am '010'
	4001: I am '011'
	2002: I am '110'
	1004: I am '001'
	1005: I am '111'
	1006: I am '101'
```

主要是实现`lib/fork.c`中的`fork()`功能，涉及到三个函数
```c
static void
pgfault(struct UTrapframe *utf)
{
	void *addr = (void *) utf->utf_fault_va;
	uint32_t err = utf->utf_err;
	int r;
	
	// LAB 4: Your code here.
	if ((err & FEC_WR) == 0 || (uvpt[PGNUM(addr)] & PTE_COW) == 0) {
		panic("pgfault: check access failed");
	}
	
	// LAB 4: Your code here.
	if((r = sys_page_alloc(0, PFTEMP, PTE_P | PTE_U | PTE_W)) < 0) {
		panic("alloc page failed: %e", r);
	}
	addr = ROUNDDOWN(addr, PGSIZE);
	memcpy(PFTEMP, addr, PGSIZE);
	if((r = sys_page_map(0, PFTEMP, 0, addr, PTE_P | PTE_U | PTE_W)) < 0) {
		panic("sys_page_map: %e", r);
	}
	if((r = sys_page_unmap(0, PFTEMP)) < 0) {
		panic("sys_page_unmap: %e", r);
	}
}

static int
duppage(envid_t envid, unsigned pn)
{
	int r;

	// LAB 4: Your code here.
	void *va = (void *)(pn << PGSHIFT);
	int perm = uvpt[pn] & 0xFFF;
	if ((perm & PTE_W) || (perm & PTE_COW)) {
		perm |= PTE_COW;
		perm &= ~PTE_W;
	}
	perm &= PTE_SYSCALL;
	if((r = sys_page_map(0, va, envid, va, perm)) < 0) {
		panic("sys_page_map: %e", r);
	}
	if((r = sys_page_map(0, va, 0, va, perm)) < 0) {
		panic("sys_page_map: %e", r);
	}
	return r;
}

envid_t
fork(void)
{
	// LAB 4: Your code here.
	set_pgfault_handler(pgfault);
	envid_t envid = sys_exofork();
	if(envid < 0) {
		panic("sys_exofork: %e", envid);
	}
	if(envid == 0) {
		thisenv = &envs[ENVX(sys_getenvid())];
		return 0;
	}
	for(uintptr_t addr = UTEXT; addr < USTACKTOP; addr += PGSIZE) {
		if (uvpd[PDX(addr)] & PTE_P && uvpt[PGNUM(addr)] & PTE_P) {
			duppage(envid, PGNUM(addr));
		}
	}
	int r;
	if((r = sys_page_alloc(envid, (void*)(UXSTACKTOP - PGSIZE), PTE_P | PTE_U | PTE_W)) < 0) {
		panic("sys_page_alloc: %e", r);
	}
	extern void _pgfault_upcall(void);
	if ((r = sys_env_set_pgfault_upcall(envid, _pgfault_upcall)) < 0){
        panic("sys_env_set_pgfault_upcall: %e", r);
	}
    if ((r = sys_env_set_status(envid, ENV_RUNNABLE)) < 0){
        panic("sys_env_set_status: %e", r);
	}
    return envid;
}
```

同时，需要设置`inc/memlayout.h`中的`JOS_USER`，所以在其中加上
```c
#define JOS_USER 1
```

运行`make grade`可以看到`Part B`全部通过则说明上述代码都是正确的

## Exercise 13
>Exercise 13. Modify kern/trapentry.S and kern/trap.c to initialize the appropriate entries in the IDT and provide handlers for IRQs 0 through 15. Then modify the code in env_alloc() in kern/env.c to ensure that user environments are always run with interrupts enabled.

>Also uncomment the sti instruction in sched_halt() so that idle CPUs unmask interrupts.

>The processor never pushes an error code when invoking a hardware interrupt handler. You might want to re-read section 9.2 of the 80386 Reference Manual, or section 5.8 of the IA-32 Intel Architecture Software Developer's Manual, Volume 3, at this time.

>After doing this exercise, if you run your kernel with any test program that runs for a non-trivial length of time (e.g., spin), you should see the kernel print trap frames for hardware interrupts. While interrupts are now enabled in the processor, JOS isn't yet handling them, so you should see it misattribute each interrupt to the currently running user environment and destroy it. Eventually it should run out of environments to destroy and drop into the monitor.

主要是初始化中断，简单

`kern/trapentry.S`里面添加，递增添加16个
```armasm
TRAPHANDLER_NOEC(irq32, IRQ_OFFSET)
```

`kern/trap.c`的`trap_init`添加。递增添加16个
```c
void irq32();
SETGATE(idt[IRQ_OFFSET], 0, GD_KT, irq32, 3);
```

取消`kern/sched.c`的`sched_halt`的代码
```c
"sti\n"
```

最后在`kern/env.c`中的`env_alloc`添加返回的flag
```c
// Enable interrupts while in user mode.
// LAB 4: Your code here.
e->env_tf.tf_eflags |= FL_IF;
```

## Exercise 14
>Exercise 14. Modify the kernel's trap_dispatch() function so that it calls sched_yield() to find and run a different environment whenever a clock interrupt takes place.

>You should now be able to get the user/spin test to work: the parent environment should fork off the child, sys_yield() to it a couple times but in each case regain control of the CPU after one time slice, and finally kill the child environment and terminate gracefully.

这个就很简单了，只要是时钟中断，就强制调用`sched_yield`就可以了
```c
static void
trap_dispatch(struct Trapframe *tf)
{
	// Handle clock interrupts. Don't forget to acknowledge the
	// interrupt using lapic_eoi() before calling the scheduler!
	// LAB 4: Your code here.
	if(tf->tf_trapno == IRQ_OFFSET + IRQ_TIMER) {
		lapic_eoi();
		sched_yield();
		return;
	}
}
```

运行`make grade`，可以有`65/80`就说明是成功的

## Exercise 15
>Exercise 15. Implement sys_ipc_recv and sys_ipc_try_send in kern/syscall.c. Read the comments on both before implementing them, since they have to work together. When you call envid2env in these routines, you should set the checkperm flag to 0, meaning that any environment is allowed to send IPC messages to any other environment, and the kernel does no special permission checking other than verifying that the target envid is valid.

>Then implement the ipc_recv and ipc_send functions in lib/ipc.c.

>Use the user/pingpong and user/primes functions to test your IPC mechanism. user/primes will generate for each prime number a new environment until JOS runs out of environments. You might find it interesting to read user/primes.c to see all the forking and IPC going on behind the scenes.

主要是实现四个函数，都在文档中有写具体的实现。所以按照文档来实现就行了
```c
static int
sys_ipc_try_send(envid_t envid, uint32_t value, void *srcva, unsigned perm)
{
	// LAB 4: Your code here.
	struct Env *e;
	int ret = envid2env(envid, &e, 0);
	if (ret) return ret;//bad env
	if (!e->env_ipc_recving) return -E_IPC_NOT_RECV;
	if (srcva < (void*)UTOP) {
		pte_t *pte;
		struct PageInfo *pg = page_lookup(curenv->env_pgdir, srcva, &pte);
		if (!pg) return -E_INVAL;
		if ((*pte & perm) != perm) return -E_INVAL;
		if ((perm & PTE_W) && !(*pte & PTE_W)) return -E_INVAL;
		if (srcva != ROUNDDOWN(srcva, PGSIZE)) return -E_INVAL;
		if (e->env_ipc_dstva < (void*)UTOP) {
			ret = page_insert(e->env_pgdir, pg, e->env_ipc_dstva, perm);
			if (ret) return ret;
			e->env_ipc_perm = perm;
		}
	}
	e->env_ipc_recving = 0;
	e->env_ipc_from = curenv->env_id;
	e->env_ipc_value = value; 
	e->env_status = ENV_RUNNABLE;
	e->env_tf.tf_regs.reg_eax = 0;
	return 0;
}
```

```c
static int
sys_ipc_recv(void *dstva)
{
	// LAB 4: Your code here.
	if((dstva < (void*)UTOP) && (dstva != ROUNDDOWN(dstva, PGSIZE)) ) return -E_INVAL;

	curenv->env_ipc_recving = 1;
	curenv->env_status = ENV_NOT_RUNNABLE;
	curenv->env_ipc_dstva = dstva;
	sys_yield();
	return 0;
}
```

```c
void
ipc_send(envid_t to_env, uint32_t val, void *pg, int perm)
{
	// LAB 4: Your code here.
	if (!pg) pg = (void*)-1;
	int ret;
	while ((ret = sys_ipc_try_send(to_env, val, pg, perm))) {
		if (ret == 0) break;
		if (ret != -E_IPC_NOT_RECV) panic("not E_IPC_NOT_RECV, %e", ret);
		sys_yield();
	}
}
```

```c
int32_t
ipc_recv(envid_t *from_env_store, void *pg, int *perm_store)
{
	// LAB 4: Your code here.
	if (from_env_store) *from_env_store = 0;
	if (perm_store) *perm_store = 0;
	if (!pg) pg = (void*) -1;
	int ret = sys_ipc_recv(pg);
	if (ret) return ret;
	if (from_env_store)
		*from_env_store = thisenv->env_ipc_from;
	if (perm_store)
		*perm_store = thisenv->env_ipc_perm;
	return thisenv->env_ipc_value;
	return 0;
}
```

最后`make grade`，完成