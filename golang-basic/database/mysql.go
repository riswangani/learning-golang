package database

var connection string


// bakal otomatis di eksekusi karena ini pake kata "init"
func init() {
	connection = "Gani Connec masql"
}

func GetConnection() string {
	return connection
}