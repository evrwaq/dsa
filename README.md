# Data Structures and Algorithms (DSA)
*[Leia em Português](README.pt-br.md)*

This project aims to demonstrate the main data structures and algorithms, as well as document and explain each one.

## Data Structures

<details>
  <summary>Linear Data Structures</summary>

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
The implementation of the Array data structure in Go can be found in the [`array.go`](src/data_structures/linear/array.go) file.

### Linked List

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
The implementation of the Linked List data structure in Go can be found in the file [`linked_list.go`](src/data_structures/linear/linked_list.go).

### Stack

A Stack is a linear data structure that follows the Last In, First Out (LIFO) principle. Elements are added and removed from the same end, called the top of the stack.

**Characteristics:**
- **LIFO**: Last In, First Out ordering.
- **Variable Size**: The stack can grow and shrink dynamically.
- **Top Access**: Access to elements is only at the top of the stack.

**Main Operations and Their Complexities:**
- **Push** (add an element to the top):
  - O(1) - constant time.
- **Pop** (remove an element from the top):
  - O(1) - constant time.
- **Peek** (view the top element without removing it):
  - O(1) - constant time.
- **Size**:
  - O(1) - constant time.
- **IsEmpty**:
  - O(1) - constant time.

**Implementation in Go:**
The implementation of the Stack data structure in Go can be found in the file [`stack.go`](src/data_structures/linear/stack.go).

### Queue

A Queue is a linear data structure that follows the First In, First Out (FIFO) principle. Elements are added at the back (rear) and removed from the front.

**Characteristics:**
- **FIFO**: First In, First Out ordering.
- **Variable Size**: The queue can grow and shrink dynamically.
- **Front and Rear Access**: Elements are added at the rear and removed from the front.

**Main Operations and Their Complexities:**
- **Enqueue** (add an element to the rear):
  - O(1) - constant time.
- **Dequeue** (remove an element from the front):
  - O(1) - constant time.
- **Peek** (view the front element without removing it):
  - O(1) - constant time.
- **Size**:
  - O(1) - constant time.
- **IsEmpty**:
  - O(1) - constant time.

**Implementation in Go:**
The implementation of the Queue data structure in Go can be found in the file [`queue.go`](src/data_structures/linear/queue.go).

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
The implementation of the Deque data structure in Go can be found in the file [`deque.go`](src/data_structures/linear/deque.go).

</details>

<details>
  <summary>Tree-based Data Structures</summary>

### Tree

A Tree is a hierarchical data structure consisting of nodes, where each node has a value and references to its left and right children.

**Characteristics:**
- **Hierarchical Structure**: Composed of nodes with parent-child relationships.
- **Root Node**: The top node with zero or more child nodes.
- **Subtrees**: Each child node can have its own children, forming subtrees.
- **Binary Trees**: Trees where each node has at most two children (left and right).

**Main operations and their complexities:**
- **Insert**:
  - Average case: O(log n) - logarithmic time.
  - Worst case: O(n) - linear time (unbalanced tree).
- **Search**:
  - Average case: O(log n) - logarithmic time.
  - Worst case: O(n) - linear time (unbalanced tree).
- **Remove**:
  - Average case: O(log n) - logarithmic time.
  - Worst case: O(n) - linear time (unbalanced tree).
- **IsEmpty**:
  - O(1) - constant time.

**Implementation in Go:**
The implementation of the Tree data structure in Go can be found in the file [`tree.go`](src/data_structures/tree-based/tree.go).

</details>

## Algorithms

(Coming soon)
