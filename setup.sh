sudo apt update
sudo apt install -y protobuf-compiler
sudo apt install  -y golang-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
sudo apt  -y install nodejs
sudo apt  -y install npm
sudo npm  -y install --global yarn
