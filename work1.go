package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"runtime"
	"strings"
	"time"
	//"github.com/entropyx/countword/algortihms"
)

func WordCount(s []string) map[string]int {
	res := make(map[string]int)
	for _, str := range s {
		res[strings.ToLower(str)]++
	}
	return res
}

func WordCount2(s []string, c chan map[string]int) {
	res := make(map[string]int)
	for _, str := range s {
		res[strings.ToLower(str)]++
	}
	c <- res
}

func main() {
	t := time.Now()
	var result2 []string
	dat, _ := ioutil.ReadFile("apocalipsis.txt")

	result := strings.Split(string(dat), " ")

	for i := 0; i < len(result); i++ {
		x := strings.Split(result[i], "\n")
		for j := 0; j < len(x); j++ {
			if x[j] != "" {
				result2 = append(result2, x[j])
			}
		}

	}
	out1 := WordCount(result2)
	fmt.Println("Time counting words 1:", time.Since(t))
	fmt.Println(out1["pero"])
	fmt.Println(out1["venida"])
	fmt.Println(out1["amor"])

	t = time.Now()
	n := runtime.NumCPU()
	N := int(math.Floor(float64(len(result2) / n)))
	c := make(chan map[string]int)
	for i := 0; i < n; i++ {
		j1 := i * N
		j2 := (i + 1) * N
		go WordCount2(result2[j1:j2], c)
	}

	if n*N != len(result2) {
		go WordCount2(result2[(n*N-1):(len(result2)-1)], c)
		n++
	}

	out := map[string]int{}
	for i := 0; i < n; i++ {
		u := <-c
		for k, v := range u {
			out[k] = out[k] + v
		}
	}

	close(c)
	fmt.Printf("Time counting words 2: %s \n", time.Since(t))
	fmt.Println(out["pero"])
	fmt.Println(out["venida"])
	fmt.Println(out["amor"])
}
