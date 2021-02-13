#!/bin/bash
set -eux

mkdir work
cd work

ach contest create --default-template contestFoo
cd contestFoo/A
ach test
