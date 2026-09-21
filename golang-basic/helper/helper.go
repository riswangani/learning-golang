package helper

var version = "1.0.0"
var Application = "Belajar Golang"

func SayHello(name string) string {
	return "Hello " + name
}

// gabisa di akses dari luar package karena hurup kecil awalnya
func sayGoodBye(name string) string {
	return "Good Bye " + name
}