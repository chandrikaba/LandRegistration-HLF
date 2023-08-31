#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Deploying Chaincode REGNET On Property Registration Network"
echo
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
VERSION="$4"
TYPE="$5"
: ${CHANNEL_NAME:="registrationchannel"}
: ${DELAY:="5"}
: ${LANGUAGE:="node"}
: ${VERSION:=1.1}
: ${TYPE="basic"}

LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
ORGS="beneficiary agent government buyer"
TIMEOUT=15

if [ "$TYPE" = "basic" ]; then
  CC_SRC_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/chaincode/"
else
  CC_SRC_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/chaincode-advanced/"
fi

echo "Channel name : "$CHANNEL_NAME

# import utils
. scripts/utils.sh

## Install new version of chaincode on peer0 of all 4 orgs making them endorsers
echo "Installing chaincode on peer0.beneficiary.property-registration-network.com.com ..."
installChaincode 0 'beneficiary' $VERSION
echo "Installing chaincode on peer0.agent.property-registration-network.com.com ..."
installChaincode 0 'agent' $VERSION
echo "Installing chaincode on peer0.buyer.property-registration-network.com.com ..."
installChaincode 0 'buyer' $VERSION
echo "Installing chaincode on peer0.government.property-registration-network.com.com ..."
installChaincode 0 'government' $VERSION
# echo "Installing chaincode on peer0.upgrad.property-registration-network.com.com ..."
# installChaincode 0 'upgrad' $VERSION

# Instantiate chaincode on the channel using peer0.beneficiary
echo "Instantiating chaincode on channel using peer0.beneficiary.property-registration-network.com.com ..."
instantiateChaincode 0 'beneficiary' $VERSION

echo
echo "========= All GOOD, Chaincode REGNET Is Now Installed & Instantiated On Registration Network =========== "
echo

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0
