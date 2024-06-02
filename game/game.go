package game

import (
	"fmt"
	"image/color"
	"my-game/assets"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player  *Player
	lasers  []*Laser
	meteors []*Meteor
	meteorSpawTimer *Timer
	score         int
}

func NewGame() *Game {
	g := &Game{
		meteorSpawTimer: NewTimer(24),
	}
	player := NewPlayer(g)
	g.player = player
	return g
}

// rode 60fps
// responsável por atualizar a logica do jogo
// 60 x por segundo
// 1 x rodando
func (g *Game) Update() error {
	g.player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

    g.meteorSpawTimer.Update()
	if g.meteorSpawTimer.IsReady() {
    g.meteorSpawTimer.Reset()

		m := NewMeteor()
		g.meteors = append(g.meteors, m)	
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if (m.Collider().Intersects(g.player.Collider())){
			fmt.Println("Você Perdeu !!!")
			g.Reset()
		}
			
		}

		for i, m := range(g.meteors){


			for j, l := range(g.lasers){


				if (m.Collider().Intersects(l.Collider())){

					g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)


					g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)

					g.score += 1
					
				}
			
			}
		}

	return nil
}

// responsável por desenhar objetos na tela
// 60 x por segundo
func (g *Game) Draw(screen *ebiten.Image) {

	g.player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	 // Draw score
    text.Draw(screen, fmt.Sprintf("Score : %d", g.score), assets.FontUi, 20, 100, color.White)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight

}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)

}

func (g *Game) Reset() {

	g.player = NewPlayer(g)
	g.meteors = nil
    g.lasers = nil
	g.meteorSpawTimer.Reset()
	g.score = 0
}
