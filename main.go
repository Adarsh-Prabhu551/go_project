package main

import (
	"fmt"
	"log"
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
	sentence := fmt.Sprintf("%v | %.2f | %v | %v\n", s.Name, s.Score, s.Grade, s.isPassing())
	return sentence
}

func printSummary(r Reporter) {
	fmt.Print(r.Summary())
}

func calcStats(students []Student) (float64, float64, error) {
	var highest float64 = 0
	var total float64 = 0
	var avg float64

	for i := range students {
		if students[i].Score < 0 || students[i].Score > 100 {
			// FIX: error message now accurately describes the problem
			return 0, 0, fmt.Errorf("score %.2f for student %q is out of range (must be 0–100)", students[i].Score, students[i].Name)
		}
		if students[i].Score > highest {
			highest = students[i].Score
		}
		total += students[i].Score
	}
	avg = float64(total) / float64(len(students))

	return highest, avg, nil
}

func printStudents(students []Student) {
	for i := range students {
		students[i].display()
	}
}

func aboveAvg(students []Student, avg float64) {
	fmt.Printf("The students who scored above avg:\n")
	for i := range students {
		// FIX: skip invalid students after logging, don't fall through to the score check
		if students[i].Score < 0 || students[i].Score > 100 {
			log.Printf("Error: score %.2f for student %q is out of range (must be 0–100)", students[i].Score, students[i].Name)
			continue
		}
		if students[i].Score > avg {
			fmt.Printf("Student name: %v, Marks: %.2f and Grade: %v\n", students[i].Name, students[i].Score, students[i].Grade)
		}
	}
}

func printStats(highest float64, avg float64) {
	fmt.Printf("Highest marks is %.2f\n", highest)
	fmt.Printf("Avg marks is %.2f\n", avg)
}

func (s Student) getGrade() (string, error) {
	const (
		gradeA = 90.0
		gradeB = 75.0
		gradeC = 60.0
	)
	if s.Score < 0 || s.Score > 100 {
		return "", fmt.Errorf("score must be between 0 and 100")
	}
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

	return grade, nil
}

func (s Student) isPassing() string {
	if s.Score >= 60 {
		return "Pass"
	}
	return "Fail"
}

func (s Student) display() {
	fmt.Printf("Name: %v | Score: %.2f | Grade: %v | Status: %v\n", s.Name, s.Score, s.Grade, s.isPassing())
}

func (s *Student) updateScore(newScore float64) {
	s.Score = newScore
	grade, err := s.getGrade()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	s.Grade = grade
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
		grade, err := students[i].getGrade()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		students[i].Grade = grade
	}

	highest, avg, err := calcStats(students)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	printStats(highest, avg)
	printStudents(students)

	// Update Bob's score, then recalculate stats so downstream calls use fresh data
	students[1].updateScore(85)
	students[1].display()

	// FIX: recalculate after score update so aboveAvg and ClassReport are accurate
	highest, avg, err = calcStats(students)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	aboveAvg(students, avg)

	classReport := ClassReport{
		Students: students,
		Highest:  highest,
		Avg:      avg,
	}
	printSummary(students[0])
	printSummary(classReport)
}
