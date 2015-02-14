// Copyright 2015 Jeff Martinez. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file
// or at http://opensource.org/licenses/MIT

/*
See README.md for full description and usage info.
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"github.com/jeffbmartinez/staticserver/handler"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1
const EXIT_USAGE_FAILURE = 2 // Same as golang's flag module uses, hardcoded at https://github.com/golang/go/blob/release-branch.go1.4/src/flag/flag.go#L812

func main() {
	randomSeed := time.Now().UnixNano()
	rand.Seed(randomSeed)

	setupExitOnCtrlC()

	allowAnyHostToConnect, listenPort, directoryToServe := getCommandLineArgs()

	verifyDirectoryOrDie(directoryToServe)

	const DIR_PREFIX = "/dir/"
	const ECHO_PREFIX = "/echo/"

	http.Handle(DIR_PREFIX, handler.NewFileServer(DIR_PREFIX, directoryToServe))
	http.Handle(ECHO_PREFIX, handler.NewEcho(ECHO_PREFIX))
	http.Handle("/random", handler.NewRandom())
	http.Handle("/increment", handler.NewIncrement())

	listenHost := "localhost"
	if allowAnyHostToConnect {
		listenHost = ""
	}

	displayServerInfo(directoryToServe, listenHost, listenPort)

	listenAddress := fmt.Sprintf("%v:%v", listenHost, listenPort)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}

func verifyDirectoryOrDie(dir string) {
	fileInfo, err := os.Stat(dir)
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot read directory '%v'", dir))
		os.Exit(EXIT_FAILURE)
	}

	if !fileInfo.IsDir() {
		log.Fatal(fmt.Sprintf("This is not a directory: '%v'", dir))
		os.Exit(EXIT_FAILURE)
	}
}

func getCanonicalDirName(dir string) string {
	canonicalDirName, err := filepath.Abs(dir)

	if err != nil {
		/* After following golang's source code, this should only happen
		in fairly odd conditions such as being unable to resolve the working
		directory. */
		log.Fatal(fmt.Sprintf("Cannot serve from directory '%v'", dir))
		os.Exit(EXIT_FAILURE)
	}

	return canonicalDirName
}

func setupExitOnCtrlC() {
	const NUM_PARALLEL_SIGNALS_TO_PROCESS = 1

	killChannel := make(chan os.Signal, NUM_PARALLEL_SIGNALS_TO_PROCESS)
	signal.Notify(killChannel, os.Interrupt, os.Kill)

	go func() {
		<-killChannel
		cleanExit()
	}()
}

func cleanExit() {
	/* \b is the equivalent of hitting the back arrow. With the two following
	   space characters they serve to hide the "^C" that is printed when
	   ctrl-c is typed.
	*/
	fmt.Println("\b\b  \n[ctrl-c] Server is shutting down")
	os.Exit(EXIT_SUCCESS)
}

func getCommandLineArgs() (allowAnyHostToConnect bool, port int, directoryToServe string) {
	const DEFAULT_PORT = 8000
	const DEFAULT_DIR = "."

	flag.BoolVar(&allowAnyHostToConnect, "a", false, "Use allow any ip address (any host) to connect. Default allows ony localhost.")
	flag.IntVar(&port, "port", DEFAULT_PORT, "Port on which to listen for connections.")
	flag.StringVar(&directoryToServe, "dir", DEFAULT_DIR, "Directory to serve. Default is current directory.")

	flag.Parse()

	/* Don't accept any positional command line arguments. flag.NArgs()
	counts only non-flag arguments. */
	if flag.NArg() != 0 {
		/* flag.Usage() isn't in the golang.org documentation,
		but it's right there in the code. It's the same one used when an
		error occurs parsing the flags so it makes sense to use it here as
		well. Hopefully the lack of documentation doesn't mean it's gonna be
		changed it soon. Worst case can always copy that code into a local
		function if it goes away :p
		Currently using go 1.4.1
		https://github.com/golang/go/blob/release-branch.go1.4/src/flag/flag.go#L411
		*/
		flag.Usage()
		os.Exit(EXIT_USAGE_FAILURE)
	}

	return
}

func displayServerInfo(directoryToServe string, listenHost string, listenPort int) {
	visibleTo := listenHost
	if visibleTo == "" {
		visibleTo = "All ip addresses"
	}

	fmt.Printf("Server is running.\n\n")
	fmt.Printf("Directory: %v\n", getCanonicalDirName(directoryToServe))
	fmt.Printf("Visible to: %v\n", visibleTo)
	fmt.Printf("Port: %v\n\n", listenPort)
	fmt.Printf("Hit [ctrl-c] to quit\n")
}
