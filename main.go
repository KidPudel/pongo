package main

import (
	"image/color"
	"log"

	"github.com/KidPudel/pongo/entities"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *entities.Player
	ball   *entities.Ball
}

func (g *Game) Update() error {
	g.player.Update()
	g.ball.Update(g.player.Position)

	return nil
}

func (g *Game) Draw(image *ebiten.Image) {
	image.Fill(color.Black)

	g.player.Draw(image)
	g.ball.Draw(image)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("pong")

	game := Game{}
	player := entities.InitPlayer()
	ball := entities.InitBall()
	game.player = player
	game.ball = ball

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
