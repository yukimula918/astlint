# astlint
A pure syntax based static analyzer with linters to identify code bugs.

## project structure

The following structures the top-level modules in `astlint`.

```
astlint
|-- cli                     # command-line entries
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
|-- |-- test                    ## unit test layer
|-- |-- util                    ## infrastructure layer
|-- script                  # script for running linter
|-- go.mod                  # module dependency defined
|-- Makefile                # binary building script
```

The package dependency strictly follows the following orders (up->down)

```
cli/*
pkg/test
pkg/linter/golang  pkg/linter/java pkg/linter/python ...
pkg/linter
pkg/astree         pkg/fsutil      pkg/report      
pkg/util           pkg/logging     pkg/ruleset
pkg/domain
```
