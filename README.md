# locker [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov]

Reader/writer mutual exclusion lock applied to certain resources.

## Installation

```shell
go get -u github.com/tracerun/locker
```

## Usage

```go
// create an instance
lock := New()
// read lock resource
release := lock.ReadLock("resource1")
... 
// release the read lock
relaase()
// write lock resource
release := lock.WriteLock("resource2")
... 
// release the write lock
relaase()
```

[ci-img]: https://travis-ci.org/tracerun/locker.svg?branch=master
[ci]: https://travis-ci.org/tracerun/locker
[cov-img]: https://coveralls.io/repos/github/tracerun/locker/badge.svg?branch=master
[cov]: https://coveralls.io/github/tracerun/locker?branch=master