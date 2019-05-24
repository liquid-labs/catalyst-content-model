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
