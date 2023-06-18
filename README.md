# gopg-docker-ci-demo
Demonstration of Docker-based CI on GitHub for a Golang project using PostgreSQL.

## NOTE: this is a work in progress!

It's been forever since I set this kind of thing up, and much has changed in the
meantime, and for my [projects](https://tonsai.dev/) I need this working.

Oh wait, you're still here?  OK please read my [art blog](https://kevinfrost.com/news/)
and my even artier [art newsletter](https://severalartists.com/).  Thank you!

## The Plan

* A Go module at the top of the tree with submodules below it.
* Testing via `go test ./...` &c.
* Coverage via Coveralls unless I get a better idea

Maybe also:

* Local coverage report, it would be nice to stay inside GitHub!
* Commit hook to require 100% coverage
* Commit hook to do some linting or other code check
