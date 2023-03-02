#!/bin/bash
dir=$(dirname $(realpath $0))

[ -f "${dir}/config.env" ] &&  {
source "${dir}/config.env"
}

cd $dir/deploy


ansible-playbook -i hosts playbook.yml
