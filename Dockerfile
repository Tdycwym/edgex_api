#
# Copyright (c) 2018-2020 VMWare, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
FROM golang:1.15


RUN mkdir -p $GOPATH/src/github.com/nju-iot/edgex_api

COPY . $GOPATH/src/github.com/nju-iot/edgex_api/

WORKDIR $GOPATH/src/github.com/nju-iot/edgex_api

RUN sh build.sh

EXPOSE 6789

ENTRYPOINT ["./output/bootstrap.sh"]
