New in this branch => [go-intro-02](https://github.com/jackterranova/go-sandbox/blob/go-intro-02/go-intro-02.md) 

# Go Sandbox

A repo for learning Go based on the docs from https://golang.org.

The master branch of this repo represents an all-in-one view of learning Go.  

But you can follow a step-wise path by examining each of the repo's branches individually, starting with branch [go-intro-01](https://github.com/jackterranova/go-sandbox/tree/go-intro-01).  

Each subsequent branch builds on the previous, layering more Go information and features. 


## Go Primer for Java Wranglers

### Go code organization - `$GOPATH`

If you are coming from a JVM background like me, one of the first oddities you'll grapple with is code organization.  

Even in non-jvm projects, we typically organize our projects in isolation.  Maybe there are projects that depend on others and they are related through a build tool or an IDE, but physically they stand alone in their own directory structure.  

> __Go programmers typically put all of their Go source (even across multiple Go projects/repositories) in a single workspace specified by `$GOPATH` on their local machines.__  

When building Go code, the `go` tool uses `$GOPATH` to locate all Go source code.  

For example, `go install github.com/jackterranova/hello`, builds and installs the source code under `$GOPATH/github.com/jackterranova/hello`, where `github.com/user/hello` is the Go package being built and installed. 

By default, `$GOPATH` is `$HOME/go` (the users local home directory).  

So if you want to cut right to writing Go code and worry about pariculars later, you can start writing and building Go code under `go` in your user's local root directory.  

Otherwise set `$GOPATH` explicitly. 

At the root of the Go workspace are 2 directories: `bin` and `src`.  Each project under `src` is source controlled independently, but, locally, all of the Go source and binary is in the same physical directory.

This has critical implications due the way the Go import path and $GOPATH works. 

### Go import path

Java package naming conventions recommend package names prefixed with what is more or less the URN of the author or an organization that the author works for in reverse order: `com.example.tools`.   This is simply to avoid package collisions when building code from different sources.  

Go retains the package concept, but 

> __in Go, the package is actually used to physically locate the source code: `github.com/jackterranova/go-sandbox/hello`, but the fully qualifed package is NOT part of the physical struture under source control.  This is a significant departure from standard Java package naming conventions__


```bash
~/go/
    bin/
        hello
        goodbye
    src/
        github.com/jackterranova/               # <-- repo at github.com
                                hello/.         # <-- begin dir structure under source control 
                                     hello.go.  # ==> maps to https://github.com/jackterranova/hello/hello.go

        bitbucket.org/jack-terranova/           # <-- repo at bitbucket.org
                                goodbye/        # <-- begin dir structure under source control 
                                     goodbye.go # ==> maps to https://bitbucket.org/jack-terranova/goodbye/goodbye.go
```

__Note the fully qualified package names, including the source control remote server__

Contrast this with how a typical Java project might be set up ...

```bash
~/MyJavaApp/               # <-- repo somewhere (?)
           build/
                ...
           src/            # <-- begin dir structure under source control 
                com/example/tools/
                                 Utils.java
~/MyOtherJavaApp/          # <-- repo somewhere (?)
           build/
                ...
           src/            # <-- begin dir structure under source control 
                org/example/models/
                                 User.java					
```

__The first project might be located at `github.com/jackterranova/MyJavaApp` with `src/com/example/tools` under source control__  

> __In Go, part of the physical package structure is OUTSIDE of source control yet it is referenced in the source.__

## Getting Go source from a remote/`go get`

> __Because Go packages are used to physically locate remote source code repositories, `go get` can be used to clone, build and install remote dependencies__

```bash
go get github.com/golang/example/hello
```

__Thanks to the Go import path, this will clone, build and install Golang's hello example and all of its dependencies under $GOPATH__

## Running Go programs

One of the really neat things about Go is that it is part of the C programming family.  As such, it compiles to executable binary just like a C program would and runs natively on the local machine.   

__Ah-ha! So its not portable ISSS it?__

Sort of... Yes, Java is a compile once, run anywhere language thanks to the JVM, and, no...

> __you can not build a Go binary on a Mac and then run it on a Windows machine__

But...

> __you can build Go binaries for *any* platform from the single platform in which you are working using the `$GOOS` variable__

... (more on that later).

So, since Go binaries run natively ...

> __you can call Go binaries directly ... they are fully executable - no classpaths or additional runtime redtape is needed__

Additonally, it is recommended that you set your `$PATH` variable to include `$GOPATH/bin` ... e.g. `export GOPATH=$(go env GOPATH)`

## Quick HelloWorld example 

1. `mkdir $GOPATH/src/github.com/<your github user name>/go-sandbox/hello`. <-- if you plan to use github.com for your remote

2. Create `hello.go` in the new directory ...

```
package main //<- must use `main` when creating executable

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

3. `go install github.com/<your github user name>/go-sandbox/hello` --note you can run this anywhere thanks to `$GOPATH`

4. `$GOPATH/bin/hello` or `hello` if $GOPATH/bin is in your executable path.

__Note the `main` package declared in hello.go despite the fact that the source is located in `github.com/<your github user name>/go-sandbox/hello`__ 

Huh?  

> __When writing Go excecutables, you must use `package main`.  Using any other package name will make your compile-output a non-executable library.__  

__Although `hello.go` declares `package main`, the source is physically located locally under `$GOPATH/src/github.com/<your github user name>/go-sandbox/hello`.  Remotely, the source is located at https://github.com/<your github user name>/go-sandbox/hello__


