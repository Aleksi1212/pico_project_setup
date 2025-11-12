# Pico_project_setup

## Install go
```
cd ~
curl -OL https://golang.org/dl/go1.21.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.21.1.linux-amd64.tar.gz
```
#### Set go paths
```
sudo vi ~/.profile

# at the end of file
export PATH=$PATH:/usr/local/go/bin

source ~/.profile
```
#### Check installation
```
go version
```

## Build and run
```
# build
chmod +x build.sh
./build.sh

# move to /usr/bin
sudo mv pico_project_setup /usr/bin

# run
pico_project_setup <project_name> <clang>
```
