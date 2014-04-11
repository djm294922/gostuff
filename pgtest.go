package main
    import "fmt"
    import "log"
    import "reflect"
	import (
		_ "github.com/lib/pq"
		"database/sql"
	)


func main() {



    fmt.Printf("pgtest starts")

	db, err := sql.Open("postgres", 
		"sslmode=disable user=5347fd143d77cee941000017 host=127.0.0.1 port=5432 dbname=template1")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM pg_stat_database")

	if rows != nil {
		cols,err2 := rows.Columns()
		if (err2 == nil) {
		if cols != nil {
			for i := range cols {
			fmt.Printf("cols were %s ", cols[i])
			}
		}
		}

		fmt.Printf("-----\n");

	vals := make([]interface{}, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
    	err = rows.Scan(vals...)

		for i:=0; i < 3; i++  {
			fmt.Printf("col= %s ", vals[i])
			typ := reflect.TypeOf(vals[i])
			typv := reflect.ValueOf(vals[i])
			fmt.Printf("interface=%s ", typv.Interface())
			fmt.Printf("ValueOf=%s ", typv)
			fmt.Printf("typeof=%s ", typ)
			fmt.Printf("typeKinid=%s\n ", typ.Kind().String())
		}
	}	
	} else {
		println("rows were nil");
	}
}
