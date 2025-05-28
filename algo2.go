package main

import (
	"fmt"
)

const nmax int = 50

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
			compatibility()
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
	fmt.Println("============================================================")
	fmt.Println("		WELCOME TO FIND YOUR CAREER!				")
	fmt.Println("============================================================")
	fmt.Printf("type 1 to input data\ntype 2 to show data\ntype 3 to delete data\ntype 4 to modify data\ntype 5 to search for available jobs\ntype 6 to see the compatibility\ntype 7 to see the average salary\ntype 8 to sort by compatibility\ntype 9 to sort by salary\ntype 10 to exit the program\n")
	fmt.Println("your input : ")
}

func inputData() {
	var y int
	fmt.Printf("type 1 to input skill data\ntype 2 to input interest data\ntype 3 to input company name\ntype 4 to input industry\n")
	fmt.Scan(&y)
	if y == 1 {
		fmt.Printf("Please input your skill (index %d): ", idxSkill)
		fmt.Scan(&x[idxSkill].skill)
		idxSkill++
	} else if y == 2 {
		fmt.Printf("Please input your interest (index %d): ", idxInterest)
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
	} else {
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
		"nvidia", "tesla", "samsung", "huawei", "adobe",
		"netflix", "intel", "ibm", "cisco", "oracle",
		"salesforce", "sap", "sony", "hp", "dell",
	}

	var positionList = [nmax]string{
		"",
		"AI_Engineer", "Backend_Developer", "Data_Analyst", "Data_Scientist",
		"Frontend_Developer", "Fullstack_Developer", "IT_Consultant",
		"Mobile_Developer", "UX_Designer", "Cyber_Security_Analyst",
		"Network_Engineer", "Cloud_Engineer", "DevOps_Engineer",
		"Machine_Learning_Engineer", "Business_Analyst", "Project_Manager",
		"Database_Administrator", "Quality_Assurance", "Technical_Writer",
	}

	foundCompany := false
	for i := 1; i < nmax; i++ {
		if companyName == companyNames[i] {
			foundCompany = true
			break
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
			high = mid + 1
		}
	}

	if foundPosition {
		fmt.Printf("You are compatible to work in %s as a %s.\n", companyName, industry)
	} else {
		fmt.Printf("There's no position available for your desired position at %s.\n", companyName)
	}
}

func filling(xy *test) {
	for i := 0; i < nmax; i++ {
		xy[i] = rec{}
	}

	xy[1] = rec{"Programming", "GameDeveloping", "Backend_Developer", 4500000, 0}
	xy[2] = rec{"Programming", "WebDeveloping", "Frontend_Developer", 4700000, 0}
	xy[3] = rec{"Data_Science", "DataAnalyzing", "AI_Engineer", 6000000, 0}
	xy[4] = rec{"Programming", "MobileDeveloping", "Fullstack_Developer", 5200000, 0}
	xy[5] = rec{"Data_Science", "GameDeveloping", "Data_Scientist", 5800000, 0}
	xy[6] = rec{"IT_Consulting", "WebDeveloping", "IT_Consultant", 5500000, 0}
	xy[7] = rec{"Data_Science", "DataAnalyzing", "Data_Analyst", 4000000, 0}
	xy[8] = rec{"Mobile_Dev", "MobileDeveloping", "Mobile_Developer", 5000000, 0}
	xy[9] = rec{"UIUX_Design", "UIUXDesigning", "UX_Designer", 4200000, 0}
	xy[10] = rec{"Cyber_Security", "NetworkSecurity", "Cyber_Security_Analyst", 6500000, 0}
	xy[11] = rec{"Network_Admin", "SystemAdministration", "Network_Engineer", 4800000, 0}
	xy[12] = rec{"Cloud_Computing", "CloudInfrastructure", "Cloud_Engineer", 7000000, 0}
	xy[13] = rec{"DevOps", "Automation", "DevOps_Engineer", 7200000, 0}
	xy[14] = rec{"Machine_Learning", "AlgorithmDevelopment", "Machine_Learning_Engineer", 7500000, 0}
	xy[15] = rec{"Business_Analysis", "ProcessImprovement", "Business_Analyst", 5300000, 0}
	xy[16] = rec{"Project_Management", "TeamLeadership", "Project_Manager", 6800000, 0}
	xy[17] = rec{"Database_Management", "DataStorage", "Database_Administrator", 5100000, 0}
	xy[18] = rec{"Quality_Assurance", "SoftwareTesting", "Quality_Assurance", 3800000, 0}
	xy[19] = rec{"Technical_Writing", "Documentation", "Technical_Writer", 3500000, 0}
	xy[20] = rec{"Embedded_Systems", "HardwareProgramming", "Embedded_Engineer", 6200000, 0}
	xy[21] = rec{"Game_Development", "GraphicsProgramming", "Game_Developer", 5600000, 0}
	xy[22] = rec{"Cyber_Security", "IncidentResponse", "Cyber_Security_Analyst", 6700000, 0}
	xy[23] = rec{"Data_Engineering", "ETLPipelines", "Data_Engineer", 6300000, 0}
	xy[24] = rec{"Sales_Engineering", "ProductDemonstration", "Sales_Engineer", 5900000, 0}
	xy[25] = rec{"ERP_Consulting", "SystemImplementation", "ERP_Consultant", 6100000, 0}
	xy[26] = rec{"Blockchain_Dev", "DecentralizedApps", "Blockchain_Developer", 7800000, 0}
	xy[27] = rec{"Robotics", "AutomationDesign", "Robotics_Engineer", 7100000, 0}
	xy[28] = rec{"Cloud_Computing", "AWS_Azure_GCP", "Cloud_Architect", 8000000, 0}
	xy[29] = rec{"Bioinformatics", "GenomicAnalysis", "Bioinformatics_Scientist", 6400000, 0}
	xy[30] = rec{"FinTech", "FinancialModeling", "FinTech_Developer", 7300000, 0}
	xy[31] = rec{"Quantum_Computing", "QuantumAlgorithms", "Quantum_Engineer", 8500000, 0}
	xy[32] = rec{"Computer_Vision", "ImageProcessing", "Computer_Vision_Engineer", 7600000, 0}
	xy[33] = rec{"Augmented_Reality", "VirtualReality", "AR_VR_Developer", 6900000, 0}
	xy[34] = rec{"IoT_Development", "SensorIntegration", "IoT_Developer", 6600000, 0}
	xy[35] = rec{"Digital_Marketing", "SEO_SEM", "Digital_Marketing_Specialist", 4000000, 0}
	xy[36] = rec{"Content_Creation", "Copywriting", "Content_Writer", 3700000, 0}
	xy[37] = rec{"Technical_Support", "CustomerService", "Tech_Support_Specialist", 3200000, 0}
	xy[38] = rec{"System_Administration", "ServerManagement", "System_Administrator", 4900000, 0}
	xy[39] = rec{"Project_Coordination", "AgileMethodology", "Scrum_Master", 6000000, 0}
	xy[40] = rec{"Data_Visualization", "DashboardDesign", "Data_Viz_Specialist", 4400000, 0}
	xy[41] = rec{"UX_Research", "UserTesting", "UX_Researcher", 4300000, 0}
	xy[42] = rec{"Mobile_Game_Dev", "Unity_Unreal", "Mobile_Game_Developer", 5700000, 0}
	xy[43] = rec{"AI_Ethics", "ResponsibleAI", "AI_Ethicist", 7400000, 0}
	xy[44] = rec{"Site_Reliability", "SystemUptime", "SRE_Engineer", 7700000, 0}
	xy[45] = rec{"Customer_Success", "ClientRetention", "Customer_Success_Manager", 5400000, 0}
	xy[46] = rec{"Sales", "ClientAcquisition", "Software_Sales_Rep", 5950000, 0}
	xy[47] = rec{"Product_Management", "ProductStrategy", "Product_Manager", 7000000, 0}
	xy[48] = rec{"Solution_Architecture", "SystemDesign", "Solution_Architect", 8200000, 0}
	xy[49] = rec{"Compliance_Analyst", "RegulatoryCompliance", "Compliance_Analyst", 4600000, 0}
}

func compatibility() {
	var xy test
	filling(&xy)

	var userSkillInput, userInterestInput, userIndustryInput string

	fmt.Print("Input your skill (e.g., Programming, Data_Science): ")
	fmt.Scan(&userSkillInput)

	fmt.Print("Input your interest (e.g., GameDeveloping, DataAnalyzing): ")
	fmt.Scan(&userInterestInput)

	fmt.Print("Input your desired industry (e.g., Backend_Developer, Data_Analyst): ") // Updated example
	fmt.Scan(&userIndustryInput)

	foundAnyMatch := false
	fmt.Println("\n--- Highly Compatible Career Recommendations ---")

	for i := 1; i < nmax; i++ {
		if xy[i].skill == "" && xy[i].interest == "" && xy[i].industry == "" {
			continue
		}

		currentCompatibility := 0

		if userSkillInput != "" && xy[i].skill == userSkillInput {
			currentCompatibility++
		}

		if userInterestInput != "" && xy[i].interest == userInterestInput {
			currentCompatibility++
		}

		if userIndustryInput != "" && xy[i].industry == userIndustryInput {
			currentCompatibility++
		}

		if currentCompatibility > 0 {
			foundAnyMatch = true
			fmt.Printf("Recommendation: Skill: %s, Interest: %s, Industry: %s (Matches: %d/3) - ",
				xy[i].skill, xy[i].interest, xy[i].industry, currentCompatibility)

			switch currentCompatibility {
			case 3:
				fmt.Println("You are **fully compatible**!")
			case 2:
				fmt.Println("You are **closely compatible**.")
			case 1:
				fmt.Println("You are **compatible enough**.")
			}
		}
	}

	if !foundAnyMatch {
		fmt.Println("No career recommendations found matching your inputs.")
	}
	fmt.Println("--------------------------------------------------")
}

func salarys() {
	var xy test
	filling(&xy)

	var salaryIndustryInput string
	fmt.Print("Input the industry to see the average salary for (e.g., Backend_Developer): ") // Updated example
	fmt.Scan(&salaryIndustryInput)

	found := false
	for i := 1; i < nmax; i++ {
		if salaryIndustryInput == xy[i].industry {
			fmt.Printf("The average salary in the %s industry is %d\n", xy[i].industry, xy[i].avgsalary)
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

	var userSkillSort, userInterestSort, userIndustrySort string
	fmt.Print("Enter your skill for sorting compatibility (e.g., Programming): ")
	fmt.Scan(&userSkillSort)
	fmt.Print("Enter your interest for sorting compatibility (e.g., GameDeveloping): ")
	fmt.Scan(&userInterestSort)
	fmt.Print("Enter your desired industry for sorting compatibility (e.g., Backend_Developer): ") // Updated example
	fmt.Scan(&userIndustrySort)

	for i := 1; i < nmax; i++ {
		xy[i].compatibility = 0
		if xy[i].skill == "" && xy[i].interest == "" && xy[i].industry == "" {
			continue
		}

		if userSkillSort != "" && xy[i].skill == userSkillSort {
			xy[i].compatibility++
		}
		if userInterestSort != "" && xy[i].interest == userInterestSort {
			xy[i].compatibility++
		}
		if userIndustrySort != "" && xy[i].industry == userIndustrySort {
			xy[i].compatibility++
		}
	}

	for i := 1; i < nmax-1; i++ {
		maxIdx := i
		for j := i + 1; j < nmax; j++ {
			if xy[j].skill == "" && xy[j].interest == "" && xy[j].industry == "" {
				continue
			}
			if xy[j].compatibility > xy[maxIdx].compatibility {
				maxIdx = j
			}
		}
		if maxIdx != i {
			xy[i], xy[maxIdx] = xy[maxIdx], xy[i]
		}
	}

	fmt.Println("Sorted by compatibility (high to low):")
	foundAnySortedMatch := false
	for i := 1; i < nmax; i++ {
		if xy[i].compatibility > 0 && !(xy[i].skill == "" && xy[i].interest == "" && xy[i].industry == "") {
			foundAnySortedMatch = true
			fmt.Printf("%d. Skill: %s, Interest: %s, Industry: %s, Compatibility: %d\n",
				i, xy[i].skill, xy[i].interest, xy[i].industry, xy[i].compatibility)
		}
	}
	if !foundAnySortedMatch {
		fmt.Println("No compatible recommendations found for sorting based on your inputs.")
	}
}

func sortSalary() {
	var xy test
	filling(&xy)

	for i := 2; i < nmax; i++ {
		temp := xy[i]
		j := i
		for j > 1 && temp.avgsalary < xy[j-1].avgsalary && !(xy[j-1].skill == "" && xy[j-1].interest == "" && xy[j-1].industry == "") {
			xy[j] = xy[j-1]
			j--
		}
		xy[j] = temp
	}

	fmt.Println("Sorted by average salary (low to high):")
	foundAnySalaryEntry := false
	for i := 1; i < nmax; i++ {
		if !(xy[i].skill == "" && xy[i].interest == "" && xy[i].industry == "") {
			foundAnySalaryEntry = true
			fmt.Printf("%d. Skill: %s, Interest: %s, Industry: %s, Salary: %d\n",
				i, xy[i].skill, xy[i].interest, xy[i].industry, xy[i].avgsalary)
		}
	}
	if !foundAnySalaryEntry {
		fmt.Println("No salary data entries to display.")
	}
}
