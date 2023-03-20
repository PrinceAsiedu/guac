// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"

	"github.com/99designs/gqlgen/graphql"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _IsDependency_id(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsDependency_package(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_package(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Package, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Package)
	fc.Result = res
	return ec.marshalNPackage2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackage(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_package(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Package_id(ctx, field)
			case "type":
				return ec.fieldContext_Package_type(ctx, field)
			case "namespaces":
				return ec.fieldContext_Package_namespaces(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Package", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsDependency_dependentPackage(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_dependentPackage(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.DependentPackage, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Package)
	fc.Result = res
	return ec.marshalNPackage2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackage(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_dependentPackage(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Package_id(ctx, field)
			case "type":
				return ec.fieldContext_Package_type(ctx, field)
			case "namespaces":
				return ec.fieldContext_Package_namespaces(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Package", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsDependency_versionRange(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_versionRange(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.VersionRange, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_versionRange(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsDependency_justification(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_justification(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Justification, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_justification(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsDependency_origin(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_origin(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Origin, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_origin(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _IsDependency_collector(ctx context.Context, field graphql.CollectedField, obj *model.IsDependency) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_IsDependency_collector(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Collector, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_IsDependency_collector(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "IsDependency",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputIsDependencyInputSpec(ctx context.Context, obj interface{}) (model.IsDependencyInputSpec, error) {
	var it model.IsDependencyInputSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"versionRange", "justification", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "versionRange":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("versionRange"))
			it.VersionRange, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "justification":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("justification"))
			it.Justification, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputIsDependencySpec(ctx context.Context, obj interface{}) (model.IsDependencySpec, error) {
	var it model.IsDependencySpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "package", "dependentPackage", "versionRange", "justification", "origin", "collector"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			it.ID, err = ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "package":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			it.Package, err = ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "dependentPackage":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("dependentPackage"))
			it.DependentPackage, err = ec.unmarshalOPkgNameSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgNameSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "versionRange":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("versionRange"))
			it.VersionRange, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "justification":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("justification"))
			it.Justification, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "origin":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("origin"))
			it.Origin, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "collector":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("collector"))
			it.Collector, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputPkgNameSpec(ctx context.Context, obj interface{}) (model.PkgNameSpec, error) {
	var it model.PkgNameSpec
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"id", "type", "namespace", "name"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "id":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
			it.ID, err = ec.unmarshalOID2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "type":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("type"))
			it.Type, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "namespace":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("namespace"))
			it.Namespace, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		case "name":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("name"))
			it.Name, err = ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var isDependencyImplementors = []string{"IsDependency", "Nodes"}

func (ec *executionContext) _IsDependency(ctx context.Context, sel ast.SelectionSet, obj *model.IsDependency) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, isDependencyImplementors)
	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("IsDependency")
		case "id":

			out.Values[i] = ec._IsDependency_id(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "package":

			out.Values[i] = ec._IsDependency_package(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "dependentPackage":

			out.Values[i] = ec._IsDependency_dependentPackage(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "versionRange":

			out.Values[i] = ec._IsDependency_versionRange(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "justification":

			out.Values[i] = ec._IsDependency_justification(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "origin":

			out.Values[i] = ec._IsDependency_origin(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "collector":

			out.Values[i] = ec._IsDependency_collector(ctx, field, obj)

			if out.Values[i] == graphql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNIsDependency2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsDependency(ctx context.Context, sel ast.SelectionSet, v model.IsDependency) graphql.Marshaler {
	return ec._IsDependency(ctx, sel, &v)
}

func (ec *executionContext) marshalNIsDependency2ᚕᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsDependencyᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.IsDependency) graphql.Marshaler {
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNIsDependency2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsDependency(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalNIsDependency2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsDependency(ctx context.Context, sel ast.SelectionSet, v *model.IsDependency) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._IsDependency(ctx, sel, v)
}

func (ec *executionContext) unmarshalNIsDependencyInputSpec2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsDependencyInputSpec(ctx context.Context, v interface{}) (model.IsDependencyInputSpec, error) {
	res, err := ec.unmarshalInputIsDependencyInputSpec(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOIsDependencySpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐIsDependencySpec(ctx context.Context, v interface{}) (*model.IsDependencySpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputIsDependencySpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) unmarshalOPkgNameSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgNameSpec(ctx context.Context, v interface{}) (*model.PkgNameSpec, error) {
	if v == nil {
		return nil, nil
	}
	res, err := ec.unmarshalInputPkgNameSpec(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
