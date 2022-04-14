# Gotbit balance monitoring task

Simple task to monitor every 30 seconds balance of specified account

## Install and build
```bash
# Install mage
git clone https://github.com/magefile/mage
cd ./mage
go run bootstrap.go
export PATH="$PATH:$(go env GOPATH)/bin"

# Download and build program
git clone https://github.com/k0tletka/gotbit-balance-monitoring-task
cd ./gotbit-balance-monitoring-task
mage build
./build/balance-monitoring
```
