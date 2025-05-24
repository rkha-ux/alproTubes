package main

import (
	"fmt"
)

const nmax int = 10

type arr [nmax]personal

type personal struct {
	interest, skill string
}

type career struct {
	position string
	company  string
}

type rec struct {
	skill, interest, industry string
	avgsalary, compatibility  int
}

type test [nmax]rec

var compatibility int
var companyName, industry string

var idxInterest int = 1
var idxSkill int = 1
var x arr

func main() {
	var n int
	for n != 10 {
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
		case 5:
			searchPath()
		case 6:
			compatibilityi()
		case 7:
			salarys()
		case 8:
			sortCompatibility()
		case 9:
			sortSalary()
		}
	}
}

func printing() {
	fmt.Println()
	fmt.Println("Finding Career Program [BETA]")
	fmt.Println("____________________")
	fmt.Printf("type 1 to input data\ntype 2 to show data\ntype 3 to delete data\ntype 4 to modify data\ntype 5 to search for available jobs\ntype 6 to see the compatibility\ntype 7 to see the average salary\ntype 8 to sort by compatibility\ntype 9 to sort by salary\ntype 10 to exit the program\n")
	fmt.Print("input : ")
}

func inputData() {
	var y int
	fmt.Printf("type 1 to input skill data\ntype 2 to input interest data\ntype 3 to input company name\ntype 4 to input industry\n")
	fmt.Scan(&y)
	if y == 1 {
		fmt.Scan(&x[idxSkill].skill)
		idxSkill++
	} else if y == 2 {
		fmt.Scan(&x[idxInterest].interest)
		idxInterest++
	} else if y == 3 {
		fmt.Println("Please input your desired company:")
		fmt.Scan(&companyName)
	} else if y == 4 {
		fmt.Printf("Please input your desired industry: ")
		fmt.Scan(&industry)
	} else {
		fmt.Print("Invalid input")
	}
}

func deleteData() {
	var y, i int
	fmt.Printf("type 1 to delete skill data\ntype 2 to delete interest data\ninput : ")
	fmt.Scan(&y)
	fmt.Printf("type which index you want to delete\ninput : ")
	fmt.Scan(&i)

	if y == 1 {
		if i >= 1 && i < idxSkill {
			for j := i; j < idxSkill-1; j++ {
				x[j].skill = x[j+1].skill
			}
			x[idxSkill-1].skill = ""
			idxSkill--
		} else {
			fmt.Println("Invalid index")
		}
	} else if y == 2 {
		if i >= 1 && i < idxInterest {
			for j := i; j < idxInterest-1; j++ {
				x[j].interest = x[j+1].interest
			}
			x[idxInterest-1].interest = ""
			idxInterest--
		} else {
			fmt.Println("Invalid index")
		}
	} else {
		fmt.Printf("Wrong input\n")
	}
}

func modifyData() {
	var y, i int
	fmt.Printf("type 1 to modify skill data\ntype 2 to modify interest data\ninput : ")
	fmt.Scan(&y)
	fmt.Printf("type which index you want to modify\ninput : ")
	fmt.Scan(&i)

	if i < 1 || i >= nmax {
		fmt.Println("Invalid index")
		return
	}

	if y == 1 {
		fmt.Scan(&x[i].skill)
		fmt.Println()
	} else if y == 2 {
		fmt.Scan(&x[i].interest)
		fmt.Println()
	} else if y == 4 {
		fmt.Print("wrong input")
	}
}

func outputData() {
	var y int
	fmt.Printf("type 1 to print skill data\ntype 2 to print interest data\ninput : ")
	fmt.Scan(&y)

	if y == 1 {
		for i := 1; i < idxSkill; i++ {
			fmt.Printf("%d. %s\n", i, x[i].skill)
		}
	} else if y == 2 {
		for i := 1; i < idxInterest; i++ {
			fmt.Printf("%d. %s\n", i, x[i].interest)
		}
	} else {
		fmt.Print("wrong input")
	}
}

func searchPath() {
	var companyNames = [nmax]string{
		"",
		"microsoft", "amazon", "google", "meta", "apple",
		"nvidia", "tesla", "samsung", "huawei",
	}

	var positionList = [nmax]string{
		"",
		"ai_engineer", "backend_developer", "data_analyst", "data_analyst",
		"data_scientist", "frontend_developer", "fullstack_developer",
		"it_consultant", "mobile_developer",
	}

	foundCompany := false
	for i := 1; i < nmax; i++ {
		if companyName == companyNames[i] {
			foundCompany = true
		}
	}

	if !foundCompany {
		fmt.Println("Your desired company is not on the list.")
		return
	}

	low := 1
	high := nmax - 1
	foundPosition := false

	for low <= high {
		mid := (low + high) / 2
		if positionList[mid] == industry {
			foundPosition = true
			break
		} else if positionList[mid] < industry {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundPosition {
		fmt.Printf("You are compatible to work in %s as a %s.\n", companyName, industry)
	} else {
		fmt.Printf("There's no position available for your desired position at %s.\n", companyName)
	}
}

func filling(xy *test) {
	xy[1] = rec{"Programming", "GameDeveloping", "backend_developer", 4500000, 0}
	xy[2] = rec{"Programming", "WebDeveloping", "frontend_developer", 4500000, 0}
	xy[3] = rec{"Programming", "DataAnalyzing", "ai_engineer", 12323, 0}
	xy[4] = rec{"Programming", "MobileDeveloping", "fullstack_developer", 123123, 0}
	xy[5] = rec{"Programming", "GameDeveloping", "data_scientist", 1, 0}
	xy[6] = rec{"Programming", "WebDeveloping", "it_consultant", 1, 0}
	xy[7] = rec{"Programming", "DataAnalyzing", "data_analyst", 1, 0}
	xy[8] = rec{"Programming", "MobileDeveloping", "data_analyst", 1, 0}
	xy[9] = rec{"Programming", "UIUXDesigning", "backend_developer", 1, 0}
}

func compatibilityi() {
	var xy test
	var y, yy int
	compatibility = 0
	filling(&xy)

	fmt.Printf("input which index skills you want to find your compatibility from\ninput: ")
	fmt.Scan(&y)
	fmt.Printf("input which index interest you want to find your compatibility from\ninput: ")
	fmt.Scan(&yy)

	for j := 1; j < nmax; j++ {
		if x[y].skill == xy[j].skill && x[y].skill != "" {
			compatibility++
		}
		if industry == xy[j].industry && industry != "" {
			compatibility++
		}
		if x[yy].interest == xy[j].interest && x[yy].interest != "" {
			compatibility++
		}
	}

	if compatibility == 3 {
		fmt.Print("You are fully compatible")
	} else if compatibility == 2 {
		fmt.Print("You are closely compatible")
	} else if compatibility == 1 {
		fmt.Print("You are compatible enough")
	} else {
		fmt.Print("You are not compatible")
	}
	fmt.Println()
}

func salarys() {
	var xy test
	filling(&xy)

	found := false
	for i := 1; i < nmax; i++ {
		if industry == xy[i].industry {
			fmt.Printf("The average salary in the %s industry is %d\n", industry, xy[i].avgsalary)
			found = true
			return
		}
	}
	if !found {
		fmt.Println("No salary data available for that industry.")
	}
}

func sortCompatibility() {
	var xy test
	filling(&xy)

	for i := 1; i < nmax; i++ {
		xy[i].compatibility = 0
		for j := 1; j < idxSkill; j++ {
			if xy[i].skill == x[j].skill && x[j].skill != "" {
				xy[i].compatibility++
			}
		}
		for j := 1; j < idxInterest; j++ {
			if xy[i].interest == x[j].interest && x[j].interest != "" {
				xy[i].compatibility++
			}
		}
		if industry == xy[i].industry && industry != "" {
			xy[i].compatibility++
		}
	}

	for i := 1; i < nmax-1; i++ {
		maxIdx := i
		for j := i + 1; j < nmax; j++ {
			if xy[j].compatibility > xy[maxIdx].compatibility {
				maxIdx = j
			}
		}
		if maxIdx != i {
			xy[i], xy[maxIdx] = xy[maxIdx], xy[i]
		}
	}

	fmt.Println("Sorted by compatibility (high to low):")
	for i := 1; i < nmax; i++ {
		fmt.Printf("%d. Skill: %s, Interest: %s, Industry: %s, Compatibility: %d\n",
			i, xy[i].skill, xy[i].interest, xy[i].industry, xy[i].compatibility)
	}
}

func sortSalary() {
	var xy test
	filling(&xy)

	for i := 2; i < nmax; i++ {
		temp := xy[i]
		j := i
		for j > 1 && temp.avgsalary < xy[j-1].avgsalary {
			xy[j] = xy[j-1]
			j--
		}
		xy[j] = temp
	}

	fmt.Println("Sorted by average salary (low to high):")
	for i := 1; i < nmax; i++ {
		fmt.Printf("%d. Skill: %s, Interest: %s, Industry: %s, Salary: %d\n",
			i, xy[i].skill, xy[i].interest, xy[i].industry, xy[i].avgsalary)
	}
}
