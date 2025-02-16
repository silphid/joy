package list

import (
	"fmt"
	"os"
	"strings"

	"github.com/nestoca/joy/api/v1alpha1"

	"github.com/olekukonko/tablewriter"

	"github.com/nestoca/joy/internal/git"
	"github.com/nestoca/joy/internal/release/filtering"
	"github.com/nestoca/joy/internal/style"
	"github.com/nestoca/joy/pkg/catalog"
)

type Opts struct {
	// CatalogDir is the path to the catalog directory.
	CatalogDir string

	// SelectedEnvs is the list of environments that were selected by user to work with.
	SelectedEnvs []string

	// Filter specifies releases to list.
	// Optional, defaults to listing all releases.
	Filter filtering.Filter
}

func List(opts Opts) error {
	err := git.EnsureCleanAndUpToDateWorkingCopy(opts.CatalogDir)
	if err != nil {
		return err
	}

	// Load catalog
	loadOpts := catalog.LoadOpts{
		Dir:             opts.CatalogDir,
		LoadEnvs:        true,
		LoadReleases:    true,
		EnvNames:        opts.SelectedEnvs,
		SortEnvsByOrder: true,
		ReleaseFilter:   opts.Filter,
	}
	cat, err := catalog.Load(loadOpts)
	if err != nil {
		return fmt.Errorf("loading catalog: %w", err)
	}

	// Set up table writer
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetAutoWrapText(true)
	table.SetBorder(false)
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetCenterSeparator("")

	// Add headers
	releases := cat.Releases
	headers := []string{"NAME"}
	for _, env := range releases.Environments {
		headers = append(headers, strings.ToUpper(env.Name))
	}
	table.SetHeader(headers)

	// Add rows
	for _, crossRelease := range releases.Items {
		row := []string{crossRelease.Name}
		for _, rel := range crossRelease.Releases {
			displayVersion := GetReleaseDisplayVersion(rel)
			row = append(row, displayVersion)
		}
		table.Append(row)
	}

	table.Render()
	return nil
}

func GetReleaseDisplayVersion(rel *v1alpha1.Release) string {
	if rel == nil {
		return style.ReleaseNotAvailable("-")
	}
	version := rel.Spec.Version
	if version == "" {
		version = "no version"
	}
	return version
}
