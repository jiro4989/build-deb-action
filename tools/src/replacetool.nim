import os, strutils

proc fix(body, package, maintainer, version, arch: string): string =
  result =
    body
      .replace("PACKAGE", package)
      .replace("MAINTAINER", maintainer)
      .replace("VERSION", version)
      .replace("ARCH", arch)

proc fixFile(file, package, maintainer, version, arch: string) =
  let
    body = readFile(file)
    fixedBody = fix(body, package=package, maintainer=maintainer, version=version, arch=arch)
  writeFile(file, fixedBody)

let
  debDir = getEnv("DEBIAN_DIR")
  controlFile = debDir/"control"
  rulesFile = debDir/"rules"
  changelog = debDir/"changelog"

  package = getEnv("PACKAGE")
  maintainer = getEnv("MAINTAINER")
  version = getEnv("VERSION")
  arch = getEnv("ARCH")

fixFile(controlFile, package=package, maintainer=maintainer, version=version, arch=arch)
fixFile(rulesFile, package=package, maintainer=maintainer, version=version, arch=arch)
fixFile(changelog, package=package, maintainer=maintainer, version=version, arch=arch)
