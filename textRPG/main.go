package main

import (
	"log"

	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

)

type Charactor struct{
  HP, MaxHP, MP, MaxMP int
  Name string
}

const (
  MONSTER_PLAYER = 1
  MONSTER_MAX = 1
)

var Player = Charactor{HP: 15, MaxHP: 15, MP: 15, MaxMP: 15, Name: "ゆうしゃ"}

var Monsters [MONSTER_MAX]Charactor = [MONSTER_MAX]Charactor{{HP: 15, MaxHP: 15, MP: 15, MaxMP: 15, Name: "すらいむ"}}

func battle () {
  // nll
}

type Game struct{}
var stateNum = 0
func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  ebitenutil.DebugPrint(screen, fmt.Sprintf("%s \nHP: %d/%d MP: %d/%d\n", Player.Name, Player.HP, Player.MaxHP, Player.MP, Player.MaxMP))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("text rpg")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
