package di

import "reflect"

// AddTransient adds a simple transient type
func AddTransient(builder *Builder, rt reflect.Type) Def {
	return AddTransientWithImplementedTypes(builder, rt, nil)
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
		SafeInject:       true, // don't panic
		Unshared:         true, // Transient
	}
	builder.Add(def)
	return def
}

// AddScoped adds a simple scoped type
func AddScoped(builder *Builder, rt reflect.Type) Def {
	return AddScopedWithImplementedTypes(builder, rt, nil)
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
		SafeInject:       true,  // don't panic
		Unshared:         false, // singleton within scope
	}
	builder.Add(def)
	return def
}

// AddScopedWithImplementedTypesWithBuilder adds a type and its implemented interfaces
func AddScopedWithImplementedTypesWithBuilder(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            Request, // Scoped
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		SafeInject:       true,  // don't panic
		Unshared:         false, // singleton within scope
		Build:            build,
	}
	builder.Add(def)
	return def
}

// AddSingleton adds a simple singleton type
func AddSingleton(builder *Builder, rt reflect.Type) Def {
	return AddSingletonWithImplementedTypes(builder, rt, nil)
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
		SafeInject:       true,  // don't panic
		Unshared:         false, // Singleton
	}
	builder.Add(def)
	return def
}

// AddSingletonWithImplementedTypesWithBuilder adds a prebuilt obj
func AddSingletonWithImplementedTypesWithBuilder(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
	implementedTypes2 := NewTypeSet()
	for _, rt := range implementedTypes {
		implementedTypes2.Add(rt)
	}
	def := Def{
		Scope:            App, // Singleton
		Type:             rt,
		ImplementedTypes: implementedTypes2,
		SafeInject:       true,  // don't panic
		Unshared:         false, // Singleton
		Build:            build,
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
		SafeInject:       true,  // don't panic
		Unshared:         false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
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
		SafeInject: true,  // don't panic
		Unshared:   false, // Singleton
		Build: func(ctn Container) (interface{}, error) {
			return obj, nil
		},
	}
	builder.Add(def)
	return def
}
