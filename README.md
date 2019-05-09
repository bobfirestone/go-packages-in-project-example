# go-packages-in-project-example

Testing how package naming in go works with multiple files in the same project

- This app expects 2 numbers to be given as command line arguments.
- The 1st number gets passed to a package called double
- The 2nd number gets passed to another package called triple

```bash
go build
./go-packages-in-project-example 5 10
=> (5 * 2) + (10 * 3) = 40
```

Each sub package is in its own directory. I tried to put it in the same as main and it complained.
