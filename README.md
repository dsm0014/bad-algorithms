# bad algorithms

This repo showcases some algorithms with very poor time complexity.
None of the approaches here are recommended for use anywhere outside of 
experimentation/research.

There are three examples in total:
- O(n!)  --  `go run main.go bad`
- O(n^n)  --  `go run main.go worse`
- O((n^n)!)  --  `go run main.go worst`

**WARNING** the `worst` algorithm is _incredibly inefficient_ and can take very long to execute.
Stepping through the math -- given the predefined input array size of just 3 we come up with `(3^3)!`;
simplified to `27!`. A total of approximately
```
10,888,869,000,000,000,000,000,000,000 iterations
```
