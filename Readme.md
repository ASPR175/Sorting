# Sorting Algorithms in Go

This repository contains clean, from-scratch implementations of fundamental sorting algorithms written in Go.
The goal of this project is to understand algorithmic tradeoffs, time complexity, space usage, and stability characteristics rather than rely on built-in library functions.

# Implemented Algorithms

1) Bubble Sort

2) Selection Sort

3) Insertion Sort

4) Merge Sort

5) Quick Sort

# Complexity Overview

---------------|------------|-----------------|------------|----------------------|--------|
Algorithm	   | Best Case	|   Average Case  |	Worst Case |	Space Complexity  |	Stable |
---------------|------------|-----------------|------------|----------------------|--------|
Bubble  Sort   |	O(n)	|    O(n²)	      |O(n²)	   |    O(1)	          |  Yes   | 
Selection Sort |	O(n²)	|    O(n²)	      |O(n²)	   |    O(1)	          |  No    | 
Insertion Sort |	O(n)	|    O(n²)	      |O(n²)	   |    O(1)	          |  Yes   |
Merge Sort	   | O(n log n)	|    O(n log n)	  |O(n log n)  |	O(n)	          |  Yes   |
Quick Sort	   | O(n log n)	|    O(n log n)	  |O(n²)	   |    O(log n)*	      |  No    |
---------------|------------|-----------------|------------|----------------------|--------|

*Quick Sort space depends on recursion depth.

# Algorithm Details



1. Bubble Sort Concept

Repeatedly compares adjacent elements and swaps them if they are in the wrong order.
After each pass, the largest unsorted element “bubbles” to the end.

Characteristics
- Simple but inefficient
- Useful for educational purposes
- Adaptive if optimized with early termination flag

When to Use
- Very small datasets
- Teaching basic sorting mechanics



2. Selection Sort Concept

Divides the array into sorted and unsorted portions.
Repeatedly selects the minimum element from the unsorted portion and places it at the beginning.

Characteristics
- Performs fewer swaps than Bubble Sort
- Not stable by default
- Time complexity remains O(n²) regardless of input

When to Use
- When memory writes are expensive
- Educational demonstration



3. Insertion Sort Concept

Builds the sorted array one element at a time by inserting each new element into its correct position in the already sorted portion.

Characteristics
- Efficient for small or nearly sorted datasets
- Stable
- In-place

When to Use
- Small arrays
- Nearly sorted data
- As a hybrid algorithm base case



4. Merge Sort Concept

Uses divide and conquer:
Split array into halves
Recursively sort each half
Merge the sorted halves

Characteristics
- Guaranteed O(n log n)
- Stable
- Requires additional memory

When to Use
- Large datasets
- When stability is required
- External sorting scenarios



5. Quick Sort Concept

Uses divide and conquer:
Select a pivot
Partition array into elements smaller and larger than pivot
Recursively sort partitions

Characteristics
- Fast in practice
- In-place (depending on implementation)
- Worst-case O(n²) without good pivot selection

When to Use
- General-purpose sorting
- Performance-critical scenarios
- When memory usage must be minimal
- What This Project Demonstrates
- Understanding of algorithmic tradeoffs
- Manual implementation without relying on standard library sort
- Complexity analysis and stability considerations
- Recursive and iterative problem solving patterns