package main
import "fmt"

func player(playerName string, ch chan string, rounds int) {
	// TODO create the player behaviour, by implementing a loop where
	// the player initially waits for a message,
	// then prints it
	// and then sends another message on the channel to unlock who's waiting for a message
	for i := 0; i < rounds; i++ {
		<- ch
		fmt.Println(playerName)
		ch <- "unlock"
	}
}

func main() {
	fmt.Println("Welcome to the Ping Pong Game!")

	rounds := 10

	// TODO create a new channel that will used to transmit strings
	var sharedChannel chan string = make(chan string)

	// TODO create two goroutines where you create two players, pinger and ponger
	go player("		ping", sharedChannel, rounds)
	go player("pong", sharedChannel, rounds)

	// TODO send an initial message "init" on the channel to start the players
	sharedChannel <- "init"

	var input string
	fmt.Scanln(&input)
}
