package v1alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
					{Type: "cicd", Revision: "rev1", Template: ""},
					{Type: "ide", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "ide", Revision: "rev1", Template: ""},
					{Type: "cicd", Revision: "rev1", Template: ""},
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
					{Type: "cicd", Revision: "rev1", Template: ""},
					{Type: "ide", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "cicd", Revision: "rev1", Template: ""},
				},
			},
			want: false,
		},

		{
			name: "ns_revision_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "cicd", Revision: "rev1", Template: ""},
					{Type: "ide", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "cicd", Revision: "rev1", Template: ""},
					{Type: "ide", Revision: "rev2", Template: ""},
				},
			},
			want: false,
		},

		{
			name: "ns_type_not_same",
			first: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "cicd", Revision: "rev1", Template: ""},
					{Type: "ide", Revision: "rev1", Template: ""},
				},
			},
			second: &NSTemplateSetSpec{
				TierName: "basic",
				Namespaces: []Namespace{
					{Type: "cicd", Revision: "rev1", Template: ""},
					{Type: "stage", Revision: "rev1", Template: ""},
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
