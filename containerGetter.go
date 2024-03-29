package di

import (
	"fmt"
	"reflect"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// buildingChan is used internally as the value of an object while it is being built.
type buildingChan chan struct{}

// containerGetter contains all the functions that are useful
// to retrieve an object from a container.
type containerGetter struct{}

func (g *containerGetter) GetByType(ctn *container, rt reflect.Type) interface{} {
	obj, err := ctn.SafeGetByType(rt)
	if err != nil {
		panic(err)
	}

	return obj
}
func (g *containerGetter) GetManyByType(ctn *container, rt reflect.Type) []interface{} {
	obj, err := ctn.SafeGetManyByType(rt)
	if err != nil {
		panic(err)
	}

	return obj
}

func (g *containerGetter) Get(ctn *container, name string) interface{} {
	obj, err := ctn.SafeGet(name)
	if err != nil {
		panic(err)
	}

	return obj
}

func (g *containerGetter) Fill(ctn *container, name string, dst interface{}) error {
	obj, err := ctn.SafeGet(name)
	if err != nil {
		return err
	}

	return fill(obj, dst)
}

func (g *containerGetter) GetDefinitionByType(ctn *container, rt reflect.Type) *Def {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	defs, ok := ctn.typeDefMap[rt]
	if ok {
		if ctn.scope == "app" {
			for _, def := range defs {
				if def.Scope == "request" {
					// special case because we may have the same type registered in both app and request scopes
					// ignore anything but app scopes
					continue
				}
				return def
			}
		} else if ctn.scope == "request" || ctn.scope == "subrequest" {
			// only go after the first request scope first
			var nonRequestdef []*Def
			for _, def := range defs {
				if def.Scope != "request" {
					nonRequestdef = append(nonRequestdef, def)
					continue
				}
				return def
			}
			return nonRequestdef[0]
		}
	}
	return nil

}

var rtIContainer = reflect.TypeOf((*Container)(nil)).Elem()

func (g *containerGetter) SafeGetByType(ctn *container, rtIn reflect.Type) (interface{}, error) {
	rt := rtIn
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	containerRequested := rtIContainer == rt
	if containerRequested {
		return ctn, nil
	}
	defs, ok := ctn.typeDefMap[rt]
	if ok {
		if ctn.scope == "app" {
			for _, def := range defs {
				if def.Scope == "request" {
					// special case because we may have the same type registered in both app and request scopes
					// ignore anything but app scopes
					continue
				}
				obj, err := g.SafeGet(ctn, def.Name)
				if err != nil {
					return nil, err
				}
				return obj, err
			}
		} else if ctn.scope == "request" || ctn.scope == "subrequest" {
			// only go after the first request scope first
			var nonRequestdef []*Def
			for _, def := range defs {
				if def.Scope != "request" {
					nonRequestdef = append(nonRequestdef, def)
					continue
				}
				obj, err := g.SafeGet(ctn, def.Name)
				if err != nil {
					return nil, err
				}
				return obj, err
			}
			// there is no request scope, so go after the first non request scope
			obj, err := g.SafeGet(ctn, nonRequestdef[0].Name)
			if err != nil {
				return nil, err
			}
			return obj, err
		}
	}
	return nil, fmt.Errorf("could not get type of `%s` because the definition does not exist", getTypeFullPath(rt))
}

func (g *containerGetter) GetDefinitionsByType(ctn *container, rt reflect.Type) []*Def {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	var result []*Def
	defs, ok := ctn.typeDefMap[rt]
	if !ok {
		return result
	}
	return defs
}
func (g *containerGetter) SafeGetManyByType(ctn *container, rt reflect.Type) ([]interface{}, error) {
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	var result []interface{}
	defs, ok := ctn.typeDefMap[rt]
	if !ok {
		return result, status.Errorf(codes.NotFound, "could not get types of `%s` because the definition does not exist", getTypeFullPath(rt))
	}
	for _, def := range defs {
		obj, err := g.SafeGet(ctn, def.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)

	}
	return result, nil
}
func (g *containerGetter) SafeGet(ctn *container, name string) (interface{}, error) {
	def, ok := ctn.definitions[name]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "could not get `%s` because the definition does not exist", name)
	}

	if ctn.builtList.HasDef(name) {
		return nil, status.Errorf(codes.NotFound, "could not get `%s` because there is a cycle in the object definitions (%v)", def.Name, ctn.builtList.OrderedList())
	}

	if ctn.scope != def.Scope {
		return g.getInParent(ctn, def)
	}

	return g.getInThisContainer(ctn, def)
}

func (g *containerGetter) getInParent(ctn *container, def Def) (interface{}, error) {
	p := ctn.containerLineage.parent(ctn)

	if p.containerCore == nil {
		return nil, status.Errorf(codes.InvalidArgument,
			"could not get `%s` because it requires `%s` scope which does not match this container scope or any of its parents scope", def.Name, def.Scope,
		)
	}

	if p.scope != def.Scope {
		return g.getInParent(p, def)
	}

	return g.getInThisContainer(p, def)
}

func (g *containerGetter) getInThisContainer(ctn *container, def Def) (interface{}, error) {
	ctn.m.Lock()

	if ctn.closed {
		ctn.m.Unlock()
		return nil, status.Errorf(codes.NotFound, "could not get `%s` because the container has been deleted", def.Name)
	}

	objKey := objectKey{defName: def.Name}
	if def.Unshared {
		ctn.lastUniqueID += 1
		objKey.uniqueID = ctn.lastUniqueID
	}

	g.addDependencyToGraph(ctn, objKey)

	obj, ok := ctn.objects[objKey]
	if !ok {
		// the object need to be created
		c := make(buildingChan)
		ctn.objects[objKey] = c
		ctn.m.Unlock()
		return g.buildInThisContainer(ctn, def, objKey, c)
	}

	ctn.m.Unlock()

	c, isBuilding := obj.(buildingChan)

	if !isBuilding {
		// the object is ready to be used
		return obj, nil
	}

	// the object is being built by another goroutine
	<-c

	return g.getInThisContainer(ctn, def)
}

func (g *containerGetter) addDependencyToGraph(ctn *container, objKey objectKey) {
	if last, ok := ctn.builtList.LastElement(); ok {
		ctn.dependencies.AddEdge(last, objKey)
		return
	}

	ctn.dependencies.AddVertex(objKey)
}

func (g *containerGetter) buildInThisContainer(ctn *container, def Def, objKey objectKey, c buildingChan) (interface{}, error) {
	obj, err := g.build(ctn, def, objKey)

	ctn.m.Lock()

	if err != nil {
		// The object could not be created. Remove the channel from the object map
		// and close it to allow the object to be created again.
		delete(ctn.objects, objKey)
		ctn.m.Unlock()
		close(c)
		return nil, err
	}

	if ctn.closed {
		// The container has been deleted while the object was being built.
		// The newly created object needs to be closed, and it will not be returned.
		ctn.m.Unlock()
		close(c)
		err = ctn.containerSlayer.closeObject(obj, def)
		return nil, status.Errorf(codes.NotFound,
			"could not get `%s` because the container has been deleted, the object has been created and closed%s",
			def.Name, g.formatCloseErr(err),
		)
	}

	ctn.objects[objKey] = obj
	ctn.m.Unlock()
	close(c)

	return obj, nil
}

func (g *containerGetter) formatCloseErr(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf(" (with an error: %+v)", err)
}

func (g *containerGetter) build(ctn *container, def Def, objKey objectKey) (obj interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "could not build `%s` because the build function panicked: %+v", def.Name, r)
		}
	}()

	obj, err = def.Build(&container{
		containerCore: ctn.containerCore,
		builtList:     ctn.builtList.Add(objKey),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not build `%s`: %+v", def.Name, err)
	}

	return obj, nil
}
