-- 002.sql
-- TEMP VALUES
-- Inserts for students
INSERT INTO students (grNo, sPwd, userRole, studName, std, section) VALUES
(1, 'pass1234', 'student', 'Alice Johnson', 5, 'A'),
(2, 'pass2345', 'student', 'Bob Smith', 6, 'B'),
(3, 'pass3456', 'student', 'Carol Lee', 7, 'A'),
(4, 'pass4567', 'student', 'David Kim', 8, 'C'),
(5, 'pass5678', 'student', 'Eva Brown', 9, 'B'),
(6, 'pass6789', 'student', 'Frank Moore', 10, 'A'),
(7, 'pass7890', 'student', 'Grace Davis', 11, 'C'),
(8, 'pass8901', 'student', 'Henry Wilson', 12, 'B'),
(9, 'pass9012', 'student', 'Ivy Martinez', 4, 'A'),
(10, 'pass0123', 'student', 'Jack Taylor', 3, 'C');

-- Inserts for teachers
INSERT INTO teachers (tId, tPwd, userRole, tName, subjectAssigned, classAllocated) VALUES
('T001', 'teach001', 'faculty', 'Mr. Anderson', 'Mathematics', '5A'),
('T002', 'teach002', 'faculty', 'Ms. Baker', 'Science', '6B'),
('T003', 'teach003', 'faculty', 'Mr. Clark', 'English', '7A'),
('T004', 'teach004', 'faculty', 'Ms. Diaz', 'History', '8C'),
('T005', 'teach005', 'faculty', 'Mr. Evans', 'Geography', '9B'),
('T006', 'teach006', 'faculty', 'Ms. Foster', 'Physics', '10A'),
('T007', 'teach007', 'faculty', 'Mr. Green', 'Chemistry', '11C'),
('T008', 'teach008', 'faculty', 'Ms. Harris', 'Biology', '12B'),
('T009', 'teach009', 'faculty', 'Mr. Irving', 'Computer Sci', '4A'),
('T010', 'teach010', 'faculty', 'Ms. Jenkins', 'Art', '3C');

-- Inserts for subjects
INSERT INTO subjects (subId, subName, levelStd, credits) VALUES
(101, 'Mathematics', 5, 4),
(102, 'Science', 6, 3),
(103, 'English', 7, 3),
(104, 'History', 8, 2),
(105, 'Geography', 9, 2),
(106, 'Physics', 10, 4),
(107, 'Chemistry', 11, 4),
(108, 'Biology', 12, 3),
(109, 'Computer Science', 4, 3),
(110, 'Art', 3, 2);

-- Inserts for teachersSubInfo
INSERT INTO teachersSubInfo (tId, subId, std, section) VALUES
('T001', 101, '5','A'),
('T002', 102, '6','B'),
('T003', 103, '7','A'),
('T004', 104, '8','C'),
('T005', 105, '9','B'),
('T006', 106, '10','A'),
('T007', 107, '11','C'),
('T008', 108, '12','B'),
('T009', 109, '4','A'),
('T010', 110, '3','C');

-- Inserts for marks
INSERT INTO marks (grNo, subId, theoryM, practicalM, grade) VALUES
(1, 101, 75, 18, 'A'),
(2, 102, 68, 15, 'B'),
(3, 103, 70, 17, 'B'),
(4, 104, 60, 14, 'C'),
(5, 105, 72, 16, 'B'),
(6, 106, 78, 19, 'A'),
(7, 107, 65, 13, 'C'),
(8, 108, 69, 15, 'B'),
(9, 109, 80, 20, 'A'),
(10, 110, 55, 10, 'D');

-- Inserts for subjectAllocation
INSERT INTO subjectAllocation (std, markLimit) VALUES
(1, 5),
(2, 5),
(3, 5),
(4, 5),
(5, 5),
(6, 6),
(7, 6),
(8, 6),
(9, 7),
(10, 7);

-- Inserts for reviews
INSERT INTO reviews (tId, grNo, comment) VALUES
('T001', 1, 'Excellent progress'),
('T002', 2, 'Needs improvement in practicals'),
('T003', 3, 'Very active in class'),
('T004', 4, 'Struggles with theory'),
('T005', 5, 'Good understanding'),
('T006', 6, 'Shows great enthusiasm'),
('T007', 7, 'Must participate more'),
('T008', 8, 'Consistent performer'),
('T009', 9, 'Very creative'),
('T010', 10, 'Improved steadily');
