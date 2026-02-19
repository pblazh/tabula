GO_CMD := "go"

# --- Build targets ---
setup: setup-go setup-vscode setup-obsidian
build-go: build-darwin-arm64 build-darwin-amd64 build-linux-arm64 build-linux-amd64 build-linux-386 build-windows-arm64 build-windows-amd64 build-windows-386
build: build-go build-vscode build-obsidian
test: test-go test-vscode test-obsidian
lint: lint-go lint-vscode lint-obsidian
clean:
	rm -rf bin
	rm -f coverage.out coverage.html
	rm -f plugins/tabula.vscode/node_modules
	rm -f plugins/tabula.obsidian/node_modules

build-darwin-arm64:
	env GOOS=darwin GOARCH=arm64 {{GO_CMD}} build -o bin/darwin/arm64/tabula ./cmd/cli

build-darwin-amd64:
	env GOOS=darwin GOARCH=amd64 {{GO_CMD}} build -o bin/darwin/amd64/tabula ./cmd/cli

build-linux-arm64:
	env GOOS=linux GOARCH=arm64 {{GO_CMD}} build -o bin/linux/arm64/tabula ./cmd/cli

build-linux-amd64:
	env GOOS=linux GOARCH=amd64 {{GO_CMD}} build -o bin/linux/amd64/tabula ./cmd/cli

build-linux-386:
	env GOOS=linux GOARCH=386 {{GO_CMD}} build -o bin/linux/386/tabula ./cmd/cli

build-windows-arm64:
	env GOOS=windows GOARCH=arm64 {{GO_CMD}} build -o bin/windows/arm64/tabula.exe ./cmd/cli

build-windows-amd64:
	env GOOS=windows GOARCH=amd64 {{GO_CMD}} build -o bin/windows/amd64/tabula.exe ./cmd/cli

build-windows-386:
	env GOOS=windows GOARCH=386 {{GO_CMD}} build -o bin/windows/386/tabula.exe ./cmd/cli

test-go:
	{{GO_CMD}} test ./...

coverage:
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo "Opening coverage report in browser..."
	@open coverage.html || xdg-open coverage.html || start coverage.html


setup-go:
  {{GO_CMD}} mod download

# --- Install ---
install:
	{{GO_CMD}} install ./cmd/cli

# --- Vet & Lint ---
lint-go:
	{{GO_CMD}} fmt ./...
	{{GO_CMD}} vet ./...
	golangci-lint run

# --------------------------------------------------
# plugins
# --------------------------------------------------

# VSCode
setup-vscode:
  cd plugins/tabula.vscode && npm ci && npm audit --omit=dev

build-vscode:
  cd plugins/tabula.vscode && npm run build

test-vscode:
  cd plugins/tabula.vscode && xvfb-run  npm run test

lint-vscode:
  cd plugins/tabula.vscode && npm run lint:fix

pack-vscode:
  #!/bin/sh
  set -euo pipefail

  VERSION=$(cat VERSION.txt)
  echo pack tabula.vscode.${VERSION}.tar.gz

  cd plugins/tabula.vscode
  mkdir -p dist
  tar -czf dist/tabula.vscode.${VERSION}.tar.gz -C out .

# Obsidian
setup-obsidian:
  cd plugins/tabula.obsidian && npm ci && npm audit --omit=dev

build-obsidian:
  cd plugins/tabula.obsidian && npm run build

test-obsidian:
  cd plugins/tabula.obsidian && npm run test

lint-obsidian:
  cd plugins/tabula.obsidian && npm run lint:fix

pack-obsidian:
  #!/bin/sh
  set -euo pipefail

  VERSION=$(cat VERSION.txt)
  echo pack tabula.obsidian.${VERSION}.tar.gz

  cd plugins/tabula.obsidian
  mkdir -p dist
  tar -czf dist/tabula.obsidian.${VERSION}.tar.gz -C out .

# Vim
pack-vim:
  #!/bin/sh
  set -euo pipefail

  VERSION=$(cat VERSION.txt)
  echo pack tabula.vim.${VERSION}.tar.gz

  cd plugins/tabula.vim
  mkdir -p dist
  tar -czf dist/tabula.vim.${VERSION}.tar.gz doc ftdetect plugin syntax README.md

github-validate:
  wrkflw validate
# --------------------------------------------------
# Version bump targets (no git here)
# --------------------------------------------------
_update_version version:
  #!/usr/bin/env bash
  echo VERSION.txt
  echo -ne "{{version}}" > ./VERSION.txt

_update_json_version version json:
  #!/usr/bin/env bash
  echo {{json}}
  sed -e "s/^\([ \t]*\)\"version\":.*/\1\"version\": \"{{version}}\",/" {{json}} > tmp && mv tmp {{json}}

_update_files_version version:
  #!/usr/bin/env bash

  sed -e "s/^const VERSION = .*/const VERSION = \"{{version}}\"/" "cmd/cli/version.go" > tmp && mv tmp "cmd/cli/version.go"

_commit_version version:
	git checkout -b release/v{{version}}
	git add "VERSION.txt" "cmd/cli/version.go" "plugins/tabula.obsidian/package.json" "plugins/tabula.vscode/package.json"

	git commit -m "chore(release): bump version to {{version}} [skip ci]"
	echo "Committed version {{version}} on branch release/v{{version}}"

major:
  #!/usr/bin/env bash
  set -euo pipefail

  CUR_VERSION=`cat ./VERSION.txt`

  MAJOR=`echo $CUR_VERSION | cut -d. -f1`

  NEW_VERSION="$(($MAJOR + 1)).0.0"
  echo $CUR_VERSION "->" $NEW_VERSION

  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/package.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.vscode/package.json"
  just _update_files_version ${NEW_VERSION}
  just _update_version ${NEW_VERSION}
  just _commit_version ${NEW_VERSION}

minor:
  #!/usr/bin/env bash
  set -euo pipefail

  CUR_VERSION=`cat ./VERSION.txt`
  MAJOR=`echo $CUR_VERSION | cut -d. -f1`
  MINOR=`echo $CUR_VERSION | cut -d. -f2`

  NEW_VERSION="${MAJOR}.$(($MINOR + 1)).0"
  echo $CUR_VERSION "->" $NEW_VERSION

  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/package.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.vscode/package.json"
  just _update_files_version ${NEW_VERSION}
  just _update_version ${NEW_VERSION}
  just _commit_version ${NEW_VERSION}

patch:
  #!/usr/bin/env bash
  set -euo pipefail

  CUR_VERSION=`cat ./VERSION.txt`

  MAJOR=`echo $CUR_VERSION | cut -d. -f1`
  MINOR=`echo $CUR_VERSION | cut -d. -f2`
  PATCH=`echo $CUR_VERSION | cut -d. -f3`

  NEW_VERSION="${MAJOR}.${MINOR}.$(($PATCH + 1))"
  echo $CUR_VERSION "->" $NEW_VERSION

  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/package.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.vscode/package.json"
  just _update_files_version ${NEW_VERSION}
  just _update_version ${NEW_VERSION}
  just _commit_version ${NEW_VERSION}


