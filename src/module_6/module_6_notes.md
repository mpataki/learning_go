### Module 6: Slices

Slices are built on arrays. An array in go is the same in many ways as an array in C or C++. It is a collection of items that are allocated in contiguous memory addresses that is of a fixed size. In go however, the values are zeroed out on creation. Also in go, an array is not a pointer. This means that if you pass around an array from function to function, you will actually be passing values, not pointers or references as is the case in many other languages.

Slices are generally more useful in go. They are pointers to arrays, which means that although the pointer gets passed by value when used as a parameter, the underlying array does not, making them easier to pass around.

When slices are declared, they are nil to begin with and can't be used. We must call the `make` function on them to initialize.

A slice itself is fixed in size, but we can use `append` to make it grow.
