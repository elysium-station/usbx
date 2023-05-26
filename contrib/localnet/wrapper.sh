#!/usr/bin/env sh

##
## Input parameters
##
BINARY=/blackfury/${BINARY:-blackfuryd}
ID=${ID:-0}
LOG=${LOG:-blackfuryd.log}

##
## Assert linux binary
##
if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found. Please add the binary to the shared folder. Please use the BINARY environment variable if the name of the binary is not 'blackfuryd' E.g.: -e BINARY=blackfuryd_my_test_version"
	exit 1
fi
BINARY_CHECK="$(file "$BINARY" | grep 'ELF 64-bit LSB executable, x86-64')"
if [ -z "${BINARY_CHECK}" ]; then
	echo "Binary needs to be OS linux, ARCH amd64"
	exit 1
fi

##
## Run binary with all parameters
##
export BLACKFURY_HOME="/blackfury/node${ID}/blackfuryd"

if [ -d "$(dirname "${BLACKFURY_HOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${BLACKFURY_HOME}" --trace "$@" | tee "${BLACKFURY_HOME}/${LOG}"
else
  "${BINARY}" --home "${BLACKFURY_HOME}" --trace "$@"
fi
