## What you need to know about pointers:
- Function arguments are pass-by-value, which means that arguments passed to a fcn are only copies of the value. This means that if fcn need to modify value that is being passed in, fcn will only alter the copy, not the original.
- The meaning and use of "pointers" and "memory address" are interchangeable.
- Pointers point to a value in memory.
- Typing `&` before a variable name gives the address of that variable's value. `&` can be read aloud as "address of". (i.e. `&myNumber` => `address of myNumber`)
- Place a `*` before a pointer to get the value at that pointer. To indicate a `pointer type`, use a `*` character, followed by the type of value it will point to. (i.e. `var x *float64`, or `func hello(pointer *pointer) {...}` when specifying argument).

### When would you want to use pointers?
1. When a fcn needs to modify the original value of the argument you're passing in.
2. When the value is large, it becomes inefficient to make a copy of said value.

## What you need to know about arrays/slices:
- Slices are used more often than arrays.
- Unlike slices, arrays are fixed length, cannot grow or shrink. 
- Slices are built on top of arrays, so it's important to understand arrays as well. Slices don't actually hold any data themselves. They're just a sort of window into the contents of the underlying array.
- `append` create a new array, copying all the elements from the original array and appending additional elements. `append` returns the new slice that points to the larger array.
- When capacity (memory allocation) of the underlying array is reached, appending an additional element will create a new array that is double the capacity of the original underlying array. This is because memory allocation is a slow process, which takes O(n) time. by doubling the capacity each time the current capacity is reached, amortized time for memory allocation becomes O(1).

### What you need to know about maps:
- You can use any type you want as a key, as long as the keys are all the same. All values must be of same type as well. 

### Value Types vs Reference Types 
- **Value types: int, float, string, boolean, struct. If you pass these types into a fcn, you want to use pointers to change these values in a fcn.**
- With value types in Go, we need to worry about pointers if we want to pass a value to a fcn and want to modify the original value inside that fcn. With reference types, we do not need to worry about pointers. 
- **Reference types: slices, maps, channels, pointers, functions. Reference types reference to another data structure in memory. You don't need to worry about pointers with these types because the copy Go makes is always going to be pointing back to the same underyling true source of data.**