# Go Sorting Algorithms

This project contains sorting algorithms I learned in Introduction to Data Structures and Algorithms course (kudos to the teacher). Decided to implement them in Go, as I am learning this language (enjoying so far :D)

---

## Algorithms

| File | Algorithm | Strategy | Time (Avg) | Space | Notes |
|------|-----------|----------|------------|-------|-------|
| `bubble.go` | **Bubble Sort** | Compare & swap adjacent elements | O(n^2) | O(1) | Includes an early-exit optimization if the slice becomes sorted |
| `insertion.go` | **Insertion Sort** | Build a sorted sub-array one item at a time | O(n^2) | O(1) | Efficient for small or nearly-sorted data |
| `selection.go` | **Selection Sort** | Repeatedly find the minimum and place it at the front | O(n^2) | O(1) | Simple, but always performs ~n^2/2 comparisons |
| `merge.go` | **Merge Sort** | Divide & conquer with auxiliary merging | O(n log n) | O(n) | Returns a new sorted slice (not in-place)|
| `quick.go` | **Quick Sort** | Divide & conquer with a pivot partition | O(n log n) | O(log n) | In-place; uses the last element as the pivot |
| `heap.go` | **Heap Sort** | Build a max-heap, then extract the root | O(n log n) | O(1) | In-place; relies on recursive `heapify` |

---

## Running the code

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...
