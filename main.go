package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StrToInt(str []string) ([]int, error) {
	var ints []int
	for _, i := range str {
		i = strings.TrimSpace(i)
		if string(i) == "" || string(i) == " " {
			break
		}
		num, err := strconv.Atoi(string(i))
		if err != nil {
			return ints, err
		}
		ints = append(ints, num)
	}
	return ints, nil
}

func FindVer(m []int, ver int) bool {
	for i := 0; i < len(m); i++ {
		if m[i]+1 == ver {
			return true
		}
	}
	return false
}

func KompSvyaz(graf [][]int, start int, m []int) []int {
	for j := range graf[start-1] {
		if graf[start-1][j] == 1 {
			if !FindVer(m, j+1) {
				m = append(m, j)
				m = KompSvyaz(graf, j+1, m)
			}
		}
	}
	return m
}

func Map() func(int, map[int]int) map[int]int {
	count := 0
	return func(key int, m map[int]int) map[int]int {
		if _, found := m[key]; !found {
			m[key] = count
			count++
		}
		return m
	}
}

func Sort(mas []int) {
	for i := 0; i < len(mas)-1; i++ {
		for j := i; j < len(mas); j++ {
			if mas[i] > mas[j] {
				buf := mas[i]
				mas[i] = mas[j]
				mas[j] = buf
			}
		}
	}
}

func String(graf [][]int) {
	for _, i := range graf {
		fmt.Println(i)
	}
}

func main() {
	f := Map()
	ma := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return
	}
	nm, err := StrToInt(strings.Split(scanner.Text(), " "))
	if err != nil {
		return
	}
	graf := make([][]int, nm[0])
	for i := 0; i < nm[0]; i++ {
		graf[i] = make([]int, nm[0])
	}
	for i := 0; i < nm[1]; i++ {
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			return
		}
		v, err := StrToInt(strings.Split(scanner.Text(), " "))
		for _, j := range v {
			ma = f(j, ma)
		}
		if err != nil {
			return
		}
		if graf[ma[v[0]]][ma[v[1]]] == 0 {
			graf[ma[v[0]]][ma[v[1]]] = 1
		}
		if graf[ma[v[1]]][ma[v[0]]] == 0 {
			graf[ma[v[1]]][ma[v[0]]] = 1
		}
	}
	var m []int
	//String(graf)
	m = KompSvyaz(graf, 1, m)
	Sort(m)
	fmt.Println(len(m))
	var a []int
	for i := range m {
		for key, value := range ma {
			if m[i] == value {
				a = append(a, key)
			}
		}
	}
	Sort(a)
	for _, i := range a {
		fmt.Print(i, " ")
	}
	fmt.Println("")
	// String(graf)
}
