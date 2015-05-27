### Module 6: Slices

Slices are build on arrays. An array in go is the same in many ways as an array in C or C++. It is a collection of items that are allocated in contiguous memory addresses that is of a fixed size. In go however, the values are zeroed out on creation. Also, in go, an array is not a pointer. This means that if you pass around an array from function to function, you will actually be passing values, not pointers or references as is the case in many other languages.
