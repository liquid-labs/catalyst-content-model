import { entityPropsModel, Model } from '@liquid-labs/catalyst-core-api'
import { personPropsModel } from '@liquid-labs/catalyst-persons-api'

const contributorPropsModel = [ 'role', 'summaryCreditOrder']
  .map((propName) => ({ propName, writable : true }))
contributorPropsModel.push(...personPropsModel)

const Contributor = class extends Model {
  get resourceName() { return 'contributors' }
}
Model.finalizeConstructor(Contributor, contributorPropsModel)

const contentTextTypePropsModel = [
  'title',
  'summary',
  'format',
  'text',
  'slug',
  'externPath',
  'lastSync',
  'version_cookie',
].map((propName) => ({ propName, writable : true }))
contributorPropsModel.push(...entityPropsModel)
contentTextTypePropsModel.push([ 'keyContributors', 'contributors']
  .map((propName) => ({ propName, model: Contributor, valueType: arrayType, writable: true }))

const ContentTypeText = class extends Model {
  get resourceName() { return 'content' }
}
Model.finalizeConstructor(ContentTypeText, contentTextTypePropsModel)
