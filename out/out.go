// Source generated for release WIN63-202312011715-920410468 (source: Sulek API)

package out

import g "github.com/b7c/goearth"

func id(name string) g.Identifier {
	return g.Identifier{Dir: g.Out, Name: name}
}

var (
	Game2ExitGame = id("Game2ExitGame")
	Game2GameChat = id("Game2GameChat")
	Game2LoadStageReady = id("Game2LoadStageReady")
	Game2PlayAgain = id("Game2PlayAgain")
	CommunityGoalVote = id("CommunityGoalVote")
	AcceptTrading = id("AcceptTrading")
	AddItemsToTrade = id("AddItemsToTrade")
	AddItemToTrade = id("AddItemToTrade")
	CloseTrading = id("CloseTrading")
	ConfirmAcceptTrading = id("ConfirmAcceptTrading")
	ConfirmDeclineTrading = id("ConfirmDeclineTrading")
	OpenTrading = id("OpenTrading")
	RemoveItemFromTrade = id("RemoveItemFromTrade")
	UnacceptTrading = id("UnacceptTrading")
	Craft = id("Craft")
	CraftSecret = id("CraftSecret")
	GetCraftableProducts = id("GetCraftableProducts")
	GetCraftingRecipe = id("GetCraftingRecipe")
	GetCraftingRecipesAvailable = id("GetCraftingRecipesAvailable")
	Game2GetFriendsLeaderboard = id("Game2GetFriendsLeaderboard")
	Game2GetTotalGroupLeaderboard = id("Game2GetTotalGroupLeaderboard")
	Game2GetTotalLeaderboard = id("Game2GetTotalLeaderboard")
	Game2GetWeeklyFriendsLeaderboard = id("Game2GetWeeklyFriendsLeaderboard")
	Game2GetWeeklyGroupLeaderboard = id("Game2GetWeeklyGroupLeaderboard")
	Game2GetWeeklyLeaderboard = id("Game2GetWeeklyLeaderboard")
	CommandBot = id("CommandBot")
	GetBotCommandConfigurationData = id("GetBotCommandConfigurationData")
	MysteryBoxWaitingCanceled = id("MysteryBoxWaitingCanceled")
	RoomNetworkOpenConnection = id("RoomNetworkOpenConnection")
	GetTalentTrackLevel = id("GetTalentTrackLevel")
	GetTalentTrack = id("GetTalentTrack")
	GuideAdvertisementRead = id("GuideAdvertisementRead")
	PeerUsersClassification = id("PeerUsersClassification")
	RoomUsersClassification = id("RoomUsersClassification")
	ChangeQueue = id("ChangeQueue")
	OpenFlatConnection = id("OpenFlatConnection")
	Quit = id("Quit")
	AmbassadorAlert = id("AmbassadorAlert")
	AssignRights = id("AssignRights")
	BanUserWithDuration = id("BanUserWithDuration")
	KickUser = id("KickUser")
	LetUserIn = id("LetUserIn")
	MuteAllInRoom = id("MuteAllInRoom")
	MuteUser = id("MuteUser")
	RemoveAllRights = id("RemoveAllRights")
	RemoveRights = id("RemoveRights")
	UnbanUserFromRoom = id("UnbanUserFromRoom")
	UnmuteUser = id("UnmuteUser")
	AddFavouriteRoom = id("AddFavouriteRoom")
	CancelEvent = id("CancelEvent")
	CanCreateRoom = id("CanCreateRoom")
	CompetitionRoomsSearch = id("CompetitionRoomsSearch")
	ConvertGlobalRoomId = id("ConvertGlobalRoomId")
	CreateFlat = id("CreateFlat")
	DeleteFavouriteRoom = id("DeleteFavouriteRoom")
	EditEvent = id("EditEvent")
	ForwardToARandomPromotedRoom = id("ForwardToARandomPromotedRoom")
	ForwardToSomeRoom = id("ForwardToSomeRoom")
	GetGuestRoom = id("GetGuestRoom")
	GetOfficialRooms = id("GetOfficialRooms")
	GetPopularRoomTags = id("GetPopularRoomTags")
	GetUserEventCats = id("GetUserEventCats")
	GetUserFlatCats = id("GetUserFlatCats")
	GuildBaseSearch = id("GuildBaseSearch")
	MyFavouriteRoomsSearch = id("MyFavouriteRoomsSearch")
	MyFrequentRoomHistorySearch = id("MyFrequentRoomHistorySearch")
	MyFriendsRoomsSearch = id("MyFriendsRoomsSearch")
	MyGuildBasesSearch = id("MyGuildBasesSearch")
	MyRecommendedRooms = id("MyRecommendedRooms")
	MyRoomHistorySearch = id("MyRoomHistorySearch")
	MyRoomRightsSearch = id("MyRoomRightsSearch")
	MyRoomsSearch = id("MyRoomsSearch")
	PopularRoomsSearch = id("PopularRoomsSearch")
	RateFlat = id("RateFlat")
	RemoveOwnRoomRightsRoom = id("RemoveOwnRoomRightsRoom")
	RoomAdEventTabAdClicked = id("RoomAdEventTabAdClicked")
	RoomAdEventTabViewed = id("RoomAdEventTabViewed")
	RoomAdSearch = id("RoomAdSearch")
	RoomsWhereMyFriendsAreSearch = id("RoomsWhereMyFriendsAreSearch")
	RoomsWithHighestScoreSearch = id("RoomsWithHighestScoreSearch")
	RoomTextSearch = id("RoomTextSearch")
	SetRoomSessionTags = id("SetRoomSessionTags")
	ToggleStaffPick = id("ToggleStaffPick")
	UpdateHomeRoom = id("UpdateHomeRoom")
	Game2CheckGameDirectoryStatus = id("Game2CheckGameDirectoryStatus")
	Game2GetAccountGameStatus = id("Game2GetAccountGameStatus")
	Game2LeaveGame = id("Game2LeaveGame")
	Game2QuickJoinGame = id("Game2QuickJoinGame")
	Game2StartSnowWar = id("Game2StartSnowWar")
	ResetUnseenItemIds = id("ResetUnseenItemIds")
	ResetUnseenItems = id("ResetUnseenItems")
	RequestFurniInventory = id("RequestFurniInventory")
	RequestFurniInventoryWhenNotInRoom = id("RequestFurniInventoryWhenNotInRoom")
	RequestRoomPropertySet = id("RequestRoomPropertySet")
	CancelTyping = id("CancelTyping")
	Chat = id("Chat")
	Shout = id("Shout")
	StartTyping = id("StartTyping")
	Whisper = id("Whisper")
	BuildersClubPlaceRoomItem = id("BuildersClubPlaceRoomItem")
	BuildersClubPlaceWallItem = id("BuildersClubPlaceWallItem")
	BuildersClubQueryFurniCount = id("BuildersClubQueryFurniCount")
	GetBonusRareInfo = id("GetBonusRareInfo")
	GetBundleDiscountRuleset = id("GetBundleDiscountRuleset")
	GetCatalogIndex = id("GetCatalogIndex")
	GetCatalogPage = id("GetCatalogPage")
	GetCatalogPageWithEarliestExpiry = id("GetCatalogPageWithEarliestExpiry")
	GetClubGift = id("GetClubGift")
	GetClubOffers = id("GetClubOffers")
	GetGiftWrappingConfiguration = id("GetGiftWrappingConfiguration")
	GetHabboClubExtendOffer = id("GetHabboClubExtendOffer")
	GetIsOfferGiftable = id("GetIsOfferGiftable")
	GetLimitedOfferAppearingNext = id("GetLimitedOfferAppearingNext")
	GetNextTargetedOffer = id("GetNextTargetedOffer")
	GetProductOffer = id("GetProductOffer")
	GetRoomAdPurchaseInfo = id("GetRoomAdPurchaseInfo")
	GetSeasonalCalendarDaily = id("GetSeasonalCalendarDaily")
	GetSellablePetPalettes = id("GetSellablePetPalettes")
	GetSnowWarGameTokensOffer = id("GetSnowWarGameTokensOffer")
	MarkCatalogNewAdditionsPageOpened = id("MarkCatalogNewAdditionsPageOpened")
	PurchaseBasicMembershipExtension = id("PurchaseBasicMembershipExtension")
	PurchaseFromCatalogAsGift = id("PurchaseFromCatalogAsGift")
	PurchaseFromCatalog = id("PurchaseFromCatalog")
	PurchaseRoomAd = id("PurchaseRoomAd")
	PurchaseSnowWarGameTokensOffer = id("PurchaseSnowWarGameTokensOffer")
	PurchaseTargetedOffer = id("PurchaseTargetedOffer")
	PurchaseVipMembershipExtension = id("PurchaseVipMembershipExtension")
	RedeemVoucher = id("RedeemVoucher")
	RoomAdPurchaseInitiated = id("RoomAdPurchaseInitiated")
	SelectClubGift = id("SelectClubGift")
	SetTargetedOfferState = id("SetTargetedOfferState")
	ShopTargetedOfferViewed = id("ShopTargetedOfferViewed")
	GetSelectedNftWardrobeOutfit = id("GetSelectedNftWardrobeOutfit")
	GetUserNftWardrobe = id("GetUserNftWardrobe")
	SaveUserNftWardrobe = id("SaveUserNftWardrobe")
	AddSpamWallPostIt = id("AddSpamWallPostIt")
	ControlYoutubeDisplayPlayback = id("ControlYoutubeDisplayPlayback")
	CreditFurniRedeem = id("CreditFurniRedeem")
	DiceOff = id("DiceOff")
	EnterOneWayDoor = id("EnterOneWayDoor")
	ExtendRentOrBuyoutFurni = id("ExtendRentOrBuyoutFurni")
	ExtendRentOrBuyoutStripItem = id("ExtendRentOrBuyoutStripItem")
	GetGuildFurniContextMenuInfo = id("GetGuildFurniContextMenuInfo")
	GetRentOrBuyoutOffer = id("GetRentOrBuyoutOffer")
	GetYoutubeDisplayStatus = id("GetYoutubeDisplayStatus")
	OpenMysteryTrophy = id("OpenMysteryTrophy")
	OpenPetPackage = id("OpenPetPackage")
	PlacePostIt = id("PlacePostIt")
	PresentOpen = id("PresentOpen")
	RentableSpaceCancelRent = id("RentableSpaceCancelRent")
	RentableSpaceRent = id("RentableSpaceRent")
	RentableSpaceStatus = id("RentableSpaceStatus")
	RoomDimmerChangeState = id("RoomDimmerChangeState")
	RoomDimmerGetPresets = id("RoomDimmerGetPresets")
	RoomDimmerSavePreset = id("RoomDimmerSavePreset")
	SetCustomStackingHeight = id("SetCustomStackingHeight")
	SetMannequinFigure = id("SetMannequinFigure")
	SetMannequinName = id("SetMannequinName")
	SetRandomState = id("SetRandomState")
	SetRoomBackgroundColorData = id("SetRoomBackgroundColorData")
	SetYoutubeDisplayPlaylist = id("SetYoutubeDisplayPlaylist")
	SpinWheelOfFortune = id("SpinWheelOfFortune")
	ThrowDice = id("ThrowDice")
	GetHotLooks = id("GetHotLooks")
	SetChatPreferences = id("SetChatPreferences")
	SetChatStylePreference = id("SetChatStylePreference")
	SetIgnoreRoomInvites = id("SetIgnoreRoomInvites")
	SetNewNavigatorWindowPreferences = id("SetNewNavigatorWindowPreferences")
	SetRoomCameraPreferences = id("SetRoomCameraPreferences")
	SetSoundSettings = id("SetSoundSettings")
	SetUIFlags = id("SetUIFlags")
	PollAnswer = id("PollAnswer")
	PollReject = id("PollReject")
	PollStart = id("PollStart")
	GetCreditsInfo = id("GetCreditsInfo")
	AddAdminRightsToMember = id("AddAdminRightsToMember")
	ApproveAllMembershipRequests = id("ApproveAllMembershipRequests")
	ApproveMembershipRequest = id("ApproveMembershipRequest")
	ApproveName = id("ApproveName")
	ChangeEmail = id("ChangeEmail")
	CreateGuild = id("CreateGuild")
	DeactivateGuild = id("DeactivateGuild")
	DeselectFavouriteHabboGroup = id("DeselectFavouriteHabboGroup")
	GetEmailStatus = id("GetEmailStatus")
	GetExtendedProfileByName = id("GetExtendedProfileByName")
	GetExtendedProfile = id("GetExtendedProfile")
	GetGuildCreationInfo = id("GetGuildCreationInfo")
	GetGuildEditInfo = id("GetGuildEditInfo")
	GetGuildEditorData = id("GetGuildEditorData")
	GetGuildMemberships = id("GetGuildMemberships")
	GetGuildMembers = id("GetGuildMembers")
	GetHabboGroupBadges = id("GetHabboGroupBadges")
	GetHabboGroupDetails = id("GetHabboGroupDetails")
	GetIgnoredUsers = id("GetIgnoredUsers")
	GetMemberGuildItemCount = id("GetMemberGuildItemCount")
	GetMOTD = id("GetMOTD")
	GetRelationshipStatusInfo = id("GetRelationshipStatusInfo")
	GetSelectedBadges = id("GetSelectedBadges")
	GiveStarGemToUser = id("GiveStarGemToUser")
	IgnoreUserId = id("IgnoreUserId")
	IgnoreUser = id("IgnoreUser")
	JoinHabboGroup = id("JoinHabboGroup")
	KickMember = id("KickMember")
	RejectMembershipRequest = id("RejectMembershipRequest")
	RemoveAdminRightsFromMember = id("RemoveAdminRightsFromMember")
	ScrGetKickbackInfo = id("ScrGetKickbackInfo")
	ScrGetUserInfo = id("ScrGetUserInfo")
	SelectFavouriteHabboGroup = id("SelectFavouriteHabboGroup")
	UnblockGroupMember = id("UnblockGroupMember")
	UnignoreUser = id("UnignoreUser")
	UpdateGuildBadge = id("UpdateGuildBadge")
	UpdateGuildColors = id("UpdateGuildColors")
	UpdateGuildIdentity = id("UpdateGuildIdentity")
	UpdateGuildSettings = id("UpdateGuildSettings")
	EventLog = id("EventLog")
	LagWarningReport = id("LagWarningReport")
	LatencyPingReport = id("LatencyPingReport")
	LatencyPingRequest = id("LatencyPingRequest")
	PerformanceLog = id("PerformanceLog")
	AcceptQuest = id("AcceptQuest")
	ActivateQuest = id("ActivateQuest")
	CancelQuest = id("CancelQuest")
	FriendRequestQuestComplete = id("FriendRequestQuestComplete")
	GetCommunityGoalHallOfFame = id("GetCommunityGoalHallOfFame")
	GetCommunityGoalProgress = id("GetCommunityGoalProgress")
	GetConcurrentUsersGoalProgress = id("GetConcurrentUsersGoalProgress")
	GetConcurrentUsersReward = id("GetConcurrentUsersReward")
	GetDailyQuest = id("GetDailyQuest")
	GetQuests = id("GetQuests")
	GetSeasonalQuestsOnly = id("GetSeasonalQuestsOnly")
	OpenQuestTracker = id("OpenQuestTracker")
	RejectQuest = id("RejectQuest")
	StartCampaign = id("StartCampaign")
	GetForumsList = id("GetForumsList")
	GetForumStats = id("GetForumStats")
	GetMessages = id("GetMessages")
	GetThread = id("GetThread")
	GetThreads = id("GetThreads")
	GetUnreadForumsCount = id("GetUnreadForumsCount")
	ModerateMessage = id("ModerateMessage")
	ModerateThread = id("ModerateThread")
	PostMessage = id("PostMessage")
	UpdateForumReadMarker = id("UpdateForumReadMarker")
	UpdateForumSettings = id("UpdateForumSettings")
	UpdateThread = id("UpdateThread")
	CreditVaultStatus = id("CreditVaultStatus")
	IncomeRewardClaim = id("IncomeRewardClaim")
	IncomeRewardStatus = id("IncomeRewardStatus")
	WithdrawCreditVault = id("WithdrawCreditVault")
	AddJukeboxDisk = id("AddJukeboxDisk")
	GetJukeboxPlayList = id("GetJukeboxPlayList")
	GetNowPlaying = id("GetNowPlaying")
	GetOfficialSongId = id("GetOfficialSongId")
	GetSongInfo = id("GetSongInfo")
	GetSoundMachinePlayList = id("GetSoundMachinePlayList")
	GetSoundSettings = id("GetSoundSettings")
	GetUserSongDisks = id("GetUserSongDisks")
	RemoveJukeboxDisk = id("RemoveJukeboxDisk")
	GetInterstitial = id("GetInterstitial")
	InterstitialShown = id("InterstitialShown")
	CloseIssueDefaultAction = id("CloseIssueDefaultAction")
	CloseIssues = id("CloseIssues")
	DefaultSanction = id("DefaultSanction")
	GetCfhChatlog = id("GetCfhChatlog")
	GetModeratorRoomInfo = id("GetModeratorRoomInfo")
	GetModeratorUserInfo = id("GetModeratorUserInfo")
	GetRoomChatlog = id("GetRoomChatlog")
	GetRoomVisits = id("GetRoomVisits")
	GetUserChatlog = id("GetUserChatlog")
	ModAlert = id("ModAlert")
	ModBan = id("ModBan")
	ModerateRoom = id("ModerateRoom")
	ModeratorAction = id("ModeratorAction")
	ModKick = id("ModKick")
	ModMessage = id("ModMessage")
	ModMute = id("ModMute")
	ModToolPreferences = id("ModToolPreferences")
	ModToolSanction = id("ModToolSanction")
	ModTradingLock = id("ModTradingLock")
	PickIssues = id("PickIssues")
	ReleaseIssues = id("ReleaseIssues")
	AvatarEffectActivated = id("AvatarEffectActivated")
	AvatarEffectSelected = id("AvatarEffectSelected")
	CancelPetBreeding = id("CancelPetBreeding")
	ConfirmPetBreeding = id("ConfirmPetBreeding")
	GetPetInventory = id("GetPetInventory")
	ApplySnapshot = id("ApplySnapshot")
	Open = id("Open")
	UpdateAction = id("UpdateAction")
	UpdateAddon = id("UpdateAddon")
	UpdateCondition = id("UpdateCondition")
	UpdateSelector = id("UpdateSelector")
	UpdateTrigger = id("UpdateTrigger")
	AvatarExpression = id("AvatarExpression")
	ChangeMotto = id("ChangeMotto")
	ChangePosture = id("ChangePosture")
	CustomizeAvatarWithFurni = id("CustomizeAvatarWithFurni")
	Dance = id("Dance")
	DropCarryItem = id("DropCarryItem")
	LookTo = id("LookTo")
	PassCarryItem = id("PassCarryItem")
	PassCarryItemToPet = id("PassCarryItemToPet")
	Sign = id("Sign")
	GetBadgePointLimits = id("GetBadgePointLimits")
	GetBadges = id("GetBadges")
	GetIsBadgeRequestFulfilled = id("GetIsBadgeRequestFulfilled")
	RequestABadge = id("RequestABadge")
	SetActivatedBadges = id("SetActivatedBadges")
	BreedPets = id("BreedPets")
	CustomizePetWithFurni = id("CustomizePetWithFurni")
	GetPetInfo = id("GetPetInfo")
	PetSelected = id("PetSelected")
	RespectPet = id("RespectPet")
	ChangeUserNameInRoom = id("ChangeUserNameInRoom")
	ChangeUserName = id("ChangeUserName")
	CheckUserName = id("CheckUserName")
	GetWardrobe = id("GetWardrobe")
	SaveWardrobeOutfit = id("SaveWardrobeOutfit")
	GetResolutionAchievements = id("GetResolutionAchievements")
	GetAchievements = id("GetAchievements")
	OpenCampaignCalendarDoorAsStaff = id("OpenCampaignCalendarDoorAsStaff")
	OpenCampaignCalendarDoor = id("OpenCampaignCalendarDoor")
	FriendFurniConfirmLock = id("FriendFurniConfirmLock")
	BuyMarketplaceOffer = id("BuyMarketplaceOffer")
	BuyMarketplaceTokens = id("BuyMarketplaceTokens")
	CancelMarketplaceOffer = id("CancelMarketplaceOffer")
	GetMarketplaceCanMakeOffer = id("GetMarketplaceCanMakeOffer")
	GetMarketplaceConfiguration = id("GetMarketplaceConfiguration")
	GetMarketplaceItemStats = id("GetMarketplaceItemStats")
	GetMarketplaceOffers = id("GetMarketplaceOffers")
	GetMarketplaceOwnOffers = id("GetMarketplaceOwnOffers")
	MakeOffer = id("MakeOffer")
	RedeemMarketplaceOfferCredits = id("RedeemMarketplaceOfferCredits")
	DeleteRoom = id("DeleteRoom")
	GetBannedUsersFromRoom = id("GetBannedUsersFromRoom")
	GetCustomRoomFilter = id("GetCustomRoomFilter")
	GetFlatControllers = id("GetFlatControllers")
	GetRoomSettings = id("GetRoomSettings")
	SaveRoomSettings = id("SaveRoomSettings")
	UpdateRoomCategoryAndTradeSettings = id("UpdateRoomCategoryAndTradeSettings")
	UpdateRoomFilter = id("UpdateRoomFilter")
	ForwardToACompetitionRoom = id("ForwardToACompetitionRoom")
	ForwardToASubmittableRoom = id("ForwardToASubmittableRoom")
	ForwardToRandomCompetitionRoom = id("ForwardToRandomCompetitionRoom")
	GetCurrentTimingCode = id("GetCurrentTimingCode")
	GetIsUserPartOfCompetition = id("GetIsUserPartOfCompetition")
	GetSecondsUntil = id("GetSecondsUntil")
	RoomCompetitionInit = id("RoomCompetitionInit")
	SubmitRoomToCompetition = id("SubmitRoomToCompetition")
	VoteForRoom = id("VoteForRoom")
	CallForHelpFromForumMessage = id("CallForHelpFromForumMessage")
	CallForHelpFromForumThread = id("CallForHelpFromForumThread")
	CallForHelpFromIM = id("CallForHelpFromIM")
	CallForHelpFromPhoto = id("CallForHelpFromPhoto")
	CallForHelpFromSelfie = id("CallForHelpFromSelfie")
	CallForHelp = id("CallForHelp")
	ChatReviewGuideDecidesOnOffer = id("ChatReviewGuideDecidesOnOffer")
	ChatReviewGuideDetached = id("ChatReviewGuideDetached")
	ChatReviewGuideVote = id("ChatReviewGuideVote")
	ChatReviewSessionCreate = id("ChatReviewSessionCreate")
	DeletePendingCallsForHelp = id("DeletePendingCallsForHelp")
	GetCfhStatus = id("GetCfhStatus")
	GetGuideReportingStatus = id("GetGuideReportingStatus")
	GetPendingCallsForHelp = id("GetPendingCallsForHelp")
	GetQuizQuestions = id("GetQuizQuestions")
	GuideSessionCreate = id("GuideSessionCreate")
	GuideSessionFeedback = id("GuideSessionFeedback")
	GuideSessionGetRequesterRoom = id("GuideSessionGetRequesterRoom")
	GuideSessionGuideDecides = id("GuideSessionGuideDecides")
	GuideSessionInviteRequester = id("GuideSessionInviteRequester")
	GuideSessionIsTyping = id("GuideSessionIsTyping")
	GuideSessionMessage = id("GuideSessionMessage")
	GuideSessionOnDutyUpdate = id("GuideSessionOnDutyUpdate")
	GuideSessionReport = id("GuideSessionReport")
	GuideSessionRequesterCancels = id("GuideSessionRequesterCancels")
	GuideSessionResolved = id("GuideSessionResolved")
	PostQuizAnswers = id("PostQuizAnswers")
	NavigatorAddCollapsedCategory = id("NavigatorAddCollapsedCategory")
	NavigatorAddSavedSearch = id("NavigatorAddSavedSearch")
	NavigatorDeleteSavedSearch = id("NavigatorDeleteSavedSearch")
	NavigatorRemoveCollapsedCategory = id("NavigatorRemoveCollapsedCategory")
	NavigatorSetSearchCodeViewMode = id("NavigatorSetSearchCodeViewMode")
	NewNavigatorInit = id("NewNavigatorInit")
	NewNavigatorSearch = id("NewNavigatorSearch")
	UpdateFigureData = id("UpdateFigureData")
	ClientHello = id("ClientHello")
	CompleteDiffieHandshake = id("CompleteDiffieHandshake")
	Disconnect = id("Disconnect")
	InfoRetrieve = id("InfoRetrieve")
	InitDiffieHandshake = id("InitDiffieHandshake")
	Pong = id("Pong")
	SSOTicket = id("SSOTicket")
	UniqueID = id("UniqueID")
	VersionCheck = id("VersionCheck")
	AcceptFriend = id("AcceptFriend")
	DeclineFriend = id("DeclineFriend")
	FindNewFriends = id("FindNewFriends")
	FollowFriend = id("FollowFriend")
	FriendListUpdate = id("FriendListUpdate")
	GetFriendRequests = id("GetFriendRequests")
	GetMessengerHistory = id("GetMessengerHistory")
	HabboSearch = id("HabboSearch")
	MessengerInit = id("MessengerInit")
	RemoveFriend = id("RemoveFriend")
	RequestFriend = id("RequestFriend")
	SendMsg = id("SendMsg")
	SendRoomInvite = id("SendRoomInvite")
	SetRelationshipStatus = id("SetRelationshipStatus")
	VisitUser = id("VisitUser")
	GetBotInventory = id("GetBotInventory")
	NewUserExperienceGetGifts = id("NewUserExperienceGetGifts")
	NewUserExperienceScriptProceed = id("NewUserExperienceScriptProceed")
	SelectInitialRoom = id("SelectInitialRoom")
	Game2MakeSnowball = id("Game2MakeSnowball")
	Game2RequestFullStatusUpdate = id("Game2RequestFullStatusUpdate")
	Game2SetUserMoveTarget = id("Game2SetUserMoveTarget")
	Game2ThrowSnowballAtHuman = id("Game2ThrowSnowballAtHuman")
	Game2ThrowSnowballAtPosition = id("Game2ThrowSnowballAtPosition")
	ResetPhoneNumberState = id("ResetPhoneNumberState")
	SetPhoneNumberVerificationStatus = id("SetPhoneNumberVerificationStatus")
	TryPhoneNumber = id("TryPhoneNumber")
	VerifyCode = id("VerifyCode")
	ClickFurni = id("ClickFurni")
	CompostPlant = id("CompostPlant")
	GetFurnitureAliases = id("GetFurnitureAliases")
	GetHeightMap = id("GetHeightMap")
	GetItemData = id("GetItemData")
	GetPetCommands = id("GetPetCommands")
	GiveSupplementToPet = id("GiveSupplementToPet")
	HarvestPet = id("HarvestPet")
	MountPet = id("MountPet")
	MoveAvatar = id("MoveAvatar")
	MoveObject = id("MoveObject")
	MovePet = id("MovePet")
	MoveWallItem = id("MoveWallItem")
	PickupObject = id("PickupObject")
	PlaceBot = id("PlaceBot")
	PlaceObject = id("PlaceObject")
	PlacePet = id("PlacePet")
	RemoveBotFromFlat = id("RemoveBotFromFlat")
	RemoveItem = id("RemoveItem")
	RemovePetFromFlat = id("RemovePetFromFlat")
	RemoveSaddleFromPet = id("RemoveSaddleFromPet")
	SetClothingChangeData = id("SetClothingChangeData")
	SetItemData = id("SetItemData")
	SetObjectData = id("SetObjectData")
	TogglePetBreedingPermission = id("TogglePetBreedingPermission")
	TogglePetRidingPermission = id("TogglePetRidingPermission")
	UseFurniture = id("UseFurniture")
	UseWallItem = id("UseWallItem")
	GetOccupiedTiles = id("GetOccupiedTiles")
	GetRoomEntryTile = id("GetRoomEntryTile")
	UpdateFloorProperties = id("UpdateFloorProperties")
	PhotoCompetition = id("PhotoCompetition")
	PublishPhoto = id("PublishPhoto")
	PurchasePhoto = id("PurchasePhoto")
	RenderRoom = id("RenderRoom")
	RenderRoomThumbnail = id("RenderRoomThumbnail")
	RequestCameraConfiguration = id("RequestCameraConfiguration")
	GetPromoArticles = id("GetPromoArticles")
)