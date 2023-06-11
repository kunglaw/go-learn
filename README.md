
# Belajar Golang by Noval Agung 


list bab yang sudah di pelajari 

- goroutine
seperti asynchronous dalam javascript . sebuah mini threading. [See More](https://dasarpemrogramangolang.novalagung.com/A-goroutine.html)

- channel 
digunakan untuk menghubungkan goroutine satu dengan goroutine lain. [See more](https://dasarpemrogramangolang.novalagung.com/A-channel.html)

### 1. Conversion in Golang 

a. tipe konversi data eksplisit 
```go
var x int = 10
var y float64 = float64(x) // Konversi int ke float64

var a int = 5
var b int64 = int64(a) // Konversi int ke int64

var s string = "123"
var i int = int(s) 
// error , tidak bisa langsung conver ke integer
```
b. konversi tipe data yang disediakan

```go
import "strconv"

var s string = "123"
i, _ := strconv.Atoi(s) // Konversi string ke int

var i int = 123
s := strconv.Itoa(i) // Konversi int ke string
```

### 2 . Membuat Array atau Slice 

```go 

// array 
var arr = [3]int{1,2,3}
// array di tentukan panjang datanya. tidak boleh melebihi panjang data yg sudah di tentukan

// slice
var arr = []int{1,2,3} 
// slice panjang data nya tidak di tentukan sehingga bebas mau memasukkan berapa data

```

### 3. Membuat Reflect 

