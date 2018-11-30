# Copyright 2018 Kubedge.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

DOCKER_NAMESPACE ?= kubedge
VERION           ?= 0.1.0
SIMU_NAME        = kubesim_epc
CURRENT_BRANCH   = ${shell git rev-parse --abbrev-ref HEAD}
IMAGE_NAME       = ${SIMU_NAME}-${CURRENT_BRANCH}
DHUBREPO         = ${DOCKER_NAMESPACE}/${IMAGE_NAME}

.PHONY: all

all: simulator

simulator: kubesim_epc/main.go
	cd kubesim_epc && CGO_ENABLED=0 GOOS=linux GOARM=7 GOARCH=arm go build -a -installsuffix cgo -o goserv-${CURRENT_BRANCH} .

clean: kubesim_epc/goserv-${CURRENT_BRANCH}
	cd kubesim_epc && rm -f goserv-${CURRENT_BRANCH}