package container

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"unsafe"
)

var (
	autoWireTagName = "autowired"
)

var (
	TypeNotRegisteredError      = errors.New("TypeNotRegistered")
	TypeAutoWireNotSupportError = errors.New("TypeAutoWireNotSupportError")
	CircularReferenceError      = errors.New("CircularReferenceDetected")
)

type threadSafeContainer struct {
	newFunctionRegistry map[reflect.Type]*NewFunction
	singletonRegistry   map[*NewFunction]interface{}
	mu                  *sync.RWMutex
}

func newThreadSafeContainer() *threadSafeContainer {
	return &threadSafeContainer{
		newFunctionRegistry: map[reflect.Type]*NewFunction{},
		singletonRegistry:   map[*NewFunction]interface{}{},
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

// findNewFunction find newFunction for type
func (c *threadSafeContainer) findNewFunction(typeNeeded reflect.Type) (*NewFunction, bool) {
	// find by type
	if newFunction, ok := c.newFunctionRegistry[typeNeeded]; ok {
		return newFunction, ok
	}
	// if not found, find first implementation
	for k := range c.newFunctionRegistry {
		if typeNeeded.Kind() == reflect.Interface && k.Implements(typeNeeded) {
			return c.newFunctionRegistry[k], true
		}
	}
	return nil, false
}

func (c *threadSafeContainer) provide(typeNeeded reflect.Type, visited map[reflect.Type]string) (interface{}, error) {
	if visited[typeNeeded] == Visiting {
		return nil, CircularReferenceError
	}

	visited[typeNeeded] = Visiting

	if newFunction, ok := c.findNewFunction(typeNeeded); ok {
		if ret, ok := c.singletonRegistry[newFunction]; ok {
			return ret, nil
		}
		ret, err := newFunction.F()
		if err != nil {
			return nil, err
		}
		retType := reflect.TypeOf(ret).Elem()
		retValue := reflect.ValueOf(ret).Elem()
		for i := 0; i < retValue.NumField(); i++ {
			rt := retValue.Field(i).Type()
			fieldType := retValue.Field(i).Type()
			autoWired := retType.Field(i).Tag.Get(autoWireTagName)
			if autoWired == "true" {
				if rt.Kind() != reflect.Ptr && rt.Kind() != reflect.Interface {
					return nil, fmt.Errorf("%w, type: %s", TypeAutoWireNotSupportError, fieldType)
				}
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
			c.singletonRegistry[newFunction] = ret
		}

		visited[typeNeeded] = Done
		return ret, nil
	} else {
		return nil, fmt.Errorf("%w, type: %s", TypeNotRegisteredError, typeNeeded)
	}
}

func checkTypeAndFunction(t interface{}, f *NewFunction, checkFunc bool) reflect.Type {
	var rtt reflect.Type
	trt := reflect.TypeOf(t)
	if rt, ok := t.(reflect.Type); ok {
		rtt = rt
	} else if trt.Kind() == reflect.Ptr {
		// 对于 Interface 类型特殊处理，存储 Interface, 而不是 *Interface
		if trt.Elem().Kind() == reflect.Interface {
			rtt = trt.Elem()
		} else {
			rtt = trt
		}
	} else if trt.Kind() == reflect.Interface {
		rtt = trt
	}
	if checkFunc {
		if f == nil || f.F == nil {
			panic(fmt.Sprintf("register type %v with invalid function\n", t))
		}
	}
	return rtt
}

func (c *threadSafeContainer) registerType(t interface{}, f *NewFunction) {
	c.mu.Lock()
	defer c.mu.Unlock()

	typeToRegister := checkTypeAndFunction(t, f, true)
	if _, ok := c.newFunctionRegistry[typeToRegister]; !ok {
		c.newFunctionRegistry[typeToRegister] = f
	} else {
		panic(fmt.Sprintf("type: %s already registered\n", typeToRegister))
	}
}

func (c *threadSafeContainer) RegisterObjectSingleton(t interface{}) {
	c.registerType(t, &NewFunction{func() (interface{}, error) {
		return t, nil
	}, true})
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

	typeNeeded := checkTypeAndFunction(t, nil, false)
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

// RegisterObjectSingleton register object t as singleton, t should be Ptr Type Or Interface
func RegisterObjectSingleton(t interface{}) { Default.RegisterObjectSingleton(t) }

// RegisterTypeSingleton register t as singleton with constructor f, t should be Ptr Type Or reflect.Type
func RegisterTypeSingleton(t interface{}, f func() (interface{}, error)) {
	Default.RegisterTypeSingleton(t, f)
}

// RegisterType register t with constructor f, t should be Ptr Type Or reflect.Type, f should return Ptr or Interface
func RegisterType(t interface{}, f func() (interface{}, error)) { Default.RegisterType(t, f) }

// Provide return object with type t if registered, t should be Ptr Type Or reflect.Type
func Provide(t interface{}) (interface{}, error) { return Default.Provide(t) }

// Provide return object with type t, panic if not registered, t should be Ptr Type Or reflect.Type
func MustProvide(t interface{}) interface{} { return Default.MustProvide(t) }
