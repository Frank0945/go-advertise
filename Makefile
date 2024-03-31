#! /usr/bin/make

MODULE=$(shell go list -m)

.PHONY: gen
gen:
	@echo GENERATING CODE...
	@goa gen $(MODULE)/api/design -o api