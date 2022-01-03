import os, strutils, parseopt

type
  Options = object
    debianDir, package, maintainer, version, installedSize, depends, arch, desc: string

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
      of "installed-size":
        result.installedSize = val
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

proc replaceTemplate(body, package, maintainer, version, installedSize, arch, depends, desc: string): string =
  result =
    body
      .replace("{PACKAGE}", package)
      .replace("{MAINTAINER}", maintainer)
      .replace("{VERSION}", version)
      .replace("{INSTALLED_SIZE}", installedSize)
      .replace("{ARCH}", arch)
      .replace("{DEPENDS}", depends)
      .replace("{DESC}", desc)

proc formatDescription(desc: string): string =
  "Description: " & desc

proc formatDepends(depends: string): string =
  result =
    if depends != "none":
      "Depends: " & depends & "\n"
    else:
      ""

proc fixFile(file, package, maintainer, version, installedSize, arch, depends, desc: string) =
  let
    body = readFile(file)
    fixedBody = replaceTemplate(body, package=package, maintainer=maintainer,
                                version=version, installedSize=installedSize, arch=arch, depends=depends, desc=desc)
  writeFile(file, fixedBody)

let
  args = commandLineParams()
  params = getCmdOpts(args)

  controlFile = params.debianDir/"control"

  package = params.package
  maintainer = params.maintainer
  version = params.version.strip(trailing = false, chars = {'v'})
  installedSize = params.installedSize
  arch = params.arch
  depends = params.depends.formatDepends
  desc = params.desc.formatDescription

fixFile(controlFile, package=package, maintainer=maintainer, version=version,
        installedSize=installedSize, arch=arch, depends=depends, desc=desc)
