package socialmedia

import (
	"time"
)

// go:generate stringer -type=MoodState
type MoodState int

// Here we define all the possible mood states using an
// iota numerator.
const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateCloudNine
)

// This is a type we embed into types we want to keep a
// check on for auditing purposes.
type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

// This is the type that represents a social media post
type Post struct {
	AuditableContent // Embeded type
	Caption          string
	MessageBody      string
	URL              string
	ImageURI         string
	ThumbnailURI     string
	Keywords         []string
	Likers           []string
	AuthorMood       MoodState
}

// Map that holds the various mood states with keys to serve as
// aliases to their respective mood state
var Moods map[string]MoodState

// The init() function is responsible for initialising the mood state
func init() {

	Moods = map[string]MoodState{"neutral": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry,
		"hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored,
		"shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateCloudNine}

}

// This is the function responsible for creating a new social media post.
func NewPost(
	username string,
	mood MoodState,
	caption string,
	messageBody string,
	url string,
	imageurl string,
	thumbnailurl string,
	keywords []string,
	likers []string) *Post {

	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageurl, ThumbnailURI: thumbnailurl,
		AuthorMood: mood, Keywords: keywords, AuditableContent: auditableContent}
}
