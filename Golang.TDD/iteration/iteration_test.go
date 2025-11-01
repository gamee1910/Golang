package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%s' but got '%s'", expected, repeated)
	}
}

/* Note: By default benchmarks are run sequentially.
Only the body of the loop is timed; it automatically excludes setup and cleanup code from benchmark timing.
A typical benchmark is structured like:

func Benchmark(b *testing.B) {
	//... setup ...
	for b.Loop() {
		//... code to measure ...
	}
	//... cleanup ...
}

*/

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
