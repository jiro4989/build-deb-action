import unittest
import ../src/replacetool

suite "proc replaceTemplate":
  test "ok: 1 line":
    let s = "sample"
    let got = s.formatDescription()
    let want = "Description: sample"
    check want == got

  test "ok: multiline":
    let s = """sample

sample2"""
    let got = s.formatDescription()
    let want = """Description: sample
 .
 sample2"""
    check want == got

  test "ok: multiline2":
    let s = """sample

sample2
sample3

sample4
"""
    let got = s.formatDescription()
    let want = """Description: sample
 .
 sample2
 sample3
 .
 sample4"""
    check want == got
