## ach contest create

creates contest directory

### Synopsis

creates contest directory.
Temporarily, current template directory is hard-coded as $HOME/projects/private/atcoder/D
D is for directory.
		

```
ach contest create [contestName] [flags]
```

### Options

```
  -d, --default-template   (required) use default contest template
  -h, --help               help for create
      --open-editor        open editor for each task (default true)
```

### Options inherited from parent commands

```
      --config string        config file (default "$HOME/.ach/config.yaml")
      --task-config string   task config file (default "./achTaskConfig.yaml")
```

### SEE ALSO

* [ach contest](ach_contest.md)	 - manipulates an AtCoder contest

