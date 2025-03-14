//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Khan/genqlient/graphql"
	"github.com/guacsec/guac/internal/testing/ptrfrom"
	model "github.com/guacsec/guac/pkg/assembler/clients/generated"
	"github.com/guacsec/guac/pkg/assembler/helpers"
	"github.com/guacsec/guac/pkg/cli"
	"github.com/guacsec/guac/pkg/logging"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	hashEqualStr        string = "hashEqual"
	scorecardStr        string = "scorecard"
	occurrenceStr       string = "occurrence"
	hasSrcAtStr         string = "hasSrcAt"
	hasSBOMStr          string = "hasSBOM"
	hasSLSAStr          string = "hasSLSA"
	certifyVulnStr      string = "certifyVuln"
	certifyLegalStr     string = "certifyLegal"
	vexLinkStr          string = "vexLink"
	badLinkStr          string = "badLink"
	goodLinkStr         string = "goodLink"
	pkgEqualStr         string = "pkgEqual"
	packageSubjectType  string = "package"
	sourceSubjectType   string = "source"
	artifactSubjectType string = "artifact"
)

type queryKnownOptions struct {
	// gql endpoint
	graphqlEndpoint string
	headerFile      string
	// package, source or artifact
	subjectType string
	// purl / source (<vcs_tool>+<transport>) / artifact (algorithm:digest)
	subject string
}

type neighbors struct {
	hashEquals    []*model.NeighborsNeighborsHashEqual
	scorecards    []*model.NeighborsNeighborsCertifyScorecard
	occurrences   []*model.NeighborsNeighborsIsOccurrence
	hasSrcAt      []*model.NeighborsNeighborsHasSourceAt
	hasSBOMs      []*model.NeighborsNeighborsHasSBOM
	hasSLSAs      []*model.NeighborsNeighborsHasSLSA
	certifyVulns  []*model.NeighborsNeighborsCertifyVuln
	certifyLegals []*model.NeighborsNeighborsCertifyLegal
	vexLinks      []*model.NeighborsNeighborsCertifyVEXStatement
	badLinks      []*model.NeighborsNeighborsCertifyBad
	goodLinks     []*model.NeighborsNeighborsCertifyGood
	pkgEquals     []*model.NeighborsNeighborsPkgEqual
}

var (
	colTitleNodeType = "Node Type"
	colTitleNodeID   = "Node ID #"
	colTitleNodeInfo = "Additional Information"
	rowHeader        = table.Row{colTitleNodeType, colTitleNodeID, colTitleNodeInfo}
)

var queryKnownCmd = &cobra.Command{
	Use:   "known [flags] <type> <subject>",
	Short: "Query for all the available information on a package, source, or artifact.",
	Long: `Query for all the available information on a package, source, or artifact.
  <type> must be either "package", "source", or "artifact".
  <subject> is in the form of "<purl>" for package, "<vcs_tool>+<transport>" for source, or "<algorithm>:<digest>" for artiact.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := logging.WithLogger(context.Background())
		logger := logging.FromContext(ctx)

		opts, err := validateQueryKnownFlags(
			viper.GetString("gql-addr"),
			viper.GetString("header-file"),
			args,
		)
		if err != nil {
			fmt.Printf("unable to validate flags: %v\n", err)
			_ = cmd.Help()
			os.Exit(1)
		}

		httpClient := http.Client{Transport: cli.HTTPHeaderTransport(ctx, opts.headerFile, http.DefaultTransport)}
		gqlclient := graphql.NewClient(opts.graphqlEndpoint, &httpClient)

		t := table.NewWriter()
		tTemp := table.Table{}
		tTemp.Render()
		t.AppendHeader(rowHeader)

		var path []string
		switch opts.subjectType {
		case packageSubjectType:
			pkgInput, err := helpers.PurlToPkg(opts.subject)
			if err != nil {
				logger.Fatalf("failed to parse PURL: %v", err)
			}

			pkgQualifierFilter := []model.PackageQualifierSpec{}
			for _, qualifier := range pkgInput.Qualifiers {
				// to prevent https://github.com/golang/go/discussions/56010
				qualifier := qualifier
				pkgQualifierFilter = append(pkgQualifierFilter, model.PackageQualifierSpec{
					Key:   qualifier.Key,
					Value: &qualifier.Value,
				})
			}

			pkgFilter := &model.PkgSpec{
				Type:       &pkgInput.Type,
				Namespace:  pkgInput.Namespace,
				Name:       &pkgInput.Name,
				Version:    pkgInput.Version,
				Subpath:    pkgInput.Subpath,
				Qualifiers: pkgQualifierFilter,
			}
			pkgResponse, err := model.Packages(ctx, gqlclient, *pkgFilter)
			if err != nil {
				logger.Fatalf("error querying for package: %v", err)
			}
			if len(pkgResponse.Packages) != 1 {
				logger.Fatalf("failed to located package based on purl")
			}

			pkgNameNeighbors, neighborsPath, err := queryKnownNeighbors(ctx, gqlclient, pkgResponse.Packages[0].Namespaces[0].Names[0].Id)
			if err != nil {
				logger.Fatalf("error querying for package name neighbors: %v", err)
			}
			t.SetTitle("Package Name Nodes")
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgNameNeighbors, hasSrcAtStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgNameNeighbors, badLinkStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgNameNeighbors, goodLinkStr, packageSubjectType))
			t.AppendSeparator()

			path = append([]string{pkgResponse.Packages[0].Namespaces[0].Names[0].Id,
				pkgResponse.Packages[0].Namespaces[0].Id,
				pkgResponse.Packages[0].Id}, neighborsPath...)

			fmt.Println(t.Render())
			fmt.Printf("Visualizer url: http://localhost:3000/?path=%v\n", strings.Join(removeDuplicateValuesFromPath(path), `,`))

			pkgVersionNeighbors, neighborsPath, err := queryKnownNeighbors(ctx, gqlclient, pkgResponse.Packages[0].Namespaces[0].Names[0].Versions[0].Id)
			if err != nil {
				logger.Fatalf("error querying for package version neighbors: %v", err)
			}

			// instantiate new table for package version nodes
			t := table.NewWriter()
			tTemp := table.Table{}
			tTemp.Render()
			t.AppendHeader(rowHeader)

			t.SetTitle("Package Version Nodes")
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, hasSrcAtStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, occurrenceStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, certifyVulnStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, certifyLegalStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, hasSBOMStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, hasSLSAStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, vexLinkStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, pkgEqualStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, badLinkStr, packageSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, pkgVersionNeighbors, goodLinkStr, packageSubjectType))
			path = append([]string{pkgResponse.Packages[0].Namespaces[0].Names[0].Versions[0].Id,
				pkgResponse.Packages[0].Namespaces[0].Names[0].Id, pkgResponse.Packages[0].Namespaces[0].Id,
				pkgResponse.Packages[0].Id}, neighborsPath...)

			fmt.Println(t.Render())
			fmt.Printf("Visualizer url: http://localhost:3000/?path=%v\n", strings.Join(removeDuplicateValuesFromPath(path), `,`))

		case sourceSubjectType:
			srcInput, err := helpers.VcsToSrc(opts.subject)
			if err != nil {
				logger.Fatalf("failed to parse source: %v", err)
			}

			srcFilter := &model.SourceSpec{
				Type:      &srcInput.Type,
				Namespace: &srcInput.Namespace,
				Name:      &srcInput.Name,
				Tag:       srcInput.Tag,
				Commit:    srcInput.Commit,
			}
			srcResponse, err := model.Sources(ctx, gqlclient, *srcFilter)
			if err != nil {
				logger.Fatalf("error querying for sources: %v", err)
			}
			if len(srcResponse.Sources) != 1 {
				logger.Fatalf("failed to located sources based on vcs")
			}
			sourceNeighbors, neighborsPath, err := queryKnownNeighbors(ctx, gqlclient, srcResponse.Sources[0].Namespaces[0].Names[0].Id)
			if err != nil {
				logger.Fatalf("error querying for source neighbors: %v", err)
			}
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, sourceNeighbors, hasSrcAtStr, sourceSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, sourceNeighbors, occurrenceStr, sourceSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, sourceNeighbors, scorecardStr, sourceSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, sourceNeighbors, badLinkStr, sourceSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, sourceNeighbors, goodLinkStr, sourceSubjectType))
			path = append([]string{srcResponse.Sources[0].Namespaces[0].Names[0].Id,
				srcResponse.Sources[0].Namespaces[0].Id, srcResponse.Sources[0].Id}, neighborsPath...)

			fmt.Println(t.Render())
			fmt.Printf("Visualizer url: http://localhost:3000/?path=%v\n", strings.Join(removeDuplicateValuesFromPath(path), `,`))
		case artifactSubjectType:
			split := strings.Split(opts.subject, ":")
			if len(split) != 2 {
				logger.Fatalf("failed to parse artifact. Needs to be in algorithm:digest form")
			}
			artifactFilter := &model.ArtifactSpec{
				Algorithm: ptrfrom.String(strings.ToLower(string(split[0]))),
				Digest:    ptrfrom.String(strings.ToLower(string(split[1]))),
			}

			artifactResponse, err := model.Artifacts(ctx, gqlclient, *artifactFilter)
			if err != nil {
				logger.Fatalf("error querying for artifacts: %v", err)
			}
			if len(artifactResponse.Artifacts) != 1 {
				logger.Fatalf("failed to located artifacts based on (algorithm:digest)")
			}
			artifactNeighbors, neighborsPath, err := queryKnownNeighbors(ctx, gqlclient, artifactResponse.Artifacts[0].Id)
			if err != nil {
				logger.Fatalf("error querying for artifact neighbors: %v", err)
			}
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, hashEqualStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, occurrenceStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, hasSBOMStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, hasSLSAStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, vexLinkStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, badLinkStr, artifactSubjectType))
			t.AppendSeparator()
			t.AppendRows(getOutputBasedOnNode(ctx, gqlclient, artifactNeighbors, goodLinkStr, artifactSubjectType))
			path = append([]string{artifactResponse.Artifacts[0].Id}, neighborsPath...)

			fmt.Println(t.Render())
			fmt.Printf("Visualizer url: http://localhost:3000/?path=%v\n", strings.Join(removeDuplicateValuesFromPath(path), `,`))
		default:
			logger.Fatalf("expected type to be either a package, source or artifact")
		}
	},
}

func queryKnownNeighbors(ctx context.Context, gqlclient graphql.Client, subjectQueryID string) (*neighbors, []string, error) {
	collectedNeighbors := &neighbors{}
	var path []string
	neighborResponse, err := model.Neighbors(ctx, gqlclient, subjectQueryID, []model.Edge{})
	if err != nil {
		return nil, nil, fmt.Errorf("error querying neighbors: %v", err)
	}
	for _, neighbor := range neighborResponse.Neighbors {
		switch v := neighbor.(type) {
		case *model.NeighborsNeighborsCertifyVuln:
			collectedNeighbors.certifyVulns = append(collectedNeighbors.certifyVulns, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsCertifyBad:
			collectedNeighbors.badLinks = append(collectedNeighbors.badLinks, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsCertifyGood:
			collectedNeighbors.goodLinks = append(collectedNeighbors.goodLinks, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsCertifyScorecard:
			collectedNeighbors.scorecards = append(collectedNeighbors.scorecards, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsCertifyVEXStatement:
			collectedNeighbors.vexLinks = append(collectedNeighbors.vexLinks, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsHasSBOM:
			collectedNeighbors.hasSBOMs = append(collectedNeighbors.hasSBOMs, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsHasSLSA:
			collectedNeighbors.hasSLSAs = append(collectedNeighbors.hasSLSAs, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsHasSourceAt:
			collectedNeighbors.hasSrcAt = append(collectedNeighbors.hasSrcAt, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsHashEqual:
			collectedNeighbors.hashEquals = append(collectedNeighbors.hashEquals, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsIsOccurrence:
			collectedNeighbors.occurrences = append(collectedNeighbors.occurrences, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsPkgEqual:
			collectedNeighbors.pkgEquals = append(collectedNeighbors.pkgEquals, v)
			path = append(path, v.Id)
		case *model.NeighborsNeighborsCertifyLegal:
			collectedNeighbors.certifyLegals = append(collectedNeighbors.certifyLegals, v)
			path = append(path, v.Id)
		default:
			continue
		}
	}
	return collectedNeighbors, path, nil
}

func getOutputBasedOnNode(ctx context.Context, gqlclient graphql.Client, collectedNeighbors *neighbors, nodeType string, subjectType string) []table.Row {
	logger := logging.FromContext(ctx)
	var tableRows []table.Row
	switch nodeType {
	case certifyVulnStr:
		for _, certVuln := range collectedNeighbors.certifyVulns {
			if certVuln.Vulnerability.Type != noVulnType {
				for _, vuln := range certVuln.Vulnerability.VulnerabilityIDs {
					tableRows = append(tableRows, table.Row{certifyVulnStr, certVuln.Id, "vulnerability ID: " + vuln.VulnerabilityID})
				}
			} else {
				tableRows = append(tableRows, table.Row{certifyVulnStr, certVuln.Id, "vulnerability ID: " + noVulnType})
			}
		}
	case badLinkStr:
		for _, bad := range collectedNeighbors.badLinks {
			tableRows = append(tableRows, table.Row{badLinkStr, bad.Id, "justification: " + bad.Justification})
		}
	case goodLinkStr:
		for _, good := range collectedNeighbors.goodLinks {
			tableRows = append(tableRows, table.Row{goodLinkStr, good.Id, "justification: " + good.Justification})
		}
	case scorecardStr:
		for _, score := range collectedNeighbors.scorecards {
			tableRows = append(tableRows, table.Row{scorecardStr, score.Id, "Overall Score: " + fmt.Sprintf("%f", score.Scorecard.AggregateScore)})
		}
	case vexLinkStr:
		for _, vex := range collectedNeighbors.vexLinks {
			tableRows = append(tableRows, table.Row{vexLinkStr, vex.Id, "Vex Status: " + vex.Status})
		}
	case hasSBOMStr:
		if len(collectedNeighbors.hasSBOMs) > 0 {
			for _, sbom := range collectedNeighbors.hasSBOMs {
				tableRows = append(tableRows, table.Row{hasSBOMStr, sbom.Id, "SBOM Download Location: " + sbom.DownloadLocation})
			}
		} else {
			// if there is an isOccurrence, check to see if there are sbom associated with it
			for _, occurrence := range collectedNeighbors.occurrences {
				neighborResponseHasSBOM, err := getAssociatedArtifact(ctx, gqlclient, occurrence, model.EdgeArtifactHasSbom)
				if err != nil {
					logger.Debugf("error querying neighbors: %v", err)
				} else {
					for _, neighborHasSBOM := range neighborResponseHasSBOM.Neighbors {
						if hasSBOM, ok := neighborHasSBOM.(*model.NeighborsNeighborsHasSBOM); ok {
							tableRows = append(tableRows, table.Row{hasSBOMStr, hasSBOM.Id, "SBOM Download Location: " + hasSBOM.DownloadLocation})
						}
					}
				}
			}
		}
	case hasSLSAStr:
		if len(collectedNeighbors.hasSLSAs) > 0 {
			for _, slsa := range collectedNeighbors.hasSLSAs {
				tableRows = append(tableRows, table.Row{hasSLSAStr, slsa.Id, "SLSA Attestation Location: " + slsa.Slsa.Origin})
			}
		} else {
			// if there is an isOccurrence, check to see if there are slsa attestation associated with it
			for _, occurrence := range collectedNeighbors.occurrences {
				neighborResponseHasSLSA, err := getAssociatedArtifact(ctx, gqlclient, occurrence, model.EdgeArtifactHasSlsa)
				if err != nil {
					logger.Debugf("error querying neighbors: %v", err)
				} else {
					for _, neighborHasSLSA := range neighborResponseHasSLSA.Neighbors {
						if hasSLSA, ok := neighborHasSLSA.(*model.NeighborsNeighborsHasSLSA); ok {
							tableRows = append(tableRows, table.Row{hasSLSAStr, hasSLSA.Id, "SLSA Attestation Location: " + hasSLSA.Slsa.Origin})
						}
					}
				}
			}
		}
	case hasSrcAtStr:
		for _, src := range collectedNeighbors.hasSrcAt {
			if subjectType == packageSubjectType {
				namespace := ""
				if !strings.HasPrefix(src.Source.Namespaces[0].Namespace, "https://") {
					namespace = "https://" + src.Source.Namespaces[0].Namespace
				} else {
					namespace = src.Source.Namespaces[0].Namespace
				}
				tableRows = append(tableRows, table.Row{hasSrcAtStr, src.Id, "Source: " + src.Source.Type + "+" + namespace + "/" +
					src.Source.Namespaces[0].Names[0].Name})
			} else {
				purl := helpers.PkgToPurl(src.Package.Type, src.Package.Namespaces[0].Namespace,
					src.Package.Namespaces[0].Names[0].Name, "", "", []string{})

				tableRows = append(tableRows, table.Row{hasSrcAtStr, src.Id, "Source for Package: " + purl})
			}
		}
	case hashEqualStr:
		for _, hash := range collectedNeighbors.hashEquals {
			tableRows = append(tableRows, table.Row{hashEqualStr, hash.Id, ""})
		}
	case occurrenceStr:
		for _, occurrence := range collectedNeighbors.occurrences {
			if subjectType == artifactSubjectType {
				switch v := occurrence.Subject.(type) {
				case *model.AllIsOccurrencesTreeSubjectPackage:
					purl := helpers.PkgToPurl(v.Type, v.Namespaces[0].Namespace,
						v.Namespaces[0].Names[0].Name, v.Namespaces[0].Names[0].Versions[0].Version, "", []string{})

					tableRows = append(tableRows, table.Row{occurrenceStr, occurrence.Id, "Occurrence for Package: " + purl})
				case *model.AllIsOccurrencesTreeSubjectSource:
					namespace := ""
					if !strings.HasPrefix(v.Namespaces[0].Namespace, "https://") {
						namespace = "https://" + v.Namespaces[0].Namespace
					} else {
						namespace = v.Namespaces[0].Namespace
					}
					tableRows = append(tableRows, table.Row{occurrenceStr, occurrence.Id, "Occurrence for Package: " + v.Type + "+" + namespace + "/" +
						v.Namespaces[0].Names[0].Name})
				}
			} else {
				tableRows = append(tableRows, table.Row{occurrenceStr, occurrence.Id, "Occurrence for Artifact: " + occurrence.Artifact.Algorithm + ":" + occurrence.Artifact.Digest})
			}
		}
	case pkgEqualStr:
		for _, equal := range collectedNeighbors.pkgEquals {
			tableRows = append(tableRows, table.Row{pkgEqualStr, equal.Id, ""})
		}
	case certifyLegalStr:
		for _, legal := range collectedNeighbors.certifyLegals {
			tableRows = append(tableRows, table.Row{
				certifyLegalStr,
				legal.Id,
				"Declared License: " + legal.DeclaredLicense +
					",\nDiscovered License: " + legal.DiscoveredLicense +
					",\nOrigin: " + legal.Origin,
			})
		}
	}

	return tableRows
}

func getAssociatedArtifact(ctx context.Context, gqlclient graphql.Client, occurrence *model.NeighborsNeighborsIsOccurrence, edge model.Edge) (*model.NeighborsResponse, error) {
	logger := logging.FromContext(ctx)
	artifactFilter := &model.ArtifactSpec{
		Algorithm: &occurrence.Artifact.Algorithm,
		Digest:    &occurrence.Artifact.Digest,
	}
	artifactResponse, err := model.Artifacts(ctx, gqlclient, *artifactFilter)
	if err != nil {
		logger.Debugf("error querying for artifacts: %v", err)
	}
	if len(artifactResponse.Artifacts) != 1 {
		logger.Debugf("failed to located artifacts based on (algorithm:digest)")
	}
	return model.Neighbors(ctx, gqlclient, artifactResponse.Artifacts[0].Id, []model.Edge{edge})
}

func validateQueryKnownFlags(graphqlEndpoint, headerFile string, args []string) (queryKnownOptions, error) {
	var opts queryKnownOptions
	opts.graphqlEndpoint = graphqlEndpoint
	opts.headerFile = headerFile

	if len(args) != 2 {
		return opts, fmt.Errorf("expected positional arguments for <type> <subject>")
	}
	opts.subjectType = args[0]
	if opts.subjectType != "package" && opts.subjectType != "source" && opts.subjectType != "artifact" {
		return opts, fmt.Errorf("expected type to be either a package, source or artifact")
	}
	opts.subject = args[1]

	return opts, nil
}

func init() {
	queryCmd.AddCommand(queryKnownCmd)
}
