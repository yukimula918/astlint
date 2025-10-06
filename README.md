# astlint
A pure syntax based static analyzer with linters to identify code bugs.

## project structure

The following structures the top-level modules in `astlint`.

```
astlint
|-- api                     # server API handlers
|-- cmd                     # command-line entries
|-- |-- astlint                 ## binary entry
|-- |-- server                  ## server entry 
|-- conf                    # configuration files
|-- pkg                     # implementation packages
|-- |-- astree                  ## syntax tree layer
|-- |-- fsutil                  ## file io layer
|-- |-- linter                  ## static analysis layer
|-- |-- |-- cpp                     ### C/C++ analyzer & linter
|-- |-- |-- java                    ### Java analyzer & linter
|-- |-- |-- golang                  ### Golang analyzer & linter
|-- |-- |-- python                  ### Python analyzer & linter
|-- |-- logging                 ## logging layer
|-- |-- report                  ## issue report layer
|-- |-- ruleset                 ## linting rule layer
|-- |-- util                    ## infrastructure layer
|-- script                  # script for running linter
|-- test                    # unit test layer
|-- go.mod                  # module dependency defined
|-- Makefile                # binary building script
```

The package dependency strictly follows the following orders (up->down).

```
cli/*
test/*
pkg/linter/golang  pkg/linter/java pkg/linter/python ...
pkg/linter
pkg/astree         pkg/fsutil      pkg/report      
pkg/logging        pkg/ruleset
pkg/domain         pkg/util
```

NOTE: Packages in the same layer should NOT depend on each other.

## build & running

Using the following command to build binaries of `astlint`.

```
make build
```
