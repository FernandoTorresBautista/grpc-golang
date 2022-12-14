## gRPC course

Patrones de comunicación <br>
- Unary 
- Server streaming 
- Client streaming 
- Bidireccional streaming 

All that you need is install <br>
- protoc (install on windows/linux) 

For golang <br>
- go get -u google.golang.org/grpc
- go install google.golang.org/protobuf/cmd/protoc-gen-go


### generate ssl files
the first step is create the folder, in this case is "ssl/" <br>
after the creation of the folder, follow the next commands to generate all the files needed: <br>
- add env variable SERVER=localhost
    - windows $env:SERVER="localhost"
    - linux   export SERVER="localhost"
- openssl genrsa --passout pass:1111 -des3 -out ca.key 4096
- openssl req --passin pass:1111 -new -x509 -days 100 -key ca.key -out ca.crt -subj "/CN=${server}"
- openssl genrsa --passout pass:1111 -des3 -out server.key 4096
- openssl req -passin pass:1111 -new -key server.key -out server.csr -subj "/CN={server}" 
- openssl x509 -req -passin pass:1111 -days=100 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt
- openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem

### run the example server_ssl.go 
go into hello\hello_server and run: <br>
- go run .\server.go

### run the example client_ssl.go 
go into hello\hello_client and run: <br>
- go run .\client.go
