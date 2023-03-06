
test:
	go test -v

bench:
	go test -bench=. -v

cpu-profile:
	GOGC=off go test -bench=Benchmark2Base_1 -cpuprofile cpu.out

mem-profile:
	GOGC=off go test -bench=Benchmark2Base_1 -memprofile mem.out

prof-svg: cpu-profile
	rm -f profile001.svg
	go tool pprof -svg 2base.test cpu.out
	firefox profile001.svg

prof-disasm: cpu-profile
	go tool pprof -disasm Benchmark2Base_1 2base.test cpu.out

prof-list: cpu-profile
	go tool pprof -list Benchmark2Base_1 2base.test cpu.out

prof: cpu-profile
	go tool pprof 2base.test cpu.out

clean:
	rm -f *.svg *.pdf *.gz 2base.test cpu.out
