import { arrayType, CommonResourceConf, entityPropsModel, Model } from '@liquid-labs/catalyst-core-api'
import { personPropsModel } from '@liquid-labs/catalyst-persons-api'

const contributorPropsModel = [ 'role', 'summaryCreditOrder' ]
  .map((propName) => ({ propName, writable : true }))
contributorPropsModel.push(...personPropsModel)

const Contributor = class extends Model {
  get resourceName() { return 'contributors' }
}
Model.finalizeConstructor(Contributor, contributorPropsModel)

const contentTextTypePropsModel = [...entityPropsModel]
contentTextTypePropsModel.push(...[
  'title',
  'summary',
  'format',
  'slug',
  'externPath',
  'lastSync',
  'version_cookie',
].map((propName) => ({ propName, writable : true })))
contentTextTypePropsModel.push({ propName : 'type', writeable : false })

contentTextTypePropsModel.push([ 'keyContributors', 'contributors' ]
  .map((propName) => ({ propName, model : Contributor, valueType : arrayType, writable : true })))

const ContentTypeText = class extends Model {
  get resourceName() { return 'content' }
}
Model.finalizeConstructor(ContentTypeText, contentTextTypePropsModel)

const contentResourceConf = new CommonResourceConf('content', {
  model       : ContentTypeText,
  sortOptions : [
    { label : 'Title (asc)',
      value : 'title-asc',
      func  : (a, b) => a.title.localeCompare(b.title) },
    { label : 'Title (desc)',
      value : 'title-desc',
      func  : (a, b) => -a.title.localeCompare(b.title) }
  ],
  sortDefault : 'title-asc'
}, { resourceName : 'content' })

contentResourceConf.urlMapper = (source) => {
  // TODO: keep current namespace if present.
  const key = `REACT_APP_CONTENT_NAMESPACE`
  const namespace = process.env[key]
  console.log(process.env)
  if (namespace === undefined) {
    throw new Error(`No '${key}' found. Try in '.env' and rebuilding application.`)
  }
  return `${source}?namespace=${encodeURIComponent(namespace)}`
}

export {
  contributorPropsModel,
  Contributor,
  contentTextTypePropsModel,
  ContentTypeText,
  contentResourceConf,
}
