package main

import "fmt"

const nmax int = 10

type arr [nmax]personal

type personal struct {
	interest, skill string
}

var idxInterest int = 1
var idxSkill int = 1
var x arr

func main() {
	var n int
	for n != 5 {
		printing()
		fmt.Scan(&n)
		switch n {
		case 1:
			inputData()
		case 2:
			outputData()
		case 3:
			deleteData()
		case 4:
			modifyData()

		}
	}

}

func printing() {
	fmt.Println()
	fmt.Println("Finding Career Program [BETA]")
	fmt.Println("____________________")
	fmt.Printf("type 1 to input data\ntype 2 to show data\ntype 3 to delete data\ntype 4 to modify data\ntype 5 to exit program\n")
	fmt.Print("input : ")
}

func inputData() {
	var y int

	fmt.Printf("type 1 to input skill data\ntype 2 to input interest data\ninput : ")
	fmt.Scan(&y)
	if y == 1 {
		fmt.Scan(&x[idxSkill].skill)
		idxSkill++
	} else if y == 2 {
		fmt.Scan(&x[idxInterest].interest)
		idxInterest++
	} else {
		fmt.Print("wrong input")
	}
}

func deleteData() {
	var y, i int

	fmt.Printf("type 1 to delete skill data\ntype 2 to delete interest data\ninput : ")
	fmt.Scan(&y)
	fmt.Printf("type which index you want to delete\ninput : ")
	fmt.Scan(&i)

	if i < 1 || i >= idxSkill {
		fmt.Println("Invalid index")
		return
	}

	if y == 1 {
		if i >= 1 && i < idxSkill {
			for j := i; j < idxSkill-1; j++ {
				x[j].skill = x[j+1].skill
			}
			x[idxSkill-1].skill = ""
			idxSkill--
		}
	} else if y == 2 {
		if i >= 1 && i < idxInterest {
			for j := i; j < idxInterest-1; j++ {
				x[j].interest = x[j+1].interest
			}
			x[idxInterest-1].interest = ""
			idxInterest--
		}
	}

}

func modifyData() {
	var y, i int

	fmt.Printf("type 1 to modify skill data\ntype 2 to modify interest data\ninput : ")
	fmt.Scan(&y)
	fmt.Printf("type which index you want to modify\ninput : ")
	fmt.Scan(&i)

	if i < 1 || i >= idxSkill {
		fmt.Println("Invalid index")
		return
	}

	if y == 1 {
		fmt.Scan(&x[i].skill)
		fmt.Println()
	} else if y == 2 {
		fmt.Scan(&x[i].interest)
		fmt.Println()
	} else {
		fmt.Print("wrong input")
	}
}

func outputData() {
	var y int
	fmt.Printf("type 1 to print skill data\ntype 2 to print interest data\ninput : ")
	fmt.Scan(&y)

	if y == 1 {
		for i := 1; i <= idxSkill-1; i++ {
			fmt.Printf("%d. %s\n", i, x[i].skill)
		}
	} else if y == 2 {
		for i := 1; i <= idxInterest-1; i++ {
			fmt.Printf("%d. %s\n", i, x[i].interest)
		}
	} else {
		fmt.Print("wrong input")
	}
}
