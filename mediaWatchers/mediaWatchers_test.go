package mediaWatchers

import (
	"meteor/mediaplayers"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockBroadcaster struct {
	mock.Mock
}

type mockStore struct {
	mock.Mock
}

func (b *mockBroadcaster) Broadcast(info interface{}) {
	b.Called(info)
}

func (s *mockStore) UpdatePosition(mediaFile string, position int) error {
	s.Called(mediaFile, position)
	return nil
}

func (s *mockStore) UpdateDuration(mediaFile string, duration int) error {
	s.Called(mediaFile, duration)
	return nil
}

func TestMediaStateBroadcasterSendsBroadcast(t *testing.T) {
	status := mediaplayers.MediaPlayerStatus{Duration: 0, Position: 0, State: 0}
	broadcaster := &mockBroadcaster{}

	broadcaster.On("Broadcast", status)

	mediaBroadcaster := MediaStateNotifier{broadcaster}

	mediaBroadcaster.Notify(status)

	broadcaster.AssertExpectations(t)
}

func TestUpdatesWhenMediaIsPlaying(t *testing.T) {
	status := mediaplayers.MediaPlayerStatus{NowPlaying: "Test.mkv", Duration: 1000, Position: 10, State: 2}
	store := &mockStore{}

	store.On("UpdatePosition", status.NowPlaying, status.Position)
	store.On("UpdateDuration", status.NowPlaying, status.Duration)

	mediaStore := MediaStateRecorder{store}

	mediaStore.Notify(status)

	store.AssertExpectations(t)
}

func TestDoesNotUpdatesWhenMediaIsNotPlaying(t *testing.T) {
	status := mediaplayers.MediaPlayerStatus{NowPlaying: "Test.mkv", Duration: 1000, Position: 10, State: 0}
	store := &mockStore{}

	mediaStore := MediaStateRecorder{store}

	mediaStore.Notify(status)

	store.AssertNotCalled(t, "UpdatePosition")
	store.AssertNotCalled(t, "UpdateDuration")
}
