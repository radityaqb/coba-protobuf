# coba-protobuf

# Install
Git Clone Protobuf
```bash
git clone https://github.com/google/protobuf.git
```
Install protoc
```bash
sudo apt-get install protobuf-compiler
```

# Build 
Build .proto file(s) to .pb.go file(s) using 
```bash
protoc --go_out=. *.proto
```

# Test
1. Build and Run the server
2. Build and Run the client
