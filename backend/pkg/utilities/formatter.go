package utilities

import "strings"

func FormatStringForDatabase(str string) string {
	formattedStr := strings.ReplaceAll(str, "'", "''")
	return formattedStr
}
