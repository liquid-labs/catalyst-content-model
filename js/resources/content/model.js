import { entityPropsModel, Model } from '@liquid-labs/catalyst-core-api'
import { personPropsModel } from '@liquid-labs/catalyst-persons-api'

const contributorPropsModel = [ 'role', 'summaryCreditOrder']
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
contributorPropsModel.push({ propName : 'type', writeable : false })

contentTextTypePropsModel.push([ 'keyContributors', 'contributors']
  .map((propName) => ({ propName, model : Contributor, valueType : arrayType, writable : true })))

const ContentTypeText = class extends Model {
  get resourceName() { return 'content' }
}
Model.finalizeConstructor(ContentTypeText, contentTextTypePropsModel)

export {
  contributorPropsModel,
  Contributor,
  contentTextTypePropsModel,
  ContentTypeText,
}
