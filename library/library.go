package library

import (
	"fmt"
	"math/rand"
)

type mahasiswa struct {
	nama string
	umur int
}

func Method() {
	var m1, m2, m3 mahasiswa

	m1.nama = "Ahmad"
	m2.nama = "Abi"
	m3.nama = "Zaka"

	m1.umur = rand.Intn(100)
	m2.umur = rand.Intn(100)
	m3.umur = rand.Intn(100)

	m1.print()
	m2.print()
	m3.print()
}

func (m mahasiswa) print() {
	fmt.Println(m.nama, m.umur)
}
