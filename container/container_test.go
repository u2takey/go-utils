package container

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/u2takey/go-utils/json"
	"github.com/u2takey/go-utils/rand"
	"github.com/stretchr/testify/assert"
)

type A struct {
	Aint int
}

type B struct {
	Aint int
	Aa   *A `autowired:"true"`
}

type D struct {
	bb   *B `autowired:"true"`
	Dint int
}

type E struct {
	Ff *F `autowired:"true"`
}

type F struct {
	Ee *E `autowired:"true"`
}

type G struct {
}

type H struct {
}

type I struct {
}

type ExampleA struct {
	d map[string]string
}

func (a *ExampleA) SayHello() {
	fmt.Printf("say hello from: %+v\n", a)
}

type SayHelloInterface interface {
	SayHello()
}

// TestContainerExample1, 注册 type + new 函数
func TestContainerExample1(t *testing.T) {
	c := newThreadSafeContainer()

	c.RegisterType(&ExampleA{}, func() (interface{}, error) {
		t.Logf("ExampleA's new function called")
		return &ExampleA{d: map[string]string{"1": "2"}}, nil
	})
	ret, err := c.Provide(&ExampleA{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"1": "2"}, ret.(*ExampleA).d)
}

// TestContainerExample2, 注册 单例 type + new 函数
func TestContainerExample2(t *testing.T) {
	c := newThreadSafeContainer()

	c.RegisterType(&ExampleA{}, func() (interface{}, error) {
		t.Logf("ExampleA's new function called")
		return &ExampleA{d: map[string]string{"1": "2"}}, nil
	})
	ret, err := c.Provide(&ExampleA{})
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"1": "2"}, ret.(*ExampleA).d)
}

// TestContainerExample3, 注册单例 object
func TestContainerExample3(t *testing.T) {
	c := newThreadSafeContainer()
	exampleA := &ExampleA{d: map[string]string{"1": rand.String(10)}}
	c.RegisterObjectSingleton(exampleA)
	ret, err := c.Provide(&ExampleA{})
	assert.Nil(t, err)
	ret2, err := c.Provide(&ExampleA{})
	assert.Nil(t, err)
	assert.Equal(t, ret, ret2)
}

// TestContainerExample4, 注册单例 type + new 函数, provide 的时候用 interface 获取
func TestContainerExample4(t *testing.T) {
	c := newThreadSafeContainer()

	c.RegisterType(&ExampleA{}, func() (interface{}, error) {
		t.Logf("ExampleA's new function called")
		return &ExampleA{d: map[string]string{"1": "2"}}, nil
	})
	ret, err := c.Provide((*SayHelloInterface)(nil))
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"1": "2"}, ret.(*ExampleA).d)

	ret2, err := c.Provide(reflect.TypeOf((*SayHelloInterface)(nil)).Elem())
	assert.Nil(t, err)
	assert.Equal(t, ret, ret2)
}

type ExampleB struct {
	d         map[string]string
	exampleA1 *ExampleA         `autowired:"true"`
	exampleA2 SayHelloInterface `autowired:"true"`
}

func (a *ExampleB) Foo() {
	fmt.Printf("foo from: %+v\n", a)
}

type FooInterface interface {
	Foo()
}

// TestContainerExample5, 注册单例 type + new 函数, provide 的时候用 interface 获取, 同时使用 autoWired 自动注入元素
func TestContainerExample5(t *testing.T) {
	c := newThreadSafeContainer()

	c.RegisterTypeSingleton(&ExampleA{}, func() (interface{}, error) {
		t.Logf("ExampleA's new function called")
		return &ExampleA{d: map[string]string{"1": rand.String(5)}}, nil
	})

	c.RegisterTypeSingleton(&ExampleB{}, func() (interface{}, error) {
		t.Logf("ExampleB's new function called")
		return &ExampleB{d: map[string]string{"3": "4"}}, nil
	})

	ret, err := c.Provide((*FooInterface)(nil))
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"3": "4"}, ret.(*ExampleB).d)
	assert.Equal(t, ret.(*ExampleB).exampleA1, ret.(*ExampleB).exampleA2)
}

func Test_threadSafeContainer_Provide(t *testing.T) {
	c := newThreadSafeContainer()
	c.RegisterType(&A{}, func() (interface{}, error) {
		return &A{1}, nil
	})
	c.RegisterType(&B{}, func() (interface{}, error) {
		t.Logf("B's new function called")
		return &B{Aint: 2}, nil
	})
	c.RegisterTypeSingleton(&D{}, func() (interface{}, error) {
		t.Logf("D's new function called")
		return &D{Dint: time.Now().Nanosecond()}, nil
	})
	c.RegisterType(&E{}, func() (interface{}, error) {
		return &E{}, nil
	})
	c.RegisterType(&F{}, func() (interface{}, error) {
		return &F{}, nil
	})
	d, _ := c.Provide(&D{})

	h := &H{}
	c.RegisterObjectSingleton(h)

	c.RegisterType(&I{}, func() (interface{}, error) {
		return nil, errors.New("new I failed")
	})
	tests := []struct {
		name    string
		t       interface{}
		want    interface{}
		wantErr bool
	}{
		{
			"test1",
			&A{},
			&A{1},
			false,
		},
		{
			"test2",
			&B{},
			&B{Aint: 2, Aa: &A{1}},
			false,
		},
		{
			"test3",
			&D{},
			d,
			false,
		},
		{
			"test4",
			&D{},
			d,
			false,
		},
		{
			"test5",
			&E{},
			nil,
			true,
		},
		{
			"test6",
			&G{},
			nil,
			true,
		},
		{
			"test7",
			&H{},
			h,
			false,
		},
		{
			"test8",
			&I{},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Provide(tt.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotB, _ := json.Marshal(got)
			wantB, _ := json.Marshal(tt.want)
			t.Logf("Provide() got = %v, want %v", string(gotB), string(wantB))
			if !reflect.DeepEqual(gotB, wantB) {
				t.Errorf("Provide() got = %v, want %v", string(gotB), string(wantB))
			}
		})
	}
}
