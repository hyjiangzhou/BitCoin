#!/bin/bash
rm *.db
rm blockchain
go build -o blockchain *.go
./blockchain createBlockChain 1GmzSeeAd91jXNuu9upyCyUVwbcM1DYhKL
./blockchain
