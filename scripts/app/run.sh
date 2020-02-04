#!/usr/bin/env bash

root=$(git rev-parse --show-toplevel)
source $root/scripts/common.bash
args=$@

host_user_id=$(id -u)
host_group_id=$(id -g)

docker run -u ${host_user_id}:${host_group_id} -ti -v mcs2lm:/app-bin -v ${PWD}:/app --rm mcs2lm /bin/bash -c "/app-bin/mcs2lm ${args}"
