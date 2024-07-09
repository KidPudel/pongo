package entities

import (
	"image/color"

	customTypes "github.com/KidPudel/pongo/custom-types"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	PlayerWidth  = 20
	PlayerHeight = 80
)

type Player struct {
	Position customTypes.Vector2[float64]
	Image    *ebiten.Image
}

func InitPlayer() *Player {
	player := &Player{}
	player.Position = customTypes.Vector2[float64]{X: 20, Y: 0}
	player.Image = ebiten.NewImage(PlayerWidth, PlayerHeight)
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
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Position.X, p.Position.Y)
	parent.DrawImage(p.Image, &op)

}
