# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

.DEFAULT_GOAL := help

# let's keep this Makefile as minimal as possible, and put all goals in dedicated ./make/*.mk files
include ./make/*.mk
