# Milestones

# done

# v0.1

- It can test testcases
  - `ach test`
- It can create a contest directory with a hard-coded contest template.
   - `ach create --defaulte-template foo`
- It can intialize each task according to the setting of the given config file.


# v0.0

- autorelease by tagging
- documented in Japanese

# not yet

# v0.2

- It can create a contest directory with specified contest templates
  - `ach contest create -t foo.yaml <name>`
  - `ach contest create ---default-contest <name>`
- It has an integration test running on a container.

# v0.3

- it can spread task-template
  - `ach contest create ...`
- it can use multiple templates and languages

# v0.4

- it can fetch the list of contests
  - `ach contest upcoming`
  - `ach contest list <matcher>`

# v0.5

- it can create a contest directory according to fetched information
  - `ach contest crete <abc001>`

# v0.6

- it can login and logout
  - `ach login`

# v0.7

- it can submit and fetch testCases
