// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	dbHost     = "localhost"
// 	dbPort     = 5432
// 	dbUser     = "postgres"
// 	dbPassword = "postgres"
// 	dbName     = "threejoin"
// )

// func main() {
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Create the table and insert one initial record with a counter value of 0
// 	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS counter_table (val INT)")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	_, err = db.Exec("INSERT INTO counter_table (val) VALUES (0)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Measure the time it takes to increment the counter within 1 second

// 	counter := 0
// 	var timeSecond = time.Second
// 	fmt.Println("start Time: ", time.Now())
// 	startTime := time.Now()
// 	for time.Since(startTime) < timeSecond {
// 		_, err = db.Exec("UPDATE counter_table SET val = val + 1")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		counter++
// 	}
// 	fmt.Println("Completion Time: ", time.Now())
// 	fmt.Printf("Counter was incremented %d times in 1 second.\n", counter)
// }

//-----------------------------------------------
//package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	dbHost     = "localhost"
// 	dbPort     = 5432
// 	dbUser     = "postgres"
// 	dbPassword = "postgres"
// 	dbName     = "threejoin"
// )

// const numGoroutines = 10 // Number of concurrent goroutines

// func main() {
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	var counter int

// 	// Create the table and insert one initial record with a counter value of 0
// 	_, err = db.Exec("INSERT INTO counter_table (val) VALUES (0) ON CONFLICT DO NOTHING")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Channel to receive completion signals from goroutines
// 	done := make(chan bool)

// 	// Start multiple goroutines to increment the counter concurrently
// 	for i := 0; i < numGoroutines; i++ {
// 		fmt.Println("start Time: ", time.Now())
// 		go func() {

// 			var timeSecond = time.Second

// 			startTime := time.Now()

// 			for time.Since(startTime) < timeSecond {
// 				_, err = db.Exec("UPDATE counter_table SET val = val + 1")
// 				if err != nil {
// 					log.Fatal(err)
// 				}
// 				counter++
// 			}
// 			done <- true
// 			fmt.Printf("Goroutine %d completed. Incremented counter %d times.\n", i, counter)
// 		}()
// 	}

// 	// Wait for all goroutines to complete
// 	for i := 0; i < numGoroutines; i++ {
// 		<-done
// 	}

//		fmt.Printf("total counterd ", counter)
//		fmt.Println("Completion Time: ", time.Now())
//	}
//

// -------------------------------------------//

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	dbHost     = "localhost"
// 	dbPort     = 5432
// 	dbUser     = "postgres"
// 	dbPassword = "postgres"
// 	dbName     = "threejoin"
// )

// const numGoroutines = 4 // Number of concurrent goroutines

// var counter int // Global counter variable

// func main() {
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Create the table and insert one initial record with a counter value of 0
// 	_, err = db.Exec("CREATE TABLE IF NOT EXISTS counter_table (val INT)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create the table and insert one initial record with a counter value of 0
// 	_, err = db.Exec("INSERT INTO counter_table (val) VALUES (0) ON CONFLICT DO NOTHING")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Channel to receive completion signals from goroutines
// 	done := make(chan bool)

// 	// Start multiple goroutines to increment the counter concurrently
// 	for i := 0; i < numGoroutines; i++ {
// 		fmt.Println("staring time ", time.Now())
// 		go func() {
// 			var timeSecond = time.Second
// 			startTime := time.Now()
// 			for time.Since(startTime) < timeSecond {
// 				_, err = db.Exec("UPDATE counter_table SET val = val + 1")
// 				if err != nil {

// 					log.Fatal(err)
// 				}
// 				counter++
// 			}
// 			done <- true
// 			fmt.Printf("Goroutine %d completed. Incremented counter %d times.\n", i, counter)
// 		}()
// 	}

// 	// Wait for all goroutines to complete
// 	for i := 0; i < numGoroutines; i++ {
// 		<-done
// 	}

// 	fmt.Println("Completion Time: ", time.Now())
// 	fmt.Printf("Final Counter Value: %d\n", counter)
// }

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "threejoin"
)

const numGoroutines = 4 // Number of concurrent goroutines

var counter int // Global counter variable

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table and insert one initial record with a counter value of 0
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS counter_table (val INT)")
	if err != nil {
		log.Fatal(err)
	}

	// Create the table and insert one initial record with a counter value of 0
	_, err = db.Exec("INSERT INTO counter_table (val) VALUES (0) ON CONFLICT DO NOTHING")
	if err != nil {
		log.Fatal(err)
	}

	// Channel to receive completion signals from goroutines
	done := make(chan bool)

	// Start multiple goroutines to increment the counter concurrently
	for i := 0; i < numGoroutines; i++ {
		fmt.Println("staring time ", time.Now())
		go func() {
			var timeSecond = time.Second
			startTime := time.Now()
			for time.Since(startTime) < timeSecond {
				_, err = db.Exec("UPDATE counter_table SET val = val + 1")
				if err != nil {

					log.Fatal(err)
				}
				counter++
			}
			done <- true
			fmt.Printf("Goroutine %d completed. Incremented counter %d times.\n", i, counter)
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	fmt.Println("Completion Time: ", time.Now())
	fmt.Printf("Final Counter Value: %d\n", counter)
}
