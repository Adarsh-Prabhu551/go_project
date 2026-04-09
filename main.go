package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score float64
	Grade string
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

func (s Student) isPassing() bool {
	return s.Score >= 60
}

func (s Student) display() {
	var status string
	if s.isPassing() {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Printf("Name:%v | Score:%v | Grade:%v | Status:%v\n", s.Name, s.Score, s.Grade, status)
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
}
