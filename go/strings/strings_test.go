package benchmark

import (
	"strings"
	"testing"
)

var str string

func Benchmark_Replace(b *testing.B) {
	b.Run("strings.Replace", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			str = strings.Replace(testString, `\n`, "\n", -1)
		}
	})

	b.Run("strings.Replacer.Replace", func(b *testing.B) {
		r := strings.NewReplacer(`\n`, "\n")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			str = r.Replace(testString)
		}
	})
}

const testString = `
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
< Content-Length: 5
< Content-Type: textplain; charset=utf-8
<
* Connection #0 to host localhost left intact
Hello%
\n@euclid-1% less tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
\n@euclid-1% tail
\n@euclid-1% tail -f tmpprofiling.log
07:21:02.118448
07:21:02.199631
07:21:04.337656
07:21:05.224804
20170429 07:22:39.003091 home\nloca
inplaygroundgoprofilingmain.go:23:
20170429 07:22:39.084298 home\nloca
inplaygroundgoprofilingmain.go:28:
`
