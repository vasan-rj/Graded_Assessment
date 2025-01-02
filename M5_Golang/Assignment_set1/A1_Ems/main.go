package main

import (
	"errors"
	"fmt"
	"strings"
)

type Staff struct {
	EmpID       int
	FullName    string
	YearsOld    int
	Team        string
}

var teamMembers []Staff

const (
	HumanResources = "HR"
	InformationTech = "IT"
)

func main() {
	addErr := registerStaff(101, "Anand", 29, InformationTech)
	if addErr != nil {
		fmt.Println(addErr)
	}
	addErr = registerStaff(102, "Bhuvana", 24, HumanResources)
	if addErr != nil {
		fmt.Println(addErr)
	}

	findStaffByID(101)
	findStaffByName("Bhuvana")
	showTeamMembers(HumanResources)
	countTeamMembers(InformationTech)
}

func registerStaff(empID int, fullName string, yearsOld int, team string) error {
	for _, member := range teamMembers {
		if member.EmpID == empID {
			return errors.New("employee ID should be unique")
		}
	}

	if yearsOld <= 18 {
		return errors.New("employee age must be above 18")
	}

	teamMembers = append(teamMembers, Staff{
		EmpID:    empID,
		FullName: fullName,
		YearsOld: yearsOld,
		Team:     team,
	})
	fmt.Println("Staff member added successfully!")
	return nil
}

func findStaffByID(empID int) {
	for _, member := range teamMembers {
		if member.EmpID == empID {
			fmt.Printf("Staff found: %+v\n", member)
			return
		}
	}
	fmt.Println("Error: Staff not found.")
}

func findStaffByName(fullName string) {
	for _, member := range teamMembers {
		if strings.EqualFold(member.FullName, fullName) {
			fmt.Printf("Staff found: %+v\n", member)
			return
		}
	}
	fmt.Println("Error: Staff not found.")
}

func showTeamMembers(team string) {
	fmt.Printf("Staff members in %s team:\n", team)
	for _, member := range teamMembers {
		if member.Team == team {
			fmt.Printf("- %+v\n", member)
		}
	}
}

func countTeamMembers(team string) {
	num := 0
	for _, member := range teamMembers {
		if member.Team == team {
			num++
		}
	}
	fmt.Printf("Total staff in %s team: %d\n", team, num)
}
