package content

import (
  "errors"

  "github.com/Liquid-Labs/catalyst-core-api/go/resources/entities"
  "github.com/Liquid-Labs/go-nullable-mysql/nulls"
)

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
  Namespace     nulls.String `json:"namespace"` // the string name
  SourceType    nulls.String `json:"-"`
  Slug          nulls.String `json:"slug"`
  Type          nulls.String `json:"type"`
  LastSync      nulls.Int64  `json:"lastSync"`
  VersionCookie nulls.String `json:"versionCookie"`
  // TODO: want to name this '(K/k)eyContributors' to be more precise, but that
  // means we need to implement custom marshalling for the 'ContentSummary'
  Contributors  ContributorSummaries `json:contributors`
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

const CreateContentQuery = `INSERT INTO content_summary (id, extern_path, slug, type, title, summary, version_cookie) VALUES(?,?,?,?,?,?,?)`

// retrieve -- partials
const CommonContentFields = `e.pub_id, e.last_updated, c.title, c.summary, ns.name, cs.source_type, c.slug, c.type `
const CommonCententTypeTextFields = `ctt.format, ctt.text, c.extern_path, c.last_sync, c.version_cookie `
const CommonContentContribFields = `e.pub_id, p.display_name, pc.role, pc.summary_credit_order `
const CommonContentFrom = `FROM entities e JOIN content_summary c ON e.id=c.id JOIN contributors pc ON c.id=pc.content JOIN persons p ON pc.id=p.id JOIN content_namespaces ns ON c.namespace=ns.id JOIN content_sources cs ON c.source=cs.id `

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

// create
const CreateContentTypeTextQuery = `INSERT INTO content_type_text (id, format, text) VALUES(?,?,?)`

// retrieve -- partials
const CommonContentTypeTextFields = CommonContentFields + `, ctt.format, ctt.text `
const CommonContentTypeTextFrom = `JOIN content_type_text ctt ON c.id=ctt.id `
const CommonContentTypeTextGet string = `SELECT ` + CommonContentTypeTextFields + CommonContentFrom + CommonContentTypeTextFrom
// retrieve -- queries
const GetContentTypeTextQuery string = CommonContentTypeTextGet + `WHERE e.pub_id=? `
const GetContentTypeTextByNSSlugQuery string = CommonContentTypeTextGet + ` WHERE ns.name=? AND c.slug=? `
const GetContentTypeTextByIDQuery string = CommonContentTypeTextGet + ` WHERE c.id=? `

// update
const UpdateContentTypeTextSansTextQuery = `UPDATE content_summary c JOIN content_type_text ctt ON c.id=ctt.id JOIN entities e ON c.id=e.id SET e.last_updated=0, c.title=?, c.summary=?, c.extern_path=?, c.slug=?, ctt.format=? WHERE e.pub_id=?`
const UpdateContentTypeTextWithTextQuery = `UPDATE content_summary c JOIN content_type_text ctt ON c.id=ctt.id JOIN entities e ON c.id=e.id SET e.last_updated=0, ctt.last_sync=0, c.title=?, c.summary=?, c.extern_path=?, c.slug=?, ctt.format=?, ctt.text=? WHERE e.pub_id=?`
const UpdateContentTypeTextOnlyTextQuery = `UPDATE content_summary c JOIN content_type_text ctt ON c.id=ctt.id JOIN entities e ON c.id=e.id SET e.last_updated=0, ctt.last_sync=0, ctt.text=? WHERE e.pub_id=?`
