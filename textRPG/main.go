package main

import "fmt"

type Character struct{
	Hp int
	MaxHp int
	Mp int
	MaxMp int
	Name string
}

const (
	MONSTER_PLAYER = iota
	MONSTER_MAX
)

const (
	CHARACTER_PLAYER = iota
	CHARACTER_MONSTER
	CHARACTER_MAX
)

var monsters = [MONSTER_MAX]Character{
	{15, 15, 15, 15, "ゆうしゃ"},
}

var characters = [CHARACTER_MAX]Character{}

func DrawBattleScreen() {
	fmt.Printf("%v\n", characters[CHARACTER_PLAYER].Name)
	fmt.Printf("H P ：　%v／%v　M P ：%v／%v\n\n", characters[CHARACTER_PLAYER], characters[CHARACTER_PLAYER].MaxHp, characters[CHARACTER_PLAYER].Mp, characters[CHARACTER_PLAYER].MaxMp)
	
}

func Init() {
	characters[CHARACTER_PLAYER] = monsters[MONSTER_PLAYER]
}

func main() {
	Battle()
}

func Battle () {
	DrawBattleScreen()
}

