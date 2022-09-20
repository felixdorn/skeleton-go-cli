# go-skeleton

This is a skeleton for a CLI. It is intended to be used as a starting point for new CLIs.

## Usage

```bash
$ git clone git@github.com:felixdorn/go-skeleton.git --depth 1 --branch main --single-branch
$ ./configure.sh
Repository URL (): 
Binary name (myapp): 
Author name (Your Name): 
Author email (Your Email): 
Modify files? (y/N): 
```

This project uses Taskfile for building and testing. You can find the available tasks by running `task -l`.

To run your app, you may:

* Run it directly `go run github.com/owner/repository/cmd/myapp`
* Alias it `alias $(task alias)` and run it as `myapp`
* Build it using `task build` and run it as `./bin/myapp`

