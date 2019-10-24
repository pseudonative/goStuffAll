package main

import "fmt"

type player struct {
	name, position, sport, team string
	age                         int
}

func main() {
	player1 := player{"KobeBryant", "ShootingGuard", "Basketball", "LALakers", 41}
	fmt.Println("Player1=", player1)
	fmt.Printf("(1) name=%s position=%s sport=%s team=%s age=%d\n\n",
		player1.name, player1.position, player1.sport, player1.team, player1.age)

	player2 := player{
		name:     "DezBryant",
		position: "WideReceiver",
		sport:    "Football",
		team:     "FreeAgent(DallasCowboys)",
		age:      30,
	}
	fmt.Println("Player2=", player2)
	fmt.Printf("(2) name=%s position=%s sport=%s team=%s age=%d\n\n",
		player2.name, player2.position, player2.sport, player2.team, player2.age)

	var player3 player
	player3.name = "MichaelJordan"
	player3.position = "ShootingGuard"
	player3.sport = "Basketball"
	player3.team = "Retired(ChicagoBulls)"
	player3.age = 56
	fmt.Println("Player3", player3)
	fmt.Printf("Player3=%+v\n", player3)

	player4 := new(player)
	player4.name = "ScottyPippen"
	player4.position = "Foreward"
	player4.sport = "Basketball"
	player4.team = "Retired(ChicagoBulls)"
	player4.age = 53
	fmt.Printf("&Player4=%v\n", *player4)
	fmt.Printf("&Player4=%+v\n", *player4)
}
