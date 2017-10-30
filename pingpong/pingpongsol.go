package pingpong
import "fmt"

func player(playerName string, ch chan string, rounds int) {
	// TODO create the player behaviour, by implementing a loop where
	// the player initially waits for a message,
	// then prints it
	// and then sends another message on the channel to unlock who's waiting for a message
}

func main() {
	fmt.Println("Welcome to the Ping Pong Game!")

	rounds := 10

	// TODO create a new channel that will used to transmit strings

	// TODO create two goroutines where you create two players, pinger and ponger

	// TODO send an initial message "init" on the channel to start the players

	var input string
	fmt.Scanln(&input)
}