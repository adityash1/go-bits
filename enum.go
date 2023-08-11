package main

import "fmt"

type WeaponType int

const (
	Axe WeaponType = iota
	Sword
	WoodenStick
	Knife
)

func (w WeaponType) String() string {
	switch w {
	case Sword:
		return "SWORD"
	case Axe:
		return "AXE"
	}

	return ""
}

func getDamage(weaponType WeaponType) int {
	switch weaponType {
	case Axe:
		return 100
	case Sword:
		return 90
	case WoodenStick:
		return 10
	case Knife:
		return 50
	default:
		panic("weapon doesn't exist")
	}
}

func main() {
	fmt.Println("damage of weapon (%s) is (%d)\n", Axe, getDamage(Axe))
}
