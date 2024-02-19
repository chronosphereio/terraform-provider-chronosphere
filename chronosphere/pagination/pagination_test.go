// Copyright 2024 Chronosphere Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pagination

import (
	"context"
	"fmt"
	"testing"

	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/client/collection"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/mocks"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/pkg/configv1/models"
	"github.com/chronosphereio/terraform-provider-chronosphere/chronosphere/sliceutil"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPagination(t *testing.T) {
	tests := []struct {
		msg      string
		setMocks func(*mocks.MockClientService)
		want     []string
		wantErr  string
	}{
		{
			msg: "single page of results",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(listCollectionResponse{
					CollectionNames: []string{"c1", "c2"},
				}.build(), nil)
			},
			want: []string{"c1", "c2"},
		},
		{
			msg: "error getting first page",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError.Error(),
		},
		{
			msg: "multiple pages of results",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(listCollectionResponse{
					CollectionNames: []string{"c1", "c2"},
					Page:            "page2",
				}.build(), nil)
				m.EXPECT().ListCollections(listCollectionPageMatcher("page2")).Return(listCollectionResponse{
					CollectionNames: []string{"c3", "c4"},
					Page:            "page3",
				}.build(), nil)
				m.EXPECT().ListCollections(listCollectionPageMatcher("page3")).Return(listCollectionResponse{
					CollectionNames: []string{"c5", "c6"},
					Page:            "",
				}.build(), nil)
			},
			want: []string{"c1", "c2", "c3", "c4", "c5", "c6"},
		},
		{
			msg: "error on second page",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(listCollectionResponse{
					CollectionNames: []string{"c1", "c2"},
					Page:            "page2",
				}.build(), nil)
				m.EXPECT().ListCollections(listCollectionPageMatcher("page2")).Return(nil, assert.AnError)
			},
			wantErr: assert.AnError.Error(),
		},
		{
			msg: "missing payload on first page",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(&collection.ListCollectionsOK{}, nil)
			},
			want: nil,
		},
		{
			msg: "missing payload on second page",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(listCollectionResponse{
					CollectionNames: []string{"c1", "c2"},
					Page:            "page2",
				}.build(), nil)
				m.EXPECT().ListCollections(listCollectionPageMatcher("page2")).Return(&collection.ListCollectionsOK{}, nil)
			},
			want: []string{"c1", "c2"},
		},
		{
			msg: "missing Page on second page",
			setMocks: func(m *mocks.MockClientService) {
				m.EXPECT().ListCollections(listCollectionPageMatcher("")).Return(listCollectionResponse{
					CollectionNames: []string{"c1", "c2"},
					Page:            "page2",
				}.build(), nil)
				m.EXPECT().ListCollections(listCollectionPageMatcher("page2")).Return(&collection.ListCollectionsOK{
					Payload: &models.Configv1ListCollectionsResponse{
						Collections: collectionsFromNames("c3", "c4"),
						Page:        nil,
					},
				}, nil)
			},
			want: []string{"c1", "c2", "c3", "c4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := mocks.NewMockClientService(ctrl)
			tt.setMocks(client)

			got, err := ListCollections(context.Background(), &configv1.Client{
				Collection: client,
			})
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, collectionsToNames(got))
		})
	}
}

type listCollectionPageMatcher string

func (m listCollectionPageMatcher) Matches(x interface{}) bool {
	params := x.(*collection.ListCollectionsParams)

	gotPage := ""
	if params.PageToken != nil {
		gotPage = *params.PageToken
	}

	return string(m) == gotPage
}

func (m listCollectionPageMatcher) String() string {
	return fmt.Sprintf("page=%q", string(m))
}

type listCollectionResponse struct {
	Page            string
	CollectionNames []string
}

func (l listCollectionResponse) build() *collection.ListCollectionsOK {
	return &collection.ListCollectionsOK{
		Payload: &models.Configv1ListCollectionsResponse{
			Collections: collectionsFromNames(l.CollectionNames...),
			Page: &models.Configv1PageResult{
				NextToken: l.Page,
			},
		},
	}
}

func collectionsFromNames(names ...string) []*models.Configv1Collection {
	return sliceutil.Map(names, func(name string) *models.Configv1Collection {
		return &models.Configv1Collection{
			Name: name,
		}
	})
}

func collectionsToNames(collections []*models.Configv1Collection) []string {
	return sliceutil.Map(collections, func(c *models.Configv1Collection) string {
		return c.Name
	})
}
