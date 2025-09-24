---
author: xnzone 
title: 实验操作
date: 2021-09-10
image: https://s2.loli.net/2025/09/24/B9kM2IQLyPN4DRw.jpg
cover: false 
weight: 1232
tags: ["xv6", "os", "protect"]
---

## Exercise 1
>Exercise 1. Modify mem_init() in kern/pmap.c to allocate and map the envs array. This array consists of exactly NENV instances of the Env structure allocated much like how you allocated the pages array. Also like the pages array, the memory backing envs should also be mapped user read-only at UENVS (defined in inc/memlayout.h) so user processes can read from this array.    
>You should run your code and make sure check_kern_pgdir() succeeds.

分配地址给`envs`，所以代码也是类似的， 在`mem_init`中填充如下代码
{{< highlight c >}}
// Make 'envs' point to an array of size 'NENV' of 'struct Env'.
// LAB 3: Your code here.
envs = (struct Env*)boot_alloc(sizeof(struct Env) * NENV);

// LAB 3: Your code here.
boot_map_region(kern_pgdir, UENVS, NENV * sizeof(struct Env), PADDR(envs), PTE_U);
{{< /highlight  >}}

编译运行后出现，`check_kern_pgdir() succeeded!`就说明是成功的

## Exercise 2
>Exercise 2. In the file env.c, finish coding the following functions:

- env_init(): Initialize all of the Env structures in the envs array and add them to the env_free_list. Also calls env_init_percpu, which configures the segmentation hardware with separate segments for privilege level 0 (kernel) and privilege level 3 (user).
- env_setup_vm(): Allocate a page directory for a new environment and initialize the kernel portion of the new environment's address space.
- region_alloc(): Allocates and maps physical memory for an environment
- load_icode(): You will need to parse an ELF binary image, much like the boot loader already does, and load its contents into the user address space of a new environment.
- env_create(): Allocate an environment with env_alloc and call load_icode to load an ELF binary into it.
- env_run(): Start a given environment running in user mode.

>As you write these functions, you might find the new cprintf verb %e useful -- it prints a description corresponding to an error code. For example,
{{< highlight c >}}
	r = -E_NO_MEM;
	panic("env_alloc: %e", r);
{{< /highlight  >}}

>will panic with the message "env_alloc: out of memory".

`env_init`初始化环境，主要是把所有`env`都放到`env_free_list`里面，然后更新`env_free_list`的值，注意从后往前遍历
{{< highlight c >}}
void
env_init(void)
{
	// Set up envs array
	// LAB 3: Your code here.
	env_free_list = NULL;
	int i;
	for(i = NENV - 1; i >= 0; i--) {
		envs[i].env_id = 0;
		envs[i].env_link = env_free_list;
		env_free_list = &envs[i];
	}
	// Per-CPU part of the initialization
	env_init_percpu();
}
{{< /highlight  >}}

`env_setup_vm`设置`env_pgdir`的值, 大部分的代码已经是有的，只是需要初始化`env_pgdir`，使用`page2kva`函数就可以了
{{< highlight c >}}
static int
env_setup_vm(struct Env *e)
{
	int i;
	struct PageInfo *p = NULL;

	if (!(p = page_alloc(ALLOC_ZERO)))
		return -E_NO_MEM;

	// LAB 3: Your code here.
	p->pp_ref++;
	e->env_pgdir = (pte_t *) page2kva(p);
	memcpy(e->env_pgdir, kern_pgdir, PGSIZE);

	// UVPT maps the env's own page table read-only.
	// Permissions: kernel R, user R
	e->env_pgdir[PDX(UVPT)] = PADDR(e->env_pgdir) | PTE_P | PTE_U;

	return 0;
}
{{< /highlight  >}}

`region_alloc` 分配和映射物理内存，需要注意的是，分配内存要按页大小分配，所以需要用`ROUNDDOWN`和`ROUNDUP`来设置初始地址和结束地址，然后按页大小分配和映射，然后插入到`e->pgdir`的目录下面就可以了
{{< highlight c >}}
static void
region_alloc(struct Env *e, void *va, size_t len)
{
	// LAB 3: Your code here.
	struct PageInfo *pp = NULL;
	uintptr_t start = ROUNDDOWN((uintptr_t)va, PGSIZE);
	uintptr_t end = ROUNDUP((uintptr_t)va + len, PGSIZE);
	for(; start < end; start += PGSIZE) {
		struct PageInfo *p = page_alloc(0);
		if(!p) {
			panic("region alloc failed");
		}
		page_insert(e->env_pgdir, p, (void*)start, PTE_W | PTE_U);
	}
}
{{< /highlight  >}}

`load_icode`是加载ELF文件，所以大部分代码可以参考`boot/main.c`的代码，只是需要处理一下加载的位置就可以了
{{< highlight c >}}
static void
load_icode(struct Env *e, uint8_t *binary)
{
	// LAB 3: Your code here.
	struct Elf *elf = (struct Elf*)binary;
	struct Proghdr *ph, *eph;
	if(elf->e_magic != ELF_MAGIC) {
		panic("not excutable");
	}
	
	ph = (struct Proghdr *)((uint8_t *) elf + elf->e_phoff);
	eph = ph + elf->e_phnum;

	lcr3(PADDR(e->env_pgdir));

	for(; ph < eph; ph++) {
		if(ph->p_type == ELF_PROG_LOAD) {
			region_alloc(e, (void *)ph->p_va, ph->p_memsz);
			memset((void *)ph->p_va, 0, ph->p_memsz);
			memcpy((void *)ph->p_va, binary + ph->p_offset, ph->p_filesz);
		}
	}
	e->env_tf.tf_eip = elf->e_entry;

	// Now map one page for the program's initial stack
	// at virtual address USTACKTOP - PGSIZE.

	// LAB 3: Your code here.
	region_alloc(e, (void *)(USTACKTOP- PGSIZE), PGSIZE);
	lcr3(PADDR(kern_pgdir));
}
{{< /highlight  >}}

`env_create` 分配一个内存，然后调用`load_icode`加载ELF文件
{{< highlight c >}}
void
env_create(uint8_t *binary, enum EnvType type)
{
	// LAB 3: Your code here.
	struct Env *penv;
	int r = env_alloc(&penv, 0);
	if(r < 0) {
		panic("env_create %e", r);
	}
	penv->env_type = type;
	load_icode(penv, binary);
}
{{< /highlight  >}}

`env_run` 就是运行，把当前的`env`停掉，然后把`env`设置成运行的。提示中也有要调用`env_pop_tf`来恢复寄存器
{{< highlight c >}}
void
env_run(struct Env *e)
{
	// LAB 3: Your code here.
	if(curenv && curenv->env_status == ENV_RUNNING) {
		curenv->env_status = ENV_RUNNABLE;
	}
	curenv = e;
	e->env_status = ENV_RUNNING;
	e->env_runs++;
	lcr3(PADDR(e->env_pgdir));
	env_pop_tf(&e->env_tf);	

	// panic("env_run not yet implemented");
}
{{< /highlight  >}}

编译运行，如果出现`Triple fault.`则说明是成功的

下面是调用步骤，确保你理解了这个过程
- start(kern/entry.S)
- i386_init(kern/init.c)
    - cons_init
    - mem_init
    - env_init
    - trap_init(现在还没有实现)
    - env_create
    - env_run
        - env_pop_tf

## Exercise 3
>Exercise 3. Read Chapter 9, Exceptions and Interrupts in the 80386 Programmer's Manual (or Chapter 5 of the IA-32 Developer's Manual), if you haven't already.

练习3主要是熟悉异常和中断机制，可以阅读[80386 Programmer's Manual](https://pdos.csail.mit.edu/6.828/2018/readings/i386/toc.htm)的第九章或者[IA-32 Developer's Manual](https://pdos.csail.mit.edu/6.828/2018/readings/ia32/IA32-3A.pdf)的第五章

## Exercise 4
>Exercise 4. Edit trapentry.S and trap.c and implement the features described above. The macros TRAPHANDLER and TRAPHANDLER_NOEC in trapentry.S should help you, as well as the T_* defines in inc/trap.h. You will need to add an entry point in trapentry.S (using those macros) for each trap defined in inc/trap.h, and you'll have to provide _alltraps which the TRAPHANDLER macros refer to. You will also need to modify trap_init() to initialize the idt to point to each of these entry points defined in trapentry.S; the SETGATE macro will be helpful here.

>Your _alltraps should:

1. push values to make the stack look like a struct Trapframe
2. load GD_KD into %ds and %es
3. pushl %esp to pass a pointer to the Trapframe as an argument to trap()
4. call trap (can trap ever return?)

>Consider using the pushal instruction; it fits nicely with the layout of the struct Trapframe.

>Test your trap handling code using some of the test programs in the user directory that cause exceptions before making any system calls, such as user/divzero. You should be able to get make grade to succeed on the divzero, softint, and badsegment tests at this point.

就是用汇编注册一下IDT的代码

`kern/trapentry.S`
{{< highlight asm >}}
/*
 * Lab 3: Your code here for generating entry points for the different traps.
 */
TRAPHANDLER_NOEC(th0, 0)
TRAPHANDLER_NOEC(th1, 1)
TRAPHANDLER_NOEC(th2, 2)
TRAPHANDLER_NOEC(th3, 3)
TRAPHANDLER_NOEC(th4, 4)
TRAPHANDLER_NOEC(th5, 5)
TRAPHANDLER_NOEC(th6, 6)
TRAPHANDLER_NOEC(th7, 7)
TRAPHANDLER(th8, 8)
TRAPHANDLER(th10, 10)
TRAPHANDLER(th11, 11)
TRAPHANDLER(th12, 12)
TRAPHANDLER(th13, 13)
TRAPHANDLER(th14, 14)
TRAPHANDLER_NOEC(th16, 16)
{{< /highlight  >}}

{{< highlight asm >}}
/*
 * Lab 3: Your code here for _alltraps
 */

.global _alltraps
_alltraps:
    pushl %ds 
    pushl %es
    pushal

    movw $GD_KD, %ax
    movw %ax, %ds
    movw %ax, %es

    pushl %esp
    call trap
{{< /highlight  >}}

`kern/trap.c`
{{< highlight c >}}
void
trap_init(void)
{
	extern struct Segdesc gdt[];

	// LAB 3: Your code here.
	void th0();
	void th1();
	void th2();
	void th3();
	void th4();
	void th5();
	void th6();
	void th7();
	void th8();
	void th10();
	void th11();
	void th12();
	void th13();
	void th14();
	void th16();


	SETGATE(idt[0], 0, GD_KT, th0, 0);
	SETGATE(idt[1], 0, GD_KT, th1, 0);
	SETGATE(idt[2], 0, GD_KT, th2, 0);
	SETGATE(idt[3], 0, GD_KT, th3, 0);
	SETGATE(idt[4], 0, GD_KT, th4, 0);
	SETGATE(idt[5], 0, GD_KT, th5, 0);
	SETGATE(idt[6], 0, GD_KT, th6, 0);
	SETGATE(idt[7], 0, GD_KT, th7, 0);
	SETGATE(idt[8], 0, GD_KT, th8, 0);
	SETGATE(idt[10], 0, GD_KT, th10, 0);
	SETGATE(idt[11], 0, GD_KT, th11, 0);
	SETGATE(idt[12], 0, GD_KT, th12, 0);
	SETGATE(idt[13], 0, GD_KT, th13, 0);
	SETGATE(idt[14], 0, GD_KT, th14, 0);
	SETGATE(idt[16], 0, GD_KT, th16, 0);

	// Per-CPU setup 
	trap_init_percpu();
}
{{< /highlight  >}}

最后运行`make grade`看到如下结果就是成功了
{{< highlight bash >}}
divzero: OK (7.1s) 
    (Old jos.out.divzero failure log removed)
softint: OK (7.1s) 
    (Old jos.out.softint failure log removed)
badsegment: OK (6.6s) 
    (Old jos.out.badsegment failure log removed)
Part A score: 30/30
{{< /highlight  >}}

## Question
>1.What is the purpose of having an individual handler function for each exception/interrupt? (i.e., if all exceptions/interrupts were delivered to the same handler, what feature that exists in the current implementation could not be provided?)

给每个中断或异常提供一个处理函数是更好的隔离和保护

>2.Did you have to do anything to make the user/softint program behave correctly? The grade script expects it to produce a general protection fault (trap 13), but softint's code says int $14. Why should this produce interrupt vector 13? What happens if the kernel actually allows softint's int $14 instruction to invoke the kernel's page fault handler (which is interrupt vector 14)?

因为如果系统运行在用户态，权限级别为 3，而 INT 指令是系统指令，权限级别为 0，因此会首先引发 Gerneral Protection Excepetion（即 trap 13）。由 SETGATE 函数定义上方注释可知，通过改变参数 dpl 可以改变调用该 interrupt 需要的权限等级。通过把原来 dpl = 0 的改成 dpl = 3，就可以让用户态程序也可以调用。

## Exercise 5
>Exercise 5. Modify trap_dispatch() to dispatch page fault exceptions to page_fault_handler(). You should now be able to get make grade to succeed on the faultread, faultreadkernel, faultwrite, and faultwritekernel tests. If any of them don't work, figure out why and fix them. Remember that you can boot JOS into a particular user program using make run-x or make run-x-nox. For instance, make run-hello-nox runs the hello user program.

很简单，就是看错误代码是不是`T_PGFLT`，如果是的话，调用`page_fault_handler`就可以了，
{{< highlight c >}}
static void
trap_dispatch(struct Trapframe *tf)
{
	// Handle processor exceptions.
	// LAB 3: Your code here.
	if(tf != NULL && tf->tf_trapno == T_PGFLT) {
		page_fault_handler(tf);
		return;
	}

	// Unexpected trap: The user process or the kernel has a bug.
	print_trapframe(tf);
	if (tf->tf_cs == GD_KT)
		panic("unhandled trap in kernel");
	else {
		env_destroy(curenv);
		return;
	}
}
{{< /highlight  >}}

最后使用`make grade`，如果有以下输出ok，就是成功的
{{< highlight bash >}}
faultread: OK (7.3s) 
    (Old jos.out.faultread failure log removed)
faultreadkernel: OK (6.9s) 
    (Old jos.out.faultreadkernel failure log removed)
faultwrite: OK (6.9s) 
    (Old jos.out.faultwrite failure log removed)
faultwritekernel: OK (7.5s) 
{{< /highlight  >}}

## Exercise 6
>Exercise 6. Modify trap_dispatch() to make breakpoint exceptions invoke the kernel monitor. You should now be able to get make grade to succeed on the breakpoint test.

跟练习5类似
{{< highlight c >}}
static void
trap_dispatch(struct Trapframe *tf)
{
	// Handle processor exceptions.
	// LAB 3: Your code here.
	if(tf != NULL && tf->tf_trapno == T_PGFLT) {
		page_fault_handler(tf);
		return;
	}
	if(tf != NULL && tf->tf_trapno == T_BRKPT) {
		monitor(tf);
		return;
	}

	// Unexpected trap: The user process or the kernel has a bug.
	print_trapframe(tf);
	if (tf->tf_cs == GD_KT)
		panic("unhandled trap in kernel");
	else {
		env_destroy(curenv);
		return;
	}
}
{{< /highlight  >}}

同时还需要把之前初始化的权限改成3，即`trap_init`函数里面。在用户模式进行int 3进入更高特权等级的内核，要求CPL<=DPL
{{< highlight c >}}
SETGATE(idt[3], 0, GD_KT, th3, 3);
{{< /highlight  >}}

最后调用`make grade`，如果出现以下内容，就是说明成功
{{< highlight bash >}}
breakpoint: OK (6.3s) 
    (Old jos.out.breakpoint failure log removed)
{{< /highlight  >}}

## Question 
>3.The break point test case will either generate a break point exception or a general protection fault depending on how you initialized the break point entry in the IDT (i.e., your call to SETGATE from trap_init). Why? How do you need to set it up in order to get the breakpoint exception to work as specified above and what incorrect setup would cause it to trigger a general protection fault?

设置DPL可以启用用户模式下的调用


>4.What do you think is the point of these mechanisms, particularly in light of what the user/softint test program does? 

主要是保护和隔离

## Exercise 7
>Exercise 7. Add a handler in the kernel for interrupt vector T_SYSCALL. You will have to edit kern/trapentry.S and kern/trap.c's trap_init(). You also need to change trap_dispatch() to handle the system call interrupt by calling syscall() (defined in kern/syscall.c) with the appropriate arguments, and then arranging for the return value to be passed back to the user process in %eax. Finally, you need to implement syscall() in kern/syscall.c. Make sure syscall() returns -E_INVAL if the system call number is invalid. You should read and understand lib/syscall.c (especially the inline assembly routine) in order to confirm your understanding of the system call interface. Handle all the system calls listed in inc/syscall.h by invoking the corresponding kernel function for each call.

>Run the user/hello program under your kernel (make run-hello). It should print "hello, world" on the console and then cause a page fault in user mode. If this does not happen, it probably means your system call handler isn't quite right. You should also now be able to get make grade to succeed on the testbss test.

主要是利用前面文档说的，针对系统调用进行代码编写，需要改的几个地方是`kern/trapentry.S`,`kern/trap.c`的`trap_init()`,`trap_dispatch()`和`kern/syscall.c`的`syscall()`

{{< highlight asm >}}
TRAPHANDLER_NOEC(th48, 48)
{{< /highlight  >}}

{{< highlight c >}}
void
trap_init(void)
{
    // 添加的内容
    void th48();
    SETGATE(idt[48], 0, GD_KT, th48, 3);
}
{{< /highlight  >}}

{{< highlight c >}}
static void
trap_dispatch(struct Trapframe *tf)
{
    // 添加的内容
    if(tf != NULL && tf->tf_trapno == T_SYSCALL) {
		tf->tf_regs.reg_eax = syscall(tf->tf_regs.reg_eax, tf->tf_regs.reg_edx, tf->tf_regs.reg_ecx,tf->tf_regs.reg_ebx, tf->tf_regs.reg_edi, tf->tf_regs.reg_esi);
		return;
	}
}
{{< /highlight  >}}

{{< highlight c >}}
int32_t
syscall(uint32_t syscallno, uint32_t a1, uint32_t a2, uint32_t a3, uint32_t a4, uint32_t a5)
{
	// Call the function corresponding to the 'syscallno' parameter.
	// Return any appropriate return value.
	// LAB 3: Your code here.

	int ret = 0;

	switch (syscallno) {
	case SYS_cputs:
		sys_cputs((char*)a1, a2);
		ret = 0;
		break;
	case SYS_cgetc:
		ret = sys_cgetc();
		break;
	case SYS_getenvid:
		ret = sys_getenvid();
		break;
	case SYS_env_destroy:
		sys_env_destroy(a1);
		ret = 0;
		break;
	default:
		return -E_INVAL;
	}
	return ret;
}
{{< /highlight  >}}

运行`make grade`如果出现以下内容，则说明成功
{{< highlight bash >}}
testbss: OK (6.9s) 
    (Old jos.out.testbss failure log removed)
{{< /highlight  >}}

## Exercise 8
>Exercise 8. Add the required code to the user library, then boot your kernel. You should see user/hello print "hello, world" and then print "i am environment 00001000". user/hello then attempts to "exit" by calling sys_env_destroy() (see lib/libmain.c and lib/exit.c). Since the kernel currently only supports one user environment, it should report that it has destroyed the only environment and then drop into the kernel monitor. You should be able to get make grade to succeed on the hello test.

主要修改`libmain`的代码，把`thisenv`赋值就可以了

{{< highlight c >}}
void
libmain(int argc, char **argv)
{
	// set thisenv to point at our Env structure in envs[].
	// LAB 3: Your code here.
	thisenv = 0;
	envid_t env_id = sys_getenvid();
	for(int i = 0; i < NENV; i++) {
		if(envs[i].env_id == env_id) {
			thisenv = &envs[i];
		}
	}
}
{{< /highlight  >}}

运行`make grade`，出现下面的结果，就是成功
{{< highlight bash >}}
hello: OK (6.6s) 
    (Old jos.out.hello failure log removed)
{{< /highlight  >}}

## Exercise 9
>Exercise 9. Change kern/trap.c to panic if a page fault happens in kernel mode.

>Hint: to determine whether a fault happened in user mode or in kernel mode, check the low bits of the tf_cs.

>Read user_mem_assert in kern/pmap.c and implement user_mem_check in that same file.

>Change kern/syscall.c to sanity check arguments to system calls.

>Boot your kernel, running user/buggyhello. The environment should be destroyed, and the kernel should not panic. You should see:

{{< highlight text >}}
	[00001000] user_mem_check assertion failure for va 00000001
	[00001000] free env 00001000
	Destroyed the only environment - nothing more to do!
{{< /highlight  >}}	

>Finally, change debuginfo_eip in kern/kdebug.c to call user_mem_check on usd, stabs, and stabstr. If you now run user/breakpoint, you should be able to run backtrace from the kernel monitor and see the backtrace traverse into lib/libmain.c before the kernel panics with a page fault. What causes this page fault? You don't need to fix it, but you should understand why it happens.

需要修改`user_member_check`
{{< highlight c >}}
int
user_mem_check(struct Env *env, const void *va, size_t len, int perm)
{
	// LAB 3: Your code here.
	uint32_t start = (uint32_t)ROUNDDOWN(va, PGSIZE);
	uint32_t end = (uint32_t)ROUNDDOWN(va + len - 1, PGSIZE);

	for(uint32_t cur = start; ; cur += PGSIZE) {
		pte_t *pte = pgdir_walk(env->env_pgdir, (void*)cur, 0);
		if (pte == NULL || (*pte & perm) != perm || cur >= ULIM) {
            user_mem_check_addr = cur == start ? (uintptr_t)va : cur;
            return -E_FAULT;
        }
		if(cur == end) {
			return 0;
		}
	}
	return 0;
}
{{< /highlight  >}}

然后在`kdebug.c`文件里面添加校验
{{< highlight c >}}
int
debuginfo_eip(uintptr_t addr, struct Eipdebuginfo *info)
{
    // LAB 3: Your code here.
    if (user_mem_check(curenv, usd, sizeof(struct UserStabData), PTE_U))
        return -1;
        
    // LAB 3: Your code here.
    if (user_mem_check(curenv, stabs, sizeof(struct Stab), PTE_U))
        return -1;

    // Your code here.
    if (user_mem_check(curenv, stabstr, stabstr_end-stabstr, PTE_U))
        return -1;
}
{{< /highlight  >}}

修改`page_fault_handler`
{{< highlight c >}}
void
page_fault_handler(struct Trapframe *tf)
{
    // LAB 3: Your code here.
	if((tf->tf_cs & 0x3) == 0) {
		panic("page fault in kernel");
	}
}
{{< /highlight  >}}

最后还要校验`sys_cputs`
{{< highlight c >}}
static void
sys_cputs(const char *s, size_t len)
{
    // LAB 3: Your code here.
	user_mem_assert(curenv, s, len, PTE_U);
}
{{< /highlight  >}}

运行`make grade`，有以下输出就结束了

{{< highlight bash >}}
buggyhello: OK (7.8s) 
    (Old jos.out.buggyhello failure log removed)
buggyhello2: OK (9.1s) 
    (Old jos.out.buggyhello2 failure log removed)
evilhello: OK (7.0s) 
    (Old jos.out.evilhello failure log removed)
Part B score: 50/50

Score: 80/80
{{< /highlight  >}}