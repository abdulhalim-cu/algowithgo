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

## 03. Reverse a string [code]

Source file: `reverse.go`
Function def: `Reverse(word string) string`

Given a string, return its reverse.

Ex

```go
Reverse("cat")      // tac
Reverse("alphabet") // tebahpla
```

## 04. FizzBuzz [code]

Source file: `fizz_buzz.go`
Function def: `FizzBuzz(n int)`

Given a number N, print out all numbers from 1 to N, but when a number is divisible by 3 print out "Fizz" instead of the number and when a number divisible by 5 print out 
"Buzz" instead of the number. For a number divisible by 3 and 5 print "Fizz Buzz".

Ex

```go
FizzBuzz(1)         // 1
FizzBuzz(5)         // 1, 2, Fizz, 4, Buzz
FizzBuzz(15)        // 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz 
```

## 05. Decimal to another base [code]

Source file: `dec_to_base.go`
Function def: `DecToBase(dec, base int) string`

You can convert a number from one base to antoher by: 
1. Take the number you have in decimal and get the  remainder when it is devided by 
the new base. The remainder is the next "digit" in the new base. You add the digit
to the LEFT of the new number as you create it.

2. Divide your number by the new base, then repeat from step(1) if the new number is > 0.

Example 1:

```
14 to base 2

Step 1: 14 % 2 = 0, so next digit is "0". Our new number is currently "0"
Step 2: 14 / 2 = 7, so we update our number to be 7. 7 > 0, so we go to
sept 1...
Step 1: 7 % 2 = 1; next digit is "1"
    Our new number is currently "10" (we add the new digit to left)
Step 2: 7 / 2 = 3 (round down always), so we update our number to 3.
    3 > 0, so go to step 1
Step 1: 3 % 2 = 1; next digit is "1"
    Our new number is "110"
Step 2: 3 / 2 = 1; Go to step 1
Step 1: 1 % 2 = 1; next digit is "1"
    Our new number is "1110"
Step 2: 1/2 = 0; top here
Final number is "1110", which is 14 in base 2.
```

Example 2:

```
15 to base 16
Step 1: 15 % 16 = 15, so next digit is "F"
    Current new number is "F"
Step 2: 15 / 16 = 0. Stop!
Final new number: "F"
```

## 06. Base to Decimal [Code]

Source file: `base_to_dec.go`
Function def: `BaseToDec(value string, base int) int`

Each digit gets multiplied by the base raise to a power, e.g. for 1110 in base 2:

```
1       2       3
10^2    10^1    10^0

1       1       1       0
2^3     2^2     2^1     2^0

So this is: 1*2^3 + 1*2^2 + 1*2^1 + 0*2^0 = 8 + 4 + 2 + 0 = 14
```

charset = "0123456789ABCDEF"
value = "1110"
base = 2
power = 0
get "0"'s index charset = 0
res = res + 0*base^power

## 07. Any base to any base [code]

Source file: `base_to_base.go`
Function def: `BaseToBase(value string, base, newBase int) string`

Break it into smaller problems we have already solved!

Ex:
```go
BaseToBase("E", 16, 2) // "1110"
BaseToBase("8831A383B", 12, 16) // "DEADBEEF"
```