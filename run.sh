#!/bin/sh
dir=$(dirname $(realpath $0))

$dir/create-vm.sh || exit 1

$dir/deploy-app.sh || exit 2
