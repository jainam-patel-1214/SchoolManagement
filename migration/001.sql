-- 001.sql
CREATE TABLE IF NOT EXISTS students (
    grNo int PRIMARY KEY,
    sPwd varchar(8) NOT NULL,
    userRole varchar(10) DEFAULT 'student',
    studName VARCHAR(50) NOT NULL,
    std int NOT NULL,
    section VARCHAR(2) NOT NULL,
    CONSTRAINT check_std_input CHECK (std BETWEEN 1 AND 12)
);
CREATE TABLE IF NOT EXISTS teachers (
    tId VARCHAR(8) PRIMARY KEY,
    tPwd varchar(8) NOT NULL,
    userRole varchar(10) DEFAULT 'faculty',
    tName VARCHAR(50) NOT NULL,
    subjectAssigned VARCHAR(25),
    classAllocated VARCHAR(4)
);
CREATE TABLE IF NOT EXISTS subjects (
    subId int PRIMARY KEY,
    subName VARCHAR(50) NOT NULL,
    levelStd int NOT NULL,
    credits int NOT NULL,
    CONSTRAINT check_level_input CHECK (levelStd BETWEEN 1 AND 12)
);
CREATE TABLE IF NOT EXISTS teachersSubInfo (
    tId VARCHAR(8) NOT NULL,
    subId int NOT NULL,
    std int NOT NULL,
    section VARCHAR(2) NOT NULL,
    FOREIGN KEY (tId) REFERENCES teachers(tId),
    FOREIGN KEY (subId) REFERENCES subjects(subId)
);
CREATE TABLE IF NOT EXISTS marks (
    grNo int NOT NULL,
    subId int NOT NULL, 
    theoryM int,
    practicalM int, 
    grade varchar(2), 
    FOREIGN KEY (grNo) REFERENCES students(grNo), 
    FOREIGN KEY (subId) REFERENCES subjects(subId),
    CONSTRAINT check_theory_marks_input CHECK (theoryM BETWEEN 0 AND 80),
    CONSTRAINT check_practical_marks_input CHECK (practicalM BETWEEN 0 AND 20)
);
CREATE TABLE IF NOT EXISTS subjectAllocation (
    std int NOT NULL,
    markLimit int NOT NULL
);
CREATE TABLE IF NOT EXISTS reviews (
    tId VARCHAR(8) NOT NULL,
    grNo int NOT NULL,
    comment VARCHAR(255),
    FOREIGN KEY (tId) REFERENCES teachers(tId),
    FOREIGN KEY (grNo) REFERENCES students(grNo)
);
CREATE TABLE IF NOT EXISTS activeSessions(
    sessionId int NOT NULL AUTO_INCREMENT,
    sessiontoken TEXT NOT NULL,
    userRole VARCHAR(10) NOT NULL,
    validtime DATETIME NOT NULL,
    PRIMARY KEY(sessionId)
)