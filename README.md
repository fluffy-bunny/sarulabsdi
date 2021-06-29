# sarulabsdi  
This is a baseline drop of [v2.4.2 of github.com/sarulabsdi](https://github.com/sarulabs/di/releases/tag/v2.4.2)   


# New Features
The following are non breaking changes to sarulabs/di.  

## SafeGetByType  
```go
// SafeGetByType retrieves the last object added from the Container.
// The object has to belong to this scope or a more generic one.
// If the object does not already exist, it is created and saved in the Container.
// If the object can not be created, it returns an error.
SafeGetByType(rt reflect.Type) (interface{}, error)
```

## GetByType  
```go
// GetByType is similar to SafeGetByType but it does not return the error.
// Instead it panics.
GetByType(rt reflect.Type) interface{}
```

## SafeGetManyByType  
```go
// SafeGetManyByType retrieves an array of objects from the Container.
// The objects have to belong to this scope or a more generic one.
// If the objects do not already exist, it is created and saved in the Container.
// If the objects can not be created, it returns an error.
SafeGetManyByType(rt reflect.Type) ([]interface{}, error)
```

## GetManyByType  
```go
// GetManyByType is similar to SafeGetManyByType but it does not return the error.
// Instead it panics.
GetManyByType(rt reflect.Type) []interface{}
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

Only the ```Type``` option is needed to register and it is automatically added as an implemented type.   The ```ImplementedTypes``` option is primarily for claiming that the added type supports a given set of interfaces.  If this option is added, the original ```Type``` is automatically added to the ```ImplementedTypes``` set.  

The following [code](https://github.com/fluffy-bunny/sarulabsdi/blob/909f303f513ce84953164cc78b311a57ae959544/builder.go#L90) will return an error if the added type **DOES NOT** implemented the ```ImplementedTypes``` that were claimed. 


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

func AddTransientService(builder *di.Builder) {
	log.Info().Msg("IoC: AddTransientService")

	types := di.NewTypeSet()
	inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))
	types.Add(inter)
	
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
inter := di.GetInterfaceReflectType((*exampleServices.ISomething)(nil))

dd := ctn.GetManyByType(inter)
for _, d := range dd {
  ds := d.(exampleServices.ISomething)
  ds.SetName("rabbit")
  log.Info().Msg(ds.GetName())
}

```


