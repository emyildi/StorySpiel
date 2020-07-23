package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type choices struct {
	cmd         string
	description string
	nextNode    *storyNode
	nextChoice  *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choices{cmd, description, nextNode, nil}
	if node.choices == nil {
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextChoice != nil {
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}

func (node *storyNode) render() {

	fmt.Println(node.text)
	currentChoice := node.choices
	for currentChoice != nil {
		fmt.Println(currentChoice.cmd, ":", currentChoice.description)
		currentChoice = currentChoice.nextChoice
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	currentChoice := node.choices
	for currentChoice != nil {
		if strings.ToLower(currentChoice.cmd) == strings.ToLower(cmd) {
			return currentChoice.nextNode
		}
		currentChoice = currentChoice.nextChoice
	}
	fmt.Println("Sorry, nix verstehen")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()

	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `
	1 Räuber tut dir folgen. Du näherst dich 1 Kreuzung. Logischerweise hast du 3 Möglichkeiten, entweder 
	Richtung Norden, Süden oder Osten. Wenn du mich fragst, ich würde nach Süden gehen.
	`}

	darkRoom := storyNode{text: "Ohhh fuckkk. Du merkst, es ist hier mega Dunkel. Kannst nix sehen, scheiß Laterne"}

	darkRoomLit := storyNode{text: "Schmeckt, du kannst wieder Umgebung erblicken. Jetzt aber 1 entscheidende Frage, willst du dein Herz folgen und weiter gehen, oder gehst du zurück und kämpfst?"}

	grue := storyNode{text: "Ein sogenannter Stein war dir im Weg, du stolperst."}

	trap := storyNode{text: "Du fällst in 1 Schlangengrube."}

	treasure := storyNode{text: "1 Polizist trifft ein. Du sagst: 'Hilfe Herr Officer, 1 Räuber ist hinter mir.'"}

	kampf := storyNode{text: "Räuber gibt Rechte, du weichst aus, machst Roundhouse-Kick - erfolglos. Plötzlich macht Räuber einen Salto, um dich zu verwirren, macht Kapoera und knockt dich aus."}

	start.addChoice("N", "Richtung Norden", &darkRoom)
	start.addChoice("S", "Richtung Süden", &darkRoom)
	start.addChoice("O", "Richtung Osten", &trap)

	darkRoom.addChoice("S", "Einfach weitergehen gen Süden", &grue)
	darkRoom.addChoice("T", "Die Taschenlampe die du immer bei dir hast einschalten", &darkRoomLit)

	darkRoomLit.addChoice("W", "Weiter gehen", &treasure)
	darkRoomLit.addChoice("K", "Kämpfen", &kampf)

	start.play()

	fmt.Println()
	fmt.Println("Ende.")

}
