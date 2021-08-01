## 01. Determine if a number in a list [code]

Given a list of numbers, determine if a specific number in that list.

Ex:
```go
NumInList([]int{1, 2, 3, 4, 5}, 5)      // true
NumInList([]int{3, 3, 3, 3, 3}, 5)      // false
NumInList([]int{3, 3, 5, 5, 5}, 5)      // true
NumInList([]int{3, 5, -10, 34}, -10)    // true

// empty list!
NumInList(nil, 5)       // false
NumInList([]int{}, 5)   // false
```

## 02. Sum all the numbers in a list

Source file: `sum.go`
Function def: `Sum(numbers []int) int`

Given a list of numbers, sum them all up and return the sum

Ex:

```go
Sum([]int{1, 2, 3, 4, 5})           // 15
Sum([]int{3, 3, 3, 3, 3})           // 15
Sum([]int{3, 5, 3, 5, 3})           // 19
Sum([]int{4, 2, 22, -10, 8})        // 26

// list's just return 0 for empty lists
Sum(nil)        // 0
Sum([]int{})    // 0
```