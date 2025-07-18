// inner_events.go provides EventsAPI particular inner events

package slackevents

import (
	"encoding/json"

	"github.com/slack-go/slack"
)

// EventsAPIInnerEvent the inner event of a EventsAPI event_callback Event.
type EventsAPIInnerEvent struct {
	Type string `json:"type"`
	Data interface{}
}

// AssistantThreadMessageEvent is an (inner) EventsAPI subscribable event.
type AssistantThreadStartedEvent struct {
	Type            string          `json:"type"`
	AssistantThread AssistantThread `json:"assistant_thread"`
	EventTimestamp  string          `json:"event_ts"`
}

// AssistantThreadChangedEvent is an (inner) EventsAPI subscribable event.
type AssistantThreadContextChangedEvent struct {
	Type            string          `json:"type"`
	AssistantThread AssistantThread `json:"assistant_thread"`
	EventTimestamp  string          `json:"event_ts"`
}

// AssistantThread is an object that represents a thread of messages between a user and an assistant.
type AssistantThread struct {
	UserID          string                 `json:"user_id"`
	Context         AssistantThreadContext `json:"context"`
	ChannelID       string                 `json:"channel_id"`
	ThreadTimeStamp string                 `json:"thread_ts"`
}

// AssistantThreadContext is an object that represents the context of an assistant thread.
type AssistantThreadContext struct {
	ChannelID    string `json:"channel_id"`
	TeamID       string `json:"team_id"`
	EnterpriseID string `json:"enterprise_id"`
}

// AppMentionEvent is an (inner) EventsAPI subscribable event.
type AppMentionEvent struct {
	Type            string `json:"type"`
	User            string `json:"user"`
	Text            string `json:"text"`
	TimeStamp       string `json:"ts"`
	ThreadTimeStamp string `json:"thread_ts"`
	Channel         string `json:"channel"`
	EventTimeStamp  string `json:"event_ts"`

	// When Message comes from a channel that is shared between workspaces
	UserTeam   string `json:"user_team,omitempty"`
	SourceTeam string `json:"source_team,omitempty"`

	// BotID is filled out when a bot triggers the app_mention event
	BotID string `json:"bot_id,omitempty"`

	// When the app is mentioned in the edited message
	Edited *Edited `json:"edited,omitempty"`
}

// AppHomeOpenedEvent Your Slack app home was opened.
type AppHomeOpenedEvent struct {
	Type           string      `json:"type"`
	User           string      `json:"user"`
	Channel        string      `json:"channel"`
	EventTimeStamp string      `json:"event_ts"`
	Tab            string      `json:"tab"`
	View           *slack.View `json:"view,omitempty"`
}

// AppUninstalledEvent Your Slack app was uninstalled.
type AppUninstalledEvent struct {
	Type string `json:"type"`
}

// ChannelCreatedEvent represents the Channel created event
type ChannelCreatedEvent struct {
	Type           string             `json:"type"`
	Channel        ChannelCreatedInfo `json:"channel"`
	EventTimestamp string             `json:"event_ts"`
}

// ChannelDeletedEvent represents the Channel deleted event
type ChannelDeletedEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelArchiveEvent represents the Channel archive event
type ChannelArchiveEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	User           string `json:"user"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelUnarchiveEvent represents the Channel unarchive event
type ChannelUnarchiveEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	User           string `json:"user"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelLeftEvent represents the Channel left event
type ChannelLeftEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelRenameEvent represents the Channel rename event
type ChannelRenameEvent struct {
	Type           string            `json:"type"`
	Channel        ChannelRenameInfo `json:"channel"`
	EventTimestamp string            `json:"event_ts"`
}

// ChannelIDChangedEvent represents the Channel identifier changed event
type ChannelIDChangedEvent struct {
	Type           string `json:"type"`
	OldChannelID   string `json:"old_channel_id"`
	NewChannelID   string `json:"new_channel_id"`
	EventTimestamp string `json:"event_ts"`
}

// ChannelCreatedInfo represents the information associated with the Channel created event
type ChannelCreatedInfo struct {
	ID        string `json:"id"`
	IsChannel bool   `json:"is_channel"`
	Name      string `json:"name"`
	Created   int    `json:"created"`
	Creator   string `json:"creator"`
}

// ChannelRenameInfo represents the information associated with the Channel rename event
type ChannelRenameInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
}

// ChannelUnsharedEvent represents a channel has been unshared with an external workspace event
type ChannelUnsharedEvent struct {
	Type                      string `json:"type"`
	PreviouslyConnectedTeamID string `json:"previously_connected_team_id"`
	Channel                   string `json:"channel"`
	IsExtShared               bool   `json:"is_ext_shared"`
	EventTimestamp            string `json:"event_ts"`
}

// GroupDeletedEvent represents the Group deleted event
type GroupDeletedEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// GroupArchiveEvent represents the Group archive event
type GroupArchiveEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// GroupUnarchiveEvent represents the Group unarchive event
type GroupUnarchiveEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// GroupLeftEvent represents the Group left event
type GroupLeftEvent struct {
	Type           string `json:"type"`
	Channel        string `json:"channel"`
	EventTimestamp string `json:"event_ts"`
}

// GroupRenameEvent represents the Group rename event
type GroupRenameEvent struct {
	Type           string          `json:"type"`
	Channel        GroupRenameInfo `json:"channel"`
	EventTimestamp string          `json:"event_ts"`
}

// GroupRenameInfo represents the information associated with the Group rename event
type GroupRenameInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created int    `json:"created"`
}

// FileChangeEvent represents the information associated with the File change
// event.
type FileChangeEvent struct {
	Type   string        `json:"type"`
	FileID string        `json:"file_id"`
	File   FileEventFile `json:"file"`
}

// FileDeletedEvent represents the information associated with the File deleted
// event.
type FileDeletedEvent struct {
	Type           string `json:"type"`
	FileID         string `json:"file_id"`
	EventTimestamp string `json:"event_ts"`
}

// FileSharedEvent represents the information associated with the File shared
// event.
type FileSharedEvent struct {
	Type           string        `json:"type"`
	ChannelID      string        `json:"channel_id"`
	FileID         string        `json:"file_id"`
	UserID         string        `json:"user_id"`
	File           FileEventFile `json:"file"`
	EventTimestamp string        `json:"event_ts"`
}

// FileUnsharedEvent represents the information associated with the File
// unshared event.
type FileUnsharedEvent struct {
	Type   string        `json:"type"`
	FileID string        `json:"file_id"`
	File   FileEventFile `json:"file"`
}

// FileEventFile represents information on the specific file being shared in a
// file-related Slack event.
type FileEventFile struct {
	ID string `json:"id"`
}

// GridMigrationFinishedEvent An enterprise grid migration has finished on this workspace.
type GridMigrationFinishedEvent struct {
	Type         string `json:"type"`
	EnterpriseID string `json:"enterprise_id"`
}

// GridMigrationStartedEvent An enterprise grid migration has started on this workspace.
type GridMigrationStartedEvent struct {
	Type         string `json:"type"`
	EnterpriseID string `json:"enterprise_id"`
}

// LinkSharedEvent A message was posted containing one or more links relevant to your application
type LinkSharedEvent struct {
	Type      string `json:"type"`
	User      string `json:"user"`
	TimeStamp string `json:"ts"`
	Channel   string `json:"channel"`
	// MessageTimeStamp can be both a numeric timestamp if the LinkSharedEvent corresponds to a sent
	// message and (contrary to the field name) a uuid if the LinkSharedEvent is generated in the
	// compose text area.
	MessageTimeStamp string        `json:"message_ts"`
	ThreadTimeStamp  string        `json:"thread_ts"`
	Links            []SharedLinks `json:"links"`
	EventTimestamp   string        `json:"event_ts"`
}

type SharedLinks struct {
	Domain string `json:"domain"`
	URL    string `json:"url"`
}

// MessageEvent occurs when a variety of types of messages has been posted.
// Parse ChannelType to see which
// if ChannelType = "group", this is a private channel message
// if ChannelType = "channel", this message was sent to a channel
// if ChannelType = "im", this is a private message
// if ChannelType = "mim", A message was posted in a multiparty direct message channel
// TODO: Improve this so that it is not required to manually parse ChannelType
type MessageEvent struct {
	// Basic Message Event - https://api.slack.com/events/message
	ClientMsgID     string `json:"client_msg_id"`
	Type            string `json:"type"`
	User            string `json:"user"`
	Text            string `json:"text"`
	ThreadTimeStamp string `json:"thread_ts"`
	TimeStamp       string `json:"ts"`
	Channel         string `json:"channel"`
	ChannelType     string `json:"channel_type"`
	EventTimeStamp  string `json:"event_ts"`

	// When Message comes from a channel that is shared between workspaces
	UserTeam   string `json:"user_team,omitempty"`
	SourceTeam string `json:"source_team,omitempty"`

	// When we get a 'message' event with no subtype, i.e. telling us about a new
	// message, the message information is stored at the top level. But when we get
	// a 'message_changed' event, the message information is stored in
	// the Message property. This is really hard to represent nicely in Go, so we use
	// a custom JSON unmarshaller to populate the Message field in both cases.
	Message *slack.Msg `json:"message,omitempty"`
	// Root is set if the SubType is `thread_broadcast`.
	Root *slack.Msg `json:"root,omitempty"`
	// Edited Message
	PreviousMessage *slack.Msg `json:"previous_message,omitempty"`

	// Deleted Message
	DeletedTimeStamp string `json:"deleted_ts,omitempty"`

	// Message Subtypes
	SubType string `json:"subtype,omitempty"`

	// bot_message (https://api.slack.com/events/message/bot_message)
	BotID    string `json:"bot_id,omitempty"`
	Username string `json:"username,omitempty"`
	Icons    *Icon  `json:"icons,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for MessageEvent.
// This custom unmarshaler handles both regular messages and message_changed events
// by normalizing the message data into the Message field.
func (e *MessageEvent) UnmarshalJSON(data []byte) error {
	// First, unmarshal into an anonymous struct to avoid infinite recursion
	// when calling json.Unmarshal on the MessageEvent type itself
	type MessageEventAlias MessageEvent
	alias := struct {
		MessageEventAlias
	}{}

	if err := json.Unmarshal(data, &alias.MessageEventAlias); err != nil {
		return err
	}

	// Copy all fields from alias to the original struct
	*e = MessageEvent(alias.MessageEventAlias)

	// Now check if there's no Message field (which would happen for regular messages)
	if e.Message == nil {
		// For regular messages, the message content is at the top level,
		// so we need to unmarshal the data again into a slack.Msg
		msg := &slack.Msg{}
		if err := json.Unmarshal(data, msg); err != nil {
			return err
		}

		// Set the Message field to the unmarshaled msg
		e.Message = msg
	}

	return nil
}

// MemberJoinedChannelEvent A member joined a public or private channel
type MemberJoinedChannelEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Channel        string `json:"channel"`
	ChannelType    string `json:"channel_type"`
	Team           string `json:"team"`
	Inviter        string `json:"inviter"`
	EventTimestamp string `json:"event_ts"`
}

// MemberLeftChannelEvent A member left a public or private channel
type MemberLeftChannelEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Channel        string `json:"channel"`
	ChannelType    string `json:"channel_type"`
	Team           string `json:"team"`
	EventTimestamp string `json:"event_ts"`
}

type pinEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Item           Item   `json:"item"`
	Channel        string `json:"channel_id"`
	EventTimestamp string `json:"event_ts"`
	HasPins        bool   `json:"has_pins,omitempty"`
}

type reactionEvent struct {
	Type           string `json:"type"`
	User           string `json:"user"`
	Reaction       string `json:"reaction"`
	ItemUser       string `json:"item_user"`
	Item           Item   `json:"item"`
	EventTimestamp string `json:"event_ts"`
}

// ReactionAddedEvent An reaction was added to a message - https://api.slack.com/events/reaction_added
type ReactionAddedEvent reactionEvent

// ReactionRemovedEvent An reaction was removed from a message - https://api.slack.com/events/reaction_removed
type ReactionRemovedEvent reactionEvent

// PinAddedEvent An item was pinned to a channel - https://api.slack.com/events/pin_added
type PinAddedEvent pinEvent

// PinRemovedEvent An item was unpinned from a channel - https://api.slack.com/events/pin_removed
type PinRemovedEvent pinEvent

type tokens struct {
	Oauth []string `json:"oauth"`
	Bot   []string `json:"bot"`
}

// TeamJoinEvent A new member joined a workspace -  https://api.slack.com/events/team_join
type TeamJoinEvent struct {
	Type           string      `json:"type"`
	User           *slack.User `json:"user"`
	EventTimestamp string      `json:"event_ts"`
}

// TokensRevokedEvent APP's API tokens are revoked - https://api.slack.com/events/tokens_revoked
type TokensRevokedEvent struct {
	Type           string `json:"type"`
	Tokens         tokens `json:"tokens"`
	EventTimestamp string `json:"event_ts"`
}

// EmojiChangedEvent is the event of custom emoji has been added or changed
type EmojiChangedEvent struct {
	Type           string `json:"type"`
	Subtype        string `json:"subtype"`
	EventTimeStamp string `json:"event_ts"`

	// filled out when custom emoji added
	Name string `json:"name,omitempty"`

	// filled out when custom emoji removed
	Names []string `json:"names,omitempty"`

	// filled out when custom emoji renamed
	OldName string `json:"old_name,omitempty"`
	NewName string `json:"new_name,omitempty"`

	// filled out when custom emoji added or renamed
	Value string `json:"value,omitempty"`
}

// MessageMetadataPostedEvent is sent, if a message with metadata is posted
type MessageMetadataPostedEvent struct {
	Type             string               `json:"type"`
	AppId            string               `json:"app_id"`
	BotId            string               `json:"bot_id"`
	UserId           string               `json:"user_id"`
	TeamId           string               `json:"team_id"`
	ChannelId        string               `json:"channel_id"`
	Metadata         *slack.SlackMetadata `json:"metadata"`
	MessageTimestamp string               `json:"message_ts"`
	EventTimestamp   string               `json:"event_ts"`
}

// MessageMetadataUpdatedEvent is sent, if a message with metadata is deleted
type MessageMetadataUpdatedEvent struct {
	Type             string               `json:"type"`
	ChannelId        string               `json:"channel_id"`
	EventTimestamp   string               `json:"event_ts"`
	PreviousMetadata *slack.SlackMetadata `json:"previous_metadata"`
	AppId            string               `json:"app_id"`
	BotId            string               `json:"bot_id"`
	UserId           string               `json:"user_id"`
	TeamId           string               `json:"team_id"`
	MessageTimestamp string               `json:"message_ts"`
	Metadata         *slack.SlackMetadata `json:"metadata"`
}

// MessageMetadataDeletedEvent is sent, if a message with metadata is deleted
type MessageMetadataDeletedEvent struct {
	Type             string               `json:"type"`
	ChannelId        string               `json:"channel_id"`
	EventTimestamp   string               `json:"event_ts"`
	PreviousMetadata *slack.SlackMetadata `json:"previous_metadata"`
	AppId            string               `json:"app_id"`
	BotId            string               `json:"bot_id"`
	UserId           string               `json:"user_id"`
	TeamId           string               `json:"team_id"`
	MessageTimestamp string               `json:"message_ts"`
	DeletedTimestamp string               `json:"deleted_ts"`
}

// JSONTime exists so that we can have a String method converting the date
type JSONTime int64

// Comment contains all the information relative to a comment
type Comment struct {
	ID        string   `json:"id,omitempty"`
	Created   JSONTime `json:"created,omitempty"`
	Timestamp JSONTime `json:"timestamp,omitempty"`
	User      string   `json:"user,omitempty"`
	Comment   string   `json:"comment,omitempty"`
}

// File is a file upload
type File struct {
	ID                 string `json:"id"`
	Created            int    `json:"created"`
	Timestamp          int    `json:"timestamp"`
	Name               string `json:"name"`
	Title              string `json:"title"`
	Mimetype           string `json:"mimetype"`
	Filetype           string `json:"filetype"`
	PrettyType         string `json:"pretty_type"`
	User               string `json:"user"`
	Editable           bool   `json:"editable"`
	Size               int    `json:"size"`
	Mode               string `json:"mode"`
	IsExternal         bool   `json:"is_external"`
	ExternalType       string `json:"external_type"`
	IsPublic           bool   `json:"is_public"`
	PublicURLShared    bool   `json:"public_url_shared"`
	DisplayAsBot       bool   `json:"display_as_bot"`
	Username           string `json:"username"`
	URLPrivate         string `json:"url_private"`
	FileAccess         string `json:"file_access"`
	URLPrivateDownload string `json:"url_private_download"`
	Thumb64            string `json:"thumb_64"`
	Thumb80            string `json:"thumb_80"`
	Thumb360           string `json:"thumb_360"`
	Thumb360W          int    `json:"thumb_360_w"`
	Thumb360H          int    `json:"thumb_360_h"`
	Thumb480           string `json:"thumb_480"`
	Thumb480W          int    `json:"thumb_480_w"`
	Thumb480H          int    `json:"thumb_480_h"`
	Thumb160           string `json:"thumb_160"`
	Thumb720           string `json:"thumb_720"`
	Thumb720W          int    `json:"thumb_720_w"`
	Thumb720H          int    `json:"thumb_720_h"`
	Thumb800           string `json:"thumb_800"`
	Thumb800W          int    `json:"thumb_800_w"`
	Thumb800H          int    `json:"thumb_800_h"`
	Thumb960           string `json:"thumb_960"`
	Thumb960W          int    `json:"thumb_960_w"`
	Thumb960H          int    `json:"thumb_960_h"`
	Thumb1024          string `json:"thumb_1024"`
	Thumb1024W         int    `json:"thumb_1024_w"`
	Thumb1024H         int    `json:"thumb_1024_h"`
	ImageExifRotation  int    `json:"image_exif_rotation"`
	OriginalW          int    `json:"original_w"`
	OriginalH          int    `json:"original_h"`
	Permalink          string `json:"permalink"`
	PermalinkPublic    string `json:"permalink_public"`
}

// Edited is included when a Message is edited
type Edited struct {
	User      string `json:"user"`
	TimeStamp string `json:"ts"`
}

// Icon is used for bot messages
type Icon struct {
	IconURL   string `json:"icon_url,omitempty"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

// Item is any type of slack message - message, file, or file comment.
type Item struct {
	Type      string       `json:"type"`
	Channel   string       `json:"channel,omitempty"`
	Message   *ItemMessage `json:"message,omitempty"`
	File      *File        `json:"file,omitempty"`
	Comment   *Comment     `json:"comment,omitempty"`
	Timestamp string       `json:"ts,omitempty"`
}

// ItemMessage is the event message
type ItemMessage struct {
	Type            string   `json:"type"`
	User            string   `json:"user"`
	Text            string   `json:"text"`
	Timestamp       string   `json:"ts"`
	PinnedTo        []string `json:"pinned_to"`
	ReplaceOriginal bool     `json:"replace_original"`
	DeleteOriginal  bool     `json:"delete_original"`
}

// IsEdited checks if the MessageEvent is caused by an edit
func (e MessageEvent) IsEdited() bool {
	return e.Message != nil &&
		e.Message.Edited != nil
}

// TeamAccessGrantedEvent is sent if access to teams was granted for your org-wide app.
type TeamAccessGrantedEvent struct {
	Type    string   `json:"type"`
	TeamIDs []string `json:"team_ids"`
}

// TeamAccessRevokedEvent is sent if access to teams was revoked for your org-wide app.
type TeamAccessRevokedEvent struct {
	Type    string   `json:"type"`
	TeamIDs []string `json:"team_ids"`
}

// UserProfileChangedEvent is sent if access to teams was revoked for your org-wide app.
type UserProfileChangedEvent struct {
	User    *slack.User `json:"user"`
	CacheTs int         `json:"cache_ts"`
	Type    string      `json:"type"`
	EventTs string      `json:"event_ts"`
}

// SharedChannelInviteApprovedEvent is sent if your invitation has been approved
type SharedChannelInviteApprovedEvent struct {
	Type            string              `json:"type"`
	Invite          *SharedInvite       `json:"invite"`
	Channel         *slack.Conversation `json:"channel"`
	ApprovingTeamID string              `json:"approving_team_id"`
	TeamsInChannel  []*SlackEventTeam   `json:"teams_in_channel"`
	ApprovingUser   *SlackEventUser     `json:"approving_user"`
	EventTs         string              `json:"event_ts"`
}

// SharedChannelInviteAcceptedEvent is sent if external org accepts a Slack Connect channel invite
type SharedChannelInviteAcceptedEvent struct {
	Type                string            `json:"type"`
	ApprovalRequired    bool              `json:"approval_required"`
	Invite              *SharedInvite     `json:"invite"`
	Channel             *SharedChannel    `json:"channel"`
	TeamsInChannel      []*SlackEventTeam `json:"teams_in_channel"`
	AcceptingUser       *SlackEventUser   `json:"accepting_user"`
	EventTs             string            `json:"event_ts"`
	RequiresSponsorship bool              `json:"requires_sponsorship,omitempty"`
}

// SharedChannelInviteDeclinedEvent is sent if external or internal org declines the Slack Connect invite
type SharedChannelInviteDeclinedEvent struct {
	Type            string            `json:"type"`
	Invite          *SharedInvite     `json:"invite"`
	Channel         *SharedChannel    `json:"channel"`
	DecliningTeamID string            `json:"declining_team_id"`
	TeamsInChannel  []*SlackEventTeam `json:"teams_in_channel"`
	DecliningUser   *SlackEventUser   `json:"declining_user"`
	EventTs         string            `json:"event_ts"`
}

// SharedChannelInviteReceivedEvent is sent if a bot or app is invited to a Slack Connect channel
type SharedChannelInviteReceivedEvent struct {
	Type    string         `json:"type"`
	Invite  *SharedInvite  `json:"invite"`
	Channel *SharedChannel `json:"channel"`
	EventTs string         `json:"event_ts"`
}

// SlackEventTeam is a struct for teams in ShareChannel events
type SlackEventTeam struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	Icon                *SlackEventIcon `json:"icon,omitempty"`
	AvatarBaseURL       string          `json:"avatar_base_url,omitempty"`
	IsVerified          bool            `json:"is_verified"`
	Domain              string          `json:"domain"`
	DateCreated         int             `json:"date_created"`
	RequiresSponsorship bool            `json:"requires_sponsorship,omitempty"`
	// TeamID              string          `json:"team_id,omitempty"`
}

// SlackEventIcon is a struct for icons in ShareChannel events
type SlackEventIcon struct {
	ImageDefault bool   `json:"image_default,omitempty"`
	Image34      string `json:"image_34,omitempty"`
	Image44      string `json:"image_44,omitempty"`
	Image68      string `json:"image_68,omitempty"`
	Image88      string `json:"image_88,omitempty"`
	Image102     string `json:"image_102,omitempty"`
	Image132     string `json:"image_132,omitempty"`
	Image230     string `json:"image_230,omitempty"`
}

// SlackEventUser is a struct for users in ShareChannel events
type SlackEventUser struct {
	ID                     string             `json:"id"`
	TeamID                 string             `json:"team_id"`
	Name                   string             `json:"name"`
	Updated                int                `json:"updated,omitempty"`
	Profile                *slack.UserProfile `json:"profile,omitempty"`
	WhoCanShareContactCard string             `json:"who_can_share_contact_card,omitempty"`
}

// SharedChannel is a struct for shared channels in ShareChannel events
type SharedChannel struct {
	ID        string `json:"id"`
	IsPrivate bool   `json:"is_private"`
	IsIm      bool   `json:"is_im"`
	Name      string `json:"name,omitempty"`
}

// SharedInvite is a struct for shared invites in ShareChannel events
type SharedInvite struct {
	ID                string          `json:"id"`
	DateCreated       int             `json:"date_created"`
	DateInvalid       int             `json:"date_invalid"`
	InvitingTeam      *SlackEventTeam `json:"inviting_team,omitempty"`
	InvitingUser      *SlackEventUser `json:"inviting_user,omitempty"`
	RecipientEmail    string          `json:"recipient_email,omitempty"`
	RecipientUserID   string          `json:"recipient_user_id,omitempty"`
	IsSponsored       bool            `json:"is_sponsored,omitempty"`
	IsExternalLimited bool            `json:"is_external_limited,omitempty"`
}

type ChannelHistoryChangedEvent struct {
	Type    string `json:"type"`
	Latest  string `json:"latest"`
	Ts      string `json:"ts"`
	EventTs string `json:"event_ts"`
}

type CommandsChangedEvent struct {
	Type    string `json:"type"`
	EventTs string `json:"event_ts"`
}

type DndUpdatedEvent struct {
	Type      string `json:"type"`
	User      string `json:"user"`
	DndStatus struct {
		DndEnabled     bool  `json:"dnd_enabled"`
		NextDndStartTs int64 `json:"next_dnd_start_ts"`
		NextDndEndTs   int64 `json:"next_dnd_end_ts"`
		SnoozeEnabled  bool  `json:"snooze_enabled"`
		SnoozeEndtime  int64 `json:"snooze_endtime"`
	} `json:"dnd_status"`
}

type DndUpdatedUserEvent struct {
	Type      string `json:"type"`
	User      string `json:"user"`
	DndStatus struct {
		DndEnabled     bool  `json:"dnd_enabled"`
		NextDndStartTs int64 `json:"next_dnd_start_ts"`
		NextDndEndTs   int64 `json:"next_dnd_end_ts"`
	} `json:"dnd_status"`
}

type EmailDomainChangedEvent struct {
	Type        string `json:"type"`
	EmailDomain string `json:"email_domain"`
	EventTs     string `json:"event_ts"`
}

type GroupCloseEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type GroupHistoryChangedEvent struct {
	Type    string `json:"type"`
	Latest  string `json:"latest"`
	Ts      string `json:"ts"`
	EventTs string `json:"event_ts"`
}

type GroupOpenEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type ImCloseEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type ImCreatedEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel struct {
		ID string `json:"id"`
	} `json:"channel"`
}

type ImHistoryChangedEvent struct {
	Type    string `json:"type"`
	Latest  string `json:"latest"`
	Ts      string `json:"ts"`
	EventTs string `json:"event_ts"`
}

type ImOpenEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type SubTeam struct {
	ID          string `json:"id"`
	TeamID      string `json:"team_id"`
	IsUsergroup bool   `json:"is_usergroup"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Handle      string `json:"handle"`
	IsExternal  bool   `json:"is_external"`
	DateCreate  int64  `json:"date_create"`
	DateUpdate  int64  `json:"date_update"`
	DateDelete  int64  `json:"date_delete"`
	AutoType    string `json:"auto_type"`
	CreatedBy   string `json:"created_by"`
	UpdatedBy   string `json:"updated_by"`
	DeletedBy   string `json:"deleted_by"`
	Prefs       struct {
		Channels []string `json:"channels"`
		Groups   []string `json:"groups"`
	} `json:"prefs"`
	Users     []string `json:"users"`
	UserCount int      `json:"user_count"`
}

type SubteamCreatedEvent struct {
	Type    string  `json:"type"`
	Subteam SubTeam `json:"subteam"`
}

type SubteamMembersChangedEvent struct {
	Type               string   `json:"type"`
	SubteamID          string   `json:"subteam_id"`
	TeamID             string   `json:"team_id"`
	DatePreviousUpdate int      `json:"date_previous_update"`
	DateUpdate         int64    `json:"date_update"`
	AddedUsers         []string `json:"added_users"`
	AddedUsersCount    int      `json:"added_users_count"`
	RemovedUsers       []string `json:"removed_users"`
	RemovedUsersCount  int      `json:"removed_users_count"`
}

type SubteamSelfAddedEvent struct {
	Type      string `json:"type"`
	SubteamID string `json:"subteam_id"`
}

type SubteamSelfRemovedEvent struct {
	Type      string `json:"type"`
	SubteamID string `json:"subteam_id"`
}

type SubteamUpdatedEvent struct {
	Type    string  `json:"type"`
	Subteam SubTeam `json:"subteam"`
}

type TeamDomainChangeEvent struct {
	Type   string `json:"type"`
	URL    string `json:"url"`
	Domain string `json:"domain"`
	TeamID string `json:"team_id"`
}

type TeamRenameEvent struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	TeamID string `json:"team_id"`
}

type UserChangeEvent struct {
	Type    string `json:"type"`
	User    User   `json:"user"`
	CacheTS int64  `json:"cache_ts"`
	EventTS string `json:"event_ts"`
}

type AppDeletedEvent struct {
	Type       string `json:"type"`
	AppID      string `json:"app_id"`
	AppName    string `json:"app_name"`
	AppOwnerID string `json:"app_owner_id"`
	TeamID     string `json:"team_id"`
	TeamDomain string `json:"team_domain"`
	EventTs    string `json:"event_ts"`
}

type AppInstalledEvent struct {
	Type       string `json:"type"`
	AppID      string `json:"app_id"`
	AppName    string `json:"app_name"`
	AppOwnerID string `json:"app_owner_id"`
	UserID     string `json:"user_id"`
	TeamID     string `json:"team_id"`
	TeamDomain string `json:"team_domain"`
	EventTs    string `json:"event_ts"`
}

type AppRequestedEvent struct {
	Type       string `json:"type"`
	AppRequest struct {
		ID  string `json:"id"`
		App struct {
			ID                     string `json:"id"`
			Name                   string `json:"name"`
			Description            string `json:"description"`
			HelpURL                string `json:"help_url"`
			PrivacyPolicyURL       string `json:"privacy_policy_url"`
			AppHomepageURL         string `json:"app_homepage_url"`
			AppDirectoryURL        string `json:"app_directory_url"`
			IsAppDirectoryApproved bool   `json:"is_app_directory_approved"`
			IsInternal             bool   `json:"is_internal"`
			AdditionalInfo         string `json:"additional_info"`
		} `json:"app"`
		PreviousResolution struct {
			Status string `json:"status"`
			Scopes []struct {
				Name        string `json:"name"`
				Description string `json:"description"`
				IsSensitive bool   `json:"is_sensitive"`
				TokenType   string `json:"token_type"`
			} `json:"scopes"`
		} `json:"previous_resolution"`
		User struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"user"`
		Team struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Domain string `json:"domain"`
		} `json:"team"`
		Enterprise interface{} `json:"enterprise"`
		Scopes     []struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			IsSensitive bool   `json:"is_sensitive"`
			TokenType   string `json:"token_type"`
		} `json:"scopes"`
		Message string `json:"message"`
	} `json:"app_request"`
}

type AppUninstalledTeamEvent struct {
	Type       string `json:"type"`
	AppID      string `json:"app_id"`
	AppName    string `json:"app_name"`
	AppOwnerID string `json:"app_owner_id"`
	UserID     string `json:"user_id"`
	TeamID     string `json:"team_id"`
	TeamDomain string `json:"team_domain"`
	EventTs    string `json:"event_ts"`
}

type CallRejectedEvent struct {
	Token    string `json:"token"`
	TeamID   string `json:"team_id"`
	APIAppID string `json:"api_app_id"`
	Event    struct {
		Type             string `json:"type"`
		CallID           string `json:"call_id"`
		UserID           string `json:"user_id"`
		ChannelID        string `json:"channel_id"`
		ExternalUniqueID string `json:"external_unique_id"`
	} `json:"event"`
	Type        string   `json:"type"`
	EventID     string   `json:"event_id"`
	AuthedUsers []string `json:"authed_users"`
}

type ChannelSharedEvent struct {
	Type            string `json:"type"`
	ConnectedTeamID string `json:"connected_team_id"`
	Channel         string `json:"channel"`
	EventTs         string `json:"event_ts"`
}

type FileCreatedEvent struct {
	Type   string `json:"type"`
	FileID string `json:"file_id"`
	File   struct {
		ID string `json:"id"`
	} `json:"file"`
}

type FilePublicEvent struct {
	Type   string `json:"type"`
	FileID string `json:"file_id"`
	File   struct {
		ID string `json:"id"`
	} `json:"file"`
}

type FunctionExecutedEvent struct {
	Type     string `json:"type"`
	Function struct {
		ID              string `json:"id"`
		CallbackID      string `json:"callback_id"`
		Title           string `json:"title"`
		Description     string `json:"description"`
		Type            string `json:"type"`
		InputParameters []struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Title       string `json:"title"`
			IsRequired  bool   `json:"is_required"`
		} `json:"input_parameters"`
		OutputParameters []struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Title       string `json:"title"`
			IsRequired  bool   `json:"is_required"`
		} `json:"output_parameters"`
		AppID       string `json:"app_id"`
		DateCreated int64  `json:"date_created"`
		DateUpdated int64  `json:"date_updated"`
		DateDeleted int64  `json:"date_deleted"`
	} `json:"function"`
	Inputs              map[string]interface{} `json:"inputs"`
	FunctionExecutionID string                 `json:"function_execution_id"`
	WorkflowExecutionID string                 `json:"workflow_execution_id"`
	EventTs             string                 `json:"event_ts"`
	BotAccessToken      string                 `json:"bot_access_token"`
}

type InviteRequestedEvent struct {
	Type          string `json:"type"`
	InviteRequest struct {
		ID            string   `json:"id"`
		Email         string   `json:"email"`
		DateCreated   int64    `json:"date_created"`
		RequesterIDs  []string `json:"requester_ids"`
		ChannelIDs    []string `json:"channel_ids"`
		InviteType    string   `json:"invite_type"`
		RealName      string   `json:"real_name"`
		DateExpire    int64    `json:"date_expire"`
		RequestReason string   `json:"request_reason"`
		Team          struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Domain string `json:"domain"`
		} `json:"team"`
	} `json:"invite_request"`
}

type StarAddedEvent struct {
	Type string `json:"type"`
	User string `json:"user"`
	Item struct {
	} `json:"item"`
	EventTS string `json:"event_ts"`
}

type StarRemovedEvent struct {
	Type string `json:"type"`
	User string `json:"user"`
	Item struct {
	} `json:"item"`
	EventTS string `json:"event_ts"`
}

type UserHuddleChangedEvent struct {
	Type    string `json:"type"`
	User    User   `json:"user"`
	CacheTS int64  `json:"cache_ts"`
	EventTS string `json:"event_ts"`
}

type User struct {
	ID                     string  `json:"id"`
	TeamID                 string  `json:"team_id"`
	Name                   string  `json:"name"`
	Deleted                bool    `json:"deleted"`
	Color                  string  `json:"color"`
	RealName               string  `json:"real_name"`
	TZ                     string  `json:"tz"`
	TZLabel                string  `json:"tz_label"`
	TZOffset               int     `json:"tz_offset"`
	Profile                Profile `json:"profile"`
	IsAdmin                bool    `json:"is_admin"`
	IsOwner                bool    `json:"is_owner"`
	IsPrimaryOwner         bool    `json:"is_primary_owner"`
	IsRestricted           bool    `json:"is_restricted"`
	IsUltraRestricted      bool    `json:"is_ultra_restricted"`
	IsBot                  bool    `json:"is_bot"`
	IsAppUser              bool    `json:"is_app_user"`
	Updated                int64   `json:"updated"`
	IsEmailConfirmed       bool    `json:"is_email_confirmed"`
	WhoCanShareContactCard string  `json:"who_can_share_contact_card"`
	Locale                 string  `json:"locale"`
}

type Profile struct {
	Title                  string                 `json:"title"`
	Phone                  string                 `json:"phone"`
	Skype                  string                 `json:"skype"`
	RealName               string                 `json:"real_name"`
	RealNameNormalized     string                 `json:"real_name_normalized"`
	DisplayName            string                 `json:"display_name"`
	DisplayNameNormalized  string                 `json:"display_name_normalized"`
	Fields                 map[string]interface{} `json:"fields"`
	StatusText             string                 `json:"status_text"`
	StatusEmoji            string                 `json:"status_emoji"`
	StatusEmojiDisplayInfo []interface{}          `json:"status_emoji_display_info"`
	StatusExpiration       int                    `json:"status_expiration"`
	AvatarHash             string                 `json:"avatar_hash"`
	FirstName              string                 `json:"first_name"`
	LastName               string                 `json:"last_name"`
	Image24                string                 `json:"image_24"`
	Image32                string                 `json:"image_32"`
	Image48                string                 `json:"image_48"`
	Image72                string                 `json:"image_72"`
	Image192               string                 `json:"image_192"`
	Image512               string                 `json:"image_512"`
	StatusTextCanonical    string                 `json:"status_text_canonical"`
	Team                   string                 `json:"team"`
}

type UserStatusChangedEvent struct {
	Type    string `json:"type"`
	User    User   `json:"user"`
	CacheTS int64  `json:"cache_ts"`
	EventTS string `json:"event_ts"`
}

type Actor struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsBot       bool   `json:"is_bot"`
	TeamID      string `json:"team_id"`
	Timezone    string `json:"timezone"`
	RealName    string `json:"real_name"`
	DisplayName string `json:"display_name"`
}

type TargetUser struct {
	Email    string `json:"email"`
	InviteID string `json:"invite_id"`
}

type TeamIcon struct {
	Image34      string `json:"image_34"`
	ImageDefault bool   `json:"image_default"`
}

type Team struct {
	ID                  string   `json:"id"`
	Icon                TeamIcon `json:"icon"`
	Name                string   `json:"name"`
	Domain              string   `json:"domain"`
	IsVerified          bool     `json:"is_verified"`
	DateCreated         int64    `json:"date_created"`
	AvatarBaseURL       string   `json:"avatar_base_url"`
	RequiresSponsorship bool     `json:"requires_sponsorship"`
}

type SharedChannelInviteRequestedEvent struct {
	Actor                       Actor        `json:"actor"`
	ChannelID                   string       `json:"channel_id"`
	EventType                   string       `json:"event_type"`
	ChannelName                 string       `json:"channel_name"`
	ChannelType                 string       `json:"channel_type"`
	TargetUsers                 []TargetUser `json:"target_users"`
	TeamsInChannel              []Team       `json:"teams_in_channel"`
	IsExternalLimited           bool         `json:"is_external_limited"`
	ChannelDateCreated          int64        `json:"channel_date_created"`
	ChannelMessageLatestCounted int64        `json:"channel_message_latest_counted_timestamp"`
}

type EventsAPIType string

const (
	// AppDeleted is an event when an app is deleted from a workspace
	AppDeleted = EventsAPIType("app_deleted")
	// AppHomeOpened Your Slack app home was opened
	AppHomeOpened = EventsAPIType("app_home_opened")
	// AppInstalled is an event when an app is installed to a workspace
	AppInstalled = EventsAPIType("app_installed")
	// AppMention is an Events API subscribable event
	AppMention = EventsAPIType("app_mention")
	// AppRequested is an event when a user requests to install an app to a workspace
	AppRequested = EventsAPIType("app_requested")
	// AppUninstalled Your Slack app was uninstalled.
	AppUninstalled = EventsAPIType("app_uninstalled")
	// AppUninstalledTeam is an event when an app is uninstalled from a team
	AppUninstalledTeam = EventsAPIType("app_uninstalled_team")
	// AssistantThreadContextChanged Your Slack AI Assistant has changed the context of a thread
	AssistantThreadContextChanged = EventsAPIType("assistant_thread_context_changed")
	// AssistantThreadStarted Your Slack AI Assistant has started a new thread
	AssistantThreadStarted = EventsAPIType("assistant_thread_started")
	// CallRejected is an event when a Slack call is rejected
	CallRejected = EventsAPIType("call_rejected")
	// ChannelArchive is sent when a channel is archived.
	ChannelArchive = EventsAPIType("channel_archive")
	// ChannelCreated is sent when a new channel is created.
	ChannelCreated = EventsAPIType("channel_created")
	// ChannelDeleted is sent when a channel is deleted.
	ChannelDeleted = EventsAPIType("channel_deleted")
	// ChannelHistoryChanged The history of a channel changed
	ChannelHistoryChanged = EventsAPIType("channel_history_changed")
	// ChannelIDChanged is sent when a channel identifier is changed.
	ChannelIDChanged = EventsAPIType("channel_id_changed")
	// ChannelLeft is sent when a channel is left.
	ChannelLeft = EventsAPIType("channel_left")
	// ChannelRename is sent when a channel is rename.
	ChannelRename = EventsAPIType("channel_rename")
	// ChannelShared is an event when a channel is shared with another workspace
	ChannelShared = EventsAPIType("channel_shared")
	// ChannelUnarchive is sent when a channel is unarchived.
	ChannelUnarchive = EventsAPIType("channel_unarchive")
	// ChannelUnshared is sent when a channel is unshared.
	ChannelUnshared = EventsAPIType("channel_unshared")
	// CommandsChanged A command was changed
	CommandsChanged = EventsAPIType("commands_changed")
	// DndUpdated Do Not Disturb settings were updated
	DndUpdated = EventsAPIType("dnd_updated")
	// DndUpdatedUser Do Not Disturb settings for a user were updated
	DndUpdatedUser = EventsAPIType("dnd_updated_user")
	// EmailDomainChanged The email domain changed
	EmailDomainChanged = EventsAPIType("email_domain_changed")
	// EmojiChanged A custom emoji has been added or changed
	EmojiChanged = EventsAPIType("emoji_changed")
	// FileChange is sent when a file is changed.
	FileChange = EventsAPIType("file_change")
	// FileCreated is an event when a file is created in a workspace
	FileCreated = EventsAPIType("file_created")
	// FileDeleted is sent when a file is deleted.
	FileDeleted = EventsAPIType("file_deleted")
	// FilePublic is an event when a file is made public in a workspace
	FilePublic = EventsAPIType("file_public")
	// FileShared is sent when a file is shared.
	FileShared = EventsAPIType("file_shared")
	// FileUnshared is sent when a file is unshared.
	FileUnshared = EventsAPIType("file_unshared")
	// FunctionExecuted is an event when a Slack function is executed
	FunctionExecuted = EventsAPIType("function_executed")
	// GridMigrationFinished An enterprise grid migration has finished on this workspace.
	GridMigrationFinished = EventsAPIType("grid_migration_finished")
	// GridMigrationStarted An enterprise grid migration has started on this workspace.
	GridMigrationStarted = EventsAPIType("grid_migration_started")
	// GroupArchive is sent when a group is archived.
	GroupArchive = EventsAPIType("group_archive")
	// GroupClose A group was closed
	GroupClose = EventsAPIType("group_close")
	// GroupDeleted is sent when a group is deleted.
	GroupDeleted = EventsAPIType("group_deleted")
	// GroupHistoryChanged The history of a group changed
	GroupHistoryChanged = EventsAPIType("group_history_changed")
	// GroupLeft is sent when a group is left.
	GroupLeft = EventsAPIType("group_left")
	// GroupOpen A group was opened
	GroupOpen = EventsAPIType("group_open")
	// GroupRename is sent when a group is renamed.
	GroupRename = EventsAPIType("group_rename")
	// GroupUnarchive is sent when a group is unarchived.
	GroupUnarchive = EventsAPIType("group_unarchive")
	// ImClose An instant message channel was closed
	ImClose = EventsAPIType("im_close")
	// ImCreated An instant message channel was created
	ImCreated = EventsAPIType("im_created")
	// ImHistoryChanged The history of an instant message channel changed
	ImHistoryChanged = EventsAPIType("im_history_changed")
	// ImOpen An instant message channel was opened
	ImOpen = EventsAPIType("im_open")
	// InviteRequested is an event when a user requests an invite to a workspace
	InviteRequested = EventsAPIType("invite_requested")
	// LinkShared A message was posted containing one or more links relevant to your application
	LinkShared = EventsAPIType("link_shared")
	// MemberJoinedChannel is sent if a member joined a channel.
	MemberJoinedChannel = EventsAPIType("member_joined_channel")
	// MemberLeftChannel is sent if a member left a channel.
	MemberLeftChannel = EventsAPIType("member_left_channel")
	// Message A message was posted to a channel, private channel (group), im, or mim
	Message = EventsAPIType("message")
	// MessageMetadataDeleted A message with metadata was deleted
	MessageMetadataDeleted = EventsAPIType("message_metadata_deleted")
	// MessageMetadataPosted A message with metadata was posted
	MessageMetadataPosted = EventsAPIType("message_metadata_posted")
	// MessageMetadataUpdated A message with metadata was updated
	MessageMetadataUpdated = EventsAPIType("message_metadata_updated")
	// PinAdded An item was pinned to a channel
	PinAdded = EventsAPIType("pin_added")
	// PinRemoved An item was unpinned from a channel
	PinRemoved = EventsAPIType("pin_removed")
	// ReactionAdded An reaction was added to a message
	ReactionAdded = EventsAPIType("reaction_added")
	// ReactionRemoved An reaction was removed from a message
	ReactionRemoved = EventsAPIType("reaction_removed")
	// SharedChannelInviteAccepted Slack connect channel invite accepted by an end user
	SharedChannelInviteAccepted = EventsAPIType("shared_channel_invite_accepted")
	// SharedChannelInviteApproved Slack connect channel invite approved
	SharedChannelInviteApproved = EventsAPIType("shared_channel_invite_approved")
	// SharedChannelInviteDeclined Slack connect channel invite declined
	SharedChannelInviteDeclined = EventsAPIType("shared_channel_invite_declined")
	// SharedChannelInviteReceived Slack connect app or bot invite received
	SharedChannelInviteReceived = EventsAPIType("shared_channel_invite_received")
	// SharedChannelInviteRequested is an event when an invitation to share a channel is requested
	SharedChannelInviteRequested = EventsAPIType("shared_channel_invite_requested")
	// StarAdded is an event when a star is added to a message or file
	StarAdded = EventsAPIType("star_added")
	// StarRemoved is an event when a star is removed from a message or file
	StarRemoved = EventsAPIType("star_removed")
	// SubteamCreated A subteam was created
	SubteamCreated = EventsAPIType("subteam_created")
	// SubteamMembersChanged The members of a subteam changed
	SubteamMembersChanged = EventsAPIType("subteam_members_changed")
	// SubteamSelfAdded The current user was added to a subteam
	SubteamSelfAdded = EventsAPIType("subteam_self_added")
	// SubteamSelfRemoved The current user was removed from a subteam
	SubteamSelfRemoved = EventsAPIType("subteam_self_removed")
	// SubteamUpdated A subteam was updated
	SubteamUpdated = EventsAPIType("subteam_updated")
	// TeamAccessGranted is sent if access to teams was granted for your org-wide app.
	TeamAccessGranted = EventsAPIType("team_access_granted")
	// TeamAccessRevoked is sent if access to teams was revoked for your org-wide app.
	TeamAccessRevoked = EventsAPIType("team_access_revoked")
	// TeamDomainChange The team's domain changed
	TeamDomainChange = EventsAPIType("team_domain_change")
	// TeamJoin A new user joined the workspace
	TeamJoin = EventsAPIType("team_join")
	// TeamRename The team was renamed
	TeamRename = EventsAPIType("team_rename")
	// TokensRevoked APP's API tokes are revoked
	TokensRevoked = EventsAPIType("tokens_revoked")
	// UserChange A user object has changed
	UserChange = EventsAPIType("user_change")
	// UserHuddleChanged is an event when a user's huddle status changes
	UserHuddleChanged = EventsAPIType("user_huddle_changed")
	// UserProfileChanged is sent if a user's profile information has changed.
	UserProfileChanged = EventsAPIType("user_profile_changed")
	// UserStatusChanged is an event when a user's status changes
	UserStatusChanged = EventsAPIType("user_status_changed")
	// WorkflowStepExecute Happens, if a workflow step of your app is invoked
	WorkflowStepExecute = EventsAPIType("workflow_step_execute")
)

// EventsAPIInnerEventMapping maps INNER Event API events to their corresponding struct
// implementations. The structs should be instances of the unmarshalling
// target for the matching event type.
var EventsAPIInnerEventMapping = map[EventsAPIType]interface{}{
	AppDeleted:                    AppDeletedEvent{},
	AppHomeOpened:                 AppHomeOpenedEvent{},
	AppInstalled:                  AppInstalledEvent{},
	AppMention:                    AppMentionEvent{},
	AppRequested:                  AppRequestedEvent{},
	AppUninstalled:                AppUninstalledEvent{},
	AppUninstalledTeam:            AppUninstalledTeamEvent{},
	AssistantThreadContextChanged: AssistantThreadContextChangedEvent{},
	AssistantThreadStarted:        AssistantThreadStartedEvent{},
	CallRejected:                  CallRejectedEvent{},
	ChannelArchive:                ChannelArchiveEvent{},
	ChannelCreated:                ChannelCreatedEvent{},
	ChannelDeleted:                ChannelDeletedEvent{},
	ChannelHistoryChanged:         ChannelHistoryChangedEvent{},
	ChannelIDChanged:              ChannelIDChangedEvent{},
	ChannelLeft:                   ChannelLeftEvent{},
	ChannelRename:                 ChannelRenameEvent{},
	ChannelShared:                 ChannelSharedEvent{},
	ChannelUnarchive:              ChannelUnarchiveEvent{},
	ChannelUnshared:               ChannelUnsharedEvent{},
	CommandsChanged:               CommandsChangedEvent{},
	DndUpdated:                    DndUpdatedEvent{},
	DndUpdatedUser:                DndUpdatedUserEvent{},
	EmailDomainChanged:            EmailDomainChangedEvent{},
	EmojiChanged:                  EmojiChangedEvent{},
	FileChange:                    FileChangeEvent{},
	FileCreated:                   FileCreatedEvent{},
	FileDeleted:                   FileDeletedEvent{},
	FilePublic:                    FilePublicEvent{},
	FileShared:                    FileSharedEvent{},
	FileUnshared:                  FileUnsharedEvent{},
	FunctionExecuted:              FunctionExecutedEvent{},
	GridMigrationFinished:         GridMigrationFinishedEvent{},
	GridMigrationStarted:          GridMigrationStartedEvent{},
	GroupArchive:                  GroupArchiveEvent{},
	GroupClose:                    GroupCloseEvent{},
	GroupDeleted:                  GroupDeletedEvent{},
	GroupHistoryChanged:           GroupHistoryChangedEvent{},
	GroupLeft:                     GroupLeftEvent{},
	GroupOpen:                     GroupOpenEvent{},
	GroupRename:                   GroupRenameEvent{},
	GroupUnarchive:                GroupUnarchiveEvent{},
	ImClose:                       ImCloseEvent{},
	ImCreated:                     ImCreatedEvent{},
	ImHistoryChanged:              ImHistoryChangedEvent{},
	ImOpen:                        ImOpenEvent{},
	InviteRequested:               InviteRequestedEvent{},
	LinkShared:                    LinkSharedEvent{},
	MemberJoinedChannel:           MemberJoinedChannelEvent{},
	MemberLeftChannel:             MemberLeftChannelEvent{},
	Message:                       MessageEvent{},
	MessageMetadataDeleted:        MessageMetadataDeletedEvent{},
	MessageMetadataPosted:         MessageMetadataPostedEvent{},
	MessageMetadataUpdated:        MessageMetadataUpdatedEvent{},
	PinAdded:                      PinAddedEvent{},
	PinRemoved:                    PinRemovedEvent{},
	ReactionAdded:                 ReactionAddedEvent{},
	ReactionRemoved:               ReactionRemovedEvent{},
	SharedChannelInviteAccepted:   SharedChannelInviteAcceptedEvent{},
	SharedChannelInviteApproved:   SharedChannelInviteApprovedEvent{},
	SharedChannelInviteDeclined:   SharedChannelInviteDeclinedEvent{},
	SharedChannelInviteReceived:   SharedChannelInviteReceivedEvent{},
	SharedChannelInviteRequested:  SharedChannelInviteRequestedEvent{},
	StarAdded:                     StarAddedEvent{},
	StarRemoved:                   StarRemovedEvent{},
	SubteamCreated:                SubteamCreatedEvent{},
	SubteamMembersChanged:         SubteamMembersChangedEvent{},
	SubteamSelfAdded:              SubteamSelfAddedEvent{},
	SubteamSelfRemoved:            SubteamSelfRemovedEvent{},
	SubteamUpdated:                SubteamUpdatedEvent{},
	TeamAccessGranted:             TeamAccessGrantedEvent{},
	TeamAccessRevoked:             TeamAccessRevokedEvent{},
	TeamDomainChange:              TeamDomainChangeEvent{},
	TeamJoin:                      TeamJoinEvent{},
	TeamRename:                    TeamRenameEvent{},
	TokensRevoked:                 TokensRevokedEvent{},
	UserChange:                    UserChangeEvent{},
	UserHuddleChanged:             UserHuddleChangedEvent{},
	UserProfileChanged:            UserProfileChangedEvent{},
	UserStatusChanged:             UserStatusChangedEvent{},
}
