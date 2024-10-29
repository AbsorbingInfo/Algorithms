### Quick Union Improvements
- Weighting
- Keep track of number of object in a tree
- Always link smaller tree to larger tree

### Analysis
- **Initialize**: O(n)
- **Union**: O(log n)
- **find**: O(log n)

### Further imporvement
- **Path compression**: While we are finding out root we are accessing all the path elements, we can simply make all the path elements to point to root to flatent the tree.