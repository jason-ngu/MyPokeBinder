package utilities

import (
	"fmt"
	"reflect"
	"strings"
)

const offsetLimitStr = " OFFSET %d LIMIT %d"

func FormatStringForDatabase(str string) string {
	formattedStr := strings.ReplaceAll(str, "'", "''")
	return formattedStr
}

func FormatSqlQueryWithSearchParams(sqlQuery string, searchParams interface{}, addOffsetLimit bool) string {
	var formattedSqlQuery string
	var paramFilters []string

	v := reflect.ValueOf(searchParams)
	for i := 0; i < v.NumField(); i++ {
		dbTagValue := v.Type().Field(i).Tag.Get("db")
		fieldValue := v.Field(i).Interface()

		if fieldValue != "" {
			paramFilters = append(paramFilters, fmt.Sprintf("%s = :%s", dbTagValue, dbTagValue))
		}
	}
	if len(paramFilters) > 0 {
		formattedSqlQuery += " WHERE " + strings.Join(paramFilters, " AND ")
	}

	if addOffsetLimit {
		formattedSqlQuery += offsetLimitStr
	}

	return sqlQuery + formattedSqlQuery
}
