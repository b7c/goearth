package main

import (
	"fmt"
	"log"
	"strings"

	g "github.com/b7c/goearth"
)

var ext = g.NewExt(g.ExtInfo{
	Title:       "Go Earth",
	Description: "example: basic",
	Version:     "1.0",
	Author:      "b7",
})

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	// Handling extension initialization
	ext.Initialized(func(e *g.InitArgs) {
		log.Printf("Extension initialized (connected=%t)", e.Connected)
	})
	// Handling extension activation
	// (when the green "play" button is clicked in G-Earth)
	ext.Activated(func() {
		log.Printf("Extension activated")
		// Launch a goroutine to retrieve the user's info
		if ext.IsConnected() {
			go getUserInfo()
		} else {
			log.Printf("Game is not connected")
		}
	})
	// Handling game connection start
	ext.Connected(func(e *g.ConnectArgs) {
		log.Printf("Game connection established (%s:%d)", e.Host, e.Port)
		log.Printf("Client %s (%s)", e.Client.Identifier, e.Client.Version)
		log.Printf("Received %d message info", len(e.Messages))
	})
	// Intercepting all packets
	ext.InterceptAll(func(e *g.InterceptArgs) {
		if e.Packet.Header.Name() == "Ping" {
			log.Printf("Received ping")
		}
	})
	// Intercepting specific packets
	ext.Intercept("Chat", "Shout", "Whisper").Out(handleChat)
	// Handling game disconnection
	ext.Disconnected(func() { log.Printf("Game connection lost") })
	ext.Run()
}

func handleChat(e *g.InterceptArgs) {
	// Reading data from packets
	msg := e.Packet.ReadString()
	action := strings.ToLower(e.Packet.Header.Name())
	if action == "chat" {
		action = "said"
	} else {
		action += "ed"
	}
	log.Printf("You %s %q", action, msg)
	if strings.Contains(msg, "block") {
		// Blocking packets
		e.Block = true
		log.Printf("Blocking message!")
	} else if strings.Contains(msg, "apple") {
		// Modifying packets
		msg = strings.ReplaceAll(msg, "apple", "orange")
		e.Packet.ReplaceStringAt(msg, 0)
		log.Printf("Replacing message: %q", msg)
	}
}

func getUserInfo() {
	log.Printf("Retrieving user info...")
	// Sending packets
	ext.Send("InfoRetrieve")
	// Receiving packets
	/*
		Block() prevents the packet from reaching the client.
		You can also chain Timeout(), TimeoutSec() or TimeoutMs()
		to adjust the timeout. The default timeout is 1 minute.
		If() can be chained to pass in a function that returns true
		if the packet should be intercepted.
	*/
	if pkt := ext.Recv("UserObject").Block().Wait(); pkt != nil {
		id, name := pkt.ReadId(), pkt.ReadString()
		msg := fmt.Sprintf("Got user info. (id:%d, name:%q)", id, name)
		log.Println(msg)
		// Sending client-side packets
		ext.SendToClient("Chat", 0, msg, 0, 34, 0, 0)
	} else {
		log.Printf("Timed out.")
	}
}
