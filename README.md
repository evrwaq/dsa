# Data Structures and Algorithms (DSA)

*[Leia em Português](README.pt-br.md)*

This project aims to demonstrate the main data structures and algorithms, as well as document and explain each one.

## Data Structures

### Array

An Array is a data structure that stores a collection of elements in contiguous memory locations. Each element can be directly accessed by its index, making access operations very fast. However, insertion and removal of elements can be time-consuming as they may require shifting other elements.

**Characteristics:**
- **Indexing**: Direct access to elements via index.
- **Fixed size**: The size of the array is defined at creation and cannot be changed.
- **Homogeneous type**: All elements in the array are of the same type.

**Main operations and their complexities:**
- **Access**:
  - Best case: O(1) - constant time.
  - Worst case: O(1) - constant time.
- **Insertion**:
  - Best case: O(1) - constant time (insertion at the end if there is space).
  - Worst case: O(n) - linear time (reallocation of elements).
- **Removal**:
  - Best case: O(1) - constant time (removal of the last element).
  - Worst case: O(n) - linear time (shifting of elements).

**Implementation in Go:**
The implementation of the Array data structure in Go can be found in the [`array.go`](src/data_structures/array.go) file.

### Linked List (English)

A Linked List is a linear data structure where each element is a node that contains a value and a reference (or pointer) to the next node in the sequence.

**Characteristics:**
- **Chain of Nodes**: Each node points to the next node.
- **Variable Size**: The list can grow and shrink dynamically.
- **Sequential Access**: Access to elements is sequential, starting from the first node.

**Main Operations and Their Complexities:**
- **Access**:
  - Best case: O(1) - constant time (first element).
  - Worst case: O(n) - linear time (last element).
- **Insertion**:
  - Best case: O(1) - constant time (insertion at the beginning).
  - Worst case: O(n) - linear time (insertion at the end).
- **Removal**:
  - Best case: O(1) - constant time (removal from the beginning).
  - Worst case: O(n) - linear time (removal from the end).

**Implementation in Go:**
The implementation of the Linked List data structure in Go can be found in the file [`linked_list.go`](src/data_structures/linked_list.go).

### Stack

A Stack is a linear data structure that follows the Last In First Out (LIFO) principle. Elements can be added and removed only from the top of the stack.

**Characteristics:**
- **LIFO**: Last In First Out ordering.
- **Variable Size**: The stack can grow and shrink dynamically.
- **Top Access**: Access to elements is only at the top of the stack.

**Main Operations and Their Complexities:**
- **Push**:
  - O(1) - constant time.
- **Pop**:
  - O(1) - constant time.
- **Peek**:
  - O(1) - constant time.
- **Size**:
  - O(1) - constant time.
- **IsEmpty**:
  - O(1) - constant time.

**Implementation in Go:**
The implementation of the Stack data structure in Go can be found in the file [`stack.go`](src/data_structures/stack.go).

### Queue

A Queue is a linear data structure that follows the First In First Out (FIFO) principle. Elements can be added at the end and removed from the front of the queue.

**Characteristics:**
- **FIFO**: First In First Out ordering.
- **Variable Size**: The queue can grow and shrink dynamically.
- **Front and Rear Access**: Access to elements is only at the front and rear of the queue.

**Main Operations and Their Complexities:**
- **Enqueue**:
  - O(1) - constant time.
- **Dequeue**:
  - O(1) - constant time.
- **Peek**:
  - O(1) - constant time.
- **Size**:
  - O(1) - constant time.
- **IsEmpty**:
  - O(1) - constant time.

**Implementation in Go:**
The implementation of the Queue data structure in Go can be found in the file [`queue.go`](src/data_structures/queue.go).

### Deque

A Deque (Double-ended Queue) is a linear data structure that allows insertion and removal of elements from both ends, front and back.

**Characteristics:**
- **Double-ended**: Elements can be added or removed from both ends.
- **Variable Size**: The deque can grow and shrink dynamically.
- **Front and Rear Access**: Access to elements is at both the front and rear of the deque.

**Main Operations and Their Complexities:**
- **AddFront**:
  - O(n) - linear time.
- **AddBack**:
  - O(1) - constant time.
- **RemoveFront**:
  - O(1) - constant time.
- **RemoveBack**:
  - O(1) - constant time.
- **Size**:
  - O(1) - constant time.
- **IsEmpty**:
  - O(1) - constant time.

**Implementation in Go:**
The implementation of the Deque data structure in Go can be found in the file [`deque.go`](src/data_structures/deque.go).

### Tree

The `Tree` data structure is a hierarchical structure consisting of nodes, where each node has a value and references to its left and right children.

**Characteristics**:
- Hierarchical structure.
- Root node with zero or more child nodes.
- Each child node can have its own children, forming a subtree.

**Main Operations**:
- `Insert`: Adds a new value to the tree.
- `Search`: Checks if a value exists in the tree.
- `Remove`: Removes a value from the tree.
- `IsEmpty`: Checks if the tree is empty.

**Complexity**:
- Insert: O(log n) on average, O(n) in the worst case.
- Search: O(log n) on average, O(n) in the worst case.
- Remove: O(log n) on average, O(n) in the worst case.

**Implementation in Go:**
The implementation of the Deque data structure in Go can be found in the file [`deque.go`](src/data_structures/tree-based/tree.go.go).

## Algorithms

(Coming soon)
