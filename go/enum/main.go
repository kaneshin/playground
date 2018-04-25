package main

import "fmt"

type Member int
type Male Member
type Female Member

const (
	male Member = iota
	female
)

func NewMale() Male {
	return Male(male)
}

func NewFemale() Female {
	return Female(female)
}

type MemberInterface interface {
	Say()
}

func (m Member) Gender() string {
	if m == male {
		return "Male"
	}
	if m == female {
		return "Female"
	}
	return "Neutral"
}

func (m Male) Say() {
	fmt.Printf("%s, 男性振る舞い\n", Member(m).Gender())
}

func (m Female) Say() {
	fmt.Printf("%s, 女性振る舞い\n", Member(m).Gender())
}

func Print(m MemberInterface) {
	m.Say()
}

func main() {
	Print(NewMale())
	Print(NewFemale())

	NewMale().Say()
	NewFemale().Say()
}
