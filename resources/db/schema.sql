CREATE
    DATABASE IF NOT EXISTS web;

USE
    web;

CREATE TABLE IF NOT EXISTS students
(
    id    INT PRIMARY KEY AUTO_INCREMENT,
    name  VARCHAR(255) NOT NULL,
    age   INT          NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(11)  NOT NULL
);

CREATE TABLE IF NOT EXISTS subjects
(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    name     VARCHAR(255) NOT NULL,
    workload INT          NOT NULL
);

CREATE TABLE IF NOT EXISTS students_subjects
(
    student_id INT NOT NULL,
    subject_id INT NOT NULL,
    CONSTRAINT FOREIGN KEY (student_id) REFERENCES students (id),
    CONSTRAINT FOREIGN KEY (subject_id) REFERENCES subjects (id),
    CONSTRAINT UNIQUE KEY (student_id, subject_id)
);

CREATE TABLE IF NOT EXISTS professors
(
    id    INT PRIMARY KEY AUTO_INCREMENT,
    name  VARCHAR(255) NOT NULL,
    age   INT          NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(11)  NOT NULL
);

CREATE TABLE IF NOT EXISTS professors_subjects
(
    professor_id INT NOT NULL,
    subject_id   INT NOT NULL,
    CONSTRAINT FOREIGN KEY (professor_id) REFERENCES professors (id),
    CONSTRAINT FOREIGN KEY (subject_id) REFERENCES subjects (id),
    CONSTRAINT UNIQUE (professor_id, subject_id)
);

CREATE TABLE IF NOT EXISTS users
(
    id       INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);
