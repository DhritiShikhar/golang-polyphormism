// Demonstration of polymorphism in Golang

package main

import "fmt"

type income interface {
	calculate() int // calculates income from source
	source() string // returns name of source
}

type fullTimeJob struct {
	companyName   string
	months        int
	monthlyIncome int
}

type partTimeJob struct {
	jobName     string
	days        int // total number of days worked in a year
	dailyIncome int
}

// Calculate full time salary
func (fullTime *fullTimeJob) calculate() int {
	return fullTime.months * fullTime.monthlyIncome
}

// Calculate part time salary
func (partTime *partTimeJob) calculate() int {
	return partTime.days * partTime.dailyIncome
}

func totalEarning(fullTimeEarning int, partTimeEarning int) int {
	return fullTimeEarning + partTimeEarning
}

func main() {
	myFullTimeJob := fullTimeJob{
		companyName:   "ABC Private Limited",
		months:        12,
		monthlyIncome: 543,
	}

	myFullTimeEarning := myFullTimeJob.calculate()

	myPartTimeJob := partTimeJob{
		jobName:     "Driving School",
		days:        30,
		dailyIncome: 10,
	}

	myPartTimeEarning := myPartTimeJob.calculate()

	total := totalEarning(myFullTimeEarning, myPartTimeEarning)

	fmt.Println("My Yearly Earning is", total)
}
