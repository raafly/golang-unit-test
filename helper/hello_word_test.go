package helper

// gunakan t.Error() dan t.Fatal() untuk menggagalkan test

// untuk melakukan pengecekan sebaiknya gunakan libray testify

import (
	"fmt"
	"runtime"
	"testing"
	"github.com/stretchr/testify/assert"
)

// before and after 
func TestMain( m *testing.M ) {
	// BEFORE 
	fmt.Println("BEFORE UNIT TEST")
	
	m.Run()
	
	// AFTER
	fmt.Println("AFTER UNIT TEST")
}


// test sub test (test didalam test)
func TestSubTest( t *testing.T ) {
	t.Run("rafly", func(t *testing.T) {
		result := HelloWord("rafly")
		assert.Equal(t, "hello rafly", result, "result must be 'hello rafly'")	
	})

	t.Run("nur", func(t *testing.T) {
		result := HelloWord("nur")
		assert.Equal(t, "hello nur", result, "result must be 'hello nur'")
	})
}
 
// menskip test
func TestSkip( t *testing.T ) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run in Mac")
	}

	result := HelloWord("raf")
	assert.NotEqual(t, "hello rafly", result, "result must be 'hello rafly'")
}

func TestHelloWithAssert( t *testing.T ) {
	result := HelloWord("raf")
	assert.NotEqual(t, "hello rafly", result, "result must be 'hello rafly'")
}

// mengagalkan test
func TestHelloWord( t *testing.T ) {
	result := HelloWord("rafly")

	if result != "hello rafly" {
		// error
		t.Error("hasil harus 'hello rafly'")
	}

	fmt.Println("Test Selesai 1")
}

func TestTambah( t *testing.T ) {
	output := Tambah(12, 12)

	if output != 24 {
		// error
		t.Fatal("hasil harus 24")
	}

	fmt.Println("Test Selesai 2")
} 

// tabel tests
func TestTableTest(t *testing.T) {
	tests := []struct {
		nama string
		request string
		expected string
	} {
		{
			nama: "rafly",
			request: "rafly",
			expected: "hello rafly",
		},
		{
			nama: "nur",
			request: "nur",
			expected: "hello nur",
		},
		{
			nama: "ramadhani",
			request: "ramadhani",
			expected: "hello ramadhani",
		},
	}

	for _, test := range tests {
		t.Run(test.nama, func(t *testing.T) {
			result := HelloWord(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}