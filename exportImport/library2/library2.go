package library2

// nama struct harus besar
type Student struct {
	Name string // property exported, karena huruf besar
	//grade int    // property unexported, karena huruf kecil
	Grade int
}
