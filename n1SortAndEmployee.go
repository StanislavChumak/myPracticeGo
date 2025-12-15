package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sort
func merge(arr []int, p, q, r int) {
	left := append([]int(nil), arr[p:q+1]...)
	right := append([]int(nil), arr[q+1:r+1]...)
	i, j := 0, 0
	for k := p; k <= r; k++ {
		if j >= len(right) || (i < len(left) && left[i] <= right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
}

func sortMergeRec(arr []int, p, r int) {
	if p < r {
		q := (p + r) / 2
		sortMergeRec(arr, p, q)
		sortMergeRec(arr, q+1, r)
		merge(arr, p, q, r)
	}
}

func sortMerge(arr []int) {
	sortMergeRec(arr, 0, len(arr)-1)
}

// Employee
type Employee struct {
	name     string
	age      int
	position string
	salary   int
}

func (e Employee) print() {
	fmt.Printf("Имя: %s, возрост: %d\nПозиция: %s, зарплата: %d\n", e.name, e.age, e.position, e.salary)
}

var commands = `
1 - Добавить нового сотрудника
2 - Удалить сотрудника
3 - Вывести список сотрудников
4 - Выйти из программы
`

func main() {
	//sort
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var arr [100]int
	for i := 0; i < len(arr); i++ {
		arr[i] = r.Intn(100)
	}

	fmt.Println(arr)

	sortMerge(arr[:])

	fmt.Println(arr)

	// Employee
	var empls []Employee
Loop:
	for {
		var cmd int
		fmt.Print(commands)
		fmt.Scanf("%d", &cmd)

		switch cmd {
		case 1:
			var empl Employee
			fmt.Println("\nИмя, возраст, позиция, зарплата")
			fmt.Scanf("%s %d %s %d", &empl.name, &empl.age, &empl.position, &empl.salary)
			empls = append(empls, empl)
		case 2:
			fmt.Println("Введите номер сотрудника")
			var n int
			fmt.Scan("%d", &n)
			empls = append(empls[:n], empls[n+1:]...)
		case 3:
			for _, empl := range empls {
				empl.print()
			}
		case 4:
			break Loop
		}
	}

}
