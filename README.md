# sarulabsdi  
This is a baseline drop of [v2.4.2 of github.com/sarulabsdi](https://github.com/sarulabs/di/releases/tag/v2.4.2)   


# New Features
The following are non breaking changes to sarulabs/di.  

## SafeGetByType  
```go
// SafeGetByType retrieves an array of objects from the Container.
// The objects have to belong to this scope or a more generic one.
// If the objects do not already exist, it is created and saved in the Container.
// If the objects can not be created, it returns an error.
SafeGetByType(rt reflect.Type) ([]interface{}, error)
```

## GetByType  
```go
// GetByType is similar to SafeGetByType but it does not return the error.
// Instead it panics.
GetByType(rt reflect.Type) []interface{}
```
Just like dotnetcore, I would like to register a bunch of services that all implement the same interface.  I would like to get back an array of all registered objects, and to do that I need to ask for them by type

For efficiency, type validation is done during registration.  After that it should just be lookups.   
When you do a ```GetByType```, don't keep calling reflect to get a type of what is a compliled type.  do it once and put it in a map.  



### Registration

type [DEF](https://github.com/fluffy-bunny/sarulabsdi/blob/8a200c4fa3aefa0a28ddc66739aac1631f2a95aa/definition.go#L19) struct  

```go
// Def contains information to build and close an object inside a Container.
type Def struct {
	Build            func(ctn Container) (interface{}, error)
	Close            func(obj interface{}) error
	Name             string
	Scope            string
	Tags             []Tag
	Type             reflect.Type
	ImplementedTypes TypeSet
	Unshared         bool
}
```
Two new fields where added to the Def struct.  
```
	Type             reflect.Type
	ImplementedTypes TypeSet
```

```Type```:              is the type of the object being registered.  
```ImplementedTypes```:  is the types that this object either is or implements  

Both are needed because an upfront expensive reflect validation is done to make sure that everything is legit.    
The following [code](https://github.com/fluffy-bunny/sarulabsdi/blob/8a200c4fa3aefa0a28ddc66739aac1631f2a95aa/builder.go#L93) will panic during a bad add.  


```go 
var diServiceName = "di-transient-service"

type ISomething interface {
	GetName() string
	SetName(name string)
}

type Service struct {
	name            string
}
// SetName ...
func (s *Service) SetName(in string) {
	s.name = in
}

// SetName ...
func (s *Service) GetName() string {
	return s.name
}

func GetInterfaceReflectType(i interface{}) reflect.Type {
	return reflect.TypeOf(i).Elem()
}

func AddTransientService(builder *di.Builder) {
	log.Info().Msg("IoC: AddTransientService")

	types := di.NewTypeSet()
	inter := GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	types.Add(reflect.TypeOf(&Service{}))

	builder.Add(di.Def{
		Name:             diServiceName,
		Scope:            di.App,
		ImplementedTypes: types,
		Type:             reflect.TypeOf(&Service{}),
		Unshared: true,
		Build: func(ctn di.Container) (interface{}, error) {
			service := &Service{
			}
			return service, nil
		},
	})
}
```

### Retrieval 

```go
// please put this into a lookup map
inter := GetInterfaceReflectType((*exampleServices.ISomething)(nil))

dd := ctn.GetByType(inter)
for _, d := range dd {
  ds := d.(exampleServices.ISomething)
  ds.SetName("rabbit")
  log.Info().Msg(ds.GetName())
}

```


