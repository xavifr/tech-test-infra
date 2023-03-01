#!/bin/sh
dir=$(dirname $(realpath $0))

virsh destroy immfly-debian10 2>&-
virsh create $dir/assets/vm.xml

virsh net-dhcp-leases default

NET=$(virsh domifaddr immfly-debian10 | awk '/ipv4/ { print $4 }')
echo "[deploy_machines]\n${NET%/*}" > $dir/deploy/hosts

echo "Detected IP Address ${NET%/*}"
