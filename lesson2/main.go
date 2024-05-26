package main

import (
	"errors"
	"fmt"
)

type person struct {
	age  int
	name string
}

var (
	add = func(i, j int) int { return i + j }
)

func main() {
	//fmt.Println(sum(1, 2, 4))
	//result, reminder, err := divAndReminder(5, 2)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//fmt.Println(result, reminder)
	//deferExample()
	//i := 2
	//s := "Hello"
	//p := person{}
	//modifyFails(i, s, p)
	//fmt.Println(i, s, p)
	//modifySuccess(&i, &s, &p)
	//fmt.Println(i, s, p)
	m := map[int]string{
		1: "hello",
		2: "goodbye",
	}
	modifyMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modifySlice(s)
	fmt.Println(s)
}

func modifyMap(m map[int]string) {
	m[2] = "hola"
	m[3] = "adios"
	delete(m, 1)
}

func modifySlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modifySuccess(i *int, s *string, p *person) {
	*i = *i * 2
	*s = "Goodbye"
	p.name = "Bob"
}
func addTo(base int, vals ...int) int {
	out := make([]int, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out[0]
}

func sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total
}

func divAndReminder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func changeAdd() {
	add = func(i, j int) int { return i + j + j }
}

//defer - отложить в конец

func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}
