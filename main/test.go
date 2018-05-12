package main

import (
	"encoding/json"
	"fmt"
)

type Foo interface {
	Bar() string
}

type Bar struct {
	bar string
}

func (b *Bar) Bar() string {
	return b.bar
}

func (b *Bar) MarshalJSON() ([]byte, error) {
	return json.Marshal(BarJSON{b.bar})
}

func (b *Bar) UnmarshalJSON(data []byte) error {
	tmp := &BarJSON{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	b.bar = tmp.Bar
	return nil
}

type BarJSON struct {
	Bar string
}

type FooBar struct {
	H    int
	Foos map[string]Foo
}

func main() {
	b := &Bar{"butts"}
	var f Foo
	if err := JSONTester(b, f); err != nil {
		panic(err)
	}

	fooMap := map[string]Foo{
		"Booty": b,
	}
	var f2 map[string]Foo
	if err := JSONTester(fooMap, f2); err != nil {
		panic(err)
	}

	fooBar := FooBar{3, fooMap}
	var f3 FooBar
	if err := JSONTester(fooBar, f3); err != nil {
		panic(err)
	}

	fooBarMap := map[string]FooBar{
		"Fooer": fooBar,
	}
	var f4 map[string]FooBar
	if err := JSONTester(fooBarMap, f4); err != nil {
		panic(err)
	}

}

func JSONTester(obj interface{}, out interface{}) error {
	fmt.Println("pre ", obj)

	byt, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	fmt.Println("byte", string(byt))

	n := out
	err = json.Unmarshal(byt, &n)
	if err != nil {
		return err
	}
	fmt.Println("post", string(byt))
	fmt.Println()
	return nil
}
