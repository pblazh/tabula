GO_CMD := "go"

# --- Build targets ---
setup: go-setup vscode-setup obsidian-setup
go-build: build-darwin-arm64 build-darwin-amd64 build-linux-arm64 build-linux-amd64 build-linux-386 build-windows-arm64 build-windows-amd64 build-windows-386
build: go-build vscode-build obsidian-build
test: go-test vscode-test obsidian-test
lint: go-lint vscode-lint obsidian-lint

go: go-setup go-lint go-test go-build
obsidian: obsidian-setup obsidian-lint obsidian-test obsidian-build obsidian-pack
vscode: vscode-setup vscode-lint vscode-test vscode-build vscode-pack
vim: vim-pack

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

go-setup:
  {{GO_CMD}} mod download

go-test:
	{{GO_CMD}} test ./...

go-coverage:
	@echo "Generating coverage report..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo "Opening coverage report in browser..."
	@open coverage.html || xdg-open coverage.html || start coverage.html


# --- Install ---
go-install:
	{{GO_CMD}} install ./cmd/cli

# --- Vet & Lint ---
go-lint:
	golangci-lint run

# --------------------------------------------------
# plugins
# --------------------------------------------------

# VSCode
vscode-setup:
  cd plugins/tabula.vscode && npm ci && npm audit --omit=dev

vscode-build:
  cd plugins/tabula.vscode && npm run build

vscode-test:
  cd plugins/tabula.vscode && npm run test

vscode-lint:
  cd plugins/tabula.vscode && npm run lint:fix

vscode-pack:
  #!/bin/sh
  set -eu

  VERSION=$(cat VERSION.txt)
  echo pack tabula.vscode.${VERSION}.tar.gz

  cd plugins/tabula.vscode
  mkdir -p dist
  tar -czf dist/tabula.vscode.${VERSION}.tar.gz -C out .

# Obsidian
obsidian-setup:
  cd plugins/tabula.obsidian && npm ci && npm audit --omit=dev

obsidian-build:
  cd plugins/tabula.obsidian && npm run build

obsidian-test:
  cd plugins/tabula.obsidian && npm run test

obsidian-lint:
  cd plugins/tabula.obsidian && npm run lint:fix

obsidian-pack:
  #!/bin/sh
  set -eu

  VERSION=$(cat VERSION.txt)
  echo pack tabula.obsidian.${VERSION}.tar.gz

  cd plugins/tabula.obsidian
  mkdir -p dist
  tar -czf dist/tabula.obsidian.${VERSION}.tar.gz -C out .

# Vim
vim-pack:
  #!/bin/sh
  set -eu

  VERSION=$(cat VERSION.txt)
  echo pack tabula.vim.${VERSION}.tar.gz

  cd plugins/tabula.vim
  mkdir -p dist
  tar -czf dist/tabula.vim.${VERSION}.tar.gz doc ftdetect plugin syntax README.md

webstorm-pack:
  #!/bin/sh
  tmp=`mktemp`
  functions=`sed -En 's/^(\t| )*"(.*)": func\(context.*/\`\2\`,/p' ./internal/core/dispatch.go | tr "\n" " "`
  echo $functions
  sed -E "s/(^.*Functions\*\*: ).*/\1${functions}/" ./plugins/tabula.webstorm/README.md > $tmp
  mv $tmp ./plugins/tabula.webstorm/README.md

github-lint:
  wrkflw validate

# --------------------------------------------------
# Version bump targets
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
  just webstorm-pack
  git checkout -b release/v{{version}}
  git add "VERSION.txt" "cmd/cli/version.go" "plugins/tabula.vscode/package.json" "plugins/tabula.obsidian/package.json" "plugins/tabula.obsidian/manifest.json"
  git add plugins/tabula.webstorm
  git commit -m "chore(release): bump version to {{version}}"
  echo "Committed version {{version}} on branch release/v{{version}}"

major:
  #!/usr/bin/env bash
  set -eu

  CUR_VERSION=`cat ./VERSION.txt`
  cp ./VERSION.txt plugins/tabula.obsidian/

  MAJOR=`echo $CUR_VERSION | cut -d. -f1`

  NEW_VERSION="$(($MAJOR + 1)).0.0"
  echo $CUR_VERSION "->" $NEW_VERSION

  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/package.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/manifest.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.vscode/package.json"
  just _update_files_version ${NEW_VERSION}
  just _update_version ${NEW_VERSION}
  just _commit_version ${NEW_VERSION}

minor:
  #!/usr/bin/env bash
  set -eu

  CUR_VERSION=`cat ./VERSION.txt`
  cp ./VERSION.txt plugins/tabula.obsidian/
  MAJOR=`echo $CUR_VERSION | cut -d. -f1`
  MINOR=`echo $CUR_VERSION | cut -d. -f2`

  NEW_VERSION="${MAJOR}.$(($MINOR + 1)).0"
  echo $CUR_VERSION "->" $NEW_VERSION

  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/package.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/manifest.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.vscode/package.json"
  just _update_files_version ${NEW_VERSION}
  just _update_version ${NEW_VERSION}
  just _commit_version ${NEW_VERSION}

patch:
  #!/usr/bin/env bash
  set -eu

  CUR_VERSION=`cat ./VERSION.txt`
  cp ./VERSION.txt plugins/tabula.obsidian/

  MAJOR=`echo $CUR_VERSION | cut -d. -f1`
  MINOR=`echo $CUR_VERSION | cut -d. -f2`
  PATCH=`echo $CUR_VERSION | cut -d. -f3`

  NEW_VERSION="${MAJOR}.${MINOR}.$(($PATCH + 1))"
  echo $CUR_VERSION "->" $NEW_VERSION

  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/package.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.obsidian/manifest.json"
  just _update_json_version ${NEW_VERSION} "plugins/tabula.vscode/package.json"
  just _update_files_version ${NEW_VERSION}
  just _update_version ${NEW_VERSION}
  just _commit_version ${NEW_VERSION}
