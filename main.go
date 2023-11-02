// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "threejoin"
// )

// func main() {
// 	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)
// 	defer db.Close()

// 	err = db.Ping()
// 	CheckError(err)
// 	fmt.Println("db Connected successfully")

// 	dataToUpdate := []struct {
// 		Keydata    string
// 		NewKeydata string
// 		NewVal     string
// 	}{
// 		{"key1", "new_key11", "new_value1"},
// 		{"key2", "new_key11", "new_value1"},
// 		{"key3", "new_key11", "new_value1"},
// 		{"key4", "new_key11", "new_value1"},
// 		{"key5", "new_key11", "new_value1"},
// 		{"key6", "new_key11", "new_value1"},
// 		{"key7", "new_key11", "new_value1"},
// 		{"key8", "new_key11", "new_value1"},
// 		{"key9", "new_key11", "new_value1"},
// 		{"key10", "new_key11", "new_value1"},
// 		{"key11", "new_key11", "new_value1"},
// 		{"key12", "new_key11", "new_value1"},
// 		{"key13", "new_key11", "new_value1"},
// 		{"key14", "new_key11", "new_value1"},
// 		{"key15", "new_key11", "new_value1"},
// 		{"key16", "new_key11", "new_value1"},
// 		{"key17", "new_key11", "new_value1"},
// 		{"key18", "new_key11", "new_value1"},
// 		{"key19", "new_key11", "new_value1"},
// 		{"key1", "new_key11", "new_value1"},
// 		{"key20", "new_key11", "new_value1"},
// 		{"key21", "new_key11", "new_value1"},
// 		{"key22", "new_key11", "new_value1"},
// 		{"key23", "new_key11", "new_value1"},
// 		{"key24", "new_key11", "new_value1"},
// 		{"key25", "new_key11", "new_value1"},
// 		{"key26", "new_key11", "new_value1"},
// 		{"key27", "new_key11", "new_value1"},
// 		{"key28", "new_key11", "new_value1"},
// 		{"key29", "new_key11", "new_value1"},
// 		{"key30", "new_key11", "new_value1"},
// 		{"key31", "new_key11", "new_value1"},
// 		{"key32", "new_key11", "new_value1"},
// 		{"key33", "new_key11", "new_value1"},
// 		{"key34", "new_key11", "new_value1"},
// 		{"key35", "new_key11", "new_value1"},
// 		{"key36", "new_key11", "new_value1"},
// 		{"key37", "new_key11", "new_value1"},
// 		{"key38", "new_key11", "new_value1"},
// 		{"key39", "new_key11", "new_value1"},
// 		{"key40", "new_key11", "new_value1"},
// 		{"key41", "new_key11", "new_value1"},
// 		{"key42", "new_key11", "new_value1"},
// 		{"key43", "new_key11", "new_value1"},
// 		{"key44", "new_key11", "new_value1"},
// 		{"key45", "new_key11", "new_value1"},
// 		{"key46", "new_key11", "new_value1"},
// 		{"key47", "new_key11", "new_value1"},
// 		{"key48", "new_key11", "new_value1"},
// 		{"key49", "new_key11", "new_value1"},
// 		{"key50", "new_key11", "new_value1"},
// 		{"key51", "new_key11", "new_value1"},
// 		{"key52", "new_key11", "new_value1"},
// 		{"key53", "new_key11", "new_value1"},
// 		{"key54", "new_key11", "new_value1"},
// 		{"key55", "new_key11", "new_value1"},
// 		{"key56", "new_key11", "new_value1"},
// 		{"key57", "new_key11", "new_value1"},
// 		{"key58", "new_key11", "new_value1"},
// 		{"key59", "new_key11", "new_value1"},
// 		{"key60", "new_key11", "new_value1"},
// 		{"key61", "new_key11", "new_value1"},
// 		{"key62", "new_key11", "new_value1"},
// 		{"key63", "new_key11", "new_value1"},
// 		{"key64", "new_key11", "new_value1"},
// 		{"key65", "new_key11", "new_value1"},
// 		{"key66", "new_key11", "new_value1"},
// 		{"key67", "new_key11", "new_value1"},
// 		{"key68", "new_key11", "new_value1"},
// 		{"key69", "new_key11", "new_value1"},
// 		{"key70", "new_key11", "new_value1"},
// 		{"key71", "new_key11", "new_value1"},
// 		{"key72", "new_key11", "new_value1"},
// 		{"key73", "new_key11", "new_value1"},
// 		{"key74", "new_key11", "new_value1"},
// 		{"key75", "new_key11", "new_value1"},
// 		{"key76", "new_key11", "new_value1"},
// 		{"key77", "new_key11", "new_value1"},
// 		{"key78", "new_key11", "new_value1"},
// 		{"key79", "new_key11", "new_value1"},
// 		{"key80", "new_key11", "new_value1"},
// 		{"key81", "new_key11", "new_value1"},
// 		{"key82", "new_key11", "new_value1"},
// 		{"key83", "new_key11", "new_value1"},
// 		{"key84", "new_key11", "new_value1"},
// 		{"key85", "new_key11", "new_value1"},
// 		{"key86", "new_key11", "new_value1"},
// 		{"key87", "new_key11", "new_value1"},
// 		{"key88", "new_key11", "new_value1"},
// 		{"key89", "new_key11", "new_value1"},
// 		{"key90", "new_key11", "new_value1"},
// 		{"key91", "new_key11", "new_value1"},
// 		{"key92", "new_key11", "new_value1"},
// 		{"key93", "new_key11", "new_value1"},
// 		{"key94", "new_key11", "new_value1"},
// 		{"key95", "new_key11", "new_value1"},
// 		{"key96", "new_key11", "new_value1"},
// 		{"key97", "new_key11", "new_value1"},
// 		{"key98", "new_key11", "new_value1"},
// 		{"key99", "new_key11", "new_value1"},
// 		{"key100", "new_key11", "new_value1"},
// 	}

// 	done := make(chan bool)
// 	count := 0

// 	go func() {
// 		for i, data := range dataToUpdate {
// 			timer := time.NewTimer(1 * time.Second) // Set a timer for 1 second for each iteration
// 			select {
// 			case <-timer.C:
// 				fmt.Println("Timer expired. Stopping updates.")
// 				done <- true
// 				return
// 			default:
// 				err := UpdateData(db, data)
// 				if err != nil {
// 					fmt.Printf("Error updating data at index %d: %v\n", i, err)
// 				} else {
// 					count++
// 					fmt.Printf("Updated data at index %d\n", i)
// 				}
// 			}
// 		}
// 		done <- true
// 	}()

// 	<-done

// 	fmt.Printf("Total updates within 1 second: %d\n", count)
// }

// func UpdateData(db *sql.DB, data struct {
// 	Keydata    string
// 	NewKeydata string
// 	NewVal     string
// }) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	updateStmt := `UPDATE "godata" SET "keydata"=$1, "val"=$2 WHERE "keydata"=$3`

// 	_, err = tx.Exec(updateStmt, data.NewKeydata, data.NewVal, data.Keydata)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	return err
// }

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

//  100 times

// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "threejoin"
// )

// func main() {
// 	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", psqlconn)
// 	CheckError(err)
// 	defer db.Close()

// 	err = db.Ping()
// 	CheckError(err)
// 	fmt.Println("db Connected successfully")

// 	dataToUpdate := []struct {
// 		Keydata    string
// 		NewKeydata string
// 		NewVal     string
// 	}{
// 		{"key1", "new_key11", "new_value1"},
// 		{"key2", "new_key11", "new_value1"},
// 		{"key3", "new_key11", "new_value1"},
// 		{"key4", "new_key11", "new_value1"},
// 		{"key5", "new_key11", "new_value1"},
// 		{"key6", "new_key11", "new_value1"},
// 		{"key7", "new_key11", "new_value1"},
// 		{"key8", "new_key11", "new_value1"},
// 		{"key9", "new_key11", "new_value1"},
// 		{"key10", "new_key11", "new_value1"},
// 		{"key11", "new_key11", "new_value1"},
// 		{"key12", "new_key11", "new_value1"},
// 		{"key13", "new_key11", "new_value1"},
// 		{"key14", "new_key11", "new_value1"},
// 		{"key15", "new_key11", "new_value1"},
// 		{"key16", "new_key11", "new_value1"},
// 		{"key17", "new_key11", "new_value1"},
// 		{"key18", "new_key11", "new_value1"},
// 		{"key19", "new_key11", "new_value1"},
// 		{"key1", "new_key11", "new_value1"},
// 		{"key20", "new_key11", "new_value1"},
// 		{"key21", "new_key11", "new_value1"},
// 		{"key22", "new_key11", "new_value1"},
// 		{"key23", "new_key11", "new_value1"},
// 		{"key24", "new_key11", "new_value1"},
// 		{"key25", "new_key11", "new_value1"},
// 		{"key26", "new_key11", "new_value1"},
// 		{"key27", "new_key11", "new_value1"},
// 		{"key28", "new_key11", "new_value1"},
// 		{"key29", "new_key11", "new_value1"},
// 		{"key30", "new_key11", "new_value1"},
// 		{"key31", "new_key11", "new_value1"},
// 		{"key32", "new_key11", "new_value1"},
// 		{"key33", "new_key11", "new_value1"},
// 		{"key34", "new_key11", "new_value1"},
// 		{"key35", "new_key11", "new_value1"},
// 		{"key36", "new_key11", "new_value1"},
// 		{"key37", "new_key11", "new_value1"},
// 		{"key38", "new_key11", "new_value1"},
// 		{"key39", "new_key11", "new_value1"},
// 		{"key40", "new_key11", "new_value1"},
// 		{"key41", "new_key11", "new_value1"},
// 		{"key42", "new_key11", "new_value1"},
// 		{"key43", "new_key11", "new_value1"},
// 		{"key44", "new_key11", "new_value1"},
// 		{"key45", "new_key11", "new_value1"},
// 		{"key46", "new_key11", "new_value1"},
// 		{"key47", "new_key11", "new_value1"},
// 		{"key48", "new_key11", "new_value1"},
// 		{"key49", "new_key11", "new_value1"},
// 		{"key50", "new_key11", "new_value1"},
// 		{"key51", "new_key11", "new_value1"},
// 		{"key52", "new_key11", "new_value1"},
// 		{"key53", "new_key11", "new_value1"},
// 		{"key54", "new_key11", "new_value1"},
// 		{"key55", "new_key11", "new_value1"},
// 		{"key56", "new_key11", "new_value1"},
// 		{"key57", "new_key11", "new_value1"},
// 		{"key58", "new_key11", "new_value1"},
// 		{"key59", "new_key11", "new_value1"},
// 		{"key60", "new_key11", "new_value1"},
// 		{"key61", "new_key11", "new_value1"},
// 		{"key62", "new_key11", "new_value1"},
// 		{"key63", "new_key11", "new_value1"},
// 		{"key64", "new_key11", "new_value1"},
// 		{"key65", "new_key11", "new_value1"},
// 		{"key66", "new_key11", "new_value1"},
// 		{"key67", "new_key11", "new_value1"},
// 		{"key68", "new_key11", "new_value1"},
// 		{"key69", "new_key11", "new_value1"},
// 		{"key70", "new_key11", "new_value1"},
// 		{"key71", "new_key11", "new_value1"},
// 		{"key72", "new_key11", "new_value1"},
// 		{"key73", "new_key11", "new_value1"},
// 		{"key74", "new_key11", "new_value1"},
// 		{"key75", "new_key11", "new_value1"},
// 		{"key76", "new_key11", "new_value1"},
// 		{"key77", "new_key11", "new_value1"},
// 		{"key78", "new_key11", "new_value1"},
// 		{"key79", "new_key11", "new_value1"},
// 		{"key80", "new_key11", "new_value1"},
// 		{"key81", "new_key11", "new_value1"},
// 		{"key82", "new_key11", "new_value1"},
// 		{"key83", "new_key11", "new_value1"},
// 		{"key84", "new_key11", "new_value1"},
// 		{"key85", "new_key11", "new_value1"},
// 		{"key86", "new_key11", "new_value1"},
// 		{"key87", "new_key11", "new_value1"},
// 		{"key88", "new_key11", "new_value1"},
// 		{"key89", "new_key11", "new_value1"},
// 		{"key90", "new_key11", "new_value1"},
// 		{"key91", "new_key11", "new_value1"},
// 		{"key92", "new_key11", "new_value1"},
// 		{"key93", "new_key11", "new_value1"},
// 		{"key94", "new_key11", "new_value1"},
// 		{"key95", "new_key11", "new_value1"},
// 		{"key96", "new_key11", "new_value1"},
// 		{"key97", "new_key11", "new_value1"},
// 		{"key98", "new_key11", "new_value1"},
// 		{"key99", "new_key11", "new_value1"},
// 		{"key100", "new_key11", "new_value1"},
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
// 	defer cancel()

// 	done := make(chan bool)
// 	count := 0

// 	go func() {
// 		for i, data := range dataToUpdate {
// 			select {
// 			case <-ctx.Done():
// 				fmt.Println("Time limit reached. Stopping updates.")
// 				done <- true
// 				return
// 			default:
// 				err := UpdateData(db, data)
// 				if err != nil {
// 					fmt.Printf("Error updating data at index %d: %v\n", i, err)
// 				} else {
// 					count++
// 					fmt.Printf("Updated data at index %d\n", i)
// 				}
// 			}
// 		}
// 		done <- true
// 	}()

// 	<-done

// 	fmt.Printf("Total updates within 1 second: %d\n", count)
// }

// // Remaining code (UpdateData, CheckError, and the dataToUpdate struct) stays the same
// func UpdateData(db *sql.DB, data struct {
// 	Keydata    string
// 	NewKeydata string
// 	NewVal     string
// }) error {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return err
// 	}
// 	defer tx.Rollback()

// 	updateStmt := `UPDATE "godata" SET "keydata"=$1, "val"=$2 WHERE "keydata"=$3`

// 	_, err = tx.Exec(updateStmt, data.NewKeydata, data.NewVal, data.Keydata)
// 	if err != nil {
// 		return err
// 	}

// 	err = tx.Commit()
// 	return err
// }

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	startTime := time.Now()

// 	var a = time.Millisecond

// 	counter := 0

// 	for time.Since(startTime) < a*2 {
// 		counter++
// 	}

// 	fmt.Printf("Counter was incremented %d times in 1 ms.\n", counter)
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	// Define your PostgreSQL connection parameters
// 	psqlHost     = "localhost"
// 	psqlPort     = 5432
// 	psqlUser     = "postgres"
// 	psqlPassword = "postgres"
// 	psqlDB       = "threejoin"
// )

// func main() {
// 	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", psqlHost, psqlPort, psqlUser, psqlPassword, psqlDB)

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Create the table if it doesn't exist
// 	createTable := `
//         CREATE TABLE IF NOT EXISTS example_table (
//             id serial PRIMARY KEY,
//             data text
//         )
//     `
// 	_, err = db.Exec(createTable)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set a timer for 1 second
// 	timer := time.NewTimer(1 * time.Second)
// 	defer timer.Stop()

// 	updateCounter := 0

// 	for {
// 		select {
// 		case <-timer.C:
// 			// 1-second timer has elapsed, print the update counter and exit
// 			fmt.Printf("Updates in 1 second: %d\n", updateCounter)
// 			return
// 		default:
// 			// Perform an update operation on the table
// 			updateQuery := "UPDATE example_table SET data = 'updated' WHERE id = 1"
// 			_, err := db.Exec(updateQuery)
// 			if err != nil {
// 				log.Println("Update error:", err)
// 			} else {
// 				updateCounter++
// 			}
// 		}
// 	}
// }

// 2/Nov/2023

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	_ "github.com/lib/pq" // Import the PostgreSQL driver
// )

// func main() {

// 	dsn := "user=postgres password=postgres dbname=threejoin host=localhost port=5432 sslmode=disable"

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		// Update the "val" column by +1
// 		_, err := db.Exec("UPDATE datacount SET val = val + 1")
// 		if err != nil {
// 			fmt.Println("Error updating counter:", err)
// 			break
// 		}
// 		fmt.Println("Counter updated")
// 	}
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

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table and insert one initial record with a counter value of 0
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS counter_table (val INT)")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	_, err = db.Exec("INSERT INTO counter_table (val) VALUES (0)")
	if err != nil {
		log.Fatal(err)
	}

	// Measure the time it takes to increment the counter within 1 second

	counter := 0
	var timeSecond = time.Second
	fmt.Println("start Time: ", time.Now())
	startTime := time.Now()
	for time.Since(startTime) < timeSecond {
		_, err = db.Exec("UPDATE counter_table SET val = val + 1")
		if err != nil {
			log.Fatal(err)
		}
		counter++
	}
	fmt.Println("Completion Time: ", time.Now())
	fmt.Printf("Counter was incremented %d times in 1 second.\n", counter)
}
