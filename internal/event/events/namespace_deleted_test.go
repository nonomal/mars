package events

import (
	"testing"

	"github.com/duc-cnzj/mars/internal/testutil"

	"github.com/duc-cnzj/mars/internal/app/instance"
	"github.com/duc-cnzj/mars/internal/mock"
	"github.com/duc-cnzj/mars/internal/models"

	"github.com/golang/mock/gomock"
)

func TestHandleNamespaceDeleted(t *testing.T) {
	ctrl := gomock.NewController(t)
	app := mock.NewMockApplicationInterface(ctrl)
	defer ctrl.Finish()
	instance.SetInstance(app)
	pubsub := mock.NewMockPubSub(ctrl)
	sender := testutil.MockWsServer(ctrl, app)
	sender.EXPECT().New("", "").Return(pubsub)
	pubsub.EXPECT().ToAll(&EventNamespaceDeletedMatcher{nsID: 1}).Times(1)
	HandleNamespaceDeleted(&models.Namespace{ID: 1}, EventNamespaceDeleted)
}
