package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Grade int

const (
	A Grade = iota // makes sure all the constants unique
	B
	C
	D
	E
)

type Data struct {
	x int
	y int
}

func switchStatement() {
	x := 7
	switch {
	case x > 3:
		fmt.Printf("1")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}
}

func stringss() {
	s := "sudhanva"
	println(strings.ToUpper(s) + strings.TrimSpace("                 aha                     "))

	if unicode.IsDigit('2') {
		println("is digit")
	}

	strconv.Atoi("12")
	strconv.Itoa(122)

	for i := range 10 {
		fmt.Printf("%s ", strconv.Itoa(i))
	}

	st := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(st)
}

func dummy() {
	fmt.Printf("Hi %s", "joe")
	fmt.Println()

	channel := make(chan bool)

	go func() {
		channel <- true
	}()

	com := complex(1, 2)

	println(com)
}

func arrays() {
	var arr [5]int
	arr[0] = 1
	fmt.Println(arr)

	t := [...]int{1, 2} // array literal

	fmt.Println(t)

	for i, v := range t {
		println(i, "-> ", v)
	}

	println()
}

func slicess() {
	var nums [6]string = [...]string{"1", "2", "3", "4", "5", "6"}
	s1 := nums[1:4]
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4, 5} // slice literal
	println(len(s2))
	s2 = append(s2, 6)
	fmt.Println(s2)
	s2 = append(s2, 7, 8, 9)
	println(cap(s2))

	sl := make([]int, 5) // size = capacity
	sl[0] = 1

	sl1 := make([]int, 5, 10) // size = 5, capacity = 10
	sl1[0] = 1
}

func hashTables() {
	var mp map[string]string
	mp = make(map[string]string)
	mp["1"] = "one"
	mp["2"] = "two"

	if mp["10"] == "" {
		mp["10"] = "ten"
	}

	for k, v := range mp {
		fmt.Printf("%s -> %s\n", k, v)
	}

	mp["1"] = "onee"

	delete(mp, "2")

	fmt.Println(mp, mp["1"])
}

func structs() {
	type Person struct {
		name string
		age  int
		addr string
	}

	p := Person{}
	p.name = "sudhanva"
	p.age = 22
	p.addr = "bangalore"
	fmt.Println(p)

	a := Person{
		name: "joe",
		age:  22,
		addr: "bangalore",
	}
	fmt.Println(a)

	p1 := new(Person)
	p1.name = "john"
	p1.age = 23
	p1.addr = "mumbai"
	fmt.Println(p1.addr)
}

type P struct {
	X string
	Y int
}

func jsonn() {
	p := P{
		X: "sudhanva",
		Y: 22,
	}

	jSon, err := json.Marshal(p)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(jSon))

	var p1 P

	json.Unmarshal(jSon, &p1)

	fmt.Println(p1)
}

type Name struct {
	Fname [20]byte
	Lname [20]byte
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func bubbleSort(nums []int) {
	for i := range nums {
		for j := range len(nums) - 1 - i {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func makeDistToOrigin(ox, oy float32) func(x, y float32) float32 {
	return func(x, y float32) float32 {
		return float32(math.Sqrt(float64((x-ox)*(x-ox) + (y-oy)*(y-oy))))
	}
}

func closure() func(a, b int) int {
	x := 100
	return func(a, b int) int {
		x += a + b
		return x
	}
}

func someFunc(vals ...int) {}

func getMax(vals ...int) int {
	max := math.MinInt32
	for _, v := range vals {
		if v > max {
			max = v
		}
	}

	slice := []int{1, 2, 3, 4, 5}

	defer fmt.Println(1, 2, 3, 4, 5)

	someFunc(slice...)

	return max
}

func deferFunc() {
	i := 1
	defer fmt.Println(i + 1)
	i++
	fmt.Println(i + 1)
}

func test(s string) int {
	return len(s)
}

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main___() {
	fn := GenDisplaceFn(9.8, 2, 1)

	fmt.Println(fn(2))

	fmt.Println(fn(5))
}
