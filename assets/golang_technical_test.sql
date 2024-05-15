CREATE DATABASE IF NOT EXISTS golang_technical_test;
USE golang_technical_test;

-- Professors Table
CREATE TABLE Professors (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Lastname VARCHAR(255),
    Email VARCHAR(255),
    Specialization VARCHAR(255)
);

-- Courses Table
CREATE TABLE Courses (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Description TEXT
);

-- ProfessorCourse Table (relationship between Professors and Courses)
CREATE TABLE ProfessorCourse (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    ProfessorID INT,
    CourseID INT,
    FOREIGN KEY (ProfessorID) REFERENCES Professors(ID),
    FOREIGN KEY (CourseID) REFERENCES Courses(ID)
);

-- Students Table
CREATE TABLE Students (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255),
    Lastname VARCHAR(255),
    DateOfBirth DATE,
    Address VARCHAR(255),
    Email VARCHAR(255)
);

-- Enrollment Table (relationship between Students and Courses)
CREATE TABLE Enrollment (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    StudentID INT,
    CourseID INT,
    FOREIGN KEY (StudentID) REFERENCES Students(ID),
    FOREIGN KEY (CourseID) REFERENCES Courses(ID)
);

-- Grades Table
CREATE TABLE Grades (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    StudentID INT,
    CourseID INT,
    ProfessorID INT,
    Grade DECIMAL(5,2),
    FOREIGN KEY (StudentID) REFERENCES Students(ID),
    FOREIGN KEY (CourseID) REFERENCES Courses(ID),
    FOREIGN KEY (ProfessorID) REFERENCES Professors(ID)
);
