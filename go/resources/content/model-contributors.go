package content

import (
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

// update by refresh only
const ContributorsDeleteQuery = `DELETE FROM contributors WHERE content=?`
const ContributorInsertQuery = `INSERT INTO contributors (id, content, role, summary_credit_order) SELECT p.id, c.id, ?, ? FROM entities pe JOIN persons p ON pe.id=p.id JOIN content_summary c JOIN entities ce ON ce.id=c.id WHERE pe.pub_id=? AND ce.pub_id=?`
const ContributorInsertWithContentIDQuery = `INSERT INTO contributors (id, content, role, summary_credit_order) SELECT p.id, ?, ?, ? FROM entities e JOIN persons p ON e.id=p.id WHERE e.pub_id=?`
