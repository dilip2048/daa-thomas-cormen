**Quicksort**

We have implemented both the randomized and tail recursion version of the quick sort.

The codes are written in go language. This language is developed by google. Since I have previous experience in go, I chosed this language. This langauge is fast and easy to code.

To run these codes, you need to have go installed in your machine. It can be done from https://go.dev/doc/install.

### A. how to run the test cases?
1. Go to the path of the project using command line.
2. Then run `go test` on that path.
```
dilipyadav@Dilips-MacBook-Pro project % go test

Time take to sort 250 numbers is 59.163µs using randomized quicksort

Time take to sort 500 numbers is 89.587µs using randomized quicksort

Time take to sort 750 numbers is 151.468µs using randomized quicksort

Time take to sort 1000 numbers is 193.969µs using randomized quicksort

Time take to sort 1250 numbers is 279.642µs using randomized quicksort

Time take to sort 1500 numbers is 320.649µs using randomized quicksort

Time take to sort 1750 numbers is 363.238µs using randomized quicksort

Time take to sort 2000 numbers is 412.688µs using randomized quicksort

Time take to sort 2250 numbers is 505.916µs using randomized quicksort

Time take to sort 2500 numbers is 547.257µs using randomized quicksort

Time take to sort 250 numbers is 1.118024ms using tail recursion

Time take to sort 500 numbers is 7.266038ms using tail recursion

Time take to sort 750 numbers is 20.573399ms using tail recursion

Time take to sort 1000 numbers is 114.504139ms using tail recursion

Time take to sort 1250 numbers is 44.67353ms using tail recursion

Time take to sort 1500 numbers is 55.87823ms using tail recursion

Time take to sort 1750 numbers is 1.372460446s using tail recursion

Time take to sort 2000 numbers is 554.481513ms using tail recursion

Time take to sort 2250 numbers is 1.632166022s using tail recursion

Time take to sort 2500 numbers is 2.524851257s using tail recursion

PASS
ok  	github.com/grill-daa/project	6.550s
```


### B. how to run the specific test file?
1. Go to the folder path
2. run command `go test file_name.go test_file_name.go -v` as below
```
dilipyadav@Dilips-MacBook-Pro project % go test randomized_quicksort.go randomized_quicksort_test.go -v
=== RUN   TestRandomizedQuickSort1
Time take to sort 250 numbers is 53.954µs using randomized quicksort

--- PASS: TestRandomizedQuickSort1 (0.00s)
=== RUN   TestRandomizedQuickSort2
Time take to sort 500 numbers is 70.399µs using randomized quicksort

--- PASS: TestRandomizedQuickSort2 (0.00s)
=== RUN   TestRandomizedQuickSort3
Time take to sort 750 numbers is 140.336µs using randomized quicksort

--- PASS: TestRandomizedQuickSort3 (0.00s)
=== RUN   TestRandomizedQuickSort4
Time take to sort 1000 numbers is 150.207µs using randomized quicksort

--- PASS: TestRandomizedQuickSort4 (0.00s)
=== RUN   TestRandomizedQuickSort5
Time take to sort 1250 numbers is 250.433µs using randomized quicksort

--- PASS: TestRandomizedQuickSort5 (0.00s)
=== RUN   TestRandomizedQuickSort6
Time take to sort 1500 numbers is 236.819µs using randomized quicksort

--- PASS: TestRandomizedQuickSort6 (0.00s)
=== RUN   TestRandomizedQuickSort7
Time take to sort 1750 numbers is 260.087µs using randomized quicksort

--- PASS: TestRandomizedQuickSort7 (0.00s)
=== RUN   TestRandomizedQuickSort8
Time take to sort 2000 numbers is 302.617µs using randomized quicksort

--- PASS: TestRandomizedQuickSort8 (0.00s)
=== RUN   TestRandomizedQuickSort9
Time take to sort 2250 numbers is 344.558µs using randomized quicksort

--- PASS: TestRandomizedQuickSort9 (0.00s)
=== RUN   TestRandomizedQuickSort10
Time take to sort 2500 numbers is 392.637µs using randomized quicksort

--- PASS: TestRandomizedQuickSort10 (0.00s)
PASS
ok  	command-line-arguments	0.177s
dilipyadav@Dilips-MacBook-Pro project %

```

### C. How to run specific test case?
1. Go to the path of the project folder
2. Run `go test -run test_case_method_name -v` as shown below:
```
dilipyadav@Dilips-MacBook-Pro project % go test -run TestTailRecursionQuickSort2 -v
=== RUN   TestTailRecursionQuickSort2
Time take to sort 500 numbers is 6.124565ms using tail recursion

--- PASS: TestTailRecursionQuickSort2 (0.01s)
PASS
ok  	github.com/grill-daa/project	0.182s
```

