-- define a select statement to get all students enrolled in a course
SELECT students.student_id, students.first_name, students.last_name, students.phone, students.email, students.city, students.state, students.zip_code
FROM courses.students
INNER JOIN courses.registrations ON students.student_id = registrations.student_id
INNER JOIN courses.registration_items ON registrations.registration_id = registration_items.registration_id
WHERE registration_items.course_id = 1;

-- write an index to improve the performance of the query
CREATE INDEX idx_course_id ON courses.registration_items (course_id);

-- define a table for student attendance to capture attendance by class
CREATE TABLE courses.attendance (
    attendance_id INT IDENTITY (1, 1) PRIMARY KEY,
    registration_id INT NOT NULL,
    attendance_date DATE NOT NULL,
    attendance_status tinyint NOT NULL,
    -- Attendance status: 1 = Present; 2 = Absent; 3 = Late; 4 = Excused
    FOREIGN KEY (registration_id) REFERENCES courses.registrations (registration_id) ON DELETE CASCADE ON UPDATE CASCADE
);

--
CREATE PROCEDURE GetCourseEnrollmentByLocation
    @CourseID INT
AS
BEGIN
    SELECT students.city, students.state, COUNT(*) AS enrollment_count
    FROM courses.students
    INNER JOIN courses.registrations ON students.student_id = registrations.student_id
    INNER JOIN courses.registration_items ON registrations.registration_id = registration_items.registration_id
    WHERE registration_items.course_id = @CourseID
    GROUP BY students.city, students.state
    ORDER BY students.state, students.city;
END;

--
SELECT * 
FROM courses.registrations 
WHERE registration_date >= '2023-09-01' 
  AND registration_date < '2023-10-01';
