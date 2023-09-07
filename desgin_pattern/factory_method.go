package desginpattern

import "log"

type Mobil struct {
	Merk  string
	Tipe  string
	Harga int64
}

func FactoryMethod() {
	avanza := Mobil{
		Merk:  "toyota",
		Tipe:  "family car",
		Harga: 10.000,
	}
	log.Println("avanza :", avanza)
	xenia := Mobil{
		Merk:  "daihatsu",
		Tipe:  "family car",
		Harga: 10.000,
	}
	log.Println("xenia :", xenia)
	rubicon := Mobil{
		Merk:  "rubicon",
		Tipe:  "jeep",
		Harga: 12.000,
	}
	log.Println("rubicon :", rubicon)
	kia := Mobil{
		Merk:  "kia",
		Tipe:  "jeep",
		Harga: 12.000,
	}
	log.Println("kia :", kia)
}
