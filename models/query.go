package models

import (
	"database/sql"
	"fmt"
	"strings"
)

func Select(db *sql.DB, query string) ([]map[string]string, error) {
	theCase := "lower" // "lower", "upper", "camel" or the orignal case if this is anything other than these three

	// data []map[string]string, error
	//data1, err := gosqljson.QueryDbToMap(db, theCase, query)
	data1, err := QueryDbToMap(db, theCase, query)

	return data1, err
}

// QueryDbToMap - run sql and return an array of maps
func QueryDbToMap(db *sql.DB, theCase string, sqlStatement string, sqlParams ...interface{}) ([]map[string]string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	results := []map[string]string{}
	rows, err := db.Query(sqlStatement, sqlParams...)
	if err != nil {
		fmt.Println("Error executing: ", sqlStatement)
		return results, err
	}
	cols, _ := rows.Columns()
	colsLower := make([]string, len(cols))
	colsCamel := make([]string, len(cols))

	if theCase == "lower" {
		for i, v := range cols {
			colsLower[i] = strings.ToLower(v)
		}
	} else if theCase == "upper" {
		for i, v := range cols {
			cols[i] = strings.ToUpper(v)
		}
	} else if theCase == "camel" {
		for i, v := range cols {
			colsCamel[i] = toCamel(v)
		}
	}

	rawResult := make([][]byte, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		result := make(map[string]string, len(cols))
		rows.Scan(dest...)
		for i, raw := range rawResult {
			if raw == nil {
				if theCase == "lower" {
					result[colsLower[i]] = ""
				} else if theCase == "upper" {
					result[cols[i]] = ""
				} else if theCase == "camel" {
					result[colsCamel[i]] = ""
				} else {
					result[cols[i]] = ""
				}
			} else {
				if theCase == "lower" {
					result[colsLower[i]] = string(raw)
				} else if theCase == "upper" {
					result[cols[i]] = string(raw)
				} else if theCase == "camel" {
					result[colsCamel[i]] = string(raw)
				} else {
					result[cols[i]] = string(raw)
				}
			}
		}
		results = append(results, result)
	}
	return results, nil
}

func toCamel(s string) (ret string) {
	s = strings.ToLower(s)
	a := strings.Split(s, "_")
	for i, v := range a {
		if i == 0 {
			ret += v
		} else {
			f := strings.ToUpper(string(v[0]))
			n := string(v[1:])
			ret += fmt.Sprint(f, n)
		}
	}
	return
}
