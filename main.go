// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	dsn := "root:admin123@tcp(127.0.0.1:3306)/goLearn"
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("Error opening DB: ", err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Error connecting to DB: ", err)
// 	}
// 	fmt.Println("Connected to MySQL ✅")

// 	// if _, err := db.Exec("CREATE TABLE moreInfo(City VARCHAR(255), Country VARCHAR(255))"); err != nil {
// 	// 	fmt.Printf("error creating table")
// 	// } else {
// 	// 	fmt.Println("successfully created")
// 	// }

// 	// if _, err := db.Exec(`INSERT INTO moreInfo (City,COUNTRY) VALUES ("Bermuda","Triangle")`); err != nil {
// 	// 	fmt.Printf("%s", err.Error())
// 	// } else {
// 	// 	fmt.Println("successfully inserted")
// 	// }

// 	//using LIMIT TO DECIDE NUMBER OF RESULTS TO BE DELETED
// 	// if _, err := db.Exec(`DELETE FROM moreInfo WHERE City="Atlantis" LIMIT 1`); err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	fmt.Println("deleted successfully")
// 	// }

// 	// using in backend from database
// 	// rows, err := db.Query("SELECT * FROM moreInfo WHERE Country IS NOT NULL")
// 	// if err != nil {
// 	// 	log.Fatal("Query failed:", err)
// 	// }
// 	// defer rows.Close()
// 	// for rows.Next() {
// 	// 	var City, Country string
// 	// 	err := rows.Scan(&City, &Country)
// 	// 	if err != nil {
// 	// 		log.Fatal("Row scan failed:", err)
// 	// 	}
// 	// 	fmt.Printf("%s, %s\n", City, Country)
// 	// }

// 	//updating
// 	// if RES, err := db.Exec(`UPDATE moreInfo SET Country="unknown" WHERE Country IS NULL`); err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	fmt.Println(RES, "UPDATED SUCCESSFULLY")
// 	// }

// 	//aggregations
// 	// if res, err := db.Query(`SELECT COUNT(PersonID) as amount FROM Persons WHERE City="Mumbai"`); err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	for res.Next() {
// 	// 		var op int
// 	// 		if err := res.Scan(&op); err != nil {
// 	// 			fmt.Println(err)
// 	// 		}
// 	// 		fmt.Printf("%d", op)
// 	// 	}
// 	// }

// 	// QUERY return multiple rows and if there is 0 output then it is still valid, use manual way tojust notify yourself, QUERYROW returns single row
// 	// rows, err := db.Query(`SELECT COUNT(PersonID) AS totalNum FROM Persons WHERE City IN (SELECT City FROM moreInfo WHERE Country="India" OR City="New York")`)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer rows.Close()

// 	// found := false
// 	// for rows.Next() {
// 	// 	var out int
// 	// 	if err := rows.Scan(&out); err != nil {
// 	// 		log.Fatal(err)
// 	// 	}
// 	// 	fmt.Println(out, "output")
// 	// 	found = true
// 	// }
// 	// if !found {
// 	// 	fmt.Println("No rows found")
// 	// }

// 	//joins
// 	// if res, err := db.Query(`SELECT FirstName, Max(PersonID) FROM Persons JOIN moreInfo ON moreInfo.City=Persons.city GROUP BY Persons.FirstName`); err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	for res.Next() {
// 	// 		var name, addr string
// 	// 		if err := res.Scan(&name, &addr); err != nil {
// 	// 			fmt.Println("error occured while scanning")
// 	// 		} else {
// 	// 			fmt.Println(name, "-", addr)
// 	// 		}
// 	// 	}
// 	// }
// 	// if res, err := db.Query(`SELECT FirstName, Persons.City FROM Persons JOIN moreInfo ON moreInfo.City=Persons.city WHERE Persons.City="New York"`); err != nil {
// 	// 	fmt.Println(err)
// 	// } else {
// 	// 	for res.Next() {
// 	// 		var name, addr string
// 	// 		if err := res.Scan(&name, &addr); err != nil {
// 	// 			fmt.Println("error occured while scanning")
// 	// 		} else {
// 	// 			fmt.Println(name, "-", addr, "-")
// 	// 		}
// 	// 	}
// 	// }

// 	// if _, err := db.Exec(`CREATE TABLE temp AS SELECT City,Country FROM moreInfo`); err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	// if _, err := db.Exec(`DELETE FROM temp`); err != nil {
// 	// 	fmt.Println(err)
// 	// }
// 	if _, err := db.Exec(`INSERT INTO temp (City,Country) SELECT Persons.City, moreInfo.Country FROM Persons JOIN moreInfo ON Persons.City = moreInfo.CitY `); err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("done")
// 	}

// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"sort"

// 	_ "github.com/lib/pq" // PostgreSQL driver
// )

// const (
// 	dbHost     = "localhost"
// 	dbPort     = 5432
// 	dbUser     = "postgres"
// 	dbPassword = "yourpassword"
// 	dbName     = "yourdbname"
// )

// func main() {
// 	// Construct connection string
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		dbHost, dbPort, dbUser, dbPassword, dbName)

// 	// Connect to the database
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatalf("Unable to connect to database: %v", err)
// 	}
// 	defer db.Close()

// 	// Verify connection
// 	if err := db.Ping(); err != nil {
// 		log.Fatalf("Unable to ping database: %v", err)
// 	}

// 	fmt.Println("✅ Connected to the database.")

// 	// Run migrations
// 	err = runMigrations(db, "./migrations")
// 	if err != nil {
// 		log.Fatalf("Migration failed: %v", err)
// 	}

// 	fmt.Println("✅ Migrations completed.")
// }

// func runMigrations(db *sql.DB, migrationsDir string) error {
// 	files, err := ioutil.ReadDir(migrationsDir)
// 	if err != nil {
// 		return fmt.Errorf("failed to read migrations directory: %w", err)
// 	}

// 	var sqlFiles []string
// 	for _, f := range files {
// 		if filepath.Ext(f.Name()) == ".sql" {
// 			sqlFiles = append(sqlFiles, f.Name())
// 		}
// 	}

// 	sort.Strings(sqlFiles) // Ensure order (e.g., 001_*, 002_*, etc.)

// 	for _, file := range sqlFiles {
// 		path := filepath.Join(migrationsDir, file)
// 		content, err := os.ReadFile(path)
// 		if err != nil {
// 			return fmt.Errorf("failed to read file %s: %w", file, err)
// 		}

// 		fmt.Printf("➡️ Running migration: %s\n", file)
// 		if _, err := db.Exec(string(content)); err != nil {
// 			return fmt.Errorf("failed to execute %s: %w", file, err)
// 		}
// 	}

// 	return nil
// }

package main

import (
	"database/sql"

	// "encoding/json"
	"fmt"
	// "io"
	"log"
	// "net/http"

	"example.com/main/migration"
	"example.com/main/routes"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:admin123@tcp(127.0.0.1:3306)/goLearn?multiStatements=true"

// type ReturnMsg struct {
// 	Code    int    `json:"statusCode" binding:"required"`
// 	Status  string `json:"status" binding:"required"`
// 	Message string `json:"response"`
// }

// type StudentInfo struct {
// 	RollNo  int    `json:"rollNo" binding:"required"`
// 	Name    string `json:"name" binding:"required"`
// 	Section string `json:"section" binding:"required"`
// }

// func sendJSONResponse(writer http.ResponseWriter, code int, status, message string) {
// 	writer.WriteHeader(code)
// 	response := ReturnMsg{Code: code, Status: status, Message: message}
// 	b, err := json.Marshal(response)
// 	if err != nil {
// 		fmt.Println("JSON Marshal error:", err)
// 		return
// 	}
// 	writer.Write(b)
// }

// func handleSubmitStudent(writer http.ResponseWriter, reader *http.Request) {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("Error opening DB: ", err)
// 	}
// 	defer db.Close()
// 	res, err := db.Query("SELECT * FROM information_schema.tables WHERE table_schema='goLearn' AND table_name='students'")
// 	if err != nil {
// 		fmt.Println("database not defined")
// 		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT FIND DATABASE")
// 		return
// 	}
// 	if !res.Next() {
// 		_, err = db.Exec("CREATE TABLE students (RollNo int NOT NULL UNIQUE, Name varchar(255), Section varchar(4), PRIMARY KEY(RollNo))")
// 		if err != nil {
// 			fmt.Println("ISSUE WHILE CREATING TABLE")
// 			sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CREATING TABLE")
// 			return
// 		}
// 		fmt.Println("table created")
// 	}
// 	defer res.Close()

// 	body, err := io.ReadAll(reader.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		var fail = ReturnMsg{Code: 404, Status: "Not found", Message: "BODY UNREADABLE"}
// 		b, err1 := json.Marshal(fail)
// 		if err1 != nil {
// 			fmt.Println(err1)
// 		}
// 		writer.Write(b)
// 		return
// 	}
// 	var data StudentInfo
// 	if err = json.Unmarshal(body, &data); err != nil {
// 		fmt.Println(err)
// 		var fail = ReturnMsg{Code: 404, Status: "Not found", Message: "REQUIRED FIELDS EMPTY"}
// 		b, err1 := json.Marshal(fail)
// 		if err1 != nil {
// 			fmt.Println(err1)
// 		}
// 		writer.Write(b)
// 		return
// 	}
// 	transisiton, err := db.Begin()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	_, err = db.Exec(`INSERT INTO students (RollNo, Name, Section, ScienceMark, EnglishMark, ClassID) VALUES (?,?,?)`, data.RollNo, data.Name, data.Section)
// 	if err != nil {
// 		fmt.Println(err)
// 		transisiton.Rollback()
// 	}
// 	err = transisiton.Commit()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("commited and saved successfully")
// }

func main() {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB: ", err)
	}
	defer db.Close()
	migration.RunMigrations(db, "./migration")
	// start := func(w http.ResponseWriter, _ *http.Request) {
	// 	if _, err := w.Write([]byte("welcome to server")); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	router := routes.InitializeRouter()
	router.Run("localhost:8090")

	// mux := http.NewServeMux()
	// mux.HandleFunc("/", start)
	// mux.HandleFunc("/addFaculty", teacher.HandleSubmitFaculty)
	// mux.HandleFunc("/addStudent", handleSubmitStudent)
	// mux.HandleFunc("/addSubject", subjects.AddSubjectInfo)

	// 	handler := cors.New(cors.Options{
	// 		AllowedOrigins: []string{"http://localhost:5173"},
	// 	}).Handler(mux)
	// for later use

	fmt.Println("Server running at port 6969")
	// http.ListenAndServe(":6969", mux)
}
