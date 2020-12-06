SHELL := /bin/bash

.PHONY: bench
bench: bench-v0

.PHONY: report
bench-report: bench-v0-report

.PHONY: bench-v0
bench-v0:
	@echo "Running Benchmarks.  This will take a while."
	@echo "  Int Funcs"
	@go test -bench=. -count=30 -timeout=2h v0/bytify/int-write_bench_test.go > tmp.txt
	@tail -n+3 tmp.txt \
	  | head -n-2 \
	  | sed -e 's/ //g;s/_/\t/g;s/-[0-9]\+//g;s/ns\/op//g' \
	  | awk '{split($$1, f, "	"); print >> "docs/benchmarks/" f[1] ".tsv"}'
	@rm tmp.txt
	@echo "  Uint Funcs"
	@go test -bench=. -count=30 -timeout=2h v0/bytify/uint-write_bench_test.go > tmp.txt
	@tail -n+3 tmp.txt \
	  | head -n-2 \
	  | sed -e 's/ //g;s/_/\t/g;s/-[0-9]\+//g;s/ns\/op//g' \
	  | awk '{split($$1, f, "	"); print >> "docs/benchmarks/" f[1] ".tsv"}'
	@rm tmp.txt

bench-v0-report:
	@echo Generating Averages Report
	@go run v0/tools/bench/merge.go > averages.tsv