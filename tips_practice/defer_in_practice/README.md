# defer maybe not work as expected

When we use defer in `for` loop, `defer` maybe not work as expected.

Our code,

```go
func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("number: %d\n", i)
		defer fmt.Printf("defer func\n")
	}

	fmt.Println("finish")
}
```


Our result,

```bash
go run main.go
number: 0
number: 1
number: 2
number: 3
number: 4
number: 5
number: 6
number: 7
number: 8
number: 9
finish
defer func
defer func
defer func
defer func
defer func
defer func
defer func
defer func
defer func
defer func
```