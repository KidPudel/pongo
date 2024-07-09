package entities

import (
	"image/color"

	customTypes "github.com/KidPudel/pongo/custom-types"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	BallWidth  = 20
	BallHeight = 20
)

type Ball struct {
	Position  customTypes.Vector2[float64]
	Direction customTypes.Vector2[float64]
	Image     *ebiten.Image
}

func InitBall() *Ball {
	ball := Ball{}
	ball.Position = customTypes.Vector2[float64]{X: 400, Y: 300}
	ball.Direction = customTypes.Vector2[float64]{X: 5, Y: 0}
	ball.Image = ebiten.NewImage(BallWidth, BallHeight)

	return &ball
}

func (b *Ball) Update(playerPosition customTypes.Vector2[float64]) {
	ballCollidePlayer := b.Position.X == playerPosition.X && (b.Position.Y >= playerPosition.Y && b.Position.Y <= playerPosition.Y+PlayerHeight) && b.Direction.X < 0
	if (b.Position.X >= 800 || b.Position.X <= 0) || ballCollidePlayer {
		b.Direction.X = b.Direction.X * -1
	}

	b.Position = b.Position.Add(b.Direction)
}

func (b *Ball) Draw(parent *ebiten.Image) {
	b.Image.Fill(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.Position.X, b.Position.Y)
	parent.DrawImage(b.Image, &op)

}
