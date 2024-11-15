package services

import (
	"context"
	"testing"

	"github.com/duc-cnzj/mars/api/v5/picture"
	"github.com/duc-cnzj/mars/v5/internal/application"
	"github.com/duc-cnzj/mars/v5/internal/repo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNewPictureSvc(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	svc := NewPictureSvc(repo.NewMockPictureRepo(m))
	assert.NotNil(t, svc)
	assert.NotNil(t, svc.(*pictureSvc).picRepo)
}

func Test_pictureSvc_Background(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	picRepo := repo.NewMockPictureRepo(m)
	svc := NewPictureSvc(picRepo)
	picRepo.EXPECT().Get(gomock.Any(), true).Return(&application.PictureItem{}, nil)
	_, err := svc.Background(context.TODO(), &picture.BackgroundRequest{Random: true})
	assert.Nil(t, err)

	picRepo.EXPECT().Get(gomock.Any(), true).Return(nil, assert.AnError)
	_, err = svc.Background(context.TODO(), &picture.BackgroundRequest{Random: true})
	assert.NotNil(t, err)
}
