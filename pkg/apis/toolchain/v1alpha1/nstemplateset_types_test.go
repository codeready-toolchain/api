package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tmpl1dev = `
	---
	apiVersion: v1
	kind: Template
	metadata:
	  labels:
		project: codeready-toolchain
	  name: codeready-toolchain-dev`

	tmpl2dev = `
	---
	apiVersion: v1
	kind: Template
	metadata:
	  labels:
		project: codeready-toolchain
	  name: codeready-toolchain-dev`

	tmpl1code = `
	---
	apiVersion: v1
	kind: Template
	metadata:
		labels:
		project: codeready-toolchain
		name: codeready-toolchain-code`

	tmpl2code = `
	---
	apiVersion: v1
	kind: Template
	metadata:
		labels:
		project: codeready-toolchain
		name: codeready-toolchain-code`
)

func TestNSTemplateSetSpecCompareTo(t *testing.T) {
	tables := []struct {
		name   string
		first  *NSTemplateSetSpec
		second NSTemplateSetSpec
		want   bool
	}{
		{
			name: "both_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: tmpl1dev},
					{Template: tmpl1code},
				},
			},
			second: NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: tmpl2code},
					{Template: tmpl2dev},
				},
			},
			want: true,
		},

		{
			name: "tier_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
			},
			second: NSTemplateSetSpec{
				TierName: "advance",
			},
			want: false,
		},

		{
			name: "ns_count_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: ""},
					{Template: ""},
				},
			},
			second: NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: ""},
				},
			},
			want: false,
		},

		{
			name: "template_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: tmpl1dev},
					{Template: tmpl1code},
				},
			},
			second: NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: tmpl2code},
					{Template: ""},
				},
			},
			want: false,
		},

		{
			name: "ns_revision_in_templateref_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: "", TemplateRef: "basic-dev-rev1"},
					{Template: "", TemplateRef: "basic-code-rev1"},
				},
			},
			second: NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: "", TemplateRef: "basic-dev-rev1"},
					{Template: "", TemplateRef: "basic-stage-rev2"},
				},
			},
			want: false,
		},

		{
			name: "ns_type_in_templateref_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: "", TemplateRef: "basic-dev-rev1"},
					{Template: "", TemplateRef: "basic-code-rev1"},
				},
			},
			second: NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []NSTemplateSetNamespace{
					{Template: "", TemplateRef: "basic-dev-rev1"},
					{Template: "", TemplateRef: "basic-stage-rev1"},
				},
			},
			want: false,
		},
	}

	for _, table := range tables {
		t.Run(table.name, func(t *testing.T) {
			got := table.first.CompareTo(table.second)
			assert.Equal(t, table.want, got)
		})
	}
}
