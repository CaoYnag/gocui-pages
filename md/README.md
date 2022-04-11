a simple markdown renderer in cui
no extra actions like collapse/expand, code edit

## TODO
### code highlight
try go-highlight.

### markdown parse
- gomarkdown/markdown
  git@github.com:russross/blackfriday.git

- yuin/goldmark

- blackfriday
  git@github.com:russross/blackfriday.git

benchmark result:
```shell
[spes@Gensoukyo bench]$ go test -bench='Benchmark*' -benchtime=2s -cpu 1,2,4,8 -v .
goos: linux
goarch: amd64
pkg: gocui-demo/test/bench
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkGoldMark
BenchmarkGoldMark                  29574             81395 ns/op
BenchmarkGoldMark-2                29354             80468 ns/op
BenchmarkGoldMark-4                28920             85170 ns/op
BenchmarkGoldMark-8                28129             85166 ns/op
BenchmarkGomarkdown
BenchmarkGomarkdown                26539             89206 ns/op
BenchmarkGomarkdown-2              26227             92031 ns/op
BenchmarkGomarkdown-4              25713             93387 ns/op
BenchmarkGomarkdown-8              24746             97752 ns/op
BenchmarkBlackfriday
BenchmarkBlackfriday               33744             70996 ns/op
BenchmarkBlackfriday-2             33568             70492 ns/op
BenchmarkBlackfriday-4             32754             73796 ns/op
BenchmarkBlackfriday-8             31867             74186 ns/op
PASS
ok      gocui-demo/test/bench   38.905s
```