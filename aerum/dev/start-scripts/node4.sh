#!/bin/bash
$PWD/../../build/bin/aerum  --datadir $PWD/nodes/node4/data --mine --gasprice "1"  --syncmode "full" --networkid 2018 --port 4444 console --unlock  $(cat $PWD/nodes/node4/account.txt|cut -f 2 -d " "|sed 's/{//g'|sed 's/}//g') --password <(echo "node4")
