package main

import (
	"fmt"
)

// Struct Student dengan field Name, Age, dan Grade
type Student struct {
	Name  string
	Age   int
	Grade string
}

// Method untuk menampilkan informasi Student
func (s *Student) DisplayInfo() {
	fmt.Printf("Nama: %s, Umur: %d tahun, Nilai: %s\n", s.Name, s.Age, s.Grade)
}

// Method untuk mengubah nilai Grade Student
func (s *Student) ChangeGrade(newGrade string) {
	s.Grade = newGrade
}

func main() {
	// Meminta input dari pengguna
	var name, grade string
	var age int

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)

	fmt.Print("Masukkan umur: ")
	fmt.Scanln(&age)

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&grade)

	// Membuat objek Student menggunakan pointer
	student := &Student{Name: name, Age: age, Grade: grade}

	// Menampilkan informasi awal
	fmt.Println("\nInformasi awal:")
	student.DisplayInfo()

	// Meminta input dari pengguna untuk mengubah nilai Grade
	fmt.Print("\nMasukkan nilai baru: ")
	fmt.Scanln(&grade)

	// Mengubah nilai Grade melalui method ChangeGrade
	student.ChangeGrade(grade)

	// Menampilkan informasi setelah perubahan
	fmt.Println("\nInformasi setelah perubahan:")
	student.DisplayInfo()
}