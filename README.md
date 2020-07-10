# build-deb-action

<!-- TODO: replace jiro4989/build-deb-action with your repo name -->
[![Test](https://github.com/jiro4989/build-deb-action/workflows/Test/badge.svg)](https://github.com/jiro4989/build-deb-action/actions?query=workflow%3ATest)
[![reviewdog](https://github.com/jiro4989/build-deb-action/workflows/reviewdog/badge.svg)](https://github.com/jiro4989/build-deb-action/actions?query=workflow%3Areviewdog)
[![release](https://github.com/jiro4989/build-deb-action/workflows/release/badge.svg)](https://github.com/jiro4989/build-deb-action/actions?query=workflow%3Arelease)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/jiro4989/build-deb-action?logo=github&sort=semver)](https://github.com/jiro4989/build-deb-action/releases)
[![action-bumpr supported](https://img.shields.io/badge/bumpr-supported-ff69b4?logo=github&link=https://github.com/haya14busa/action-bumpr)](https://github.com/haya14busa/action-bumpr)

`build-deb-action` builds a simple debian package.

## Input

```yaml
inputs:
  package:
    description: 'Package name of debian package.'
    required: true
  package_root:
    description: 'Directory of release files.'
    required: true
  maintainer:
    description: 'Package maintainer name.'
    required: true
  version:
    description: 'Package version.'
    required: true
  arch:
    description: 'Package architecture.'
    default: 'amd64'
  desc:
    description: 'Package description.'
    default: ''
```

## Usage

```yaml
name: build

on:
  push:
    tags:
      - 'v*.*.*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Set tag
        id: vars
        run: echo ::set-output name=tag::${GITHUB_REF:10}

      - name: create sample script
        run: |
          mkdir -p .debpkg/usr/bin
          mkdir -p .debpkg/usr/lib/samplescript
          echo -e "echo sample" > .debpkg/usr/bin/samplescript
          chmod +x .debpkg/usr/bin/samplescript
          echo -e "a=1" > .debpkg/usr/lib/samplescript/samplescript.conf
      - uses: jiro4989/build-deb-action@v2
        with:
          package: samplescript
          package_root: .debpkg
          maintainer: your_name
          version: ${{ steps.vars.outputs.tag }} # vX.X.X
          arch: 'amd64'
          desc: 'this is sample package.'
```

## Development

### Release

#### [haya14busa/action-bumpr](https://github.com/haya14busa/action-bumpr)
You can bump version on merging Pull Requests with specific labels (bump:major,bump:minor,bump:patch).
Pushing tag manually by yourself also work.

#### [haya14busa/action-update-semver](https://github.com/haya14busa/action-update-semver)

This action updates major/minor release tags on a tag push. e.g. Update v1 and v1.2 tag when released v1.2.3.
ref: https://help.github.com/en/articles/about-actions#versioning-your-action

### Lint - reviewdog integration

This reviewdog action template itself is integrated with reviewdog to run lints
which is useful for Docker container based actions.

![reviewdog integration](https://user-images.githubusercontent.com/3797062/72735107-7fbb9600-3bde-11ea-8087-12af76e7ee6f.png)

Supported linters:

- [reviewdog/action-shellcheck](https://github.com/reviewdog/action-shellcheck)
- [reviewdog/action-hadolint](https://github.com/reviewdog/action-hadolint)
- [reviewdog/action-misspell](https://github.com/reviewdog/action-misspell)

