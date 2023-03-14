package unit_testing

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"fmt"
)

// Membuat SubTest Unit Testing
func TestSubData(t *testing.T){
	t.Run("Iki", func(t *testing.T){
		hasil := ModulSay("Iki")
		assert.Equal(t, "Wellcome To Unit Testing Iki", hasil, "Result is must be 'Wellcome To Unit Testing Iki'")
		fmt.Println("TesSubData 1 is Done")
	})
	t.Run("Rizki", func(t *testing.T){
		hasil := ModulSay("Rizki")
		require.Equal(t, "Wellcome To Unit Testing Iki", hasil, "Result is must be 'Wellcome To Unit Testing Iki'")
		fmt.Println("TesSubData 2 is Done")
	})
}