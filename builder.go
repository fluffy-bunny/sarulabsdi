package di

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// Builder can be used to create a Container.
// The Builder should be created with NewBuilder.
// Then you can add definitions with the Add method,
// and finally build the Container with the Build method.
type Builder struct {
	definitions DefMap
	scopes      ScopeList
}

// NewBuilder is the only way to create a working Builder.
// It initializes a Builder with a list of scopes.
// The scopes are ordered from the most generic to the most specific.
// If no scope is provided, the default scopes are used:
// [App, Request, SubRequest]
// It can return an error if the scopes are not valid.
func NewBuilder(scopes ...string) (*Builder, error) {
	if len(scopes) == 0 {
		scopes = []string{App, Request, SubRequest}
	}

	if err := checkScopes(scopes); err != nil {
		return nil, err
	}

	return &Builder{
		definitions: DefMap{},
		scopes:      scopes,
	}, nil
}

func checkScopes(scopes []string) error {
	if len(scopes) == 0 {
		return errors.New("at least one scope is required")
	}

	for i, scope := range scopes {
		if scope == "" {
			return errors.New("a scope can not be an empty string")
		}
		if ScopeList(scopes[i+1:]).Contains(scope) {
			return fmt.Errorf("at least two scopes are identical")
		}
	}

	return nil
}

// Scopes returns the list of available scopes.
func (b *Builder) Scopes() ScopeList {
	return ScopeList(b.scopes).Copy()
}

// Definitions returns a map with the all the objects definitions
// registered with the Add method.
// The key of the map is the name of the Definition.
func (b *Builder) Definitions() DefMap {
	return b.definitions.Copy()
}

// IsDefined returns true if there is a definition with the given name.
func (b *Builder) IsDefined(name string) bool {
	_, ok := b.definitions[name]
	return ok
}

// Add adds one or more definitions in the Builder.
// It returns an error if a definition can not be added.
// If a definition with the same name has already been added,
// it will be replaced by the new one, as if the first one never existed.
func (b *Builder) Add(defs ...Def) error {
	for _, def := range defs {
		if err := b.add(def); err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) add(def Def) error {
	if def.Type != nil {
		def.Name = GenerateUniqueServiceKeyFromType(def.Type.Elem())
	}

	if def.Name == "" {
		return errors.New("name can not be empty")
	}
	fmt.Println(def.Name)
	for rt := range def.ImplementedTypes {
		fmt.Println(rt.String())
		if rt.Kind() == reflect.Interface {
			if !def.Type.Implements(rt) {
				panic(fmt.Errorf("%v does not implement %v", def.Name, rt))
			}
		} else {
			if def.Type != rt {
				panic(fmt.Errorf("%v does not implement %v", def.Name, rt))
			}
		}
	}

	// note that an empty scope is allowed
	// it will be replaced in the Build method by the most generic scope
	if def.Scope != "" && !b.scopes.Contains(def.Scope) {
		return fmt.Errorf("scope `%s` is not allowed", def.Scope)
	}

	if def.Build == nil {
		return errors.New("Build can not be nil")
	}

	b.definitions[def.Name] = def

	return nil
}

// Set is a shortcut to add a definition for an already built object.
func (b *Builder) Set(name string, obj interface{}) error {
	return b.add(Def{
		Name: name,
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	})
}

// Build creates a Container in the most generic scope
// with all the definitions registered in the Builder.
func (b *Builder) Build() Container {
	if err := checkScopes(b.scopes); err != nil {
		return nil
	}

	defs := b.Definitions()

	rtDefMap := make(map[string]deflist)

	// efficiency map.  Build out a fast lookup for types to defs
	for name, _ := range defs {
		def := defs[name]

		if def.Scope == "" {
			def.Scope = b.scopes[0]
			defs[name] = def
		}

		for rt := range def.ImplementedTypes {
			var key string
			if rt.Kind() == reflect.Interface {
				key = GenerateReproducableTypeKey(rt)
			} else {
				key = GenerateReproducableTypeKey(rt.Elem())
			}
			fmt.Println(key)
			rtDefMap[key] = append(rtDefMap[key], &def)
		}
	}

	for k, _ := range rtDefMap {
		eList := rtDefMap[k]
		sort.Sort(deflist(eList))
	}
	return &container{
		containerCore: &containerCore{
			scopes:       b.scopes,
			scope:        b.scopes[0],
			definitions:  defs,
			parent:       nil,
			children:     map[*containerCore]struct{}{},
			objects:      map[objectKey]interface{}{},
			dependencies: newGraph(),
			typeDefMap:   rtDefMap,
		},
	}
}
