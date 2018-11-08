# Callstats.io Server Team Assignment
## Usage
Use the following commands to install & build.

```
go get github/pykmi/go-sliding-window
go build
```

From there you can run the program.

```
./go-sliding-window --size 100 --input=input/test2.csv > output.csv
```

All three test files have been included to make running the tests easier.
To alternate between tests and window sizes, you can manipulate the --size and --input flags.

## Objective
Write a Sliding Window program that reads input from .cvs files. 

* Sliding Window must have an adjustable size constraint
* Input source file must also be adjustable
* Include addDelay & getMedian interfaces to the program
* Output should be directed to a file and median calculated each time a value is added.

## Analysis
As with any problem, I tried to break the problem set down into simple list of requirements. After that I broke the design of the program into its component parts.

It was clear that control of the sliding window would be crucial part. Median is relatively easy to calculate and the formulas had been given in the assignment.

## Problem
Obviously the solution was to save values into an array but the principle problem was that to remove items from a sliding window in the correct order you need the array to be essentially unordered; the values presented as they were added. But to calculate median you need the array to be ordered, from smallest to largest.

The first choice was to maintain the array unordered and sort it each time a median is calculated. This would require almost no maintenance and the easiest (quickest) to write. The assignment made it clear that time-complexity and performance should be considered however, so I took the more complicated approach.

## Tools
I decided to write this program in Go.

It would have, likely, been easier to do so in Javascript but I wanted the challenge of doing it in Go and also because I am heavily biased towards Go, because:

* Go is a language of choice in CallStat.io
* It performs much better than a node application.
* It is type safe.
* The Go binary can be compiled, specific to the target system, and transported as a single binary file. It will not require any additional installations.
* I also liked the challenge of doing it in Go, which, as a language, is less forgiving than Javascript.
* I also had not written Go in almost a year, so it was good practice to brush off on those skills.

Additional tools were writing unit tests while developing the code. I only wrote the main function at the very end, once the sliding window was complete and working.

Also I relied heavily on the Go repl at repl.it, and obviously Google / Stack Overflow.

## Solution
The program maintains two arrays instead of one; one ordered (for calculating median) and one unordered, to maintain the order in which items are cut from the window, as the window reaches its maximum length. The program also maintains the length of the array, to similarly avoid recalculating the length of array.

### addDelay Interface
The list of operations, when a value is added to the window, is as follows:

1. Check if the array is full and cut the oldest value if necessary
2. Check if the array there are any values in the array yet
	* Also check if the new value is the larger than any previous value
	* This is simply because the response for both checks is the same and we can avoid code repetition
3. Check if the new value is smaller than any previously appended value
4. Finally we know the value belongs somewhere in middle of the array
	* It was important to determine this because of the way that Go handles appending values to an array
	* The corresponding code had a bug where it would fail to append to the beginning or the end of the array and turned out it was less time consuming to make a couple of extra checks than to debug the line

There are a few helper functions, like cut() and copy() but these are more to avoid code repetition and also to keep addDelay() clean and readable.

### getMedian Interface
This was a simple method to write; the algorithm had already been given in the assignment.

1. Check to make sure there is more than 1 value in the array or return -1.
2. Check if the number of values in the array is an odd number
	* Apply the algorithm Median (M) = ((n + 1)/2th item from the sorted array, where n = length of the array.
3. Check if the number of values in the array is an even number
	* Apply the algorithm Median (M) = [ ((n)/2)/th item + ((n)/2 + 1)/th item ] / 2, where n = length of the array.

There are a few minor deviations from the assignment that have to do with idiomatic Go, such as function names.

## Improvements
I am reasonably pleased with the way the program turned out. Although the Go standard library, no doubt, uses the most efficient sorting algorithms, I am still convinced that this approach was better than continuously relying on those sorting methods.

But there are still ways the program could be improved.

### Unit Tests
As often is the case, when on a schedule, tests suffer. I used unit tests to help debug my code and test it in small pieces but I simply ran out of allotted time (5 hours) to write acceptable unit tests for the different parts of the program.

### Binary Search
As the size of the window goes up to thousands of even tens of thousands of values the way the correct values are becomes inefficient. A binary search, for example, could significantly improve the program’s performance.

The presumption is that the sliding window would be used to calculate any number of measurements from the incoming data, in which case the way the window’s internal array is handled and searched becomes even more crucial.

### Turned Into a Package
I would have preferred to turn the sliding window into a separate package, independent from the main program that runs it. This way the process of calculating measurements could be spread into smaller processes. It also makes it easier to maintain, separate from the main program that could have its own development process.

### Real-time Processing
As it stands, the current program only runs once but the presumption is that the system is constantly adding new values to the input source. The program could be improved by making it run as a background process, monitoring for changes in the input source and processing new values as they are added.

### Use Goroutines
As interfaces are added to the library, in case of a monolithic design (multiple / all measurements are calculated by one runtime), it would be beneficial to have measurements being calculated asynchronously, instead of waiting one to finish before the next can be started.

