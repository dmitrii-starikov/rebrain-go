## Useful links:
- [The Laws of Reflection](https://go.dev/blog/laws-of-reflection)
- [Reflection in Go](https://vporoshok.me/post/2019/01/reflection/)

## Pros and Cons of Reflection
Pros of reflection include the ability to separate certain logic into separate aspects without
duplicating or complicating the code. At the same time, of course, complexity does not disappear
but moves to runtime. This leads to the cons of reflection: operating with untyped objects at 
execution time. For reliable work with reflection, tests are simply necessary — they will check 
the correctness of applying aspects directly to the business logic code, as well as good coverage 
of the aspect itself. This is similar to how in dynamically typed languages it's necessary to 
cover all possible combinations of input parameter types with tests. Of course, covering absolutely 
all possible combinations doesn't make sense, but checking the normal case, edge cases, and 
exceptions will be useful for the aspect, while for the business logic code, a straightforward
scenario will suffice.

Additionally, the cons of reflection include a decrease in performance. This is related to 
additional memory allocation. As mentioned above, interfaces represent tuples of type and value
— thus, as soon as we cast a value to an interface, memory is allocated for this pair. 
Furthermore, the reflection objects themselves — reflect.Type and reflect.Value — also require 
memory. But in practice, for a known number of objects, this is negligible, especially compared 
to network latencies. Of course, if you need to write a real-time application, reflection won't 
be the best choice, but in most cases, using reflection is justified.

- [Go Data Structures: Interfaces](https://research.swtch.com/interfaces)
- [The Laws of Reflection](https://blog.golang.org/laws-of-reflection)
- [Godoc reflect](https://godoc.org/reflect#StructField)