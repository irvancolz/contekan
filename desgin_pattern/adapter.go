package desginpattern

import "log"

type Story interface {
	TampilkanStory()
}

type listStory []Story

type Video struct {
	Link  string
	Judul string
}

func (v Video) TampilkanStory() {
	log.Println(v.Judul, " link video ", v.Link)
}

type Text struct {
	Konten string
}

func (t Text) TampilkanStory() {
	log.Println(t.Konten)
}

type Musik struct {
	Penyanyi string
	Durasi   string
}

func (m Musik) TampilkanStory() {
	log.Printf("sedang mendengarkan : %s dengan durasi: %s ", m.Penyanyi, m.Durasi)
}

func Adapter() {
	var storyList listStory
	storyList = append(storyList, Text{Konten: "selamat Hari jumat"})
	storyList = append(storyList, Video{Judul: "liburan ke bali", Link: "link1"})
	storyList = append(storyList, Video{Judul: "liburan ke purwokerto", Link: "link2"})
	storyList = append(storyList, Musik{Penyanyi: "baby metal", Durasi: "5 menit"})

	for i := 0; i < len(storyList); i++ {
		storyList[i].TampilkanStory()
	}
}
