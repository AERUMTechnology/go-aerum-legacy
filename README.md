## Go AERUMTechnology

Official golang implementation of the AERUMTechnology protocol featuring our custom DxPoS Consensus Algorithm (Delegated Cross-chain Proof-of-stake).

If you need any assistance please do not hesitate to contact one of the AERUM team via our twitter @aerum_official or if you would like more information 
in person feel free to buy us a beer (or coffee).


## Installation

1. clone repo, ```cd``` into it and run ```make all``` this will create all the aerum binaries on your os (it may also be necessary to add them to your path vars).

2. Generate a genesis file: the easiest way to do this is to run the ```puppeth``` binary (located here ./go-aerum/build/bin/aerum). It will populate the genesis with our custom ATMOS configuration parameters as well as call our governance contract to get the current list of our bootstrap delegates and populate it with them.

3. create a data dir on your OS and generally you can run something like this for a normal standard node (replace ./go-aerum/build/bin/aerum if you update your $PATH variables and your data dir):

``` ./go-aerum/build/bin/aerum --datadir=./datadir/ --atmos.ethereum.endpoint="https://rinkeby.infura.io" --syncmode 'full' --networkid 418313827693 --ws --wsaddr="0.0.0.0" --wsorigins "*" --wsapi admin,db,eth,debug,miner,net,txpool,clique,web3 --rpc --rpcapi admin,db,eth,debug,miner,net,txpool,clique,web3 --rpcaddr "0.0.0.0" --rpccorsdomain "*" --verbosity="3" --fast --targetgaslimit="25000000" --gasprice="1000000000" --txpool.globalslots="65536" --txpool.accountqueue="1024" --txpool.globalqueue="16384" --cache=1024 console```

or if you are a delegate you would run something similar to this:

```./go-aerum/build/bin/aerum --datadir=./datadir/ --atmos.ethereum.endpoint="https://rinkeby.infura.io" --mine --unlock b7668f414ccb7ba45b5ac1a61ce5201e3910ba17 --password <(echo "your password goes here") --rpc --rpcapi admin,db,eth,debug,miner,net,txpool,clique,web3 --rpcport "8547" --rpcaddr "0.0.0.0" -rpcvhosts="*" --rpccorsdomain "*" --networkid="418313827693" --verbosity="3" --fast --targetgaslimit="25000000" --gasprice="1000000000" --txpool.globalslots="65536" --txpool.accountqueue="1024" --txpool.globalqueue="16384" --cache=1024 console```

4. If you need to tweak anything to do with the ATMOS consensus check out the file ```./params/atmos_params.go```