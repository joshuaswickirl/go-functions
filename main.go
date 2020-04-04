package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strings"

	"github.com/joshuaswickirl/go-functions/simplemath"
)

func main() {
	answer, err := simplemath.Divide(6, 3)
	if err != nil {
		fmt.Printf("an error occured %s\n", err.Error())
	} else {
		fmt.Printf("%f\n", answer)
	}

	numbers := []float64{12.2, 14, 16, 22.4}
	total := simplemath.Sum(numbers...)
	fmt.Printf("total of sum: %f\n", total)

	sv := simplemath.NewSemanticVerison(1, 2, 3)
	sv.IncrementMajor()
	p := &sv
	p.IncrementMajor()
	println(sv.String())

	tripper := &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com",
		strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
	println(tripper.count)
	_, _ = tripper.RoundTrip(r)
	println(tripper.count)

	anonymous()

	expr := mathExpression(SubtractExpr)
	println(expr(2, 3))
	fmt.Printf("%f\n", double(3, 2, mathExpression(MultiplyExpr)))

	p2 := powerOfTwo()
	value := p2()
	println(value)
	value = p2()
	println(value)
	value = p2()
	println(value)

	var funcs []func() int64
	for i := 0; i < 10; i++ {
		cleanI := i
		funcs = append(funcs, func() int64 {
			return int64(math.Pow(float64(cleanI), 2))
		})
	}
	for _, f := range funcs {
		println(f())
	}

	if err := ReadFullFile(); err != nil {
		fmt.Printf("something bad occured: %s\n", err)
	}

}

func ReadFullFile() (err error) {
	var r io.ReadCloser = &SimpleReader{}
	defer func() {
		_ = r.Close()
		if p := recover(); p == errCatastrophicReader {
			println(p)
			err = errors.New("a panic occurred but it is okay")
		} else if p != nil {
			panic("an unexpected error occured and we do not want to recover")
		}
	}()
	defer func() {
		println("before for-loop")
	}()
	for {
		value, readerErr := r.Read([]byte("text that does something"))
		if err == io.EOF {
			println("finished reading file, breaking out of loop")
			break
		} else if readerErr != nil {
			return readerErr
		}
		println(value)
	}
	defer func() {
		println("after for-loop")
	}()
	return nil
}

func ReadSomethingBad() error {
	var r io.Reader = BadReader{errors.New("my nonsense reader")}
	value, err := r.Read([]byte("test something"))
	if err != nil {
		fmt.Printf("an error occured %s", err)
		return err
	}
	println(value)
	return nil
}

type BadReader struct {
	err error
}

func (br BadReader) Read(p []byte) (n int, err error) {
	return -1, br.err
}

type SimpleReader struct {
	count int
}

var errCatastrophicReader = errors.New("something catastrophic occured in the reader")

func (br *SimpleReader) Read(b []byte) (n int, err error) {
	if br.count == 2 {
		panic(errors.New("just another error"))
	}
	if br.count > 3 {
		return 0, io.EOF //errors.New("random error")
	}
	br.count++
	return br.count, nil
}

func (br *SimpleReader) Close() error {
	println("closing reader")
	return nil
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

func anonymous() {
	a := func(name string) string {
		fmt.Printf("my first %s function\n", name)
		return name
	}
	value := a("anonymous")
	println(value)
}

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	switch expr {
	case AddExpr:
		return simplemath.Add
	case SubtractExpr:
		return simplemath.Subtract
	case MultiplyExpr:
		return simplemath.Multiply
	default:
		panic("an invalid math expression was used")
	}
}

func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}

func powerOfTwo() func() int64 {
	x := 1.0
	return func() int64 {
		x++
		return int64(math.Pow(x, 2))
	}
}
