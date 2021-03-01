# Build Instructions
If you would like to build the code yourself, here are the steps. There were some oddities in building for Windows because of the sqlite3 package and the fact that it uses cgo.

## Linux
Linux is very simple. A standard `go build main.go` will do the trick. The sqlite3 github says you will need the `CGO_ENABLED=1` environment variable set, but it compiled without issue for me with no additional options.

## Windows
Cross-compiling for Windows from Linux should be possible, but I could not get it working. Building directly on Windows was much easier, but still more complicated than Linux... *what else is new*

According the cross-compile directions in the go-sqlite3 repo, you should be able to [cross compile](https://github.com/mattn/go-sqlite3#cross-compile).
I first tried just using the standard method of `env GOOS=windows GOARCH=amd64 go build main.go`
This did compile, but when I tested the resulting .exe on Windows, it was giving me erors about bindings because I was missing that CGO_ENABLED option mentioned above.
So I then tried `env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build main.go`
This would not compile. It was giving an error about unrecognized command line option ‘-mthreads’. I found the issue called out in the go-sqlite repo [Issue 303](https://github.com/mattn/go-sqlite3/issues/303)
The fix identified was to install some gcc's and add some more command line options. That did not work for me.

I decided it would be easier to just build in Windows. You will need a gcc installed. I tried a couple but the only one that actually worked was the one called out on the go-sqlite3 site... *image that* one day I will just read and make my life easier.
Following the directions pulled directly from the [go-sqlite3 repo](https://github.com/mattn/go-sqlite3#windows) you will need to:
1. Install a Windows gcc toolchain. [This is the one that worked](https://jmeubank.github.io/tdm-gcc/)
2. Add the bin folders to the Windows path if the installer did not do this by default.
3. Open a terminal for the TDM-GCC toolchain, can be found in the Windows Start menu.
4. Navigate to your project folder and run the `go build main.go` command for this package.
This gave me a working .exe

## Apple
I do not have a Mac currently to build on and I tried to cross-compile from linux to darwin, but it did not work and I learned a lesson from Windows that I should just compile on the target OS. [Here are the steps](https://github.com/mattn/go-sqlite3#mac-osx)