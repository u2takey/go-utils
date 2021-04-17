package container

import (
	"reflect"
	"testing"
	"time"

	"github.com/u2takey/go-utils/json"
)

type A struct {
	Aint int
}

type B struct {
	Aint int
	Aa   *A `autoWired:"true"`
}

type D struct {
	bb   *B `autoWired:"true"`
	Dint int
}

type E struct {
	Ff *F `autoWired:"true"`
}

type F struct {
	Ee *E `autoWired:"true"`
}

type G struct {
}

type H struct {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Provide(tt.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotb, _ := json.Marshal(got)
			wantb, _ := json.Marshal(tt.want)
			t.Logf("Provide() got = %v, want %v", string(gotb), string(wantb))
			if !reflect.DeepEqual(gotb, wantb) {
				t.Errorf("Provide() got = %v, want %v", string(gotb), string(wantb))
			}
		})
	}
}
