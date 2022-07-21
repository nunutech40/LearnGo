
- How to run go file go in terminal:

cd MyProjectGo -> "go run ."Cancel changes

# Membuat go mod init -> go mod init dibutuhkan untuk setiap module untuk mentrack/melacak dependecies yang digunakan di kode kita.
- go mod init example.com/hello
- perintah di atas, di lakukan dalam direktori folder module go yang kita buat
- akan tercipta file go.mod, yang pada awalnya, isinya kosong, hanya ada versi go yang kita pake, tapi setelah kita menambahkan dependecy ke dalam modul kita, file2 dependecy akan di tuliskan/dicatat di dalam file go.mod


# How to install go?
- masuk ke go.dev untuk download sdk sesuai sistem operasi yang dipake.
- baca dokumentasi di golang indonesia : https://golang-id.org/doc

# Menambahkan Dependecies
- langsung import modul yang akan dipake di file .go, contoh: import "rsc.io/quote"
- masuk ke terminal: $ go mod tidy

# Create and run unit test
- create file _test for identity test
- example unit test:

```go

// Test function Hello dengan sebuah nama, dan memeriksa return yang valid
func TestHelloName(t *testing.T) {
	name := "Dorothy"
	want := regexp.MustCompile(`\b` +name+ `\b`)
	msg, err := Hello("Dorothy")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Dorothy") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// Test function Hello dengan mengisi empty string, dan memeriksa jika ada error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}

```

- run this in terminal: go test or go test -v
