#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

source /usr/local/opt/chruby/share/chruby/chruby.sh

chruby $(< ~/.ruby-version)

"${DIR}/wrapper.rb" "${@}"
