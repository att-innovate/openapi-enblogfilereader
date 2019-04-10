# Pro-Active Wireless Network Resource Management and Control

## eNB Logfile Reader 

## (Cross-) Compile and Run

    $ go build openapi-enblogfilereader

    $ env GOOS=linux GOARCH=amd64 go build openapi-enblogfilereader && scp ./openapi-enblogfilereader <USER>:<HOST>/<PATH>
    

    $ Ctrl+A then Shift+H in screen mode to enable output to logfile


    $ ./openapi-enblogfilereader --screen=<SCREEN NUMBER>