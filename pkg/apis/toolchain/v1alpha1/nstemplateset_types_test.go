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
		second *NSTemplateSetSpec
		want   bool
	}{
		{
			name: "both_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: tmpl1dev},
					{Type: "code", Revision: "rev1", Template: tmpl1code},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "code", Revision: "rev1", Template: tmpl2code},
					{Type: "dev", Revision: "rev1", Template: tmpl2dev},
				},
			},
			want: true,
		},

		{
			name: "tier_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
			},
			second: &NSTemplateSetSpec{
				TierName: "advance",
			},
			want: false,
		},

		{
			name: "ns_count_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: ""},
					{Type: "code", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: ""},
				},
			},
			want: false,
		},

		{
			name: "ns_revision_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: ""},
					{Type: "code", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: ""},
					{Type: "code", Revision: "rev2", Template: ""},
				},
			},
			want: false,
		},

		{
			name: "ns_type_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: ""},
					{Type: "code", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: ""},
					{Type: "stage", Revision: "rev1", Template: ""},
				},
			},
			want: false,
		},

		{
			name: "template_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "dev", Revision: "rev1", Template: tmpl1dev},
					{Type: "code", Revision: "rev1", Template: tmpl1code},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "code", Revision: "rev1", Template: tmpl2code},
					{Type: "dev", Revision: "rev1", Template: ""},
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
