# staticserver

A quick and easy to use web server that can serve the contents of any directory.

The **default behavior** is to serve the contents of the current directory,
accessible on port 8000 only accessible from localhost (in other words,
**other machines on the network can't hit your server**).

You can allow other machines on your network to access the server by using
the -a flag. This enables anyone to see the contents of the server on the
specified port.

You can change the port with -port

	Example:
	-port=54321

You can change the directory served with -dir

	Examples:
	-dir=my/relative/path       <- relative path  
	-dir=../..                  <- another relative path  
	-dir=/absolute/path         <- absolute path  

Exit codes are as follows:

	0: success ([ctrl-c] used to terminate server)  
	1: failure code, possibly bad directory name or permissions issue  
	2: Incorrect usage. Will display usage (equivalent to using -h or --help)  

Complete examples:

	staticserver -port=8080 -dir=do/what/I/want
	staticserver -a -port=51515 -dir=/sharing/with/coworkers

For license info, see LICENSE.txt (Short version: MIT license).

## Endpoints

/dir/ - Shows the contents of the current directory

/echo/ - The response echos whatever comes after the `/echo/` portion of the url, for example, hitting `/echo/hello` will respond with `hello`

/increment - Responds with ever incrementing numbers, based on the number of times you've hit the endpoint. The first time responds with `1`, then `2`, and so on.

/random - Responds with a random integer.
