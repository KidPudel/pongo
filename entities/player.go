package entities

import (
	"fmt"
	"image/color"

	customTypes "github.com/KidPudel/pongo/custom-types"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position customTypes.Vector2[float64]
	Image    *ebiten.Image
}

func InitPlayer() *Player {
	player := &Player{}
	player.Image = ebiten.NewImage(20, 60)
	return player
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Position = p.Position.Add(customTypes.Vector2[float64]{X: 0, Y: -20})
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Position = p.Position.Add(customTypes.Vector2[float64]{X: 0, Y: 20})
	}
}

func (p *Player) Draw(parent *ebiten.Image) {
	p.Image.Fill(color.RGBA{255, 255, 255, 255})
	fmt.Println(*p.Image)
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	parent.DrawImage(p.Image, &op)
	fmt.Println(p.Position)

}
