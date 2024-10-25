# Linked lists

HEAD -> Node A -> Node B -> Node C -> nil

A singly linked list is a linear data structure consisting of a sequence of elements, where each element, called a node, contains data and a reference (or link) to the next node in the sequence. Unlike arrays, *linked lists do not require contiguous memory allocation, and their size can grow or shrink dynamically*.

## Characteristics:
- Dynamic Size: Can grow or shrink as elements are added or removed.
- Memory Efficiency: Uses memory more efficiently than arrays for frequent insertions and deletions, especially in the middle of the list.
- Sequential Access: Accessing elements requires traversing from the head to the desired node, making it less efficient for random access compared to arrays.

## Applications:
- Implementing Stacks and Queues: Can be used to implement these data structures efficiently.
- Undo Mechanisms: Used in applications where operations need to be undone in sequence.
- Dynamic Memory Allocation: Useful when the number of elements is unknown and changes frequently.