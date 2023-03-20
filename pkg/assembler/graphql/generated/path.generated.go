// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
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

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputPackageSourceArtifactBuilderOsvCveOrGhsaFilter(ctx context.Context, obj interface{}) (model.PackageSourceArtifactBuilderOsvCveOrGhsaFilter, error) {
	var it model.PackageSourceArtifactBuilderOsvCveOrGhsaFilter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"package", "source", "artifact", "builder", "osv", "cve", "ghsa"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "package":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("package"))
			it.Package, err = ec.unmarshalOPkgSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPkgSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "source":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("source"))
			it.Source, err = ec.unmarshalOSourceSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐSourceSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "artifact":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("artifact"))
			it.Artifact, err = ec.unmarshalOArtifactSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐArtifactSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "builder":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("builder"))
			it.Builder, err = ec.unmarshalOBuilderSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐBuilderSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "osv":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("osv"))
			it.Osv, err = ec.unmarshalOOSVSpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐOSVSpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "cve":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("cve"))
			it.Cve, err = ec.unmarshalOCVESpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐCVESpec(ctx, v)
			if err != nil {
				return it, err
			}
		case "ghsa":
			var err error

			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("ghsa"))
			it.Ghsa, err = ec.unmarshalOGHSASpec2ᚖgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐGHSASpec(ctx, v)
			if err != nil {
				return it, err
			}
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

func (ec *executionContext) _Nodes(ctx context.Context, sel ast.SelectionSet, obj model.Nodes) graphql.Marshaler {
	switch obj := (obj).(type) {
	case nil:
		return graphql.Null
	case model.Package:
		return ec._Package(ctx, sel, &obj)
	case *model.Package:
		if obj == nil {
			return graphql.Null
		}
		return ec._Package(ctx, sel, obj)
	case model.Source:
		return ec._Source(ctx, sel, &obj)
	case *model.Source:
		if obj == nil {
			return graphql.Null
		}
		return ec._Source(ctx, sel, obj)
	case model.Artifact:
		return ec._Artifact(ctx, sel, &obj)
	case *model.Artifact:
		if obj == nil {
			return graphql.Null
		}
		return ec._Artifact(ctx, sel, obj)
	case model.Builder:
		return ec._Builder(ctx, sel, &obj)
	case *model.Builder:
		if obj == nil {
			return graphql.Null
		}
		return ec._Builder(ctx, sel, obj)
	case model.Osv:
		return ec._OSV(ctx, sel, &obj)
	case *model.Osv:
		if obj == nil {
			return graphql.Null
		}
		return ec._OSV(ctx, sel, obj)
	case model.Cve:
		return ec._CVE(ctx, sel, &obj)
	case *model.Cve:
		if obj == nil {
			return graphql.Null
		}
		return ec._CVE(ctx, sel, obj)
	case model.Ghsa:
		return ec._GHSA(ctx, sel, &obj)
	case *model.Ghsa:
		if obj == nil {
			return graphql.Null
		}
		return ec._GHSA(ctx, sel, obj)
	case model.IsOccurrence:
		return ec._IsOccurrence(ctx, sel, &obj)
	case *model.IsOccurrence:
		if obj == nil {
			return graphql.Null
		}
		return ec._IsOccurrence(ctx, sel, obj)
	case model.IsDependency:
		return ec._IsDependency(ctx, sel, &obj)
	case *model.IsDependency:
		if obj == nil {
			return graphql.Null
		}
		return ec._IsDependency(ctx, sel, obj)
	case model.IsVulnerability:
		return ec._IsVulnerability(ctx, sel, &obj)
	case *model.IsVulnerability:
		if obj == nil {
			return graphql.Null
		}
		return ec._IsVulnerability(ctx, sel, obj)
	case model.CertifyVEXStatement:
		return ec._CertifyVEXStatement(ctx, sel, &obj)
	case *model.CertifyVEXStatement:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyVEXStatement(ctx, sel, obj)
	case model.HashEqual:
		return ec._HashEqual(ctx, sel, &obj)
	case *model.HashEqual:
		if obj == nil {
			return graphql.Null
		}
		return ec._HashEqual(ctx, sel, obj)
	case model.CertifyBad:
		return ec._CertifyBad(ctx, sel, &obj)
	case *model.CertifyBad:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyBad(ctx, sel, obj)
	case model.CertifyPkg:
		return ec._CertifyPkg(ctx, sel, &obj)
	case *model.CertifyPkg:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyPkg(ctx, sel, obj)
	case model.CertifyScorecard:
		return ec._CertifyScorecard(ctx, sel, &obj)
	case *model.CertifyScorecard:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyScorecard(ctx, sel, obj)
	case model.CertifyVuln:
		return ec._CertifyVuln(ctx, sel, &obj)
	case *model.CertifyVuln:
		if obj == nil {
			return graphql.Null
		}
		return ec._CertifyVuln(ctx, sel, obj)
	case model.HasSourceAt:
		return ec._HasSourceAt(ctx, sel, &obj)
	case *model.HasSourceAt:
		if obj == nil {
			return graphql.Null
		}
		return ec._HasSourceAt(ctx, sel, obj)
	case model.HasSbom:
		return ec._HasSBOM(ctx, sel, &obj)
	case *model.HasSbom:
		if obj == nil {
			return graphql.Null
		}
		return ec._HasSBOM(ctx, sel, obj)
	case model.HasSlsa:
		return ec._HasSLSA(ctx, sel, &obj)
	case *model.HasSlsa:
		if obj == nil {
			return graphql.Null
		}
		return ec._HasSLSA(ctx, sel, obj)
	default:
		panic(fmt.Errorf("unexpected type %T", obj))
	}
}

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNNodes2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐNodes(ctx context.Context, sel ast.SelectionSet, v model.Nodes) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Nodes(ctx, sel, v)
}

func (ec *executionContext) marshalNNodes2ᚕgithubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐNodesᚄ(ctx context.Context, sel ast.SelectionSet, v []model.Nodes) graphql.Marshaler {
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
			ret[i] = ec.marshalNNodes2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐNodes(ctx, sel, v[i])
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

func (ec *executionContext) unmarshalNPackageSourceArtifactBuilderOsvCveOrGhsaFilter2githubᚗcomᚋguacsecᚋguacᚋpkgᚋassemblerᚋgraphqlᚋmodelᚐPackageSourceArtifactBuilderOsvCveOrGhsaFilter(ctx context.Context, v interface{}) (model.PackageSourceArtifactBuilderOsvCveOrGhsaFilter, error) {
	res, err := ec.unmarshalInputPackageSourceArtifactBuilderOsvCveOrGhsaFilter(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
