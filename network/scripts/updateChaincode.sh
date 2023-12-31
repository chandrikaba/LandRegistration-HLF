#!/bin/bash

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo
echo "Updating Chaincode REGNET On Property Registration Network"
echo
CHANNEL_NAME="$1"
DELAY="$2"
LANGUAGE="$3"
VERSION="$4"
TYPE="$5"
: ${CHANNEL_NAME:="registrationchannel"}
: ${DELAY:="5"}
: ${LANGUAGE:="node"}
: ${VERSION:=1.2}
: ${TYPE="basic"}

LANGUAGE=`echo "$LANGUAGE" | tr [:upper:] [:lower:]`
ORGS="beneficiary agent government buyer"
TIMEOUT=15

if [ "$TYPE" = "basic" ]; then
  CC_SRC_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/chaincode/"
else
  CC_SRC_PATH="/opt/gopath/src/github.com/hyperledger/fabric/peer/chaincode-advanced/"
fi

echo "New Version : "$VERSION

# import utils
. scripts/utils.sh

## Install new version of chaincode on peer0 of all 3 orgs making them endorsers
echo "Updating chaincode on peer0.beneficiary.property-registration-network.com.com ..."
installChaincode 0 'beneficiary' $VERSION
echo "Updating chaincode on peer0.agent.property-registration-network.com.com ..."
installChaincode 0 'agent' $VERSION
echo "Updating chaincode on peer0.buyer.property-registration-network.com.com ..."
installChaincode 0 'buyer' $VERSION
echo "Updating chaincode on peer0.government.property-registration-network.com.com ..."
installChaincode 0 'government' $VERSION

# Upgrade chaincode on the channel using peer0.iit
echo "Upgrading chaincode on channel using peer0.beneficiary.property-registration-network.com ..."
upgradeChaincode 0 'beneficiary' $VERSION

echo
echo "========= All GOOD, Chaincode REGNET Is Now Updated To Version '$VERSION' =========== "
echo

echo
echo " _____   _   _   ____   "
echo "| ____| | \ | | |  _ \  "
echo "|  _|   |  \| | | | | | "
echo "| |___  | |\  | | |_| | "
echo "|_____| |_| \_| |____/  "
echo

exit 0
