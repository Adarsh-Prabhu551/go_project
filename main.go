package main

import (
	"fmt"
)

type Reporter interface {
	Summary() string
}

type Student struct {
	Name  string
	Score float64
	Grade string
}

type ClassReport struct {
	Students []Student
	Highest  float64
	Avg      float64
}

func (c ClassReport) Summary() string {
	sentence := fmt.Sprintf("Class Report: %v students | Highest: %.2f | Avg: %.2f\n", len(c.Students), c.Highest, c.Avg)
	return sentence
}

func (s Student) Summary() string {
	sentence := fmt.Sprintf("%v | %.2f | %v | %v \n", s.Name, s.Score, s.Grade, s.isPassing())
	return sentence
}

func printSummary(r Reporter) {
	fmt.Print(r.Summary())

}
func calcStats(students []Student) (float64, float64) {
	var highest float64 = -1
	var total float64 = 0
	var avg float64

	for i := range students {
		if students[i].Score > highest {
			highest = students[i].Score //calculate highest score
		}
		total += students[i].Score
	}
	avg = float64(total) / float64(len(students))

	return highest, avg
}

func printStudents(students []Student) {

	for i := range students {
		students[i].display()
	}
}

func aboveAvg(students []Student, avg float64) {
	fmt.Printf("The students who scored above avg :\n")
	for i := range students {
		if students[i].Score > avg {
			fmt.Printf("Student name: %v, Marks:%v and Grade:%v\n", students[i].Name, students[i].Score, students[i].Grade)
		}
	}
}

func printStats(highest float64, avg float64) {
	fmt.Printf("Highest marks is %.2f\n", highest)
	fmt.Printf("Avg marks is %.2f\n", avg)
}

func (s Student) getGrade() string {
	const (
		gradeA = 90.0
		gradeB = 75.0
		gradeC = 60.0
	)
	var grade string
	switch {
	case s.Score >= gradeA:
		grade = "A"
	case s.Score >= gradeB:
		grade = "B"
	case s.Score >= gradeC:
		grade = "C"
	default:
		grade = "F"
	}

	return grade
}

func (s Student) isPassing() string {
	var status string
	if s.Score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	return status
}

func (s Student) display() {
	fmt.Printf("Name:%v | Score:%v | Grade:%v | Status:%v\n", s.Name, s.Score, s.Grade, s.isPassing())
}

func (s *Student) updateScore(newScore float64) {
	s.Score = newScore
	s.Grade = s.getGrade()
}

func main() {
	students := []Student{
		{Name: "Alice", Score: 95.5},
		{Name: "Bob", Score: 87.0},
		{Name: "Charlie", Score: 72.0},
		{Name: "Diana", Score: 91.0},
		{Name: "Eve", Score: 58.0},
	}

	for i := range students {
		students[i].Grade = students[i].getGrade()
	}

	highest, avg := calcStats(students)
	printStats(highest, avg)

	printStudents(students)
	students[1].updateScore(85)
	students[1].display()

	aboveAvg(students, avg)

	classReport := ClassReport{
		Students: students,
		Highest:  highest,
		Avg:      avg,
	}
	printSummary(students[0])
	printSummary(classReport)

}
