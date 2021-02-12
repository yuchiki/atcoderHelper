## ach test

tests sample cases

### Synopsis

tests sample cases.

The specification is not fixed. The below is the current temporal behaviour.


- build.sh is run once/
- when there exists case{n}.input for n in 1..N, tests are done for 1..N.
- in each test,
  - the command executes "cat case{n}.input | ./run.sh > case{n}.actual".
  - Then, it compares case{n}.actual and case{n}.expected.
  - If case{n}.input is "[skip ach test]\n", the case is skipped.


```
ach test [flags]
```

### Options

```
  -h, --help   help for test
```

### Options inherited from parent commands

```
      --config string        config file (default "$HOME/.ach/config.yaml")
      --task-config string   task config file (default "./achTaskConfig.yaml")
```

### SEE ALSO

* [ach](ach.md)	 - ach automates routine work you does when you participate AtCoder contests

