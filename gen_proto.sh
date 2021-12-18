echo 'export GOPATH=$HOME/Go' >> $HOME/.bashrc source $HOME/.bashrc

protoc --go_out=. --go_opt=paths=source_relative ./protos/inventories.proto

echo "DONE"