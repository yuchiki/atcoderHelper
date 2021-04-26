#!/bin/bash
set -eux

mkdir work
cd work

# show incoming contests
ach contest incoming

# create a contest working directory
ach contest create --default-template contestFoo

# in a contest directory, ...

# test command should succeeds
cd contestFoo/A
ach test
