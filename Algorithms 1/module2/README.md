### Union find problem
A set of algorithms for solving dynamic connectivity problem.

#### Dynamic connectivity
Given a set of N objects.
- **Union command**: connect two objects.
- **Find/connected query**: Is there a path connecting the two objects?

#### Algorithms
- Quick find (eager approach)
- Quick union (lazy approach)
- Quick union improvement

#### Analysis
- **Quick Find**: O(n)
- **Quick Union**: O(n)
- **Weighted Quick Union with path compression**: O(log n)