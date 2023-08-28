.PHONY: test
bench:
	go test -v -run=none -benchmem -bench . -count 1