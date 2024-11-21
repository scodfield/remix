#!/usr/bin/env bash
geth --datadir ./data --state.scheme=hash init genesis.json
geth --datadir ./data --networkid 1234 --port 30303 --http --http.addr 127.0.0.1 --http.port 8545 --http.api eth,web3,personal,admin,txpool --nodiscover --dev --dev.period 1 --gcmode archive console 2>>geth.log