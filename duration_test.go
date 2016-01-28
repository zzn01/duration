package duration

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func test_parse(l []time.Duration, t *testing.T) {
	for _, a := range l {
		if d, err := Parse(fmt.Sprintf("%s", a)); err == nil {
			if d != a {
				t.Errorf("%s!=%s", d, a)
			}
		} else {
			t.Errorf("err:%s, %s", err, a)
		}
	}
}

func TestParse1(t *testing.T) {
	var l []time.Duration
	for i := 0; i < 15; i++ {
		l = append(l, (time.Duration(math.Pow10(i))+1)*time.Nanosecond)
	}
	test_parse(l, t)
}

func TestParse2(t *testing.T) {
	var l []time.Duration
	for i := 0; i < 15; i++ {
		l = append(l, (time.Duration(math.Pow10(i)))*time.Nanosecond)
	}
	test_parse(l, t)
}

func TestParse3(t *testing.T) {
	l := []time.Duration{
		0,
		2 * time.Hour,
		25 * time.Hour,
		59 * time.Minute,
		61 * time.Minute,
		-100 * time.Minute,
	}
	test_parse(l, t)
}

func TestParse4(t *testing.T) {
	if _, err := Parse("1000s10ms"); err == nil {
		t.Error("should be error")
	}
}

func BenchmarkParse1(b *testing.B) {
	s := fmt.Sprintf("%s", 1*time.Second+time.Nanosecond)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(s)
	}
}

func BenchmarkParse2(b *testing.B) {
	s := fmt.Sprintf("%s", 61*time.Minute)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(s)
	}
}

func BenchmarkParseMalform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse("1000s10ms")
	}
}
