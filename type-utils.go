package di

import "reflect"

// AddTransient adds a simple transient type
func AddTransient(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	return AddTransientWithImplementedTypes(builder, rt, implementedTypes...)
}

// AddTransientWithMetadata adds a simple transient type
func AddTransientWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	return AddTransientWithImplementedTypesWithMetadata(builder, rt, metaData, implementedTypes...)
}

// AddTransientWithImplementedTypes adds a type and its implemented interfaces
func AddTransientWithImplementedTypes(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
	}
	builder.Add(def)
	return def
}

// AddTransientWithImplementedTypes adds a type and its implemented interfaces
func AddTransientWithImplementedTypesWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddTransientWithImplementedTypesByFunc adds a type and its implemented interfaces
func AddTransientWithImplementedTypesByFunc(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
		Build:            build,
	}
	builder.Add(def)
	return def
}

// AddTransientWithImplementedTypesByFuncWithMetadata adds a type and its implemented interfaces
func AddTransientWithImplementedTypesByFuncWithMetadata(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         true, // Transient
		Build:            build,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddScoped adds a simple scoped type
func AddScoped(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	return AddScopedWithImplementedTypes(builder, rt, implementedTypes...)
}

// AddScopedWithMetadata adds a simple scoped type
func AddScopedWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	return AddScopedWithImplementedTypesWithMetadata(builder, rt, metaData, implementedTypes...)
}

// AddScopedWithImplementedTypes adds a type and its implemented interfaces
func AddScopedWithImplementedTypes(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
	}
	builder.Add(def)
	return def
}

// AddScopedWithImplementedTypesWithMetadata adds a type and its implemented interfaces
func AddScopedWithImplementedTypesWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddScopedWithImplementedTypesByFunc adds a type and its implemented interfaces
func AddScopedWithImplementedTypesByFunc(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
		Build:            build,
	}
	builder.Add(def)
	return def
}

// AddScopedWithImplementedTypesByFuncWithMetadata adds a type and its implemented interfaces
func AddScopedWithImplementedTypesByFuncWithMetadata(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // singleton within scope
		Build:            build,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddFunc adds a simple singleton type
func AddFunc(builder *Builder, fnc interface{}) Def {
	rt := reflect.TypeOf(fnc)
	implementedTypes2 := NewTypeSet()
	implementedTypes2.Add(rt)
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Func:             fnc,
	}
	builder.Add(def)
	return def
}

// AddFunc adds a simple singleton type
func AddFuncWithMetadata(builder *Builder, fnc interface{}, metaData map[string]interface{}) Def {
	rt := reflect.TypeOf(fnc)
	implementedTypes2 := NewTypeSet()
	implementedTypes2.Add(rt)
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Func:             fnc,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddSingleton adds a simple singleton type
func AddSingleton(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	return AddSingletonWithImplementedTypes(builder, rt, implementedTypes...)
}

// AddSingletonWithMetadata adds a simple singleton type
func AddSingletonWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	return AddSingletonWithImplementedTypesWithMetadata(builder, rt, metaData, implementedTypes...)
}

// AddSingletonWithImplementedTypes adds a prebuilt obj
func AddSingletonWithImplementedTypes(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
	}
	builder.Add(def)
	return def
}

// AddSingletonWithImplementedTypesWithMetadata adds a prebuilt obj
func AddSingletonWithImplementedTypesWithMetadata(builder *Builder, rt reflect.Type, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddSingletonWithImplementedTypesByFunc adds a prebuilt obj
func AddSingletonWithImplementedTypesByFunc(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build:            build,
	}
	builder.Add(def)
	return def
}

// AddSingletonWithImplementedTypesByFuncWithMetadata adds a prebuilt obj
func AddSingletonWithImplementedTypesByFuncWithMetadata(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build:            build,
		MetaData:         metaData,
	}
	builder.Add(def)
	return def
}

// AddSingletonWithImplementedTypesByObj adds a prebuilt obj
func AddSingletonWithImplementedTypesByObj(builder *Builder, obj interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             reflect.TypeOf(obj),
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	}
	builder.Add(def)
	return def
}

// AddSingletonWithImplementedTypesByObjWithMetadata adds a prebuilt obj
func AddSingletonWithImplementedTypesByObjWithMetadata(builder *Builder, obj interface{}, metaData map[string]interface{}, implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             reflect.TypeOf(obj),
		ImplementedTypes: implementedTypes2,
		Unshared:         false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
		MetaData: metaData,
	}
	builder.Add(def)
	return def
}

// AddSingletonTypeByObj adds a prebuilt object by its type
func AddSingletonTypeByObj(builder *Builder, obj interface{}) Def {
	def := Def{
		Scope: App, // Singleton
		Type:  reflect.TypeOf(obj),
		// SafeInject MUST be true because the following CAN be nil
		// Dialer kafkaContracts.IKafkaDialer `inject:""`
		Unshared: false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	}
	builder.Add(def)
	return def
}

// AddSingletonTypeByObjWithMetadata adds a prebuilt object by its type
func AddSingletonTypeByObjWithMetadata(builder *Builder, obj interface{}, metaData map[string]interface{}) Def {
	def := Def{
		Scope: App, // Singleton
		Type:  reflect.TypeOf(obj),
		// SafeInject MUST be true because the following CAN be nil
		// Dialer kafkaContracts.IKafkaDialer `inject:""`
		Unshared: false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
		MetaData: metaData,
	}
	builder.Add(def)
	return def
}
