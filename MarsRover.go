package main


import (
	"fmt"
)

type MarsRover struct{
	
	North int
	South int
	East int
	West int

	HeadingDirection int
	PosX int
	PosY int
	

}


func (mR *MarsRover) curPos(x,y,North int){
	//fmt.Println("x...", x,y,North)
	mR.PosX = x
	mR.PosY = y
	mR.HeadingDirection = North
	
}

func (mR *MarsRover) explore(instructions string){
	for _,inst := range instructions {
		//fmt.Println("Explore....", inst)
		
		if inst=='L' {
			//fmt.Println("explore...", string(inst))
			//spin 90 Degrees Left
			if (mR.HeadingDirection-1) < mR.North {
				mR.HeadingDirection = mR.West
	
			}else{
				mR.HeadingDirection--

			}
		}else if inst=='M'{
			//move forward one grid point
			switch mR.HeadingDirection {
				case  mR.West  : mR.PosX--
				case  mR.East  : mR.PosX++
				case  mR.South : mR.PosY--
				case  mR.North : mR.PosY++
				default: fmt.Println("Heading Direction - invalid input")
			}
		}else if inst=='R'{
			fmt.Println("explore...", string(inst))
			//spin 90 Degrees Right
			if (mR.HeadingDirection-1) > mR.West {
				mR.HeadingDirection = mR.North
	
			}else{
				mR.HeadingDirection++

			}
		}
		
	}

}

func (mR *MarsRover) OUTPUT_curPosAfterExplored(){
	headingDirection := "N"
	switch mR.HeadingDirection {
		case  mR.West  : headingDirection = "W"
		case  mR.East  : headingDirection = "E"
		case  mR.South : headingDirection = "S"
		case  mR.North : headingDirection = "N"
		default: fmt.Println("Heading Direction - invalid")
		
	}
	//OUTPUT
	fmt.Println(mR.PosX,mR.PosY, headingDirection)

}

func main() {
	//fmt.Println("In Main.")
	
	marsRover := MarsRover{ HeadingDirection:1, North:1, East:2, South:3, West:4}

	marsRover.curPos(1, 2, marsRover.North)
	marsRover.explore("LMLMLMLMM")
	marsRover.OUTPUT_curPosAfterExplored()
}
