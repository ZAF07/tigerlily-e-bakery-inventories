echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc source $HOME/.bashrc

# protoc --go_out=. --go_opt=paths=source_relative ./protos/*
protoc --go_out=api/rpc --go_opt=paths=source_relative --go-grpc_out=api/rpc  --go-grpc_opt=paths=source_relative ./proto/inventory/*

echo "Copying into api/rpc"
cp -R api/rpc/proto/inventory/* api/rpc

# echo "Cleaning api/rpc directory"
# rm -rf api/rpc/proto

echo "DONE"