package content

import (
  "github.com/Liquid-Labs/catalyst-core-api/go/resources/entities"
  "github.com/Liquid-Labs/go-nullable-mysql/nulls"
)

type ContentSource struct {
  entities.Entity
  ContentNamespace
  SourceType nulls.String                  `json:"sourceType"`
  Config     map[string]nulls.String `json:"config"`
}

const CreateContentSourceQuery = `INSERT INTO content_sources (id, source_type) VALUES(?,?)`
const GetContentSourceQuery = `SELECT e.id, e.pub_id, e.last_updated, ns.name, cs.source_type, csc.key, csc.value FROM entites e JOIN content_sources cs ON e.id=cs.id JOIN content_namespaces ns ON cs.namespace=ns.id JOIN content_sources_config csc ON csc.source=cs.id WHERE e.pub_id=?`
// config refreshed
const UpdateContentSourceConfigDeleteQuery = `DELETE FROM content_sources_config WHERE source=?`
// the refresh is built up as a single, batch insert query.
