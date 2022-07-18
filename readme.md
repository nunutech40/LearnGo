
- How to run go file go in terminal:

cd MyProjectGo -> "go run ."

#Menambahkan Dependecies
- langsung import modul yang akan dipake di file .go, contoh: import "rsc.io/quote"
- masuk ke terminal: $ go mod tidy

#Membuat go mod init -> go mod init dibutuhkan untuk setiap module untuk mentrack/melacak dependecies yang digunakan di kode kita.
- go mod init example.com/hello
- perintah di atas, di lakukan dalam direktori folder module go yang kita buat
- akan tercipta file go.mod, yang pada awalnya, isinya kosong, hanya ada versi go yang kita pake, tapi setelah kita menambahkan dependecy ke dalam modul kita, file2 dependecy akan di tuliskan/dicatat di dalam file go.mod

#How to install go?
- masuk ke go.dev untuk download sdk sesuai sistem operasi yang dipake.
- baca dokumentasi di golang indonesia : https://golang-id.org/doc
