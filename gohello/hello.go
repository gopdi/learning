package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"os"
	"runtime"
	"strings"
)

//Triangle bbb
type Triangle struct {
	base, alt float64
}

//Rect bbbb
type Rect struct {
	len, bre float64
}

func rec() {
	str := recover()
	fmt.Println(str)
}

//Shape bbbbb
type Shape interface {
	area() float64
}

func (t *Triangle) area() (mul float64) {
	mul = 0.5 * t.base * t.alt
	return
}

func (r *Rect) area() (mul float64) {
	mul = r.len * r.bre
	return
}

func shapearea(s Shape) float64 {
	return s.area()
}

func main() {
	runtime.GOMAXPROCS(20)
	fmt.Println(runtime.NumCPU())

	fmt.Println("hello world")

	var x, y, z int

	fmt.Println("Enter value of x")
	fmt.Scanf("%d", &x)

	fmt.Println("Enter value of y")
	fmt.Scanf("%d", &y)

	z = x + y

	fmt.Println("value of z is = ", z)

	counts := make(map[string]int)
	//_ = counts

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, count := range counts {
		fmt.Println(line, "-", count)
	}

	var s, spac string

	clargs := os.Args[1:]

	fmt.Println(clargs)

	fmt.Printf("|")

	fmt.Printf("%s", strings.Join(clargs, " "))

	fmt.Printf("|\n")

	fmt.Printf("|")
	for _, str := range clargs {

		s += spac + str
		spac = " "
		//fmt.Printf("%s ", str)
	}
	fmt.Printf("%s", s)
	fmt.Printf("|\n")
	str := []byte{65, 'b', 'c', 'd'}

	fmt.Printf("%s\n", str)

	//_ = str

	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)

	var t = Triangle{50.0, 20.0}
	var r = Rect{50.0, 20.0}

	//fmt.Println(t.area())
	//fmt.Println(r.area())
	fmt.Println(shapearea(&t))
	fmt.Println(shapearea(&r))

	//fmt.Println(runtime.NumCPU())

	defer rec()

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	el, ok := elements["Li"]

	if ok {
		fmt.Println(el["name"], el["state"])
	}

	//var x map[string]int
	// x := make(map[string]int)
	// x["pardeep"] = 31
	// x["vinant"] = 32
	// x["santosh"] = 35
	// fmt.Println(x)

	// name, ok := x["santosh"]
	// fmt.Println(name, ok)

	// delete(x, "santosh")

	// fmt.Println(x)

	//	slice1 := []int{1, 2, 3}

	//fmt.Println(slice1, slice2)
	//a := 5

	//s0.5:= tAbaseOP"t.alt//str2 := str

	//y := 0

	var p [10]int

	for i := 1; i <= 10; i++ {
		p[i-1] = i
	}

	//x := []int{1, 2, 3, 4, 5, 6, 7, 8}

	//slice2 := append(x, 1, 2, 3, 4)

	//copy(x, slice2)

	//fmt.Println(slice2)

	// for _, value := range x {
	// 	y += value
	// }

	//for i := 1; i <= 10; i++ {
	//fmt.Printf("%x - length=%d - total=%d\n", x, len(x), y)
	//}

	/*	for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				for j := 10; j >= i; j-- {
					fmt.Printf("*")
				}
			}
			fmt.Printf("\n")
		}
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				for j := 1; j <= i; j++ {
					fmt.Printf("*")
				}
			}
			fmt.Printf("\n")
		}
		/*fmt.Printf("Hello, world.- %s - %d - %d\n", str, a, b)
		fmt.Println(stringutil.Reverse("\nHello, world."))
		if str == str2 {
			fmt.Println("HOOLA-"[0] > 0 && str[0] > 0)
		}
		fmt.Printf("1+1=%c\n", "hello"[0])*/

}
