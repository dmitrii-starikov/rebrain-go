package counter

import (
	"errors"
	"module06/internal/app/services/post"
	"module06/test/gomock/mocks/postmock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestPostCountTable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name          string
		mockPosts     []post.Post
		mockError     error
		expectedCount int
		expectedError error
	}{
		{
			name: "success with multiple posts",
			mockPosts: []post.Post{
				{ID: 1, Title: "Post 1"},
				{ID: 2, Title: "Post 2"},
			},
			mockError:     nil,
			expectedCount: 2,
			expectedError: nil,
		},
		{
			name:          "success with empty posts",
			mockPosts:     []post.Post{},
			mockError:     nil,
			expectedCount: 0,
			expectedError: nil,
		},
		{
			name:          "error from client",
			mockPosts:     nil,
			mockError:     errors.New("api error"),
			expectedCount: 0,
			expectedError: errors.New("api error"),
		},
	}

	for _, tc := range testCases {
		tc := tc // closure
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // all the expected calls - Times(1) - are done

			mockClient := postmock.NewMockClient(ctrl)
			mockClient.EXPECT().
				GetList().
				Times(1).
				// Times(2). // = missing call(s) to *postmock.MockClient.GetList()
				Return(tc.mockPosts, tc.mockError)

			count, err := PostCount(mockClient)

			// Assert
			if tc.expectedError != nil {
				require.Error(t, err)
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tc.expectedCount, count)
		})
	}
}
