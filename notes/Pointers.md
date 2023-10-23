# Pointers

Memory is characterized by two parts

- Address, the location of this memory in the hard-drive
- Value, the contents of the memory

A pointer is simply a memory which contains an address of another memory, hence it is pointing to that original memory.

`&i` gives you the address of the variable i

`*` has two uses,

- if it is present before a type `*int` it is a pointer type where int is the base <br/>
- if it is present before a variable `*p`, the star acts as an operator and returns the value that the pointer is pointing to,<br/>
  here p is the address and \*p is the value present at that address

---

## Why pointers?

``` Go
j=j/37 // is an ineffective assignment 
p := &j
*p = *p/37 // is valid and will perform the expected operation
```

> *Core idea* <br/> It is efficient to use the variable in one place and access it with pointers rather than copying it repeatedly to multiple places.

Go by default adheres to this and doesn't allow you to mutate the value outside the scope the function that the variable is declared in.<br/>
To mutate the data across the function calls we need to use pointers.

> Go passes arguments by value, not by reference, unless you use a pointer.

## Memory Allocation

When we execute some go code, a goroutine is created and each goroutine gets it's own stack of memory.<br/>
A goroutine is an independent path of execution can be treated as a light weight thread.<br/>
When a function call is made a part of the memory stack is allocated to the function which is called a frame.<br/>
The goroutine only works in the current active frame and cannot access memory that's outside the frame. This guaranties immutability.<br/>
To operate on values that are outside of the current function we use pointers.

---

## Heap
