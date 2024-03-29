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

# build go app
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o ./bin/app .


# install argocd cli
FROM alpine:3.19 As argocdInstaller
WORKDIR /app
RUN apk update && apk add --no-cache \
    curl \
    && rm -rf /var/cache/apk/* \
    && curl -sSL -o /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64 \
    && chmod +x /usr/local/bin/argocd


# final image
FROM alpine:3.19

# copy argocd cli
COPY --from=argocdInstaller /usr/local/bin/argocd /usr/local/bin/

# create a non-root user with id 1000
RUN adduser -D -u 1000 user1

# set work dir
WORKDIR /home/user1/app

# copy go app
COPY --from=builder /app/bin/app /home/user1/app/run

USER user1

# run
ENTRYPOINT /home/user1/app/run
