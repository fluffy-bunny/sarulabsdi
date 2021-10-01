package di

import "reflect"

// AddTransientByType adds a simple transient type
func AddTransientByType(builder *Builder, rt reflect.Type) Def {
	return AddTransientWithImplementedTypesByType(builder, rt, nil)
}

// AddTransientWithImplementedTypesByType adds a type and its implemented interfaces
func AddTransientWithImplementedTypesByType(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
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

// AddScopedByType adds a simple scoped type
func AddScopedByType(builder *Builder, rt reflect.Type) Def {
	return AddScopedWithImplementedTypesByType(builder, rt, nil)
}

// AddScopedWithImplementedTypesByType adds a type and its implemented interfaces
func AddScopedWithImplementedTypesByType(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
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

// AddScopedWithImplementedTypesWithBuilderByType adds a type and its implemented interfaces
func AddScopedWithImplementedTypesWithBuilderByType(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
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

// AddSingletonByType adds a simple singleton type
func AddSingletonByType(builder *Builder, rt reflect.Type) Def {
	return AddSingletonWithImplementedTypesByType(builder, rt, nil)
}

// AddSingletonWithImplementedTypesByType adds a prebuilt obj
func AddSingletonWithImplementedTypesByType(builder *Builder, rt reflect.Type, implementedTypes ...reflect.Type) Def {
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

// AddSingletonWithImplementedTypesWithBuilderByType adds a prebuilt obj
func AddSingletonWithImplementedTypesWithBuilderByType(builder *Builder, rt reflect.Type, build func(ctn Container) (interface{}, error), implementedTypes ...reflect.Type) Def {
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
