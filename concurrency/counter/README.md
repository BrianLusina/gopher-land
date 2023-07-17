# Data Race checker

An example of how data race can be found in code via tests. This shows a simple example where 5 goroutines are launched all trying to
update a single `counter` variable.

A test has been written to test that sometimes some of the writes from some of the goroutines are lost while others succeed which means this code
will not always be correct.

In order to check this, the test can be run with:

```shell
go test --race .
```

> Running this in the current directory should give an output like below:

```plain
==================
WARNING: DATA RACE
Read at 0x00c00001e248 by goroutine 11:
  gopherland/counter.getCounter.func1()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:13 +0x46

Previous write at 0x00c00001e248 by goroutine 12:
  gopherland/counter.getCounter.func1()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:13 +0x58

Goroutine 11 (running) created at:
  gopherland/counter.getCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:11 +0x8d
  gopherland/counter.TestGetCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter_test.go:6 +0x2b
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1576 +0x216
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1629 +0x47

Goroutine 12 (running) created at:
  gopherland/counter.getCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:11 +0x8d
  gopherland/counter.TestGetCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter_test.go:6 +0x2b
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1576 +0x216
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1629 +0x47
==================
==================
WARNING: DATA RACE
Write at 0x00c00001e248 by goroutine 9:
  gopherland/counter.getCounter.func1()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:13 +0x58

Previous write at 0x00c00001e248 by goroutine 11:
  gopherland/counter.getCounter.func1()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:13 +0x58

Goroutine 9 (running) created at:
  gopherland/counter.getCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:11 +0x8d
  gopherland/counter.TestGetCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter_test.go:6 +0x2b
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1576 +0x216
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1629 +0x47

Goroutine 11 (running) created at:
  gopherland/counter.getCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter.go:11 +0x8d
  gopherland/counter.TestGetCounter()
      /home/lusinabrian/Projects/brianlusina/gopher-land/concurrency/counter/counter_test.go:6 +0x2b
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1576 +0x216
  testing.(*T).Run.func1()
      /usr/local/go/src/testing/testing.go:1629 +0x47
==================
--- FAIL: TestGetCounter (0.00s)
    counter_test.go:8: unpexpected counter 3542
    testing.go:1446: race detected during execution of test
FAIL
FAIL    gopherland/counter      0.014s
FAIL
```

> Note that this output might not entirely match yours but should be somewhat similar
