# randomXY

randomXY by [LitFill](https://github.com/LitFill).
A cli returning a random combination of 2 ints for some purposes.

has 2 flags `-r` and `-c`.

## Program Random Address Grid

Program ini menghasilkan sekumpulan alamat dalam bentuk koordinat (x, y) untuk sebuah grid yang dimensinya ditentukan oleh pengguna.
Program ini memiliki kemampuan untuk mengacak urutan alamat-alamat tersebut dan mencetaknya satu per satu ke output standar.

### Fitur Utama

#### Pembuatan Grid:

Menghasilkan koordinat (x, y) untuk grid dengan dimensi yang ditentukan pengguna.Dimensi grid dapat diatur melalui flag baris perintah -r (untuk baris) dan -c (untuk kolom).

#### Pengacakan Alamat:

Mengacak urutan alamat dalam grid menggunakan seed dari waktu saat ini untuk memastikan hasil yang berbeda setiap kali program dijalankan.

#### Pencetakan Alamat:

Mencetak alamat yang diacak ke output standar dalam format `addr: X ⊗ Y`.

### Cara Penggunaan

#### Mengatur Dimensi Grid:

Gunakan flag `-r` atau `--rows` untuk menentukan jumlah baris.
Gunakan flag `-c` atau `--cols` untuk menentukan jumlah kolom.

Contoh:

    go run main.go -r 10 -c 5

akan menghasilkan grid 10x5.

#### Menjalankan Program:

Jalankan program menggunakan perintah `go run`, dan hasil alamat akan dicetak ke terminal.
Contoh Eksekusi:

    go run main.go -r 4 -c 3

Hasil Output (acak):

```
addr: 2 ⊗ 1
addr: 0 ⊗ 0
addr: 1 ⊗ 2
addr: 3 ⊗ 1
...
```

Program ini berguna untuk berbagai aplikasi yang memerlukan pengacakan koordinat atau pengolahan data berbasis grid.
