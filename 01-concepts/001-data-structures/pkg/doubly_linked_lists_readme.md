# Doubly Linked lists

 Aspect                | Singly Linked List                            | Doubly Linked List                            |
|-----------------------|-----------------------------------------------|-----------------------------------------------|
| **Node Structure**    | Data + Next pointer                           | Data + Next pointer + Previous pointer        |
| **Direction**         | One-way traversal (head to tail)              | Two-way traversal (head to tail and vice versa)|
| **Memory Usage**      | Lower (stores only one pointer per node)      | Higher (stores two pointers per node)         |
| **Insertion/Deletion**| Efficient at beginning (`O(1)` time)          | Efficient at both beginning and end (`O(1)` time) |
| **Insertion/Deletion (Middle)** | `O(n)` time due to traversal        | `O(1)` time if the node is known              |
| **Traversal**         | Only forward                                  | Forward and backward                          |
| **Implementation Complexity** | Simpler (fewer pointers to manage)    | More complex (managing two pointers per node) |
| **Use Cases**         | Simple stacks, queues, or one-way lists       | Deques, navigation systems, LRU caches        |
| **Advantages**        | Less memory, simpler to implement             | Flexible traversal, efficient operations in both directions |
| **Disadvantages**     | Limited to one-way traversal, less flexible   | More memory consumption, more complex implementation |


- Singly Linked List: Simpler, requires less memory, but only supports one-way traversal.
- Doubly Linked List: More complex, requires more memory, but supports efficient two-way traversal and more flexible node insertion and deletion.

- **Note**: Traversal is the process of visiting each node or element in a data structure in a systematic way. In the context of linked lists, traversal involves starting at the head (or the first node) and moving from one node to the next by following the references (pointers) until you reach the end of the list.

