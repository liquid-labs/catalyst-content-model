// Here we define the content resource model. The primary resource type is
// logically 'Content', but that is effectively an "abstract" type and cannot
// be directly instantiated. Instead, the user instantiates either a summary
// type containing general meta-data, but no actual content, or a concrete
// detail item type, such as 'ContentTypeText', which is currently the only
// concrete type supported.
//
// 'Contributors' is a relative-data type that connects a person resource to a
// particular piece of content through a 'role'. There are a number of standard
// roles, but generally any short string may be used.
//
// You may notice that there is a 'ContributorSummary', but no 'Contributor'
// type. The to retrive contributor details, you must access the corresponding
// person resource.
package content

import (
  "errors"

  "github.com/Liquid-Labs/catalyst-core-api/go/resources/entities"
  "github.com/Liquid-Labs/catalyst-persons-api/go/resources/persons"
  "github.com/Liquid-Labs/go-nullable-mysql/nulls"
)

type ContributorSummary struct {
  persons.PersonSummary
  Role               nulls.String `json:"role"`
  SummaryCreditOrder nulls.Int64  `json:"summaryCreditOrder"`
}

type ContributorSummaries []*ContributorSummary

func (c *ContributorSummary) GetRole() nulls.String { return c.Role }
func (c *ContributorSummary) SetRole(r nulls.String) { c.Role = r }

func (c *ContributorSummary) GetSummaryCreditORder() nulls.Int64 {
  return c.SummaryCreditOrder
}
func (c *ContributorSummary) SetSummaryCreditOrder(i nulls.Int64) {
  c.SummaryCreditOrder = i
}

// Content is an abstract type interface.
type Content interface {
  entities.EntityIface

  GetTitle() nulls.String
  SetTitle(nulls.String)

  GetSummary() nulls.String
  SetSummary(nulls.String)

  GetNamespace() nulls.String
  SetNamespace(nulls.String)

  GetSlug() nulls.String
  SetSlug(nulls.String)

  GetType() nulls.String
  SetType(nulls.String) error
}

type ContentSummary struct {
  entities.Entity
  Title         nulls.String `json:"title"`
  Summary       nulls.String `json:"summary"`
  ExternPath    nulls.String `json:"externPath"`
  Namespace     nulls.String `json:"summary"`
  Slug          nulls.String `json:"slug"`
  Type          nulls.String `json:"type"`
  LastSync      nulls.Int64  `json:"lastSync"`
  VersionCookie nulls.String `json:"versionCookie"`
  // TODO: want to name this '(K/k)eyContributors' to be more precise, but that
  // means we need to implement custom marshalling for the 'ContentSummary'
  Contributors ContributorSummaries `json:contributors`
}

func (c *ContentSummary) GetTitle() nulls.String { return c.Title }
func (c *ContentSummary) SetTitle(t nulls.String) { c.Title = t }

func (c *ContentSummary) GetSummary() nulls.String { return c.Summary }
func (c *ContentSummary) SetSummary(s nulls.String) { c.Summary = s }

func (c *ContentSummary) GetNamespace() nulls.String { return c.Namespace }
func (c *ContentSummary) SetNamespace(n nulls.String) { c.Namespace = n }

func (c *ContentSummary) GetSlug() nulls.String { return c.Slug }
func (c *ContentSummary) SetSlug(s nulls.String) { c.Slug = s }

func (c *ContentSummary) GetType() nulls.String { return c.Type }
func (c *ContentSummary) SetType(t nulls.String) error {
  if c.PubId.IsValid() {
    return errors.New("Cannot change 'type' after creation.")
  } else {
    c.Type = t
    return nil
  }
}

type ContentTypeText struct {
  ContentSummary
  Format        nulls.String `json:"format"`
  Text          nulls.String `json:"text"`
  // Contributors  ContributorSummaries `json:contributors`,
}

func (c *ContentTypeText) GetFormat() nulls.String { return c.Format }
func (c *ContentTypeText) SetFormat(f nulls.String) { c.Format = f }

func (c *ContentTypeText) GetText() nulls.String { return c.Text }
func (c *ContentTypeText) SetText(t nulls.String) { c.Text = t }
