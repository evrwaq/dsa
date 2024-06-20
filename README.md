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

## Algorithms

(Coming soon)
