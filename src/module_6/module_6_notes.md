### Module 6: Slices

Slices are built on arrays. An array in go is the same in many ways as an array in C or C++. It is a collection of items that are allocated in contiguous memory addresses that is of a fixed size. In go however, the values are zeroed out on creation. Also in go, an array is not a pointer. This means that if you pass around an array from function to function, you will actually be passing values, not pointers or references as is the case in many other languages.

Slices are generally more useful in go. They are pointers to arrays, which means that although the pointer gets passed by value when used as a parameter, the underlying array does not, making them easier to pass around.

When slices are declared, they are nil to begin with and can't be used. We must call the `make` function on them to initialize. Here is an example using the long form:
```go
  var s []int
  s = make([]int, 3, 10)
  s[0] = 1
  s[1] = 2
  s[2] = 3
```
Here `[]int`, notice the brackets without the number in it, declares the slice. We then call `make` with the slice type, a length, and a capacity. The length is how many items will be in the array to begin with, and the capacity indicates the size of the underlying array. This means that we can `append` to this slice until it is of length 10 before we need to reallocate.

A slice itself is fixed in size, but we can use `append` to make it grow.

We can also declare and initialize a slice using a literal, which makes it much shorter
```go
s2 := []int { 1, 2, 3 }
```

We often want to divide slices up into smaller parts. Here's how we might do that using the slice defined above:
```go
s2 = s[0:1]
```
Here we are assigning `s2` to `0` through `1` indices inclusively. If we wanted to we could also leave off the index before or after the colon. For example:
```go
s2 = s[:1]
```
This would assign everything up to index `1` exclusively.

We often want to append to a slice. Here's how that goes:
```go
s = append(s, 4)
```
If the slice doesn't have the capacity to hold what we are trying to append, it will allocate a new underlying array, copy all of the existing elements over, and finally, copy the new new value as well. We can also append other slices:
```go
s = append(s, s2...)
```
The `...` indicates a variable number of arguments.

We also may want to delete from a slice. This is done as follows:
```go
slice = append(slice[:1], slice[2:]...)
```
This will delete index `1` from the slice.
