//go:generate mockgen -source=$GOFILE -destination=../linebotmock/mock_$GOFILE -package=mock_$GOPACKAGE
package lime

import (
	"io"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBotClient interface {
	GetNumberMessagesDelivery(date string) *linebot.GetNumberMessagesDeliveryCall
	GetNumberFollowers(date string) *linebot.GetNumberFollowersCall
	GetFriendDemographics() *linebot.GetFriendDemographicsCall
	GetUserInteractionStats(requestID string) *linebot.GetUserInteractionStatsCall
	GetGroupSummary(groupID string) *linebot.GetGroupSummaryCall
	GetMessageQuota() *linebot.GetMessageQuotaCall
	GetMessageQuotaConsumption() *linebot.GetMessageQuotaCall
	GetMessageConsumption() *linebot.GetMessageConsumptionCall
	LeaveGroup(groupID string) *linebot.LeaveGroupCall
	LeaveRoom(roomID string) *linebot.LeaveRoomCall
	NewRawCall(method string, endpoint string) (*linebot.RawCall, error)
	NewRawCallWithBody(method string, endpoint string, body io.Reader) (*linebot.RawCall, error)
	GetLIFF() *linebot.GetLIFFAllCall
	AddLIFF(app linebot.LIFFApp) *linebot.AddLIFFCall
	UpdateLIFF(liffID string, app linebot.LIFFApp) *linebot.UpdateLIFFCall
	DeleteLIFF(liffID string) *linebot.DeleteLIFFCall
	GetBotInfo() *linebot.GetBotInfoCall
	UploadAudienceGroup(description string, options ...linebot.IUploadAudienceGroupOption) *linebot.UploadAudienceGroupCall
	UploadAudienceGroupByFile(description string, audiences []string, options ...linebot.IUploadAudienceGroupByFileOption) *linebot.UploadAudienceGroupByFileCall
	AddAudiences(audienceGroupID int, audiences []string, options ...linebot.IAddAudiencesOption) *linebot.AddAudiencesCall
	AddAudiencesByFile(audienceGroupID int, audiences []string, options ...linebot.IAddAudiencesByFileOption) *linebot.AddAudiencesByFileCall
	ClickAudienceGroup(description, requestID string, options ...linebot.IClickAudienceGroupOption) *linebot.ClickAudienceGroupCall
	IMPAudienceGroup(description, requestID string) *linebot.IMPAudienceGroupCall
	UpdateAudienceGroupDescription(audienceGroupID int, description string) *linebot.UpdateAudienceGroupDescriptionCall
	ActivateAudienceGroup(audienceGroupID int) *linebot.ActivateAudienceGroupCall
	DeleteAudienceGroup(audienceGroupID int) *linebot.DeleteAudienceGroupCall
	GetAudienceGroup(audienceGroupID int) *linebot.GetAudienceGroupCall
	ListAudienceGroup(page int, options ...linebot.IListAudienceGroupOption) *linebot.ListAudienceGroupCall
	GetAudienceGroupAuthorityLevel() *linebot.GetAudienceGroupAuthorityLevelCall
	ChangeAudienceGroupAuthorityLevel(authorityLevel linebot.AudienceAuthorityLevelType) *linebot.ChangeAudienceGroupAuthorityLevelCall
	GetFollowerIDs(continuationToken string) *linebot.GetFollowerIDsCall
	GetGroupMemberCount(groupID string) *linebot.GetGroupMemberCountCall
	GetRoomMemberCount(roomID string) *linebot.GetRoomMemberCountCall
	PushMessage(to string, messages ...linebot.SendingMessage) *linebot.PushMessageCall
	ReplyMessage(replyToken string, messages ...linebot.SendingMessage) *linebot.ReplyMessageCall
	Multicast(to []string, messages ...linebot.SendingMessage) *linebot.MulticastCall
	BroadcastMessage(messages ...linebot.SendingMessage) *linebot.BroadcastMessageCall
	Narrowcast(messages ...linebot.SendingMessage) *linebot.NarrowcastCall
	ValidatePushMessage(messages ...linebot.SendingMessage) *linebot.ValidatePushMessageCall
	ValidateReplyMessage(messages ...linebot.SendingMessage) *linebot.ValidateReplyMessageCall
	ValidateMulticastMessage(messages ...linebot.SendingMessage) *linebot.ValidateMulticastMessageCall
	ValidateBroadcastMessage(messages ...linebot.SendingMessage) *linebot.ValidateBroadcastMessageCall
	ValidateNarrowcastMessage(messages ...linebot.SendingMessage) *linebot.ValidateNarrowcastMessageCall
	IssueAccessTokenV2(clientAssertion string) *linebot.IssueAccessTokenV2Call
	GetAccessTokensV2(clientAssertion string) *linebot.GetAccessTokensV2Call
	RevokeAccessTokenV2(channelID, channelSecret, accessToken string) *linebot.RevokeAccessTokenV2Call
	IssueAccessToken(channelID, channelSecret string) *linebot.IssueAccessTokenCall
	RevokeAccessToken(accessToken string) *linebot.RevokeAccessTokenCall
	VerifyAccessToken(accessToken string) *linebot.VerifyAccessTokenCall
	GetGroupMemberIDs(groupID, continuationToken string) *linebot.GetGroupMemberIDsCall
	GetRoomMemberIDs(roomID, continuationToken string) *linebot.GetRoomMemberIDsCall
	IssueLinkToken(userID string) *linebot.IssueLinkTokenCall
	GetNumberReplyMessages(date string) *linebot.GetNumberMessagesCall
	GetNumberPushMessages(date string) *linebot.GetNumberMessagesCall
	GetNumberMulticastMessages(date string) *linebot.GetNumberMessagesCall
	GetNumberBroadcastMessages(date string) *linebot.GetNumberMessagesCall
	GetProgressNarrowcastMessages(requestID string) *linebot.GetProgressMessagesCall
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	GetWebhookInfo() *linebot.GetWebhookInfo
	SetWebhookEndpointURL(webhookEndpoint string) *linebot.SetWebhookEndpointURLCall
	TestWebhook() *linebot.TestWebhook
	GetProfile(userID string) *linebot.GetProfileCall
	GetGroupMemberProfile(groupID, userID string) *linebot.GetGroupMemberProfileCall
	GetRoomMemberProfile(roomID, userID string) *linebot.GetRoomMemberProfileCall
	GetRichMenu(richMenuID string) *linebot.GetRichMenuCall
	GetUserRichMenu(userID string) *linebot.GetUserRichMenuCall
	CreateRichMenu(richMenu linebot.RichMenu) *linebot.CreateRichMenuCall
	DeleteRichMenu(richMenuID string) *linebot.DeleteRichMenuCall
	LinkUserRichMenu(userID, richMenuID string) *linebot.LinkUserRichMenuCall
	UnlinkUserRichMenu(userID string) *linebot.UnlinkUserRichMenuCall
	SetDefaultRichMenu(richMenuID string) *linebot.SetDefaultRichMenuCall
	CancelDefaultRichMenu() *linebot.CancelDefaultRichMenuCall
	GetDefaultRichMenu() *linebot.GetDefaultRichMenuCall
	GetRichMenuList() *linebot.GetRichMenuListCall
	DownloadRichMenuImage(richMenuID string) *linebot.DownloadRichMenuImageCall
	UploadRichMenuImage(richMenuID, imgPath string) *linebot.UploadRichMenuImageCall
	BulkLinkRichMenu(richMenuID string, userIDs ...string) *linebot.BulkLinkRichMenuCall
	BulkUnlinkRichMenu(userIDs ...string) *linebot.BulkUnlinkRichMenuCall
	CreateRichMenuAlias(richMenuAliasID, richMenuID string) *linebot.CreateRichMenuAliasCall
	UpdateRichMenuAlias(richMenuAliasID, richMenuID string) *linebot.UpdateRichMenuAliasCall
	DeleteRichMenuAlias(richMenuAliasID string) *linebot.DeleteRichMenuAliasCall
	GetRichMenuAlias(richMenuAliasID string) *linebot.GetRichMenuAliasCall
	GetRichMenuAliasList() *linebot.GetRichMenuAliasListCall
	ValidateRichMenuObject(richMenu linebot.RichMenu) *linebot.ValidateRichMenuObjectCall
	GetMessageContent(messageID string) *linebot.GetMessageContentCall
}
