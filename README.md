# devserver

Updated February 14, 2015

devserver is a server designed for use by developers. It contains a number of resources useful during the testing of web related projects, such as a file server and an echo endpoint.

## Usage

Run devserver via `devserver` on the command line. It will listen for requests to the available endpoints.

One of the most useful features of devserver is that it will automatically make available the contents of the directory from which you ran `devserver`. For example, if `devserver` is executed while in */tmp/myfiles/*, the contents of */tmp/myfiles/* will be accessible through devserver at *localhost:8000/dir/*. The file server can be turned off with the `-nodir` flag.

The default behavior of devserver is to allow connections only from localhost. Since devserver's default behavior includes a fileserver which can serve the contents of a local directory, access from other machines must be explicitly enabled with the `-a` ("allow" or "all") flag. To reiterate, the **default behavior prevents other machines on the network from reaching your server**.

### Usage Flags

#### Port Number

Change the **default port** of 8000 with `-port`.

	-port=12345

#### File Server Directory

Change the directory served with -dir. Default is the directory where `devserver` is run.

	-dir=my/relative/path       <- relative path  
	-dir=../..                  <- another relative path  
	-dir=/absolute/path         <- absolute path

#### Disable File Server

Disable the file server with `-nodir`. If both `-nodir` and `-dir` are used, `-nodir` takes precedence and the file server is still disabled.

	-nodir

#### Allow All IP Address Connections

By default, only access via localhost is allowed. To allow other machines to connect to devserver, use `-a`.

	-a

#### Complete Examples of Flag Usage:

Allow access to my/private/test/files on port 8080, only via localhost (other machines can't access):

	devserver -port=8080 -dir=my/private/test/files
	
Allow access to anyone on the network, but disable the file server:

	devserver -a -nodir

Allow access to anyone on the network over port 51515. Expose */sharing/with/coworkers* directory:
	
	devserver -a -port=51515 -dir=/sharing/with/coworkers

## Exit Codes

Exit codes for devserver as as follows:

	0: success ([ctrl-c] used to terminate server)  
	1: failure code, possibly bad directory name or permissions issue  
	2: Incorrect usage. Will display usage (equivalent to using -h or --help)  

## Licence

For license info, see LICENSE.txt (Short version: MIT license).

## Available Resources

All the examples assume devserver is running locally on port 8000, as is the default behavior.

### File Server

#### Directory Contents

`localhost:8000/dir/[path to directory]`

Will return a simple html page with links to the underlying directories and files of the directory being served. The root directory as far as the file server is concerned is the one specified either by the `-dir` flag or based on the working directory at the time `devserver` was executed.

#### File Contents

`localhost:8000/dir/[path to file]`

Respond with the contents of the file.

### Echo

`localhost:8000/echo/[message]`

Responds with whatever [message] you send.

### Random

`localhost:8000/random`

Responds with a random integer

### Counter

`localhost:8000/counter`

Responds with a number, based on the number of times you've called the endpoint. The first time is 1, the second time is 2, etc.

## Available Header Options

Unless otherwise specified, any of the resources in devserver will accept http request headers to modify the behavior of the call. The following are accepted headers:

### delay-milliseconds

Adds a minimum delay to sending the response after receiving the request. This is useful for simulating higher latencies across real networks.

Example to wait one second before sending the response:

    $ curl -H 'delay-milliseconds: 1000' localhost:8000/echo/hello
