//go:generate mockgen -source=$GOFILE -destination=../linebotmock/mock_$GOFILE -package=mock_$GOPACKAGE
package lime

import (
	"io"

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
}
