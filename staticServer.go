// Copyright 2015 Jeff Martinez. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file
// or at http://opensource.org/licenses/MIT

/*
A quick and easy to use web server that can be pointed to serve the contents
of any directory.

The default behavior is to serve the contents of the current directory,
accessible on port 8000 only accessible from localhost (in other words,
other machines on the network can't hit your server).

You can allow other machines on your network to access the server by using
the -a flag. This enables anyone to see the contents of the server on the
specified port.

You can change the port with -port.
Example:
	-port=54321

You can change the directory served with -dir.
Examples:
	-dir=my/relative/path       <- relative path
	-dir=../..                  <- another relative path
	-dir=/absolute/path         <- absolute path

Exit codes are as follows:

0: success ([ctrl-c] used to terminate server)
1: failure code, possibly bad directory name or permissions issue
2: Incorrect usage. Will display usage (equivalent to using -h or --help)
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
)

const EXIT_SUCCESS = 0
const EXIT_FAILURE = 1
const EXIT_USAGE_FAILURE = 2 // Same as golang's flag module uses

func main() {
	setupExitOnCtrlC()

	allowAnyHostToConnect, listenPort, directoryToServe := getCommandLineArgs()

	verifyDirectoryOrDie(directoryToServe)

	listenHost := "localhost"
	if allowAnyHostToConnect {
		listenHost = ""
	}

	listenAddress := fmt.Sprintf("%v:%v", listenHost, listenPort)

	httpDirectory := http.Dir(directoryToServe)
	fileServer := http.FileServer(httpDirectory)

	canonicalDirName := getCanonicalDirName(directoryToServe)

	visibleTo := listenHost
	if visibleTo == "" {
		visibleTo = "All ip addresses"
	}

	fmt.Printf("Server is running.\n\n")
	fmt.Printf("Directory: %v\n", canonicalDirName)
	fmt.Printf("Visible to: %v\n", visibleTo)
	fmt.Printf("Port: %v\n\n", listenPort)
	fmt.Printf("Hit [ctrl-c] to quit\n")

	log.Fatal(http.ListenAndServe(listenAddress, fileServer))
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
