package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// 10.2023-001 - BITTY KNIGHT

// NOTE ALL THE MARK: NOTES ARE FOR VS CODE MARK JUMPING AND CAN BE IGNORED

// MARK: VAR
var (

	//TIMES
	times             []int
	besttime, timeson bool
	besttimesT        int32

	//AUDIO
	music      rl.Music
	backMusic  []rl.Music
	sfx        []rl.Sound
	musicon    bool
	volume     = float32(0.2)
	bgMusicNum int

	//END LEVEL
	bosses              []xboss
	bossnum             int
	endgameT, endPauseT int32
	endgopherrec        rl.Rectangle

	//STARTSCREEN
	startscreen, intro, introcount bool
	introT1, introT2, introT3      = fps * 2, fps * 2, fps * 3
	introF1, introF2, introF3      = float32(0), float32(0), float32(0)

	//MARIO
	marioon, mariojump                            bool
	marioT, mariojumpT                            int32
	mariorecs, mariocoins                         []rl.Rectangle
	marioPL, marioScreenRec, patternRec, marioImg rl.Rectangle
	marioV2L, marioV2R                            rl.Vector2
	marioCols                                     []rl.Color
	mariocoinonoff                                []bool

	//SHOP
	shopon    bool
	shopExitY float32
	shopitems []xblok
	shopnum   int
	shopExitT int32

	//OPTIONS CREDITS
	optionnum int
	txtSize   = txU2
	optionT   int32

	shaderon, optionson, hpBarsOn, artifactson, scanlineson, creditson, helpon, invincible, resettimes, restarton, optionsChange, controllerDisconnect, controllerWasOn bool

	//ENEMIES
	enProj []xproj

	enSpikes, enGhost, enSlime, enRock, enMushroom = xenemy{}, xenemy{}, xenemy{}, xenemy{}, xenemy{}

	//INVEN
	inven   []xblok
	invenon bool

	//FX & TEXT & MODS
	fx                         []xfx
	txtSoldlist                []xtxt
	snow                       []ximg
	scanlinev2                 []rl.Vector2
	gametxt                    []xtxt
	chainV2                    []rl.Vector2
	chainLightOn               bool
	chainLightTimer            int32
	rain                       []rl.Rectangle
	floodRec, floodImg         rl.Rectangle
	fish1, fish2               rl.Rectangle
	fishV2, fish2V2            rl.Vector2
	fishSiz, fishSiz2          float32
	fishLR, fish2LR            bool
	fishRec, fishRec2          rl.Rectangle
	waterLR, waterUP           bool
	airstrikeT, airstrikebombT int32
	airstrikeDir               int
	airstrikeOn                bool
	airstrikeV2                []rl.Vector2
	fireworksCnt               rl.Vector2
	fadeblinkon, fadeblinkon2  bool
	fadeblink                  = float32(0.5)
	fadeblink2                 = float32(0.5)

	//LEVEL
	level                        []xroom
	levRec, levRecInner, wallT   rl.Rectangle
	levW                         = float32(720)
	borderWallBlokSiz            = bsU
	levX, levY                   float32
	roomNum, levBorderBlokNum    int
	levMap                       []rl.Rectangle
	levelnum                     = 1
	exitRoomNum, shopRoomNum     int
	exited, nextlevelscreen      bool
	secs, mins, minsEND, secsEND int

	roomChangedTimer, anchorT, runT, diedscrT, nextlevelT int32

	levMapOn, roomChanged, night, flipcam, shader2on, shader3on, exiton, exitLR, endgame, hardcore bool

	//COMPANIONS
	mrplanty, mralien, mrcarrot = xcompanion{}, xcompanion{}, xcompanion{}

	//PLAYER
	pl                          = xplayer{}
	mods                        = xmod{}
	max                         = xmax{}
	plProj                      []xproj
	kills                       = xkills{}
	hpHitY, reviveY, waterY     float32
	hpHitF, reviveF, waterF     = float32(1), float32(1), float32(1)
	plVineRec, diedRec, diedIMG rl.Rectangle
	teleportRoomNum             int
	teleportRadius              []float32
	startdmgT                   int32

	escaped, escapeRoomFound, teleporton, platkrecon, chainLightingSwingOnOff, died bool

	//CONTROLLER
	useController, isController bool
	contolleron                 = true

	//CORE
	imgs                      rl.Texture2D
	shader, shader2, shader3  rl.Shader
	renderTarget              rl.RenderTexture2D
	frames, scrW, scrH        int
	scrW32, scrH32, keypressT int32
	scrWF32, scrHF32          float32
	cam2                      rl.Camera2D
	fps                       = int32(60)
	mouseV2, mousev2cam, cnt  rl.Vector2
	ori                       = rl.NewVector2(0, 0)
	debug, pause              bool

	//IMGS
	walltiles, floortiles, bats, knight, etc, shrines, plants, skulls, candles, signs, splats, statues, mushrooms, alien, patterns, gems []rl.Rectangle

	coin = rl.NewRectangle(1120, 250, 16, 16)

	rabbit1, fireballPlayer, burn, star, wateranim, plantBull, spikes, spring, posiongas, mushBull, blades, spear, firetrailanim, orbitalanim, floodanim, fishR, fishL, airstrikeanim, boss1anim, boss2anim = xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}, xanim{}

	//UNITS
	bsU, bsU2, bsU3, bsU4, bsU5, bsU6, bsU7, bsU8, bsU9, bsU10, bsU11, bsU12 = float32(16), bsU * 2, bsU * 3, bsU * 4, bsU * 5, bsU * 6, bsU * 7, bsU * 8, bsU * 9, bsU * 10, bsU * 11, bsU * 12

	bsUi, bsU2i, bsU3i, bsU4i, bsU5i, bsU6i, bsU7i, bsU8i, bsU9i, bsU10i = 16, bsUi * 2, bsUi * 3, bsUi * 4, bsUi * 5, bsUi * 6, bsUi * 7, bsUi * 8, bsUi * 9, bsUi * 10

	bsUi32, bsU2i32, bsU3i32, bsU4i32, bsU5i32, bsU6i32, bsU7i32, bsU8i32, bsU9i32, bsU10i32 = int32(16), bsUi32 * 2, bsUi32 * 3, bsUi32 * 4, bsUi32 * 5, bsUi32 * 6, bsUi32 * 7, bsUi32 * 8, bsUi32 * 9, bsUi32 * 10

	txU, txU2, txU3, txU4, txU5, txU6, txU7, txU8, txU9, txU10 = int32(10), txU * 2, txU * 3, txU * 4, txU * 5, txU * 6, txU * 7, txU * 8, txU * 9, txU * 10
)

// MARK:STRUCT

type xboss struct {
	img, rec, crec            rl.Rectangle
	yt, xl, vel, velX, velY   float32
	hppause, timer            int32
	hp, hpmax, direc, atkType int
	cnt                       rl.Vector2
	off                       bool
}
type xmod struct {
	axe, santa, snowon, fireball, vine, key, apple, planty, medikit, wallet, exitmap, firetrail, hppotion, invisible, orbital, chainlightning, recharge, anchor, umbrella, socks, flood, peace, alien, airstrike, fireworks, carrot bool

	axeN, fireballN, bounceN, keyN, appleN, firetrailN, hppotionN, coffeeN, atkrangeN, atkdmgN, orbitalN, hpringN, armorN, cherryN, cakeN int

	axeT, axeT2, santaT int32
}
type xmax struct {
	axe, fireball, bounce, key, apple, firetrail, hppotion, coffee, atkrange, atkdmg, orbital, hpring, armor, cherry, cake int
}
type xenemy struct {
	img, rec, imgl, imgr, crec, arec       rl.Rectangle
	cnt, ori                               rl.Vector2
	ro, vel, velX, velY, xImg, xImg2, fade float32
	frameNum, direc, hp, hpmax, spawnN     int
	fly, anim, off                         bool
	name                                   string
	col                                    rl.Color
	hppause, T1                            int32
}
type xplayer struct {
	cnt, ori, orbital1, orbital2 rl.Vector2

	img, rec, crec, arec, atkrec, orbimg1, orbimg2, orbrec1, orbrec2 rl.Rectangle

	atkTimer, hppause, slideT, poisonT, poisonCollisT, hppotionT, peaceT, waterT int32

	move, atk, slide, escape, revived, poison, armorHit, underWater bool

	siz, ro, vel, imgAtkX, imgWalkX, sizImg, angle float32

	direc, framesAtk, framesWalk, hp, hpmax, coins, atkDMG, slideDIR, poisonCount, armor, armorMax, rechargeN int
}
type xblok struct {
	name, desc                  string
	img, rec, crec, crec2, drec rl.Rectangle
	color                       rl.Color
	fade                        float32
	cnt, ori                    rl.Vector2
	velX, velY, vel, ro         float32
	v2s                         []rl.Vector2
	timer, txtT                 int32

	movType, numof, slideDIR, numType, numCoins, shopprice int

	bump, onoff, solid, onoffswitch, fadeswitch, shopoff bool
}

type xcompanion struct {
	img, imgl, imgr, rec, crec rl.Rectangle
	hp, hpmax, frames          int
	vel, velx, vely            float32
	cnt                        rl.Vector2
	timer                      int32
}
type xroom struct {
	floor, walls, movBloks, etc, innerBloks, spikes []xblok

	doorSides     []int
	nextRooms     []int
	visited, exit bool
	floorT, wallT rl.Rectangle
	enemies       []xenemy
	doorExitRecs  []rl.Rectangle
}

type ximg struct {
	img, rec rl.Rectangle
	ro, fade float32
	cnt, ori rl.Vector2
	col      rl.Color
	off      bool
}
type xkills struct {
	bunnies, bats, mushrooms, rocks, slimes, spikehogs, ghosts int
}
type xtxt struct {
	txt, txt2 string
	x, y      int32
	fade      float32
	col       rl.Color
	onoff     bool
}
type xproj struct {
	cnt                  rl.Vector2
	drec, rec, crec, img rl.Rectangle
	ori                  rl.Vector2
	onoff                bool
	col                  rl.Color
	dmg, bounceN         int
	name                 string

	ro, vel, velx, vely, fade float32
}
type xfx struct {
	timer    int32
	onoff    bool
	name     string
	cnt      rl.Vector2
	rec, img rl.Rectangle
	col      rl.Color
	fade     float32

	recs []xrec
}
type xrec struct {
	rec              rl.Rectangle
	col              rl.Color
	fade, velX, velY float32
}
type xanim struct {
	xl, yt, frames, W float32
	recTL             rl.Rectangle
}

// MARK: DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW DRAW
func drawcam() { //MARK:DRAW CAM

	//INVEN OPTIONS SHOP ON
	if optionson {
		drawOptions()
	} else if timeson {
		drawTimes()
	} else if shopon {
		drawShop()
	} else if marioon {
		drawUpMario()
	} else if endgame {
		drawEndGame()
	} else if nextlevelscreen {
		drawnextlevelscreen()
	} else if died {
		drawDied()
	}

	//DRAW GAME LEVEL
	if !pause {

		//FLOORS
		dFloors := level[roomNum].floor
		for a := 0; a < len(dFloors); a++ {
			drawBlok(dFloors[a], false, false, 0)
		}
		//rl.DrawRectangleRec(levRec, rl.Fade(rl.Black, 0.5))

		//SPIKES
		if len(level[roomNum].spikes) > 0 {
			for a := 0; a < len(level[roomNum].spikes); a++ {
				shadowRec := level[roomNum].spikes[a].rec
				shadowRec.X -= 5
				shadowRec.Y += 5
				rl.DrawTexturePro(imgs, level[roomNum].spikes[a].img, shadowRec, ori, 0, rl.Fade(rl.Black, 0.8))
				col := ranRed()
				rl.DrawTexturePro(imgs, level[roomNum].spikes[a].img, level[roomNum].spikes[a].rec, ori, 0, col)

				if frames%3 == 0 {
					level[roomNum].spikes[a].img.X += 32
					if level[roomNum].spikes[a].img.X >= spikes.xl+spikes.frames*32 {
						level[roomNum].spikes[a].img.X = spikes.xl
					}
				}

				//CHECK PLAYER SPIKES COLLISION
				if rl.CheckCollisionRecs(pl.crec, level[roomNum].spikes[a].rec) {
					hitPL(0, 2)
				}
			}
		}

		//WALLS
		dWalls := level[roomNum].walls
		for a := 0; a < len(dWalls); a++ {
			drawBlok(dWalls[a], false, false, 0)
			//WALL BLOCKNUMS
			if debug {
				rl.DrawText(fmt.Sprint(a), dWalls[a].rec.ToInt32().X+4, dWalls[a].rec.ToInt32().Y+4, txU, rl.White)
			}
		}

		//INNER BLOKS
		if len(level[roomNum].innerBloks) > 0 {
			for a := 0; a < len(level[roomNum].innerBloks); a++ {
				drawBlok(level[roomNum].innerBloks[a], false, true, 4)
			}
		}

		//ENEMY PROJECTILES
		if len(enProj) > 0 {
			drawUpEnProj()
		}

		//FX
		if len(fx) > 0 {
			drawupfx()
		}

		//ETC
		if len(level[roomNum].etc) > 0 {
			drawUpEtc()
		}

		//PLAYER PROJECTILES
		if len(plProj) > 0 {
			drawUpPlayerProj()
		}

		//CHAIN LIGHTNING
		if chainLightOn {
			drawChainLight()
		}

		//ENEMIES
		if len(level[roomNum].enemies) > 0 {
			drawUpEnemies()
		}
		//BOSS
		if levelnum == 6 {
			drawUpBoss()
		}

		//MOVE BLOKS
		dMovBloks := level[roomNum].movBloks
		for a := 0; a < len(dMovBloks); a++ {
			drawBlok(dMovBloks[a], true, false, 0)
		}

		//PLAYER
		drawPlayer()

		//GAME TEXT
		if len(gametxt) > 0 {
			clear := false
			for a := 0; a < len(gametxt); a++ {
				if gametxt[a].onoff {
					rl.DrawText(gametxt[a].txt, gametxt[a].x, gametxt[a].y, 20, rl.Fade(gametxt[a].col, gametxt[a].fade))

					gametxt[a].fade -= 0.02
					gametxt[a].y--
					if gametxt[a].fade <= 0 {
						gametxt[a].onoff = false
					}
				} else {
					clear = true
				}
			}
			if clear {
				for a := 0; a < len(gametxt); a++ {
					if !gametxt[a].onoff {
						gametxt = remTxt(gametxt, a)
					}
				}
			}
		}

		//AIR STRIKE
		if airstrikeOn {
			drawUpAirStrike()
		}

		//SNOW
		if len(snow) > 0 {
			clear := false
			for a := 0; a < len(snow); a++ {
				if !snow[a].off {
					rl.DrawTexturePro(imgs, snow[a].img, snow[a].rec, snow[a].ori, snow[a].ro, rl.Fade(snow[a].col, snow[a].fade))

					snow[a].ro += 2
					snow[a].rec.Y += 2

					if snow[a].rec.Y > scrHF32 {
						snow[a].off = true
						clear = true
					}
				}
			}
			if clear {
				for a := 0; a < len(snow); a++ {
					if snow[a].off {
						snow = remImg(snow, a)
					}
				}
			}
		}

		//INVENTORY
		if len(inven) > 0 {
			drawInven()
		}
		//PLAYER INFO
		drawPlayerInfo()

		if debug {
			//NEXT ROOM DOOR NUMS
			for a := 0; a < len(level[roomNum].doorSides); a++ {
				if level[roomNum].doorSides[a] == 1 {
					rl.DrawText("up  "+fmt.Sprint(level[roomNum].nextRooms[a]), levRec.ToInt32().X+levRec.ToInt32().Width/2, levRec.ToInt32().Y+bsUi32, txU2, rl.White)
				}
				if level[roomNum].doorSides[a] == 2 {
					rl.DrawText("right"+fmt.Sprint(level[roomNum].nextRooms[a]), levRec.ToInt32().X+levRec.ToInt32().Width-bsU2i32, levRec.ToInt32().Y+levRec.ToInt32().Width/2, txU2, rl.White)
				}
				if level[roomNum].doorSides[a] == 3 {
					rl.DrawText("down  "+fmt.Sprint(level[roomNum].nextRooms[a]), levRec.ToInt32().X+levRec.ToInt32().Width/2, levRec.ToInt32().Y+levRec.ToInt32().Width-bsU2i32, txU2, rl.White)
				}
				if level[roomNum].doorSides[a] == 4 {
					rl.DrawText("left  "+fmt.Sprint(level[roomNum].nextRooms[a]), levRec.ToInt32().X+bsU2i32, levRec.ToInt32().Y+levRec.ToInt32().Width/2, txU2, rl.White)
				}
			}
			//DOOR EXIT RECS

			for a := 0; a < len(level[roomNum].doorExitRecs); a++ {
				rl.DrawRectangleLinesEx(level[roomNum].doorExitRecs[a], 0.5, rl.Magenta)
			}

			//LEVEL BORDER RECS
			rl.DrawRectangleLinesEx(levRec, 2, rl.Green)
			rl.DrawRectangleLinesEx(levRecInner, 2, rl.Magenta)
		}
	}

	//ARTIFACTS
	if artifactson {
		num := 100

		for {
			x := (levX - bsU6) + rF32(0, levW+bsU12)
			y := levY + rF32(0, levW)
			siz := rF32(1, 3)
			rec := rl.NewRectangle(x, y, siz, siz)
			rl.DrawRectangleRec(rec, rl.Black)

			num--
			if num == 0 {
				break
			}
		}
	}

}
func drawnextlevelscreen() { //MARK:DRAW NEXT LEVEL SCREEN

	txt := "prepare for level " + fmt.Sprint(levelnum)
	txtlen := rl.MeasureText(txt, 40)
	x := int32(cnt.X) - txtlen/2
	y := int32(cnt.Y - 20)

	rl.DrawText(txt, x, y, 40, rl.White)

	txt = "press space or button to continue"
	txtlen = rl.MeasureText(txt, 20)
	x = int32(cnt.X) - txtlen/2
	y = int32(cnt.Y + 30)

	rl.DrawText(txt, x, y, 20, rl.White)

	if nextlevelT > 0 {
		nextlevelT--
	} else {
		startdmgT = fps * 5
		if rl.IsKeyPressed(rl.KeySpace) {
			nextlevelscreen = false
			pause = false
		}
		if useController {
			if rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
				nextlevelscreen = false
				pause = false
			}
		}
	}

}
func drawEndGame() { //MARK:DRAW END GAME
	pause = true
	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)
	shaderon = false
	if endPauseT > 0 {
		endPauseT--
	}
	if endgameT > 0 {
		endgameT--
		if endgopherrec.Y > levRec.Y+levRec.Height-endgopherrec.Height {
			endgopherrec.Y -= 2
		}
	}

	rl.DrawTexturePro(imgs, etc[55], endgopherrec, rl.Vector2Zero(), 0, rl.White)
	rl.PlaySound(sfx[19])
	txt := "you"
	txtlen := rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2-3, int32(cnt.Y)-txU8+3, txU8, rl.Black)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2, int32(cnt.Y)-txU8, txU8, rl.White)
	txt = "win"
	txtlen = rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2-3, int32(cnt.Y)+txU+3, txU8, rl.Black)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2, int32(cnt.Y)+txU, txU8, rl.White)

	if keypressT == 0 {
		if endPauseT == 0 {
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) || endgameT == 0 {
				endgame = false
				endgameT = 0
				endgopherrec = rl.NewRectangle(cnt.X-bsU4, levRec.Y+levRec.Height, bsU8, bsU8)
				restartgame()
			}
		}
	}
}
func drawShop() { //MARK:DRAW SHOP

	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)

	if rl.IsKeyPressed(rl.KeyA) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
		if optionT == 0 {
			optionT = fps / 5
			shopnum--
			if shopnum < 0 {
				shopnum = 4
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyD) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
		if optionT == 0 {
			optionT = fps / 5
			shopnum++
			if shopnum > 4 {
				shopnum = 0
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyS) || rl.GetGamepadAxisMovement(0, 1) > 0 || rl.IsGamepadButtonDown(0, 3) {
		if optionT == 0 {
			optionT = fps / 5
			if shopnum == 0 {
				shopnum = 2
			} else if shopnum == 2 {
				shopnum = 4
			} else if shopnum == 4 {
				shopnum = 0
			}
			if shopnum == 1 {
				shopnum = 3
			} else if shopnum == 3 {
				shopnum = 4
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.GetGamepadAxisMovement(0, 1) < 0 || rl.IsGamepadButtonDown(0, 1) {
		if optionT == 0 {
			optionT = fps / 5
			if shopnum == 0 {
				shopnum = 4
			} else if shopnum == 2 {
				shopnum = 0
			}
			if shopnum == 1 {
				shopnum = 4
			} else if shopnum == 3 {
				shopnum = 1
			}
			if shopnum == 4 {
				shopnum = 0
			}
		}
	}
	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {

		if shopnum == 4 {
			shopExitT = fps * 2
			shopon = false
			pl.cnt.Y = shopExitY
			upPlayerRec()
			pause = false
		} else {
			if mods.wallet {
				mods.wallet = false
				clearinven("wallet")
				shopitems[shopnum].shopoff = true
				addshopitem(shopnum)
				rl.PlaySound(sfx[21])
			} else if !shopitems[shopnum].shopoff && pl.coins >= shopitems[shopnum].shopprice {
				if !shopitems[shopnum].shopoff {
					shopitems[shopnum].shopoff = true
					pl.coins -= shopitems[shopnum].shopprice
				}
				shopitems[shopnum].shopoff = true
				addshopitem(shopnum)
				rl.PlaySound(sfx[21])
			} else {
				rl.PlaySound(sfx[22])
			}
		}
	}

	txt := "shop"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(levY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	siz := bsU5
	y := levY + bsU6
	x := cnt.X
	x -= siz * 2

	rec := rl.NewRectangle(x, y, siz, siz)
	if shopnum == 0 {
		if shopitems[0].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, fadeblink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, fadeblink))
		}
	}
	if shopitems[0].shopoff {
		col := ranRed()
		rl.DrawTexturePro(imgs, shopitems[0].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(imgs, shopitems[0].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {

		rl.DrawTexturePro(imgs, shopitems[0].img, rec, rl.Vector2Zero(), 0, shopitems[0].color)
		rl.DrawTexturePro(imgs, shopitems[0].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(shopitems[0].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(shopitems[0].shopprice), txtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(shopitems[0].shopprice), txtx, txty, txtSize, rl.White)
		rl.DrawTexturePro(imgs, coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(shopitems[0].name, txtSize)
	rl.DrawText(shopitems[0].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, txtSize, rl.White)

	rec.X += (siz * 2) + siz/2
	if shopnum == 1 {
		if shopitems[1].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, fadeblink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, fadeblink))
		}
	}

	if shopitems[1].shopoff {
		col := ranRed()
		rl.DrawTexturePro(imgs, shopitems[1].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(imgs, shopitems[1].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {
		rl.DrawTexturePro(imgs, shopitems[1].img, rec, rl.Vector2Zero(), 0, shopitems[1].color)
		rl.DrawTexturePro(imgs, shopitems[1].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(shopitems[1].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(shopitems[1].shopprice), txtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(shopitems[1].shopprice), txtx, txty, txtSize, rl.White)
		rl.DrawTexturePro(imgs, coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(shopitems[1].name, txtSize)
	rl.DrawText(shopitems[1].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, txtSize, rl.White)

	rec.Y += siz * 2
	if shopnum == 3 {
		if shopitems[3].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, fadeblink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, fadeblink))
		}
	}

	if shopitems[3].shopoff {
		col := ranRed()
		rl.DrawTexturePro(imgs, shopitems[3].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(imgs, shopitems[3].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {
		rl.DrawTexturePro(imgs, shopitems[3].img, rec, rl.Vector2Zero(), 0, shopitems[3].color)
		rl.DrawTexturePro(imgs, shopitems[3].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(shopitems[3].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(shopitems[3].shopprice), txtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(shopitems[3].shopprice), txtx, txty, txtSize, rl.White)
		rl.DrawTexturePro(imgs, coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(shopitems[3].name, txtSize)
	rl.DrawText(shopitems[3].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, txtSize, rl.White)

	rec.X -= (siz * 2) + siz/2
	if shopnum == 2 {
		if shopitems[2].shopoff {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Red, fadeblink))
		} else {
			rl.DrawRectangleRec(rec, rl.Fade(rl.Green, fadeblink))
		}
	}

	if shopitems[2].shopoff {
		col := ranRed()
		rl.DrawTexturePro(imgs, shopitems[2].img, rec, rl.Vector2Zero(), 0, col)
		rl.DrawTexturePro(imgs, shopitems[2].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(col, 0.2))
	} else {
		rl.DrawTexturePro(imgs, shopitems[2].img, rec, rl.Vector2Zero(), 0, shopitems[2].color)
		rl.DrawTexturePro(imgs, shopitems[2].img, BlurRec(rec, 2), rl.Vector2Zero(), 0, rl.Fade(shopitems[2].color, 0.2))
		coinx := rec.X + rec.Width + bsU
		coiny := rec.Y + bsU
		txtlen = rl.MeasureText("x"+fmt.Sprint(shopitems[2].shopprice), txtSize)
		txtx = int32(coinx+siz/4) - txtlen/2
		txty = int32(coiny+siz/2) + bsUi32/3
		rl.DrawText("x"+fmt.Sprint(shopitems[2].shopprice), txtx, txty, txtSize, rl.White)
		rl.DrawTexturePro(imgs, coin, rl.NewRectangle(coinx, coiny, siz/2, siz/2), ori, 0, rl.White)
	}
	txtlen = rl.MeasureText(shopitems[2].name, txtSize)
	rl.DrawText(shopitems[2].name, rec.ToInt32().X+rec.ToInt32().Width/2-txtlen/2, rec.ToInt32().Y+rec.ToInt32().Height+bsUi32/4, txtSize, rl.White)

	//WALLET
	if mods.wallet {
		walletx := cnt.X - siz*2 + siz/8
		wallety := rec.Y + rec.Height + bsU4 + siz/8
		rl.DrawTexturePro(imgs, etc[11], rl.NewRectangle(walletx, wallety, siz-siz/4, siz-siz/4), ori, 0, ranBrown())
	}
	//PL COINS
	coinx := cnt.X - siz
	coiny := rec.Y + rec.Height + bsU4
	rl.DrawTexturePro(imgs, coin, rl.NewRectangle(coinx, coiny, siz, siz), ori, 0, rl.White)

	txtx = int32(coinx + siz)
	txty = int32(coiny + siz/4)
	rl.DrawText("x"+fmt.Sprint(pl.coins), txtx, txty, txtSize*2, rl.White)

	if frames%6 == 0 {
		coin.X += 16
		if coin.X >= 1200 {
			coin.X = 1120
		}
	}

	//EXIT
	txtlen = rl.MeasureText("exit", txtSize*2)
	txtx = int32(cnt.X) - txtlen/2
	txty = int32(coiny + siz + bsU2)

	wid := float32(txtlen) + bsU2
	heig := float32(txtSize*2) + bsU/2

	rec = rl.NewRectangle(cnt.X-wid/2, float32(txty)-bsU/4, wid, heig)

	if shopnum == 4 {
		rl.DrawRectangleRec(rec, ranRed())
	} else {
		rl.DrawRectangleLinesEx(rec, 2, ranCol())
	}

	rl.DrawText("exit", txtx-2, txty+2, txtSize*2, rl.Black)
	rl.DrawText("exit", txtx, txty, txtSize*2, rl.White)

}
func drawHelp() { //MARK:DRAW HELP

	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)
	txt := "help"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(levY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	txtx = int32(cnt.X - bsU9)
	txty += txU7

	txt = "five levels collect power ups"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "kill enemies avoid traps"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "last level defeat boss"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txty += txtSize + txU/4
	txt = "WASD keys / xbox left stick/dpad > move"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "SPACE key / xbox a > attack"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "TAB key / xbox y > inventory"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "RIGHT CTRL key / xbox b > map"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "ESC key / xbox menu > options/exit"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "END key > exits game"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		helpon = false
	}

}
func drawDied() { //MARK:DRAW DIED

	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)

	if diedscrT > 0 {
		diedscrT--
	}

	rl.DrawTexturePro(imgs, diedIMG, diedRec, rl.Vector2Zero(), 0, ranRed())
	rl.DrawTexturePro(imgs, diedIMG, BlurRec(diedRec, 10), rl.Vector2Zero(), 0, rl.Fade(ranRed(), rF32(0.1, 0.4)))
	diedRec.X -= 2
	diedRec.Y -= 2
	diedRec.Width += 4
	diedRec.Height += 4

	if diedRec.Y <= levRecInner.Y || rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) && diedscrT == 0 {
		died = false
		besttime = false
		timeson = true
		besttimesT = fps
	}
	txt := "you"
	txtlen := rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2-3, int32(cnt.Y)-txU8+3, txU8, rl.Black)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2, int32(cnt.Y)-txU8, txU8, rl.White)
	txt = "died"
	txtlen = rl.MeasureText(txt, txU8)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2-3, int32(cnt.Y)+txU+3, txU8, rl.Black)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2, int32(cnt.Y)+txU, txU8, rl.White)

	txt = "new best time"
	txtlen = rl.MeasureText(txt, txU4)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2-3, scrH32-txU5+3, txU4, rl.Black)
	rl.DrawText(txt, int32(cnt.X)-txtlen/2, scrH32-txU5, txU4, rl.White)

}
func drawTimes() { //MARK:DRAW TIMES

	if besttimesT > 0 {
		besttimesT--
	}

	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)
	txt := "best times"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(levY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	txty += txU7

	for i := 0; i < len(times); i++ {
		minutes, seconds := times[i]/60, times[i]%60
		minTXT := fmt.Sprint(minutes)
		secsTXT := fmt.Sprint(seconds)
		if seconds < 10 {
			if seconds == 0 {
				secsTXT = "00"
			} else {
				secsTXT = "0" + secsTXT
			}
		}
		if minutes < 10 {
			if minutes == 0 {
				minTXT = "00"
			} else {
				minTXT = "0" + minTXT
			}
		}
		timesTXT := minTXT + ":" + secsTXT
		txtlen := rl.MeasureText(timesTXT, txtSize*2)
		rl.DrawText(timesTXT, int32(cnt.X)-txtlen/2, txty, txtSize*2, rl.White)
		txty += txtSize*2 + txU/2
	}

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		if besttimesT <= 0 {
			timeson = false
			if !optionson {
				restartgame()
			}
		}
	}

}

func drawCredits() { //MARK:DRAW CREDITS

	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)
	txt := "credits"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(levY) + txU
	rl.DrawText(txt, txtx, txty, txU5, rl.White)

	txtx = int32(cnt.X - bsU9)
	txty += txU7

	txt = "kenney.nl"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "laredgames.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "pixelfelix.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "piiixl.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "stealthix.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "rad-potato.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "pixelfrog-assets.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "free-game-assets.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "bdragon1727.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "kamioo.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "luquigames.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "nebelstern.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "bit-by-bit-sound.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "ironchestgames.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "pixeljad.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "magory.itch.io"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "opengameart.org/users/subspaceaudio"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)
	txty += txtSize + txU/4
	txt = "opengameart.org/users/rubberduck"
	rl.DrawText(txt, txtx, txty, txtSize, rl.White)

	if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		creditson = false
	}

}
func drawresettimes() { //MARK:DRAW RESET TIMES
	txt := "reset times"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(cnt.Y) - txU8
	rl.DrawText(txt, txtx, txty, txU5, rl.White)
	txty += txU6
	rec := rl.NewRectangle(cnt.X-bsU3, float32(txty), bsU3, bsU2)
	recX := rec.ToInt32().X
	if exitLR {
		rec.X += rec.Width
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			for i := 0; i < len(times); i++ {
				times[i] = 600
			}
			savetimes()
			resettimes = false
		}
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleRec(rec, rl.Red)
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			resettimes = false
		}
	}

	txtlen = rl.MeasureText("Y", txU3)
	txtx = recX + rec.ToInt32().Width + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("Y", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("Y", txtx, txty, txU3, rl.White)

	txtlen = rl.MeasureText("N", txU3)
	txtx = recX + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("N", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("N", txtx, txty, txU3, rl.White)

	if rl.IsKeyPressed(rl.KeyA) || rl.GetGamepadAxisMovement(0, 0) < 0 && rl.GetGamepadAxisMovement(0, 0) > -0.3 || rl.IsGamepadButtonDown(0, 4) {
		if optionT == 0 {
			optionT = fps / 5
			exitLR = !exitLR
		}
	} else if rl.IsKeyPressed(rl.KeyD) || rl.GetGamepadAxisMovement(0, 0) > 0 && rl.GetGamepadAxisMovement(0, 0) < 0.3 || rl.IsGamepadButtonDown(0, 2) {
		if optionT == 0 {
			optionT = fps / 5
			exitLR = !exitLR
		}
	}

	if rl.IsKeyPressed(rl.KeyY) {
		for i := 0; i < len(times); i++ {
			times[i] = 600
		}
		savetimes()
		resettimes = false
	} else if rl.IsKeyPressed(rl.KeyN) {
		resettimes = false
	}

}
func drawrestartconfirm() { //MARK:DRAW RESTART CONFIRM
	txt := "restart"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(cnt.Y) - txU8
	rl.DrawText(txt, txtx, txty, txU5, rl.White)
	txty += txU6
	rec := rl.NewRectangle(cnt.X-bsU3, float32(txty), bsU3, bsU2)
	recX := rec.ToInt32().X
	if exitLR {
		rec.X += rec.Width
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			restartgame()
			optionson = false
			restarton = false
		}
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleRec(rec, rl.Red)
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			restarton = false
		}
	}

	txtlen = rl.MeasureText("Y", txU3)
	txtx = recX + rec.ToInt32().Width + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("Y", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("Y", txtx, txty, txU3, rl.White)

	txtlen = rl.MeasureText("N", txU3)
	txtx = recX + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("N", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("N", txtx, txty, txU3, rl.White)

	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 && rl.GetGamepadAxisMovement(0, 0) > -0.3 || rl.IsGamepadButtonDown(0, 4) {
		if optionT == 0 {
			optionT = fps / 5
			exitLR = !exitLR
		}
	} else if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 && rl.GetGamepadAxisMovement(0, 0) < 0.3 || rl.IsGamepadButtonDown(0, 2) {
		if optionT == 0 {
			optionT = fps / 5
			exitLR = !exitLR
		}
	}

	if rl.IsKeyPressed(rl.KeyY) {
		restartgame()
		optionson = false
		restarton = false
	} else if rl.IsKeyPressed(rl.KeyN) {
		restarton = false
	}

}
func drawExit() { //MARK:DRAW EXIT

	txt := "exit"
	txtlen := rl.MeasureText(txt, txU5)
	txtx := int32(cnt.X) - txtlen/2
	txty := int32(cnt.Y) - txU8
	rl.DrawText(txt, txtx, txty, txU5, rl.White)
	txty += txU6
	rec := rl.NewRectangle(float32(txtx), float32(txty), bsU3, bsU2)
	recX := rec.ToInt32().X
	if exitLR {
		rec.X += rec.Width
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			exitgame()
		}
		rl.DrawRectangleRec(rec, rl.Green)
	} else {
		rl.DrawRectangleRec(rec, rl.Red)
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			exiton = false
		}
	}

	txtlen = rl.MeasureText("Y", txU3)
	txtx = recX + rec.ToInt32().Width + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("Y", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("Y", txtx, txty, txU3, rl.White)

	txtlen = rl.MeasureText("N", txU3)
	txtx = recX + rec.ToInt32().Width/2 - txtlen/2
	txty = rec.ToInt32().Y + 2
	rl.DrawText("N", txtx-2, txty+2, txU3, rl.Black)
	rl.DrawText("N", txtx, txty, txU3, rl.White)

	if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 && rl.GetGamepadAxisMovement(0, 0) > -0.3 || rl.IsGamepadButtonDown(0, 4) {
		if optionT == 0 {
			optionT = fps / 5
			exitLR = !exitLR
		}
	} else if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 && rl.GetGamepadAxisMovement(0, 0) < 0.3 || rl.IsGamepadButtonDown(0, 2) {
		if optionT == 0 {
			optionT = fps / 5
			exitLR = !exitLR
		}
	}

	if rl.IsKeyPressed(rl.KeyY) {
		exitgame()
	} else if rl.IsKeyPressed(rl.KeyN) {
		exiton = false
	}

}
func drawOptions() { //MARK:DRAW OPTIONS
	//rl.ShowCursor()

	rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)
	if creditson {
		drawCredits()
	} else if timeson {
		drawTimes()
	} else if helpon {
		drawHelp()
	} else if exiton {
		drawExit()
	} else if resettimes {
		drawresettimes()
	} else if restarton {
		drawrestartconfirm()
	} else {

		txt := "options"
		txtlen := rl.MeasureText(txt, txU5)
		txtx := int32(cnt.X) - txtlen/2
		txty := int32(levY) + txU
		rl.DrawText(txt, txtx, txty, txU5, rl.White)

		txty += txU7
		txtx = int32(cnt.X - bsU7)
		onoffx := txtx + int32(levRec.Width/3) - txtSize*2

		rec := rl.NewRectangle(float32(txtx)-bsU/2, float32(txty)-bsU/4, bsU*17, bsU2-bsU/4)
		rec.Y += float32(optionnum) * float32(txtSize+txtSize/2)
		if optionnum == 10 || optionnum == 11 || optionnum == 12 || optionnum == 13 || optionnum == 14 || optionnum == 15 || optionnum == 16 {
			rec.Y += float32(txtSize + txtSize/2)
		}

		//KEYS GAMEPAD INP
		if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) || rl.GetGamepadAxisMovement(0, 1) < 0 || rl.IsGamepadButtonDown(0, 1) {
			if optionT == 0 {
				optionT = fps / 5
				optionnum--
				if optionnum < 0 {
					optionnum = 16
				}
			}
		}
		if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) || rl.GetGamepadAxisMovement(0, 1) > 0 || rl.IsGamepadButtonDown(0, 3) {
			if optionT == 0 {
				optionT = fps / 5
				optionnum++
				if optionnum > 16 {
					optionnum = 0
				}
			}
		}
		if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
			switch optionnum {
			case 0:
				hpBarsOn = !hpBarsOn
			case 1:
				scanlineson = !scanlineson
			case 2:
				artifactson = !artifactson
			case 3:
				shaderon = !shaderon
			case 4:
				platkrecon = !platkrecon
			case 5:
				invincible = !invincible
			case 6:
				if isController && useController {
					contolleron = false
					useController = false
				} else if isController && !useController {
					contolleron = true
					useController = true
				} else if !isController {
					contolleron = false
					useController = false
				}
			case 7:
				if musicon {
					musicon = false
				} else {
					rl.StopMusicStream(music)
					music = backMusic[bgMusicNum]
					music.Looping = true
					musicon = true
					rl.PlayMusicStream(music)
					musicon = true
				}

			case 10:
				restarton = true
			case 11:
				timeson = true
			case 12:
				helpon = true
			case 13:
				creditson = true
			case 14:
				resettimes = true
			case 15:
				hardcore = !hardcore
				restartgame()
				optionson = false
			case 16:
				exiton = true
				exitLR = false
			}
			optionsChange = true
		}

		//OPTIONS LIST
		rl.DrawRectangleRec(rec, rl.Fade(ranCol(), fadeblink2))

		txt = "hp bars"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		hpBarsOn = onoff(onoffx, txty, float32(txtSize), hpBarsOn)
		txty += txtSize + txtSize/2

		txt = "scan lines"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		scanlineson = onoff(onoffx, txty, float32(txtSize), scanlineson)
		txty += txtSize + txtSize/2

		txt = "pixel artifacts"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		artifactson = onoff(onoffx, txty, float32(txtSize), artifactson)
		txty += txtSize + txtSize/2

		txt = "bloom 'fuzzy'"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		shaderon = onoff(onoffx, txty, float32(txtSize), shaderon)
		txty += txtSize + txtSize/2

		txt = "player atk range"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		platkrecon = onoff(onoffx, txty, float32(txtSize), platkrecon)
		txty += txtSize + txtSize/2

		txt = "invincible"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		invincible = onoff(onoffx, txty, float32(txtSize), invincible)
		txty += txtSize + txtSize/2

		txt = "use controller"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		useController = onoff(onoffx, txty, float32(txtSize), useController)
		if !isController && optionnum == 6 {
			txt = "no controller detected"
			rl.DrawText(txt, txtx, txty+(txtSize+txtSize/2)*4, txtSize, ranCol())
		}
		txty += txtSize + txtSize/2
		txt = "music"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		musicon = onoff(onoffx, txty, float32(txtSize), musicon)
		txty += txtSize + txtSize/2

		txt = "music track"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		bgMusicNum = int(updownswitch(onoffx, txty, float32(txtSize), float32(bgMusicNum), 2))
		txty += txtSize + txtSize/2

		//CHANGE MUSIC TRACK
		if optionnum == 8 {
			txt = "use left / right to adjust"
			rl.DrawText(txt, txtx, txty+txtSize+txtSize/2, txtSize, ranCol())
			if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
				if optionT == 0 {
					optionT = fps / 5
					bgMusicNum++
					if bgMusicNum > 2 {
						bgMusicNum = 0
					}
					rl.StopMusicStream(music)
					music = backMusic[bgMusicNum]
					music.Looping = true
					musicon = true
					rl.PlayMusicStream(music)
					optionsChange = true
				}
			} else if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
				if optionT == 0 {
					optionT = fps / 5
					bgMusicNum--
					if bgMusicNum < 0 {
						bgMusicNum = 2
					}
					rl.StopMusicStream(music)
					music = backMusic[bgMusicNum]
					music.Looping = true
					musicon = true
					rl.PlayMusicStream(music)
					optionsChange = true
				}
			}
		}

		txt = "volume - 0 is off"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		volume = updownswitch(onoffx, txty, float32(txtSize), volume, 1)
		txty += txtSize + txtSize/2

		//VOLUME LR
		if optionnum == 9 {
			txt = "use left / right to adjust"
			rl.DrawText(txt, txtx, txty, txtSize, ranCol())
			if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
				if optionT == 0 {
					optionT = fps / 5
					if volume < 1 {
						volume += 0.1
					}
					optionsChange = true
				}
			} else if rl.IsKeyPressed(rl.KeyA) || rl.IsKeyPressed(rl.KeyLeft) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
				if optionT == 0 {
					optionT = fps / 5
					if volume > 0 {
						volume -= 0.1
					}
					if volume < 0 {
						volume = 0
					}
					optionsChange = true
				}
			}
		}

		txty += txtSize + txtSize/2
		txt = "restart game"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		txty += txtSize + txtSize/2
		txt = "best times"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		txty += txtSize + txtSize/2
		txt = "help"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		txty += txtSize + txtSize/2
		txt = "credits"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		txty += txtSize + txtSize/2
		txt = "reset times"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		txty += txtSize + txtSize/2
		txt = "hardcore"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		hardcore = onoff(onoffx, txty, float32(txtSize), hardcore)
		if optionnum == 15 {
			txt = "more enemies > game will restart"
			rl.DrawText(txt, txtx, txty-(txtSize+txtSize/2)*6, txtSize, ranCol())
		}
		txty += txtSize + txtSize/2
		txt = "exit"
		rl.DrawText(txt, txtx, txty, txtSize, rl.White)
		txty += txtSize + txtSize/2

	}

}
func drawUpMario() { //MARK:DRAW UP MARIO

	//TIMER
	rl.DrawText(fmt.Sprint(marioT), marioScreenRec.ToInt32().X+bsUi32, marioScreenRec.ToInt32().Y+bsUi32, txtSize, rl.White)
	marioT--

	//INP
	if rl.IsKeyDown(rl.KeyD) || rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
		marioPL.X += 8
		marioV2L.X += 8
		marioV2R.X += 8
		marioImg.Y = knight[0].Y

	} else if rl.IsKeyDown(rl.KeyA) || rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
		marioPL.X -= 8
		marioV2L.X -= 8
		marioV2R.X -= 8
		marioImg.Y = knight[2].Y
	}
	if rl.IsKeyPressed(rl.KeyW) || rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
		if !mariojump {
			mariojump = true
			mariojumpT = fps / 3
		}
	}

	//DRAW PATTERN
	x := marioScreenRec.X
	y := marioScreenRec.Y
	siz := bsU10
	for {
		rl.DrawTexturePro(imgs, patternRec, rl.NewRectangle(x, y, siz, siz), ori, 0, rl.Fade(ranCol(), 0.05))
		x += siz
		if x >= marioScreenRec.X+marioScreenRec.Width {
			x = marioScreenRec.X
			y += siz
		}
		if y >= marioScreenRec.Y+marioScreenRec.Height {
			break
		}
	}

	//DRAW BLOKS
	for a := 0; a < len(mariorecs); a++ {
		rl.DrawTexturePro(imgs, wallT, mariorecs[a], ori, 0, marioCols[a])
		rl.DrawTexturePro(imgs, wallT, BlurRec(mariorecs[a], 2), ori, 0, rl.Fade(marioCols[a], 0.2))
	}

	//DRAW COINS
	for i := 0; i < len(mariocoinonoff); i++ {
		if mariocoinonoff[i] {
			rl.DrawTexturePro(imgs, coin, mariocoins[i], ori, 0, rl.White)
			if marioT%6 == 0 {
				coin.X += 16
				if coin.X >= 1200 {
					coin.X = 1120
				}
				if rl.CheckCollisionRecs(marioPL, mariocoins[i]) {
					if mariocoinonoff[i] {
						pl.coins++
						mariocoinonoff[i] = false
						rl.PlaySound(sfx[18])
					}
				}
			}
		}

	}

	//DRAW PLAYER
	drec := marioPL
	drec.X -= drec.Width / 4
	drec.Width += drec.Width / 2
	drec.Y -= drec.Height / 4
	drec.Height += drec.Height / 2
	//ori2 := rl.NewVector2(drec.Width/2,drec.Height/2)
	rl.DrawTexturePro(imgs, marioImg, drec, ori, 0, rl.White)
	if debug {
		rl.DrawRectangleLinesEx(marioPL, 1, rl.White)
	}

	if frames%4 == 0 {
		marioImg.X += pl.sizImg
	}
	if marioImg.X > pl.imgWalkX+(float32(pl.framesWalk-1)*pl.sizImg) {
		marioImg.X = pl.imgWalkX
	}

	//JUMP FALL
	if mariojumpT > 0 {
		mariojumpT--
		marioPL.Y -= 12
		marioV2L.Y -= 12
		marioV2R.Y -= 12
	} else {
		collides := false
		for a := 0; a < len(mariorecs); a++ {
			if rl.CheckCollisionPointRec(marioV2L, mariorecs[a]) || rl.CheckCollisionPointRec(marioV2R, mariorecs[a]) {
				collides = true
				mariojump = false
			}
		}
		if !collides {
			marioPL.Y += 8
			marioV2L.Y += 8
			marioV2R.Y += 8
		}

	}

	//EXIT SCREEN
	if marioT == 0 || marioPL.X+marioPL.Width < marioScreenRec.X || marioPL.X > marioScreenRec.X+marioScreenRec.Width {
		marioon = false
		pause = false
	}

}
func drawInven() { //MARK:DRAW INVENTORY

	x := levX - bsU2
	y := levY + bsU/2
	siz := bsU + bsU/2

	for a := 0; a < len(inven); a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		rl.DrawTexturePro(imgs, inven[a].img, rec, ori, 0, inven[a].color)

		rl.DrawTexturePro(imgs, inven[a].img, BlurRec(rec, 1), ori, 0, rl.Fade(inven[a].color, rF32(0.1, 0.2)))

		if inven[a].numof > 1 {
			txtx := int32(x+siz) - txU/2
			txty := int32(y+siz) - txU/2
			rl.DrawText(fmt.Sprint(inven[a].numof), txtx-2, txty-2, txU, rl.Black)
			rl.DrawText(fmt.Sprint(inven[a].numof), txtx, txty, txU, rl.White)
		}

		y += siz + bsU/4
	}

}
func drawInvenDetail() { //MARK:DRAW INVENTORY DETAIL

	if len(inven) > 0 {

		x := levX
		y := levY + bsU/4
		siz := bsU

		txtcntr := int32(cnt.X)
		txty := int32(y) + txU/4

		for a := 0; a < len(inven); a++ {

			txt := inven[a].name + " - " + inven[a].desc
			txtlen := rl.MeasureText(txt, txU)

			txtx := txtcntr - txtlen/2
			x = float32(txtx) - (siz + bsU/2)

			rl.DrawText(txt, txtx-1, txty+1, txU, rl.Black)
			rl.DrawText(txt, txtx, txty, txU, rl.White)

			rl.DrawTexturePro(imgs, inven[a].img, rl.NewRectangle(x, y, siz, siz), ori, 0, inven[a].color)

			if inven[a].numof > 1 {
				txtx2 := int32(x+siz) - txU/2
				txty2 := int32(y+siz) - txU/2
				rl.DrawText(fmt.Sprint(inven[a].numof), txtx2-2, txty2-2, txU, rl.Black)
				rl.DrawText(fmt.Sprint(inven[a].numof), txtx2, txty2, txU, rl.White)
			}

			y += siz + bsU/4
			txty += int32(siz + bsU/4)

		}

	} else {

		txt := "you have nothing..."
		txtlen := rl.MeasureText(txt, txU8)
		rl.DrawText(txt, int32(cnt.X)-txtlen/2, int32(cnt.Y)-txU8, txU8, rl.White)
		txt = "find something"
		txtlen = rl.MeasureText(txt, txU8)
		rl.DrawText(txt, int32(cnt.X)-txtlen/2, int32(cnt.Y)+txU, txU8, rl.White)
	}

}
func drawPlayerInfo() { //MARK:DRAW PLAYER INFO

	//TIMER
	y := levY + bsU/4
	txtx := int32(levRec.X + levRec.Width + bsU/2)
	txty := int32(y)

	minT := "0"
	if mins == 0 {
		minT = "00"
	} else if mins < 10 {
		minT = "0" + fmt.Sprint(mins)
	} else {
		minT = fmt.Sprint(mins)
	}
	secsT := "00"
	if secs > 0 && secs < 10 {
		secsT = "0" + fmt.Sprint(secs)
	} else if secs > 9 {
		secsT = fmt.Sprint(secs)
	}
	txt := minT + ":" + secsT
	rl.DrawText(txt, txtx, txty, txU3, rl.White)

	//HP ARMOR
	x := levX + levW + bsU/2
	y = levY + bsU2 + bsU/2
	siz := bsU + bsU/2

	for a := 0; a < pl.hpmax; a++ {
		rl.DrawTexturePro(imgs, etc[2], rl.NewRectangle(x, y, siz, siz), ori, 0, rl.DarkGray)
		y += siz
	}
	yorig := y

	y = levY + bsU2 + bsU/2
	for a := 0; a < pl.hp; a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		if pl.poison {
			rl.DrawTexturePro(imgs, etc[2], rec, ori, 0, rl.Green)
			rl.DrawTexturePro(imgs, etc[2], BlurRec(rec, 2), ori, 0, rl.Fade(rl.Green, rF32(0.1, 0.3)))
		} else {
			rl.DrawTexturePro(imgs, etc[2], rec, ori, 0, rl.Red)
			rl.DrawTexturePro(imgs, etc[2], BlurRec(rec, 2), ori, 0, rl.Fade(rl.Red, rF32(0.1, 0.3)))
		}
		y += siz
	}

	y = yorig + bsU/2

	for a := 0; a < pl.armorMax; a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		rl.DrawTexturePro(imgs, etc[38], rec, ori, 0, rl.DarkGray)
		y += siz
	}

	y = yorig + bsU/2

	for a := 0; a < pl.armor; a++ {
		rec := rl.NewRectangle(x, y, siz, siz)
		col := ranCyan()
		rl.DrawTexturePro(imgs, etc[38], rec, ori, 0, col)
		rl.DrawTexturePro(imgs, etc[38], BlurRec(rec, 2), ori, 0, rl.Fade(col, rF32(0.1, 0.3)))
		y += siz

	}

	//COINS
	y = levY + levW - (bsU2 + bsU/2)
	siz = bsU2
	rl.DrawTexturePro(imgs, coin, rl.NewRectangle(x, y, siz, siz), ori, 0, rl.White)
	if frames%6 == 0 {
		coin.X += 16
		if coin.X >= 1200 {
			coin.X = 1120
		}
	}

	txt = "x "
	txtx = int32(x+siz) + bsUi32/2
	txty = int32(y)
	rl.DrawText(txt, txtx, txty, txU3, rl.White)
	txtlen := rl.MeasureText(txt, txU3)
	txtx += txtlen
	txty += 2
	txt = fmt.Sprint(pl.coins)
	rl.DrawText(txt, txtx, txty, txU3, rl.White)

	//TXT SOLD
	if len(txtSoldlist) > 0 {
		clear := false
		for a := 0; a < len(txtSoldlist); a++ {
			if txtSoldlist[a].onoff {
				rl.DrawText(txtSoldlist[a].txt, txtSoldlist[a].x, txtSoldlist[a].y, txU2, rl.Fade(txtSoldlist[a].col, txtSoldlist[a].fade))

				rl.DrawText(txtSoldlist[a].txt2, txtSoldlist[a].x, txtSoldlist[a].y+txU2, txU2, rl.Fade(txtSoldlist[a].col, txtSoldlist[a].fade))

				txtSoldlist[a].y--
				txtSoldlist[a].fade -= 0.01
				if txtSoldlist[a].fade <= 0 {
					txtSoldlist[a].onoff = false
				}
			} else {
				clear = true
			}
		}

		if clear {
			for a := 0; a < len(txtSoldlist); a++ {
				if !txtSoldlist[a].onoff {
					txtSoldlist = remTxt(txtSoldlist, a)
				}
			}
		}
	}

}
func drawChainLight() { //MARK:DRAW CHAIN LIGHTNING

	for a := 1; a < len(chainV2); a++ {
		rl.DrawLineEx(chainV2[a], chainV2[a-1], rF32(2, 12), rl.Fade(ranCyan(), rF32(0.2, 0.5)))
	}
	chainV2 = nil
	for a := 0; a < len(level[roomNum].enemies); a++ {
		if level[roomNum].enemies[a].cnt.X > levRecInner.X && level[roomNum].enemies[a].cnt.Y > levRecInner.Y {
			chainV2 = append(chainV2, level[roomNum].enemies[a].cnt)
		}
	}
	chainLightTimer--

	if chainLightTimer <= 0 {
		for a := 0; a < len(level[roomNum].enemies); a++ {
			level[roomNum].enemies[a].hppause = fps / 2
			level[roomNum].enemies[a].hp -= 1
			if level[roomNum].enemies[a].hp <= 0 {
				cntr := level[roomNum].enemies[a].cnt
				addkill(a)
				level[roomNum].enemies[a].off = true
				makeFX(2, cntr)
			} else {
				playenemyhit()
			}
		}
		chainLightOn = false
	}

}
func drawPlayer() { //MARK:DRAW PLAYER

	drawRec := pl.rec
	drawRec.X += pl.rec.Width / 2
	drawRec.Y += pl.rec.Height / 2
	shadowRec := drawRec
	shadowRec.X -= 7
	shadowRec.Y += 7

	//NOT MOVING BOUNCE
	if !pl.move && !pl.atk {
		if roll18() == 18 {
			drawRec.Y -= 2
			shadowRec.Y -= 2
		}
	}

	//ATK REC
	if platkrecon && pl.atkTimer > 0 {
		rl.DrawRectangleRec(pl.atkrec, rl.Fade(darkRed(), 0.3))
	}

	//ESCAPE VINE
	if pl.escape {
		siz := bsU2
		plVineRec = rl.NewRectangle(pl.cnt.X-siz/2, levRec.Y, siz, pl.rec.Y-levRec.Y)
		if debug {
			rl.DrawRectangleLinesEx(plVineRec, 0.5, rl.Red)
		}
		y := plVineRec.Y
		x := plVineRec.X
		for {
			rl.DrawTexturePro(imgs, etc[6], rl.NewRectangle(x, y, siz, siz), ori, 0, rl.DarkGreen)
			y += siz
			if y > plVineRec.Y+plVineRec.Height {
				break
			}
		}
	}
	//DRAW BUBBLES
	if mods.flood && pl.underWater {
		waterRec := rl.NewRectangle((pl.crec.X+pl.crec.Width/2)-bsU/2, pl.crec.Y-waterY, bsU, bsU)
		rl.DrawTexturePro(imgs, plantBull.recTL, waterRec, ori, 0, rl.Fade(ranCyan(), waterF))
		//rl.DrawRectangleRec(waterRec,ranCol())
		waterRec2 := waterRec
		change := rF32(2, bsU/2)
		waterRec2.Width -= change
		waterRec2.Height -= change

		if waterLR {
			waterRec2.X += bsU
		} else {
			waterRec2.X -= bsU
		}
		if waterUP {
			waterRec2.Y += bsU
		} else {
			waterRec2.Y -= bsU
		}

		rl.DrawTexturePro(imgs, plantBull.recTL, waterRec2, ori, 0, rl.Fade(ranCyan(), waterF))

		waterRec.X += rF32(-2, 2)
		waterY += rF32(2, 5)
		waterF -= 0.02
		if waterY > bsU10 {
			waterY = 0
			waterF = 1
			waterLR = flipcoin()
			waterUP = flipcoin()
		}

		if roll18() == 18 {
			rl.PlaySound(sfx[24])
		}

	}
	//DRAW HP HIT
	if pl.hppause != 0 && !pl.revived && pl.armorHit { //ARMOR
		hpRec := rl.NewRectangle((pl.crec.X+pl.crec.Width/2)-bsU/2, pl.crec.Y-hpHitY, bsU, bsU)
		rl.DrawTexturePro(imgs, etc[38], hpRec, ori, 0, rl.Fade(ranCyan(), hpHitF))
		hpHitY += 2
		hpHitF -= 0.05
	} else if pl.hppause != 0 && !pl.revived && !pl.armorHit { //HP
		hpRec := rl.NewRectangle((pl.crec.X+pl.crec.Width/2)-bsU/2, pl.crec.Y-hpHitY, bsU, bsU)
		rl.DrawTexturePro(imgs, etc[2], hpRec, ori, 0, rl.Fade(rl.Red, hpHitF))
		hpHitY += 2
		hpHitF -= 0.05
	} else if pl.hppause != 0 && pl.revived { //REVIVED
		txty := int32(pl.rec.Y-reviveY) - txU2
		txt := "revived"
		txtlen := rl.MeasureText(txt, txU2)
		txtx := int32(pl.rec.X+pl.rec.Width/2) - txtlen/2
		rl.DrawText(txt, txtx, txty, txU2, rl.Fade(rl.White, reviveF))
		reviveY += 2
		reviveF -= 0.05
	}
	if pl.peaceT > 0 {
		peceRec := rl.NewRectangle((pl.crec.X+pl.crec.Width/2)-bsU/2, pl.crec.Y-(bsU), bsU, bsU)
		rl.DrawTexturePro(imgs, etc[46], peceRec, ori, 0, rl.White)
	}
	//DRAW PLAYER IMG
	rl.DrawTexturePro(imgs, pl.img, shadowRec, pl.ori, pl.ro, rl.Fade(rl.Black, 0.7))
	if pl.hppause > 0 {
		rl.DrawTexturePro(imgs, pl.img, drawRec, pl.ori, pl.ro, ranCol())
	} else {
		rl.DrawTexturePro(imgs, pl.img, drawRec, pl.ori, pl.ro, rl.White)
	}
	//ORBITAL
	if mods.orbital {
		if mods.orbitalN >= 1 {
			if roomChanged {
				pl.orbital1 = rl.NewVector2(pl.cnt.X+bsU4, pl.cnt.Y+bsU4)
			} else {
				pl.angle = pl.angle * (math.Pi / 180)
				newx := float32(math.Cos(float64(pl.angle)))*(pl.orbital1.X-pl.cnt.X) - float32(math.Sin(float64(pl.angle)))*(pl.orbital1.Y-pl.cnt.Y) + pl.cnt.X
				newy := float32(math.Sin(float64(pl.angle)))*(pl.orbital1.X-pl.cnt.X) + float32(math.Cos(float64(pl.angle)))*(pl.orbital1.Y-pl.cnt.Y) + pl.cnt.Y
				pl.orbital1 = rl.NewVector2(newx, newy)
				pl.angle += 4
				if getabs(pl.orbital1.X-pl.cnt.X) > bsU4 {
					if pl.orbital1.X > pl.cnt.X {
						pl.orbital1.X -= bsU / 8
					} else {
						pl.orbital1.X += bsU / 8
					}
				}
				if getabs(pl.orbital1.Y-pl.cnt.Y) > bsU4 {
					if pl.orbital1.Y > pl.cnt.Y {
						pl.orbital1.Y -= bsU / 8
					} else {
						pl.orbital1.Y += bsU / 8
					}
				}
			}

			siz := bsU + bsU/2
			pl.orbrec1 = rl.NewRectangle(pl.orbital1.X-siz/2, pl.orbital1.Y-siz/2, siz, siz)
			rl.DrawTexturePro(imgs, pl.orbimg1, pl.orbrec1, ori, 0, rl.White)

			if frames%3 == 0 {
				pl.orbimg1.X += orbitalanim.W
				if pl.orbimg1.X > orbitalanim.xl+orbitalanim.frames*orbitalanim.W {
					pl.orbimg1.X = orbitalanim.xl
				}
			}

		}
		if mods.orbitalN == 2 {
			if roomChanged {
				pl.orbital2 = rl.NewVector2(pl.cnt.X-bsU7, pl.cnt.Y-bsU7)
			} else {
				pl.angle = pl.angle * (math.Pi / 180)
				newx := float32(math.Cos(float64(pl.angle)))*(pl.orbital2.X-pl.cnt.X) - float32(math.Sin(float64(pl.angle)))*(pl.orbital2.Y-pl.cnt.Y) + pl.cnt.X
				newy := float32(math.Sin(float64(pl.angle)))*(pl.orbital2.X-pl.cnt.X) + float32(math.Cos(float64(pl.angle)))*(pl.orbital2.Y-pl.cnt.Y) + pl.cnt.Y
				pl.orbital2 = rl.NewVector2(newx, newy)
				pl.angle += 6

				if getabs(pl.orbital2.X-pl.cnt.X) > bsU7 {
					if pl.orbital2.X > pl.cnt.X {
						pl.orbital2.X -= bsU / 8
					} else {
						pl.orbital2.X += bsU / 8
					}
				}
				if getabs(pl.orbital2.Y-pl.cnt.Y) > bsU7 {
					if pl.orbital2.Y > pl.cnt.Y {
						pl.orbital2.Y -= bsU / 8
					} else {
						pl.orbital2.Y += bsU / 8
					}
				}
			}

			siz := bsU + bsU/2
			pl.orbrec2 = rl.NewRectangle(pl.orbital2.X-siz/2, pl.orbital2.Y-siz/2, siz, siz)
			rl.DrawTexturePro(imgs, pl.orbimg2, pl.orbrec2, ori, 0, rl.White)

			if frames%3 == 0 {
				pl.orbimg2.X += orbitalanim.W
				if pl.orbimg2.X > orbitalanim.xl+orbitalanim.frames*orbitalanim.W {
					pl.orbimg2.X = orbitalanim.xl
				}
			}
		}

	}

	//COMPANIONS
	if mods.planty || mods.alien || mods.carrot {
		drawUpCompanions()
	}

	//NIGHT REC
	if night {
		rec := rl.NewRectangle(pl.cnt.X-etc[25].Width/2, pl.cnt.Y-etc[25].Height/2, etc[25].Width, etc[25].Height)
		rl.DrawTexturePro(imgs, etc[25], rec, ori, 0, rl.White)

		recT := rl.NewRectangle(levRec.X, levRec.Y, levRec.Width, rec.Y-levRec.Y)
		recB := rl.NewRectangle(levRec.X, rec.Y+rec.Height, levRec.Width, (levRec.Y+levRec.Height)-(rec.Y+rec.Height))
		recL := rl.NewRectangle(levRec.X, rec.Y, rec.X-levRec.X, rec.Height)
		recR := rl.NewRectangle(rec.X+rec.Width, rec.Y, (levRec.X+levRec.Width)-(rec.X+rec.Width), rec.Height)

		rl.DrawRectangleRec(recT, rl.Fade(rl.Black, 0.6))
		rl.DrawRectangleRec(recB, rl.Fade(rl.Black, 0.6))
		rl.DrawRectangleRec(recL, rl.Fade(rl.Black, 0.6))
		rl.DrawRectangleRec(recR, rl.Fade(rl.Black, 0.6))

	}

	//DEBUG
	if debug {
		rl.DrawRectangleLinesEx(pl.arec, 1, rl.Red)
		rl.DrawRectangleLinesEx(pl.crec, 1, rl.Blue)
		rl.DrawRectangleLinesEx(pl.atkrec, 1, rl.White)
		rl.DrawPixelV(pl.cnt, rl.Red)
	}

}

func drawUpCompanions() { //MARK:DRAW UP COMPANIONS

	if mods.carrot {

		mrcarrot.timer--
		if mrcarrot.timer == 0 {
			mrcarrot.timer = fps * rI32(1, 5)
			makeProjectile("mrcarrot")
		}
		//MOVE
		if checkNextMove(mrcarrot.rec, mrcarrot.velx, mrcarrot.vely, false) {
			mrcarrot.rec.X += mrcarrot.velx
			mrcarrot.rec.Y += mrcarrot.vely
		} else {
			mrcarrot.velx = rF32(-mrcarrot.vel, mrcarrot.vel)
			mrcarrot.vely = rF32(-mrcarrot.vel, mrcarrot.vel)
		}

		//IMG

		shadowRec := mrcarrot.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		if mrcarrot.velx > 0 {
			rl.DrawTexturePro(imgs, mrcarrot.imgr, shadowRec, ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(imgs, mrcarrot.imgr, mrcarrot.rec, ori, 0, rl.White)
		} else {
			rl.DrawTexturePro(imgs, mrcarrot.imgl, shadowRec, ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(imgs, mrcarrot.imgl, mrcarrot.rec, ori, 0, rl.White)
		}

		//ANIM
		if frames%6 == 0 {
			mrcarrot.imgl.X += mrcarrot.imgl.Width
			if mrcarrot.imgl.X > mrcarrot.imgl.Width*float32(mrcarrot.frames) {
				mrcarrot.imgl.X = 0
			}
			mrcarrot.imgr.X += mrcarrot.imgr.Width
			if mrcarrot.imgr.X > 228+(mrcarrot.imgr.Width*float32(mrcarrot.frames)) {
				mrcarrot.imgr.X = 228
			}
		}
	}

	if mods.alien {
		mralien.timer--
		if mralien.timer == 0 {
			mralien.timer = fps * rI32(3, 8)
			makeProjectile("mralien")
		}
		//MOVE
		if checkNextMove(mralien.rec, mralien.velx, mralien.vely, false) {
			mralien.rec.X += mralien.velx
			mralien.rec.Y += mralien.vely
		} else {
			mralien.velx = rF32(-mralien.vel, mralien.vel)
			mralien.vely = rF32(-mralien.vel, mralien.vel)
		}

		//IMG
		shadowRec := mralien.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		rl.DrawTexturePro(imgs, mralien.img, shadowRec, ori, 0, rl.Fade(rl.Black, 0.8))
		rl.DrawTexturePro(imgs, mralien.img, mralien.rec, ori, 0, rl.White)

		if getabs(mralien.velx) > getabs(mralien.vely) {
			if mralien.velx > 0 {
				mralien.img.Y = 898
			} else {
				mralien.img.Y = 834
			}
		} else {
			if mralien.vely > 0 {
				mralien.img.Y = 770
			} else {
				mralien.img.Y = 962
			}
		}

		if frames%3 == 0 {
			mralien.img.X += mralien.img.Width
			if mralien.img.X >= 1200 {
				mralien.img.X = 1008
			}
		}
	}

	if mods.planty {
		//MOVE
		if checkNextMove(mrplanty.rec, mrplanty.velx, mrplanty.vely, false) {
			mrplanty.rec.X += mrplanty.velx
			mrplanty.rec.Y += mrplanty.vely
		} else {
			mrplanty.velx = rF32(-mrplanty.vel, mrplanty.vel)
			mrplanty.vely = rF32(-mrplanty.vel, mrplanty.vel)
		}
		mrplanty.cnt = rl.NewVector2(mrplanty.rec.Width/2, mrplanty.rec.Height/2)

		//IMG
		shadowRec := mrplanty.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		if mrplanty.velx > 0 {
			rl.DrawTexturePro(imgs, mrplanty.imgr, shadowRec, ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(imgs, mrplanty.imgr, mrplanty.rec, ori, 0, rl.White)
		} else {
			rl.DrawTexturePro(imgs, mrplanty.imgl, shadowRec, ori, 0, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(imgs, mrplanty.imgl, mrplanty.rec, ori, 0, rl.White)
		}

		//ANIM
		if frames%6 == 0 {
			mrplanty.imgl.X += mrplanty.imgl.Width
			if mrplanty.imgl.X > mrplanty.imgl.Width*float32(mrplanty.frames) {
				mrplanty.imgl.X = 0
			}
			mrplanty.imgr.X += mrplanty.imgr.Width
			if mrplanty.imgr.X > 352+(mrplanty.imgr.Width*float32(mrplanty.frames)) {
				mrplanty.imgr.X = 352
			}
		}

		//BULLETS
		if frames%30 == 0 {
			makeProjectile("plantbull")
		}

	}

}

func drawnocamBG() { //MARK:DRAW NO CAM BACKGROUND

}

func drawnocam() { //MARK:DRAW NO CAM

	//INTRO
	if intro && !optionson {
		if introT1 > 0 {
			shaderon = false
			x := cnt.X - etc[56].Width/2
			y := cnt.Y - etc[56].Height/2
			rl.DrawTexturePro(imgs, etc[56], rl.NewRectangle(x, y, etc[56].Width, etc[56].Height), ori, 0, rl.Fade(rl.White, introF1))
			if introF1 < 1 {
				introF1 += 0.01
			}

			if introF1 > 0.5 {
				txt := "raylib.com"
				txtlen := rl.MeasureText(txt, 20)
				txtx := int32(cnt.X) - txtlen/2
				txty := int32(cnt.Y + (etc[56].Height / 2) + bsU/2)
				rl.DrawText(txt, txtx, txty, 20, rl.White)
			}
			introT1--
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) {
				introT1 = 0
			}
		} else if introT2 > 0 {
			shaderon = false
			x := cnt.X - etc[55].Width/2
			y := cnt.Y - etc[55].Height/2
			rl.DrawTexturePro(imgs, etc[55], rl.NewRectangle(x, y, etc[55].Width, etc[55].Height), ori, 0, rl.Fade(rl.White, introF2))
			if introF2 < 1 {
				introF2 += 0.01
			}
			if introF2 > 0.5 {
				txt := "go.dev"
				txtlen := rl.MeasureText(txt, 20)
				txtx := int32(cnt.X) - txtlen/2
				txty := int32(cnt.Y + (etc[55].Height / 2) + bsU/2)
				rl.DrawText(txt, txtx, txty, 20, rl.White)
			}
			introT2--
			if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) {
				introT2 = 0

			}
		} else {
			shaderon = true

			if introcount {

				introT3--
				if introT3 > fps*2 {
					txt := "3"
					txtlen := rl.MeasureText(txt, 200)
					txty := int32(cnt.Y) - 100
					txtx := int32(cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 200, rl.Green)
				} else if introT3 > fps {
					txt := "2"
					txtlen := rl.MeasureText(txt, 200)
					txty := int32(cnt.Y) - 100
					txtx := int32(cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 200, rl.Green)
				} else if introT3 > 0 {
					txt := "1"
					txtlen := rl.MeasureText(txt, 200)
					txty := int32(cnt.Y) - 100
					txtx := int32(cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 200, rl.Green)
				} else {
					intro = false
					pause = false
					runT = 0
					mins = 0
					secs = 0
					if musicon {
						rl.PlayMusicStream(music)
					}
				}

			} else {
				x := cnt.X - etc[54].Width/2
				y := cnt.Y - etc[54].Height/2
				rl.DrawTexturePro(imgs, etc[54], rl.NewRectangle(x, y, etc[54].Width, etc[54].Height), ori, 0, rl.Fade(rl.White, introF3))
				if introF3 < 1 {
					introF3 += 0.01
				} else {
					txt := "press space or a button to start"
					txtlen := rl.MeasureText(txt, 20)
					txty := int32(y+etc[54].Height) + bsU2i32
					txtx := int32(cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 20, rl.Green)

					txt = "press esc or menu button for options"
					txtlen = rl.MeasureText(txt, 20)
					txty += 30
					txtx = int32(cnt.X) - txtlen/2
					rl.DrawText(txt, txtx, txty, 20, rl.Green)
				}
				if rl.IsKeyPressed(rl.KeySpace) || rl.IsGamepadButtonPressed(0, 7) {
					introcount = true
					startdmgT = fps * 7
					rl.PlaySound(sfx[13])
				}

				if introF3 > 0.5 {
					txt := "unklnik.com"
					txtlen := rl.MeasureText(txt, 40)
					txtx := int32(cnt.X) - txtlen/2
					txty := int32(scrH) - 50
					rl.DrawText(txt, txtx, txty, 40, rl.Green)
				}
			}

		}

	}

	//FLOOD
	if mods.flood {
		rl.DrawRectangleRec(floodRec, rl.Fade(rl.SkyBlue, rF32(0.07, 0.12)))

		//WATER ANIM
		siz := bsU2
		x := floodRec.X
		y := floodRec.Y - siz/2
		for {
			rec := rl.NewRectangle(x, y, siz, siz)
			rl.DrawTexturePro(imgs, floodImg, rec, ori, 0, rl.SkyBlue)
			x += siz
			if x >= scrWF32 {
				break
			}
		}
		if frames%3 == 0 {
			floodImg.X += floodanim.W
			if floodImg.X > floodanim.xl+floodanim.frames*floodanim.W {
				floodImg.X = floodanim.xl
			}
		}

		//FISH
		rec := rl.NewRectangle(fishV2.X-fishSiz/2, fishV2.Y-fishSiz/2, fishSiz, fishSiz)
		rl.DrawTexturePro(imgs, fish1, rec, ori, 0, rl.SkyBlue)
		rec = rl.NewRectangle(fish2V2.X-fishSiz2/2, fish2V2.Y-fishSiz2/2, fishSiz2, fishSiz2)
		rl.DrawTexturePro(imgs, fish2, rec, ori, 0, rl.SkyBlue)

		if frames%10 == 0 {
			if fishLR {
				fish1.X += fishL.W
				if fish1.X > fishL.xl+fishL.frames*fishL.W {
					fish1.X = fishL.xl
				}
			} else {
				fish1.X += fishR.W
				if fish1.X > fishR.xl+fishR.frames*fishR.W {
					fish1.X = fishR.xl
				}
			}
			if fish2LR {
				fish2.X += fishR.W
				if fish2.X > fishR.xl+fishR.frames*fishR.W {
					fish2.X = fishR.xl
				}
			} else {
				fish2.X += fishL.W
				if fish2.X > fishL.xl+fishL.frames*fishL.W {
					fish2.X = fishL.xl
				}
			}

		}

	}

	//RAIN
	if mods.umbrella {
		for a := 0; a < len(rain); a++ {
			rl.DrawRectangleRec(rain[a], rl.Fade(ranCyan(), rF32(0.4, 0.7)))
			rain[a].Y += 8
			if rain[a].Y > scrHF32 {
				rain[a].Y = rF32(-scrHF32, -bsU)
			}
		}

	}

	//TELEPORT
	if teleporton {
		for a := 0; a < len(teleportRadius); a++ {
			rl.DrawCircleLines(int32(cnt.X), int32(cnt.Y), teleportRadius[a], ranCol())
			teleportRadius[a] -= bsU
			if teleportRadius[a] <= 0 {
				teleporton = false
				pl.cnt = rl.NewVector2(levRecInner.X+levRecInner.Width/2, levRecInner.Y+bsU3)
				upPlayerRec()
				for i := 0; i < len(level[teleportRoomNum].innerBloks); i++ {
					if rl.CheckCollisionRecs(pl.crec, level[teleportRoomNum].innerBloks[i].rec) {
						pl.rec.X = level[teleportRoomNum].innerBloks[i].rec.X + level[teleportRoomNum].innerBloks[i].rec.Width + bsU/4
						upPlayerRec()
						break
					}
				}

				cntCompanion := pl.cnt
				if mods.carrot {
					mrcarrot.rec = rl.NewRectangle(cntCompanion.X-mrcarrot.rec.Width/2, cntCompanion.Y-mrcarrot.rec.Width/2, mrcarrot.rec.Width, mrcarrot.rec.Width)
				}
				if mods.alien {
					mralien.rec = rl.NewRectangle(cntCompanion.X-mralien.rec.Width/2, cntCompanion.Y-mralien.rec.Width/2, mralien.rec.Width, mralien.rec.Width)
				}
				if mods.planty {
					mrplanty.rec = rl.NewRectangle(cntCompanion.X-mrplanty.rec.Width/2, cntCompanion.Y-mrplanty.rec.Width/2, mrplanty.rec.Width, mrplanty.rec.Width)
				}

				roomNum = teleportRoomNum
				break
			}
		}
	}

	//SCANLINES
	if scanlineson {
		for a := 0; a < len(scanlinev2); a++ {
			v2 := scanlinev2[a]
			v2.X += scrWF32
			rl.DrawLineEx(scanlinev2[a], v2, 1, rl.Fade(rl.Black, 0.5))
			scanlinev2[a].Y++
			if scanlinev2[a].Y > scrHF32+2 {
				scanlinev2[a].Y = 0
			}
		}
	}

	//MARK: DRAW MAP
	if levMapOn {
		txt := "level " + fmt.Sprint(levelnum)
		txtlen := rl.MeasureText(txt, txU4)
		txtx := int32(cnt.X) - txtlen/2
		txty := txU
		rl.DrawText(txt, txtx, txty, txU4, rl.White)

		for a := 0; a < len(levMap); a++ {
			if debug {
				rl.DrawText(fmt.Sprint(a), levMap[a].ToInt32().X+4, levMap[a].ToInt32().Y+4, txU, rl.White)
			}

			if shopRoomNum == a && roomNum == a {
				rec := levMap[a]
				rec.Width = rec.Width / 2
				rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.2))
				rec.X += rec.Width
				rl.DrawRectangleRec(rec, rl.Fade(rl.Magenta, 0.2))
			} else if exitRoomNum == a && roomNum == a {
				rec := levMap[a]
				rec.Width = rec.Width / 2
				rl.DrawRectangleRec(rec, rl.Fade(rl.Green, 0.2))
				rec.X += rec.Width
				rl.DrawRectangleRec(rec, rl.Fade(rl.Yellow, 0.2))
			} else if exitRoomNum == a && mods.exitmap || exitRoomNum == a && level[a].visited {
				rl.DrawRectangleRec(levMap[a], rl.Fade(rl.Yellow, 0.2))
			} else if shopRoomNum == a {
				rl.DrawRectangleRec(levMap[a], rl.Fade(rl.Magenta, 0.2))
			} else if roomNum == a {
				rl.DrawRectangleRec(levMap[a], rl.Fade(rl.Green, 0.2))
			} else {
				if level[a].visited {
					rl.DrawRectangleRec(levMap[a], rl.Fade(rl.Blue, 0.2))
				} else {
					rl.DrawRectangleRec(levMap[a], rl.Fade(rl.Red, 0.2))
				}
			}
			rl.DrawRectangleLinesEx(levMap[a], 1, rl.Black)
		}

		txt = "player visited shop exit"
		txtlen = rl.MeasureText(txt, txU4)
		txtx = int32(cnt.X) - txtlen/2
		txty = scrH32 - txU5
		txtlen = rl.MeasureText("player ", txU4)
		rl.DrawText("player ", txtx, txty, txU4, rl.Green)
		txtx += txtlen
		txtlen = rl.MeasureText("visited ", txU4)
		rl.DrawText("visited ", txtx, txty, txU4, rl.Blue)
		txtx += txtlen
		txtlen = rl.MeasureText("shop ", txU4)
		rl.DrawText("shop ", txtx, txty, txU4, rl.Magenta)
		txtx += txtlen
		txtlen = rl.MeasureText("exit ", txU4)
		rl.DrawText("exit ", txtx, txty, txU4, rl.Yellow)

	}

}
func drawnoRender() { //MARK:DRAW NO RENDER

	if invenon {
		rl.BeginMode2D(cam2)

		drawInvenDetail()

		rl.EndMode2D()

	}

	if debug {
		drawDebug()
	}

}
func drawDebug() { //MARK:DRAW DEBUG

	siderec := rl.NewRectangle(0, 0, 300, scrHF32)
	rl.DrawRectangleRec(siderec, rl.Fade(darkRed(), 0.3))

	txtX, txtY := txU, txU

	rl.DrawText("pl.cnt.X"+" "+fmt.Sprint(pl.cnt.X), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("pl.cnt.Y"+" "+fmt.Sprint(pl.cnt.Y), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("levBorderBlokNum"+" "+fmt.Sprint(levBorderBlokNum), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len FX"+" "+fmt.Sprint(len(fx)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("axeT2"+" "+fmt.Sprint(mods.axeT2), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("santaT"+" "+fmt.Sprint(mods.santaT), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("snowOn"+" "+fmt.Sprint(mods.snowon), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len(enProj)"+" "+fmt.Sprint(len(enProj)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len(plProj)"+" "+fmt.Sprint(len(plProj)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.fireballN"+" "+fmt.Sprint(mods.fireballN), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.bounceN"+" "+fmt.Sprint(mods.bounceN), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("pl.atkDMG"+" "+fmt.Sprint(pl.atkDMG), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.orbitalN"+" "+fmt.Sprint(mods.orbitalN), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.alien"+" "+fmt.Sprint(mods.alien), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("mods.carrot"+" "+fmt.Sprint(mods.carrot), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("cam2.Zoom"+" "+fmt.Sprint(cam2.Zoom), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("isController"+" "+fmt.Sprint(isController), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("useController"+" "+fmt.Sprint(useController), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("GamepadAxisMovement 0"+" "+fmt.Sprint(rl.GetGamepadAxisMovement(0, 0)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("GamepadAxisMovement 1"+" "+fmt.Sprint(rl.GetGamepadAxisMovement(0, 1)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("bosses[bossnum].timer"+" "+fmt.Sprint(bosses[bossnum].timer), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("shopitems[0].shopoff"+" "+fmt.Sprint(shopitems[0].shopoff), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("floodRec.Y"+" "+fmt.Sprint(floodRec.Y), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("pl.crec.Y"+" "+fmt.Sprint(pl.crec.Y), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("minsEND"+" "+fmt.Sprint(minsEND), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("secsEND"+" "+fmt.Sprint(secsEND), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("len(mariocoins)"+" "+fmt.Sprint(len(mariocoins)), txtX, txtY, txU, rl.White)
	txtY += txU
	rl.DrawText("hardcore"+" "+fmt.Sprint(hardcore), txtX, txtY, txU, rl.White)
	txtY += txU

}
func drawBlokDrec(blok xblok, shadow, blur bool, blurDist float32) { //MARK:DRAW BLOK DREC

	if shadow {
		shadowRec := blok.drec
		shadowRec.X -= 5
		shadowRec.Y += 5
		rl.DrawTexturePro(imgs, blok.img, shadowRec, rl.NewVector2(blok.drec.Width/2, blok.drec.Height/2), blok.ro, rl.Fade(rl.Black, 0.8))
	}

	rl.DrawTexturePro(imgs, blok.img, blok.drec, rl.NewVector2(blok.drec.Width/2, blok.drec.Height/2), blok.ro, rl.Fade(blok.color, blok.fade))
	if blur {
		blurRec := blok.drec
		blurRec.X -= rF32(-blurDist, blurDist)
		blurRec.Y -= rF32(-blurDist, blurDist)
		rl.DrawTexturePro(imgs, blok.img, blurRec, rl.NewVector2(blok.drec.Width/2, blok.drec.Height/2), blok.ro, rl.Fade(blok.color, rF32(0.05, 0.2)))
	}

	if debug {
		rl.DrawRectangleLinesEx(blok.rec, 0.5, rl.Green)
		rl.DrawRectangleLinesEx(blok.crec, 1, rl.Blue)
		rl.DrawRectangleLinesEx(blok.crec2, 2, rl.Red)
		rl.DrawCircleV(blok.cnt, 4, rl.Red)
		rl.DrawText(blok.name, blok.rec.ToInt32().X, blok.rec.ToInt32().Y-10, 10, rl.White)
		rl.DrawText(fmt.Sprint(blok.timer), blok.rec.ToInt32().X, blok.rec.ToInt32().Y+blok.rec.ToInt32().Height, 10, rl.White)
	}

}
func drawBlok(blok xblok, shadow, blur bool, blurDist float32) { //MARK:DRAW BLOK

	if shadow {
		shadowRec := blok.rec
		shadowRec.X -= 5
		shadowRec.Y += 5
		rl.DrawTexturePro(imgs, blok.img, shadowRec, rl.NewVector2(0, 0), blok.ro, rl.Fade(rl.Black, 0.8))
	}

	//CANDLE LIGHT
	if blok.name == "candle" {
		radius := rF32(bsU3, bsU5)
		rl.DrawCircleGradient(int32(blok.cnt.X), int32(blok.cnt.Y), radius, rl.Fade(ranYellow(), rF32(0.05, 0.3)), rl.Blank)
	}

	//DRAW BLOK IMG
	rl.DrawTexturePro(imgs, blok.img, blok.rec, rl.NewVector2(0, 0), blok.ro, rl.Fade(blok.color, blok.fade))
	if blur {
		if blok.name == "skull" || blok.name == "candle" { //HALF BLUR
			blurRec := blok.rec
			blurRec.X -= rF32(-blurDist/2, blurDist/2)
			blurRec.Y -= rF32(-blurDist/2, blurDist/2)
			rl.DrawTexturePro(imgs, blok.img, blurRec, rl.NewVector2(0, 0), blok.ro, rl.Fade(blok.color, rF32(0.05, 0.2)))
		} else { //FULL BLUR
			blurRec := blok.rec
			blurRec.X -= rF32(-blurDist, blurDist)
			blurRec.Y -= rF32(-blurDist, blurDist)
			rl.DrawTexturePro(imgs, blok.img, blurRec, rl.NewVector2(0, 0), blok.ro, rl.Fade(blok.color, rF32(0.05, 0.2)))
		}
	}

	if debug {
		rl.DrawRectangleLinesEx(blok.rec, 0.5, rl.Green)
		rl.DrawRectangleLinesEx(blok.crec, 2, rl.Blue)
		rl.DrawRectangleLinesEx(blok.crec2, 2, rl.Red)
		rl.DrawText(blok.name, blok.rec.ToInt32().X, blok.rec.ToInt32().Y-10, 10, rl.White)
	}

}
func drawUpEnProj() { //MARK:DRAW UP ENEMY PROJECTILES

	clear := false
	for a := 0; a < len(enProj); a++ {

		if enProj[a].onoff {
			//	rl.DrawRectangleRec(enProj[a].rec, enProj[a].col)

			if enProj[a].name == "mushbull" {
				shadowrec := makeDrec(enProj[a].rec)
				shadowrec.X -= 5
				shadowrec.Y += 5
				rl.DrawTexturePro(imgs, enProj[a].img, shadowrec, origin(enProj[a].rec), enProj[a].ro, rl.Fade(rl.Black, 0.7))
				if flipcoin() {
					rl.DrawTexturePro(imgs, enProj[a].img, makeDrec(enProj[a].rec), origin(enProj[a].rec), enProj[a].ro, ranCyan())
				} else {
					rl.DrawTexturePro(imgs, enProj[a].img, makeDrec(enProj[a].rec), origin(enProj[a].rec), enProj[a].ro, ranRed())
				}
			} else if enProj[a].name == "boss3" {
				rl.DrawTexturePro(imgs, enProj[a].img, makeDrec(enProj[a].rec), enProj[a].ori, enProj[a].ro, enProj[a].col)

			} else if enProj[a].name == "ninja" {

				rl.DrawTexturePro(imgs, enProj[a].img, makeDrec(enProj[a].rec), rl.NewVector2(enProj[a].rec.Width/2, enProj[a].rec.Height/2), enProj[a].ro, enProj[a].col)

			} else {
				rl.DrawTexturePro(imgs, enProj[a].img, enProj[a].rec, ori, enProj[a].ro, enProj[a].col)
				if enProj[a].name == "boss2" {
					rl.DrawTexturePro(imgs, enProj[a].img, BlurRec(enProj[a].rec, 7), ori, enProj[a].ro, rl.Fade(enProj[a].col, 0.5))
				}
			}

			switch enProj[a].name {
			case "boss3":
				enProj[a].ro += 12
				if frames%4 == 0 {
					if enProj[a].rec.Width < bsU8 {
						enProj[a].rec.X -= 4
						enProj[a].rec.Y -= 4
						enProj[a].rec.Width += 8
						enProj[a].rec.Height += 8
						enProj[a].ori = rl.NewVector2(enProj[a].rec.Width/2, enProj[a].rec.Height/2)
					}
				}
				if roll18() == 18 {
					enProj[a].velx, enProj[a].vely = moveFollow(enProj[a].cnt, pl.cnt, enProj[a].vel)
				}
			case "boss2":
				if frames%4 == 0 {
					enProj[a].img.X += enProj[a].img.Width
					if enProj[a].img.X > boss2anim.xl+(float32(boss2anim.frames)*enProj[a].img.Width) {
						enProj[a].img.X = boss2anim.xl
					}
				}
			case "boss1":
				if roll12() == 12 {
					enProj[a].velx, enProj[a].vely = moveFollow(enProj[a].cnt, pl.cnt, enProj[a].vel)
				}

				if frames%3 == 0 {
					enProj[a].img.X += enProj[a].img.Width
					if enProj[a].img.X > boss1anim.xl+(float32(boss1anim.frames)*enProj[a].img.Width) {
						enProj[a].img.X = boss1anim.xl
					}
				}
			case "mushbull":
				if frames%2 == 0 {
					enProj[a].img.X += enProj[a].img.Width
					if enProj[a].img.X > mushBull.xl+(float32(mushBull.frames)*enProj[a].img.Width) {
						enProj[a].img.X = mushBull.xl
					}
				}

				if roll18() == 18 {
					enProj[a].velx, enProj[a].vely = moveFollow(enProj[a].cnt, pl.cnt, enProj[a].vel)
				}
			case "ninja":
				enProj[a].ro += bsU / 2
			}

			if debug {
				rl.DrawRectangleLinesEx(enProj[a].rec, 0.5, rl.Red)
				rl.DrawRectangleLinesEx(enProj[a].crec, 0.5, rl.White)
			}

			if checkNextMove(enProj[a].rec, enProj[a].velx, enProj[a].vely, true) {
				enProj[a].cnt.X += enProj[a].velx
				enProj[a].cnt.Y += enProj[a].vely
				enProj[a].rec = rl.NewRectangle(enProj[a].cnt.X-enProj[a].rec.Width/2, enProj[a].cnt.Y-enProj[a].rec.Height/2, enProj[a].rec.Width, enProj[a].rec.Height)
				switch enProj[a].name {
				case "boss1":
					enProj[a].crec = enProj[a].rec
					enProj[a].crec.X += enProj[a].rec.Width / 4
					enProj[a].crec.Y += enProj[a].rec.Height / 4
					enProj[a].crec.Width = enProj[a].rec.Width / 2
					enProj[a].crec.Height = enProj[a].rec.Height / 2
				}
			} else {
				enProj[a].onoff = false
			}

			if !rl.CheckCollisionPointRec(enProj[a].cnt, levRecInner) {
				enProj[a].onoff = false
			}

			if enProj[a].name == "boss1" {
				if rl.CheckCollisionRecs(enProj[a].crec, pl.crec) {
					hitPL(a, 1)
				}
			} else {
				if rl.CheckCollisionRecs(enProj[a].rec, pl.crec) {
					hitPL(a, 1)
				}
			}

		} else {
			clear = true
		}
	}

	if clear {
		for a := 0; a < len(enProj); a++ {
			if !enProj[a].onoff {
				enProj = remProj(enProj, a)
			}
		}
	}

}
func drawUpAirStrike() { //MARK:DRAW UP AIR STRIKE

	for a := 0; a < len(airstrikeV2); a++ {
		siz := bsU2
		rec := rl.NewRectangle(airstrikeV2[a].X-siz/2, airstrikeV2[a].Y-siz/2, siz, siz)
		switch airstrikeDir {
		case 1:
			rl.DrawTexturePro(imgs, etc[48], makeDrec(rec), origin(rec), 180, rl.White)
			airstrikeV2[a].Y += 7
			if a == 1 {
				if airstrikeV2[a].Y > levRec.Y+levRec.Width {
					airstrikeOn = false
				}
			}
		case 2:
			rl.DrawTexturePro(imgs, etc[48], makeDrec(rec), origin(rec), 270, rl.White)
			airstrikeV2[a].X -= 7
			if a == 1 {
				if airstrikeV2[a].X < levRec.X-rec.Width {
					airstrikeOn = false
				}
			}
		case 3:
			rl.DrawTexturePro(imgs, etc[48], makeDrec(rec), origin(rec), 0, rl.White)
			airstrikeV2[a].Y -= 7
			if a == 1 {
				if airstrikeV2[a].Y < levRec.Y-rec.Width {
					airstrikeOn = false
				}
			}
		case 4:
			rl.DrawTexturePro(imgs, etc[48], makeDrec(rec), origin(rec), 90, rl.White)
			airstrikeV2[a].X += 7
			if a == 1 {
				if airstrikeV2[a].X > levRec.X+levRec.Width {
					airstrikeOn = false
				}
			}
		}
	}

	airstrikebombT--
	if airstrikebombT <= 0 && rl.CheckCollisionPointRec(airstrikeV2[0], levRecInner) {

		siz := bsU8

		switch airstrikeDir {
		case 1:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = airstrikeV2[0]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[1]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[2]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
		case 2:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = airstrikeV2[0]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[1]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[2]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
		case 3:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = airstrikeV2[0]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[1]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[2]
			zblok.cnt.X += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
		case 4:
			zblok := makeBlokGenNoRecNoCntr()
			zblok.cnt = airstrikeV2[0]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.name = "airbomb"
			zblok.img = airstrikeanim.recTL
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[1]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			zblok.cnt = airstrikeV2[2]
			zblok.cnt.Y += rF32(-bsU4, bsU4)
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
			zblok.crec = zblok.rec
			zblok.crec.Width = zblok.crec.Width / 2
			zblok.crec.Height = zblok.crec.Width
			zblok.crec.X += zblok.crec.Width / 2
			zblok.crec.Y += zblok.crec.Width / 2
			level[roomNum].etc = append(level[roomNum].etc, zblok)
		}

		airstrikebombT = rI32(int(fps/4), int(fps*2))
	}

}
func drawUpPlayerProj() { //MARK:DRAW UP PLAYER PROJECTILES

	//MOVE DRAW
	clear := false
	for a := 0; a < len(plProj); a++ {
		if plProj[a].onoff {
			shadowRec := plProj[a].drec
			shadowRec.X -= 5
			shadowRec.Y += 5
			rl.DrawTexturePro(imgs, plProj[a].img, shadowRec, plProj[a].ori, plProj[a].ro, rl.Fade(rl.Black, 0.8))
			rl.DrawTexturePro(imgs, plProj[a].img, plProj[a].drec, plProj[a].ori, plProj[a].ro, rl.Fade(plProj[a].col, plProj[a].fade))

			if debug {
				rl.DrawRectangleLinesEx(plProj[a].rec, 0.5, rl.Magenta)
			}

			plProj[a].cnt.X += plProj[a].velx
			plProj[a].cnt.Y += plProj[a].vely
			plProj[a].rec = rl.NewRectangle(plProj[a].cnt.X-plProj[a].rec.Width/2, plProj[a].cnt.Y-plProj[a].rec.Height/2, plProj[a].rec.Width, plProj[a].rec.Height)
			plProj[a].drec = plProj[a].rec
			plProj[a].drec.X += plProj[a].rec.Width / 2
			plProj[a].drec.Y += plProj[a].rec.Height / 2

			switch plProj[a].name {
			case "mrcarrot":
				plProj[a].ro += 8
			case "plantbull":
				if frames%4 == 0 {
					plProj[a].img.X += plantBull.recTL.Width
					if plProj[a].img.X >= plantBull.xl+(plantBull.recTL.Width*plantBull.frames) {
						plProj[a].img.X = plantBull.xl
					}
				}
			case "fireball":
				if frames%3 == 0 {
					plProj[a].img.X += 17
					if plProj[a].img.X >= fireballPlayer.xl+(17*fireballPlayer.frames) {
						plProj[a].img.X = fireballPlayer.xl
					}
				}
			}

			//COLLISION ETC
			if len(level[roomNum].etc) > 0 {
				for b := 0; b < len(level[roomNum].etc); b++ {
					if level[roomNum].etc[b].onoff {
						if rl.CheckCollisionRecs(level[roomNum].etc[b].rec, plProj[a].rec) {
							switch level[roomNum].etc[b].name {
							case "powerupBlok":
								destroyPowerupBlok(b)
								makeFX(3, level[roomNum].etc[b].cnt)
								level[roomNum].etc[b].onoff = false
								rl.PlaySound(sfx[6])
							case "oilbarrel":
								makeFX(4, level[roomNum].etc[b].cnt)
								level[roomNum].etc[b].onoff = false
								rl.PlaySound(sfx[10])
								rl.PlaySound(sfx[5])
							}
						}
					}
				}
			}

			//COLLISION BOUNDARY
			if !rl.CheckCollisionPointRec(plProj[a].cnt, levRecInner) {

				if plProj[a].bounceN > 0 {
					plProj[a].bounceN--
					plProj[a].velx *= -1
					plProj[a].vely *= -1
					if plProj[a].name == "fireball" || plProj[a].name == "fireworks" {
						plProj[a].ro += 180
					}
				} else {
					plProj[a].onoff = false
				}
			}
			//COLLISION INNER RECS
			if plProj[a].onoff {
				for b := 0; b < len(level[roomNum].innerBloks); b++ {
					if rl.CheckCollisionRecs(plProj[a].rec, level[roomNum].innerBloks[b].rec) {
						if plProj[a].bounceN > 0 {
							plProj[a].bounceN--
							plProj[a].velx *= -1
							plProj[a].vely *= -1
							if plProj[a].name == "fireball" || plProj[a].name == "fireworks" {
								plProj[a].ro += 180
							}
						} else {
							plProj[a].onoff = false
						}
					}
				}
			}

			switch plProj[a].name {
			case "axe":
				plProj[a].ro += 8
			}

		} else {
			clear = true
		}

	}

	if clear {
		for a := 0; a < len(plProj); a++ {
			if !plProj[a].onoff {
				plProj = remProj(plProj, a)
			}
		}
	}

	//CHECK ENEMY PROJ COLLIS
	if len(level[roomNum].enemies) > 0 {
		for a := 0; a < len(plProj); a++ {
			if plProj[a].onoff {
				for b := 0; b < len(level[roomNum].enemies); b++ {
					if rl.CheckCollisionRecs(plProj[a].rec, level[roomNum].enemies[b].rec) && level[roomNum].enemies[b].hppause == 0 {
						level[roomNum].enemies[b].hppause = fps / 2
						level[roomNum].enemies[b].hp -= plProj[a].dmg
						playenemyhit()
						if level[roomNum].enemies[b].hp <= 0 {
							cntr := level[roomNum].enemies[b].cnt
							level[roomNum].enemies[b].off = true
							addkill(b)
							makeFX(2, cntr)
						}

					}
				}

			}
		}
	}
	//CHECK BOSS PROJ COLLIS
	if levelnum == 6 {
		for a := 0; a < len(plProj); a++ {
			if plProj[a].onoff {

				if rl.CheckCollisionRecs(plProj[a].rec, bosses[bossnum].crec) && bosses[bossnum].hppause == 0 {
					bosses[bossnum].hppause = fps
					bosses[bossnum].hp -= plProj[a].dmg
					if bosses[bossnum].hp <= 0 {
						cntr := bosses[bossnum].cnt
						bosses[bossnum].off = true
						makeFX(2, cntr)
						if !endgame {
							pause = true
							minsEND = mins
							secsEND = secs
							addtime()
							endgopherrec = rl.NewRectangle(cnt.X-bsU4, levRec.Y+levRec.Height, bsU8, bsU8)
							endgameT = fps * 3
							endgame = true
						}
					}
					rl.PlaySound(sfx[11])
				}

			}
		}

	}

}
func drawUpEtc() { //MARK:DRAW UP ETC

	clear := false
	for a := 0; a < len(level[roomNum].etc); a++ {

		if level[roomNum].etc[a].onoff {

			//TIMERS
			if level[roomNum].etc[a].timer > 0 {
				level[roomNum].etc[a].timer--
			}
			if level[roomNum].etc[a].txtT > 0 {
				level[roomNum].etc[a].txtT--
			}
			//SHOP ARROWS
			if level[roomNum].etc[a].name == "shop" {
				rl.DrawTriangle(level[roomNum].etc[a].v2s[4], level[roomNum].etc[a].v2s[3], level[roomNum].etc[a].v2s[5], rl.Fade(rl.SkyBlue, 0.8))

				if frames%3 == 0 {
					level[roomNum].etc[a].v2s[3].Y--
					level[roomNum].etc[a].v2s[4].Y--
					level[roomNum].etc[a].v2s[5].Y--
				}

				if level[roomNum].etc[a].v2s[4].Y < level[roomNum].etc[a].rec.Y+level[roomNum].etc[a].rec.Height {

					level[roomNum].etc[a].v2s[3] = level[roomNum].etc[a].v2s[0]
					level[roomNum].etc[a].v2s[4] = level[roomNum].etc[a].v2s[1]
					level[roomNum].etc[a].v2s[5] = level[roomNum].etc[a].v2s[2]
				}
			}
			//DRAW BLOK
			if level[roomNum].etc[a].name == "turret" {
				drawBlokDrec(level[roomNum].etc[a], true, false, 0)

				level[roomNum].etc[a].ro++
				if level[roomNum].etc[a].timer <= 0 {
					level[roomNum].etc[a].timer = rI32(1, 3) * fps
					makeProjectileEnemy(1, level[roomNum].etc[a].cnt)
				}
			} else if level[roomNum].etc[a].name == "spring" {

				drawBlokDrec(level[roomNum].etc[a], false, true, 4)

				if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].crec) && level[roomNum].etc[a].timer == 0 {
					level[roomNum].etc[a].timer = fps * 3
					level[roomNum].etc[a].onoffswitch = true

					pl.slide = true
					pl.slideT = fps / 2
					pl.slideDIR = level[roomNum].etc[a].slideDIR

					rl.PlaySound(sfx[3])
				}
				if level[roomNum].etc[a].onoffswitch {
					if frames%4 == 0 {
						level[roomNum].etc[a].img.X += spring.W
						if level[roomNum].etc[a].img.X > spring.xl+spring.frames*spring.W {
							level[roomNum].etc[a].img.X = spring.xl
							level[roomNum].etc[a].onoffswitch = false
						}
					}
				} else {
					if level[roomNum].etc[a].img.X > spring.xl {
						if frames%4 == 0 {
							level[roomNum].etc[a].img.X -= spring.W
							if level[roomNum].etc[a].img.X < spring.xl {
								level[roomNum].etc[a].img.X = spring.xl
							}
						}
					}
				}

			} else if level[roomNum].etc[a].name == "gascloudtrap" {
				drawBlokDrec(level[roomNum].etc[a], false, true, 4)
				level[roomNum].etc[a].ro++

				if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].rec) && level[roomNum].etc[a].timer == 0 {
					level[roomNum].etc[a].timer = fps * 3
					makegascloud(level[roomNum].etc[a].cnt)
					rl.PlaySound(sfx[20])
				}

			} else if level[roomNum].etc[a].name == "gascloud" {

				rl.DrawTexturePro(imgs, etc[58], level[roomNum].etc[a].drec, level[roomNum].etc[a].ori, level[roomNum].etc[a].ro, rl.Fade(ranGreen(), rF32(0.3, 0.8)))

				if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].crec) && !pl.poison && pl.poisonCollisT == 0 {
					pl.poisonCollisT = fps * 3
					if mods.apple {
						mods.appleN--
						if mods.appleN == 0 {
							mods.apple = false
							clearinven("apple")
						}
					} else {
						pl.poison = true
						pl.poisonCount = 2
						pl.poisonT = fps * 3
						txtHere("poisoned", pl.rec)
						rl.PlaySound(sfx[23])
					}
				}

				if frames%3 == 0 {
					level[roomNum].etc[a].img.X += posiongas.W
					if level[roomNum].etc[a].img.X > posiongas.xl+posiongas.frames*posiongas.W {
						level[roomNum].etc[a].img.X = posiongas.xl
					}
				}

				rl.DrawTexturePro(imgs, etc[58], BlurRec(level[roomNum].etc[a].drec, 5), level[roomNum].etc[a].ori, level[roomNum].etc[a].ro, rl.Fade(ranGreen(), rF32(0.1, 0.3)))
				rl.DrawTexturePro(imgs, etc[58], BlurRec(level[roomNum].etc[a].drec, 7), level[roomNum].etc[a].ori, level[roomNum].etc[a].ro, rl.Fade(ranGreen(), rF32(0.1, 0.3)))

				if debug {
					rl.DrawRectangleLinesEx(level[roomNum].etc[a].crec, 0.5, rl.White)
				}

				level[roomNum].etc[a].ro += 5

				if checknextmoveV2innerRec(level[roomNum].etc[a].cnt, level[roomNum].etc[a].velX, level[roomNum].etc[a].velY) {
					level[roomNum].etc[a].rec.X += level[roomNum].etc[a].velX
					level[roomNum].etc[a].rec.Y += level[roomNum].etc[a].velY
					level[roomNum].etc[a].drec.X += level[roomNum].etc[a].velX
					level[roomNum].etc[a].drec.Y += level[roomNum].etc[a].velY
					level[roomNum].etc[a].crec.X += level[roomNum].etc[a].velX
					level[roomNum].etc[a].crec.Y += level[roomNum].etc[a].velY
					level[roomNum].etc[a].cnt.X += level[roomNum].etc[a].velX
					level[roomNum].etc[a].cnt.Y += level[roomNum].etc[a].velY
				} else {
					level[roomNum].etc[a].velX = rF32(-level[roomNum].etc[a].vel, level[roomNum].etc[a].vel)
					level[roomNum].etc[a].velY = rF32(-level[roomNum].etc[a].vel, level[roomNum].etc[a].vel)
				}

			} else if level[roomNum].etc[a].name == "slimetrail" || level[roomNum].etc[a].name == "playerblood" {

				drawBlok(level[roomNum].etc[a], false, true, 7)

				level[roomNum].etc[a].fade -= 0.005
				if level[roomNum].etc[a].fade <= 0 {
					level[roomNum].etc[a].onoff = false
				}

				if level[roomNum].etc[a].name == "slimetrail" {
					if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].crec) && !pl.poison && pl.poisonCollisT == 0 {
						pl.poisonCollisT = fps * 3
						if mods.apple {
							mods.appleN--
							if mods.appleN == 0 {
								mods.apple = false
								clearinven("apple")
							}
						} else {
							pl.poison = true
							pl.poisonCount = 2
							pl.poisonT = fps * 3
							txtHere("poisoned", pl.rec)
						}
					}
				}

			} else if level[roomNum].etc[a].name == "footprints" {

				drawBlokDrec(level[roomNum].etc[a], false, true, 4)

				level[roomNum].etc[a].fade -= 0.01
				if level[roomNum].etc[a].fade <= 0 {
					level[roomNum].etc[a].onoff = false
				}

			} else if level[roomNum].etc[a].name == "blades" {
				drawBlok(level[roomNum].etc[a], true, true, 4)
				if frames%2 == 0 {
					level[roomNum].etc[a].img.X += blades.W
					if level[roomNum].etc[a].img.X > blades.xl+blades.frames*blades.W {
						level[roomNum].etc[a].img.X = blades.xl
						level[roomNum].etc[a].onoffswitch = false
					}
				}
				if checkNextMove(level[roomNum].etc[a].rec, level[roomNum].etc[a].velX, level[roomNum].etc[a].velY, true) {
					level[roomNum].etc[a].rec.X += level[roomNum].etc[a].velX
					level[roomNum].etc[a].rec.Y += level[roomNum].etc[a].velY
					level[roomNum].etc[a].cnt = rl.NewVector2(level[roomNum].etc[a].rec.X+level[roomNum].etc[a].rec.Width/2, level[roomNum].etc[a].rec.Y+level[roomNum].etc[a].rec.Width/2)
				} else {
					level[roomNum].etc[a].velX = rF32(-level[roomNum].etc[a].velX, level[roomNum].etc[a].velX)
					level[roomNum].etc[a].velY = rF32(-level[roomNum].etc[a].velY, level[roomNum].etc[a].velY)
				}

				if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].rec) {
					if !rl.IsSoundPlaying(sfx[30]) {
						rl.PlaySound(sfx[30])
					}
					hitPL(0, 2)
				}

			} else if level[roomNum].etc[a].name == "spear" {
				drawBlokDrec(level[roomNum].etc[a], false, true, 4)

				if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].crec) {
					hitPL(0, 2)
					if frames%4 == 0 {
						if level[roomNum].etc[a].img.X == spear.xl {
							rl.PlaySound(sfx[26])
						}
						level[roomNum].etc[a].img.X += spear.W
						if level[roomNum].etc[a].img.X > spear.xl+spear.frames*spear.W {
							level[roomNum].etc[a].img.X = spear.xl
						}
					}

				} else {
					if level[roomNum].etc[a].img.X != spear.xl {
						if frames%4 == 0 {
							level[roomNum].etc[a].img.X -= spear.W
						}
					}
				}

			} else if level[roomNum].etc[a].name == "powerupBlok" {
				drawBlok(level[roomNum].etc[a], true, false, 4)
			} else if level[roomNum].etc[a].name == "flamefiretrail" {
				drawBlok(level[roomNum].etc[a], false, true, 4)
				if frames%3 == 0 {
					level[roomNum].etc[a].img.X += firetrailanim.W
					if level[roomNum].etc[a].img.X > firetrailanim.xl+firetrailanim.frames*firetrailanim.W {
						level[roomNum].etc[a].img.X = firetrailanim.xl
					}
				}

				level[roomNum].etc[a].fade -= 0.005
				if level[roomNum].etc[a].fade <= 0 {
					level[roomNum].etc[a].onoff = false
				}

			} else if level[roomNum].etc[a].name == "airbomb" {
				drawBlok(level[roomNum].etc[a], false, true, 4)
				if frames%5 == 0 {
					if level[roomNum].etc[a].img.X == airstrikeanim.xl {
						rl.PlaySound(sfx[28])
					}
					level[roomNum].etc[a].img.X += airstrikeanim.W
					if level[roomNum].etc[a].img.X > airstrikeanim.xl+airstrikeanim.frames*airstrikeanim.W {
						level[roomNum].etc[a].img.X = airstrikeanim.xl
					}
				}

				level[roomNum].etc[a].fade -= 0.03
				if level[roomNum].etc[a].fade <= 0 {
					level[roomNum].etc[a].onoff = false
				}

			} else { //DRAW OTHER BLOKS

				drawBlok(level[roomNum].etc[a], true, true, 4)

				if level[roomNum].etc[a].name == "exit" {
					txt := "exit"
					txtlen := rl.MeasureText(txt, 20)
					txtx := int32(level[roomNum].etc[a].rec.X+level[roomNum].etc[a].rec.Width/2) - txtlen/2
					txtx++
					txty := int32(level[roomNum].etc[a].rec.Y) - 18
					rl.DrawText(txt, txtx-4, txty+4, 20, rl.Black)
					rl.DrawText(txt, txtx, txty, 20, ranCol())

					if rl.CheckCollisionRecs(level[roomNum].etc[a].rec, pl.crec) {
						rl.PlaySound(sfx[15])
						exited = true
					}
				}

				if level[roomNum].etc[a].name == "chest" && level[roomNum].etc[a].img.X < 493 {

					if rl.CheckCollisionRecs(pl.arec, level[roomNum].etc[a].crec) && mods.key && !level[roomNum].etc[a].onoffswitch {
						rl.PlaySound(sfx[17])
						level[roomNum].etc[a].onoffswitch = true
						mods.keyN--
						if mods.keyN == 0 {
							mods.key = false
						}
						delInven("key")
						makechestitem(a)
					} else if rl.CheckCollisionRecs(pl.arec, level[roomNum].etc[a].crec) && !mods.key && !level[roomNum].etc[a].onoffswitch && level[roomNum].etc[a].txtT == 0 {
						txtHere("locked", level[roomNum].etc[a].rec)
						level[roomNum].etc[a].txtT = fps * 3
						rl.PlaySound(sfx[25])
					}

					if level[roomNum].etc[a].onoffswitch {
						if frames%6 == 0 {
							level[roomNum].etc[a].img.X += 24
						}
					}

				}
			}
			//COLLISIONS PLAYER CREC > ETC REC
			if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].rec) {
				switch level[roomNum].etc[a].name {

				case "gem":
					if level[roomNum].etc[a].onoff {
						pl.coins += level[roomNum].etc[a].numCoins
						level[roomNum].etc[a].onoff = false
						txtAddCoinMulti()
					}

				case "switch": //MARK: SWITCHES ON/OFF
					if level[roomNum].etc[a].timer == 0 {
						rl.PlaySound(sfx[12])
						level[roomNum].etc[a].timer = fps * 2
						level[roomNum].etc[a].onoffswitch = !level[roomNum].etc[a].onoffswitch
						if level[roomNum].etc[a].img == etc[21] {
							level[roomNum].etc[a].img = etc[22]
						} else {
							level[roomNum].etc[a].img = etc[21]
						}

						switch level[roomNum].etc[a].numType {
						case 1:
							makeSwitchArrows()
						case 2:
							night = !night
						case 3:
							flipcam = !flipcam
							cams()
						case 4:
							shader2on = !shader2on
							if shader3on {
								shader3on = false
							}
						case 5:
							shader3on = !shader3on
							if shader2on {
								shader2on = false
							}
						case 6:
							if level[roomNum].etc[a].numCoins > 0 {
								level[roomNum].etc[a].numCoins--
								txtAddCoin()
							}

						}
					}

				//MARK:INGAME COLLECT INVEN
				case "health potion", "throwing axe", "recharge hp", "medi kit", "wallet", "mr planty", "apple", "key", "vine", "bounce", "santa", "fireball", "map", "firetrail", "invisible", "coffee", "teleport", "attack range", "attack damage", "orbital", "chain lightning", "health ring", "armor", "recharge", "anchor", "umbrella", "moldy socks", "cherry", "fish", "birthday cake", "peace", "mr alien", "air strike", "fireworks", "mr carrot", "mario":
					collectInven(a)
					level[roomNum].etc[a].onoff = false
				}
			}
			//COLLISIONS PLAYER CREC > ETC CREC2
			if rl.CheckCollisionRecs(pl.crec, level[roomNum].etc[a].crec2) && level[roomNum].etc[a].name == "shop" && !pl.escape && shopExitT == 0 {
				rl.PlaySound(sfx[16])
				shopExitY = level[roomNum].etc[a].rec.Y + level[roomNum].etc[a].rec.Height + bsU2
				shopon = true
				pause = true
			}

			//COLLISIONS PLAYER ATTACK AREA
			if rl.CheckCollisionRecs(pl.atkrec, level[roomNum].etc[a].rec) {

				switch level[roomNum].etc[a].name {
				case "oilbarrel":
					if pl.atk {
						makeFX(4, level[roomNum].etc[a].cnt)
						level[roomNum].etc[a].onoff = false
						rl.PlaySound(sfx[10])
						rl.PlaySound(sfx[5])
					}
				case "powerupBlok":
					if pl.atk {
						if mods.fireworks {
							fireworksCnt = level[roomNum].etc[a].cnt
							makeProjectile("fireworks")
						}
						destroyPowerupBlok(a)
						level[roomNum].etc[a].onoff = false
						rl.PlaySound(sfx[6])
					}
				}
			}
		} else {
			clear = true
		}
	}

	if clear {
		for a := 0; a < len(level[roomNum].etc); a++ {
			if !level[roomNum].etc[a].onoff {
				level[roomNum].etc = remBlok(level[roomNum].etc, a)
			}
		}

	}

	if exited {
		pause = true
		rl.DrawRectangle(0, 0, scrW32, scrH32, rl.Black)
		makenewlevel()
		exited = false
	}

}
func drawupfx() { //MARK:DRAW UP FX

	clear := false
	for a := 0; a < len(fx); a++ {

		switch fx[a].name {

		case "fxBurnWoodBarrel", "fxBurnOilBarrel":
			col := ranOrange()
			rl.DrawTexturePro(imgs, fx[a].img, fx[a].rec, ori, 0, col)
			rl.DrawTexturePro(imgs, fx[a].img, BlurRec(fx[a].rec, 4), ori, 0, rl.Fade(col, rF32(0.2, 0.5)))
			if frames%3 == 0 {
				fx[a].img.X += 17
				if fx[a].img.X >= burn.xl+(17*burn.frames) {
					fx[a].img.X = burn.xl
				}
			}

			if rl.CheckCollisionRecs(pl.crec, fx[a].rec) {
				hitPL(0, 2)
			}

			if debug {
				rl.DrawRectangleLinesEx(fx[a].rec, 0.5, rl.Blue)
			}

		case "fxEnemy":

			for b := 0; b < len(fx[a].recs); b++ {
				v2 := rl.NewVector2(fx[a].recs[b].rec.X+fx[a].recs[b].rec.Width/2, fx[a].recs[b].rec.Y+fx[a].recs[b].rec.Height/2)
				rl.DrawCircleV(v2, fx[a].recs[b].rec.Width/2, rl.Fade(fx[a].recs[b].col, fx[a].recs[b].fade))
				fx[a].recs[b].rec.X += fx[a].recs[b].velX
				fx[a].recs[b].rec.Y += fx[a].recs[b].velY
				fx[a].recs[b].fade -= 0.05
				fx[a].recs[b].rec.Width += 0.1
				fx[a].recs[b].rec.Height += 0.1
			}

		case "fxBarrel":

			for b := 0; b < len(fx[a].recs); b++ {
				rl.DrawRectangleRec(fx[a].recs[b].rec, rl.Fade(fx[a].recs[b].col, fx[a].recs[b].fade))
				fx[a].recs[b].rec.X += fx[a].recs[b].velX
				fx[a].recs[b].rec.Y += fx[a].recs[b].velY
				fx[a].recs[b].fade -= 0.05
				fx[a].recs[b].rec.Width -= 0.05
				fx[a].recs[b].rec.Height -= 0.05
			}

		}

		if fx[a].onoff {

			fx[a].timer--
			if fx[a].timer == 0 {
				fx[a].onoff = false
			}

		} else {
			clear = true
		}

	}

	if clear {
		for a := 0; a < len(fx); a++ {
			if !fx[a].onoff {
				fx = remFX(fx, a)
			}
		}
	}

}
func drawUpEnemies() { //MARK:DRAW UP ENEMIES

	clear := false
	for a := 0; a < len(level[roomNum].enemies); a++ {

		if !level[roomNum].enemies[a].off {

			//TIMERS
			if level[roomNum].enemies[a].T1 > 0 {
				level[roomNum].enemies[a].T1--
			}
			if level[roomNum].enemies[a].hppause > 0 {
				level[roomNum].enemies[a].hppause--
			}

			//PLAYER ENEMY CREC COLLIS
			if rl.CheckCollisionRecs(pl.crec, level[roomNum].enemies[a].crec) && pl.hppause == 0 {
				hitPL(0, 2)
			}

			shadowRec := level[roomNum].enemies[a].rec
			shadowRec.X -= 5
			shadowRec.Y += 5
			rl.DrawTexturePro(imgs, level[roomNum].enemies[a].img, shadowRec, ori, level[roomNum].enemies[a].ro, rl.Fade(rl.Black, 0.8))

			if level[roomNum].enemies[a].hppause > 0 {
				rl.DrawTexturePro(imgs, level[roomNum].enemies[a].img, level[roomNum].enemies[a].rec, ori, level[roomNum].enemies[a].ro, rl.Fade(ranCol(), level[roomNum].enemies[a].fade))
			} else {
				rl.DrawTexturePro(imgs, level[roomNum].enemies[a].img, level[roomNum].enemies[a].rec, ori, level[roomNum].enemies[a].ro, rl.Fade(level[roomNum].enemies[a].col, level[roomNum].enemies[a].fade))
			}

			if anchorT > 0 {
				siz := bsU
				rec := rl.NewRectangle(level[roomNum].enemies[a].crec.X+level[roomNum].enemies[a].crec.Width/2-siz/2, level[roomNum].enemies[a].crec.Y+level[roomNum].enemies[a].crec.Height/2-siz/2, siz, siz)
				rec.Y -= siz * 2
				rl.DrawTexturePro(imgs, etc[39], rec, ori, 0, ranCyan())
			}

			if debug {
				rl.DrawRectangleLinesEx(level[roomNum].enemies[a].rec, 1, rl.Red)
				rl.DrawRectangleLinesEx(level[roomNum].enemies[a].crec, 1, rl.White)
				rl.DrawRectangleLinesEx(level[roomNum].enemies[a].arec, 1, rl.Magenta)
				rl.DrawCircleV(level[roomNum].enemies[a].cnt, 4, rl.Red)
			}

			//ANIM
			switch level[roomNum].enemies[a].name {

			case "ghost", "slime", "rock":
				if frames%4 == 0 {
					if level[roomNum].enemies[a].imgr.X < level[roomNum].enemies[a].xImg2+(level[roomNum].enemies[a].img.Width*float32(level[roomNum].enemies[a].frameNum)) {
						level[roomNum].enemies[a].imgr.X += level[roomNum].enemies[a].img.Width
					} else {
						level[roomNum].enemies[a].imgr.X = level[roomNum].enemies[a].xImg2
					}
					if level[roomNum].enemies[a].imgl.X < level[roomNum].enemies[a].xImg+(level[roomNum].enemies[a].img.Width*float32(level[roomNum].enemies[a].frameNum)) {
						level[roomNum].enemies[a].imgl.X += level[roomNum].enemies[a].img.Width
					} else {
						level[roomNum].enemies[a].imgl.X = level[roomNum].enemies[a].xImg
					}
				}

				if level[roomNum].enemies[a].velX > 0 {
					level[roomNum].enemies[a].img = level[roomNum].enemies[a].imgr
				} else {
					level[roomNum].enemies[a].img = level[roomNum].enemies[a].imgl
				}

			case "mushroom":
				if frames%3 == 0 {
					if level[roomNum].enemies[a].imgr.X < level[roomNum].enemies[a].xImg2+(level[roomNum].enemies[a].img.Width*float32(level[roomNum].enemies[a].frameNum)) {
						level[roomNum].enemies[a].imgr.X += level[roomNum].enemies[a].img.Width
					} else {
						level[roomNum].enemies[a].imgr.X = level[roomNum].enemies[a].xImg2
					}
					if level[roomNum].enemies[a].imgl.X < level[roomNum].enemies[a].xImg+(level[roomNum].enemies[a].img.Width*float32(level[roomNum].enemies[a].frameNum)) {
						level[roomNum].enemies[a].imgl.X += level[roomNum].enemies[a].img.Width
					} else {
						level[roomNum].enemies[a].imgl.X = level[roomNum].enemies[a].xImg
					}
				}

				if level[roomNum].enemies[a].velX > 0 {
					level[roomNum].enemies[a].img = level[roomNum].enemies[a].imgr
				} else {
					level[roomNum].enemies[a].img = level[roomNum].enemies[a].imgl
				}

			case "spikehog":
				if frames%4 == 0 {
					if level[roomNum].enemies[a].img.X < level[roomNum].enemies[a].xImg+(level[roomNum].enemies[a].img.Width*float32(level[roomNum].enemies[a].frameNum)) {
						level[roomNum].enemies[a].img.X += level[roomNum].enemies[a].img.Width
					} else {
						level[roomNum].enemies[a].img.X = level[roomNum].enemies[a].xImg
					}
				}

			case "rabbit1":
				if frames%8 == 0 {
					if level[roomNum].enemies[a].img.X < rabbit1.xl+(rabbit1.frames*rabbit1.W) {
						level[roomNum].enemies[a].img.X += rabbit1.W
					} else {
						level[roomNum].enemies[a].img.X = rabbit1.xl
					}
				}
				switch level[roomNum].enemies[a].direc {
				case 1:
					level[roomNum].enemies[a].img.Y = rabbit1.yt + rabbit1.W
				case 2:
					level[roomNum].enemies[a].img.Y = rabbit1.yt + rabbit1.W*2
				case 3:
					level[roomNum].enemies[a].img.Y = rabbit1.yt
				case 4:
					level[roomNum].enemies[a].img.Y = rabbit1.yt + rabbit1.W*3
				}
			case "bat":
				if frames%8 == 0 {
					if level[roomNum].enemies[a].img.X < level[roomNum].enemies[a].xImg+(level[roomNum].enemies[a].img.Width*float32(level[roomNum].enemies[a].frameNum)) {
						level[roomNum].enemies[a].img.X += level[roomNum].enemies[a].img.Width
					} else {
						level[roomNum].enemies[a].img.X = level[roomNum].enemies[a].xImg
					}
				}
			}

			//ENEMY HP BAR
			if hpBarsOn {
				hpX := level[roomNum].enemies[a].rec.X + level[roomNum].enemies[a].rec.Width/2
				siz := float32(4)
				wid := float32(level[roomNum].enemies[a].hpmax) * (siz + 1)
				hpX -= wid / 2
				hpY := level[roomNum].enemies[a].rec.Y + level[roomNum].enemies[a].rec.Height + 5

				rec := rl.NewRectangle(hpX, hpY, siz, siz)
				for b := 0; b < level[roomNum].enemies[a].hpmax; b++ {
					rl.DrawRectangleLinesEx(rec, 1, rl.White)
					rec.X += siz + 1
				}
				rec = rl.NewRectangle(hpX, hpY, siz, siz)
				for b := 0; b < level[roomNum].enemies[a].hp; b++ {
					rl.DrawRectangleRec(rec, rl.Red)
					rec.X += siz + 1
				}
			}

			//MOVE
			if anchorT == 0 {
				moveenemy(a)
			}

			//MARK:PLAYER ENEMY ATTACK
			if rl.CheckCollisionRecs(pl.atkrec, level[roomNum].enemies[a].rec) {
				if level[roomNum].enemies[a].hppause == 0 {
					if pl.atk {
						level[roomNum].enemies[a].hppause = fps / 2
						level[roomNum].enemies[a].hp -= pl.atkDMG
						if level[roomNum].enemies[a].hp <= 0 {
							cntr := level[roomNum].enemies[a].cnt
							addkill(a)
							level[roomNum].enemies[a].off = true
							makeFX(2, cntr)
						} else {
							playenemyhit()
						}

						if mods.chainlightning && !chainLightingSwingOnOff {
							chainLightingSwingOnOff = true
							if roll6() > 2 {
								makeChainLightning()

							}
						}
					}
				}
			}

		} else {
			clear = true
		}
	}

	if clear {
		for a := 0; a < len(level[roomNum].enemies); a++ {
			if level[roomNum].enemies[a].off {
				level[roomNum].enemies = remEnemy(level[roomNum].enemies, a)
			}
		}
	}

}
func drawUpBoss() { //MARK: DRAW UP BOSS

	//IMG
	shadowrec := bosses[bossnum].rec
	shadowrec.X -= 4
	shadowrec.Y += 4
	rl.DrawTexturePro(imgs, bosses[bossnum].img, shadowrec, ori, 0, rl.Fade(rl.Black, 0.7))

	if bosses[bossnum].hppause > 0 {
		rl.DrawTexturePro(imgs, bosses[bossnum].img, bosses[bossnum].rec, ori, 0, ranCol())
	} else {
		rl.DrawTexturePro(imgs, bosses[bossnum].img, bosses[bossnum].rec, ori, 0, rl.White)
	}

	if frames%6 == 0 {
		bosses[bossnum].img.X += 48
		if bosses[bossnum].img.X > bosses[bossnum].xl+bosses[bossnum].img.Width*2 {
			bosses[bossnum].img.X = bosses[bossnum].xl
		}
	}
	switch bosses[bossnum].direc {
	case 1:
		bosses[bossnum].img.Y = bosses[bossnum].yt + bosses[bossnum].img.Height*3
	case 2:
		bosses[bossnum].img.Y = bosses[bossnum].yt + bosses[bossnum].img.Height*2
	case 3:
		bosses[bossnum].img.Y = bosses[bossnum].yt
	case 4:
		bosses[bossnum].img.Y = bosses[bossnum].yt + bosses[bossnum].img.Height
	}
	if debug {
		rl.DrawRectangleLinesEx(bosses[bossnum].rec, 1, rl.White)
		rl.DrawRectangleLinesEx(bosses[bossnum].crec, 1, rl.White)
	}

	//HP BARS
	if hpBarsOn {
		hpX := bosses[bossnum].rec.X + bosses[bossnum].rec.Width/2
		siz := float32(4)
		wid := float32(bosses[bossnum].hpmax) * (siz + 1)
		hpX -= wid / 2
		hpY := bosses[bossnum].rec.Y + bosses[bossnum].rec.Height + 5

		rec := rl.NewRectangle(hpX, hpY, siz, siz)
		for b := 0; b < bosses[bossnum].hpmax; b++ {
			rl.DrawRectangleLinesEx(rec, 1, rl.White)
			rec.X += siz + 1
		}
		rec = rl.NewRectangle(hpX, hpY, siz, siz)
		for b := 0; b < bosses[bossnum].hp; b++ {
			rl.DrawRectangleRec(rec, rl.Red)
			rec.X += siz + 1
		}
	}

	//TIMERS
	bosses[bossnum].timer--
	if bosses[bossnum].timer <= 0 {
		bosses[bossnum].timer = fps * rI32(1, 4)
		switch bosses[bossnum].atkType {
		case 1:
			makeProjectileEnemy(7, bosses[bossnum].cnt)
		case 2:
			makeProjectileEnemy(8, bosses[bossnum].cnt)
		case 3:
			makeProjectileEnemy(9, bosses[bossnum].cnt)
		}
	}
	if bosses[bossnum].hppause > 0 {
		bosses[bossnum].hppause--
	}

	//MARK:PLAYER BOSS ATTACK
	if rl.CheckCollisionRecs(pl.atkrec, bosses[bossnum].crec) {
		if bosses[bossnum].hppause == 0 {
			if pl.atk {
				bosses[bossnum].hppause = fps
				bosses[bossnum].hp -= pl.atkDMG
				if bosses[bossnum].hp <= 0 {
					cntr := bosses[bossnum].cnt
					bosses[bossnum].off = true
					makeFX(2, cntr)
					if !endgame {
						minsEND = mins
						secsEND = secs
						addtime()
						endgame = true
					}
				}
				rl.PlaySound(sfx[29])
			}
		}
	}

	//MOVE
	moveboss()

}

// MARK: CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK CHECK
func checkcontroller() { //MARK:CHECK CONTROLLER

	isController = rl.IsGamepadAvailable(0)
	if isController && contolleron {
		useController = true
	} else if !isController {
		useController = false
		controllerDisconnect = true
	}

	if isController {
		contolleron = true
		controllerDisconnect = false
		controllerWasOn = true
	}

	if controllerDisconnect && controllerWasOn && !pause {
		contolleron = false
		controllerWasOn = false
		optionson = true
		optionnum = 0
		pause = true
	}
}
func checknextmoveV2innerRec(cnt rl.Vector2, velx, velxy float32) bool { //MARK:CHECK NEXT MOVE V2 POINT INNER REC
	canmove := true
	cnt.X += velx
	cnt.Y += velxy
	if !rl.CheckCollisionPointRec(cnt, levRecInner) {
		canmove = false
	}
	return canmove
}

func checkNextMove(rec rl.Rectangle, velx, vely float32, destroyEtc bool) bool { //MARK:CHECK NEXT MOVE

	canmove := true

	nextRec := rec
	nextRec.X += velx
	nextRec.Y += vely

	tl, tr, br, bl := rl.NewVector2(nextRec.X, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y+nextRec.Height), rl.NewVector2(nextRec.X, nextRec.Y+nextRec.Height)

	//CHECK INNER REC

	if !rl.CheckCollisionPointRec(tl, levRecInner) || !rl.CheckCollisionPointRec(tr, levRecInner) || !rl.CheckCollisionPointRec(br, levRecInner) || !rl.CheckCollisionPointRec(bl, levRecInner) {
		canmove = false
	}

	//CHECK MOVE BLOCKS
	if canmove {
		if len(level[roomNum].movBloks) > 0 {
			for a := 0; a < len(level[roomNum].movBloks); a++ {
				if rl.CheckCollisionRecs(nextRec, level[roomNum].movBloks[a].rec) {
					canmove = false
				}
			}
		}
	}

	//CHECK ETC
	if canmove {
		if len(level[roomNum].etc) > 0 {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].solid && level[roomNum].etc[a].name != "turret" {
					if rl.CheckCollisionRecs(level[roomNum].etc[a].rec, nextRec) {
						canmove = false
						if destroyEtc {
							switch level[roomNum].etc[a].name {
							case "powerupBlok":
								destroyPowerupBlok(a)
								makeFX(3, level[roomNum].etc[a].cnt)
								level[roomNum].etc[a].onoff = false
								rl.PlaySound(sfx[6])
							case "oilbarrel":
								makeFX(4, level[roomNum].etc[a].cnt)
								level[roomNum].etc[a].onoff = false
								rl.PlaySound(sfx[10])
								rl.PlaySound(sfx[5])
							}
						}
					}
				}
			}
		}
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(level[roomNum].innerBloks) > 0 {
			for a := 0; a < len(level[roomNum].innerBloks); a++ {
				if level[roomNum].innerBloks[a].solid {
					if rl.CheckCollisionRecs(level[roomNum].innerBloks[a].crec, nextRec) {
						canmove = false
					}
				}
			}
		}
	}

	return canmove

}
func checkInnerBloksExits(roomN int, rec rl.Rectangle) bool { //MARK:CHECK INNER BLOK EXITS

	canadd := true

	for a := 0; a < len(level[roomN].doorExitRecs); a++ {
		if rl.CheckCollisionRecs(rec, level[roomN].doorExitRecs[a]) {
			canadd = false
		}
	}

	if canadd {
		for a := 0; a < len(level[roomN].innerBloks); a++ {
			if rl.CheckCollisionRecs(rec, level[roomN].innerBloks[a].rec) {
				canadd = false
			}
		}
	}

	return canadd
}
func checkMoveBlok(blokNum int) bool { //MARK:CHECK MOVE BLOK

	canmove := true
	checkBlok := level[roomNum].movBloks[blokNum]
	nextRec := checkBlok.rec
	if checkBlok.velX != 0 {
		nextRec.X += checkBlok.velX
	}
	if checkBlok.velY != 0 {
		nextRec.Y += checkBlok.velY
	}

	//CHECK PLAYER
	if rl.CheckCollisionRecs(nextRec, pl.crec) {
		canmove = false
	}

	//CHECK MOVE BLOKS
	if canmove {
		if len(level[roomNum].movBloks) > 1 {
			for a := 0; a < len(level[roomNum].movBloks); a++ {
				if blokNum != a {
					if rl.CheckCollisionRecs(nextRec, level[roomNum].movBloks[a].rec) {
						canmove = false
					}
				}
			}
		}
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(level[roomNum].innerBloks) > 0 {
			for a := 0; a < len(level[roomNum].innerBloks); a++ {
				if rl.CheckCollisionRecs(nextRec, level[roomNum].innerBloks[a].crec) {
					canmove = false
				}
			}
		}
	}
	//CHECK BOUNDARY
	if canmove {
		checkV1 := rl.NewVector2(nextRec.X, nextRec.Y)
		checkV2 := checkV1
		checkV2.X += nextRec.Width
		checkV3 := checkV2
		checkV3.Y += nextRec.Height
		checkV4 := checkV1
		checkV4.Y += nextRec.Height

		if !rl.CheckCollisionPointRec(checkV1, levRecInner) || !rl.CheckCollisionPointRec(checkV2, levRecInner) || !rl.CheckCollisionPointRec(checkV3, levRecInner) || !rl.CheckCollisionPointRec(checkV4, levRecInner) {
			canmove = false
		}
	}

	return canmove

}
func checkplayermove(direc int) bool { //MARK:CHECK PLAYER MOVE

	canmove := true
	nextRec := pl.crec

	switch direc {
	case 1:
		nextRec.Y -= pl.vel
	case 2:
		nextRec.X += pl.vel
	case 3:
		nextRec.Y += pl.vel
	case 4:
		nextRec.X -= pl.vel
	}

	tl, tr, br, bl := rl.NewVector2(nextRec.X, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y), rl.NewVector2(nextRec.X+nextRec.Width, nextRec.Y+nextRec.Height), rl.NewVector2(nextRec.X, nextRec.Y+nextRec.Height)

	//CHECK BOUNDARY WALLS
	dWalls := level[roomNum].walls
	for a := 0; a < len(dWalls); a++ {
		if rl.CheckCollisionPointRec(tl, dWalls[a].rec) || rl.CheckCollisionPointRec(tr, dWalls[a].rec) || rl.CheckCollisionPointRec(br, dWalls[a].rec) || rl.CheckCollisionPointRec(bl, dWalls[a].rec) {
			canmove = false
		}
		if rl.CheckCollisionRecs(nextRec, dWalls[a].rec) {
			canmove = false
		}
	}

	//CHECK MOVE BLOCKS
	if canmove {
		if len(level[roomNum].movBloks) > 0 {
			for a := 0; a < len(level[roomNum].movBloks); a++ {
				if rl.CheckCollisionRecs(nextRec, level[roomNum].movBloks[a].rec) {
					canmove = false
					if level[roomNum].movBloks[a].bump {
						switch direc {
						case 1:
							level[roomNum].movBloks[a].velY += -pl.vel / 8
						case 2:
							level[roomNum].movBloks[a].velX += pl.vel / 8
						case 3:
							level[roomNum].movBloks[a].velY += pl.vel / 8
						case 4:
							level[roomNum].movBloks[a].velX -= pl.vel / 8
						}
					}
				}

			}
		}
	}

	//CHECK ETC
	if canmove {
		if len(level[roomNum].etc) > 0 {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].solid {
					if rl.CheckCollisionRecs(level[roomNum].etc[a].rec, nextRec) {
						canmove = false
					}
				}
			}
		}
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(level[roomNum].innerBloks) > 0 {
			for a := 0; a < len(level[roomNum].innerBloks); a++ {
				if level[roomNum].innerBloks[a].solid {
					if rl.CheckCollisionRecs(level[roomNum].innerBloks[a].crec, nextRec) {
						canmove = false
					}
				}
			}
		}

	}

	return canmove
}

// MARK: FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND FIND
func findRecPoswithSpacing(wid, space float32, numRoom int) (rec rl.Rectangle, found bool) { //MARK: FIND REC POS WITH SPACING

	countbreak := 500
	found = true

	wid2 := wid + (space * 2)

	for {
		rec = rl.NewRectangle(rF32(levRecInner.X, levRecInner.X+levRecInner.Width-wid2), rF32(levRecInner.Y, levRecInner.Y+levRecInner.Width-wid2), wid2, wid2)
		canadd := true

		for a := 0; a < len(level[numRoom].doorExitRecs); a++ {
			if rl.CheckCollisionRecs(rec, level[numRoom].doorExitRecs[a]) {
				canadd = false
			}
		}
		if canadd {
			for a := 0; a < len(level[numRoom].movBloks); a++ {
				if rl.CheckCollisionRecs(rec, level[numRoom].movBloks[a].rec) {
					canadd = false
				}
			}
		}
		if canadd {
			for a := 0; a < len(level[numRoom].innerBloks); a++ {
				if rl.CheckCollisionRecs(rec, level[numRoom].innerBloks[a].rec) {
					canadd = false
				}
			}
		}
		if canadd {
			for a := 0; a < len(level[numRoom].etc); a++ {
				if rl.CheckCollisionRecs(rec, level[numRoom].etc[a].rec) && level[numRoom].etc[a].solid {
					canadd = false
				}
			}
		}

		countbreak--
		if countbreak == 0 || canadd {

			rec.X += space
			rec.Y += space
			rec.Width -= space * 2
			rec.Height -= space * 2

			if countbreak == 0 {
				found = false
			}
			break
		}

	}

	return rec, found
}
func findRecPos(wid float32, numRoom int) (rec rl.Rectangle, found bool) { //MARK: FIND REC POS

	countbreak := 500
	found = true

	for {
		rec = rl.NewRectangle(rF32(levRecInner.X, levRecInner.X+levRecInner.Width-wid), rF32(levRecInner.Y, levRecInner.Y+levRecInner.Width-wid), wid, wid)
		canadd := true
		for a := 0; a < len(level[numRoom].movBloks); a++ {
			if rl.CheckCollisionRecs(rec, level[numRoom].movBloks[a].rec) {
				canadd = false
			}
		}
		if canadd {
			for a := 0; a < len(level[numRoom].innerBloks); a++ {
				if rl.CheckCollisionRecs(rec, level[numRoom].innerBloks[a].rec) {
					canadd = false
				}
			}
		}
		if canadd {
			for a := 0; a < len(level[numRoom].etc); a++ {
				if rl.CheckCollisionRecs(rec, level[numRoom].etc[a].rec) && level[numRoom].etc[a].solid {
					canadd = false
				}
			}
		}

		countbreak--
		if countbreak == 0 || canadd {
			if countbreak == 0 {
				found = false
			}
			break
		}

	}

	return rec, found
}
func findRanCntV2() rl.Vector2 { //MARK: FIND RANDOM CNTR V2
	v2 := cnt
	wid := levW / 2
	wid -= bsU2
	v2.X += rF32(-wid, wid)
	v2.Y += rF32(-wid, wid)
	return v2
}
func findRanRecLoc(w, h float32, roomNum int) (tl rl.Vector2) { //MARK: FIND RANDOM RECTANGLE LOCATION

	v2 := rl.Vector2{}
	countbreak := 100
	for {
		v2 = rl.NewVector2(levRecInner.X+bsU2, levRecInner.Y+bsU2)
		wid := levRecInner.Width - bsU4
		heig := wid
		wid -= w
		heig -= h
		v2.X += rF32(0, wid)
		v2.Y += rF32(0, heig)

		checkrec := rl.NewRectangle(v2.X, v2.Y, w, h)
		canadd := true
		for a := 0; a < len(level[roomNum].doorExitRecs); a++ {
			if rl.CheckCollisionRecs(checkrec, level[roomNum].doorExitRecs[a]) {
				canadd = false
			}
		}

		countbreak--
		if canadd || countbreak == 0 {
			break
		}
	}

	return v2
}

// MARK: MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE MOVE
func moveboss() { //MARK:MOVE BOSS

	checkRec := bosses[bossnum].crec
	checkRec.X += bosses[bossnum].velX
	checkRec.Y += bosses[bossnum].velY

	canmove := true

	tl := rl.NewVector2(checkRec.X, checkRec.Y)
	tr := tl
	tr.X += checkRec.Width
	br := tr
	br.Y += checkRec.Height
	bl := tl
	bl.Y += checkRec.Height

	//CHECK BOUNDARY
	if !rl.CheckCollisionPointRec(tl, levRecInner) || !rl.CheckCollisionPointRec(tr, levRecInner) || !rl.CheckCollisionPointRec(br, levRecInner) || !rl.CheckCollisionPointRec(bl, levRecInner) {
		canmove = false
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(level[roomNum].innerBloks) > 0 {
			for a := 0; a < len(level[roomNum].innerBloks); a++ {
				if rl.CheckCollisionRecs(level[roomNum].innerBloks[a].crec, checkRec) {
					canmove = false
				}
			}
		}
	}

	//MOVE
	if canmove {

		bosses[bossnum].rec.X += bosses[bossnum].velX
		bosses[bossnum].rec.Y += bosses[bossnum].velY
		bosses[bossnum].crec = bosses[bossnum].rec
		bosses[bossnum].crec.X += bosses[bossnum].rec.Width / 4
		bosses[bossnum].crec.Y += bosses[bossnum].rec.Height / 5
		bosses[bossnum].crec.Width -= bosses[bossnum].rec.Width / 2
		bosses[bossnum].crec.Height -= bosses[bossnum].rec.Height / 5
		bosses[bossnum].cnt = rl.NewVector2(bosses[bossnum].rec.X+bosses[bossnum].rec.Width/2, bosses[bossnum].rec.Y+bosses[bossnum].rec.Height/2)

		//CHECK SOCKS COLLIS
		if mods.socks && bosses[bossnum].hppause == 0 {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].name == "footprints" && bosses[bossnum].hppause == 0 && rl.CheckCollisionRecs(level[roomNum].etc[a].rec, bosses[bossnum].crec) {
					bosses[bossnum].hppause = fps
					bosses[bossnum].hp--
					if bosses[bossnum].hp <= 0 {
						cntr := bosses[bossnum].cnt
						bosses[bossnum].off = true
						makeFX(2, cntr)
						if !endgame {
							minsEND = mins
							secsEND = secs
							addtime()
							endgame = true
						}
					}
					rl.PlaySound(sfx[29])
				}
			}
		}
		//CHECK ORBITAL COLLIS
		if mods.orbital && bosses[bossnum].hppause == 0 {
			if rl.CheckCollisionRecs(pl.orbrec1, bosses[bossnum].crec) {
				bosses[bossnum].hppause = fps
				bosses[bossnum].hp--
				if bosses[bossnum].hp <= 0 {
					cntr := bosses[bossnum].cnt
					bosses[bossnum].off = true
					makeFX(2, cntr)
					if !endgame {
						minsEND = mins
						secsEND = secs
						addtime()
						endgame = true
					}
					rl.PlaySound(sfx[29])
				}
			}

			if mods.orbitalN == 2 {
				if rl.CheckCollisionRecs(pl.orbrec2, bosses[bossnum].crec) {
					bosses[bossnum].hppause = fps
					bosses[bossnum].hp--
					if bosses[bossnum].hp <= 0 {
						cntr := bosses[bossnum].cnt
						bosses[bossnum].off = true
						makeFX(2, cntr)
						if !endgame {
							minsEND = mins
							secsEND = secs
							addtime()
							endgame = true
						}
					}
					rl.PlaySound(sfx[29])
				}
			}

		}
		//CHECK AIRSTRIKE COLLIS
		if mods.airstrike && bosses[bossnum].hppause == 0 {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].name == "airbomb" && rl.CheckCollisionRecs(level[roomNum].etc[a].crec, bosses[bossnum].crec) {
					bosses[bossnum].hppause = fps
					bosses[bossnum].hp--
					if bosses[bossnum].hp <= 0 {
						cntr := bosses[bossnum].cnt
						bosses[bossnum].off = true
						makeFX(2, cntr)
						if !endgame {
							minsEND = mins
							secsEND = secs
							addtime()
							endgame = true
						}
						rl.PlaySound(sfx[29])
					}
				}
			}
		}
		//CHECK FIRETRAIL COLLIS
		if mods.firetrail {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].name == "flamefiretrail" && bosses[bossnum].hppause == 0 && rl.CheckCollisionRecs(level[roomNum].etc[a].rec, bosses[bossnum].crec) {
					bosses[bossnum].hppause = fps
					bosses[bossnum].hp--
					if bosses[bossnum].hp <= 0 && !endgame {
						pause = true
						endPauseT = fps * 5
						cntr := bosses[bossnum].cnt
						bosses[bossnum].off = true
						makeFX(2, cntr)
						if !endgame {
							minsEND = mins
							secsEND = secs
							addtime()
							endgame = true
						}
					}
					rl.PlaySound(sfx[29])
				}
			}
		}

	} else {
		bosses[bossnum].velX = rF32(-bosses[bossnum].vel, bosses[bossnum].vel)
		bosses[bossnum].velY = rF32(-bosses[bossnum].vel, bosses[bossnum].vel)
	}

	//FIND DIREC FOR ANIM
	if getabs(bosses[bossnum].velX) > getabs(bosses[bossnum].velY) {
		if bosses[bossnum].velX > 0 {
			bosses[bossnum].direc = 2
		} else {
			bosses[bossnum].direc = 4
		}
	} else {
		if bosses[bossnum].velY > 0 {
			bosses[bossnum].direc = 3
		} else {
			bosses[bossnum].direc = 1
		}

	}

}

func moveFollow(v2, targetV2 rl.Vector2, vel float32) (velx, vely float32) { //MARK:MOVE FOLLOW

	xdiff := absdiff(v2.X, targetV2.X)
	ydiff := absdiff(v2.Y, targetV2.Y)

	if xdiff > ydiff {
		velx = vel
		vely = ydiff / (xdiff / vel)
		if v2.X > targetV2.X {
			velx = -velx
		}
		if v2.Y > targetV2.Y {
			vely = -vely
		}
	} else {
		vely = vel
		velx = xdiff / (ydiff / vel)
		if v2.X > targetV2.X {
			velx = -velx
		}
		if v2.Y > targetV2.Y {
			vely = -vely
		}
	}

	return velx, vely
}
func moveenemy(num int) { //MARK:MOVE ENEMY

	checkRec := level[roomNum].enemies[num].rec
	checkRec.X += level[roomNum].enemies[num].velX
	checkRec.Y += level[roomNum].enemies[num].velY

	canmove := true

	tl := rl.NewVector2(checkRec.X, checkRec.Y)
	tr := tl
	tr.X += checkRec.Width
	br := tr
	br.Y += checkRec.Height
	bl := tl
	bl.Y += checkRec.Height

	//CHECK BOUNDARY
	if !rl.CheckCollisionPointRec(tl, levRecInner) || !rl.CheckCollisionPointRec(tr, levRecInner) || !rl.CheckCollisionPointRec(br, levRecInner) || !rl.CheckCollisionPointRec(bl, levRecInner) {
		canmove = false
	}

	//CHECK INNER BLOKS
	if canmove {
		if len(level[roomNum].innerBloks) > 0 {
			for a := 0; a < len(level[roomNum].innerBloks); a++ {
				if rl.CheckCollisionRecs(level[roomNum].innerBloks[a].crec, checkRec) {
					canmove = false
				}
			}
		}
	}

	//MOVE
	if canmove {

		level[roomNum].enemies[num].rec = checkRec
		level[roomNum].enemies[num].crec = level[roomNum].enemies[num].rec

		//MOVE COLLIS REC & AREA REC
		switch level[roomNum].enemies[num].name {
		case "mushroom":
			level[roomNum].enemies[num].crec = level[roomNum].enemies[num].rec
			level[roomNum].enemies[num].crec.Y += level[roomNum].enemies[num].crec.Height / 2
			level[roomNum].enemies[num].crec.Height -= level[roomNum].enemies[num].crec.Height / 2
		case "slime", "rock":
			level[roomNum].enemies[num].crec = level[roomNum].enemies[num].rec
			level[roomNum].enemies[num].crec.Y += level[roomNum].enemies[num].crec.Height / 3
			level[roomNum].enemies[num].crec.Height -= level[roomNum].enemies[num].crec.Height / 3
		case "spikehog":
			level[roomNum].enemies[num].crec.Height = level[roomNum].enemies[num].crec.Height / 2
			level[roomNum].enemies[num].crec.Y += level[roomNum].enemies[num].crec.Height
		case "ghost":
			level[roomNum].enemies[num].crec.Y += level[roomNum].enemies[num].crec.Height / 3
			level[roomNum].enemies[num].crec.Height -= level[roomNum].enemies[num].crec.Height / 3
			level[roomNum].enemies[num].crec.X += level[roomNum].enemies[num].crec.Width / 8
			level[roomNum].enemies[num].crec.Width -= level[roomNum].enemies[num].crec.Width / 4
			level[roomNum].enemies[num].arec = level[roomNum].enemies[num].crec
			level[roomNum].enemies[num].arec.X -= level[roomNum].enemies[num].arec.Width * 2
			level[roomNum].enemies[num].arec.Y -= level[roomNum].enemies[num].arec.Width * 2
			level[roomNum].enemies[num].arec.Width = level[roomNum].enemies[num].arec.Width * 5
			level[roomNum].enemies[num].arec.Height = level[roomNum].enemies[num].arec.Height * 5

		}

		level[roomNum].enemies[num].cnt = rl.NewVector2(checkRec.X+checkRec.Width/2, checkRec.Y+checkRec.Height/2)

		//CHECK SOCKS COLLIS
		if mods.socks && level[roomNum].enemies[num].hppause == 0 && !level[roomNum].enemies[num].fly {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].name == "footprints" && level[roomNum].enemies[num].hppause == 0 && rl.CheckCollisionRecs(level[roomNum].etc[a].rec, level[roomNum].enemies[num].crec) {
					level[roomNum].enemies[num].hppause = fps / 2
					level[roomNum].enemies[num].hp--
					if level[roomNum].enemies[num].hp <= 0 {
						cntr := level[roomNum].enemies[num].cnt
						addkill(num)
						level[roomNum].enemies[num].off = true
						makeFX(2, cntr)
					} else {
						playenemyhit()
					}
				}
			}
		}
		//CHECK ORBITAL COLLIS
		if mods.orbital && level[roomNum].enemies[num].hppause == 0 {

			if rl.CheckCollisionRecs(pl.orbrec1, level[roomNum].enemies[num].crec) {
				level[roomNum].enemies[num].hppause = fps / 2
				level[roomNum].enemies[num].hp--
				if level[roomNum].enemies[num].hp <= 0 {
					cntr := level[roomNum].enemies[num].cnt
					addkill(num)
					level[roomNum].enemies[num].off = true
					makeFX(2, cntr)
				} else {
					playenemyhit()
				}
			}

			if mods.orbitalN == 2 {
				if rl.CheckCollisionRecs(pl.orbrec2, level[roomNum].enemies[num].crec) {
					level[roomNum].enemies[num].hppause = fps / 2
					level[roomNum].enemies[num].hp--
					if level[roomNum].enemies[num].hp <= 0 {
						cntr := level[roomNum].enemies[num].cnt
						addkill(num)
						level[roomNum].enemies[num].off = true
						makeFX(2, cntr)
					} else {
						if flipcoin() {
							rl.PlaySound(sfx[1])
						} else {
							rl.PlaySound(sfx[2])
						}
					}
				}
			}

		}
		//CHECK AIRSTRIKE COLLIS
		if mods.airstrike {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].name == "airbomb" && rl.CheckCollisionRecs(level[roomNum].etc[a].crec, level[roomNum].enemies[num].crec) {
					level[roomNum].enemies[num].hp = 0
					cntr := level[roomNum].enemies[num].cnt
					addkill(num)
					level[roomNum].enemies[num].off = true
					makeFX(2, cntr)
				}
			}
		}
		//CHECK FIRETRAIL COLLIS
		if mods.firetrail && !level[roomNum].enemies[num].fly {
			for a := 0; a < len(level[roomNum].etc); a++ {
				if level[roomNum].etc[a].name == "flamefiretrail" && level[roomNum].enemies[num].hppause == 0 && rl.CheckCollisionRecs(level[roomNum].etc[a].rec, level[roomNum].enemies[num].crec) {
					level[roomNum].enemies[num].hppause = fps / 2
					level[roomNum].enemies[num].hp--
					if level[roomNum].enemies[num].hp <= 0 {
						cntr := level[roomNum].enemies[num].cnt
						addkill(num)
						level[roomNum].enemies[num].off = true
						makeFX(2, cntr)
					} else {
						playenemyhit()
					}
				}
			}
		}
		//PLAYER ENEMY AREC COLLIS
		if !mods.invisible {
			if rl.CheckCollisionRecs(pl.rec, level[roomNum].enemies[num].arec) {
				if level[roomNum].enemies[num].name == "ghost" {
					level[roomNum].enemies[num].velX, level[roomNum].enemies[num].velY = moveFollow(level[roomNum].enemies[num].cnt, pl.cnt, level[roomNum].enemies[num].vel)
				}
			}
		}

		switch level[roomNum].enemies[num].name {

		case "mushroom":
			if level[roomNum].enemies[num].T1 == 0 {
				level[roomNum].enemies[num].T1 = rI32(int(fps*2), int(fps*5))
				zproj := xproj{}
				zproj.img = mushBull.recTL
				zproj.onoff = true
				zproj.cnt = level[roomNum].enemies[num].cnt
				zproj.vel = bsU / 5
				siz := bsU2
				zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
				zproj.velx, zproj.vely = moveFollow(level[roomNum].enemies[num].cnt, pl.cnt, zproj.vel)
				zproj.fade = 1
				zproj.col = ranRed()
				zproj.name = "mushbull"
				enProj = append(enProj, zproj)
			}

		case "rock":
			if level[roomNum].enemies[num].T1 == 0 && level[roomNum].enemies[num].spawnN > 0 {
				level[roomNum].enemies[num].spawnN--
				level[roomNum].enemies[num].T1 = rI32(int(fps*2), int(fps*5))
				zen := xenemy{}
				zen = enRock
				zen.spawnN = 0
				zen.rec = level[roomNum].enemies[num].rec
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				level[roomNum].enemies = append(level[roomNum].enemies, zen)
			}
			if roll18() == 18 {
				level[roomNum].enemies[num].velX = rF32(-level[roomNum].enemies[num].vel, level[roomNum].enemies[num].vel)
				level[roomNum].enemies[num].velY = rF32(-level[roomNum].enemies[num].vel, level[roomNum].enemies[num].vel)
			}
		case "slime":
			if level[roomNum].enemies[num].T1 == 0 {
				level[roomNum].enemies[num].T1 = rI32(int(fps/2), int(fps*2))
				zblok := xblok{}
				zblok.rec = level[roomNum].enemies[num].rec
				zblok.crec = zblok.rec
				zblok.crec.X += zblok.crec.Width / 8
				zblok.crec.Y += zblok.crec.Width / 8
				zblok.crec.Width -= zblok.crec.Width / 4
				zblok.crec.Height -= zblok.crec.Height / 4
				zblok.cnt = level[roomNum].enemies[num].cnt
				zblok.img = splats[rInt(0, len(splats))]
				zblok.color = ranGreen()
				zblok.fade = 0.7
				zblok.onoff = true
				zblok.name = "slimetrail"
				level[roomNum].etc = append(level[roomNum].etc, zblok)
			}
			if roll36() == 36 {
				level[roomNum].enemies[num].velX = rF32(-level[roomNum].enemies[num].vel, level[roomNum].enemies[num].vel)
				level[roomNum].enemies[num].velY = rF32(-level[roomNum].enemies[num].vel, level[roomNum].enemies[num].vel)
			}
		case "spikehog":
			if roll18() == 18 {
				level[roomNum].enemies[num].velX = 0
				level[roomNum].enemies[num].velY = 0
				if flipcoin() {
					level[roomNum].enemies[num].velX = level[roomNum].enemies[num].vel
					if flipcoin() {
						level[roomNum].enemies[num].velX *= -1
					}
				} else {
					level[roomNum].enemies[num].velY = level[roomNum].enemies[num].vel
					if flipcoin() {
						level[roomNum].enemies[num].velY *= -1
					}
				}
			}
		}

	} else {
		switch level[roomNum].enemies[num].name {

		case "ghost", "slime", "rock", "mushroom":
			level[roomNum].enemies[num].velX = rF32(-level[roomNum].enemies[num].vel, level[roomNum].enemies[num].vel)
			level[roomNum].enemies[num].velY = rF32(-level[roomNum].enemies[num].vel, level[roomNum].enemies[num].vel)
		case "spikehog":
			level[roomNum].enemies[num].velX = 0
			level[roomNum].enemies[num].velY = 0
			if flipcoin() {
				level[roomNum].enemies[num].velX = level[roomNum].enemies[num].vel
				if flipcoin() {
					level[roomNum].enemies[num].velX *= -1
				}
			} else {
				level[roomNum].enemies[num].velY = level[roomNum].enemies[num].vel
				if flipcoin() {
					level[roomNum].enemies[num].velY *= -1
				}
			}
		case "bat", "rabbit1":
			if level[roomNum].enemies[num].velX > 0 {
				level[roomNum].enemies[num].velX = rF32(-level[roomNum].enemies[num].vel, 0)
			} else if level[roomNum].enemies[num].velX < 0 {
				level[roomNum].enemies[num].velX = rF32(0, level[roomNum].enemies[num].vel)
			}

			if level[roomNum].enemies[num].velY > 0 {
				level[roomNum].enemies[num].velY = rF32(-level[roomNum].enemies[num].vel, 0)
			} else if level[roomNum].enemies[num].velY < 0 {
				level[roomNum].enemies[num].velY = rF32(0, level[roomNum].enemies[num].vel)
			}
		}
	}

	//FIND DIREC FOR ANIM
	if getabs(level[roomNum].enemies[num].velX) > getabs(level[roomNum].enemies[num].velY) {
		if level[roomNum].enemies[num].velX > 0 {
			level[roomNum].enemies[num].direc = 2
		} else {
			level[roomNum].enemies[num].direc = 4
		}
	} else {
		if level[roomNum].enemies[num].velY > 0 {
			level[roomNum].enemies[num].direc = 3
		} else {
			level[roomNum].enemies[num].direc = 1
		}
	}

}

func movebloks() { //MARK:MOVE BLOKS

	for a := 0; a < len(level[roomNum].movBloks); a++ {

		if checkMoveBlok(a) {

			level[roomNum].movBloks[a].cnt.X += level[roomNum].movBloks[a].velX
			level[roomNum].movBloks[a].cnt.Y += level[roomNum].movBloks[a].velY
			level[roomNum].movBloks[a].rec = rl.NewRectangle(level[roomNum].movBloks[a].cnt.X-level[roomNum].movBloks[a].rec.Width/2, level[roomNum].movBloks[a].cnt.Y-level[roomNum].movBloks[a].rec.Height/2, level[roomNum].movBloks[a].rec.Width, level[roomNum].movBloks[a].rec.Height)

		} else {
			switch level[roomNum].movBloks[a].movType {
			case 1, 2: //LR UD
				if level[roomNum].movBloks[a].velX > 0 {
					level[roomNum].movBloks[a].velX = -level[roomNum].movBloks[a].velX
				} else if level[roomNum].movBloks[a].velX < 0 {
					level[roomNum].movBloks[a].velX = getabs(level[roomNum].movBloks[a].velX)
				}
				if level[roomNum].movBloks[a].velY > 0 {
					level[roomNum].movBloks[a].velY = -level[roomNum].movBloks[a].velY
				} else if level[roomNum].movBloks[a].velY < 0 {
					level[roomNum].movBloks[a].velY = getabs(level[roomNum].movBloks[a].velY)
				}
			}
		}

		if rl.CheckCollisionRecs(pl.crec, level[roomNum].movBloks[a].rec) {
			level[roomNum].movBloks[a].cnt.X -= level[roomNum].movBloks[a].rec.Width * 2
			level[roomNum].movBloks[a].cnt.Y -= level[roomNum].movBloks[a].rec.Width * 2
			level[roomNum].movBloks[a].rec = rl.NewRectangle(level[roomNum].movBloks[a].cnt.X-level[roomNum].movBloks[a].rec.Width/2, level[roomNum].movBloks[a].cnt.Y-level[roomNum].movBloks[a].rec.Height/2, level[roomNum].movBloks[a].rec.Width, level[roomNum].movBloks[a].rec.Height)
		}
	}

}

// MARK: ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC ETC
func savesettings() { //MARK:SAVE SETTINGS

	f, err := os.Create("etc/st.000")
	if err != nil {
		fmt.Println(err)
		return
	}

	settingsTXT := ""

	if hpBarsOn {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if scanlineson {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if artifactson {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if shaderon {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if platkrecon {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if invincible {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if useController {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}
	if musicon {
		settingsTXT = settingsTXT + "1,"
	} else {
		settingsTXT = settingsTXT + "0,"
	}

	musicTXT := fmt.Sprint(bgMusicNum)
	settingsTXT = settingsTXT + musicTXT + ","

	voltxt := fmt.Sprintf("%.0f", volume*10)
	settingsTXT = settingsTXT + voltxt + ","

	if hardcore {
		settingsTXT = settingsTXT + "1"
	} else {
		settingsTXT = settingsTXT + "0"
	}

	_, err = f.WriteString(settingsTXT)
	if err != nil {
		fmt.Printf("Failed to write settings: %s", err)
	}
}

func addtime() { //MARK:ADD TIME

	totaltime := minsEND*60 + secsEND

	checktime := totaltime
	canadd := false
	changenum := 0

	for i := 0; i < len(times); i++ {
		if checktime < times[i] {
			canadd = true
			checktime = times[i]
			changenum = i
		}
	}
	if canadd {
		times[changenum] = totaltime
		sort.Ints(times)
		besttime = true
		savetimes()
	}

}
func restartgame() { //MARK: RESTART GAME

	for a := 0; a < len(level); a++ {
		level[a].etc = nil
		level[a].enemies = nil
		level[a].doorExitRecs = nil
		level[a].doorSides = nil
		level[a].floor = nil
		level[a].innerBloks = nil
		level[a].movBloks = nil
		level[a].nextRooms = nil
		level[a].spikes = nil
		level[a].visited = false
		level[a].walls = nil
	}

	kills = xkills{}
	inven = []xblok{}
	mods = xmod{}
	airstrikeT = 0
	airstrikeOn = false
	fx = nil
	plProj = nil
	enProj = nil
	flipcam = false
	cam2.Rotation = 0

	pl.armor = 0
	pl.coins = 0
	mods.armorN = 0
	pl.armorMax = 0
	pl.armor = 0
	floodRec.Y = scrHF32 + bsU

	pl.atk, pl.slide, pl.escape, pl.revived, pl.poison = false, false, false, false, false

	night = false

	level = nil
	shader2on = false
	shader3on = false

	levelnum = 1

	makeplayer()
	makelevel()

	pl.cnt = cnt
	roomNum = 0

	introcount = true
	introT3 = fps * 3

	intro = true
	startdmgT = fps * 7
	rl.PlaySound(sfx[13])

}
func savetimes() { //MARK: SAVE TIMES

	f, err := os.Create("etc/sc.000")
	if err != nil {
		fmt.Println(err)
		return
	}

	scoresTXT := ""
	for i := 0; i < len(times); i++ {
		if i == len(times)-1 {
			scoresTXT = scoresTXT + fmt.Sprint(times[i])
		} else {
			scoresTXT = scoresTXT + fmt.Sprint(times[i]) + ","
		}
	}

	_, err = f.WriteString(scoresTXT)
	if err != nil {
		fmt.Printf("Failed to write times: %s", err)
	}
}

func exitgame() { //MARK: EXIT GAME

	if optionsChange {
		savesettings()
	}
	savetimes()

	rl.EndDrawing()
	unload()
	rl.CloseAudioDevice()
	rl.CloseWindow()

}
func addshopitem(num int) { //MARK: ADD SHOP ITEM

	//CHECK FOR SAME ADD TO INVEN
	sold := false
	if len(inven) > 0 {
		foundsame := false
		for a := 0; a < len(inven); a++ {
			if shopitems[num].name == inven[a].name {

				switch shopitems[num].name {
				case "fireworks":
					pl.coins++
					txtSold("fireworks")
					sold = true
				case "peace":
					pl.coins++
					txtSold("peace")
					sold = true
				case "anchor":
					pl.coins++
					txtSold("anchor")
					sold = true
				case "recharge":
					pl.coins++
					txtSold("recharge")
					sold = true
				case "orbital":
					if mods.orbitalN < max.orbital {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("orbital")
						sold = true
					}
				case "coffee":
					if mods.coffeeN < max.coffee {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("coffee")
						sold = true
					}
				case "invisible":
					pl.coins++
					txtSold("invisible")
					sold = true
				case "health potion":
					if mods.hppotionN < max.hppotion {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("health potion")
						sold = true
					}
				case "firetrail":
					if mods.firetrailN < max.firetrail {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("firetrail")
						sold = true
					}
				case "map":
					pl.coins++
					txtSold("map")
					sold = true
				case "apple":
					if mods.appleN < max.apple {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("apple")
						sold = true
					}
				case "key":
					if mods.keyN < max.key {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("key")
						sold = true
					}
				case "bounce":
					if mods.bounceN < max.bounce {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("bounce")
						sold = true
					}
				case "fireball":
					if mods.fireballN < max.fireball {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("fireball")
						sold = true
					}
				case "throwing axe":
					if mods.axeN < max.axe {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("axe")
						sold = true
					}
				}
				foundsame = true
			}
		}
		if !foundsame {
			inven = append(inven, shopitems[num])
			rl.PlaySound(sfx[8])
		}
	} else {
		inven = append(inven, shopitems[num])
		rl.PlaySound(sfx[8])
	}

	//UP MODS
	if !sold {
		switch shopitems[num].name {
		case "fireworks":
			mods.fireworks = true
		case "peace":
			mods.peace = true
		case "anchor":
			mods.anchor = true
		case "recharge":
			mods.recharge = true
		case "orbital":
			mods.orbital = true
			if mods.orbitalN < max.orbital {
				mods.orbitalN++
				if mods.orbitalN == 1 {
					pl.orbital1 = rl.NewVector2(pl.cnt.X+bsU4, pl.cnt.Y+bsU4)
				}
				if mods.orbitalN == 2 {
					pl.orbital2 = rl.NewVector2(pl.cnt.X-bsU7, pl.cnt.Y-bsU7)
				}
			}
		case "coffee":
			if mods.coffeeN < max.coffee {
				mods.coffeeN++
				pl.vel++
			}
		case "invisible":
			mods.invisible = true
		case "health potion":
			mods.hppotion = true
			if mods.hppotionN < max.hppotion {
				mods.hppotionN++
			}
		case "firetrail":
			mods.firetrail = true
			if mods.firetrailN < max.firetrail {
				mods.firetrailN++
			}
		case "bounce":
			if mods.bounceN < max.bounce {
				mods.bounceN++
			}
		case "map":
			mods.exitmap = true
		case "apple":
			mods.apple = true
			if mods.appleN < max.apple {
				mods.appleN++
			}
		case "key":
			mods.key = true
			if mods.keyN < max.key {
				mods.keyN++
			}
		case "fireball":
			mods.fireball = true
			if mods.fireballN < max.fireball {
				mods.fireballN++
			}
		case "santa":
			mods.santa = true
			mods.santaT = rI32(7, 21) * fps
		case "throwing axe":
			mods.axe = true
			if mods.axeN < max.axe {
				mods.axeN++
				mods.axeT = (int32(max.axe) * fps) - (int32(mods.axeN) * fps)
			}
			if mods.axeT < fps {
				mods.axeT = fps
			}

		}
	}
}

func playenemyhit() { //MARK:PLAYER ENEMY HIT SOUND
	if flipcoin() {
		rl.PlaySound(sfx[1])
	} else {
		rl.PlaySound(sfx[2])
	}
}
func updownswitch(x32, y32 int32, siz, value float32, numType int) float32 { //MARK:UP DOWN SWITCH

	x := float32(x32)
	y := float32(y32)

	rec := rl.NewRectangle(x, y, siz*3, siz)

	rec2 := rec
	rec2.Width = rec2.Width / 3

	rec3 := rec2
	rec3.X += rec3.Width * 2
	rec3.Width -= 2
	rec3.X += 2
	rec2.Width -= 2

	rl.DrawRectangleRec(rec, rl.Black)
	rl.DrawRectangleRec(rec2, rl.Red)
	rl.DrawRectangleRec(rec3, rl.Green)

	rl.DrawRectangleLinesEx(rec, 1, rl.White)

	switch numType {
	case 2:
		txt := fmt.Sprint(value)
		txtlen := rl.MeasureText(txt, txtSize)

		rl.DrawText(txt, (rec.ToInt32().X+rec.ToInt32().Width/2)-txtlen/2, rec.ToInt32().Y+1, txtSize, rl.White)

	case 1:
		txt := fmt.Sprintf("%.0f", value*10)
		txtlen := rl.MeasureText(txt, txtSize)
		rl.DrawText(txt, (rec.ToInt32().X+rec.ToInt32().Width/2)-txtlen/2, rec.ToInt32().Y+1, txtSize, rl.White)
		upvolume()
	}

	return value
}

func destroyPowerupBlok(blokNum int) { //MARK:DESTROY POWERUP BLOK

	makeFX(1, level[roomNum].etc[blokNum].cnt)

	choose := rInt(1, 36)
	//choose = 6

	zblok := makeBlokGeneric(bsU+bsU/2, level[roomNum].etc[blokNum].cnt)
	zblok.onoff = true
	zblok.numof = 1

	switch choose {
	case 35: //MARIO
		zblok.name = "mario"
		zblok.desc = "collect extra coins"
		zblok.color = rl.White
		zblok.img = etc[51]
	case 34: //MR CARROT
		zblok.name = "mr carrot"
		zblok.desc = "your friend the root vegetable"
		zblok.color = rl.White
		zblok.img = etc[50]
	case 33: //FIREWORKS
		zblok.name = "fireworks"
		zblok.desc = "shoot fireworks when activating powerup block"
		zblok.color = rl.White
		zblok.img = etc[49]
	case 32: //AIR STRIKE
		zblok.name = "air strike"
		zblok.desc = "support from above at random intervals"
		zblok.color = rl.White
		zblok.img = etc[48]
	case 31: //ALIEN
		zblok.name = "mr alien"
		zblok.desc = "a strange companion"
		zblok.color = rl.White
		zblok.img = etc[47]
	case 30: //PEACE
		zblok.name = "peace"
		zblok.desc = "take no damage for 2 seconds on entering room"
		zblok.color = rl.White
		zblok.img = etc[46]
	case 29: //CAKE
		zblok.name = "birthday cake"
		zblok.desc = "get a random present"
		zblok.color = rl.White
		zblok.img = etc[45]
	case 28: //FLOOD
		zblok.name = "fish"
		zblok.desc = "a not quite biblical flood"
		zblok.color = ranCyan()
		zblok.img = etc[44]
	case 27: //CHERRY
		zblok.name = "cherry"
		zblok.desc = "coin jackpot - random coin amount added"
		zblok.color = rl.White
		zblok.img = etc[43]
	case 26: //SOCKS
		zblok.name = "moldy socks"
		zblok.desc = "leaves a trail of damaging footprints"
		zblok.color = rl.White
		zblok.img = etc[42]
	case 25: //UMBRELLA
		zblok.name = "umbrella"
		zblok.desc = "rain when you don't need it"
		zblok.color = rl.White
		zblok.img = etc[40]
	case 24: //ANCHOR
		zblok.name = "anchor"
		zblok.desc = "enemies pause for 2 seconds on entering room"
		zblok.color = rl.White
		zblok.img = etc[39]
	case 23: //RECHARGE
		zblok.name = "recharge"
		zblok.desc = "only works with armor - recharges 1 armor every 2 rooms"
		zblok.color = rl.White
		zblok.img = etc[12]
	case 22: //ARMOR
		zblok.name = "armor"
		zblok.desc = "protection that can be recharged"
		zblok.color = ranCyan()
		zblok.img = etc[38]
	case 21: //HP RING
		zblok.name = "health ring"
		zblok.desc = "adds another hp heart"
		zblok.color = rl.White
		zblok.img = etc[37]
	case 20: //CHAIN LIGHTNING
		zblok.name = "chain lightning"
		zblok.desc = "chance to damage all enemies on screen"
		zblok.color = rl.White
		zblok.img = etc[36]
	case 19: //ORBITAL
		zblok.name = "orbital"
		zblok.desc = "erratic revolving orbs that damage enemies"
		zblok.color = rl.White
		zblok.img = etc[35]
	case 18: //ATTACK DAMAGE
		zblok.name = "attack damage"
		zblok.desc = "increases damage of sword swing"
		zblok.color = rl.White
		zblok.img = etc[34]
	case 17: //ATTACK RANGE
		zblok.name = "attack range"
		zblok.desc = "increases range of sword swing"
		zblok.color = rl.White
		zblok.img = etc[33]
	case 16: //TELEPORT
		zblok.name = "teleport"
		zblok.desc = "transport to another room - destroyed on use"
		zblok.color = rl.White
		zblok.img = etc[32]
	case 15: //COFFEE
		zblok.name = "coffee"
		zblok.desc = "move faster - collect more = faster movement"
		zblok.color = rl.White
		zblok.img = etc[31]
	case 14: //INVISIBLE
		zblok.name = "invisible"
		zblok.desc = "enemies will not follow you"
		zblok.color = rl.White
		zblok.img = etc[30]
	case 13: //FIRE TRAIL
		zblok.name = "firetrail"
		zblok.desc = "trail of fire - does not effect flying enemies"
		zblok.color = ranOrange()
		zblok.img = etc[29]
	case 12: //FIREBALL
		zblok.name = "fireball"
		zblok.desc = "fires on attack - collect more = more fireballs"
		zblok.color = ranOrange()
		zblok.img = fireballPlayer.recTL
	case 11: //MAP
		zblok.name = "map"
		zblok.desc = "reveals location of exit room"
		zblok.color = ranGrey()
		zblok.img = etc[28]
	case 10: //WALLET
		zblok.name = "wallet"
		zblok.desc = "purchase one shop item free - destroyed on use"
		zblok.color = rl.Brown
		zblok.img = etc[11]
	case 9: //MEDI KIT
		zblok.name = "medi kit"
		zblok.desc = "resurrect from death - destroyed on use"
		zblok.color = rl.White
		zblok.img = etc[10]
	case 8: //PLANT COMPANION
		zblok.name = "mr planty"
		zblok.desc = "a companion to assist"
		zblok.color = rl.White
		zblok.img = etc[9]
	case 7: //APPLE
		zblok.name = "apple"
		zblok.desc = "prevents poisoning - destroyed on use"
		zblok.color = rl.White
		zblok.img = etc[8]
	case 6: //KEY
		zblok.name = "key"
		zblok.desc = "open locked chests"
		zblok.color = rl.White
		zblok.img = etc[7]
	case 5: //ESCAPE VINE
		zblok.name = "vine"
		zblok.desc = "automatically escape room at low hp - destroyed on use"
		zblok.color = rl.DarkGreen
		zblok.img = etc[6]
	case 4: //BOUNCE PROJECTILE
		zblok.name = "bounce"
		zblok.desc = "projectiles bounce > collect more = more bounces"
		zblok.color = rl.Yellow
		zblok.img = etc[5]
	case 3: //SANTA
		zblok.name = "santa"
		zblok.desc = "snow when you don't need it"
		zblok.color = rl.White
		zblok.img = etc[4]
	case 2: //THROWING AXE
		zblok.name = "throwing axe"
		zblok.desc = "fires at interval > collect more = faster fire rate"
		zblok.color = rl.SkyBlue
		zblok.img = etc[3]
	case 1: //HP POTION
		zblok.name = "health potion"
		zblok.desc = "automatically used when health < 2"
		zblok.color = rl.Red
		zblok.img = etc[1]

	}

	level[roomNum].etc = append(level[roomNum].etc, zblok)

}
func collectInven(blokNum int) { //MARK:COLLECT INVENTORY

	//CHECK FOR SAME ADD TO INVEN
	sold := false
	if len(inven) > 0 {
		foundsame := false
		for a := 0; a < len(inven); a++ {
			if level[roomNum].etc[blokNum].name == inven[a].name {

				switch level[roomNum].etc[blokNum].name {
				case "mr carrot":
					pl.coins++
					txtSold("mr carrot")
					sold = true
				case "fireworks":
					pl.coins++
					txtSold("fireworks")
					sold = true
				case "air strike":
					pl.coins++
					txtSold("air strike")
					sold = true
				case "mr alien":
					pl.coins++
					txtSold("mr alien")
					sold = true
				case "peace":
					pl.coins++
					txtSold("peace")
					sold = true
				case "birthday cake":
					if mods.cakeN < max.cake {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("birthday cake")
						sold = true
					}
				case "fish":
					pl.coins++
					txtSold("fish")
					sold = true
				case "cherry":
					if mods.cherryN < max.cherry {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("cherry")
						sold = true
					}
				case "socks":
					pl.coins++
					txtSold("socks")
					sold = true
				case "umbrella":
					pl.coins++
					txtSold("umbrella")
					sold = true
				case "anchor":
					pl.coins++
					txtSold("anchor")
					sold = true
				case "recharge":
					pl.coins++
					txtSold("recharge")
					sold = true
				case "armor":
					if mods.armorN < max.armor {
						inven[a].numof++
						if pl.armor < pl.armorMax {
							pl.armor++
						}
						rl.PlaySound(sfx[8])
					} else if mods.armorN == max.armor && pl.armor < pl.armorMax {
						pl.armor++
					} else {
						pl.coins++
						txtSold("armor")
						sold = true
					}
				case "health ring":
					if mods.hpringN < max.hpring {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("health ring")
						sold = true
					}
				case "chain lightning":
					pl.coins++
					txtSold("chain lightning")
					sold = true
				case "orbital":
					if mods.orbitalN < max.orbital {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("orbital")
						sold = true
					}
				case "attack damage":
					if mods.atkdmgN < max.atkdmg {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("attack damage")
						sold = true
					}
				case "attack range":
					if mods.atkrangeN < max.atkrange {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("attack range")
						sold = true
					}
				case "coffee":
					if mods.coffeeN < max.coffee {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("coffee")
						sold = true
					}
				case "invisible":
					pl.coins++
					txtSold("invisible")
					sold = true
				case "health potion":
					if mods.hppotionN < max.hppotion {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("health potion")
						sold = true
					}
				case "firetrail":
					if mods.firetrailN < max.firetrail {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("firetrail")
						sold = true
					}
				case "map":
					pl.coins++
					txtSold("map")
					sold = true
				case "wallet":
					pl.coins++
					txtSold("wallet")
					sold = true
				case "medi kit":
					pl.coins++
					txtSold("medi kit")
					sold = true
				case "mr planty":
					pl.coins++
					txtSold("mr planty")
					sold = true
				case "apple":
					if mods.appleN < max.apple {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("apple")
						sold = true
					}
				case "key":
					if mods.keyN < max.key {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("key")
						sold = true
					}
				case "vine":
					pl.coins++
					txtSold("vine")
					sold = true
				case "bounce":
					if mods.bounceN < max.bounce {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("bounce")
						sold = true
					}
				case "fireball":
					if mods.fireballN < max.fireball {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("fireball")
						sold = true
					}
				case "santa":
					pl.coins++
					txtSold("santa")
					sold = true
				case "throwing axe":
					if mods.axeN < max.axe {
						inven[a].numof++
						rl.PlaySound(sfx[8])
					} else {
						pl.coins++
						txtSold("axe")
						sold = true
					}
				}
				foundsame = true
			}
		}
		if !foundsame && level[roomNum].etc[blokNum].name != "teleport" {
			inven = append(inven, level[roomNum].etc[blokNum])
			rl.PlaySound(sfx[8])
		}
	} else {
		if level[roomNum].etc[blokNum].name != "teleport" {
			inven = append(inven, level[roomNum].etc[blokNum])
			rl.PlaySound(sfx[8])
		}
	}

	//UP MODS
	if !sold {
		switch level[roomNum].etc[blokNum].name {
		case "mario":
			makemario()
			pause = true
			marioon = true
		case "mr carrot":
			if !mods.alien && !mods.planty {
				mods.carrot = true
				mrcarrot.rec = rl.NewRectangle(pl.cnt.X-mrcarrot.rec.Width/2, pl.cnt.Y-mrcarrot.rec.Width/2, mrcarrot.rec.Width, mrcarrot.rec.Width)
			} else {
				txtCompanion()
			}
		case "fireworks":
			mods.fireworks = true
		case "air strike":
			mods.airstrike = true
			airstrikeT = fps * rI32(3, 8)
		case "mr alien":
			if !mods.carrot && !mods.planty {
				mods.alien = true
				mralien.rec = rl.NewRectangle(pl.cnt.X-mralien.rec.Width/2, pl.cnt.Y-mralien.rec.Width/2, mralien.rec.Width, mralien.rec.Width)
			} else {
				txtCompanion()
			}

		case "peace":
			mods.peace = true
		case "birthday cake":
			if mods.cakeN < max.cake {
				mods.cakeN++
				birthdaycake()
			}
		case "fish":
			mods.flood = true
			floodRec = rl.NewRectangle(0, scrHF32+bsU, scrWF32, scrHF32)
			makefish()
		case "cherry":
			if mods.cherryN < max.cherry {
				mods.cherryN++
				num2 := rInt(1, 6)
				pl.coins += num2
				txtAddCoinMulti()
			}
		case "moldy socks":
			mods.socks = true
		case "umbrella":
			mods.umbrella = true
		case "anchor":
			mods.anchor = true
		case "recharge":
			mods.recharge = true
		case "armor":
			if mods.armorN < max.armor {
				mods.armorN++
				pl.armorMax++
				pl.armor++
			}
		case "health ring":
			if mods.hpringN < max.hpring {
				mods.hpringN++
				pl.hpmax++
				if pl.hp < pl.hpmax {
					pl.hp++
				}
			}
		case "chain lightning":
			mods.chainlightning = true
		case "orbital":
			mods.orbital = true
			if mods.orbitalN < max.orbital {
				mods.orbitalN++
				if mods.orbitalN == 1 {
					pl.orbital1 = rl.NewVector2(pl.cnt.X+bsU4, pl.cnt.Y+bsU4)
				}
				if mods.orbitalN == 2 {
					pl.orbital2 = rl.NewVector2(pl.cnt.X-bsU7, pl.cnt.Y-bsU7)
				}
			}
		case "attack damage":
			if mods.atkdmgN < max.atkdmg {
				mods.atkdmgN++
			}
			pl.atkDMG++
		case "attack range":
			if mods.atkrangeN < max.atkrange {
				mods.atkrangeN++
			}
			pl.atkrec.X -= bsU
			pl.atkrec.Y -= bsU
			pl.atkrec.Width += bsU2
			pl.atkrec.Height += bsU2
		case "teleport":
			pl.hppause = fps * 5
			maketeleport()
			teleportRoomNum = rInt(0, len(level))
			teleporton = true
			rl.PlaySound(sfx[9])
		case "coffee":
			if mods.coffeeN < max.coffee {
				mods.coffeeN++
				pl.vel++
			}
		case "invisible":
			mods.invisible = true
		case "health potion":
			mods.hppotion = true
			if mods.hppotionN < max.hppotion {
				mods.hppotionN++
			}
		case "firetrail":
			mods.firetrail = true
			if mods.firetrailN < max.firetrail {
				mods.firetrailN++
			}
		case "bounce":
			if mods.bounceN < max.bounce {
				mods.bounceN++
			}
		case "map":
			mods.exitmap = true
		case "wallet":
			mods.wallet = true
		case "medi kit":
			mods.medikit = true
		case "mr planty":
			if !mods.alien && !mods.carrot {
				mods.planty = true
				mrplanty.rec = rl.NewRectangle(pl.cnt.X-mrplanty.rec.Width/2, pl.cnt.Y-mrplanty.rec.Width/2, mrplanty.rec.Width, mrplanty.rec.Width)
			} else {
				txtCompanion()
			}

		case "apple":
			mods.apple = true
			if mods.appleN < max.apple {
				mods.appleN++
			}
		case "key":
			mods.key = true
			if mods.keyN < max.key {
				mods.keyN++
			}
		case "vine":
			mods.vine = true
		case "fireball":
			mods.fireball = true
			if mods.fireballN < max.fireball {
				mods.fireballN++
			}
		case "santa":
			mods.santa = true
			mods.santaT = rI32(7, 21) * fps
		case "throwing axe":
			mods.axe = true
			if mods.axeN < max.axe {
				mods.axeN++
				mods.axeT = (int32(max.axe) * fps) - (int32(mods.axeN) * fps)
			}
			if mods.axeT < fps {
				mods.axeT = fps
			}

		}
	}

}
func birthdaycake() { //MARK:BIRTHDAY CAKE

	found := false
	for {

		choose := rInt(1, 6)
		switch choose {
		case 1:
			if pl.hp < pl.hpmax {
				pl.hp = pl.hpmax
				found = true
			}
		case 2:
			if pl.hpmax < 10 {
				pl.hpmax++
				found = true
			}
		case 3:
			if mods.armorN == 0 {
				if mods.armorN < max.armor {
					mods.armorN++
					pl.armorMax++
					pl.armor++
					found = true
				}
			}
		case 4:
			if mods.armorN > 0 && mods.armorN < max.armor {
				mods.armorN++
				pl.armorMax++
				pl.armor++
				found = true
			}
		case 5:
			if pl.armor > 0 && pl.armor < pl.armorMax {
				pl.armor = pl.armorMax
				found = true
			}
		case 6:
			num2 := rInt(1, 6)
			pl.coins += num2
			txtAddCoinMulti()
			found = true
		case 7:
			if pl.poison {
				pl.poisonT = 0
				pl.poisonCollisT = 0
				if pl.hp < pl.hpmax {
					pl.hp++
				}
			}
		}

		if found {
			break
		}

	}

}
func txtHere(txt string, rec rl.Rectangle) { //MARK:TEXT HERE

	ztxt := xtxt{}
	ztxt.txt = txt
	txtlen := rl.MeasureText(txt, 20)
	x := int32(rec.X+rec.Width/2) - txtlen/2
	y := int32(rec.Y - 20)
	ztxt.fade = 1
	ztxt.col = rl.White
	if txt == "poisoned" {
		ztxt.col = rl.Green
	}
	ztxt.x = x
	ztxt.y = y
	ztxt.onoff = true

	gametxt = append(gametxt, ztxt)

}
func delInven(name string) { //MARK:DEL INVEN
	clear := false
	for a := 0; a < len(inven); a++ {
		if name == inven[a].name {
			inven[a].numof--
			if inven[a].numof == 0 {
				clear = true
			}
		}
	}

	if clear {
		clearinven(name)
	}

}
func onoff(x32, y32 int32, siz float32, name bool) bool { //MARK: ON/OFF

	x := float32(x32)
	y := float32(y32)

	rec := rl.NewRectangle(x, y, siz*2, siz)
	rl.DrawRectangleRec(rec, rl.Black)
	rec2 := rec
	rec2.Width = rec2.Width / 2
	if name {
		rec2.X += siz
		rl.DrawRectangleRec(rec2, rl.Green)
	} else {
		rl.DrawRectangleRec(rec2, rl.Red)
	}

	rl.DrawRectangleLinesEx(rec, 1, rl.White)

	if rl.CheckCollisionPointRec(mousev2cam, rec) {

		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			name = !name
		}
	}

	return name

}
func clearinven(name string) { //MARK:CLEAR INVEN

	num := 0
	clear := false

	for a := 0; a < len(inven); a++ {
		if inven[a].name == name {
			num = a
			clear = true
		}
	}
	if clear {
		inven = remBlok(inven, num)
	}

}
func escapeplayer() { //MARK:ESCAPE PLAYER

	if pl.cnt.Y > levRecInner.Y && !escaped {
		pl.cnt.Y -= bsU / 4
		if pl.cnt.X > cnt.X {
			pl.cnt.X -= bsU / 4
		} else if pl.cnt.X < cnt.X {
			pl.cnt.X += bsU / 4
		}

		upPlayerRec()
		if pl.cnt.Y <= levRecInner.Y {
			escaped = true
			upRoomChange()
			pl.hp = 2
		}

	}

	if escaped {
		if !escapeRoomFound {
			choose := 0
			for {
				choose = rInt(0, len(level))
				if choose != roomNum {
					break
				}
			}
			roomNum = choose
			escapeRoomFound = true
		}
		if pl.cnt.Y < levRecInner.Y+bsU3 {
			pl.cnt.Y += bsU / 4
		}
		if pl.cnt.X > cnt.X {
			pl.cnt.X -= bsU / 4
		} else if pl.cnt.X < cnt.X {
			pl.cnt.X += bsU / 4
		}
		upPlayerRec()

		if pl.cnt.Y >= levRecInner.Y+bsU3 {
			pl.escape = false
		}
	}

}
func cleanlevel() { //MARK:CLEAN LEVEL

	//EMPTY ETC BLOKS
	for a := 0; a < len(level); a++ {

		var blokstoClear []int
		for b := 0; b < len(level[a].etc); b++ {
			if level[a].etc[b].name == "" {
				blokstoClear = append(blokstoClear, b)
			}
		}
		if len(blokstoClear) > 0 {
			for b := 0; b < len(blokstoClear); b++ {
				level[a].etc = remBlok(level[a].etc, blokstoClear[b])
			}
		}

	}

	//START ROOM CENTER BLOK CLEAR
	num := 0
	clear := false
	checkrec := rl.NewRectangle(cnt.X-bsU4, cnt.Y-bsU4, bsU8, bsU8)
	if len(level[0].innerBloks) > 0 {
		for a := 0; a < len(level[0].innerBloks); a++ {
			if rl.CheckCollisionRecs(checkrec, level[0].innerBloks[a].rec) {
				num = a
				clear = true
			}
		}
	}
	if clear {
		level[0].innerBloks = remBlok(level[0].innerBloks, num)
	}

	//CLEAR ETC CENTER
	for i := 0; i < len(level[0].etc); i++ {
		checkrec := rl.NewRectangle(cnt.X-bsU2, cnt.Y-bsU2, bsU4, bsU4)
		if rl.CheckCollisionRecs(checkrec, level[0].etc[i].rec) {
			level[0].etc[i].onoff = false
		}
	}

	//CLEAR MOVE BLOCKS CENTER
	for i := 0; i < len(level[0].movBloks); i++ {
		checkrec := rl.NewRectangle(cnt.X-bsU2, cnt.Y-bsU2, bsU4, bsU4)
		if rl.CheckCollisionRecs(checkrec, level[0].movBloks[i].rec) {
			level[0].movBloks[i].rec.X -= bsU4
		}
	}

}
func hitPL(numEnProj, numType int) { //MARK:HIT PLAYER

	if pl.hppause == 0 && pl.peaceT == 0 && startdmgT == 0 {
		pl.hppause = fps * 2
		hpHitY = 0
		hpHitF = 1

		if pl.armor > 0 {
			pl.armor--
			pl.armorHit = true
		} else {
			switch numType {
			case 2: // BURN POISON ENEMY COLLIS WATER
				if !invincible {
					pl.hp--

				}
			case 1: //ENEMY PROJECTILE
				if !invincible {
					pl.hp -= enProj[numEnProj].dmg

				}
			}
			zblok := xblok{}
			zblok.rec = pl.rec
			zblok.cnt = pl.cnt
			zblok.img = splats[rInt(0, len(splats))]
			zblok.color = ranRed()
			zblok.fade = 0.5
			zblok.onoff = true
			zblok.name = "playerblood"
			level[roomNum].etc = append(level[roomNum].etc, zblok)
			if pl.hp <= 0 && !invincible {
				pl.hp = 0
				if mods.medikit {
					pl.hp = pl.hpmax
					pl.revived = true
					pl.hppause = fps * 3
					reviveY = 0
					reviveF = 1
					mods.medikit = false
					clearinven("medi kit")
				} else {
					diedscrT = fps * 3
					pause = true
					died = true
					diedRec = rl.NewRectangle(cnt.X-bsU, cnt.Y-bsU, bsU2, bsU2)
					diedIMG = splats[rInt(0, len(splats))]
				}

			}
			rl.PlaySound(sfx[14])

		}
	}

}

func addkill(num int) { //MARK:ADD KILL
	switch level[roomNum].enemies[num].name {
	case "rabbit1":
		kills.bunnies++
	case "bat":
		kills.bats++
	case "mushroom":
		kills.mushrooms++
	case "ghost":
		kills.ghosts++
	case "spikehog":
		kills.spikehogs++
	case "rock":
		kills.rocks++
	case "slime":
		kills.slimes++
	}
	rl.PlaySound(sfx[7])
}
func txtSold(name string) { //MARK:TEXT SOLD
	rl.PlaySound(sfx[18])
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(levY+levW) - bsU5i32
	ztxt.x = int32(levX+levW) + bsUi32

	ztxt.txt = name + " max"
	ztxt.txt2 = "extra sold"

	txtSoldlist = append(txtSoldlist, ztxt)

}
func txtCompanion() { //MARK:TEXT COMPANION
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(levY+levW) - bsU5i32
	ztxt.x = int32(levX+levW) + bsUi32

	ztxt.txt = "1 companion"
	ztxt.txt2 = "allowed"

	txtSoldlist = append(txtSoldlist, ztxt)
}
func txtAddCoin() { //MARK:TEXT ADD 1 COIN
	rl.PlaySound(sfx[18])
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(levY+levW) - bsU5i32
	ztxt.x = int32(levX+levW) + bsUi32

	ztxt.txt = "+1 coin"

	txtSoldlist = append(txtSoldlist, ztxt)
	pl.coins++

}
func txtAddCoinMulti() { //MARK:TEXT ADD MULTIPLE COINS
	rl.PlaySound(sfx[18])
	ztxt := xtxt{}
	ztxt.onoff = true
	ztxt.col = rl.White
	ztxt.fade = 1
	ztxt.y = int32(levY+levW) - bsU5i32
	ztxt.x = int32(levX+levW) + bsUi32

	ztxt.txt = "jackpot"

	txtSoldlist = append(txtSoldlist, ztxt)
	pl.coins++

}

// MARK: UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP UP
func up() { //MARK:UP

	if !pause {
		uplevel()
		upplayer()
		timers()
	}

	checkcontroller()

	if keypressT > 0 {
		keypressT--
	}
	if optionT > 0 {
		optionT--
	}
	if startdmgT > 0 {
		startdmgT--
	}

	inp()
	mouseV2 = rl.GetMousePosition()
	mousev2cam = rl.GetScreenToWorld2D(mouseV2, cam2)

	if musicon {
		upaudio()
	}

	if fadeblinkon {
		if fadeblink < 0.5 {
			fadeblink += 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink > 0.1 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = true
		}
	}

	if fadeblinkon2 {
		if fadeblink2 < 0.3 {
			fadeblink2 += 0.03
		} else {
			fadeblinkon2 = false
		}
	} else {
		if fadeblink2 > 0.1 {
			fadeblink2 -= 0.03
		} else {
			fadeblinkon2 = true
		}
	}

}
func upaudio() { //MARK:UP AUDIO
	rl.UpdateMusicStream(music)
}
func upvolume() { //MARK:UP VOLUME

	rl.SetMusicVolume(music, volume)

	for a := 0; a < len(sfx); a++ {
		rl.SetSoundVolume(sfx[a], volume+0.1)
	}

	rl.SetMasterVolume(volume)

}
func upPlayerMods() { //MARK:UP PLAYER MODS

	//AIR STRIKE
	if mods.airstrike {
		airstrikeT--
		if airstrikeT <= 0 {
			airstrikeT = fps * rI32(3, 8)
			makeairstrike()
		}

	}

	//FLOOD
	if mods.flood {
		if marioon {
			floodRec.Y = scrHF32 + bsU
		} else if levelnum == 6 {
			marioon = false
		} else {
			levrecV2WorldtoScreen := rl.GetWorldToScreen2D(rl.NewVector2(levRecInner.X, levRecInner.Y), cam2)
			if floodRec.Y > levrecV2WorldtoScreen.Y+bsU12 {
				floodRec.Y -= 4
			} else {
				if fishLR {
					fishV2.X -= 5
					if fishV2.X < -fishSiz {
						fishLR = false
						fish1 = fishR.recTL
						fishV2.Y = rF32(scrHF32/3, scrHF32)
					}
				} else {
					fishV2.X += 7
					if fishV2.X > scrWF32 {
						fishLR = true
						fish1 = fishL.recTL
						fishV2.Y = rF32(scrHF32/3, scrHF32)
					}
				}
				if fish2LR {
					fish2V2.X += 4
					if fish2V2.X > scrWF32 {
						fish2LR = false
						fish2 = fishL.recTL
						fish2V2.Y = rF32(scrHF32/3, scrHF32)
					}
				} else {
					fish2V2.X -= 7
					if fish2V2.X <= 0 {
						fish2LR = true
						fish2 = fishR.recTL
						fish2V2.Y = rF32(scrHF32/3, scrHF32)
					}
				}
			}
		}
	}
	//SOCKS
	if mods.socks {
		sockstimer := 30
		if frames%sockstimer == 0 {
			zblok := makeBlokGenNoRecNoCntr()
			zblok.fade = 0.7
			zblok.color = ranGreen()
			zblok.cnt = pl.cnt
			siz := bsU2
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y, siz, siz)
			zblok.cnt = makeCnt(zblok)
			zblok.drec = makeDrec(zblok.rec)
			zblok.img = etc[52]
			zblok.name = "footprints"
			switch pl.direc {
			case 2:
				zblok.ro = 90
			case 3:
				zblok.ro = 180
			case 4:
				zblok.ro = 270
			}
			level[roomNum].etc = append(level[roomNum].etc, zblok)
		}
	}

	//FIRETRAIL
	if mods.firetrail {

		flametimer := 60
		if mods.firetrailN == 2 {
			flametimer = 30
		} else if mods.firetrailN == 3 {
			flametimer = 15
		}

		if frames%flametimer == 0 {
			zblok := makeBlokGenNoRecNoCntr()
			zblok.color = ranOrange()
			zblok.cnt = pl.cnt
			siz := bsU2
			zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y, siz, siz)
			zblok.img = firetrailanim.recTL
			zblok.name = "flamefiretrail"
			level[roomNum].etc = append(level[roomNum].etc, zblok)
		}

	}

	//HP POTION
	if mods.hppotion && pl.hppotionT == 0 {
		if pl.hp <= 2 {
			pl.hppotionT = fps
			pl.hp = pl.hpmax
			mods.hppotionN--
			pl.hppause = fps * 3
			if mods.hppotionN <= 0 {
				mods.hppotion = false
				clearinven("health potion")
			}
		}
	}

	//ESCAPE VINE
	if pl.hp == 1 && mods.vine && levelnum != 6 {
		rl.PlaySound(sfx[27])
		mods.vine = false
		clearinven("vine")
		escaped = false
		escapeRoomFound = false
		pl.escape = true
		pl.hppause = fps * 5
	}

	//SANTA
	if mods.santa {
		mods.santaT--
		if mods.santaT <= 0 {
			if mods.snowon {
				mods.snowon = false
			} else {
				makesnow()
				mods.snowon = true
			}
			mods.santaT = rI32(7, 21) * fps
		}
	}
	//AXE
	if mods.axe {
		mods.axeT2--
		if mods.axeT2 <= 0 {
			mods.axeT2 = mods.axeT
			makeProjectile("axe")
		}
	}

}

func upRoomChange() { //MARK:UP ROOM CHANGE

	plProj = nil
	enProj = nil
	fx = nil

	//FLOOD
	if mods.flood {
		pl.underWater = false
		floodRec = rl.NewRectangle(0, scrHF32+bsU, scrWF32, scrHF32)
		makefish()
	}
	//ARMOR RECHARGE
	if mods.recharge && mods.armorN > 0 {
		pl.rechargeN++
		if pl.rechargeN == 2 {
			pl.rechargeN = 0
			if pl.armor < pl.armorMax {
				pl.armor++
			}
		}
	}

	//PEACE PAUSE
	if mods.peace {
		pl.peaceT = fps * 2
	}
	//ANCHOR PAUSE
	if mods.anchor {
		anchorT = fps * 2
	}

}

func uplevel() { //MARK:UP LEVEL

	movebloks()

}
func upplayer() { //MARK:UP PLAYER

	//TIMERS
	if mods.flood {

		v2 := rl.GetScreenToWorld2D(rl.NewVector2(floodRec.Y, floodRec.Y), cam2)
		if pl.crec.Y > v2.Y {
			pl.underWater = true
			pl.waterT++
			if pl.waterT == fps*3 {
				hitPL(0, 2)
				pl.waterT = 0
			}
		} else {
			pl.waterT = 0
			pl.underWater = false
			waterY = 0
		}
	}
	if pl.peaceT > 0 {
		pl.peaceT--
	}
	if pl.hppotionT > 0 {
		pl.hppotionT--
	}
	if pl.hppause > 0 {
		pl.hppause--
		if pl.hppause == 1 && pl.revived {
			pl.revived = false
		}
		if pl.hppause == 1 && pl.armorHit {
			pl.armorHit = false
		}
	}
	if pl.poisonCollisT > 0 {
		pl.poisonCollisT--
	}

	if pl.poison {
		pl.poisonT--
		if pl.poisonT == 0 {
			hitPL(0, 2)
			pl.poisonCount--
			if pl.poisonCount == 0 {
				pl.poisonT = 0
				pl.poison = false
			} else {
				pl.poisonT = fps * 3
			}
		}
	}

	//ESCAPE
	if pl.escape {
		escapeplayer()
	}

	//SLIDE
	if pl.slide {
		switch pl.slideDIR {
		case 1:
			if checkplayermove(1) {
				pl.cnt.Y -= pl.vel * 2
			}
		case 2:
			if checkplayermove(2) {
				pl.cnt.X += pl.vel * 2
			}
		case 3:
			if checkplayermove(3) {
				pl.cnt.Y += pl.vel * 2
			}
		case 4:
			if checkplayermove(4) {
				pl.cnt.X -= pl.vel * 2
			}
		}
		upPlayerRec()
		pl.slideT--
		if pl.slideT <= 0 {
			pl.slide = false
		}
	}

	//UP MODS
	upPlayerMods()

	//UP IMG
	switch pl.direc {
	case 1:
		pl.img.Y = knight[1].Y
	case 2:
		pl.img.Y = knight[0].Y
	case 3:
		pl.img.Y = knight[3].Y
	case 4:
		pl.img.Y = knight[2].Y
	}

	if !pl.move && !pl.atk {
		pl.img.X = pl.imgWalkX
	} else if pl.move && !pl.atk {
		if frames%4 == 0 {
			pl.img.X += pl.sizImg
		}
		if pl.img.X > pl.imgWalkX+(float32(pl.framesWalk-1)*pl.sizImg) {
			pl.img.X = pl.imgWalkX
		}
	} else if !pl.move && pl.atk {
		if frames%4 == 0 {
			pl.img.X += pl.sizImg
		}
		if pl.img.X > pl.imgAtkX+(float32(pl.framesAtk-1)*pl.sizImg) {
			pl.img.X = pl.imgAtkX
		}
	}

	//FIND NEXT ROOM ON MOVEMENT
	if !roomChanged {
		if !rl.CheckCollisionPointRec(pl.cnt, levRec) {
			roomChanged = true
			roomChangedTimer = fps / 2
			upRoomChange()
			cntCompanion := pl.cnt
			if pl.cnt.X <= levX {
				for a := 0; a < len(level[roomNum].doorSides); a++ {
					if level[roomNum].doorSides[a] == 4 {
						roomNum = level[roomNum].nextRooms[a]
						level[roomNum].visited = true
						break
					}
				}
				pl.cnt.X = levRecInner.X + levRecInner.Width - bsU
				cntCompanion = pl.cnt
				cntCompanion.X -= bsU2
			} else if pl.cnt.X >= levX+levW {
				for a := 0; a < len(level[roomNum].doorSides); a++ {
					if level[roomNum].doorSides[a] == 2 {
						roomNum = level[roomNum].nextRooms[a]
						level[roomNum].visited = true
						break
					}
				}
				pl.cnt.X = levRecInner.X + bsU
				cntCompanion = pl.cnt
				cntCompanion.X += bsU2
			} else if pl.cnt.Y <= levY {
				for a := 0; a < len(level[roomNum].doorSides); a++ {
					if level[roomNum].doorSides[a] == 1 {
						roomNum = level[roomNum].nextRooms[a]
						level[roomNum].visited = true
						break
					}
				}
				pl.cnt.Y = levRecInner.Y + levRecInner.Width - bsU
				cntCompanion = pl.cnt
				cntCompanion.Y -= bsU2
			} else if pl.cnt.Y >= levY+levW {
				for a := 0; a < len(level[roomNum].doorSides); a++ {
					if level[roomNum].doorSides[a] == 3 {
						roomNum = level[roomNum].nextRooms[a]
						level[roomNum].visited = true
						break
					}
				}
				pl.cnt.Y = levRecInner.Y + bsU
				cntCompanion = pl.cnt
				cntCompanion.Y += bsU2
			}

			if mods.carrot {
				mrcarrot.rec = rl.NewRectangle(cntCompanion.X-mrcarrot.rec.Width/2, cntCompanion.Y-mrcarrot.rec.Width/2, mrcarrot.rec.Width, mrcarrot.rec.Width)
			}
			if mods.alien {
				mralien.rec = rl.NewRectangle(cntCompanion.X-mralien.rec.Width/2, cntCompanion.Y-mralien.rec.Width/2, mralien.rec.Width, mralien.rec.Width)
			}
			if mods.planty {
				mrplanty.rec = rl.NewRectangle(cntCompanion.X-mrplanty.rec.Width/2, cntCompanion.Y-mrplanty.rec.Width/2, mrplanty.rec.Width, mrplanty.rec.Width)
			}
		}
	}

	//TIMERS
	if pl.atkTimer > 0 {
		pl.atkTimer--
		if pl.atkTimer == 1 {
			pl.img.X = pl.imgWalkX
			pl.atk = false
		}
	}
}
func upPlayerRec() { //MARK:UP PLAYER REC CENTER CHANGED
	pl.rec = rl.NewRectangle(pl.cnt.X-pl.rec.Width/2, pl.cnt.Y-pl.rec.Width/2, pl.rec.Width, pl.rec.Width)
	pl.crec = rl.NewRectangle(pl.cnt.X-pl.crec.Width/2, pl.cnt.Y-pl.crec.Height/2, pl.crec.Width, pl.crec.Height)
	pl.arec = rl.NewRectangle(pl.cnt.X-pl.arec.Width/2, pl.cnt.Y-pl.arec.Height/2, pl.arec.Width, pl.arec.Height)
	pl.atkrec = rl.NewRectangle(pl.cnt.X-pl.atkrec.Width/2, pl.cnt.Y-pl.atkrec.Height/2, pl.atkrec.Width, pl.atkrec.Height)
}

// MARK:MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE MAKE
func makeaudio() { //MARK:MAKE AUDIO

	backMusic = append(backMusic, rl.LoadMusicStream("audio/1.ogg"))  //0
	backMusic = append(backMusic, rl.LoadMusicStream("audio/2.ogg"))  //1
	backMusic = append(backMusic, rl.LoadMusicStream("audio/16.ogg")) //2

	sfx = append(sfx, rl.LoadSound("audio/3.ogg"))  //0 PLAYER SWORD
	sfx = append(sfx, rl.LoadSound("audio/4.ogg"))  //1 ENEMY HIT1
	sfx = append(sfx, rl.LoadSound("audio/5.ogg"))  //2 ENEMY HIT2
	sfx = append(sfx, rl.LoadSound("audio/6.ogg"))  //3 SPRING
	sfx = append(sfx, rl.LoadSound("audio/7.ogg"))  //4 SPEAR
	sfx = append(sfx, rl.LoadSound("audio/8.ogg"))  //5 OIL BARREL BURN
	sfx = append(sfx, rl.LoadSound("audio/9.ogg"))  //6 POWER UP BLOCK DESTROY
	sfx = append(sfx, rl.LoadSound("audio/10.ogg")) //7 ENEMY DEATH
	sfx = append(sfx, rl.LoadSound("audio/11.ogg")) //8 ITEM PICKUP
	sfx = append(sfx, rl.LoadSound("audio/12.ogg")) //9 TELEPORT
	sfx = append(sfx, rl.LoadSound("audio/13.ogg")) //10 OIL BARREL EXPLODE
	sfx = append(sfx, rl.LoadSound("audio/14.ogg")) //11 LIGHTNING
	sfx = append(sfx, rl.LoadSound("audio/15.ogg")) //12 SWITCH
	sfx = append(sfx, rl.LoadSound("audio/17.ogg")) //13 COUNTDOWN
	sfx = append(sfx, rl.LoadSound("audio/18.ogg")) //14 PLAYER HIT
	sfx = append(sfx, rl.LoadSound("audio/19.ogg")) //15 STAIRS
	sfx = append(sfx, rl.LoadSound("audio/20.ogg")) //16 SHOP DOOR
	sfx = append(sfx, rl.LoadSound("audio/21.ogg")) //17 CHEST OPENING
	sfx = append(sfx, rl.LoadSound("audio/22.ogg")) //18 COIN ADDED
	sfx = append(sfx, rl.LoadSound("audio/23.ogg")) //19 WIN GAME
	sfx = append(sfx, rl.LoadSound("audio/24.ogg")) //20 GAS TRAP
	sfx = append(sfx, rl.LoadSound("audio/25.ogg")) //21 SHOP PURCHASE
	sfx = append(sfx, rl.LoadSound("audio/26.ogg")) //22d NO MONEY SHOP PURCHASE
	sfx = append(sfx, rl.LoadSound("audio/27.ogg")) //23 POISONED
	sfx = append(sfx, rl.LoadSound("audio/28.ogg")) //24 UNDERWATER
	sfx = append(sfx, rl.LoadSound("audio/29.ogg")) //25 LOCKED CHEST
	sfx = append(sfx, rl.LoadSound("audio/30.ogg")) //26 SPEAR
	sfx = append(sfx, rl.LoadSound("audio/31.ogg")) //27 VINE
	sfx = append(sfx, rl.LoadSound("audio/32.ogg")) //28 AIR STRIKE EPLOSION
	sfx = append(sfx, rl.LoadSound("audio/33.ogg")) //29 BOSS HIT
	sfx = append(sfx, rl.LoadSound("audio/34.ogg")) //30 CIRCULAR SAW

	upvolume()

}
func makechestitem(num int) { //MARK:MAKE CHEST ITEM

	newcnt := level[roomNum].etc[num].cnt
	newcnt.Y -= bsU2
	zblok := makeBlokGeneric(bsU+bsU/2, newcnt)
	choose := rInt(0, len(gems))
	zblok.onoff = true
	zblok.numof = 1
	zblok.name = "gem"
	zblok.color = rl.White
	zblok.img = gems[choose]
	switch choose {
	case 0:
		zblok.numCoins = 5
	case 1:
		zblok.numCoins = 6
	case 2:
		zblok.numCoins = 7
	case 3:
		zblok.numCoins = 8
	case 4:
		zblok.numCoins = 9
	case 5:
		zblok.numCoins = 10
	case 6:
		zblok.numCoins = 11
	case 7:
		zblok.numCoins = 12
	}

	level[roomNum].etc = append(level[roomNum].etc, zblok)
}
func maketimes() { //MARK: MAKE TIMES

	contents, err := os.ReadFile("etc/sc.000")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	txtTimes := strings.Split(string(contents), ",")

	for i := 0; i < len(txtTimes); i++ {
		num, _ := strconv.Atoi(txtTimes[i])
		times = append(times, num)
		sort.Ints(times)
	}

}
func makesettings() { //MARK: MAKE SETTINGS

	contents, err := os.ReadFile("etc/st.000")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	txtSettings := strings.Split(string(contents), ",")

	if txtSettings[0] == "1" {
		hpBarsOn = true
	} else {
		hpBarsOn = false
	}
	if txtSettings[1] == "1" {
		scanlineson = true
	} else {
		scanlineson = false
	}
	if txtSettings[2] == "1" {
		artifactson = true
	} else {
		artifactson = false
	}
	if txtSettings[3] == "1" {
		shaderon = true
	} else {
		shaderon = false
	}
	if txtSettings[4] == "1" {
		platkrecon = true
	} else {
		platkrecon = false
	}
	if txtSettings[5] == "1" {
		invincible = true
	} else {
		invincible = false
	}
	if txtSettings[6] == "1" {
		useController = true
		contolleron = true
	} else {
		useController = true
		contolleron = true
	}
	if txtSettings[7] == "1" {
		musicon = true
	} else {
		musicon = false
	}
	if txtSettings[10] == "1" {
		hardcore = true
	} else {
		hardcore = false
	}

	//BG MUSIC
	num, _ := strconv.Atoi(txtSettings[8])
	bgMusicNum = num
	//VOLUME
	num, _ = strconv.Atoi(txtSettings[9])
	volume = float32(num) / 10

}
func makeshop() { //MARK: MAKE SHOP

	shopitems = nil
	shopnum = 0
	countbreak := 100

	//SHOP IMG
	for {
		shopRoomNum = rInt(0, len(level))

		zblok := xblok{}
		zblok.name = "shop"
		zblok.solid = true
		zblok.onoff = true
		zblok.img = etc[14]
		zblok.cnt = cnt
		zblok.cnt.Y -= bsU8

		siz := bsU8

		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(siz, bsU4, shopRoomNum)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.crec = zblok.rec
		zblok.crec.X += zblok.crec.Width / 8
		zblok.crec.Width = (zblok.crec.Width / 4) * 3
		zblok.crec2 = zblok.crec
		zblok.crec2.Width += bsU
		zblok.crec2.Height += bsU
		zblok.crec2.X -= bsU / 2
		zblok.crec2.Y -= bsU / 2

		zblok.color = rl.White
		zblok.fade = 1

		siz = bsU / 2
		v1 := rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Height+bsU)
		v2 := v1
		v2.Y += siz
		v2.X += siz / 2
		v3 := v2
		v3.X -= siz

		zblok.v2s = append(zblok.v2s, v1)
		zblok.v2s = append(zblok.v2s, v2)
		zblok.v2s = append(zblok.v2s, v3)
		zblok.v2s = append(zblok.v2s, v1)
		zblok.v2s = append(zblok.v2s, v2)
		zblok.v2s = append(zblok.v2s, v3)

		rl.DrawTriangle(v2, v1, v3, rl.Red)

		countbreak--
		if canadd || countbreak == 0 {
			level[shopRoomNum].etc = append(level[shopRoomNum].etc, zblok)
			break
		}
	}

	//SHOP ITEMS
	zblok := xblok{}
	zblok.fade = 1
	zblok.onoff = true
	num := 4

	for num > 0 {

		choose := rInt(1, 16)

		switch choose {
		case 15: //FIREWORKS
			zblok.name = "fireworks"
			zblok.desc = "shoot fireworks when activating powerup block"
			zblok.color = rl.White
			zblok.img = etc[49]
		case 14: //PEACE
			zblok.name = "peace"
			zblok.desc = "take no damage for 2 seconds on entering room"
			zblok.color = rl.White
			zblok.img = etc[46]
		case 13: //ANCHOR
			zblok.name = "anchor"
			zblok.desc = "enemies pause for 2 seconds on entering room"
			zblok.color = rl.White
			zblok.img = etc[39]
		case 12: //RECHARGE
			zblok.name = "recharge"
			zblok.desc = "only works with armor - recharges 1 armor every 2 rooms"
			zblok.color = rl.White
			zblok.img = etc[12]
		case 11: //ORBITAL
			zblok.name = "orbital"
			zblok.desc = "erratic revolving orbs that damage enemies"
			zblok.color = rl.White
			zblok.img = etc[35]
		case 10: //COFFEE
			zblok.name = "coffee"
			zblok.desc = "move faster - collect more = faster movement"
			zblok.color = rl.White
			zblok.img = etc[31]
		case 9: //INVISIBLE
			zblok.name = "invisible"
			zblok.desc = "enemies will not follow you"
			zblok.color = rl.White
			zblok.img = etc[30]
		case 8: //FIRE TRAIL
			zblok.name = "firetrail"
			zblok.desc = "trail of fire - does not effect flying enemies"
			zblok.color = ranOrange()
			zblok.img = etc[29]
		case 7: //FIREBALL
			zblok.name = "fireball"
			zblok.desc = "fires on attack - collect more = more fireballs"
			zblok.color = ranOrange()
			zblok.img = fireballPlayer.recTL
		case 6: //MAP
			zblok.name = "map"
			zblok.desc = "reveals location of exit room"
			zblok.color = ranGrey()
			zblok.img = etc[28]
		case 5: //APPLE
			zblok.name = "apple"
			zblok.desc = "prevents poisoning - destroyed on use"
			zblok.color = rl.White
			zblok.img = etc[8]
		case 4: //SKELETON KEY
			zblok.name = "key"
			zblok.desc = "open locked things"
			zblok.color = rl.White
			zblok.img = etc[7]
		case 3: //BOUNCE PROJECTILE
			zblok.name = "bounce"
			zblok.desc = "projectiles bounce > collect more = more bounces"
			zblok.color = rl.Yellow
			zblok.img = etc[5]
		case 2: //THROWING AXE
			zblok.name = "throwing axe"
			zblok.desc = "fires at interval > collect more = faster fire rate"
			zblok.color = rl.SkyBlue
			zblok.img = etc[3]
		case 1: //HP POTION
			zblok.name = "health potion"
			zblok.desc = "automatically used when health < 2"
			zblok.color = rl.Red
			zblok.img = etc[1]
		}

		if len(shopitems) > 0 {
			canadd := true
			for a := 0; a < len(shopitems); a++ {
				if shopitems[a].name == zblok.name {
					canadd = false
				}
			}
			if canadd {
				zblok.shopprice = rInt(2, 8)
				shopitems = append(shopitems, zblok)
				num--
			}
		} else {
			zblok.shopprice = rInt(2, 8)
			shopitems = append(shopitems, zblok)
			num--
		}

	}

}

func makeProjectileEnemy(num int, cnt rl.Vector2) { //MARK:MAKE PROJECTILE ENEMY

	switch num {

	case 1: //TURRET

		siz := bsU
		zproj := xproj{}
		zproj.name = "ninja"
		zproj.cnt = cnt
		//zproj.cnt.X -= bsU / 2
		//zproj.cnt.Y -= bsU / 2
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.vel = bsU / 4
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = ranCyan()
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = etc[19]
		zproj.onoff = true

		enProj = append(enProj, zproj)

	case 9: //BOSS ATK 3
		siz := bsU3
		zproj := xproj{}
		zproj.name = "boss3"
		zproj.cnt = cnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.vel = bsU / 2
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = rl.White
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = etc[57]
		zproj.onoff = true
		enProj = append(enProj, zproj)

	case 8: //BOSS ATK 2
		siz := bsU3
		zproj := xproj{}
		zproj.name = "boss2"
		zproj.cnt = cnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.vel = bsU / 2
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = rl.White
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = boss2anim.recTL
		zproj.onoff = true
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)

	case 7: //BOSS ATK 1
		siz := bsU3
		zproj := xproj{}
		zproj.name = "boss1"
		zproj.cnt = cnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.crec = zproj.rec
		zproj.crec.X += zproj.rec.Width / 4
		zproj.crec.Y += zproj.rec.Height / 4
		zproj.crec.Width = zproj.rec.Width / 2
		zproj.crec.Height = zproj.rec.Height / 2
		zproj.vel = bsU / 2
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		zproj.col = rl.White
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = boss1anim.recTL
		zproj.onoff = true
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)
		zproj.velx = rF32(-zproj.vel, zproj.vel)
		zproj.vely = rF32(-zproj.vel, zproj.vel)
		enProj = append(enProj, zproj)

	case 2, 3, 4, 5: //SWITCH ARROWS

		siz := bsU2
		zproj := xproj{}
		zproj.name = "switcharrow"
		zproj.cnt = cnt
		//zproj.cnt.X -= bsU / 2
		//zproj.cnt.Y -= bsU / 2
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.vel = bsU / 2
		if num == 2 {
			zproj.vely = zproj.vel
			zproj.ro = 135
		} else if num == 3 {
			zproj.velx = zproj.vel
			zproj.ro = 45
		} else if num == 4 {
			zproj.vely = -zproj.vel
			zproj.ro = -45
		} else if num == 5 {
			zproj.velx = -zproj.vel
			zproj.ro = -135
		}
		zproj.col = ranOrange()
		zproj.dmg = 1
		zproj.fade = 1
		zproj.img = etc[24]
		zproj.onoff = true

		enProj = append(enProj, zproj)

	}

}
func makeendlevel() { //MARK: MAKE END LEVEL

	level = nil

	roomNum = 0

	//MAKE LEVEL ROOMS

	bossnum = rInt(0, len(bosses))

	countedBorderBlocks := false

	floorT := floortiles[rInt(0, len(floortiles))]
	wallT = walltiles[rInt(0, len(walltiles))]

	zroom := xroom{}
	zroom.floorT = floorT
	zroom.wallT = wallT

	//BOUNDARY WALLS
	x := levX
	y := levY

	for x < levX+levW {
		if !countedBorderBlocks {
			levBorderBlokNum++
		}
		zblok := xblok{}
		zblok.fade = 1
		zblok.img = zroom.wallT
		switch levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zblok.rec = rl.NewRectangle(x, y, borderWallBlokSiz, borderWallBlokSiz)
		zroom.walls = append(zroom.walls, zblok)
		zblok.rec.Y = levY + levW - borderWallBlokSiz
		switch levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zroom.walls = append(zroom.walls, zblok)
		x += borderWallBlokSiz
	}
	if !countedBorderBlocks {
		countedBorderBlocks = true
	}
	x = levX
	y = levY + borderWallBlokSiz
	for y < levY+levW-borderWallBlokSiz {
		zblok := xblok{}
		zblok.fade = 1
		zblok.img = zroom.wallT
		switch levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zblok.rec = rl.NewRectangle(x, y, borderWallBlokSiz, borderWallBlokSiz)
		zroom.walls = append(zroom.walls, zblok)
		zblok.rec.X = levX + levW - borderWallBlokSiz
		switch levelnum {
		case 1:
			zblok.color = ranBlue()
		case 2:
			zblok.color = ranBrown()
		case 3:
			zblok.color = ranOrange()
		case 4:
			zblok.color = ranDarkBlue()
		case 5:
			zblok.color = ranCol()
		case 6:
			zblok.color = ranRed()
		}
		zroom.walls = append(zroom.walls, zblok)
		y += borderWallBlokSiz
	}

	//FLOOR
	x = levX
	y = levY
	siz := bsU3
	zblok := xblok{}
	zblok.img = zroom.floorT
	for {
		zblok.rec = rl.NewRectangle(x, y, siz, siz)
		zblok.fade = rF32(0.1, 0.25)
		switch levelnum {
		case 1:
			zblok.color = ranRed()
		case 2:
			zblok.color = ranDarkGreen()
		case 3:
			zblok.color = ranDarkBlue()
		case 4:
			zblok.color = ranDarkGrey()
		case 5:
			zblok.color = ranDarkBlue()
		case 6:
			zblok.color = ranRed()
		}
		zroom.floor = append(zroom.floor, zblok)

		x += siz
		if x >= levX+levW {
			x = levX
			y += siz
		}
		if y >= levY+levW {
			break
		}
	}

	level = append(level, zroom)

	//SWITCHES
	if flipcoin() {
		zblok := makeBlokGenNoRecNoCntr()
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.name = "switch"
		zblok.numType = roll6()
		if zblok.numType == 6 {
			zblok.numCoins = rInt(3, 11)
		}
		zblok.onoffswitch = flipcoin()
		zblok.color = rl.SkyBlue
		if zblok.onoffswitch {
			zblok.img = etc[21]
		} else {
			zblok.img = etc[22]
		}
		if canadd {
			level[0].etc = append(level[0].etc, zblok)
		}
	}
	if flipcoin() {
		zblok := makeBlokGenNoRecNoCntr()
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.name = "switch"
		zblok.numType = roll6()
		if zblok.numType == 6 {
			zblok.numCoins = rInt(3, 11)
		}
		zblok.onoffswitch = flipcoin()
		zblok.color = rl.SkyBlue
		if zblok.onoffswitch {
			zblok.img = etc[21]
		} else {
			zblok.img = etc[22]
		}
		if canadd {
			level[0].etc = append(level[0].etc, zblok)
		}
	}

	//SKULLS
	num := rInt(1, 5)
	for {
		zblok := makeBlokGenNoRecNoCntr()
		siz := rF32((bsU/4)*3, bsU+bsU/2)
		canadd := true
		zblok.rec, canadd = findRecPos(siz, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.name = "skull"
		zblok.fade = rF32(0.3, 0.6)
		zblok.color = ranGrey()
		zblok.img = skulls[rInt(0, len(skulls))]
		if canadd {
			level[0].etc = append(level[0].etc, zblok)
			num--
		}
		if num <= 0 {
			break
		}
	}

	//OIL BARRELS
	num = rInt(2, 7)
	for {
		zblok := makeBlokGenNoRecNoCntr()
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, 0)
		zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
		zblok.solid = true
		zblok.onoff = true
		zblok.name = "oilbarrel"
		zblok.color = rl.DarkGreen
		zblok.img = etc[20]
		if canadd {
			level[0].etc = append(level[0].etc, zblok)
			num--
		}

		if num <= 0 {
			break
		}
	}

	makeInnerBloks()
	makemovebloks()
	makeblades()
	maketurrets()
	makespears()
	makespikes()

	//REMOVE BLOKS COLLIDING WITH BOSS REC
	checkrec := rl.NewRectangle(cnt.X-bsU4, levRecInner.Y+bsU2, bsU8, bsU8)
	numcollis := 0
	found := false
	for i := 0; i < len(level[0].innerBloks); i++ {
		if rl.CheckCollisionRecs(checkrec, level[0].innerBloks[i].rec) {
			numcollis = i
			found = true
		}
	}
	if found {
		level[0].innerBloks = remBlok(level[0].innerBloks, numcollis)
	}

	cleanlevel()

}
func makelevel() { //MARK:MAKE LEVEL

	level = nil

	//MAKE LEVEL ROOMS
	numRooms := rInt(7, 13)
	orignumRooms := numRooms
	countedBorderBlocks := false

	floorT := floortiles[rInt(0, len(floortiles))]
	wallT = walltiles[rInt(0, len(walltiles))]

	for numRooms > 0 {
		zroom := xroom{}
		zroom.floorT = floorT
		zroom.wallT = wallT

		//BOUNDARY WALLS
		x := levX
		y := levY

		for x < levX+levW {
			if !countedBorderBlocks {
				levBorderBlokNum++
			}
			zblok := xblok{}
			zblok.fade = 1
			zblok.img = zroom.wallT
			switch levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zblok.rec = rl.NewRectangle(x, y, borderWallBlokSiz, borderWallBlokSiz)
			zroom.walls = append(zroom.walls, zblok)
			zblok.rec.Y = levY + levW - borderWallBlokSiz
			switch levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zroom.walls = append(zroom.walls, zblok)
			x += borderWallBlokSiz
		}
		if !countedBorderBlocks {
			countedBorderBlocks = true
		}
		x = levX
		y = levY + borderWallBlokSiz
		for y < levY+levW-borderWallBlokSiz {
			zblok := xblok{}
			zblok.fade = 1
			zblok.img = zroom.wallT
			switch levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zblok.rec = rl.NewRectangle(x, y, borderWallBlokSiz, borderWallBlokSiz)
			zroom.walls = append(zroom.walls, zblok)
			zblok.rec.X = levX + levW - borderWallBlokSiz
			switch levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zroom.walls = append(zroom.walls, zblok)
			y += borderWallBlokSiz
		}

		//FLOOR
		x = levX
		y = levY
		siz := bsU3
		zblok := xblok{}
		zblok.img = zroom.floorT
		for {
			zblok.rec = rl.NewRectangle(x, y, siz, siz)
			zblok.fade = rF32(0.1, 0.25)
			switch levelnum {
			case 1:
				zblok.color = ranRed()
			case 2:
				zblok.color = ranDarkGreen()
			case 3:
				zblok.color = ranDarkBlue()
			case 4:
				zblok.color = ranDarkGrey()
			case 5:
				zblok.color = ranDarkBlue()
			}
			zroom.floor = append(zroom.floor, zblok)

			x += siz
			if x >= levX+levW {
				x = levX
				y += siz
			}
			if y >= levY+levW {
				break
			}
		}

		level = append(level, zroom)

		numRooms--
	}
	numRooms = orignumRooms

	//MAKE LEVEL MAP
	levMap = nil
	mapRecSize := float32(96)
	rec := rl.NewRectangle(cnt.X-mapRecSize/2, cnt.Y-mapRecSize/2, mapRecSize, mapRecSize)
	levMap = append(levMap, rec)
	numRooms--
	countbreak := 0
	for numRooms > 0 {

		choose := levMap[0]
		if len(levMap) > 1 {
			choose = levMap[rInt(0, len(levMap))]
		}

		side := rInt(1, 5)
		checkV2 := rl.NewVector2(choose.X+choose.Width/2, choose.Y+choose.Width/2)
		switch side {
		case 1:
			checkV2.Y -= mapRecSize
		case 2:
			checkV2.X += mapRecSize
		case 3:
			checkV2.Y += mapRecSize
		case 4:
			checkV2.X -= mapRecSize
		}

		canadd := true
		for a := 0; a < len(levMap); a++ {
			if rl.CheckCollisionPointRec(checkV2, levMap[a]) {
				canadd = false
			}
		}

		if canadd {
			switch side {
			case 1:
				rec = choose
				rec.Y -= mapRecSize
			case 2:
				rec = choose
				rec.X += mapRecSize
			case 3:
				rec = choose
				rec.Y += mapRecSize
			case 4:
				rec = choose
				rec.X -= mapRecSize
			}

			levMap = append(levMap, rec)
			numRooms--
		} else {
			countbreak++
		}

		if countbreak == 1000 {
			break
		}
	}

	//FIND DOORS
	for a := 0; a < len(levMap); a++ {

		checkV2 := rl.NewVector2(levMap[a].X+levMap[a].Width/2, levMap[a].Y+levMap[a].Width/2)
		ocheckV2 := checkV2
		checkV2.Y -= levMap[a].Width
		for b := 0; b < len(levMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, levMap[b]) {
					level[a].doorSides = append(level[a].doorSides, 1)
					level[a].nextRooms = append(level[a].nextRooms, b)
				}
			}
		}
		checkV2 = ocheckV2
		checkV2.X += levMap[a].Width
		for b := 0; b < len(levMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, levMap[b]) {
					level[a].doorSides = append(level[a].doorSides, 2)
					level[a].nextRooms = append(level[a].nextRooms, b)
				}
			}
		}
		checkV2 = ocheckV2
		checkV2.Y += levMap[a].Width
		for b := 0; b < len(levMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, levMap[b]) {
					level[a].doorSides = append(level[a].doorSides, 3)
					level[a].nextRooms = append(level[a].nextRooms, b)
				}
			}
		}
		checkV2 = ocheckV2
		checkV2.X -= levMap[a].Width
		for b := 0; b < len(levMap); b++ {
			if a != b {
				if rl.CheckCollisionPointRec(checkV2, levMap[b]) {
					level[a].doorSides = append(level[a].doorSides, 4)
					level[a].nextRooms = append(level[a].nextRooms, b)
				}
			}
		}

	}

	//REMOVE DOOR BLOKS
	for a := 0; a < len(level); a++ {

		for b := 0; b < len(level[a].doorSides); b++ {

			if level[a].doorSides[b] == 1 {
				checkV2 := rl.NewVector2(levX+levW/2, levY+borderWallBlokSiz/2)
				checkV2L1 := checkV2
				checkV2L1.X -= borderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.X -= borderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.X += borderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.X += borderWallBlokSiz

				exitRec := rl.NewRectangle(checkV2L2.X-borderWallBlokSiz/2, checkV2L2.Y-borderWallBlokSiz/2, borderWallBlokSiz*5, bsU5)
				level[a].doorExitRecs = append(level[a].doorExitRecs, exitRec)

				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
			}
			if level[a].doorSides[b] == 2 {
				checkV2 := rl.NewVector2(levX+levW-borderWallBlokSiz/2, levY+levW/2)
				checkV2L1 := checkV2
				checkV2L1.Y -= borderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.Y -= borderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.Y += borderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.Y += borderWallBlokSiz

				exitRec := rl.NewRectangle((checkV2L2.X+borderWallBlokSiz/2)-bsU5, checkV2L2.Y-borderWallBlokSiz/2, bsU5, borderWallBlokSiz*5)
				level[a].doorExitRecs = append(level[a].doorExitRecs, exitRec)

				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
			}

			if level[a].doorSides[b] == 3 {
				checkV2 := rl.NewVector2(levX+levW/2, levY+levW-borderWallBlokSiz/2)
				checkV2L1 := checkV2
				checkV2L1.X -= borderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.X -= borderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.X += borderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.X += borderWallBlokSiz

				exitRec := rl.NewRectangle(checkV2L2.X-borderWallBlokSiz/2, (checkV2L2.Y+borderWallBlokSiz/2)-bsU5, borderWallBlokSiz*5, bsU5)
				level[a].doorExitRecs = append(level[a].doorExitRecs, exitRec)

				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
			}

			if level[a].doorSides[b] == 4 {
				checkV2 := rl.NewVector2(levX+borderWallBlokSiz/2, levY+levW/2)
				checkV2L1 := checkV2
				checkV2L1.Y -= borderWallBlokSiz
				checkV2L2 := checkV2L1
				checkV2L2.Y -= borderWallBlokSiz
				checkV2R1 := checkV2
				checkV2R1.Y += borderWallBlokSiz
				checkV2R2 := checkV2R1
				checkV2R2.Y += borderWallBlokSiz

				exitRec := rl.NewRectangle((checkV2L2.X - borderWallBlokSiz/2), checkV2L2.Y-borderWallBlokSiz/2, bsU5, borderWallBlokSiz*5)
				level[a].doorExitRecs = append(level[a].doorExitRecs, exitRec)

				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2L2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R1, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
				for c := 0; c < len(level[a].walls); c++ {
					if rl.CheckCollisionPointRec(checkV2R2, level[a].walls[c].rec) {
						level[a].walls = remBlok(level[a].walls, c)
					}
				}
			}

		}

	}

	level[0].visited = true

	makeInnerBloks()
	makemovebloks()
	makestatues()
	makeshop()
	makeetc()

	if levelnum > 3 {
		makeblades()
		maketurrets()
	}
	if levelnum > 2 {
		makespears()
	}
	if levelnum > 1 {
		makespikes()
	}

	makeenemies()

	makeexit()

	cleanlevel()

}
func makenewlevel() { //MARK:MAKE NEW LEVEL

	nextlevelscreen = true
	levelnum++

	for a := 0; a < len(level); a++ {
		level[a].etc = nil
		level[a].enemies = nil
		level[a].doorExitRecs = nil
		level[a].doorSides = nil
		level[a].floor = nil
		level[a].innerBloks = nil
		level[a].movBloks = nil
		level[a].nextRooms = nil
		level[a].spikes = nil
		level[a].visited = false
		level[a].walls = nil
	}

	level = nil
	fx = nil
	plProj = nil
	enProj = nil
	floodRec.Y = scrHF32 + bsU

	if levelnum == 6 {
		mods.planty = false
		mods.carrot = false
		mods.alien = false
		mods.vine = false
		mods.airstrike = false
		makeendlevel()
	} else {
		makelevel()
	}
	nextlevelT = fps

	pl.cnt = cnt
	cntCompanion := pl.cnt

	if mods.carrot {
		mrcarrot.rec = rl.NewRectangle(cntCompanion.X-mrcarrot.rec.Width/2, cntCompanion.Y-mrcarrot.rec.Width/2, mrcarrot.rec.Width, mrcarrot.rec.Width)
	}
	if mods.alien {
		mralien.rec = rl.NewRectangle(cntCompanion.X-mralien.rec.Width/2, cntCompanion.Y-mralien.rec.Width/2, mralien.rec.Width, mralien.rec.Width)
	}
	if mods.planty {
		mrplanty.rec = rl.NewRectangle(cntCompanion.X-mrplanty.rec.Width/2, cntCompanion.Y-mrplanty.rec.Width/2, mrplanty.rec.Width, mrplanty.rec.Width)
	}

	roomNum = 0

}
func makeInnerBloks() { //MARK: MAKE INNER BLOKS

	for a := 0; a < len(level); a++ {

		num := rInt(3, 7)

		for num > 0 {
			countbreak := 100
			for {
				siz := rF32(bsU4, bsU8)
				zblok := xblok{}
				tl := findRanRecLoc(siz, siz, a)
				zblok.rec = rl.NewRectangle(tl.X, tl.Y, siz, siz)
				zblok.img = wallT
				switch levelnum {
				case 1:
					zblok.color = ranBlue()
				case 2:
					zblok.color = ranBrown()
				case 3:
					zblok.color = ranOrange()
				case 4:
					zblok.color = ranDarkBlue()
				case 5:
					zblok.color = ranCol()
				case 6:
					zblok.color = ranOrange()
				}
				zblok.fade = 1
				zblok.crec = zblok.rec
				zblok.solid = true

				canadd := true
				for b := 0; b < len(level[a].innerBloks); b++ {
					if rl.CheckCollisionRecs(zblok.rec, level[a].innerBloks[b].rec) {
						canadd = false
					}
					if a == 0 {
						checkrec := rl.NewRectangle(cnt.X-bsU4, cnt.Y-bsU4, bsU8, bsU8)
						if rl.CheckCollisionRecs(zblok.rec, checkrec) {
							canadd = false
						}
					}
				}
				countbreak--

				if canadd || countbreak == 0 {
					level[a].innerBloks = append(level[a].innerBloks, zblok)
					break
				}
			}
			num--
		}
	}
}

func makemario() { //MARK:MAKE MARIO

	mariorecs = nil
	marioCols = nil
	mariocoins = nil
	mariocoinonoff = nil
	marioT = fps * 5

	//IMG
	marioImg = knight[0]

	//BORDER REC
	marioScreenRec = rl.NewRectangle(cnt.X-scrWF32/(cam2.Zoom*2), cnt.Y-scrHF32/(cam2.Zoom*2), scrWF32/cam2.Zoom, scrHF32/cam2.Zoom)

	//BACK PATTERN
	patternRec = patterns[rInt(0, len(patterns))]

	//FLOOR
	siz := bsU4
	x := float32(0)
	y := levRecInner.Y + levRec.Width - (siz + bsU)
	for {
		mariorecs = append(mariorecs, rl.NewRectangle(x, y, siz, siz))
		marioCols = append(marioCols, ranBrown())
		if roll6() > 4 {
			mariocoins = append(mariocoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			mariocoinonoff = append(mariocoinonoff, true)
		}
		x += siz
		if x >= scrWF32 {
			break
		}
	}

	//PLAYER REC
	marioPL = rl.NewRectangle(scrWF32/2, y-siz, siz, siz)
	marioV2L = rl.NewVector2(marioPL.X, marioPL.Y+marioPL.Width+2)
	marioV2R = rl.NewVector2(marioPL.X+marioPL.Width, marioPL.Y+marioPL.Width+2)

	//PLATFORMS
	y -= siz * 3
	num := float32(rInt(5, 9))
	x = marioScreenRec.X + rF32(0, (marioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		mariorecs = append(mariorecs, rl.NewRectangle(x, y, siz, siz))
		marioCols = append(marioCols, ranGreen())
		if roll6() > 4 {
			mariocoins = append(mariocoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			mariocoinonoff = append(mariocoinonoff, true)
		}
		x += siz
		num--
	}
	num = float32(rInt(5, 9))
	x = marioScreenRec.X + marioScreenRec.Width/2 + rF32(0, (marioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		mariorecs = append(mariorecs, rl.NewRectangle(x, y, siz, siz))
		marioCols = append(marioCols, ranOrange())
		if roll6() > 4 {
			mariocoins = append(mariocoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			mariocoinonoff = append(mariocoinonoff, true)
		}
		x += siz
		num--
	}
	y -= siz * 3
	num = float32(rInt(5, 9))
	x = marioScreenRec.X + rF32(0, (marioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		mariorecs = append(mariorecs, rl.NewRectangle(x, y, siz, siz))
		marioCols = append(marioCols, ranCyan())
		if roll6() > 4 {
			mariocoins = append(mariocoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			mariocoinonoff = append(mariocoinonoff, true)
		}
		x += siz
		num--
	}
	num = float32(rInt(5, 9))
	x = marioScreenRec.X + marioScreenRec.Width/2 + rF32(0, (marioScreenRec.Width/2)-(num*siz))
	for num > 0 {
		mariorecs = append(mariorecs, rl.NewRectangle(x, y, siz, siz))
		marioCols = append(marioCols, ranPink())
		if roll6() > 4 {
			mariocoins = append(mariocoins, rl.NewRectangle(x, y-siz, siz/2, siz/2))
			mariocoinonoff = append(mariocoinonoff, true)
		}
		x += siz
		num--
	}

}
func makeairstrike() { //MARK:MAKE AIR STRIKE

	airstrikeV2 = nil
	airstrikeDir = rInt(1, 5)
	//airstrikeDir = 4

	switch airstrikeDir {
	case 1:
		v2 := rl.NewVector2(rF32(levRecInner.X+bsU4, levRecInner.X+levRecInner.Width-bsU8), levRecInner.Y-bsU4)
		airstrikeV2 = append(airstrikeV2, v2)
		v2.Y -= bsU4
		v2.X -= bsU4
		airstrikeV2 = append(airstrikeV2, v2)
		v2.X += bsU8
		airstrikeV2 = append(airstrikeV2, v2)
	case 2:
		v2 := rl.NewVector2(levRecInner.X+levRecInner.Width+bsU4, rF32(levRecInner.Y+bsU4, levRecInner.Y+levRecInner.Width-bsU8))
		airstrikeV2 = append(airstrikeV2, v2)
		v2.Y -= bsU4
		v2.X += bsU4
		airstrikeV2 = append(airstrikeV2, v2)
		v2.Y += bsU8
		airstrikeV2 = append(airstrikeV2, v2)
	case 3:
		v2 := rl.NewVector2(rF32(levRecInner.X+bsU4, levRecInner.X+levRecInner.Width-bsU8), levRecInner.Y+levRecInner.Width+bsU4)
		airstrikeV2 = append(airstrikeV2, v2)
		v2.Y += bsU4
		v2.X -= bsU4
		airstrikeV2 = append(airstrikeV2, v2)
		v2.X += bsU8
		airstrikeV2 = append(airstrikeV2, v2)
	case 4:
		v2 := rl.NewVector2(levRecInner.X-bsU4, rF32(levRecInner.Y+bsU4, levRecInner.Y+levRecInner.Width-bsU8))
		airstrikeV2 = append(airstrikeV2, v2)
		v2.Y -= bsU4
		v2.X -= bsU4
		airstrikeV2 = append(airstrikeV2, v2)
		v2.Y += bsU8
		airstrikeV2 = append(airstrikeV2, v2)
	}

	airstrikebombT = rI32(int(fps/4), int(fps*2))
	airstrikeOn = true
}
func makefish() { //MARK:MAKE FISH

	fishV2 = rl.Vector2{}
	fish2V2 = rl.Vector2{}

	fishSiz = rF32(bsU4, bsU7)
	fishSiz2 = rF32(bsU4, bsU7)

	fishV2.X = -fishSiz
	fish2V2.X = scrWF32 + fishSiz2

	fishV2.Y = rF32(scrHF32/3, scrHF32)
	fish2V2.Y = rF32(scrHF32/3, scrHF32)

	fishRec = rl.NewRectangle(fishV2.X, fishV2.Y, fishSiz, fishSiz)
	fishRec2 = rl.NewRectangle(fish2V2.X, fish2V2.Y, fishSiz2, fishSiz2)

	fish1 = fishR.recTL
	fish2 = fishL.recTL

}
func makeplayer() { //MARK:MAKE PLAYER

	pl.atkDMG = 1
	pl.cnt = cnt
	pl.hp = 5
	pl.hpmax = 5
	pl.vel = 4
	pl.siz = 72
	pl.rec = rl.NewRectangle(pl.cnt.X-pl.siz/2, pl.cnt.Y-pl.siz/2, pl.siz, pl.siz)
	pl.crec = pl.rec
	pl.arec = pl.rec
	pl.atkrec = pl.rec
	pl.atkrec.X -= bsU
	pl.atkrec.Y -= bsU
	pl.atkrec.Width += bsU2
	pl.atkrec.Height += bsU2
	pl.crec.X += pl.crec.Width / 3
	pl.crec.Y += pl.crec.Height / 3
	pl.crec.Width = pl.crec.Width / 3
	pl.crec.Height = pl.crec.Height / 2
	pl.framesAtk = 6
	pl.framesWalk = 8
	pl.sizImg = 32
	pl.img = knight[0]
	pl.ori = rl.NewVector2(pl.rec.Width/2, pl.rec.Height/2)
	pl.orbimg1 = orbitalanim.recTL
	pl.orbimg2 = pl.orbimg1

	max.axe = 10
	max.fireball = 8
	max.bounce = 4
	max.key = 3
	max.apple = 3
	max.firetrail = 3
	max.hppotion = 3
	max.coffee = 5
	max.atkrange = 3
	max.atkdmg = 2
	max.orbital = 2
	max.hpring = 3
	max.armor = 3
	max.cherry = 99
	max.cake = 99

}
func makeEnemyTypes() { //MARK: MAKE ENEMY TYPES

	siz := bsU4

	enSpikes.img = rl.NewRectangle(0, 404, 44, 44)
	enSpikes.fade = 1
	enSpikes.col = rl.White
	enSpikes.xImg = 0
	enSpikes.frameNum = 15
	enSpikes.vel = bsU / 5
	enSpikes.name = "spikehog"
	enSpikes.hp = 5
	enSpikes.hpmax = enSpikes.hp
	enSpikes.rec = rl.NewRectangle(0, 0, siz, siz)

	enGhost.imgl = rl.NewRectangle(4, 506, 44, 44)
	enGhost.imgr = rl.NewRectangle(454, 506, 44, 44)
	enGhost.img = enGhost.imgr
	enGhost.col = rl.White
	enGhost.fade = 1
	enGhost.xImg = 4
	enGhost.xImg2 = 454
	enGhost.frameNum = 9
	enGhost.vel = bsU / 5
	enGhost.name = "ghost"
	enGhost.hp = 4
	enGhost.hpmax = enGhost.hp
	enGhost.rec = rl.NewRectangle(0, 0, siz, siz)

	enSlime.imgl = rl.NewRectangle(4, 556, 44, 44)
	enSlime.imgr = rl.NewRectangle(458, 556, 44, 44)
	enSlime.img = enSlime.imgr
	enSlime.col = rl.White
	enSlime.fade = 1
	enSlime.xImg = 4
	enSlime.xImg2 = 458
	enSlime.frameNum = 9
	enSlime.vel = bsU / 5
	enSlime.name = "slime"
	enSlime.hp = 5
	enSlime.hpmax = enSlime.hp
	enSlime.rec = rl.NewRectangle(0, 0, siz, siz)

	enRock.imgl = rl.NewRectangle(4, 608, 22, 22)
	enRock.imgr = rl.NewRectangle(318, 608, 22, 22)
	enRock.img = enRock.imgr
	enRock.col = rl.White
	enRock.fade = 1
	enRock.xImg = 4
	enRock.xImg2 = 318
	enRock.frameNum = 13
	enRock.vel = bsU / 4
	enRock.name = "rock"
	enRock.hp = 3
	enRock.hpmax = enRock.hp
	enRock.rec = rl.NewRectangle(0, 0, bsU3, bsU3)

	enMushroom.imgl = rl.NewRectangle(4, 634, 32, 32)
	enMushroom.imgr = rl.NewRectangle(4, 668, 32, 32)
	enMushroom.img = enMushroom.imgr
	enMushroom.col = rl.White
	enMushroom.fade = 1
	enMushroom.xImg = 4
	enMushroom.xImg2 = 4
	enMushroom.frameNum = 13
	enMushroom.vel = bsU / 5
	enMushroom.name = "mushroom"
	enMushroom.hp = 5
	enMushroom.hpmax = enMushroom.hp
	enMushroom.rec = rl.NewRectangle(0, 0, siz, siz)

}
func makeChainLightning() { //MARK:MAKE CHAIN LIGHTNING

	chainV2 = nil
	if len(level[roomNum].enemies) > 1 {
		for a := 0; a < len(level[roomNum].enemies); a++ {
			if level[roomNum].enemies[a].cnt.X > levRecInner.X && level[roomNum].enemies[a].cnt.Y > levRecInner.Y {
				chainV2 = append(chainV2, level[roomNum].enemies[a].cnt)
			}
		}
		chainLightTimer = fps / 3
		chainLightOn = true
		rl.PlaySound(sfx[11])
	}
}
func makeenemies() { //MARK:MAKE ENEMIES

	makeEnemyTypes()

	zen := xenemy{}
	siz := bsU2

	for a := 0; a < len(level); a++ {

		if levelnum > 4 {
			if flipcoin() { //MUSHROOM
				zen = xenemy{}
				zen = enMushroom
				zen.T1 = rI32(int(fps*2), int(fps*5))
				canadd := true
				zen.rec, canadd = findRecPos(enMushroom.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 2
				zen.crec.Height -= zen.crec.Height / 2
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				if canadd {
					level[a].enemies = append(level[a].enemies, zen)
				}
			}
			if hardcore {
				if flipcoin() { //MUSHROOM
					zen = xenemy{}
					zen = enMushroom
					zen.T1 = rI32(int(fps*2), int(fps*5))
					canadd := true
					zen.rec, canadd = findRecPos(enMushroom.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 2
					zen.crec.Height -= zen.crec.Height / 2
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}
				if flipcoin() { //MUSHROOM
					zen = xenemy{}
					zen = enMushroom
					zen.T1 = rI32(int(fps*2), int(fps*5))
					canadd := true
					zen.rec, canadd = findRecPos(enMushroom.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 2
					zen.crec.Height -= zen.crec.Height / 2
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}
			}
		}
		if levelnum > 3 {

			if flipcoin() { //GHOST
				zen = xenemy{}
				zen = enGhost
				zen.fly = true
				canadd := true
				zen.rec, canadd = findRecPos(enGhost.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.crec.X += zen.crec.Width / 8
				zen.crec.Width -= zen.crec.Width / 4
				zen.arec = zen.crec
				zen.arec.X -= zen.arec.Width * 2
				zen.arec.Y -= zen.arec.Width * 2
				zen.arec.Width = zen.arec.Width * 5
				zen.arec.Height = zen.arec.Height * 5

				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)

				if canadd {
					level[a].enemies = append(level[a].enemies, zen)
				}
			}
			if hardcore {
				if flipcoin() { //GHOST
					zen = xenemy{}
					zen = enGhost
					zen.fly = true
					canadd := true
					zen.rec, canadd = findRecPos(enGhost.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.crec.X += zen.crec.Width / 8
					zen.crec.Width -= zen.crec.Width / 4
					zen.arec = zen.crec
					zen.arec.X -= zen.arec.Width * 2
					zen.arec.Y -= zen.arec.Width * 2
					zen.arec.Width = zen.arec.Width * 5
					zen.arec.Height = zen.arec.Height * 5

					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)

					if canadd {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}
				if flipcoin() { //GHOST
					zen = xenemy{}
					zen = enGhost
					zen.fly = true
					canadd := true
					zen.rec, canadd = findRecPos(enGhost.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.crec.X += zen.crec.Width / 8
					zen.crec.Width -= zen.crec.Width / 4
					zen.arec = zen.crec
					zen.arec.X -= zen.arec.Width * 2
					zen.arec.Y -= zen.arec.Width * 2
					zen.arec.Width = zen.arec.Width * 5
					zen.arec.Height = zen.arec.Height * 5

					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)

					if canadd {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}

			}

		}
		if levelnum > 1 {

			if flipcoin() { //SLIME
				zen = xenemy{}
				zen = enSlime
				zen.T1 = rI32(int(fps/2), int(fps*2))
				canadd := true
				zen.rec, canadd = findRecPos(enSlime.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				if canadd {
					level[a].enemies = append(level[a].enemies, zen)
				}
			}

			if hardcore {
				if flipcoin() { //SLIME
					zen = xenemy{}
					zen = enSlime
					zen.T1 = rI32(int(fps/2), int(fps*2))
					canadd := true
					zen.rec, canadd = findRecPos(enSlime.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}
				if flipcoin() { //SLIME
					zen = xenemy{}
					zen = enSlime
					zen.T1 = rI32(int(fps/2), int(fps*2))
					canadd := true
					zen.rec, canadd = findRecPos(enSlime.rec.Width, a)
					zen.crec = zen.rec
					zen.crec.Y += zen.crec.Height / 3
					zen.crec.Height -= zen.crec.Height / 3
					zen.velX = rF32(-zen.vel, zen.vel)
					zen.velY = rF32(-zen.vel, zen.vel)
					if canadd {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}
			}
		}

		if flipcoin() { //ROCK
			zen = xenemy{}
			zen = enRock
			zen.spawnN = rInt(2, 11)
			zen.T1 = rI32(int(fps*2), int(fps*5))
			canadd := true
			zen.rec, canadd = findRecPos(enRock.rec.Width, a)
			zen.crec = zen.rec
			zen.crec.Y += zen.crec.Height / 3
			zen.crec.Height -= zen.crec.Height / 3
			zen.velX = rF32(-zen.vel, zen.vel)
			zen.velY = rF32(-zen.vel, zen.vel)
			if canadd {
				level[a].enemies = append(level[a].enemies, zen)
			}
		} else { //SPIKEHOG
			zen = xenemy{}
			zen = enSpikes
			canadd := true
			zen.rec, canadd = findRecPos(enSpikes.rec.Width, a)
			zen.crec = zen.rec
			zen.crec.Height = zen.crec.Height / 2
			zen.crec.Y += zen.crec.Height
			if flipcoin() {
				zen.velX = zen.vel
				if flipcoin() {
					zen.velX *= -1
				}
			} else {
				zen.velY = zen.vel
				if flipcoin() {
					zen.velY *= -1
				}
			}
			if canadd {
				level[a].enemies = append(level[a].enemies, zen)
			}
		}
		if hardcore {
			if flipcoin() { //ROCK
				zen = xenemy{}
				zen = enRock
				zen.spawnN = rInt(2, 11)
				zen.T1 = rI32(int(fps*2), int(fps*5))
				canadd := true
				zen.rec, canadd = findRecPos(enRock.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Y += zen.crec.Height / 3
				zen.crec.Height -= zen.crec.Height / 3
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				if canadd {
					level[a].enemies = append(level[a].enemies, zen)
				}
			} else { //SPIKEHOG
				zen = xenemy{}
				zen = enSpikes
				canadd := true
				zen.rec, canadd = findRecPos(enSpikes.rec.Width, a)
				zen.crec = zen.rec
				zen.crec.Height = zen.crec.Height / 2
				zen.crec.Y += zen.crec.Height
				if flipcoin() {
					zen.velX = zen.vel
					if flipcoin() {
						zen.velX *= -1
					}
				} else {
					zen.velY = zen.vel
					if flipcoin() {
						zen.velY *= -1
					}
				}
				if canadd {
					level[a].enemies = append(level[a].enemies, zen)
				}
			}
		}

		//BATS
		if flipcoin() {
			zen.name = "bat"
			zen.hp = 2
			zen.hpmax = zen.hp
			zen.fly = true
			zen.vel = bsU / 3
			zen.velX = rF32(-zen.vel, zen.vel)
			zen.velY = rF32(-zen.vel, zen.vel)
			zen.img = bats[rInt(0, len(bats))]
			zen.fade = 1
			zen.col = rl.White
			zen.frameNum = 3
			zen.xImg = zen.img.X
			canadd := true
			zen.rec, canadd = findRecPos(siz, a)
			zen.ori = rl.NewVector2(0, 0)
			zen.crec = zen.rec
			if canadd {
				level[a].enemies = append(level[a].enemies, zen)
			}

			if flipcoin() {
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				zen.img = bats[rInt(0, len(bats))]
				zen.xImg = zen.img.X
				canadd = true
				zen.rec, canadd = findRecPos(siz, a)
				zen.crec = zen.rec
				if canadd {
					level[a].enemies = append(level[a].enemies, zen)
				}
			}
		}

		//RABBITS
		if flipcoin() {
			zen = xenemy{}
			zen.hp = 3
			zen.hpmax = zen.hp
			zen.name = "rabbit1"
			zen.anim = true
			zen.vel = bsU / 4
			zen.velX = rF32(-zen.vel, zen.vel)
			zen.velY = rF32(-zen.vel, zen.vel)
			zen.img = rabbit1.recTL
			zen.fade = 1
			zen.col = ranCol()
			canadd := true
			zen.rec, canadd = findRecPos(siz, a)
			zen.ori = rl.NewVector2(0, 0)
			zen.rec = rl.NewRectangle(zen.cnt.X-siz/2, zen.cnt.Y-siz/2, siz, siz)
			zen.crec = zen.rec
			if canadd {
				if zen.rec.X > levRecInner.X && zen.rec.Y > levRecInner.Y {
					level[a].enemies = append(level[a].enemies, zen)
				}
			}

			if flipcoin() {
				zen.velX = rF32(-zen.vel, zen.vel)
				zen.velY = rF32(-zen.vel, zen.vel)
				zen.col = ranCol()
				canadd := true
				zen.rec, canadd = findRecPos(siz, a)
				zen.crec = zen.rec
				if canadd {
					if zen.rec.X > levRecInner.X && zen.rec.Y > levRecInner.Y {
						level[a].enemies = append(level[a].enemies, zen)
					}
				}
			}
		}

	}

}
func maketeleport() { //MARK:MAKE TELEPORT
	teleportRadius = nil
	length := scrWF32 / 2
	for a := 0; a < 10; a++ {
		teleportRadius = append(teleportRadius, length)
		length -= bsU4
	}

}
func makegascloud(cnt rl.Vector2) { //MARK:MAKE GAS CLOUD

	siz := bsU10

	zblok := xblok{}
	zblok.fade = 1
	zblok.cnt = cnt
	zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
	zblok.crec = zblok.rec
	zblok.crec.X += zblok.crec.Width / 4
	zblok.crec.Y += zblok.crec.Width / 4
	zblok.crec.Width = zblok.crec.Width / 2
	zblok.crec.Height = zblok.crec.Height / 2
	zblok.drec = zblok.rec
	zblok.drec.X += zblok.rec.Width / 2
	zblok.drec.Y += zblok.rec.Height / 2
	zblok.ori = rl.NewVector2(zblok.rec.Width/2, zblok.rec.Width/2)
	zblok.name = "gascloud"
	zblok.onoff = true
	zblok.color = ranGreen()
	zblok.img = posiongas.recTL
	zblok.vel = bsU / 4
	zblok.velX = rF32(-zblok.vel, zblok.vel)
	zblok.velY = rF32(-zblok.vel, zblok.vel)
	level[roomNum].etc = append(level[roomNum].etc, zblok)

}
func makeSwitchArrows() { //MARK:MAKE SWITCH ARROWS

	choose := rInt(1, 5)

	switch choose {
	case 1:
		x := levRecInner.X + bsU
		y := levRecInner.Y
		for {
			makeProjectileEnemy(2, rl.NewVector2(x+bsU, y+bsU))
			x += bsU2
			if x > levRecInner.X+levRecInner.Width-bsU2 {
				break
			}
		}

	case 2:
		x := levRecInner.X + levRecInner.Width - bsU2
		y := levRecInner.Y + bsU
		for {
			makeProjectileEnemy(5, rl.NewVector2(x+bsU, y+bsU))
			y += bsU2
			if y > levRecInner.Y+levRecInner.Width-bsU2 {
				break
			}
		}

	case 3:
		x := levRecInner.X + bsU
		y := levRecInner.Y + levRecInner.Width - bsU2
		for {
			makeProjectileEnemy(4, rl.NewVector2(x+bsU, y+bsU))
			x += bsU2
			if x > levRecInner.X+levRecInner.Width-bsU2 {
				break
			}
		}

	case 4:
		x := levRecInner.X
		y := levRecInner.Y + bsU
		for {
			makeProjectileEnemy(3, rl.NewVector2(x+bsU, y+bsU))
			y += bsU2
			if y > levRecInner.Y+levRecInner.Width-bsU2 {
				break
			}
		}

	}

}

func makeexit() { //MARK:MAKE EXIT

	for {
		exitRoomNum = rInt(1, len(level))
		if exitRoomNum != shopRoomNum {
			level[exitRoomNum].exit = true
			break
		}
	}

	zblok := makeBlokGenNoRecNoCntr()
	zblok.img = etc[53]
	zblok.name = "exit"
	zblok.color = ranBrown()
	siz := bsU3
	for {
		canadd := true
		zblok.rec, canadd = findRecPoswithSpacing(siz, bsU2, exitRoomNum)
		if canadd {
			break
		}
	}

	level[exitRoomNum].etc = append(level[exitRoomNum].etc, zblok)

}
func makestatues() { //MARK:MAKE STATUES

	for a := 0; a < len(level); a++ {

		if roll6() > 4 {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU3, bsU5)
			found := true
			zblok.rec, found = findRecPoswithSpacing(siz, bsU2, a)
			zblok.img = statues[rInt(0, len(statues))]
			zblok.name = "statue"
			zblok.solid = true

			if found {
				level[a].etc = append(level[a].etc, zblok)
			}
		}
	}

}
func makespears() { //MARK:MAKE SPEARS

	for a := 0; a < len(level); a++ {

		zblok := makeBlokGenNoRecNoCntr()
		multi := float32(2)
		choose := rInt(1, 5)
		switch choose {
		case 1:
			zblok.ro = 180
			zblok.rec = rl.NewRectangle(levRecInner.X, levRecInner.Y, multi*spear.recTL.Width, multi*spear.recTL.Height)
			for {
				zblok.rec.X += rF32(0, levRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}
			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = zblok.rec
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)
		case 2:
			zblok.ro = 270
			zblok.rec = rl.NewRectangle(levRecInner.X+levRecInner.Width-bsU-(multi*spear.recTL.Height)/2, levRecInner.Y+bsU-(multi*spear.recTL.Height)/2, multi*spear.recTL.Width, multi*spear.recTL.Height)

			for {
				zblok.rec.Y += rF32(0, levRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}

			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = rl.NewRectangle(levRecInner.X+levRecInner.Width-multi*spear.recTL.Height, (zblok.rec.Y+zblok.rec.Height/2)-bsU, multi*spear.recTL.Height, multi*spear.recTL.Width)
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)

		case 3:
			zblok.ro = 0
			zblok.rec = rl.NewRectangle(levRecInner.X, levRecInner.Y+levRecInner.Height-(multi*spear.recTL.Height), multi*spear.recTL.Width, multi*spear.recTL.Height)
			for {
				zblok.rec.X += rF32(0, levRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}
			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = zblok.rec
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)

		case 4:
			zblok.ro = 90
			zblok.rec = rl.NewRectangle(levRecInner.X+bsU3, levRecInner.Y+bsU-(multi*spear.recTL.Height)/2, multi*spear.recTL.Width, multi*spear.recTL.Height)

			for {
				zblok.rec.Y += rF32(0, levRecInner.Width-bsU2)
				if checkInnerBloksExits(a, zblok.rec) {
					break
				}
			}

			zblok.drec = zblok.rec
			zblok.drec.X += zblok.drec.Width / 2
			zblok.drec.Y += zblok.drec.Height / 2
			zblok.crec = rl.NewRectangle(zblok.rec.X-bsU3, (zblok.rec.Y+zblok.rec.Height/2)-bsU, multi*spear.recTL.Height, multi*spear.recTL.Width)
			zblok.ori = rl.NewVector2(zblok.drec.Width/2, zblok.drec.Height/2)
		}

		zblok.img = spear.recTL
		zblok.name = "spear"
		if zblok.rec.X > levRecInner.X && zblok.rec.X < levRecInner.Width+levRecInner.X && zblok.rec.Y > levRecInner.Y && zblok.rec.Y < levRecInner.Y+levRecInner.Width {
			level[a].etc = append(level[a].etc, zblok)
		}

	}
}
func makeblades() { //MARK:MAKE BLADES

	for a := 0; a < len(level); a++ {

		if flipcoin() {
			siz := rF32(bsU2, bsU4)
			zblok := makeBlokGenNoRecNoCntr()
			found := true
			zblok.rec, found = findRecPos(siz, a)
			zblok.img = blades.recTL
			zblok.name = "blades"
			zblok.vel = bsU / 4
			zblok.velX = rF32(-zblok.vel, zblok.vel)
			zblok.velY = rF32(-zblok.vel, zblok.vel)

			if found {
				level[a].etc = append(level[a].etc, zblok)
			}
		}
	}
}
func makeshrines() { //MARK:MAKE SHRINES

	choose := rInt(0, len(level))

	siz := rF32(bsU3, bsU5)
	zblok := xblok{}
	for {
		zblok = makeBlokGenRandom(siz)
		if checkInnerBloksExits(choose, zblok.rec) {
			break
		}
	}

	zblok.img = shrines[rInt(0, len(shrines))]
	zblok.name = "shrine"
	zblok.color = ranCol()
	zblok.fade = 0.2
	zblok.solid = true
	zblok.onoff = true
	level[choose].etc = append(level[choose].etc, zblok)

}
func makeetc() { //MARK:MAKE ETC

	for a := 0; a < len(level); a++ {

		//GAS CLOUD TRAPS
		if roll6() == 6 {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.drec = zblok.rec
			zblok.drec.X += zblok.rec.Width / 2
			zblok.drec.Y += zblok.rec.Height / 2
			zblok.ori = rl.NewVector2(zblok.rec.Width/2, zblok.rec.Width/2)
			zblok.name = "gascloudtrap"
			zblok.color = rl.SkyBlue
			zblok.img = etc[26]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
			}
		}

		//SWITCHES
		if roll6() == 6 {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "switch"
			zblok.numType = roll6()
			if zblok.numType == 6 {
				zblok.numCoins = rInt(3, 11)
			}
			zblok.onoffswitch = flipcoin()
			zblok.color = rl.SkyBlue
			if zblok.onoffswitch {
				zblok.img = etc[21]
			} else {
				zblok.img = etc[22]
			}
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
			}
		}
		//CHESTS
		if roll6() == 6 {

			siz := bsU2 + bsU/2
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(siz, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "chest"
			zblok.fade = rF32(0.5, 0.8)
			zblok.color = ranOrange()
			zblok.crec = zblok.rec
			zblok.crec.Width += bsU
			zblok.crec.Height += bsU
			zblok.crec.X -= bsU / 2
			zblok.crec.Y -= bsU / 2
			zblok.img = etc[23]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
			}
		}
		//SKULLS
		num := rInt(1, 5)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32((bsU/4)*3, bsU+bsU/2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "skull"
			zblok.fade = rF32(0.3, 0.6)
			zblok.color = ranGrey()
			zblok.img = skulls[rInt(0, len(skulls))]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
				num--
			}
			if num <= 0 {
				break
			}
		}

		//CANDLES
		num = rInt(0, 3)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU+bsU/2, bsU2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.onoff = true
			zblok.name = "candle"
			zblok.fade = rF32(0.5, 0.8)
			zblok.color = rl.White
			zblok.img = candles[rInt(0, len(candles))]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//SIGNS
		num = rInt(0, 2)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU+bsU/2, bsU2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "sign"
			zblok.fade = rF32(0.4, 0.7)
			zblok.color = ranBrown()
			zblok.img = signs[rInt(0, len(signs))]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//PLANTS
		num = rInt(1, 5)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU+bsU/2, bsU2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "plant"
			zblok.fade = rF32(0.3, 0.6)
			zblok.color = ranGreen()
			zblok.img = plants[rInt(0, len(plants))]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//MUSHROOMS
		num = rInt(1, 5)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			siz := rF32(bsU, bsU+bsU/2)
			canadd := true
			zblok.rec, canadd = findRecPos(siz, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.name = "mushroom"
			zblok.fade = rF32(0.3, 0.6)
			zblok.color = rl.White
			zblok.img = mushrooms[rInt(0, len(mushrooms))]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}
		//POWERUP BLOK
		if roll6() > 2 {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU+bsU/2, bsU/4, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.solid = true
			zblok.name = "powerupBlok"
			zblok.color = brightYellow()
			zblok.img = etc[27]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)

			}
		}

		//OIL BARRELS
		num = rInt(0, 3)
		for {
			zblok := makeBlokGenNoRecNoCntr()
			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(bsU2, bsU/2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.solid = true
			zblok.onoff = true
			zblok.name = "oilbarrel"
			zblok.color = rl.DarkGreen
			zblok.img = etc[20]
			if canadd {
				level[a].etc = append(level[a].etc, zblok)
				num--
			}

			if num <= 0 {
				break
			}
		}

		//SPRINGS
		if flipcoin() {
			num := rInt(1, 4)
			siz := bsU2
			zblok := xblok{}
			zblok.onoff = true
			zblok.name = "spring"
			zblok.fade = 1
			zblok.color = rl.White
			zblok.rec = rl.NewRectangle(0, 0, siz, siz)
			zblok.img = spring.recTL

			addtolevel := false

			for num > 0 {
				choose := rInt(1, 5)
				switch choose {
				case 1:
					zblok.rec.Y = levRecInner.Y
					for {
						zblok.rec.X = levRecInner.X + rF32(0, levRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 180
					zblok.slideDIR = 3
				case 2:
					zblok.rec.X = levRecInner.X + levRecInner.Width - siz
					for {
						zblok.rec.Y = levRecInner.Y + rF32(0, levRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 270
					zblok.slideDIR = 4
				case 3:
					zblok.rec.Y = levRecInner.Y + levRecInner.Width - siz
					for {
						zblok.rec.X = levRecInner.X + rF32(0, levRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 0
					zblok.slideDIR = 1
				case 4:
					zblok.rec.X = levRecInner.X
					for {
						zblok.rec.Y = levRecInner.Y + rF32(0, levRecInner.Width-bsU2)
						canadd := true
						for b := 0; b < len(level[a].doorExitRecs); b++ {
							if rl.CheckCollisionRecs(zblok.rec, level[a].doorExitRecs[b]) {
								canadd = false
							}
						}
						if canadd {
							addtolevel = true
							break
						}
					}
					zblok.ro = 90
					zblok.slideDIR = 2
				}

				if addtolevel {
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec

					level[a].etc = append(level[a].etc, zblok)

					num--
				}
			}
		}

		//SPRING BLOKS
		if flipcoin() {
			siz := bsU2
			zblok := xblok{}
			zblok = makeBlokGenNoRecNoCntr()

			canadd := true
			zblok.rec, canadd = findRecPoswithSpacing(siz, bsU2, a)
			zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Width/2)
			zblok.solid = true
			zblok.name = "springblok"
			zblok.fade = 1
			switch levelnum {
			case 1:
				zblok.color = ranBlue()
			case 2:
				zblok.color = ranBrown()
			case 3:
				zblok.color = ranOrange()
			case 4:
				zblok.color = ranDarkBlue()
			case 5:
				zblok.color = ranCol()
			}
			zblok.img = wallT
			if canadd {
				level[a].etc = append(level[a].etc, zblok)

				zblok.name = "spring"
				zblok.solid = false
				zblok.color = rl.White
				zblok.img = spring.recTL

				if flipcoin() {
					zblok.rec.X += siz
					zblok.slideDIR = 2
					zblok.ro = 90
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					level[a].etc = append(level[a].etc, zblok)
				} else {
					zblok.rec.X += siz
				}

				if flipcoin() {
					zblok.rec.X -= siz * 2
					zblok.slideDIR = 4
					zblok.ro = 270
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					level[a].etc = append(level[a].etc, zblok)
				} else {
					zblok.rec.X -= siz * 2
				}

				if flipcoin() {
					zblok.rec.Y += siz
					zblok.rec.X += siz
					zblok.slideDIR = 3
					zblok.ro = 180
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					level[a].etc = append(level[a].etc, zblok)
				} else {
					zblok.rec.Y += siz
					zblok.rec.X += siz
				}

				if flipcoin() {
					zblok.rec.Y -= siz * 2
					zblok.slideDIR = 1
					zblok.ro = 0
					zblok.drec = zblok.rec
					zblok.drec.X += zblok.rec.Width / 2
					zblok.drec.Y += zblok.rec.Width / 2
					zblok.crec = zblok.rec
					level[a].etc = append(level[a].etc, zblok)
				} else {
					zblok.rec.Y -= siz * 2
				}
			}
		}
	}
}
func BlurRec(rec rl.Rectangle, dist float32) rl.Rectangle { //MARK:MAKE BLUR REC
	rec.X += rF32(-dist, dist)
	rec.Y += rF32(-dist, dist)
	return rec
}

func makeDrec(rec rl.Rectangle) rl.Rectangle { //MARK:MAKE DREC
	rec.X += rec.Width / 2
	rec.Y += rec.Height / 2
	return rec

}
func origin(rec rl.Rectangle) rl.Vector2 { //MARK:MAKE ORIGIN
	return rl.NewVector2(rec.Width/2, rec.Height/2)
}

func makesnow() { //MARK:MAKE SNOW

	num := rInt(50, 100)

	for {

		x := rF32(0, scrWF32)
		y := levY - rF32(bsU, scrHF32)
		siz := rF32(bsU, bsU3)

		zimg := ximg{}
		zimg.rec = rl.NewRectangle(x, y, siz, siz)
		zimg.cnt = rl.NewVector2(zimg.rec.X+zimg.rec.Width/2, zimg.rec.Y+zimg.rec.Height/2)
		choose := rInt(1, 4)
		switch choose {
		case 1:
			zimg.img = etc[15]
		case 2:
			zimg.img = etc[16]
		case 3:
			zimg.img = etc[17]
		}
		zimg.ori = rl.NewVector2(zimg.rec.Width/2, zimg.rec.Height/2)
		zimg.ro = rF32(0, 360)
		zimg.col = rl.White
		zimg.fade = rF32(0.3, 0.7)
		snow = append(snow, zimg)

		num--
		if num == 0 {
			break
		}
	}

}

func makeProjectile(name string) { //MARK:MAKE PROJECTILE

	zproj := xproj{}
	zproj.name = name
	zproj.cnt = pl.cnt
	zproj.onoff = true
	zproj.vel = bsU / 2
	zproj.fade = 1
	zproj.dmg = 1
	zproj.bounceN = mods.bounceN

	siz := bsU + bsU/2

	switch name {
	case "mrcarrot":
		siz = bsU2
		zproj.cnt = rl.NewVector2(mrcarrot.rec.X+mrcarrot.rec.Width/2, mrcarrot.rec.Y+mrcarrot.rec.Height/2)
		zproj.rec = rl.NewRectangle(mrcarrot.cnt.X-siz/2, mrcarrot.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.img = etc[50]
		zproj.col = rl.White
		zproj.vel = bsU / 4

		choose := rInt(1, 9)

		switch choose {
		case 1:
			zproj.velx = -zproj.vel
			zproj.vely = -zproj.vel
			plProj = append(plProj, zproj)
		case 2:
			zproj.vely = -zproj.vel
			plProj = append(plProj, zproj)
		case 3:
			zproj.velx = +zproj.vel
			zproj.vely = -zproj.vel
			plProj = append(plProj, zproj)
		case 4:
			zproj.velx = zproj.vel
			plProj = append(plProj, zproj)
		case 5:
			zproj.velx = +zproj.vel
			zproj.vely = +zproj.vel
			plProj = append(plProj, zproj)
		case 6:
			zproj.vely = +zproj.vel
			plProj = append(plProj, zproj)
		case 7:
			zproj.velx = -zproj.vel
			zproj.vely = +zproj.vel
			plProj = append(plProj, zproj)
		case 8:
			zproj.velx = -zproj.vel
			plProj = append(plProj, zproj)
		}

		choose2 := rInt(1, 9)
		for {
			choose = rInt(1, 9)
			if choose != choose2 {
				break
			}
		}

		switch choose2 {
		case 1:
			zproj.velx = -zproj.vel
			zproj.vely = -zproj.vel
			plProj = append(plProj, zproj)
		case 2:
			zproj.vely = -zproj.vel
			plProj = append(plProj, zproj)
		case 3:
			zproj.velx = +zproj.vel
			zproj.vely = -zproj.vel
			plProj = append(plProj, zproj)
		case 4:
			zproj.velx = zproj.vel
			plProj = append(plProj, zproj)
		case 5:
			zproj.velx = +zproj.vel
			zproj.vely = +zproj.vel
			plProj = append(plProj, zproj)
		case 6:
			zproj.vely = +zproj.vel
			plProj = append(plProj, zproj)
		case 7:
			zproj.velx = -zproj.vel
			zproj.vely = +zproj.vel
			plProj = append(plProj, zproj)
		case 8:
			zproj.velx = -zproj.vel
			plProj = append(plProj, zproj)
		}

	case "fireworks":
		siz = bsU2
		zproj.cnt = fireworksCnt
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.img = etc[49]
		zproj.col = rl.White

		zproj.velx = bsU / 2
		zproj.vely = -bsU / 2
		plProj = append(plProj, zproj)
		zproj.velx = bsU / 2
		zproj.vely = +bsU / 2
		zproj.ro = 90
		plProj = append(plProj, zproj)
		zproj.velx = -bsU / 2
		zproj.vely = -bsU / 2
		zproj.ro = 270
		plProj = append(plProj, zproj)
		zproj.velx = -bsU / 2
		zproj.vely = +bsU / 2
		zproj.ro = 180
		plProj = append(plProj, zproj)

	case "mralien":
		siz = bsU
		zproj.cnt = rl.NewVector2(mralien.rec.X+mralien.rec.Width/2, mralien.rec.Y+mralien.rec.Height/2)
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.img = plantBull.recTL
		zproj.col = ranGreen()

		zproj.velx = bsU / 4
		plProj = append(plProj, zproj)
		zproj.velx = -bsU / 4
		plProj = append(plProj, zproj)
		zproj.velx = 0

		zproj.vely = bsU / 4
		plProj = append(plProj, zproj)
		zproj.vely = -bsU / 4
		plProj = append(plProj, zproj)
		zproj.vely = 0

		zproj.velx = bsU / 4
		zproj.vely = -bsU / 4
		plProj = append(plProj, zproj)
		zproj.velx = bsU / 4
		zproj.vely = +bsU / 4
		plProj = append(plProj, zproj)
		zproj.velx = -bsU / 4
		zproj.vely = -bsU / 4
		plProj = append(plProj, zproj)
		zproj.velx = -bsU / 4
		zproj.vely = +bsU / 4
		plProj = append(plProj, zproj)

	case "plantbull":
		siz = bsU
		zproj.cnt = rl.NewVector2(mrplanty.rec.X+mrplanty.rec.Width/2, mrplanty.rec.Y+mrplanty.rec.Height/2)
		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		zproj.drec = zproj.rec
		zproj.img = plantBull.recTL
		zproj.col = ranGreen()
		if mrplanty.velx > 0 {
			zproj.velx = bsU / 4
		} else {
			zproj.velx = -bsU / 4
		}
		plProj = append(plProj, zproj)

	case "fireball":
		zproj.img = fireballPlayer.recTL
		zproj.col = ranOrange()

		switch pl.direc {
		case 1:
			zproj.vely = -zproj.vel
			zproj.ro = 270
		case 2:
			zproj.velx = zproj.vel
		case 3:
			zproj.vely = zproj.vel
			zproj.ro = 90
		case 4:
			zproj.velx = -zproj.vel
			zproj.ro = 180
		default:
			zproj.velx = zproj.vel
		}

		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		plProj = append(plProj, zproj)

		if mods.fireballN > 1 {
			zproj.col = ranOrange()
			if mods.fireballN >= 2 {
				zproj.ro += 180
				zproj.velx *= -1
				zproj.vely *= -1
				plProj = append(plProj, zproj)
			}
			if mods.fireballN >= 3 {
				zproj.col = ranOrange()
				zproj.ro += 90
				zproj.vely = 0
				zproj.velx = 0
				switch pl.direc {
				case 1:
					zproj.velx = zproj.vel
				case 2:
					zproj.vely = zproj.vel
				case 3:
					zproj.velx = -zproj.vel
				case 4:
					zproj.vely = -zproj.vel
				}
				plProj = append(plProj, zproj)
			}
			if mods.fireballN >= 4 {
				zproj.col = ranOrange()
				zproj.ro += 180
				zproj.velx *= -1
				zproj.vely *= -1
				plProj = append(plProj, zproj)
			}

			if mods.fireballN >= 5 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch pl.direc {
				case 1:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				case 2:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				case 3:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				case 4:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				}
				plProj = append(plProj, zproj)
			}

			if mods.fireballN >= 6 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch pl.direc {
				case 1:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				case 2:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				case 3:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				case 4:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				}
				plProj = append(plProj, zproj)
			}

			if mods.fireballN >= 7 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch pl.direc {
				case 1:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				case 2:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				case 3:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				case 4:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				}
				plProj = append(plProj, zproj)
			}

			if mods.fireballN >= 8 {
				zproj.col = ranOrange()
				zproj.ro = 0
				switch pl.direc {
				case 1:
					zproj.velx = -zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 225
				case 2:
					zproj.velx = zproj.vel
					zproj.vely = -zproj.vel
					zproj.ro = 315
				case 3:
					zproj.velx = zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 45
				case 4:
					zproj.velx = -zproj.vel
					zproj.vely = zproj.vel
					zproj.ro = 135
				}
				plProj = append(plProj, zproj)
			}

		}

	case "axe":
		zproj.img = etc[3]
		zproj.col = rl.SkyBlue

		for {
			zproj.velx = rF32(-zproj.vel, zproj.vel)
			zproj.vely = rF32(-zproj.vel, zproj.vel)
			if getabs(zproj.vely) > zproj.vel/2 || getabs(zproj.velx) > zproj.vel/2 {
				break
			}
		}

		zproj.rec = rl.NewRectangle(zproj.cnt.X-siz/2, zproj.cnt.Y-siz/2, siz, siz)
		zproj.drec = zproj.rec
		zproj.drec.X += zproj.rec.Width / 2
		zproj.drec.Y += zproj.rec.Height / 2
		zproj.ori = rl.NewVector2(zproj.rec.Width/2, zproj.rec.Height/2)
		plProj = append(plProj, zproj)
	}

}

func makespikes() { //MARK: MAKE SPIKES

	siz := bsU2
	zblok := xblok{}
	zblok.name = "spikes"
	zblok.fade = 1
	zblok.img = spikes.recTL

	for a := 0; a < len(level); a++ {
		if flipcoin() {
			zblok.cnt = findRanCntV2()
			num := rInt(10, 30)
			for num > 0 {
				zblok.color = ranBlue()
				zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
				v1 := rl.NewVector2(zblok.rec.X, zblok.rec.Y)
				v2 := v1
				v2.X += siz
				v3 := v2
				v3.Y += siz
				v4 := v1
				v4.Y += siz

				if rl.CheckCollisionPointRec(v1, levRecInner) && rl.CheckCollisionPointRec(v2, levRecInner) && rl.CheckCollisionPointRec(v3, levRecInner) && rl.CheckCollisionPointRec(v4, levRecInner) {
					level[a].spikes = append(level[a].spikes, zblok)
				}

				choose := rInt(1, 5)
				switch choose {
				case 1:
					zblok.cnt.Y -= siz
				case 2:
					zblok.cnt.X += siz
				case 3:
					zblok.cnt.Y += siz
				case 43:
					zblok.cnt.X -= siz
				}

				num--
			}
		}
	}

}
func maketurrets() { //MARK: MAKE TURRETS

	siz := bsU2
	zblok := xblok{}
	zblok.name = "turret"
	zblok.img = etc[18]
	zblok.solid = true
	zblok.fade = 1
	zblok.onoff = true

	for a := 0; a < len(level); a++ {

		if flipcoin() {
			num := rInt(1, 4)
			for num > 0 {
				for {
					zblok.cnt = findRanCntV2()
					zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
					if checkInnerBloksExits(a, zblok.rec) {
						break
					}
				}
				zblok.drec = zblok.rec
				zblok.drec.X += zblok.rec.Width / 2
				zblok.drec.Y += zblok.rec.Height / 2
				zblok.ori = rl.NewVector2(zblok.drec.Width, zblok.drec.Height)
				zblok.color = rl.White
				zblok.crec = zblok.rec
				zblok.ro = rF32(0, 360)
				zblok.timer = rI32(1, 3) * fps
				level[a].etc = append(level[a].etc, zblok)
				num--
			}
		}
	}

}

func makeFX(numType int, cnt rl.Vector2) { //MARK: MAKE FX

	zfx := xfx{}
	zfx.onoff = true
	zfx.timer = fps * 2
	zfx.cnt = cnt

	switch numType {
	case 4:
		zfx.name = "fxBurnOilBarrel"
		zfx.timer = fps * 4
		siz := bsU2
		cnt2 := cnt
		cnt2.X -= siz
		origX := cnt2.X
		cnt2.Y -= siz
		count := 0
		for a := 0; a < 9; a++ {
			if a != 4 {
				if flipcoin() {
					if rl.CheckCollisionPointRec(cnt2, levRecInner) {
						zfx.rec = rl.NewRectangle(cnt2.X-siz/2, cnt2.Y-siz/2, siz, siz)
						zfx.img = burn.recTL
						zfx.col = ranOrange()
						zfx.fade = rF32(0.7, 1.1)
						fx = append(fx, zfx)
					}
				}
			}
			cnt2.X += siz
			count++
			if count == 3 {
				count = 0
				cnt2.X = origX
				cnt2.Y += siz
			}
		}

		cnt2 = cnt
		cnt2.X -= siz * 2
		origX = cnt2.X
		cnt2.Y -= siz * 2
		count = 0
		for a := 0; a < 25; a++ {
			if a != 12 {
				if flipcoin() {
					if rl.CheckCollisionPointRec(cnt2, levRecInner) {
						zfx.rec = rl.NewRectangle(cnt2.X-siz/2, cnt2.Y-siz/2, siz, siz)
						zfx.img = burn.recTL
						zfx.col = ranOrange()
						zfx.fade = rF32(0.7, 1.1)
						fx = append(fx, zfx)
					}
				}
			}
			cnt2.X += siz
			count++
			if count == 5 {
				count = 0
				cnt2.X = origX
				cnt2.Y += siz
			}
		}

	case 3: //BURN WOOD BARREL
		zfx.name = "fxBurnWoodBarrel"
		zfx.timer = fps * 4
		siz := bsU2
		cnt2 := cnt

		num := rInt(1, 4)

		for num > 0 {
			countbreak := 20
			for {
				choose := rInt(1, 9)
				switch choose {
				case 1:
					cnt2.Y -= bsU2
					cnt2.X -= bsU2
				case 2:
					cnt2.Y -= bsU2
				case 3:
					cnt2.Y -= bsU2
					cnt2.X += bsU2
				case 4:
					cnt2.X += bsU2
				case 5:
					cnt2.Y += bsU2
					cnt2.X += bsU2
				case 6:
					cnt2.Y += bsU2
				case 7:
					cnt2.Y += bsU2
					cnt2.X -= bsU2
				case 8:
					cnt2.X -= bsU2
				}
				if rl.CheckCollisionPointRec(cnt2, levRecInner) || countbreak == 0 {
					break
				}
				countbreak--
			}

			zfx.rec = rl.NewRectangle(cnt2.X-siz/2, cnt2.Y-siz/2, siz, siz)
			zfx.img = burn.recTL
			zfx.col = ranOrange()
			zfx.fade = rF32(0.7, 1.1)

			fx = append(fx, zfx)
			num--
		}

	case 2: //ENEMY
		zfx.name = "fxEnemy"
		num := 20
		for {

			zrec := xrec{}
			wid := rF32(bsU, bsU3)
			zrec.rec = rl.NewRectangle(cnt.X-wid/2, cnt.Y-wid/2, wid, wid)
			zrec.velX = rF32(-bsU, bsU)
			zrec.velY = rF32(-bsU, bsU)
			zrec.col = ranRed()
			zrec.fade = rF32(0.7, 1.1)

			zfx.recs = append(zfx.recs, zrec)

			num--
			if num == 0 {
				break
			}
		}

		fx = append(fx, zfx)
	case 1: //BARREL
		zfx.name = "fxBarrel"
		num := 20
		for {

			zrec := xrec{}
			zrec.rec = rl.NewRectangle(cnt.X, cnt.Y, bsU/2, bsU/2)
			zrec.velX = rF32(-bsU, bsU)
			zrec.velY = rF32(-bsU, bsU)
			zrec.col = rl.Brown
			zrec.fade = 0.8

			zfx.recs = append(zfx.recs, zrec)

			num--
			if num == 0 {
				break
			}
		}

		fx = append(fx, zfx)

	}

}
func makeBlokGenNoRecNoCntr() xblok { //MARK:MAKE GENERIC BLOCK NO CENTER NO REC
	zblok := xblok{}
	zblok.fade = 1
	zblok.color = rl.White
	zblok.onoff = true

	return zblok
}
func makeBlokGeneric(siz float32, cnt rl.Vector2) xblok { //MARK:MAKE GENERIC BLOCK

	zblok := xblok{}
	zblok.fade = 1
	zblok.color = ranCol()
	zblok.cnt = cnt
	zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
	return zblok
}
func makeBlokGenRandom(siz float32) xblok { //MARK:MAKE GENERIC BLOCK RANDOM POSITION
	zblok := xblok{}
	zblok.fade = 1
	zblok.color = ranCol()

	x := levRecInner.X + bsU
	y := levRecInner.Y + bsU
	x = rF32(x, x+levRecInner.Width-bsU3)
	y = rF32(y, y+levRecInner.Width-bsU3)

	zblok.rec = rl.NewRectangle(x, y, siz, siz)
	zblok.cnt = rl.NewVector2(zblok.rec.X+zblok.rec.Width/2, zblok.rec.Y+zblok.rec.Height/2)
	zblok.crec = zblok.rec

	return zblok

}
func makeBlokGenericRanCentr(siz float32) xblok { //MARK:MAKE GENERIC BLOCK RANDOM CENTRE

	zblok := xblok{}
	zblok.fade = 1
	zblok.color = ranCol()
	for {
		zblok.cnt = findRanCntV2()
		if zblok.cnt.X+siz/2 < levRec.X+levRec.Width && zblok.cnt.Y+siz/2 < levRec.Y+levRec.Width {
			break
		}
	}

	zblok.rec = rl.NewRectangle(zblok.cnt.X-siz/2, zblok.cnt.Y-siz/2, siz, siz)
	zblok.crec = zblok.rec
	return zblok
}
func makeCnt(blok xblok) rl.Vector2 {
	v2 := rl.NewVector2(blok.rec.X+blok.rec.Width/2, blok.rec.Y+blok.rec.Height/2)
	return v2
}
func makemovebloks() { //MARK:MAKE MOVE BLOKS

	zblok := xblok{}
	bloksiz := bsU2
	blokimg := walltiles[rInt(0, len(walltiles))]
	zblok.fade = 1

	for a := 0; a < len(level); a++ {

		bloknum := rInt(0, 5)

		for b := 0; b < bloknum; b++ {
			addtolevel := false
			countbreak := 100
			for {

				zblok = makeBlokGenNoRecNoCntr()

				canadd := true
				zblok.rec, canadd = findRecPos(bloksiz, a)
				zblok.cnt = makeCnt(zblok)

				zblok.color = ranPink()
				zblok.img = blokimg
				zblok.velX = 0
				zblok.velY = 0
				zblok.vel = bsU / 2

				if len(level[a].movBloks) > 0 {
					for c := 0; c < len(level[a].movBloks); c++ {
						if rl.CheckCollisionRecs(zblok.rec, level[a].movBloks[c].rec) {
							canadd = false
						}
					}
				}
				if canadd {
					countbreak = 0
					addtolevel = true
				}
				countbreak--
				if countbreak <= 0 {
					break
				}
			}
			if addtolevel {
				choose := rInt(1, 3)
				switch choose {

				case 2: //UD
					zblok.velY = rF32(zblok.vel/4, zblok.vel)
					if roll12() == 12 {
						zblok.velX = rF32(zblok.vel/4, zblok.vel)
					}
					zblok.movType = 2 //UD
					level[a].movBloks = append(level[a].movBloks, zblok)
				case 1: //LR
					zblok.velX = rF32(zblok.vel/4, zblok.vel)
					if roll12() == 12 {
						zblok.velY = rF32(zblok.vel/4, zblok.vel)
					}
					zblok.movType = 1 //LR
					level[a].movBloks = append(level[a].movBloks, zblok)
				}
			}

		}

	}

}

func makeshaders() { //MARK:MAKE SHADERS
	shader = rl.LoadShader("", "shaders/bloom.fs")
	shader2 = rl.LoadShader("", "shaders/grayscale.fs")
	shader3 = rl.LoadShader("", "shaders/sobel.fs")
	renderTarget = rl.LoadRenderTexture(scrW32, scrH32)
}
func makecompanions() { //MARK: MAKE COMPANIONS

	//MR CARROT
	siz := bsU3
	mrcarrot.rec = rl.NewRectangle(cnt.X, cnt.Y, siz, siz)
	mrcarrot.imgl = rl.NewRectangle(0, 248, 38, 38)
	mrcarrot.imgr = rl.NewRectangle(228, 248, 38, 38)
	mrcarrot.vel = bsU / 5
	mrcarrot.velx = rF32(-mrcarrot.vel, mrcarrot.vel)
	mrcarrot.vely = rF32(-mrcarrot.vel, mrcarrot.vel)
	mrcarrot.hp = 5
	mrcarrot.hpmax = mrcarrot.hp
	mrcarrot.frames = 5
	mrcarrot.timer = fps * rI32(1, 5)

	//MR PLANTY
	mrplanty.rec = rl.NewRectangle(cnt.X, cnt.Y, siz, siz)
	mrplanty.imgl = rl.NewRectangle(0, 454, 44, 44)
	mrplanty.imgr = rl.NewRectangle(352, 454, 44, 44)
	mrplanty.vel = bsU / 5
	mrplanty.velx = rF32(-mrplanty.vel, mrplanty.vel)
	mrplanty.vely = rF32(-mrplanty.vel, mrplanty.vel)
	mrplanty.hp = 5
	mrplanty.hpmax = mrplanty.hp
	mrplanty.frames = 7

	//ALIEN
	mralien.rec = rl.NewRectangle(cnt.X, cnt.Y, siz, siz)
	mralien.img = alien[0]
	mralien.vel = bsU / 5
	mralien.velx = rF32(-mralien.vel, mralien.vel)
	mralien.vely = rF32(-mralien.vel, mralien.vel)
	mralien.hp = 5
	mralien.hpmax = mralien.hp
	mralien.timer = fps * rI32(3, 8)

}
func makefxinitial() { //MARK:MAKE FX INITAL

	//RAIN
	num := 300
	for a := 0; a < num; a++ {
		siz := rF32(2, 5)
		rec := rl.NewRectangle(rF32(0, scrWF32), rF32(-scrHF32, -bsU), siz, siz)
		rain = append(rain, rec)
	}

	//SCAN LINES
	y := float32(-2)
	x := float32(0)
	change := float32(3)

	for {
		scanlinev2 = append(scanlinev2, rl.NewVector2(x, y))
		y += change
		if y >= scrHF32+1 {
			break
		}
	}

}
func makebosses() { //MARK:MAKE BOSSES
	siz := bsU8
	zboss := xboss{}
	zboss.hp = 20
	zboss.timer = fps * 4
	zboss.hpmax = zboss.hp
	zboss.img = rl.NewRectangle(308, 1561, 48, 48)
	zboss.xl = zboss.img.X
	zboss.yt = zboss.img.Y
	zboss.vel = rF32(bsU/8, bsU/4)
	zboss.velX = rF32(-zboss.vel, zboss.vel)
	zboss.velY = rF32(-zboss.vel, zboss.vel)
	zboss.rec = rl.NewRectangle(cnt.X-siz/2, levRecInner.Y+bsU2, siz, siz)
	zboss.cnt = cnt
	zboss.crec = zboss.rec
	zboss.crec.X += zboss.rec.Width / 4
	zboss.crec.Y += zboss.rec.Height / 5
	zboss.crec.Width -= zboss.rec.Width / 2
	zboss.crec.Height -= zboss.rec.Height / 5
	zboss.atkType = rInt(1, 4)
	bosses = append(bosses, zboss)
	zboss.img = rl.NewRectangle(450, 1561, 48, 48)
	zboss.xl = zboss.img.X
	zboss.vel = rF32(bsU/8, bsU/4)
	zboss.velX = rF32(-zboss.vel, zboss.vel)
	zboss.velY = rF32(-zboss.vel, zboss.vel)
	zboss.atkType = rInt(1, 4)
	bosses = append(bosses, zboss)
	zboss.img = rl.NewRectangle(593, 1561, 48, 48)
	zboss.xl = zboss.img.X
	zboss.vel = rF32(bsU/8, bsU/4)
	zboss.velX = rF32(-zboss.vel, zboss.vel)
	zboss.velY = rF32(-zboss.vel, zboss.vel)
	zboss.atkType = rInt(1, 4)
	bosses = append(bosses, zboss)
}
func makeimgs() { //MARK:MAKE IMGS

	etc = append(etc, rl.NewRectangle(2, 36, 18, 18))       // 0 BARREL
	etc = append(etc, rl.NewRectangle(24, 36, 16, 16))      // 1 HP POTION
	etc = append(etc, rl.NewRectangle(42, 36, 16, 16))      // 2 PLAYER HP ICON
	etc = append(etc, rl.NewRectangle(60, 36, 14, 14))      // 3 THROWING AXE
	etc = append(etc, rl.NewRectangle(77, 36, 18, 18))      // 4 SANTA
	etc = append(etc, rl.NewRectangle(99, 36, 13, 13))      // 5 BOUNCE PROJECTILE
	etc = append(etc, rl.NewRectangle(115, 36, 16, 16))     // 6 ESCAPE VINE
	etc = append(etc, rl.NewRectangle(132, 35, 17, 17))     // 7 SKELETON KEY
	etc = append(etc, rl.NewRectangle(151, 35, 16, 16))     // 8 APPLE
	etc = append(etc, rl.NewRectangle(170, 36, 17, 17))     // 9 PLANT COMPANION
	etc = append(etc, rl.NewRectangle(192, 36, 18, 18))     // 10 MEDI KIT
	etc = append(etc, rl.NewRectangle(215, 38, 16, 16))     // 11 WALLET
	etc = append(etc, rl.NewRectangle(235, 38, 15, 15))     // 12 RECHARGE
	etc = append(etc, rl.NewRectangle(1082, 234, 32, 32))   // 13 SHOP1
	etc = append(etc, rl.NewRectangle(993, 156, 104, 104))  // 14 SHOP2
	etc = append(etc, rl.NewRectangle(1132, 228, 17, 17))   // 15 SNOW1
	etc = append(etc, rl.NewRectangle(1157, 229, 15, 15))   // 16 SNOW2
	etc = append(etc, rl.NewRectangle(1182, 228, 15, 15))   // 17 SNOW3
	etc = append(etc, rl.NewRectangle(254, 37, 16, 16))     // 18 TURRET
	etc = append(etc, rl.NewRectangle(274, 38, 14, 14))     // 19 NINJA STAR
	etc = append(etc, rl.NewRectangle(292, 36, 16, 16))     // 20 OIL BARREL
	etc = append(etc, rl.NewRectangle(344, 60, 15, 15))     // 21 SWITCH 1
	etc = append(etc, rl.NewRectangle(361, 60, 15, 15))     // 22 SWITCH 2
	etc = append(etc, rl.NewRectangle(445, 58, 17, 17))     // 23 CHEST
	etc = append(etc, rl.NewRectangle(717, 37, 12, 12))     // 24 SWITCH ARROW
	etc = append(etc, rl.NewRectangle(998, 414, 200, 200))  // 25 NIGHT REC
	etc = append(etc, rl.NewRectangle(731, 33, 18, 18))     // 26 SPINNING TRAP
	etc = append(etc, rl.NewRectangle(256, 16, 16, 16))     // 27 POWERUP BLOK
	etc = append(etc, rl.NewRectangle(820, 34, 18, 18))     // 28 MAP
	etc = append(etc, rl.NewRectangle(842, 32, 18, 18))     // 29 FIRE TRAIL
	etc = append(etc, rl.NewRectangle(863, 35, 16, 16))     // 30 INVISIBILITY
	etc = append(etc, rl.NewRectangle(883, 33, 17, 17))     // 31 COFFEE
	etc = append(etc, rl.NewRectangle(905, 34, 18, 18))     // 32 TELEPORT
	etc = append(etc, rl.NewRectangle(927, 35, 16, 16))     // 33 ATTACK RANGE
	etc = append(etc, rl.NewRectangle(945, 35, 16, 16))     // 34 ATTACK DAMAGE
	etc = append(etc, rl.NewRectangle(587, 58, 15, 15))     // 35 ORBITAL
	etc = append(etc, rl.NewRectangle(607, 58, 16, 16))     // 36 CHAIN LIGHTNING
	etc = append(etc, rl.NewRectangle(624, 57, 16, 16))     // 37 HP RING
	etc = append(etc, rl.NewRectangle(641, 57, 16, 16))     // 38 ARMOR SHIELD
	etc = append(etc, rl.NewRectangle(661, 56, 18, 18))     // 39 ANCHOR PAUSE ENEMY
	etc = append(etc, rl.NewRectangle(681, 55, 18, 18))     // 40 UMBRELLA RAIN
	etc = append(etc, rl.NewRectangle(705, 57, 14, 14))     // 41 DASH
	etc = append(etc, rl.NewRectangle(724, 56, 18, 18))     // 42 SOCKS POISON GAS
	etc = append(etc, rl.NewRectangle(748, 56, 17, 17))     // 43 CHERRIES
	etc = append(etc, rl.NewRectangle(770, 58, 15, 15))     // 44 FISH FLOOD
	etc = append(etc, rl.NewRectangle(789, 57, 18, 18))     // 45 CAKE BIRTHDAY
	etc = append(etc, rl.NewRectangle(811, 58, 16, 16))     // 46 PEACE NO DAMAGE
	etc = append(etc, rl.NewRectangle(831, 58, 16, 16))     // 47 ALIEN SPECIAL ENEMY
	etc = append(etc, rl.NewRectangle(850, 57, 16, 16))     // 48 AIR STRIKE
	etc = append(etc, rl.NewRectangle(893, 58, 14, 14))     // 49 POWERUP FIREWORKS
	etc = append(etc, rl.NewRectangle(909, 58, 14, 14))     // 50 CARROT COMPANION
	etc = append(etc, rl.NewRectangle(870, 55, 18, 18))     // 51 MARIO PLATFORMER LEVEL
	etc = append(etc, rl.NewRectangle(1132, 80, 64, 64))    // 52 FOOTPRINTS
	etc = append(etc, rl.NewRectangle(966, 35, 14, 14))     // 53 STAIRCASE
	etc = append(etc, rl.NewRectangle(41, 1219, 565, 300))  // 54 BITTY KNIGHT LOGO
	etc = append(etc, rl.NewRectangle(748, 1214, 308, 305)) // 55 GO LOGO
	etc = append(etc, rl.NewRectangle(10, 1532, 256, 256))  // 56 RAYLIB
	etc = append(etc, rl.NewRectangle(938, 104, 20, 20))    // 57 BOSS 3 PROJ
	etc = append(etc, rl.NewRectangle(841, 818, 128, 128))  // 58 POISON GAS

	plants = append(plants, rl.NewRectangle(315, 36, 13, 13))
	plants = append(plants, rl.NewRectangle(329, 38, 11, 11))
	plants = append(plants, rl.NewRectangle(341, 37, 12, 12))
	plants = append(plants, rl.NewRectangle(354, 31, 18, 18))
	plants = append(plants, rl.NewRectangle(373, 31, 18, 18))
	plants = append(plants, rl.NewRectangle(392, 33, 16, 16))
	plants = append(plants, rl.NewRectangle(409, 33, 16, 16))
	plants = append(plants, rl.NewRectangle(425, 33, 16, 16))
	plants = append(plants, rl.NewRectangle(442, 35, 14, 14))
	plants = append(plants, rl.NewRectangle(457, 35, 14, 14))

	x := float32(0)
	y := float32(0)

	//FLOOR
	for {
		floortiles = append(floortiles, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 896 {
			break
		}
	}
	//INNER WALLS
	x = 0
	y = 16
	for {
		walltiles = append(walltiles, rl.NewRectangle(x, y, 16, 16))
		x += 16
		if x >= 256 {
			break
		}
	}
	//BATS
	x = 1
	y = 351
	for {
		bats = append(bats, rl.NewRectangle(x, y, 16, 16))
		y += 16
		if y > 383 {
			break
		}
	}
	//KNIGHT MOVE
	x = 935
	pl.imgWalkX = x
	y = 268
	for a := 0; a < 4; a++ {
		knight = append(knight, rl.NewRectangle(x, y, 32, 32))
		y += 32
	}
	//KNIGHT ATTACK
	x = 735
	pl.imgAtkX = x
	y = 268
	for a := 0; a < 4; a++ {
		knight = append(knight, rl.NewRectangle(x, y, 32, 32))
		y += 32
	}
	//BUILDINGS
	x = 0
	y = 132
	for {
		shrines = append(shrines, rl.NewRectangle(x, y, 32, 32))
		x += 32
		if x > 64 {
			break
		}
	}
	//SKULLS
	x = 474
	y = 34
	for a := 0; a < 7; a++ {
		skulls = append(skulls, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//CANDLES
	x = 586
	y = 34
	for a := 0; a < 3; a++ {
		candles = append(candles, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//SIGNS
	x = 634
	y = 34
	for a := 0; a < 5; a++ {
		signs = append(signs, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//SPLATS
	x = 0
	y = 1072
	for a := 0; a < 6; a++ {
		splats = append(splats, rl.NewRectangle(x, y, 128, 128))
		x += 128
	}
	//STATUES
	x = 0
	y = 175
	for a := 0; a < 16; a++ {
		statues = append(statues, rl.NewRectangle(x, y, 48, 48))
		x += 48
	}
	//MUSHROOMS
	x = 751
	y = 34
	for a := 0; a < 4; a++ {
		mushrooms = append(mushrooms, rl.NewRectangle(x, y, 16, 16))
		x += 16
	}
	//PATTERNS
	x = 2
	y = 798
	for a := 0; a < 6; a++ {
		patterns = append(patterns, rl.NewRectangle(x, y, 128, 128))
		patterns = append(patterns, rl.NewRectangle(x, y+128, 128, 128))
		x += 128
	}
	//GEMS
	x = 0
	y = 297
	for a := 0; a < 8; a++ {
		gems = append(gems, rl.NewRectangle(x, y, 32, 32))
		x += 32
	}

	//RABBIT1
	rabbit1.frames = 3
	rabbit1.xl = 68
	rabbit1.yt = 335
	rabbit1.W = 16
	rabbit1.recTL = rl.NewRectangle(rabbit1.xl, rabbit1.yt, rabbit1.W, rabbit1.W)

	//FIREBALL PLAYER
	fireballPlayer.frames = 3
	fireballPlayer.xl = 1
	fireballPlayer.yt = 60
	fireballPlayer.W = 16
	fireballPlayer.recTL = rl.NewRectangle(fireballPlayer.xl, fireballPlayer.yt, fireballPlayer.W, fireballPlayer.W)

	//BURN
	burn.frames = 3
	burn.xl = 72
	burn.yt = 60
	burn.W = 16
	burn.recTL = rl.NewRectangle(burn.xl, burn.yt, burn.W, burn.W)

	//STAR
	star.frames = 3
	star.xl = 144
	star.yt = 60
	star.W = 16
	star.recTL = rl.NewRectangle(star.xl, star.yt, star.W, star.W)

	//WATER
	wateranim.frames = 3
	wateranim.xl = 208
	wateranim.yt = 60
	wateranim.W = 16
	wateranim.recTL = rl.NewRectangle(wateranim.xl, wateranim.yt, wateranim.W, wateranim.W)

	//PLANT BULL
	plantBull.frames = 4
	plantBull.xl = 276
	plantBull.yt = 60
	plantBull.W = 16
	plantBull.recTL = rl.NewRectangle(plantBull.xl, plantBull.yt, plantBull.W, plantBull.W)

	//SPIKES
	spikes.frames = 13
	spikes.xl = 0
	spikes.yt = 90
	spikes.W = 32
	spikes.recTL = rl.NewRectangle(spikes.xl, spikes.yt, spikes.W, spikes.W)

	//SPRING
	spring.frames = 2
	spring.xl = 382
	spring.yt = 59
	spring.W = 16
	spring.recTL = rl.NewRectangle(spring.xl, spring.yt, spring.W, spring.W)

	//POISON GAS
	posiongas.frames = 3
	posiongas.xl = 686
	posiongas.yt = 638
	posiongas.W = 128
	posiongas.recTL = rl.NewRectangle(posiongas.xl, posiongas.yt, posiongas.W, posiongas.W)

	//MUSHROOM BULL
	mushBull.frames = 3
	mushBull.xl = 516
	mushBull.yt = 59
	mushBull.W = 16
	mushBull.recTL = rl.NewRectangle(mushBull.xl, mushBull.yt, mushBull.W, mushBull.W)

	//BLADES
	blades.frames = 7
	blades.xl = 456
	blades.yt = 92
	blades.W = 32
	blades.recTL = rl.NewRectangle(blades.xl, blades.yt, blades.W, blades.W)

	//SPEAR
	spear.frames = 7
	spear.xl = 1004
	spear.yt = 5
	spear.W = 16
	spear.recTL = rl.NewRectangle(spear.xl, spear.yt, spear.W, 64)

	//FIRETRAIL
	firetrailanim.frames = 3
	firetrailanim.xl = 724
	firetrailanim.yt = 82
	firetrailanim.W = 16
	firetrailanim.recTL = rl.NewRectangle(firetrailanim.xl, firetrailanim.yt, firetrailanim.W, 16)

	//ORBITAL
	orbitalanim.frames = 5
	orbitalanim.xl = 800
	orbitalanim.yt = 84
	orbitalanim.W = 16
	orbitalanim.recTL = rl.NewRectangle(orbitalanim.xl, orbitalanim.yt, orbitalanim.W, 16)

	//FLOOD
	floodanim.frames = 3
	floodanim.xl = 900
	floodanim.yt = 82
	floodanim.W = 16
	floodanim.recTL = rl.NewRectangle(floodanim.xl, floodanim.yt, floodanim.W, 16)
	floodImg = floodanim.recTL

	//FISH R
	fishR.frames = 3
	fishR.xl = 724
	fishR.yt = 105
	fishR.W = 16
	fishR.recTL = rl.NewRectangle(fishR.xl, fishR.yt, fishR.W, 16)

	//FISH L
	fishL.frames = 3
	fishL.xl = 794
	fishL.yt = 105
	fishL.W = 16
	fishL.recTL = rl.NewRectangle(fishL.xl, fishL.yt, fishL.W, 16)

	//AIRSTRIKE
	airstrikeanim.frames = 3
	airstrikeanim.xl = 664
	airstrikeanim.yt = 614
	airstrikeanim.W = 32
	airstrikeanim.recTL = rl.NewRectangle(airstrikeanim.xl, airstrikeanim.yt, airstrikeanim.W, 32)

	//BOSS1 PROJ
	boss1anim.frames = 3
	boss1anim.xl = 977
	boss1anim.yt = 83
	boss1anim.W = 16
	boss1anim.recTL = rl.NewRectangle(boss1anim.xl, boss1anim.yt, boss1anim.W, 16)

	//BOSS2 PROJ
	boss2anim.frames = 3
	boss2anim.xl = 866
	boss2anim.yt = 107
	boss2anim.W = 16
	boss2anim.recTL = rl.NewRectangle(boss2anim.xl, boss2anim.yt, boss2anim.W, 16)

	//ALIEN
	x = 1008
	y = 770

	for {
		alien = append(alien, rl.NewRectangle(x, y, 64, 64))
		x += 64
		if x >= 1200 {
			x = 1008
			y += 64
		}
		if y >= 1026 {
			break
		}
	}

}

// MARK: CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE CORE
func inp() { //MARK:INP

	//OPTIONS ON

	if rl.IsKeyPressed(rl.KeyEscape) && !intro && !marioon && !died && !nextlevelscreen {

		if optionson && !exiton && !shopon && !levMapOn && !invenon && !creditson && !helpon && !timeson {
			if optionsChange {
				savesettings()
			}
			optionson = false
			creditson = false
			helpon = false
			exiton = false
			pause = false
		} else if optionson && exiton {
			exiton = false
		} else if !optionson && !shopon && !levMapOn && !invenon && !creditson && !helpon && !timeson {
			pause = true
			creditson = false
			helpon = false
			exiton = false
			optionson = true
			optionsChange = false
			optionnum = 0
		} else if !optionson && shopon && !levMapOn && !invenon && !creditson && !helpon && !timeson {
			shopon = false
			pl.cnt.Y = shopExitY
			upPlayerRec()
			pause = false
		} else if !optionson && !shopon && levMapOn && !invenon && !creditson && !helpon && !timeson {
			levMapOn = false
			pause = false
		} else if !optionson && !shopon && !levMapOn && invenon && !creditson && !helpon && !timeson {
			invenon = false
			pause = false
		} else if optionson && !shopon && !levMapOn && !invenon && creditson && !helpon && !timeson {
			creditson = false
			optionson = true
		} else if optionson && !shopon && !levMapOn && !invenon && !creditson && helpon && !timeson {
			helpon = false
		} else if !optionson && !shopon && !levMapOn && !invenon && !creditson && !helpon && timeson {
			timeson = false
			restartgame()
		}

	} else if intro {
		if rl.IsKeyPressed(rl.KeyEscape) || rl.IsGamepadButtonPressed(0, rl.GamepadButtonMiddleRight) {
			if optionson {
				if optionsChange {
					savesettings()
				}
				optionson = false
			} else {
				optionson = true
				optionsChange = false
				optionnum = 0
			}
		}

		if optionson && rl.IsGamepadButtonPressed(0, 6) {
			if optionsChange {
				savesettings()
			}
			optionson = false
		}
	}
	if useController {
		if rl.IsGamepadButtonPressed(0, rl.GamepadButtonMiddleRight) && !invenon && !shopon && !levMapOn && !intro && !marioon && !died && !nextlevelscreen {
			if optionson {
				if optionsChange {
					savesettings()
				}
				creditson = false
				helpon = false
				exiton = false
				optionson = false
				pause = false
			} else {
				creditson = false
				helpon = false
				exiton = false
				optionsChange = false
				optionson = true
				optionnum = 0
				pause = true
			}
		}
	}

	if !intro && !marioon && !died {

		//INVEN ON
		if rl.IsKeyPressed(rl.KeyTab) && !optionson && !levMapOn && !shopon && !nextlevelscreen {
			if invenon {
				invenon = false
				pause = false
			} else {
				invenon = true
				pause = true
			}
		}

		//MAP ON
		if rl.IsKeyPressed(rl.KeyRightControl) && !optionson && !invenon && !shopon && !nextlevelscreen {

			if levMapOn {
				levMapOn = false
				pause = false
			} else {
				levMapOn = true
				pause = true
			}
		}
		if useController {
			if rl.IsGamepadButtonPressed(0, 6) && !optionson && !invenon && !shopon && !timeson && !nextlevelscreen && !helpon && !creditson {
				if levMapOn {
					levMapOn = false
					pause = false
				} else {
					levMapOn = true
					pause = true
				}
			} else if rl.IsGamepadButtonPressed(0, 6) && optionson && !shopon && !timeson && !nextlevelscreen && !helpon && !creditson {
				if optionsChange {
					savesettings()
				}
				optionson = false
				pause = false

			} else if rl.IsGamepadButtonPressed(0, 6) && invenon && !shopon && !timeson && !nextlevelscreen && !helpon && !creditson {
				invenon = false
				pause = false
			} else if rl.IsGamepadButtonPressed(0, 6) && optionson && !invenon && !shopon && timeson && !nextlevelscreen && !helpon && !creditson {
				timeson = false
			} else if rl.IsGamepadButtonPressed(0, 6) && optionson && !invenon && !shopon && !timeson && !nextlevelscreen && helpon && !creditson {
				helpon = false
			} else if rl.IsGamepadButtonPressed(0, 6) && optionson && !invenon && !shopon && !timeson && !nextlevelscreen && !helpon && creditson {
				creditson = false
			}
			//INVEN ON
			if rl.IsGamepadButtonPressed(0, 5) && !optionson && !levMapOn && !shopon && !timeson && !nextlevelscreen && !helpon && !creditson {
				if invenon {
					invenon = false
					pause = false
				} else {
					invenon = true
					pause = true
				}
			}
		}

		//PLAYER
		if !pause {
			if !pl.escape {
				if rl.IsKeyPressed(rl.KeySpace) {
					if pl.atkTimer == 0 {
						rl.PlaySound(sfx[0])
						chainLightingSwingOnOff = false
						pl.atk = true
						pl.atkTimer = fps / 3
						pl.img.X = pl.imgAtkX
						if mods.fireball {
							makeProjectile("fireball")
						}
					}
					keypressT = fps / 2
				}
				pl.move = false

				if !pl.slide {
					if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
						if checkplayermove(1) {
							pl.cnt.Y -= pl.vel
						}
						pl.direc = 1
						if pl.atkTimer == 0 {
							pl.move = true
						}
					}
					if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
						if checkplayermove(3) {
							pl.cnt.Y += pl.vel
						}
						pl.direc = 3
						if pl.atkTimer == 0 {
							pl.move = true
						}
					}
					if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
						if checkplayermove(4) {
							pl.cnt.X -= pl.vel
						}
						pl.direc = 4
						if pl.atkTimer == 0 {
							pl.move = true
						}
					}
					if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
						if checkplayermove(2) {
							pl.cnt.X += pl.vel
						}
						pl.direc = 2
						if pl.atkTimer == 0 {
							pl.move = true
						}
					}
				}

				//CONTROLLER
				if useController {
					if rl.IsGamepadButtonPressed(0, 7) || rl.IsGamepadButtonPressed(0, 12) {
						if pl.atkTimer == 0 {
							rl.PlaySound(sfx[0])
							chainLightingSwingOnOff = false
							pl.atk = true
							pl.atkTimer = fps / 3
							pl.img.X = pl.imgAtkX
							if mods.fireball {
								makeProjectile("fireball")
							}
						}
						keypressT = fps / 2
					}
					pl.move = false

					if !pl.slide {

						if rl.GetGamepadAxisMovement(0, 1) < 0 || rl.IsGamepadButtonDown(0, 1) {
							if checkplayermove(1) {
								pl.cnt.Y -= pl.vel
							}
							pl.direc = 1
							if pl.atkTimer == 0 {
								pl.move = true
							}
						}
						if rl.GetGamepadAxisMovement(0, 1) > 0 || rl.IsGamepadButtonDown(0, 3) {
							if checkplayermove(3) {
								pl.cnt.Y += pl.vel
							}
							pl.direc = 3
							if pl.atkTimer == 0 {
								pl.move = true
							}
						}

						if rl.GetGamepadAxisMovement(0, 0) < 0 || rl.IsGamepadButtonDown(0, 4) {
							if checkplayermove(4) {
								pl.cnt.X -= pl.vel
							}
							pl.direc = 4
							if pl.atkTimer == 0 {
								pl.move = true
							}
						}
						if rl.GetGamepadAxisMovement(0, 0) > 0 || rl.IsGamepadButtonDown(0, 2) {
							if checkplayermove(2) {
								pl.cnt.X += pl.vel
							}
							pl.direc = 2
							if pl.atkTimer == 0 {
								pl.move = true
							}
						}

					}
				}

				upPlayerRec()
			}
		}
	}

	/*
		//DEBUG
		if rl.IsKeyPressed(rl.KeyF1) {
			if debug {
				debug = false
			} else {
				debug = true
			}
		}

		//ZOOM
		if rl.IsKeyPressed(rl.KeyKpAdd) {
			if cam2.Zoom == 1 {
				cam2.Zoom = 2
			} else if cam2.Zoom == 2 {
				cam2.Zoom = 3
			} else if cam2.Zoom == 3 {
				cam2.Zoom = 4
			} else if cam2.Zoom == 4 {
				cam2.Zoom = 1
			}
			cams()
		}
		if rl.IsKeyPressed(rl.KeyKpSubtract) {
			if cam2.Zoom == 1 {
				cam2.Zoom = 4
			} else if cam2.Zoom == 2 {
				cam2.Zoom = 1
			} else if cam2.Zoom == 3 {
				cam2.Zoom = 2
			} else if cam2.Zoom == 4 {
				cam2.Zoom = 3
			}
			cams()
		}
	*/

}

func timers() { //MARK:TIMERS

	if shopExitT > 0 {
		shopExitT--
	}

	runT++
	if runT%fps == 0 {
		secs++
		runT = 0
	}
	if secs == 60 {
		secs = 0
		mins++
	}

	if anchorT > 0 {
		anchorT--
	}

	if roomChangedTimer > 0 {
		roomChangedTimer--
		if roomChangedTimer == 1 {
			roomChanged = false
		}
	}

}
func cams() { //MARK:CAMS

	cam2.Target = cnt

	cam2.Offset.X = scrWF32 / 2
	cam2.Offset.Y = scrHF32 / 2

	if flipcam && !optionson {
		cam2.Rotation = 180
	} else if flipcam && optionson {
		cam2.Rotation = 0
	} else {
		cam2.Rotation = 0
	}

}

func initialWindow() { //MARK:INITIAL WINDOW
	rl.SetExitKey(rl.KeyEnd)
	rl.HideCursor()
	imgs = rl.LoadTexture("img/imgs.png")
	renderTarget = rl.LoadRenderTexture(scrW32, scrH32)

	pause = true
	intro = true
	//shaderon = true

	//endgame = true

	hpBarsOn = true
	makesettings()
	makeshaders()
	makeaudio()
	makefxinitial()
	makeimgs()
	makebosses()
	makelevel()
	makecompanions()
	makeplayer()
	maketimes()
	cams()

	music = backMusic[bgMusicNum]
	music.Looping = true
	rl.SetMasterVolume(4) //CHANGE VOLUME
	rl.SetMusicVolume(music, volume)
	upvolume()

	diedRec = rl.NewRectangle(cnt.X-bsU, cnt.Y-bsU, bsU2, bsU2)
	diedIMG = splats[rInt(0, len(splats))]
	endgameT = fps * 5
	endgopherrec = rl.NewRectangle(cnt.X-bsU4, levRec.Y+levRec.Height, bsU8, bsU8)

}

func unload() { //MARK:UNLOAD
	rl.UnloadShader(shader)
	rl.UnloadShader(shader2)
	rl.UnloadRenderTexture(renderTarget)
	rl.UnloadTexture(imgs)

	rl.StopMusicStream(music)
	rl.UnloadMusicStream(music)
	rl.UnloadMusicStream(backMusic[bgMusicNum])

	for a := 0; a < len(sfx); a++ {
		rl.UnloadSound(sfx[a])
	}
	rl.CloseAudioDevice()

}

func main() { //MARK:MAIN

	rl.SetConfigFlags(rl.FlagMsaa4xHint)
	//rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetTraceLogLevel(rl.LogError)

	rl.InitWindow(0, 0, "BITTY KNIGHT - unklnik.com") //GET SCREEN SIZE
	rl.InitAudioDevice()

	rl.SetWindowState(rl.FlagBorderlessWindowedMode)
	scrW, scrH = rl.GetScreenWidth(), rl.GetScreenHeight()

	//rl.ToggleFullscreen()

	rl.SetWindowSize(scrW, scrH)

	scrW32, scrH32 = int32(scrW), int32(scrH)
	scrWF32, scrHF32 = float32(scrW), float32(scrH)
	cnt = rl.NewVector2(scrWF32/2, scrHF32/2)
	if scrH >= 2160 {
		cam2.Zoom = 3
	} else if scrH >= 1440 && scrH < 2160 {
		cam2.Zoom = 2
	} else if scrH == 1200 {
		cam2.Zoom = 1.65
	} else if scrH > 1050 && scrH < 1200 {
		cam2.Zoom = 1.5
	} else if scrH >= 990 && scrH <= 1050 {
		cam2.Zoom = 1.35
	} else if scrH >= 900 && scrH < 990 {
		cam2.Zoom = 1.2
	} else if scrH >= 720 && scrH < 900 {
		cam2.Zoom = 1
	} else if scrH >= 600 && scrH < 720 {
		cam2.Zoom = 0.8
	} else if scrH < 600 && scrH > 300 {
		cam2.Zoom = 0.5
	} else if scrH < 300 {
		cam2.Zoom = 0.2
	}

	levX = cnt.X - levW/2
	levY = cnt.Y - levW/2
	levRec = rl.NewRectangle(levX, levY, levW, levW)
	levRecInner = levRec
	levRecInner.X += borderWallBlokSiz
	levRecInner.Y += borderWallBlokSiz
	levRecInner.Width -= borderWallBlokSiz * 2
	levRecInner.Height -= borderWallBlokSiz * 2

	initialWindow() //INITIAL INSIDE WINDOW

	rl.SetTargetFPS(fps)

	for !rl.WindowShouldClose() {
		frames++
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if shaderon && !shader2on && !shader3on { //BLOOM SHADER
			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginShaderMode(shader)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()

		} else if shaderon && shader2on && !shader3on {
			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.BeginShaderMode(shader)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
			rl.EndTextureMode()
			rl.BeginShaderMode(shader2)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else if shaderon && !shader2on && shader3on {
			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.BeginShaderMode(shader)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
			rl.EndTextureMode()
			rl.BeginShaderMode(shader3)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else if !shaderon && shader2on && !shader3on {
			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginShaderMode(shader2)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else if !shaderon && !shader2on && shader3on {
			rl.BeginTextureMode(renderTarget) // Enable drawing to texture
			rl.ClearBackground(rl.Black)
			drawnocamBG()

			rl.BeginMode2D(cam2)

			drawcam()

			rl.EndMode2D()

			drawnocam()

			rl.EndTextureMode()

			rl.BeginShaderMode(shader3)
			rl.DrawTextureRec(renderTarget.Texture, rl.NewRectangle(0, 0, float32(renderTarget.Texture.Width), float32(-renderTarget.Texture.Height)), rl.NewVector2(0, 0), rl.White)
			rl.EndShaderMode()
		} else { //NO BLOOM SHADER
			rl.ClearBackground(rl.Black)
			drawnocamBG()
			rl.BeginMode2D(cam2)
			drawcam()
			rl.EndMode2D()
			drawnocam()
		}

		drawnoRender()

		up() //UPDATE
		rl.EndDrawing()
	}

	unload() //UNLOAD

	rl.CloseAudioDevice()

	rl.CloseWindow()
}

// MARK:HELPERS
func absdiff(num1, num2 float32) float32 {
	num := float32(0)
	if num1 == num2 {
		num = 0
	} else {
		if num1 <= 0 && num2 <= 0 {
			num1 = getabs(num1)
			num2 = getabs(num2)
			if num1 > num2 {
				num = num1 - num2
			} else {
				num = num2 - num1
			}
		} else if num1 <= 0 && num2 >= 0 {
			num = num2 + getabs(num1)
		} else if num2 <= 0 && num1 >= 0 {
			num = num1 + getabs(num2)
		} else if num2 >= 0 && num1 >= 0 {
			if num1 > num2 {
				num = num1 - num2
			} else {
				num = num2 - num1
			}
		}
	}
	return num
}
func getabs(value float32) float32 {
	value2 := float64(value)
	value = float32(math.Abs(value2))
	return value
}
func remImg(slice []ximg, s int) []ximg {
	return append(slice[:s], slice[s+1:]...)
}
func remEnemy(slice []xenemy, s int) []xenemy {
	return append(slice[:s], slice[s+1:]...)
}
func remTxt(slice []xtxt, s int) []xtxt {
	return append(slice[:s], slice[s+1:]...)
}
func remProj(slice []xproj, s int) []xproj {
	return append(slice[:s], slice[s+1:]...)
}
func remFX(slice []xfx, s int) []xfx {
	return append(slice[:s], slice[s+1:]...)
}
func remBlok(slice []xblok, s int) []xblok {
	return append(slice[:s], slice[s+1:]...)
}
func flipcoin() bool {
	onoff := false
	choose := rInt(0, 100001)
	if choose > 50000 {
		onoff = true
	}
	return onoff
}
func roll6() int {
	return rInt(1, 7)
}
func roll12() int {
	return rInt(1, 13)
}
func roll18() int {
	return rInt(1, 19)
}
func roll24() int {
	return rInt(1, 25)
}
func roll30() int {
	return rInt(1, 31)
}
func roll36() int {
	return rInt(1, 37)
}
func rInt(min, max int) int {
	return min + rand.Intn(max-min)
}
func rI32(min, max int) int32 {
	return int32(min + rand.Intn(max-min))
}
func rF32(min, max float32) float32 {
	min2 := float64(min)
	max2 := float64(max)
	return float32(min2 + rand.Float64()*(max2-min2))
}

// MARK:COLORS
func brightYellow() rl.Color {
	return rl.NewColor(uint8(255), uint8(240), uint8(31), 255)
}
func ranCol() rl.Color {
	return rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
}
func ranDarkGreen() rl.Color {
	return rl.NewColor(uint8(rInt(0, 30)), uint8(rInt(40, 90)), uint8(rInt(0, 40)), 255)
}
func ranGreen() rl.Color {
	return rl.NewColor(uint8(rInt(0, 60)), uint8(rInt(140, 256)), uint8(rInt(0, 60)), 255)
}
func ranRed() rl.Color {
	return rl.NewColor(uint8(rInt(140, 256)), uint8(rInt(0, 60)), uint8(rInt(0, 60)), 255)
}
func ranPink() rl.Color {
	return rl.NewColor(uint8(rInt(200, 256)), uint8(rInt(10, 110)), uint8(rInt(130, 180)), 255)
}
func ranBlue() rl.Color {
	return rl.NewColor(uint8(rInt(0, 60)), uint8(rInt(0, 60)), uint8(rInt(140, 256)), 255)
}
func ranDarkBlue() rl.Color {
	return rl.NewColor(uint8(rInt(0, 20)), uint8(rInt(0, 20)), uint8(rInt(100, 160)), 255)
}
func ranCyan() rl.Color {
	return rl.NewColor(uint8(rInt(0, 120)), uint8(rInt(200, 256)), uint8(rInt(150, 256)), 255)
}
func ranYellow() rl.Color {
	return rl.NewColor(uint8(rInt(245, 256)), uint8(rInt(200, 256)), uint8(rInt(0, 100)), 255)
}
func ranOrange() rl.Color {
	return rl.NewColor(uint8(255), uint8(rInt(70, 170)), uint8(rInt(0, 50)), 255)
}
func ranBrown() rl.Color {
	return rl.NewColor(uint8(rInt(100, 150)), uint8(rInt(50, 120)), uint8(rInt(30, 90)), 255)
}
func ranGrey() rl.Color {
	return rl.NewColor(uint8(rInt(170, 220)), uint8(rInt(170, 220)), uint8(rInt(170, 220)), 255)
}
func ranDarkGrey() rl.Color {
	return rl.NewColor(uint8(rInt(90, 120)), uint8(rInt(90, 120)), uint8(rInt(90, 120)), 255)
}
func darkRed() rl.Color {
	return rl.NewColor(uint8(139), uint8(0), uint8(0), 255)
}
