# Decorator

The Decorator Pattern in Go (Golang) is a `structural design pattern` that allows behavior to be added to individual objects, either statically or dynamically, without affecting the behavior of other objects from the same class. This is typically done by "wrapping" objects with additional functionality.

It works like a wrapper. It envolves the behavior of a function without needing to modify it.

## Key Concepts:
- Component Interface: Defines the basic behavior of the object.
- Concrete Component: The object that needs additional functionality.
- Decorator: Wraps the concrete component and adds extra behavior.

## Use Cases:
- Adding functionality to objects dynamically.
- Enhancing the behavior of an object without modifying its original code.


## Examples
- a middleware is a good example.

## PROS
- Existing services and behaviors are kept intact.
- runtime behaviours
- composition of decotarors
- easy to test
- Decorators can work in isolation

## CONS
- The order of composition matters.
- It's difficult to remove a decorator.
- Not as elegant as a builder pattern.
