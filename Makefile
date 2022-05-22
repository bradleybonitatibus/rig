# Copyright 2022 Bradley Bonitatibus

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#    http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

local-pg:
	docker run \
		--env POSTGRES_USER=postgres \
		--env POSTGRES_PASSWORD=postgres \
		--env POSTGRES_DB=postgres \
		--env POSTGRES_PORT=5432 \
		-p "5432:5432" \
		-d \
		postgres:13 || exit 0;
	sleep 1

sast-scanning:
	semgrep --error --metrics=on --strict --config=p/golang -o sast.json --json

static-analysis:
	go fmt ./...
	go vet ./...
	staticcheck ./...

test:
	POSTGRES_HOST=localhost \
	POSTGRES_USER=postgres \
	POSTGRES_PASSWORD=postgres \
	POSTGRES_DB=postgres \
	POSTGRES_PORT=5432 \
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

clean:
	rm -f coverage.out

ci: clean local-pg static-analysis sast-scanning test

check:
	pre-commit run
