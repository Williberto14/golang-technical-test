package handlers

import (
	"database/sql"
	"golang-technical-test/src/common/common_configs"
	"golang-technical-test/src/common/utils"
	"golang-technical-test/src/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = common_configs.ConnectDB()

// Get all students
func GetStudents(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM Students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query students"})
		return
	}
	defer rows.Close()

	students := []models.Student{}
	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan student"})
			return
		}
		students = append(students, student)
	}

	if len(students) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No students found"})
		return
	}

	c.JSON(http.StatusOK, students)
}

// Get student by ID
func GetStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	err := db.QueryRow("SELECT * FROM Students WHERE student_id = ?", id).Scan(&student.StudentID, &student.FirstName, &student.LastName, &student.DateOfBirth)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No student found with given ID"})
			return
		}
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, student)
}

// Create student
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate that the fields are not empty
	if student.FirstName == "" || student.LastName == "" || student.DateOfBirth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields cannot be empty"})
		return
	}

	result, err := db.Exec("INSERT INTO Students (first_name, last_name, date_of_birth) VALUES (?, ?, ?)", student.FirstName, student.LastName, student.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve student ID"})
		return
	}

	student.StudentID = int(id)

	c.JSON(http.StatusCreated, student)
}

// Update student
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate that the fields are not empty
	if student.FirstName == "" || student.LastName == "" || student.DateOfBirth == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields cannot be empty"})
		return
	}

	_, err := db.Exec("UPDATE Students SET first_name = ?, last_name = ?, date_of_birth = ? WHERE student_id = ?", student.FirstName, student.LastName, student.DateOfBirth, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// Delete student
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	var studentID int
	err := db.QueryRow("SELECT student_id FROM Students WHERE student_id = ?", id).Scan(&studentID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "No student found with given ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query student"})
		return
	}

	_, err = db.Exec("DELETE FROM Students WHERE student_id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

// LoginHandler handles login
func LoginHandler(c *gin.Context) {
	var login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	c.BindJSON(&login)

	// Aquí deberías verificar el nombre de usuario y la contraseña con tu base de datos
	if login.Username == "admin" && login.Password == "password" {
		tokenString, _ := utils.CreateToken(login.Username)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
