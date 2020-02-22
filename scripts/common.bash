#!/usr/bin/env bash

# Colors
reset='\e[0m'
black_txt='\e[30m'
black_bg='\e[40m'
red_txt='\e[31m'
red_bg='\e[41m'
green_txt='\e[32m'
green_bg='\e[42m'
yellow_txt='\e[33m'
yellow_bg='\e[43m'
blue_txt='\e[34m'
blue_bg='\e[44m'
grey_txt='\e[37m'
grey_bg='\e[47m'

# Text decoration
normal='\e[0m'
bold='\e[1m'
underline='\e[4m'
blinking='\e[5m'

# Store error code internally.
# On the first instance of error code other than 0
# the commit fails.
error=0
_check_error_code() {
    if [[ $1 != 0 ]]; then
        error=1
        printf "\n"
    fi
}
