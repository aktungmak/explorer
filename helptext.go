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
goto <uri>
    navigate directly to the specified URI (don't need hostname)
help
    show this help text
home
    return the the service root
jump <name>
    navigate directly to the location saved as <name>
mark <name>
    save the current location as <name>
man <method> <uri> <body>
    send a custom request to <uri> using <method> and <body>
quit
    disconnect and end session
`
