## Knuth’s method

Knuth’s method, also known as the “reservoir sampling” algorithm, is a simple yet efficient technique for selecting a random sample from a stream of data without storing the elements.

This is a method of randomly sampling n items from a set of M items, with equal probability; where M >= n and M, the number of items is unknown until the end. This means that the equal probability sampling should be maintained for all successive items > n as they become available (although the content of successive samples can change).
