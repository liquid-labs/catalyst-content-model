package content

import (
  "github.com/Liquid-Labs/catalyst-core-api/go/resources/entities"
  "github.com/Liquid-Labs/go-nullable-mysql/nulls"
)

type ContentNamespace struct {
  entities.Entity
  Name       nulls.String    `json:"name"`
}

const CreateContentNamespaceQuery = `INSERT INTO content_namespaces (id, name) VALUES(?,?)`
const GetContentNamespaceQuery = `SELECT e.id, e.pub_id, e.last_updated, ns.name FROM entities e JOIN content_namespaces ns ON e.id=ns.id WHERE e.pub_id=?`
