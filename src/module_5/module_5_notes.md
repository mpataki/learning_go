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

Insert and lookup is pretty straight forward:
```go
prefixMap = make(map[string]string)
prefixMap["Bob"] = "Mr."
name := "Bob"
fmt.Println(prefixMap[name], name) // prints "Mr. Bob"
```

We can declare a map with content using the initialization operator:
```go
prefixMap := map[string]string{
    "Bob":  "Mr.",
    "Joe":  "Dr.",
    "Amy":  "Dr.",
    "Mary": "Ms.",
  }
```
Notice the final comma even on the last line. This is required.

Updating a map is also exactly how you would think:
```go
prefixMap["Joe"] = "Jr."
```

Deleting from a map is done as follows:
```go
delete(prefixMap, "Mary")
```
It should be noted that `delete` can be called on a map with a key that doesn't exist as well. Older versions of go don't have the `delete` method and instead use:
```go
prefixMap["Mary"] = "", false
```
But this is no longer the preferred way of deleting from maps.

You can check for key existence in a map as follows:
```go
if value, exists := prefixMap[name]; exists {
  return value
}
```
This might give some insight into how the old way of deleting works.
