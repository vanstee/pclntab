# pclntab

Recover symbols from stripped go executables using information from the embedded 
[pclntab](https://docs.google.com/document/d/1lyPIbmsYbXnpNj57a261hgOYVpNRcgydurVQIyZOz_o/pub).

Intended to be used for executables built with `-ldflags="-s -w"` or stripped
via simple methods. Anything run through a tool that compresses or renames
section headers (like [upx](https://upx.github.io)) will fail. Use a more
extreme tool like [GoReSym](https://github.com/mandiant/GoReSym) in those
cases.

## Work In-progress

* [x] Support Mach-O
* [x] Support ELF
* [ ] Support ELF with -buildmode=pie
* [ ] Support PE
* [ ] Support scanning executables compressed with [upx](https://upx.github.io)
* [ ] Documented supported go versions

## Install

```
go install github.com/vanstee/pclntab/cmd/pclntab
```

## Usage

```
$ pclntab $(which kubelet) | head
0000000000402420 internal/abi.(*RegArgs).Dump
00000000004026a0 internal/abi.(*RegArgs).IntRegArgAddr
0000000000402720 internal/abi.(*IntArgRegBitmap).Set
00000000004027a0 internal/abi.(*IntArgRegBitmap).Get
0000000000402820 internal/abi.Kind.String
0000000000402880 internal/abi.(*Type).Kind
00000000004028a0 internal/abi.(*Type).HasName
00000000004028c0 internal/abi.(*Type).Pointers
00000000004028e0 internal/abi.(*Type).IfaceIndir
0000000000402900 internal/abi.(*Type).IsDirectIface
```
