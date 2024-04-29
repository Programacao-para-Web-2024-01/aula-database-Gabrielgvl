CREATE
    DATABASE IF NOT EXISTS web;

USE
    web;

CREATE TABLE students
(
    id    INT PRIMARY KEY AUTO_INCREMENT,
    name  VARCHAR(255) NOT NULL,
    age   INT          NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(11)  NOT NULL
);

CREATE TABLE subjects
(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(255) NOT NULL,
    workload INT          NOT NULL
);

CREATE TABLE students_subjects
(
    student_id INT NOT NULL,
    subject_id INT NOT NULL,
    CONSTRAINT FOREIGN KEY (student_id) REFERENCES students (id),
    CONSTRAINT FOREIGN KEY (subject_id) REFERENCES subjects (id)
);

CREATE TABLE professors
(
    id    INT PRIMARY KEY AUTO_INCREMENT,
    name  VARCHAR(255) NOT NULL,
    age   INT          NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(11)  NOT NULL
);

CREATE TABLE professors_subjects
(
    professor_id INT NOT NULL,
    subject_id   INT NOT NULL,
    CONSTRAINT FOREIGN KEY (professor_id) REFERENCES professors (id),
    CONSTRAINT FOREIGN KEY (subject_id) REFERENCES subjects (id)
);

CREATE TABLE users
(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
