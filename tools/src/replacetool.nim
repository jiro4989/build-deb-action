import os, strutils, parseopt

type
  Options = object
    debianDir, package, maintainer, version, depends, arch, desc: string

proc getCmdOpts(params: seq[string]): Options =
  var optParser = initOptParser(params)
  for kind, key, val in optParser.getopt():
    case kind
    of cmdLongOption, cmdShortOption:
      case key
      of "debian-dir":
        result.debianDir = val
      of "package":
        result.package = val
      of "maintainer":
        result.maintainer = val
      of "version":
        result.version = val
      of "depends":
        result.depends = val
      of "arch":
        result.arch = val
      of "description":
        result.desc = val
    of cmdEnd:
      assert false # cannot happen
    else:
      assert false

proc replaceTemplate(body, package, maintainer, version, depends, arch, desc: string): string =
  result =
    body
      .replace("PACKAGE", package)
      .replace("MAINTAINER", maintainer)
      .replace("VERSION", version)
      .replace("DEPENDS", depends)
      .replace("ARCH", arch)
      .replace("DESC", desc)

proc formatDescription(desc: string): string =
  "Description: " & desc

proc fixFile(file, package, maintainer, version, depends, arch, desc: string) =
  let
    body = readFile(file)
    fixedBody = replaceTemplate(body, package=package, maintainer=maintainer,
                                version=version, depends=depends, arch=arch, desc=desc)
  writeFile(file, fixedBody)

let
  args = commandLineParams()
  params = getCmdOpts(args)

  controlFile = params.debianDir/"control"

  package = params.package
  maintainer = params.maintainer
  version = params.version.strip(trailing = false, chars = {'v'})
  depends = params.depends
  arch = params.arch
  desc = params.desc.formatDescription

fixFile(controlFile, package=package, maintainer=maintainer, version=version,
        depends=depends, arch=arch, desc=desc)
