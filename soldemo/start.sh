#!/usr/bin/env bash
geth --datadir ./data init genesis.json
geth --datadir ./data --networkid 12345 --http --http.addr 127.0.0.1 --http.port 8545 --http.api eth,web3,net,personal,admin,txpool --nodiscover --dev --dev.period 1 --password ./password.txt console 2>>geth.log