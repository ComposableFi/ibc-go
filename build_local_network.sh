#!/bin/bash
set -eu

PATH=.:$PATH

MONIKER="local"

simd init $MONIKER --chain-id ibcgo-1

simd keys add k1 --keyring-backend test

VALIDATOR_ADDRESS=$(simd keys show k1 -a --keyring-backend test) # cosmos1z5a2mrllta4yhkwef7vjqkcf9pwwdg6lx9ajhz
LOCAL_ADDRESS="cosmos1nnypkcfrvu3e9dhzeggpn4kh622l4cq7wwwrn0"
# 0x636f736d6f73316e6e79706b636672767533653964687a656767706e346b683632326c34637137777777726e30

simd add-genesis-account $VALIDATOR_ADDRESS 100000000000000stake
simd add-genesis-account $LOCAL_ADDRESS 100000000000stake

simd gentx k1 100000000stake --chain-id ibcgo-1 --keyring-backend test

simd collect-gentxs

cp app.toml ~/.simapp/config/
cp config.toml ~/.simapp/config/

# simd start
# simd tx bank send k1 cosmos1nnypkcfrvu3e9dhzeggpn4kh622l4cq7wwwrn0 1000000stake --keyring-backend test
# simd query bank balances cosmos1nnypkcfrvu3e9dhzeggpn4kh622l4cq7wwwrn0