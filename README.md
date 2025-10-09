# register api

- api -> post to <http://localhost:8090/register> body - user, pwd and role you want  
- JSON TAGS REQUIRED - (roleReq,userName,pwd)

# login api

- api -> post to <http://localhost:8090/login> body - userid accordingly and pwd
- JSON TAGS REQUIRED - (userId,password) "after this copy the token string which is returned in output and send as  **COOKIE -> userCookie="token string you got"**"

# STUDENTS API

## Display GET to <http://localhost:8090/student/display>

- it displays other students data according to information required
- JSON TAGS (using all isnt compulsory) - (viewByStd,viewBySection,minPercent,maxPercent)

## display list of subjects GET to <http://localhost:8090/student/displaySub>

- it displays subjects in particular standard
- JSON TAGS - (std)

## display report of logged in student GET to <http://localhost:8090/student/report>

- it displays report of that student who is logged in (token string jiski ho uska report)
- JSON TAGS - nothing, just send token in userCookie

# TEACHERS API

## GET to <http://localhost:8090/teacher/displayPerformance>

- it displays performance of teacher logged in + performance of other teachers who are in same std assigned as logged in teacher eg- all teachers of standard x are displayed with total marks of students
- JSON TAGS - (std)

## POST to <http://localhost:8090/teacher/createStud>

- it created student
- JSON TAGS - (grNo,studPwd, userRole,studName std, section)

## PUT to <http://localhost:8090/teacher/updateStud>

- it updates student
- JSON TAGS - (grNo)required,  optional tags (studPwd, userRole,studName std, section)

## POST to <http://localhost:8090/teacher/createSub>

- it create subject
- JSON TAGS - (subId,subName,levelStd,credits)

## PUT to <http://localhost:8090/teacher/updateSub>

- it updates subject
- JSON TAGS - (subId)required , optional tags (subName,levelStd,credits)

## POST to <http://localhost:8090/teacher/enterMarks>

- it help to enter mark for particular student and respective subject
- JSON TAGS - (grNo, subId,theoryMarks,practicalMarks) all required

## PUT to <http://localhost:8090/teacher/updateMarks>

- it help to update mark for particular student and respective subject
- JSON TAGS - (grNo, subId) required ones , optional ones (theoryMarks,practicalMarks) use anyone or both or none

## GET to <http://localhost:8090/teacher/displaySub>

- it displays subjects in subjects
- JSON TAGS - (std)

## POST to <http://localhost:8090/teacher/addReview>

- it help to add review for student
- JSON TAGS - (grNo,comment)

## DELETE to <http://localhost:8090/teacher/delSubject>

- it help to DELETE SUBJECTS
- JSON TAGS - (subid)

## DELETE to <http://localhost:8090/teacher/delStudent>

- it help to DELETE students
- JSON TAGS - (grNo)

# TEACHERS API

## GET to <http://localhost:8090/admin/pendingRequest>

- it displays registeration requests from register
- JSON TAGS - ()nothing just hit the api

## POST to <http://localhost:8090/admin/acceptRequest>

- it displays registeration requests from register
- JSON TAGS - (uName,uPwd,uRole)required ones,  tags filled according to role (std,section,subId)

## DELETE to <http://localhost:8090/admin/rejectRequest>

- it displays registeration requests from register
- JSON TAGS - (uName,uPwd,uRole)required ones,  tags filled according to role (std,section,subId)

## DELETE to <http://localhost:8090/admin/delTeacher>

- it deleted teacher out of existence
- JSON TAGS - (teacherId)