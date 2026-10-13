package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

// func TestExecSql(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	script := "INSERT INTO customer(id, name) VALUES('gani', 'Riswan Gani')"
// 	_,err := db.ExecContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}

//		fmt.Println("Success insert new customer")
//	}
// func TestQuerySql(t *testing.T) {
// 	db := GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()
// 	script := "SELECT id, name FROM customer"
// 	rows, err := db.QueryContext(ctx, script)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var id, name string
// 		err := rows.Scan(&id, &name)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println("Id:", id)
// 		fmt.Println("Name:", name)

// 	}
// 	if err := rows.Err(); err != nil {
// 		panic(err)
// 	}
// }


func TestQueryComplexSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT id, name, email, balance,rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool
		
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		fmt.Println("Email:", email)
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birthdate:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("CreatedAt:", createdAt)

	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
}
