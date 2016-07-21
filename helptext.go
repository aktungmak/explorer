package explorer

const __HELP_TEXT__ = `
select one of the available links using its number
or use one of the below commands:

back
    return to the previous location
body
    print the response body for the current request
dump <filename>
    write the last response body to <filename>
freq <method> <uri> <filename>
    make request to <uri> using <method>, body is contents of <filename> 
goto <uri>
    navigate directly to the specified URI (don't need hostname)
help
    show this help text
home
    return the the service root
jump <name>
    navigate directly to the location saved as <name>
last
    print the status line of the last request
mark <name>
    save the current location as <name>
quit
    disconnect and end session
req <method> <uri> <body>
    send a custom request to <uri> using <method> and optional <body>
save <filename>
    save the current session setup to a file, so it can be resumed later 
set <body|opts> <on|off>
    enable or disable automatic printing of the options or response body
where
    show the current location
`
