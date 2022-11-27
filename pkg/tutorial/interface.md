- When implementing an interface, if you use a value type, the methods that implement the interface _have_ to all have value receivers. If you're implementing the interface with a pointer, then the methods just need to exist, regardless of the
  receiver type.
- The "method set" for a value type is the set of all methods that have value receivers.
- The "method set" for a pointer type is all of the methods with value receivers AND all of the methods with pointer receivers.

Need the underlying data? Pointers.

---

# BEST PRACTICES when working with interfaces

- Use many, small interfaces.

  - If you need to use a large interface, create it out of small interfaces.
  - The smaller your interfaces, the more useful and powerful they're going to be.
  - Single method interfaces, LFG 🔥

- Don't export interfaces for types that will be consumed.

  - If you don't have a _reason_ to export it, _don't_.
    Perfectly acceptable to export the concrete type. (the struct ...?)
  - **By NOT exporting an interface, it allows the _consumer_ of that struct to define their own interface.**

- Do export interfaces for types that will be used by the package

  - Export interfaces for types that you will be using.
  - If you're gonna pull a value in, go ahead and accept an interface instead of a concrete type, if at all possible.

- ^ Opposite of other languages because of Go's implicit implementation of interfaces.
- Define ƒns and methods to receive interface types instead of concrete types whenever possible
