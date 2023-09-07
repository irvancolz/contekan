package desginpattern

import (
	"log"
	"testing"
)

func TestStory(t *testing.T) {
	Adapter()
}

func TestFactoryMethod(t *testing.T) {
	FactoryMethod()
}

func TestTemplateMthod(t *testing.T) {
	TemplateMethod()
}

func TestSingleTon(t *testing.T) {
	log.Println(GetCurrentTime().Unix())
}
