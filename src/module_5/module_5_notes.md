### Module 5: Maps

A couple of points to start us off:
1.  You can make almost any variable a key in a map. The only constraint is that it must have the `==` operator defined on it.
2.  Maps are reference types. This means that when you pass it around, it is passed by reference.
3.  Maps are not thread safe.

This is how we declare a map:
```go
  var someMapName map[string]string
```
The type int the brackets `[]`, in this case `[string]` is our key type, and the type listed after is our value type. You can see that this closely follows the existing conventions that we've seen so far, where `map[string]string` is sort of an extended type. However, one key difference here is that we can't begin using the map quite yet. To begin with , `someMapName` is `nil`. We need to call a method `make` on it in order to initialize, or make the map, like so:
```go
someMapName = make(map[string]string)
```
