# astlint
A pure syntax based static analyzer with linters to identify code bugs.

### Project Structure

The project is structured by following packages.

```
astlint
|-- cmd                     ## command-line entries
|-- domain                  ## domain model module
|-- parser                  ## syntax parse module
|-- static                  ## static analysis module
|-- linter                  ## bug linter module
|-- test                    ## unit test module
|-- util                    ## infrastructure module
|-- go.mod                  ## go module dependency 
|-- Makefile                ## building script
```
