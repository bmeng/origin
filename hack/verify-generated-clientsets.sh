#!/bin/bash

set -o nounset
set -o pipefail

OS_ROOT=$(dirname "${BASH_SOURCE}")/..
source "${OS_ROOT}/hack/lib/init.sh"

cd "${OS_ROOT}"
echo "===== Verifying Generated Client sets ====="
output=$(VERIFY=true ${OS_ROOT}/hack/update-generated-clientsets.sh 2>&1)

if [[ "$?" == "0" ]]; then
  echo "SUCCESS: Generated client sets up to date."
else
  echo  $output
  echo "FAILURE: Generated client sets out of date. Please run hack/update-generated-clientsets.sh"
  exit 1
fi
