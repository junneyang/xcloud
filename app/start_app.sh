#!/bin/bash
#set -eux

export PARAMA=123456
export PARAMA=654321
./app
read -rsp $'Press enter to continue...\n'

