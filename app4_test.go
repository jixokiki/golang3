package unit_testing

import (
	"testing"
)

func TestTableData(t *testing.T){
	dataTable := []struct{
		Name string
		request string
	}{
		{
			Name : "Rizki",
			request : "Rizki",
		},
		{
			Name : "Iki",
			request : "Iki",
		},
	}

	for _, test := range dataTable{
		//methode subtest dibuat simple
		t.Run(test.Name, func(t *testing.T){
			ModulSay("Rizki")
		})
		t.Run(test.Name, func(t *testing.T){
			ModulSay("Iki")
		})
	}
}