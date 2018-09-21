package mediaWatchers

import (
	"meteor/mediaplayers"
)

type mediaStore interface {
	UpdatePosition(mediaFile string, position int) error
	UpdateDuration(mediaFile string, duration int) error
}

//MediaStateRecorder updates a media store whenever the state of a media item changes
type MediaStateRecorder struct {
	store mediaStore
}

//NewMediaStateRecorder returns a new instance MediaStateRecorder
func NewMediaStateRecorder(store mediaStore) MediaStateRecorder {
	return MediaStateRecorder{store: store}
}

//Notify updates the media store with the current state of the media if the media is in the playing state
func (m MediaStateRecorder) Notify(status mediaplayers.MediaPlayerStatus) {
	if status.State == 2 {
		m.store.UpdatePosition(status.NowPlaying, status.Position)
		m.store.UpdateDuration(status.NowPlaying, status.Duration)
	}
}
