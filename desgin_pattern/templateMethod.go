package desginpattern

import "log"

type Seduhan interface {
	SiapkanBahan()
	TambahHiasan()
	Tuang()
	RebusAir()
}

func Masak(resep Seduhan) {
	log.Println("==== proses dimulai ====")
	resep.SiapkanBahan()
	resep.RebusAir()
	resep.Tuang()
	resep.TambahHiasan()
	log.Println("==== proses selesai ====")
}

type Kopi struct {
	Merk          string
	SuhuPenyajian int
	Wadah         string
	Topping       string
}

func (k Kopi) SiapkanBahan() {
	log.Println("proses menyeduh kopi ", k.Merk)
}
func (k Kopi) RebusAir() {
	log.Println("rebus air sampai suhu ", k.SuhuPenyajian)
}
func (k Kopi) Tuang() {
	log.Println("masukkan kopi dan air kedalam", k.Wadah)
}
func (k Kopi) TambahHiasan() {
	log.Println("tambahkan ", k.Topping)
}

type SayurRebus struct {
	Bahan, Wadah, Hiasan, LamaMerebus string
}

func (s SayurRebus) SiapkanBahan() {
	log.Println("siapkan bahan membuat sayur ", s.Bahan)
}
func (s SayurRebus) RebusAir() {
	log.Println("rebus air selama ", s.LamaMerebus)
}
func (s SayurRebus) Tuang() {
	log.Println("masukkan sayur yand sudah matang kedalam ", s.Wadah)
}
func (s SayurRebus) TambahHiasan() {
	log.Printf("tambahkan %s sebagai hiasan", s.Hiasan)
}

func TemplateMethod() {
	cappucino := Kopi{
		Merk:          "cappucino",
		SuhuPenyajian: 100,
		Wadah:         "cangkir kecil",
		Topping:       "susu",
	}

	sayurBayam := SayurRebus{
		Bahan:       "bayam , wortel",
		Wadah:       "mangkok",
		Hiasan:      "tomat",
		LamaMerebus: "10 menit",
	}

	Masak(cappucino)
	Masak(sayurBayam)
}
