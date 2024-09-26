package main

import ("fmt"
	"log"
	"os/exec"
	"os"
)
//main function is basically a functioning level. I could copy and paste it and make more levels. that'd be pretty cool, and it'd make it easier for people to make their own levels. although I might be the only person who would.
func main() {

	var chr = "\033[41mX"

	var fuck [20][20]string

	fuck = newMap()
	
	printMap(fuck)
	var choice string
	var x int
	var y int
	x = 3
	y = 5

	for {
		
		fmt.Println("type command up")
		fmt.Scanln(&choice)
		
		switch choice {
			

			case "w":
				x--
				if x == 0 {
					x++
					continue
				}
			case  "s":
				x++
				if x == 19 {
                                        x--
                                        continue
                                }

			case "d":
				y++
				if y == 19 {
                                        y--
                                        continue
                                }

			case "a":
				y--
				if y == 0 {
                                        y++
                                        continue
                                }


		}


		

	clear()
		
	fuck = newMap()
	//reminder to make collision work	
	fuck[3][7] = "\033[45m="
	fuck[3][8] = "\033[45m="
        fuck[3][9] = "\033[45m="
	fuck[5][7] = "\033[44mH"
	fuck[5][8] = "\033[49mn" 
	fuck[5][9] = "\033[44mH"
	fuck[4][7] = "\033[44mH"
        fuck[4][8] = "\033[44mH"
        fuck[4][9] = "\033[44mH"


	fuck[5][4] = "\033[41mE"
	
	fuck[x][y] = chr
	
	if fuck[5][4] == chr {
                log.Fatal()
        }

	if fuck[5][8] == chr {
		levelTwo()
	}

	printMap(fuck)
	//I fucked up and var x measures the vertical position. var y measures the horizontal position. I switched it in the print statement so it's technically right.
	fmt.Println("Y: ", x)
	fmt.Println("X: ", y)
	continue
	}
	


}


func levelTwo() {
var chr = "8"
        var fuck [20][20]string

        fuck = newMap()

        printMap(fuck)

        var choice string
        var x int
        var y int

        for {
		
                fmt.Println("type command up")
                fmt.Scanln(&choice)

                switch choice {

                        case "up", "w":
                                x--
                        case "down", "s":
                                x++
                        case "left", "a":
                                y--
                        case "right", "d":
                                y++

                        }


        if fuck[5][4] == chr {
                log.Fatal()
        }

        if fuck[5][8] == chr {

        }
	
	clear()
        fuck = newMap()

        fuck[1][5] = "%"
        fuck[1][6] = "%"
        fuck[1][7] = "%"
        fuck[3][5] = "H"
        fuck[3][6] = "n"
        fuck[3][7] = "H"
        fuck[2][5] = "H"
        fuck[2][6] = "H"
        fuck[2][7] = "H"


        fuck[5][4] = "E"

        fuck[x][y] = chr


        printMap(fuck)

        }


}

func clear() {

        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}


//prints 2d array
func printMap(fuck [20][20]string) {

/*	for d := 0; d <= 9; d++ {

                fmt.Println("", fuck[d], " ", "\n")

        }
	fmt.Println()*/

	for i := 0; i < len(fuck); i++ {    // Loop through each row
    for j := 0; j < len(fuck[i]); j++ { // Loop through each element in the row
        fmt.Print(fuck[i][j], " ")  // Print element with a space
    }
    fmt.Println() // Move to the next line after printing a row
}

}
//THIS IS WHERE YOU CAN CHANGE THE BACKGROUND SHIT
//returns 2d array of all Xs
func newMap() [20][20]string  {
	
	borderColor := "\033[40mX"
	backGroundColor := "\033[42m0"

	var fuck [20][20]string

        for j := 0; j <= 19; j++ {
                for i := 0; i <= 19; i++ {

                        fuck[j][i] = borderColor

                }
        }
	//this part fills the inside with 0s
	for d := 1; d <= 18; d++ {
		for f := 1; f <= 18; f++ {
			fuck[d][f] = backGroundColor
		}
        }



return fuck

}


