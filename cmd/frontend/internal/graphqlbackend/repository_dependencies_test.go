package graphqlbackend

import (
	"testing"

	"github.com/neelance/graphql-go/gqltesting"

	"sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/backend"
	"sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/db"
	"sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/pkg/types"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/api"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/vcs"
)

func TestRepositoryResolver_Dependencies(t *testing.T) {
	resetMocks()

	backend.Mocks.Dependencies.List = func(*types.Repo, api.CommitID) ([]*api.DependencyReference, error) {
		return []*api.DependencyReference{{
			Language: "go",
			RepoID:   1,
			DepData: map[string]interface{}{
				"name": "d",
			},
		}}, nil
	}
	backend.Mocks.Repos.MockResolveRev_NoCheck(t, "cccccccccccccccccccccccccccccccccccccccc")
	backend.Mocks.Repos.MockGetCommit_Return_NoCheck(t, &vcs.Commit{})
	db.Mocks.Repos.MockGetByURI(t, "r", 1)

	gqltesting.RunTests(t, []*gqltesting.Test{
		{
			Schema: GraphQLSchema,
			Query: `
				{
					repository(uri: "r") {
						dependencies {
							nodes {
								language
								data {
									key
									value
								}
								dependingCommit {
									repository {
										uri
									}
								}
							}
							totalCount
							pageInfo {
								hasNextPage
							}
						}
					}
				}
		`,
			ExpectedResult: `
			{
				"repository": {
					"dependencies": {
						"nodes": [{
							"language": "go",
							"data": [
								{
									"key": "name",
									"value": "d"
								}
							],
							"dependingCommit": {
								"repository": {
									"uri": "r"
								}
							}
						}],
						"totalCount": 1,
						"pageInfo": {
							"hasNextPage": false
						}
					}
				}
			}
		`,
		},
	})
}
