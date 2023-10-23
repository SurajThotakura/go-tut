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

```Go
j=j/37 // is an ineffective assignment
p := &j
*p = *p/37 // is valid and will perform the expected operation
```

> _Core idea_ <br/> It is efficient to use the variable in one place and access it with pointers rather than copying it repeatedly to multiple places.

Go by default adheres to this and doesn't allow you to mutate the value outside the scope the function that the variable is declared in.<br/>
To mutate the data across the function calls we need to use pointers.

> Go passes arguments by value, not by reference, unless you use a pointer.

## Memory Allocation

When we execute some go code, a goroutine is created and each goroutine gets it's own stack of memory.<br/>
A goroutine is an independent path of execution can be treated as a light weight thread.<br/>
When a function call is made a part of the memory stack is allocated to the function which is called a frame.<br/>
The goroutine only works in the current active frame and cannot access memory that's outside the frame. This guaranties immutability.<br/>
To operate on values that are outside of the current function we use pointers.
<br/>
As we move on from our current frame to the next the data will be overwritten so that we won't be using garbage values by accident.<br/>

---

## Heap

Heaps come into play when when we are returning a pointer at the end of the function.<br/>
In this case we can't overwrite all the data that this function is using after the frame moves in the stack. <br/>
Hence the compiler will analyze that the returned pointer cannot be overwritten and put's that data into a "Heap"<br/>
![Heap Diagram]("../assets/HeapDiag.png")

> Heap allocation is a burden for the garbage collector and cost us performance. <br/>
 Memory stack doesn't need garbage collector as it is self-cleaning.

---

## When to use pointers

- To update the state, go is primarily designed for the procedural programming paradigm like "C"
- To optimize the memory for large structures (objects) that are being used repeatedly

> It is better to use pointers conservatively as they can lead to errors and unpredictable code in a large project
