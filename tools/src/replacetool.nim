import os
import strutils
import parseopt
import streams

type
  Options = object
    debianDir, package, maintainer, version, installedSize, depends, arch, desc,
      homepage, section: string

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
      of "homepage":
        result.homepage = val
      of "section":
        result.section = val
    of cmdEnd:
      assert false # cannot happen
    else:
      assert false

proc replaceTemplate(body, package, maintainer, version, installedSize, arch,
    depends, desc, homepage, section: string): string =
  result =
    body
      .replace("{PACKAGE}", package)
      .replace("{MAINTAINER}", maintainer)
      .replace("{VERSION}", version)
      .replace("{INSTALLED_SIZE}", installedSize)
      .replace("{ARCH}", arch)
      .replace("{DEPENDS}", depends)
      .replace("{DESC}", desc)
      .replace("{HOMEPAGE}", homepage)
      .replace("{SECTION}", section)

proc formatDescription*(desc: string): string =
  var strm = newStringStream(desc)
  var line: string
  var str2: seq[string]
  var i: int
  while strm.readLine(line):
    var prefix: string
    if 0 < i:
      prefix.add(" ")
    if line == "":
      prefix.add(".")
    str2.add(prefix & line)
    inc(i)
  "Description: " & str2.join("\n")

proc formatDepends(depends: string): string =
  result =
    if depends != "none":
      "Depends: " & depends & "\n"
    else:
      ""

proc formatHomepage(homepage: string): string =
  result =
    if homepage != "none":
      "Homepage: " & homepage & "\n"
    else:
      ""

proc formatSection(homepage: string): string =
  result =
    if homepage != "none":
      "Section: " & homepage & "\n"
    else:
      ""

proc fixFile(file, package, maintainer, version, installedSize, arch, depends,
    desc, homepage, section: string) =
  let
    body = readFile(file)
    fixedBody = replaceTemplate(body, package = package, maintainer = maintainer,
                                version = version,
                                installedSize = installedSize, arch = arch,
                                depends = depends, desc = desc,
                                homepage = homepage, section = section)
  writeFile(file, fixedBody)

when isMainModule:
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
    homepage = params.homepage.formatHomepage
    section = params.section.formatSection
    desc = params.desc.formatDescription

  fixFile(controlFile, package = package, maintainer = maintainer, version = version,
          installedSize = installedSize, arch = arch, depends = depends,
          desc = desc, homepage = homepage, section = section)
