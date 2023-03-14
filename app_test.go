package unit_testing

import (
	"testing"
	"fmt"
)

// Modul Sederhana Unit Testing

func TestApp(t *testing.T){
	result := ModulSay("Rizki")
	if result != "Wellcome To Unit Testing Rizki"{
		panic("Result must be Wellcome To Unit Testing ")
	}
}

func TestApp2(t *testing.T){
	result:= ModulSay("Iki")
	if result !=  "Wellcome To Unit Testing Rizki"{
		t.Error()
	}
	fmt.Println("Result Akhir Seharusnya "+result)
}

func TestApp3(t * testing.T){
	result := ModulSay("Rizki")
	if result != "Wellcome To Unit Testing Iki" {
		t.Fatal()
		//hasilnya akan eror dan code dibawahnya tidak akan di eksekusi
	}
	fmt.Println("Result Akhir Seharusnya "+result)
}

func TestApp4(t *testing.T){
	result := ModulSay("Owen")
	if result != "Wellcome To Unit Testing Rizki" {
		t.Fail()
	}
	fmt.Println("Result Akhir Seharusnya "+result)
}

func TestApp5(t *testing.T){
	result := ModulSay("Hito")
	if result != "Wellcome To Unit Testing Rizki" {
		t.FailNow()
	}
}