package container

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var (
	TypeNotRegisteredError = errors.New("TypeNotRegistered")
	CircularReferenceError = errors.New("CircularReferenceDetected")
)

type threadSafeContainer struct {
	newFunctionRegistry map[reflect.Type]*NewFunction
	singletonRegistry   map[reflect.Type]interface{}
	mu                  *sync.RWMutex
}

func newThreadSafeContainer() *threadSafeContainer {
	return &threadSafeContainer{
		newFunctionRegistry: map[reflect.Type]*NewFunction{},
		singletonRegistry:   map[reflect.Type]interface{}{},
		mu:                  &sync.RWMutex{},
	}
}

type NewFunction struct {
	F         func() (interface{}, error)
	Singleton bool
}

const (
	Visiting string = "visiting"
	Done     string = "done"
)

func (c *threadSafeContainer) provide(typeNeeded reflect.Type, visited map[reflect.Type]string) (interface{}, error) {
	if ret, ok := c.singletonRegistry[typeNeeded]; ok {
		return ret, nil
	}

	if visited[typeNeeded] == Visiting {
		return nil, CircularReferenceError
	}

	visited[typeNeeded] = Visiting

	if newFunction, ok := c.newFunctionRegistry[typeNeeded]; ok {
		ret, err := newFunction.F()
		if err != nil {
			return nil, err
		}
		retType := reflect.TypeOf(ret).Elem()
		retValue := reflect.ValueOf(ret).Elem()
		for i := 0; i < retValue.NumField(); i++ {
			rt := retValue.Field(i).Type()
			if rt.Kind() != reflect.Ptr {
				continue
			}
			fieldType := retValue.Field(i).Type().Elem()
			autoWired := retType.Field(i).Tag.Get("autoWired")
			if autoWired == "true" {
				fieldValue, err := c.provide(fieldType, visited)
				if err != nil {
					return nil, err
				}
				rf := retValue.Field(i)
				if rf.CanAddr() {
					rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
					rf.Set(reflect.ValueOf(fieldValue))
				}
			}
		}

		if newFunction.Singleton {
			c.singletonRegistry[typeNeeded] = ret
		}

		visited[typeNeeded] = Done
		return ret, nil
	} else {
		return nil, TypeNotRegisteredError
	}
}

func (c *threadSafeContainer) registerType(t interface{}, f *NewFunction) {
	if t == nil && reflect.TypeOf(t).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("register object type not valid\n"))
	}
	if f == nil || f.F == nil {
		panic(fmt.Sprintf("register type %v with invalid function\n", t))
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	typeToRegister := reflect.TypeOf(t).Elem()
	if _, ok := c.newFunctionRegistry[typeToRegister]; !ok {
		c.newFunctionRegistry[typeToRegister] = f
	} else {
		panic(fmt.Sprintf("type: %s already registered\n", typeToRegister))
	}
}

func (c *threadSafeContainer) RegisterObjectSingleton(t interface{}) {
	if t == nil && reflect.TypeOf(t).Kind() != reflect.Ptr {
		panic(fmt.Sprintf("register object not valid\n"))
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.singletonRegistry[reflect.TypeOf(t).Elem()] = t
}

func (c *threadSafeContainer) RegisterTypeSingleton(t interface{}, f func() (interface{}, error)) {
	c.registerType(t, &NewFunction{f, true})
}

func (c *threadSafeContainer) RegisterType(t interface{}, f func() (interface{}, error)) {
	c.registerType(t, &NewFunction{f, false})
}

func (c *threadSafeContainer) Provide(t interface{}) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	typeNeeded := reflect.TypeOf(t).Elem()
	return c.provide(typeNeeded, map[reflect.Type]string{})
}

func (c *threadSafeContainer) MustProvide(t interface{}) interface{} {
	r, e := c.Provide(t)
	if e != nil {
		panic(fmt.Sprintf("get type %s fail", reflect.TypeOf(t)))
	}
	return r
}

var Default = newThreadSafeContainer()

// RegisterObjectSingleton register object t as singleton
func RegisterObjectSingleton(t interface{}) { Default.RegisterObjectSingleton(t) }

// RegisterTypeSingleton register t as singleton with constructor f
func RegisterTypeSingleton(t interface{}, f func() (interface{}, error)) {
	Default.RegisterTypeSingleton(t, f)
}

// RegisterType register t with constructor f
func RegisterType(t interface{}, f func() (interface{}, error)) { Default.RegisterType(t, f) }

// Provide return object with type t if registered
func Provide(t interface{}) (interface{}, error) { return Default.Provide(t) }

// Provide return object with type t, panic if not registered
func MustProvide(t interface{}) interface{} { return Default.MustProvide(t) }
