#!/bin/sh
dir=$(dirname $(realpath $0))

[ -f "${dir}/config.env" ] && . ${dir}/config.env

cd $dir/deploy

ansible-playbook -i hosts playbook.yml
