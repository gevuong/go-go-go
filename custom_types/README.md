### What you need to know about custom type declarations
- A new custom type has an underlying type that you want to base your custom type on. For example, in `type Title string`, `Title` is a custom type that is based on an underlying type `string`. 
- Underlying types determine what intrinsic operations a custom type supports.
`type Minutes int` will support all the operations int does: addition, subtraction, etc.
`type Title string` will support the operations string does, like concatenation of other strings.

### Methods 
- Methods are functions with a special receiver argument before the function name. These methods allow you to define new behaviors for values of a type.
- No `this` or `self`, just the value of the type used to invoke a method