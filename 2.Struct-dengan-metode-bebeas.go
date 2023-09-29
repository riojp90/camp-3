package main

import (
    "fmt"
)

// Struct Person dengan dua field: Name dan Age
type Person struct {
    Name string
    Age  int
}

// Method untuk menampilkan informasi Person
func (p Person) DisplayInfo() {
    fmt.Printf("Nama: %s, Umur: %d tahun\n", p.Name, p.Age)
}

// Method untuk mengubah umur Person
func (p *Person) ChangeAge(newAge int) {
    p.Age = newAge
}

func main() {
    // Meminta input dari pengguna
    var name string
    var age int

    fmt.Print("Masukkan nama: ")
    fmt.Scanln(&name)

    fmt.Print("Masukkan umur: ")
    fmt.Scanln(&age)

    // Membuat objek Person
    person := Person{Name: name, Age: age}

    // Memanggil method DisplayInfo pada objek Person
    person.DisplayInfo()

    // Meminta input dari pengguna untuk mengubah umur
    var newAge int
    fmt.Print("Masukkan umur baru: ")
    fmt.Scanln(&newAge)

    // Mengubah umur Person melalui method ChangeAge
    person.ChangeAge(newAge)

    // Menampilkan informasi setelah perubahan umur
    person.DisplayInfo()
}