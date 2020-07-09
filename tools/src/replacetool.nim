import os, strutils, parseopt

type
  Options = object
    debianDir, package, maintainer, version, arch: string

proc getCmdOpts(params: seq[string]): Options =
  var optParser = initOptParser(params)
  for kind, key, val in optParser.getopt():
    case kind
    of cmdLongOption, cmdShortOption:
      case key
      of "debian-dir", "d":
        result.debianDir = val
      of "package", "p":
        result.package = val
      of "maintainer", "m":
        result.maintainer = val
      of "version", "v":
        result.version = val
      of "arch", "a":
        result.arch = val
    of cmdEnd:
      assert false # cannot happen
    else:
      assert false

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
  args = commandLineParams()
  params = getCmdOpts(args)

  controlFile = params.debianDir/"control"
  rulesFile = params.debianDir/"rules"
  changelog = params.debianDir/"changelog"

  package = params.package
  maintainer = params.maintainer
  version = params.version
  arch = params.arch

fixFile(controlFile, package=package, maintainer=maintainer, version=version, arch=arch)
fixFile(rulesFile, package=package, maintainer=maintainer, version=version, arch=arch)
fixFile(changelog, package=package, maintainer=maintainer, version=version, arch=arch)
