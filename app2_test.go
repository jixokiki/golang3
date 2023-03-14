package unit_testing

import (
	
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"fmt"
)

func TestData(t *testing.T){
	hasil := ModulData("Michele 10804220200")
	if hasil != "Mahasiswa Michele 10804220200" {
		panic("Can not run testing")
	}
}

//hasilnya akan gagal testing
func TestData2(t *testing.T){
	hasil := ModulData("Michele 10804220200")
	assert.Equal(t, "Mahasiswa Michele 108040200", hasil, "Result is must be 'Mahasiswa Michele 10804220200'")
	fmt.Println("TestData2 is Done")
}

//hasilnya akan berhasil testing
func TestData3(t *testing.T){
	hasil := ModulData("Michele 10804220200")
	require.Equal(t, "Mahasiswa Michele 10804220200", hasil, "Result is must be 'Mahasiswa Michele 10804220200'")
	fmt.Println("TestData3 is Done")
}