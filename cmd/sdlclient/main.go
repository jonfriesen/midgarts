package main

import (
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/EngoEngine/ecs"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	_ "github.com/joho/godotenv/autoload"
	"github.com/project-midgard/midgarts/internal/entity"
	"github.com/project-midgard/midgarts/internal/opengl"
	"github.com/project-midgard/midgarts/internal/system"
	"github.com/project-midgard/midgarts/internal/window"
	"github.com/project-midgard/midgarts/pkg/camera"
	"github.com/project-midgard/midgarts/pkg/common/character"
	"github.com/project-midgard/midgarts/pkg/common/character/directiontype"
	"github.com/project-midgard/midgarts/pkg/common/character/jobspriteid"
	"github.com/project-midgard/midgarts/pkg/common/character/statetype"
	"github.com/project-midgard/midgarts/pkg/common/fileformat/grf"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowWidth  = 1280
	WindowHeight = 720
	AspectRatio  = float32(WindowWidth) / float32(WindowHeight)
)

var (
	GrfFilePath = os.Getenv("GRF_FILE_PATH")
)

func main() {
	runtime.LockOSThread()

	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	var win *sdl.Window
	if win, err = sdl.CreateWindow(
		"Midgarts Client",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WindowWidth,
		WindowHeight,
		sdl.WINDOW_OPENGL,
	); err != nil {
		panic(err)
	}
	defer func() {
		_ = win.Destroy()
	}()

	context, err := win.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	gls := opengl.InitOpenGL()

	var grfFile *grf.File
	if grfFile, err = grf.Load(GrfFilePath); err != nil {
		log.Fatal(err)
	}

	gl.Viewport(0, 0, WindowWidth, WindowHeight)

	cam := camera.NewPerspectiveCamera(0.638, AspectRatio, 0.1, 1000.0)
	cam.ResetAngleAndY(WindowWidth, WindowHeight)

	ks := window.NewKeyState(win)

	w := ecs.World{}
	renderSys := system.NewCharacterRenderSystem(grfFile)
	actionSystem := system.NewCharacterActionSystem(grfFile)

	c1 := entity.NewCharacter(character.Male, jobspriteid.Novice, rand.Intn(31-1)+1)
	c2 := entity.NewCharacter(character.Female, jobspriteid.Knight, rand.Intn(31-1)+1)
	c3 := entity.NewCharacter(character.Male, jobspriteid.Swordsman, rand.Intn(31-1)+1)
	c4 := entity.NewCharacter(character.Female, jobspriteid.Alchemist, rand.Intn(31-1)+1)
	c5 := entity.NewCharacter(character.Male, jobspriteid.Alcolyte, rand.Intn(31-1)+1)
	c6 := entity.NewCharacter(character.Female, jobspriteid.MonkH, rand.Intn(31-1)+1)
	c7 := entity.NewCharacter(character.Male, jobspriteid.Crusader2, rand.Intn(31-1)+1)
	c8 := entity.NewCharacter(character.Male, jobspriteid.Assassin, rand.Intn(31-1)+1)
	c9 := entity.NewCharacter(character.Male, jobspriteid.Alchemist, rand.Intn(31-1)+1)
	c10 := entity.NewCharacter(character.Female, jobspriteid.Wizard, rand.Intn(31-1)+1)

	c1.SetPosition(mgl32.Vec3{0, 42, 0})
	c2.SetPosition(mgl32.Vec3{4, 42, 0})
	c3.SetPosition(mgl32.Vec3{8, 42, 0})
	c4.SetPosition(mgl32.Vec3{0, 38, 0})
	c5.SetPosition(mgl32.Vec3{4, 38, 0})
	c6.SetPosition(mgl32.Vec3{8, 38, 0})
	c7.SetPosition(mgl32.Vec3{0, 34, 0})
	c8.SetPosition(mgl32.Vec3{4, 34, 0})
	c9.SetPosition(mgl32.Vec3{8, 34, 0})
	c10.SetPosition(mgl32.Vec3{12, 34, 0})

	var actionable *system.CharacterActionable
	var renderable *system.CharacterRenderable
	w.AddSystemInterface(actionSystem, actionable, nil)
	w.AddSystemInterface(renderSys, renderable, nil)
	w.AddSystem(system.NewOpenGLRenderSystem(gls, cam, renderSys.RenderCommands))

	w.AddEntity(c1)
	//w.AddEntity(c2)
	//w.AddEntity(c3)
	//w.AddEntity(c4)
	//w.AddEntity(c5)
	//w.AddEntity(c6)
	//w.AddEntity(c7)
	//w.AddEntity(c8)
	//w.AddEntity(c9)
	//w.AddEntity(c10)

	shouldStop := false

	frameStart := time.Now()

	for !shouldStop {
		gl.ClearColor(0, 0.5, 0.8, 1.0)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				shouldStop = true
				break
			}
		}

		const MovementRate = float32(0.075)

		p1 := c1.Position()
		// char controls
		if ks.Pressed(sdl.K_UP) && ks.Pressed(sdl.K_RIGHT) {
			c1.Direction = directiontype.NorthEast
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X() - MovementRate, p1.Y() + MovementRate, p1.Z()})
		} else if ks.Pressed(sdl.K_UP) && ks.Pressed(sdl.K_LEFT) {
			c1.Direction = directiontype.NorthWest
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X() + MovementRate, p1.Y() + MovementRate, p1.Z()})
		} else if ks.Pressed(sdl.K_DOWN) && ks.Pressed(sdl.K_RIGHT) {
			c1.Direction = directiontype.SouthEast
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X() - MovementRate, p1.Y() - MovementRate, p1.Z()})
		} else if ks.Pressed(sdl.K_DOWN) && ks.Pressed(sdl.K_LEFT) {
			c1.Direction = directiontype.SouthWest
			c1.SetPosition(mgl32.Vec3{p1.X() + MovementRate, p1.Y() - MovementRate, p1.Z()})
			c1.SetState(statetype.Walking)
		} else if ks.Pressed(sdl.K_UP) {
			c1.Direction = directiontype.North
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X(), p1.Y() + MovementRate, p1.Z()})
		} else if ks.Pressed(sdl.K_DOWN) {
			c1.Direction = directiontype.South
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X(), p1.Y() - MovementRate, p1.Z()})
		} else if ks.Pressed(sdl.K_RIGHT) {
			c1.Direction = directiontype.East
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X() - MovementRate, p1.Y(), p1.Z()})
		} else if ks.Pressed(sdl.K_LEFT) {
			c1.Direction = directiontype.West
			c1.SetState(statetype.Walking)
			c1.SetPosition(mgl32.Vec3{p1.X() + MovementRate, p1.Y(), p1.Z()})
		} else {
			c1.SetState(statetype.Idle)
		}

		// camera controls
		if ks.Pressed(sdl.K_w) {
			cam.SetPosition(mgl32.Vec3{cam.Position().X(), cam.Position().Y(), cam.Position().Z() + 0.2})
		} else if ks.Pressed(sdl.K_s) {
			cam.SetPosition(mgl32.Vec3{cam.Position().X(), cam.Position().Y(), cam.Position().Z() - 0.2})
		}

		now := time.Now()
		frameDelta := now.Sub(frameStart)
		frameStart = now

		w.Update(float32(frameDelta.Seconds()))

		win.GLSwap()
	}
}
