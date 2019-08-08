package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"

	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters
var (
	Token    string
	dbAdd    *sql.Stmt
	dbCheck  *sql.Stmt
	dbTopTen *sql.Stmt
	dev      bool
)

func init() {
	flag.BoolVar(&dev, "d", false, "Dev mode")
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(puddle.CmdHelp)
	dg.AddHandler(puddle.CmdPing)
	dg.AddHandler(puddle.CmdAdd)
	dg.AddHandler(puddle.CmdSub)
	dg.AddHandler(puddle.CmdCheck)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
