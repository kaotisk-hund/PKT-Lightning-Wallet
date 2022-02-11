#!  /usr/bin/bash

################################################################################
#   smoke test for pld / pldctl commands
################################################################################

export  PKT_HOME="$( pwd )"
export  PLD="${PKT_HOME}/bin/pld"
export  PLD_OPTIONS=""
export  PLD_OUTPUT_FILE="./pld.out"
export  PLD_PID=
export  PLDCTL="${PKT_HOME}/bin/pldctl"
export  PLDCTL_OPTIONS=""
export  PLDCTL_ERRORS_FILE="./pldctl.err"
export  JSON_OUTPUT=""

#   start pld deamon in background
startPldDeamon() {

    ${PLD} ${PLD_OPTIONS} > ${PLD_OUTPUT_FILE} &

    PLD_PID=$!

    echo ">>> ${PLD} daemon up and running: PID: ${PLD_PID}"

    sleep 10s
}

#   stop pld deamon
stopPldDeamon() {

    executeCommand 'stop'

    kill ${PLD_PID} 2> /dev/null

    sleep 10s
}

#   send a command to create a wallet
createWallet() {

    local OUTPUT=$( perl -w ./test/createWallet.pl )

    if [ -z "$( echo ${OUTPUT} | grep 'pld successfully initialized!' )" ]
    then
        kill ${PLD_PID}
        exit "error: fail attempting to run command create wallet"
    fi

    echo ">>> create: command successfully executed"
}

#   send a command to unlock the wallet
unlockWallet() {

    local OUTPUT=$( perl -w ./test/unlockWallet.pl )

    if [ -z "$( echo ${OUTPUT} | grep 'lnd successfully unlocked!' )" ]
    then
        kill ${PLD_PID}
        exit "error: fail attempting to run command unlock wallet"
    fi

    echo ">>> unlock: command successfully executed"
}

#   use pldctl to execute a command
executeCommand() {
    local COMMAND="${1}"
    local ARGUMENTS="${2}"

    JSON_OUTPUT=$( ${PLDCTL} ${PLDCTL_OPTIONS} ${COMMAND} ${ARGUMENTS} 2>> ${PLDCTL_ERRORS_FILE} )
    if [ $? -eq 0 ]
    then
        echo ">>> ${COMMAND} ${ARGUMENTS}: command successfully executed"
    else
        echo "error: fail attempting to run command \"${COMMAND} ${ARGUMENTS}\": $?"
        return 1
    fi
}

#   splash screen
echo ">>>>> Testing pld and pldctl"
echo

#   check if jq is available

output=$( which jq 2> /dev/null )
if [ $? -ne 0 ]
then
    exit "error: 'jq' is required to run this script"
fi

#   parse CLI arguments
CREATE_WALLET="false"

while [ true ]
do
    ARG=${1}

    if [ -z "${ARG}" ]
    then
        break
    fi

    if [ "${ARG}" == "--createWallet" ]
    then
        CREATE_WALLET="true"
    fi

    shift
done

#   create wallet when requested
if [ "${CREATE_WALLET}" == "true" ]
then
    #   clean things up by removing previous wallet
    rm -rf ~/.lncli ~/.pki ~/.pktd ~/.pktwallet
    rm -rf ${PLD_OUTPUT_FILE}
    rm -rf ${PLDCTL_ERRORS_FILE}

    #   start pld deamon, create a wallet and stop the deamon, because first test is unlock wallet
    startPldDeamon
    createWallet
    stopPldDeamon
fi

#   star pld daemon and test the command to unlock the wallet
startPldDeamon
unlockWallet

#   test commands to get info about the running pld daemon
executeCommand 'getinfo'
if [ $? -eq 0 ]
then
    echo -e "\t#neutrino peers: $( echo ${JSON_OUTPUT} | jq '.neutrino.peers | length' )"
fi

executeCommand 'getrecoveryinfo'
if [ $? -eq 0 ]
then
    echo -e "\trecovery mode: $( echo ${JSON_OUTPUT} | jq '.recovery_mode' )"
fi

executeCommand 'debuglevel' '--level info --show'

executeCommand 'version'
if [ $? -eq 0 ]
then
    echo -e "\tpld version: $( echo ${JSON_OUTPUT} | jq '.pld | .version' )"
    echo -e "\tpldctl version: $( echo ${JSON_OUTPUT} | jq '.pldctl | .version' )"
fi

#   test commands to manage channels
#executeCommand 'openchannel'
#executeCommand 'closechannel'
#executeCommand 'closeallchannels'
#executeCommand 'abandonchannel'

executeCommand 'channelbalance'
if [ $? -eq 0 ]
then
    echo -e "\tchannel balance: $( echo ${JSON_OUTPUT} | jq '.balance' )"
fi

executeCommand 'pendingchannels'
if [ $? -eq 0 ]
then
    echo -e "\tlimbo balance: $( echo ${JSON_OUTPUT} | jq '.total_limbo_balance' )"
fi

executeCommand 'listchannels'
if [ $? -eq 0 ]
then
    echo -e "\t#open channels: $( echo ${JSON_OUTPUT} | jq '.channels | length' )"
fi

executeCommand 'closedchannels'
if [ $? -eq 0 ]
then
    echo -e "\t#closed channels: $( echo ${JSON_OUTPUT} | jq '.channels | length' )"
fi

executeCommand 'getnetworkinfo'
if [ $? -eq 0 ]
then
    echo -e "\t#nodes: $( echo ${JSON_OUTPUT} | jq '.num_nodes' )"
    echo -e "\t#channels: $( echo ${JSON_OUTPUT} | jq '.num_channels' )"
fi

executeCommand 'feereport'
if [ $? -eq 0 ]
then
    echo -e "\tweek fee sum: $( echo ${JSON_OUTPUT} | jq '.week_fee_sum' )"
fi

#executeCommand 'updatechanpolicy' '10 10 20'

executeCommand 'exportchanbackup' '--all'
if [ $? -eq 0 ]
then
    MULTI_BACKUP=$( echo ${JSON_OUTPUT} | jq '.multi_chan_backup.multi_chan_backup' | tr --delete '"' )
    echo -e "\tmulti backup: ${MULTI_BACKUP}"
fi

executeCommand 'verifychanbackup' "--multi_backup=${MULTI_BACKUP}"
executeCommand 'restorechanbackup'

#   test commands to get graph info
executeCommand 'describegraph'
#executeCommand 'getnodemetrics'
#executeCommand 'getchaninfo'
executeCommand 'getnodeinfo'

#   test commands to manage invoices
executeCommand 'addinvoice'
#executeCommand 'lookupinvoice'
executeCommand 'listinvoices'
#executeCommand 'decodepayreq'

#   test commands to manage on-chain transactions
#executeCommand 'estimatefee'
#executeCommand 'sendmany'
#executeCommand 'sendcoins'
executeCommand 'listunspent'

#   test commands to deal with profile
executeCommand 'profile' 'add pld_test'
executeCommand 'profile' 'list'
executeCommand 'profile' 'setdefault pld_test'
executeCommand 'profile' 'remove pld_test'

#   remove profile file created by profile commands
rm ~/.lncli/profiles.json

#   test commands to deal with the wallet
executeCommand 'newaddress' 'p2wkh'
executeCommand 'walletbalance'
executeCommand 'resync'
executeCommand 'getaddressbalances'
executeCommand 'getwalletseed'
executeCommand 'getsecret'
executeCommand 'getnewaddress'

#   show any eventual error during command execution
if [ -f "${PLDCTL_ERRORS_FILE}" -a $( stat --format='%s' "${PLDCTL_ERRORS_FILE}" ) -gt 0 ]
then
    echo ">>> errors executing some command "
    echo "+++++++++++++++"
    cat ${PLDCTL_ERRORS_FILE}
fi

rm -rf ${PLDCTL_ERRORS_FILE}

#   stop pld daemon
stopPldDeamon
