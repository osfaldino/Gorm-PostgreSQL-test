package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
1. Show all students
2. Show all groups

3. Add the student
4. Delete the student
5. Edit student

6. Add group
7. Delete group
8. Edit group

0. Exit
*/

type Student struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name;not null"`
	Email     string `gorm:"column:email;unique"`
	GroupID   int32  `gorm:"column:group_id"`
}

type Group struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"column:group_name"`
}

func main() {
	dsn := "host=localhost user=postgres password=1234567890 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error with DB: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Choose your action: ")
		if !scanner.Scan() {
			return
		}

		chInput := strings.TrimSpace(scanner.Text())
		ch, err := strconv.Atoi(chInput)
		if err != nil {
			ch = 0
		}

		switch ch {
		case 1:
			selectAllStudents(db)
		case 2:
			selectAllGroups(db)
		case 3:
			addStudent(db, scanner)
		case 4:
			deleteStudent(db, scanner)
		case 5:
			editStudent(db, scanner)
		case 6:
			addGroup(db, scanner)
		case 7:
			deleteGroup(db, scanner)
		case 8:
			editGroup(db, scanner)
		default:
			fmt.Println("Exiting program...")
			return
		}
	}
}

// Students
func addStudent(db *gorm.DB, scanner *bufio.Scanner) {
	var newStudent Student

	fmt.Print("\nEnter the name: ")
	if scanner.Scan() { // Scan the input
		newStudent.FirstName = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter the lastname: ")
	if scanner.Scan() {
		newStudent.LastName = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter the email: ")
	if scanner.Scan() {
		newStudent.Email = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter the group ID: ")
	if scanner.Scan() {
		gIdInput := strings.TrimSpace(scanner.Text())
		gId, err := strconv.ParseInt(gIdInput, 10, 32)
		if err != nil {
			fmt.Println("Incorrect group ID!")
			return
		}
		newStudent.GroupID = int32(gId)
	}

	result := db.Create(&newStudent)
	if result.Error != nil {
		fmt.Printf("Database error: %v\n", result.Error)
		return
	}

	fmt.Printf("Student has been added successfully! ID: %d\n\n", newStudent.ID)
}

func deleteStudent(db *gorm.DB, scanner *bufio.Scanner) {
	fmt.Print("\nEnter student ID: ")
	if !scanner.Scan() { // Scan the input
		return
	}

	idInput := strings.TrimSpace(scanner.Text())
	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect ID!")
		return
	}

	var student Student
	result := db.Delete(&student, id)

	if result.Error != nil {
		fmt.Printf("Database error during deletion: %v\n", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Printf("Student with ID %d wasn't found in DB.\n\n", id)
	} else {
		fmt.Printf("Student with ID %d has been successfully deleted!\n\n", id)
	}
}

func editStudent(db *gorm.DB, scanner *bufio.Scanner) {
	fmt.Print("\nEnter student ID: ")
	if !scanner.Scan() { // Scan the input
		return
	}
	idInput := strings.TrimSpace(scanner.Text())
	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect ID!")
		return
	}

	var student Student
	result := db.First(&student, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Student with ID %d wasn't found in DB.\n\n", id)
			return
		}
		fmt.Printf("Database error during search: %v\n", result.Error)
		return
	}

	var newStudent Student

	fmt.Print("\nEnter the name: ")
	if scanner.Scan() {
		newStudent.FirstName = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter the lastname: ")
	if scanner.Scan() {
		newStudent.LastName = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter the email: ")
	if scanner.Scan() {
		newStudent.Email = strings.TrimSpace(scanner.Text())
	}

	fmt.Print("Enter the group ID: ")
	if scanner.Scan() {
		gIdInput := strings.TrimSpace(scanner.Text())
		gId, err := strconv.ParseInt(gIdInput, 10, 32)
		if err != nil {
			fmt.Println("Incorrect group ID!")
			return
		}
		newStudent.GroupID = int32(gId)
	}

	result = db.Model(&student).Updates(newStudent)

	if result.Error != nil {
		fmt.Printf("Database error during update: %v\n", result.Error)
		return
	}

	fmt.Println("Student updated successfully!")
}

// Groups
func addGroup(db *gorm.DB, scanner *bufio.Scanner) {
	var newGroup Group

	fmt.Print("\nEnter group name: ")
	if scanner.Scan() { // Scan the input
		newGroup.Name = strings.TrimSpace(scanner.Text())
	}

	result := db.Create(&newGroup)
	if result.Error != nil {
		fmt.Printf("Database error during saving: %v\n", result.Error)
		return
	}

	fmt.Printf("Group has been added successfully! ID: %d\n\n", newGroup.ID)
}

func deleteGroup(db *gorm.DB, scanner *bufio.Scanner) {
	fmt.Print("\nEnter group ID: ")
	if !scanner.Scan() {
		return
	}

	idInput := strings.TrimSpace(scanner.Text())
	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect ID!")
		return
	}

	var group Group
	result := db.Delete(&group, id)

	if result.Error != nil {
		fmt.Printf("Database error during deletion: %v\n", result.Error)
		return
	}

	if result.RowsAffected == 0 {
		fmt.Printf("Group with ID %d wasn't found in DB.\n\n", id)
	} else {
		fmt.Printf("Group with ID %d has been successfully deleted!\n\n", id)
	}
}

func editGroup(db *gorm.DB, scanner *bufio.Scanner) {
	fmt.Print("\nEnter group ID: ")

	if !scanner.Scan() { // Scan the input
		return
	}

	idInput := strings.TrimSpace(scanner.Text())
	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("Incorrect group ID!")
		return
	}

	var group Group
	result := db.First(&group, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Printf("Group with ID %d wasn't found in DB.\n\n", id)
			return
		}
		fmt.Printf("Database error during search: %v\n", result.Error)
		return
	}

	// Updating the group

	var newGroup Group

	fmt.Print("\nEnter group name: ")
	if scanner.Scan() {
		newGroup.Name = strings.TrimSpace(scanner.Text())
	}
	result = db.Model(&group).Updates(newGroup) // Model = WHERE (UPDATE groups SET group_name = 'new name' WHERE id = 1)

	if result.Error != nil {
		fmt.Printf("Database error during update: %v\n", result.Error)
		return
	}

	fmt.Println("Group updated successfully!")
}

func selectAllStudents(db *gorm.DB) {
	var students []Student
	result := db.Find(&students)
	if result.Error != nil {
		log.Println("Query error:", result.Error)
		return
	}
	fmt.Printf("\n--- Student List ---\n")
	for _, student := range students {
		fmt.Printf("ID: %d\t Name: %s\t Lastname: %s\t Email: %s\t Group ID: %d\n", student.ID, student.FirstName, student.LastName, student.Email, student.GroupID)
	}
	fmt.Println()
}

func selectAllGroups(db *gorm.DB) {
	var groups []Group
	result := db.Find(&groups)
	if result.Error != nil {
		log.Println("Query error:", result.Error)
		return
	}
	fmt.Printf("\n--- Group List ---\n")
	for _, group := range groups {
		fmt.Printf("ID: %d\t Name: %s\n", group.ID, group.Name)
	}
	fmt.Println()
}
