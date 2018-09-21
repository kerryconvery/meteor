package mediaWatchers

import (
	"meteor/mediaplayers"
)

type notificationListener interface {
	Broadcast(payload interface{})
}

//MediaStateNotifier wraps a mediaBroadcaster that can listen for media changes and broadcast the media state
type MediaStateNotifier struct {
	listener notificationListener
}

//NewStateNotifier returns a new instance of media mediaBroadcaster
func NewStateNotifier(listener notificationListener) MediaStateNotifier {
	return MediaStateNotifier{listener}
}

//Notify broadcasts the media status
func (m MediaStateNotifier) Notify(status mediaplayers.MediaPlayerStatus) {
	m.listener.Broadcast(status)
}
