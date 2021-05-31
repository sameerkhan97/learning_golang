# Go language 
Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written 
in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go.
On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, 
it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, 
program construction, and so on, so that programs you write will be easy for other Go programmers to understand. 
This document gives tips for writing clear, idiomatic Go code. It augments the language specification, the Tour of Go, and How to Write Go Code, all of which you should read first.


## Names
* [Package names](https://github.com/sameerkhan97/learning_golang/blob/master/names/1.1packageNames.txt)
* [Getters](https://github.com/sameerkhan97/learning_golang/blob/master/names/1.2getter.txt)
* [Interface names](https://github.com/sameerkhan97/learning_golang/blob/master/names/1.3interfaceNames.txt)

## Semicolons 
* [Use of Semicolons](https://github.com/sameerkhan97/learning_golang/blob/master/semicolon/semicolonInGo.txt)

## Control structures
* [If](https://github.com/sameerkhan97/learning_golang/blob/master/control%20structures/3.1if.go)
* [Redeclaration and reassignment](https://github.com/sameerkhan97/learning_golang/blob/master/control%20structures/3.2redeclarationAndReassignment.go)
* [For](https://github.com/sameerkhan97/learning_golang/blob/master/control%20structures/3.3for.go)
* [Switch](https://github.com/sameerkhan97/learning_golang/blob/master/control%20structures/3.4switch.go)
* [Type switch](https://github.com/sameerkhan97/learning_golang/blob/master/control%20structures/3.5typeSwitch.go)

## Functions
* [Multiple return values](https://github.com/sameerkhan97/learning_golang/blob/master/functions/4.1multipleReturnValue.go)
* [Named result parameters](https://github.com/sameerkhan97/learning_golang/blob/master/functions/4.2namedReturnParameters.go)
* [Defer](https://github.com/sameerkhan97/learning_golang/blob/master/functions/4.3defer.go)

## Data
* [Allocation with new](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.1AllocationWithnew.go)
* [Allocation with make](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.2AllocationWithmake.go)
* [Arrays](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.3Array.go)
* [Slices](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.4Slices.go)
* [Maps-I](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.6.1Map.go)
* [Map-II](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.6.2Map.go)
* [Printing](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.7printing.go)
* [Append](https://github.com/sameerkhan97/learning_golang/blob/master/data/5.5Append.go)

## Initialization
* [Constants](https://github.com/sameerkhan97/learning_golang/blob/master/initialization/6.1constant.go)
* [Variables](https://github.com/sameerkhan97/learning_golang/blob/master/initialization/6.2variables.go)
* [The init function](https://github.com/sameerkhan97/learning_golang/blob/master/initialization/6.3init.go)


## Interfaces and other types
* [Interfaces](https://github.com/sameerkhan97/learning_golang/blob/master/interfaces/8.1Interfce.go)

## The blank identifier
* [The blank identifier in multiple assignment](https://github.com/sameerkhan97/learning_golang/blob/master/blank%20identifier/9.1multipleAssignment.go)
* [Unused imports and variables](https://github.com/sameerkhan97/learning_golang/blob/master/blank%20identifier/9.2unusedImportsAndVariables.go)
* [Import for side effect](https://github.com/sameerkhan97/learning_golang/blob/master/blank%20identifier/9.3importsForSideEffect.go)
* [Interface checks](https://github.com/sameerkhan97/learning_golang/blob/master/blank%20identifier/9.4interfaceChecks.go)

## Embedding
* [Embedding](https://github.com/sameerkhan97/learning_golang/blob/master/embedding/1Embedding.go)

## Concurrency
* [Goroutines](https://github.com/sameerkhan97/learning_golang/blob/master/concurrency/11.1goroutine.go)
* [Synchronization of Goroutines](https://github.com/sameerkhan97/learning_golang/blob/master/concurrency/11.2synchropnizationOfGoroutine.go)
* [Channels](https://github.com/sameerkhan97/learning_golang/blob/master/concurrency/11.3channels.go)
* [Block On Send and Recieve on Channel](https://github.com/sameerkhan97/learning_golang/blob/master/concurrency/11.4blockOnSendandRecieveOnChannel.go)
* [Buffered Channel](https://github.com/sameerkhan97/learning_golang/blob/master/concurrency/11.5bufferedChannel.go)

## Errors
* [Errors](https://github.com/sameerkhan97/learning_golang/blob/master/errors/12.1errors.go)
* [Panic](https://github.com/sameerkhan97/learning_golang/blob/master/errors/12.2panic.go)
* [Recover](https://github.com/sameerkhan97/learning_golang/blob/master/errors/12.3recover.go)
