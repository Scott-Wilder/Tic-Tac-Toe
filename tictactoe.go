package main

import (
	"fmt"
	"math/rand"
	"time"
)

type player struct {
	name string
	winner bool
	piece string
	npc bool
}

func main() {
	fmt.Println("Welcome to Scott's Tic-Tac-Toe Game")
	var players [2]string = collectPlayersName()
	fmt.Println("Players: ", players)
	var gameBoard = createGameBoard()
	var player1, player2 = newPlayer(&players)
	printGameBoard(gameBoard)
	takeTurn(gameBoard, player1, player2)
}

func collectPlayersName() [2]string {
	var player1, player2 string
	var players [2]string
	fmt.Println("Player 1, please enter your Name:")
	fmt.Scan(&player1)
	fmt.Println("Player 2, please enter your Name or type NPC:")
	fmt.Scan(&player2)
	players[0] = player1
	players[1] = player2
		return players
}

func newPlayer(players *[2]string) (*player, *player) {
	var player1 = player{players[0], false, "X", false}
	var player2 = player{players[1], false, "O", false}
	if(players[1] == "NPC") {
		player2 = player{players[1], false, "O", true}
	}
	return &player1, &player2
}

func createGameBoard() *[3][3]string {
	var gameBoard = [3][3]string {{"*", "*", "*"}, {"*", "*", "*"}, {"*", "*", "*"}}
	return &gameBoard
}

func printGameBoard(gameBoard *[3][3]string) {
	var i int
	for  i = 0; i < 3; i++ {
		fmt.Printf("GameBoard = %v\n", gameBoard[i])
	}
}

func takeTurn(gameBoard *[3][3]string, player1 *player, player2 *player) [3][3]string{
	var move [2]int
	out:
	for {
		fmt.Println(player1.name, "please enter your move as coordinates [x][y]")
		fmt.Scan(&move[0])
		fmt.Scan(&move[1])
		var x = move[0]
		var y = move[1]
		if (validMove(move, gameBoard, &x, &y) == true) {
			fmt.Printf("%s move [%d][%d]\n", player1.name,x,y)
			gameBoard[x][y] = player1.piece
			printGameBoard(gameBoard)
			if (gameOver(gameBoard, &move, &x, &y, player1, player2) == true) {
				fmt.Println("GAME OVER!")
				break out
			} 
			for {
				// NPC logic
				if(player2.npc == true) {
					rand.Seed(time.Now().UnixNano())
					min := 0
					max := 2
					// random number between 0-2 for both x,y.
					var x = rand.Intn(max - min + 1) + min
					var y = rand.Intn(max - min + 1) + min
					if (validMove(move, gameBoard, &x, &y) == true) {
						fmt.Printf("NPC move [%d][%d]\n", x,y)
						gameBoard[x][y] = player2.piece
						printGameBoard(gameBoard)
						if (gameOver(gameBoard, &move, &x, &y, player1, player2) == true) {
							fmt.Println("GAME OVER!")
							break out
						}
						break
					} else {
						fmt.Println(player2.name, "redo turn.")
					}
				}
				if(player2.npc == false) {
					fmt.Println(player2.name, "please enter your move as coordinates [x][y]")
					fmt.Scan(&move[0])
					fmt.Scan(&move[1])
					var x = move[0]
					var y = move[1]
					if (validMove(move, gameBoard, &x, &y) == true) {
						fmt.Printf("%s move [%d][%d]\n", player2.name,x,y)
						gameBoard[x][y] = player2.piece
						printGameBoard(gameBoard)
						if (gameOver(gameBoard, &move, &x, &y, player1, player2) == true) {
							fmt.Println("GAME OVER!")
							break out
						}
						break
					} else {
						fmt.Println(player2.name, "redo turn.")
					}
				}
			}
		} else {
				fmt.Println(player1.name, "redo turn.")
		}
	}
	return *gameBoard
}

func validMove(move [2]int, gameBoard *[3][3]string, x *int, y *int) bool{
	if (*x > 2 || *y > 2)  {
		printGameBoard(gameBoard)
		fmt.Println("Out of bounds move.")
		return false
	}
	if (*x < 0 || *y < 0)  {
		printGameBoard(gameBoard)
		fmt.Println("Out of bounds move.")
		return false
	}
	if (gameBoard[*x][*y] == "*") {
		return true
	} else {
		printGameBoard(gameBoard)
		fmt.Println("Space already taken.")
		return false
	}
} 

func gameOver(gameBoard *[3][3]string, move *[2]int, x *int, y *int, player1 *player, player2 *player) bool {
	//[0,0], [1,0], [2,0]
	if(gameBoard[0][0] == player1.piece && gameBoard[1][0] == player1.piece && gameBoard[2][0] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[0][0] == player2.piece && gameBoard[1][0] == player2.piece && gameBoard[2][0] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	//[0,0], [0,1], [0,2]
	if(gameBoard[0][0] == player1.piece && gameBoard[0][1] == player1.piece && gameBoard[0][2] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[0][0] == player2.piece && gameBoard[0][1] == player2.piece && gameBoard[0][2] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	//[0,0], [1,1], [2,2]
	if(gameBoard[0][0] == player1.piece && gameBoard[1][1] == player1.piece && gameBoard[2][2] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[0][0] == player2.piece && gameBoard[1][1] == player2.piece && gameBoard[2][2] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	//[0,1], [1,1], [2,1]
	if(gameBoard[0][1] == player1.piece && gameBoard[1][1] == player1.piece && gameBoard[2][1] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[0][1] == player2.piece && gameBoard[1][1] == player2.piece && gameBoard[2][1] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	//[1,0], [1,1], [1,2]
	if(gameBoard[1][0] == player1.piece && gameBoard[1][1] == player1.piece && gameBoard[1][2] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[1][0] == player2.piece && gameBoard[1][1] == player2.piece && gameBoard[1][2] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	//[0,2], [1,2], [2,2]
	if(gameBoard[0][2] == player1.piece && gameBoard[1][2] == player1.piece && gameBoard[2][2] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[0][2] == player2.piece && gameBoard[1][2] == player2.piece && gameBoard[2][2] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	//[0,2], [1,1], [2,0]
	if(gameBoard[0][2] == player1.piece && gameBoard[1][1] == player1.piece && gameBoard[2][0] == player1.piece) {
		fmt.Println(player1.name, "is the Winner!")
		return true
	} else if(gameBoard[0][2] == player2.piece && gameBoard[1][1] == player2.piece && gameBoard[2][0] == player2.piece) {
		fmt.Println(player2.name, "is the Winner!")
		return true
	}
	// no winner, all spaces taken.
	if(gameBoard[0][0] != "*" && gameBoard[0][1] != "*"  && gameBoard[0][2] != "*" && gameBoard[1][0] != "*" && gameBoard[1][1] != "*" && gameBoard[1][2] != "*" && gameBoard[2][0] != "*" && gameBoard[2][1] != "*" && gameBoard[2][2] != "*"){
		fmt.Println("Cats Game, No Winner!")
		return true
	}
	return false
}
