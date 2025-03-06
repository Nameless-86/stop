# Go Exercises

## Recursion

1. **Factorial**: Write a recursive function to calculate the factorial of a given number.
2. **Fibonacci Sequence**: Implement a recursive function to find the nth number in the Fibonacci sequence.
3. **Sum of Digits**: Write a recursive function to find the sum of digits of a given integer.
4. **Palindrome Check**: Create a recursive function to check if a string is a palindrome.
5. **Power Calculation**: Write a recursive function to calculate `x^n` (x raised to the power n).

## Concurrency

1. **Parallel Factorial**: Compute factorials of multiple numbers concurrently and combine the results.
2. **Prime Numbers**: Create a program where goroutines check if numbers in a given range are prime.
3. **Web Scraper**: Write a program where each goroutine fetches the content of a URL from a list.
4. **Channel-based Sum**: Use multiple goroutines to calculate the sum of different parts of a slice, then combine results using channels.
5. **Task Queue**: Simulate a task queue where workers (goroutines) handle tasks from a shared channel.

## Interfaces

1. **Shape Area Calculator**: Define an interface `Shape` with a method `Area()`. Implement it for `Circle` and `Rectangle`.
2. **Sortable Slice**: Create an interface for sorting slices of different types (e.g., integers and strings).
3. **Animal Sounds**: Define an interface `Animal` with a method `Speak()`. Implement it for different animals like `Dog`, `Cat`, and `Cow`.
4. **Payment System**: Define a `Payment` interface with a `Pay()` method. Implement it for `CreditCard`, `PayPal`, and `Cash`.
5. **Logger**: Create a logging interface with a method `Log(message string)`. Implement it for logging to a file and to the console.

## Slices

1. **Reverse Slice**: Write a function to reverse the elements of a slice of integers.
2. **Unique Elements**: Create a function to return only the unique elements in a slice.
3. **Filter Slice**: Write a function that filters elements of a slice based on a predicate function.
4. **Merge and Sort**: Combine two slices and return a sorted version of the merged slice.
5. **Rotate Slice**: Implement a function to rotate a slice (e.g., move the first element to the end).

---

## Go Template for Exercises

```go
package main

import (
	"fmt"
)

// Your function implementation here
func ExampleFunction() {
	// Implement the logic here
}

func main() {
	// Call your function and print the result
	fmt.Println("Result:", ExampleFunction())
}
```
