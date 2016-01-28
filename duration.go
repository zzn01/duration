package duration

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/shopspring/decimal"
)

type ParseError error

var ErrMalformed ParseError = errors.New("malformed")

type parser struct {
	re   *regexp.Regexp
	unit time.Duration
}

var parsers [4]parser

func init() {

	parsers[0].unit = time.Nanosecond
	parsers[1].unit = time.Microsecond
	parsers[2].unit = time.Millisecond
	parsers[3].unit = time.Second

	re_str := []string{
		`(\d+)ns`,
		`(\d+(?:\.\d+)?)µs`,
		`(\d+(?:\.\d+)?)ms`,
		`(?:(?:(\d+)h)?(?:(\d+)m))?(\d+(?:\.\d+)?)s`,
	}

	for i, s := range re_str {
		parsers[i].re, _ = regexp.Compile("^(-)?" + s + "$")
	}
}

func Parse(s string) (time.Duration, error) {

	var d time.Duration
	var p parser
	if s[len(s)-1] == 's' {
		switch c := s[len(s)-2]; c {
		case 'n': //ns
			p = parsers[0]
		case 'µ': //µs
			p = parsers[1]
		case 'm': //ms
			p = parsers[2]
		default:
			if '0' <= c && c <= '9' { //s
				p = parsers[3]
			} else {
				return d, ErrMalformed
			}
		}
	} else {
		return d, ErrMalformed
	}

	sub := p.re.FindStringSubmatch(s)
	//	fmt.Println(len(sub), sub)
	switch len(sub) {
	case 5:
		i, _ := strconv.Atoi(sub[2])
		d += time.Duration(i) * time.Hour
		i, _ = strconv.Atoi(sub[3])
		d += time.Duration(i) * time.Minute

		f, err := decimal.NewFromString(sub[4])
		if err != nil {
			panic(err)
		}
		f = f.Mul(decimal.New(int64(p.unit), 0))
		d += time.Duration(f.IntPart())
	case 3:
		f, err := decimal.NewFromString(sub[2])
		if err != nil {
			panic(err)
		}
		f = f.Mul(decimal.New(int64(p.unit), 0))
		d += time.Duration(f.IntPart())
	default:
		return d, ErrMalformed
	}

	if sub[1] != "-" {
		return d, nil
	} else {
		return -d, nil
	}
}
