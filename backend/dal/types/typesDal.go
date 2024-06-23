package typesDal

// https://dakaii.medium.com/repository-pattern-in-golang-d22d3fa76d91
// https://www.alexedwards.net/blog/organising-database-access

type Types struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
}

func searchTypes() {

}

func getTypeById(id int) Types {

}

func createType(typeName string) {

}
