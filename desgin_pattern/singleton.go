package desginpattern

import (
	"sync"
	"time"
)

type KontrakKoneksiDatabase interface {
	Insert(query string)
	Update(query string)
}
type KoneksiDatabase struct {
	host, username, pasword string
}

func (k KoneksiDatabase) Insert(query string) {}
func (k KoneksiDatabase) Update(query string) {}

func BuatKoneksiDatabase(host, username, pasword string) *KoneksiDatabase {
	return &KoneksiDatabase{
		host:     host,
		username: username,
		pasword:  pasword,
	}
}

func Singleton() {
	mouse := Barang{
		Nama:  "Mouse",
		Stok:  10,
		Harga: 10000,
	}
	mouse.BuatBaru()
	mouse.TambahStok(5)
	mouse.Beli(1)
}

type Barang struct {
	Nama  string
	Stok  int
	Harga int64
}

func (b Barang) BuatBaru() {
	koneksi := BuatKoneksiDatabase("host:123", "admin", "admin")
	koneksi.Insert("upload barang baru.....")
}
func (b Barang) TambahStok(jumlah int) {
	koneksi := BuatKoneksiDatabase("host:123", "admin", "admin")
	koneksi.Update("tambah persediaan barang.....")

}
func (b Barang) Beli(jumlah int) {
	koneksi := BuatKoneksiDatabase("host:123", "admin", "admin")
	koneksi.Update("kurangi persediaan barang.....")
}

func GetCurrentTime() *time.Time {
	var result *time.Time

	if result == nil {
		currentTime := time.Now()
		result = &currentTime
	}

	return result
}

func GetCurrentTimeWithSync() *time.Time {
	var once sync.Once
	var result time.Time

	once.Do(func() {
		result = time.Now()
	})

	return &result
}
