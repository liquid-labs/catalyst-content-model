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

func (c *ContributorSummary) SetRole(val string) {
  c.Role = nulls.NewString(val)
}

func (c *ContributorSummary) ClearRole() {
  c.Role = nulls.NewNullString()
}

func (c *ContributorSummary) SetSummaryCreditOrder(val int64) {
  c.SummaryCreditOrder = nulls.NewInt64(val)
}

func (c *ContributorSummary) ClearSummaryCreditOrder() {
  c.SummaryCreditOrder = nulls.NewNullInt64()
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

func (c *ContentSummary) SetTitle(val string) {
  c.Title = nulls.NewString(val)
}

func (c *ContentSummary) ClearTitle() {
  c.Title = nulls.NewNullString()
}

func (c *ContentSummary) SetSummary(val string) {
  c.Summary = nulls.NewString(val)
}

func (c *ContentSummary) ClearSummary() {
  c.Summary = nulls.NewNullString()
}

func (c *ContentSummary) SetNamespace(val string) {
  c.Namespace = nulls.NewString(val)
}

func (c *ContentSummary) ClearNamespace() {
  c.Namespace = nulls.NewNullString()
}

func (c *ContentSummary) SetSlug(val string) {
  c.Slug = nulls.NewString(val)
}

func (c *ContentSummary) ClearSlug() {
  c.Slug = nulls.NewNullString()
}

func (c *ContentSummary) SetType(val string) error {
  if c.PubId.IsValid() {
    return errors.New("Cannot change 'type' after creation.")
  } else {
    c.Type = nulls.NewString(val)
    return nil
  }
}

type ContentTypeText struct {
  ContentSummary
  Format        nulls.String `json:"format"`
  Text          nulls.String `json:"text"`
  // Contributors  ContributorSummaries `json:contributors`,
}

func (c *ContentTypeText) SetFormat(val string) {
  c.Format = nulls.NewString(val)
}

func (c *ContentTypeText) ClearFormat() {
  c.Format = nulls.NewNullString()
}

func (c *ContentTypeText) SetText(val string) {
  c.Text = nulls.NewString(val)
}

func (c *ContentTypeText) ClearText() {
  c.Text = nulls.NewNullString()
}
