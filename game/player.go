package game

import (
	"my-game/assets"
    
	"github.com/hajimehoshi/ebiten/v2"
	
)        

type Player struct {
	image *ebiten.Image
	position Vector
	game *Game
	laserLoadingTimer *Timer
}
func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
		bounds := image.Bounds()
		halfW  := float64(bounds.Dx()) / 2
		position := Vector{
			X: (float64(screenWidth) / 2) - halfW,
			Y: 500,
		}
    return &Player{
		image: image,
		position: position,	
		game: game,
		laserLoadingTimer: NewTimer(12),
	}
}
func (p *Player) Update(){

	speed := 8.0

	if(ebiten.IsKeyPressed(ebiten.KeyLeft)){
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight){
		p.position.X += speed
	}

	p.laserLoadingTimer.Update()

	if (ebiten.IsKeyPressed(ebiten.KeySpace)) && p.laserLoadingTimer.IsReady(){

		p.laserLoadingTimer.Reset()

		bounds := p.image.Bounds()
	    halfW  := float64(bounds.Dx()) / 2 // Metade da largura da imagem do laser
        halfH  := float64(bounds.Dy()) / 2 // Metade da altura da imagem do laser

		spawPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}

		laser := NewLaser(spawPos)
		p.game.AddLasers(laser)
		
	}

}
func (p *Player) Draw(screen *ebiten.Image){
	op := &ebiten.DrawImageOptions{}

	//Posição X e Y que a imagem sera desenhada a tela
	op.GeoM.Translate(p.position.X, p.position.Y)
	//
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect{
	bounds := p.image.Bounds()

	return NewRect(p.position.X, p.position.Y,
	float64(bounds.Dx()), 
	float64(bounds.Dy()))
}