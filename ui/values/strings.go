package values

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"code.cryptopower.dev/group/cryptopower/ui/values/localizable"
)

const (
	DefaultLangauge = localizable.ENGLISH
	commentPrefix   = "/"
)

var rex = regexp.MustCompile(`(?m)("(?:\\.|[^"\\])*")\s*=\s*("(?:\\.|[^"\\])*")`) // "key"="value"
var Languages = []string{localizable.ENGLISH, localizable.CHINESE, localizable.FRENCH, localizable.SPANISH}
var UserLanguages = []string{DefaultLangauge} // order of preference

var languageStrings map[string]map[string]string

func init() {

	readIntoMap := func(m map[string]string, localizableStrings string) {
		scanner := bufio.NewScanner(strings.NewReader(localizableStrings))
		for scanner.Scan() {
			line := scanner.Text()
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, commentPrefix) {
				continue
			}

			matches := rex.FindAllStringSubmatch(line, -1)
			if len(matches) == 0 {
				continue
			}

			kv := matches[0]
			key := trimQuotes(kv[1])
			value := trimQuotes(kv[2])

			m[key] = value
		}
	}

	en := make(map[string]string)
	zh := make(map[string]string)
	fr := make(map[string]string)
	es := make(map[string]string)
	languageStrings = make(map[string]map[string]string)

	readIntoMap(en, localizable.EN)
	languageStrings[localizable.ENGLISH] = en

	readIntoMap(zh, localizable.ZH)
	languageStrings[localizable.CHINESE] = zh

	readIntoMap(fr, localizable.FR)
	languageStrings[localizable.FRENCH] = fr

	readIntoMap(es, localizable.ES)
	languageStrings[localizable.SPANISH] = es
}

func hasLanguage(language string) bool {
	for _, lang := range Languages {
		if lang == language {
			return true
		}
	}

	return false
}

func SetUserLanguage(lang string) {
	if hasLanguage(lang) {
		languages := []string{lang}
		if lang != DefaultLangauge {
			languages = append(languages, DefaultLangauge)
		}

		UserLanguages = languages
	}
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func String(key string) string {
	for _, lang := range UserLanguages {
		languageMap := languageStrings[lang]
		str, ok := languageMap[key]
		if ok {
			return str
		}
	}

	return ""
}

func StringF(key string, a ...interface{}) string {
	str := String(key)
	if str == "" {
		return str
	}

	return fmt.Sprintf(str, a...)
}

const (
	StrAbandoned                       = "abandoned"
	StrAbout                           = "about"
	StrAbstain                         = "abstain"
	StrAccount                         = "account"
	StrAccountMixer                    = "accountMixer"
	StrAcctCreated                     = "acctCreated"
	StrAcctDetailsKey                  = "acctDetailsKey"
	StrAcctName                        = "acctName"
	StrAcctNum                         = "acctNum"
	StrAcctRenamed                     = "accRenamed"
	StrAddAcctWarn                     = "addAcctWarn"
	StrAddDexServer                    = "addDexServer"
	StrAddNewAccount                   = "addNewAccount"
	StrAddress                         = "address"
	StrAddressCopied                   = "addressCopied"
	StrAddressDiscoveryInProgress      = "addressDiscoveryInProgress"
	StrAddressDiscoveryStarted         = "addressDiscoveryStarted"
	StrAddressDiscoveryStartedBody     = "addressDiscoveryStartedBody"
	StrAddrNotOwned                    = "addrNotOwned"
	StrAddVSP                          = "addVSP"
	StrAddWallet                       = "addWallet"
	StrAdminToTriggerVoting            = "adminToTriggerVoting"
	StrAgo                             = "ago"
	StrAll                             = "all"
	StrAllowSpendingFromUnmixedAccount = "allowSpendingFromUnmixedAccount"
	StrAllowUnspendUnmixedAcct         = "allowUnspendUnmixedAcct"
	StrAllTickets                      = "allTickets"
	StrAmount                          = "amount"
	StrAppName                         = "appName"
	StrApproved                        = "approved"
	StrAppTitle                        = "appTitle"
	StrAppWallet                       = "appWallet"
	StrAskedEnterSeedWords             = "askedEnterSeedWords"
	StrAuthorToAuthorizeVoting         = "authorToAuthorizeVoting"
	StrAutomatic                       = "automatic"
	StrAutoSetUp                       = "autoSetUp"
	StrAutoSync                        = "autoSync"
	StrAutoTicketInfo                  = "autoTicketInfo"
	StrAutoTicketPurchase              = "autoTicketPurchase"
	StrAutoTicketWarn                  = "autoTicketWarn"
	StrAwareOfRisk                     = "imawareOfRisk"
	StrBackAndRename                   = "backAndRename"
	StrBackStaking                     = "backStaking"
	StrBackToWallets                   = "backToWallets"
	StrBackupInfo                      = "backupInfo"
	StrBackupLater                     = "backupLater"
	StrBackupNow                       = "backupNow"
	StrBackupSeedPhrase                = "backupSeedPhrase"
	StrBackupWarning                   = "backupWarning"
	StrBalance                         = "balance"
	StrBalanceAfter                    = "balanceAfter"
	StrBalToMaintain                   = "balToMaintain"
	StrBalToMaintainValue              = "balToMaintainValue"
	StrBeepForNewBlocks                = "beepForNewBlocks"
	StrBestBlockAge                    = "bestBlockAge"
	StrBestBlocks                      = "bestBlocks"
	StrBestBlockTimestamp              = "bestBlockTimestamp"
	StrBlockHeaderFetched              = "blockHeaderFetched"
	StrBlockHeaderFetchedCount         = "blockHeaderFetchedCount"
	StrBlocksLeft                      = "blocksLeft"
	StrBlocksScanned                   = "blocksScanned"
	StrBuild                           = "build"
	StrBuildDate                       = "buildDate"
	StrCanBuy                          = "canBuy"
	StrCancel                          = "cancel"
	StrCanceling                       = "canceling"
	StrCancelMixer                     = "cancelMixer"
	StrChange                          = "change"
	StrChangeAccount                   = "changeAccount"
	StrChangeSpecificPeer              = "changeSpecificPeer"
	StrChangeSpendingPass              = "changeSpendingPass"
	StrChangeStartupPassword           = "changeStartupPassword"
	StrChangeUserAgent                 = "changeUserAgent"
	StrChangeWalletName                = "changeWalletName"
	StrCheckGovernace                  = "checkGovernace"
	StrCheckMixerStatus                = "checkMixerStatus"
	StrCheckStatistics                 = "checkStatistics"
	StrCheckWalletLog                  = "checkWalletLog"
	StrClear                           = "clear"
	StrClearAll                        = "clearAll"
	StrClosingWallet                   = "closingWallet"
	StrCoinSelection                   = "coinSelection"
	StrColon                           = "colon"
	StrComplete                        = "complete"
	StrConfirm                         = "confirm"
	StrConfirmDexReset                 = "confirmDexReset"
	StrConfirmed                       = "confirmed"
	StrConfirmNewSpendingPassword      = "confirmNewSpendingPassword"
	StrConfirmNewStartupPass           = "confirmNewStartupPass"
	StrConfirmOrder                    = "confirmOrder"
	StrConfirmPending                  = "confirmPending"
	StrConfirmPurchase                 = "confirmPurchase"
	StrConfirmRemoveStartupPass        = "confirmRemoveStartupPass"
	StrConfirmSend                     = "confirmSend"
	StrConfirmShowSeed                 = "confirmToShowSeed"
	StrConfirmSpendingPassword         = "confirmSpendingPassword"
	StrConfirmStartupPass              = "confirmStartupPass"
	StrConfirmToCreateAccs             = "confirmtoCreateAccs"
	StrConfirmToMixAccount             = "confirmToMixAcc"
	StrConfirmToRemove                 = "confirmToRemove"
	StrConfirmToSetMixer               = "confirmToSetMixer"
	StrConfirmToSign                   = "confirmToSign"
	StrConfirmToVerifySeed             = "confirmToVerifySeed"
	StrConfirmUmixedSpending           = "confirmUmixedSpending"
	StrConfirmVote                     = "confirmVote"
	StrConfirmYourOrder                = "confirmYourOrder"
	StrConfStatus                      = "confStatus"
	StrConnectedPeersCount             = "connectedPeersCount"
	StrConnectedTo                     = "connectedTo"
	StrConnecting                      = "connecting"
	StrConnection                      = "connection"
	StrConnectToSpecificPeer           = "connectToSpecificPeer"
	StrConsensusChange                 = "consensusChange"
	StrConsensusDashboard              = "consensusDashboard"
	StrContinue                        = "continue"
	StrCoordinationServer              = "coordinationServer"
	StrCopied                          = "copied"
	StrCopy                            = "copy"
	StrCopyBlockLink                   = "copyBlockLink"
	StrCopySeed                        = "copyseed"
	StrCopyLink                        = "copyLink"
	StrCost                            = "cost"
	StrCreate                          = "create"
	StrCreateANewWallet                = "createANewWallet"
	StrCreateNewAccount                = "createNewAccount"
	StrCreateNewOrder                  = "createNewOrder"
	StrCreateNSetUpAccs                = "createNSetUpAccs"
	StrCreateOrder                     = "createOrder"
	StrCreateOrderPageInfo             = "createOrderPageInfo"
	StrCreateStartupPassword           = "createStartupPassword"
	StrCurrentSpendingPassword         = "currentSpendingPassword"
	StrCurrentStartupPass              = "currentStartupPass"
	StrCurrentTotalBalance             = "currentTotalBalance"
	StrCustomUserAgent                 = "CustomUserAgent"
	StrDangerZone                      = "dangerZone"
	StrDarkMode                        = "darkMode"
	StrDateSize                        = "dateSize"
	StrDayAgo                          = "dayAgo"
	StrDays                            = "days"
	StrDaysAgo                         = "daysAgo"
	StrDaysToMiss                      = "daysToMiss"
	StrDaysToVote                      = "daysToVote"
	StrDCRCaps                         = "dcrCaps"
	StrDcrReceived                     = "dcrReceived"
	StrDebug                           = "debug"
	StrDefault                         = "default"
	StrDeleted                         = "delete"
	StrDestAddr                        = "destAddr"
	StrDestination                     = "destination"
	StrDestinationModalInfo            = "destinationModalInfo"
	StrDestinationWalletNotSynced      = "destinationWalletNotSynced"
	StrDex                             = "dex"
	StrDexDataReset                    = "dexDataReset"
	StrDexDataResetFalse               = "dexDataResetFalse"
	StrDexResetInfo                    = "dexResetInfo"
	StrDexStartupErr                   = "dexStartupErr"
	StrDisable                         = "disable"
	StrDisabled                        = "disabled"
	StrDisconnect                      = "disconnect"
	StrDiscoverAddressUsage            = "discoverAddressUsage"
	StrDiscoveringWalletAddress        = "discoveringWalletAddress"
	StrDiscussions                     = "discussions"
	StrDocumentation                   = "documentation"
	StrDuration                        = "duration"
	StrEdit                            = "edit"
	StrEmptyMsg                        = "emptyMsg"
	StrEmptySign                       = "emptySign"
	StrEnabled                         = "enabled"
	StrEnglish                         = "english"
	StrEnterAddressToSign              = "enterAddressToSign"
	StrEnterExtendedPubKey             = "enterXpubKey"
	StrEnterHex                        = "enterHex"
	StrEnterSeedPhrase                 = "enterSeedPhrase"
	StrEnterSpendingPassword           = "enterSpendingPassword"
	StrEnterValidAddress               = "enterValidAddress"
	StrEnterValidMsg                   = "enterValidMsg"
	StrEnterWalDetails                 = "enterWalletDetails"
	StrEnterWalletName                 = "enterWalletName"
	StrEnterWalletSeed                 = "enterWalletSeed"
	StrErrorMovingFunds                = "errorMovingFunds"
	StrErrPassEmpty                    = "errPassEmpty"
	StrEstimatedSize                   = "estimatedSize"
	StrEstimatedTime                   = "estimatedTime"
	StrExchange                        = "exchange"
	StrExchangeIntro                   = "exchangeIntro"
	StrExchangeRate                    = "exchangeRate"
	StrExistingWalletName              = "existingWalletName"
	StrExit                            = "exit"
	StrExpired                         = "expired"
	StrExpiredInfo                     = "expiredInfo"
	StrExpiredInfoDisc                 = "expiredInfoDisc"
	StrExpiredInfoDiscSub              = "expiredInfoDiscSub"
	StrExpiredOn                       = "expiredOn"
	StrExpiresIn                       = "expiresIn"
	StrExplorerURL                     = "explorerURL"
	StrExtendedCopied                  = "extendedKeyCopied"
	StrExtendedInfo                    = "extendedInfo"
	StrExtendedKey                     = "extendedKey"
	StrExtendedPubKey                  = "extendedPubKey"
	StrExternal                        = "external"
	StrFailed                          = "failed"
	StrFee                             = "fee"
	StrFeeRates                        = "feerates"
	StrFetchingAgenda                  = "fetchingAgenda"
	StrFetchingBlockHeaders            = "fetchingBlockHeaders"
	StrFetchingOrders                  = "fetchingOrders"
	StrFetchingPolicies                = "fetchingPolicies"
	StrFetchingProposals               = "fetchingProposals"
	StrFetchProposals                  = "fetchProposals"
	StrFetchRateError                  = "fetchRateError"
	StrFetchRates                      = "fetchRates"
	StrFinished                        = "finished"
	StrFrench                          = "french"
	StrFrom                            = "from"
	StrFunctionUnavailable             = "functionUnavailable"
	StrGapLimit                        = "gapLimit"
	StrGapLimitInputErr                = "gapLimitInputErr"
	StrGeneral                         = "general"
	StrGenerateAddress                 = "generateAddress"
	StrGotIt                           = "gotIt"
	StrGovernance                      = "governance"
	StrGovernanceInfo                  = "governanceInfo"
	StrGovernanceSettingsInfo          = "governanceSettingsInfo"
	StrHash                            = "hash"
	StrHDPath                          = "hdPath"
	StrHelp                            = "help"
	StrHelpInfo                        = "helpInfo"
	StrHex                             = "hex"
	StrHideDetails                     = "hideDetails"
	StrHideSeedPhrase                  = "hideSeedPhrase"
	StrHint                            = "hint"
	StrHistory                         = "history"
	StrHourAgo                         = "hourAgo"
	StrHours                           = "hours"
	StrHoursAgo                        = "hoursAgo"
	StrHowGovernanceWork               = "howGovernanceWork"
	StrHowNotToStoreSeedPhrase         = "howNotToStoreSeedPhrase"
	StrHowToCopy                       = "howToCopy"
	StrHowToStoreSeedPhrase            = "howToStoreSeedPhrase"
	StrHTTPRequest                     = "httpReq"
	StrImmature                        = "immature"
	StrImmatureInfo                    = "immatureInfo"
	StrImmatureRewards                 = "immatureRewards"
	StrImmatureStakeGen                = "immatureStakeGen"
	StrImport                          = "Import"
	StrImportantSeedPhrase             = "importantSeedPhrase"
	StrImported                        = "import"
	StrImportExistingWallet            = "importExistingWallet"
	StrImportWatchingOnlyWallet        = "importWatchingOnlyWallet"
	StrIncludedInBlock                 = "includedInBlock"
	StrInDiscussion                    = "inDiscussion"
	StrInfo                            = "info"
	StrInitiateSetup                   = "initiateSetup"
	StrInProgress                      = "inprogress"
	StrInsufficentFund                 = "insufficentFund"
	StrInvalidAddress                  = "invalidAddress"
	StrInvalidAmount                   = "invalidAmount"
	StrInvalidHex                      = "invalidHex"
	StrInvalidPassphrase               = "invalidPassphrase"
	StrInvalidSeedPhrase               = "invalidSeedPhrase"
	StrInvalidSignature                = "invalidSignature"
	StrIPAddress                       = "ipAddress"
	StrJustNow                         = "justNow"
	StrKeepAppOpen                     = "keepAppOpen"
	StrKeepInMind                      = "keepInMind"
	StrKey                             = "key"
	StrLabelSpendable                  = "labelSpendable"
	StrLanguage                        = "language"
	StrLastBlockHeight                 = "lastBlockHeight"
	StrLatestBlock                     = "latestBlock"
	StrLicense                         = "license"
	StrLifeSpan                        = "lifeSpan"
	StrLive                            = "live"
	StrLiveIn                          = "liveIn"
	StrLiveInfo                        = "liveInfo"
	StrLiveInfoDisc                    = "liveInfoDisc"
	StrLiveInfoDiscSub                 = "liveInfoDiscSub"
	StrLiveTickets                     = "liveTickets"
	StrLoading                         = "loading"
	StrLoadingPrice                    = "loadingPrice"
	StrLocked                          = "locked"
	StrLockedByTickets                 = "lockedByTickets"
	StrLockedIn                        = "lockedin"
	StrManualSetUp                     = "manualSetUp"
	StrMaturity                        = "maturity"
	StrMax                             = "max"
	StrMessage                         = "message"
	StrMinimumAssetType                = "minimumAssetType"
	StrMinMax                          = "minMax"
	StrMinuteAgo                       = "minuteAgo"
	StrMinutes                         = "mins"
	StrMinutesAgo                      = "minutesAgo"
	StrMissedOn                        = "missedOn"
	StrMissedTickets                   = "missedTickets"
	StrMix                             = "mix"
	StrMixed                           = "mixed"
	StrMixedAccDisabled                = "mixedAccDisabled"
	StrMixedAccHidden                  = "mixedAccHidden"
	StrMixedAccount                    = "mixedAccount"
	StrMixer                           = "mixer"
	StrMixerAccErrorMsg                = "mixerAccErrorMsg"
	StrMixerRunning                    = "mixerRunning"
	StrMixerShutdown                   = "mixerShutdown"
	StrMixerStart                      = "mixerStart"
	StrMixingActivity                  = "mixingActivity"
	StrMonthAgo                        = "monthAgo"
	StrMonthsAgo                       = "monthsAgo"
	StrMore                            = "more"
	StrMoveFundsFrmDefaultToUnmixed    = "moveFundsFrmDefaultToUnmixed"
	StrMoveToUnmixed                   = "moveToUnmixed"
	StrMultipleMixerAccNeeded          = "multipleMixerAccNeeded"
	StrMyAcct                          = "myAcct"
	StrNConfirmations                  = "nConfirmations"
	StrNetwork                         = "network"
	StrNewest                          = "newest"
	StrNewProposalUpdate               = "newProposalUpdate"
	StrNewSpendingPassword             = "newSpendingPassword"
	StrNewStartupPass                  = "newStartupPass"
	StrNewWallet                       = "newWallet"
	StrNext                            = "next"
	StrNo                              = "no"
	StrNoActiveTickets                 = "noActiveTickets"
	StrNoAgendaYet                     = "noAgendaYet"
	StrNoConnectedPeer                 = "noConnectedPeer"
	StrNoInternet                      = "noInternet"
	StrNoMixable                       = "errNoMixable"
	StrNonAccSelector                  = "nonAccSelector"
	StrNone                            = "none"
	StrNoOrders                        = "noOrders"
	StrNoPoliciesYet                   = "noPoliciesYet"
	StrNoProposals                     = "noProposal"
	StrNoReward                        = "noReward"
	StrNotAvailable                    = "notAvailable"
	StrNotBackedUp                     = "notBackedUp"
	StrNotConnected                    = "notConnected"
	StrNotEnoughVotes                  = "notEnoughVotes"
	StrNoTickets                       = "noTickets"
	StrNotifications                   = "notifications"
	StrNotOwned                        = "notOwned"
	StrNoTransactions                  = "noTransactions"
	StrNotSameAccoutMixUnmix           = "notSameAccoutMixUnmix"
	StrNoValidAccountFound             = "noValidAccountFound"
	StrnoValidWalletFound              = "noValidWalletFound"
	StrNoVSPLoaded                     = "noVSPLoaded"
	StrNoWalletLoaded                  = "noWalletLoaded"
	StrNumberOfVotes                   = "numberOfVotes"
	StrOffChainVote                    = "offChainVote"
	StrOffline                         = "offline"
	StrOk                              = "ok"
	StrOK                              = "ok"
	StrOldest                          = "oldest"
	StrOnChainVote                     = "onChainVote"
	StrOnline                          = "online"
	StrOpeningWallet                   = "openingWallet"
	StrOrderCeated                     = "orderCreated"
	StrOrderDetails                    = "orderDetails"
	StrOrderReceivingTo                = "orderReceivingTo"
	StrOrderSendingFrom                = "orderSendingFrom"
	StrOrderSettingsSaved              = "orderSettingsSaved"
	StrOverview                        = "overview"
	StrOwned                           = "owned"
	StrPageWarningNotSync              = "pageWarningNotSync"
	StrPageWarningSync                 = "pageWarningSync"
	StrPasswordNotMatch                = "passwordNotMatch"
	StrPasteSeedWords                  = "pasteSeedWords"
	StrPeer                            = "peer"
	StrPeers                           = "peers"
	StrPeersConnected                  = "peersConnected"
	StrPending                         = "pending"
	StrPercentageMixed                 = "percentageMixed"
	StrPiKey                           = "piKey"
	StrPolicySetSuccessful             = "policySetSuccessfully"
	StrPriority                        = "priority"
	StrPrivacyInfo                     = "privacyInfo"
	StrPropFetching                    = "propFetching"
	StrPropNotif                       = "propNotif"
	StrPropNotification                = "propNotification"
	StrProposal                        = "proposals"
	StrProposalAddedNotif              = "proposalAddedNotif"
	StrProposalInfo                    = "proposalInfo"
	StrProposalVoteDetails             = "proposalVoteDetails"
	StrPublished                       = "published"
	StrPublished2                      = "published2"
	StrPurchased                       = "purchased"
	StrPurchasedOn                     = "purchasedOn"
	StrPurchasingAcct                  = "purchasingAcct"
	StrQuorumRequirement               = "quorumRequirement"
	StrRate                            = "rate"
	StrReadyToMix                      = "readyToMix"
	StrRebroadcast                     = "rebroadcast"
	StrReceive                         = "receive"
	StrReceived                        = "received"
	StrReceiveInfo                     = "receiveInfo"
	StrReceiving                       = "receiving"
	StrReceivingAddress                = "receivingAddress"
	StrRecentProposals                 = "recentProposals"
	StrRecentTransactions              = "recentTransactions"
	StrReconnect                       = "reconnect"
	StrRefresh                         = "refresh"
	StrRejected                        = "rejected"
	StrRemove                          = "remove"
	StrRemovePeer                      = "removePeer"
	StrRemovePeerWarn                  = "removePeerWarn"
	StrRemoveUserAgent                 = "removeUserAgent"
	StrRemoveUserAgentWarn             = "removeUserAgentWarn"
	StrRemoveWallet                    = "removeWallet"
	StrRemoveWalletInfo                = "removeWalletInfo"
	StrRename                          = "rename"
	StrRenameAcct                      = "renameAcct"
	StrRenameWalletSheetTitle          = "renameWalletSheetTitle"
	StrRepublished                     = "republished"
	StrRescan                          = "rescan"
	StrRescanBlockchain                = "rescanBlockchain"
	StrRescanInfo                      = "rescanInfo"
	StrRescanningBlocks                = "rescanningBlocks"
	StrRescanningHeaders               = "rescanningHeaders"
	StrRescanProgressNotification      = "rescanProgressNotification"
	StrRestore                         = "restore"
	StrRestoreExistingWallet           = "restoreExistingWallet"
	StrRestoreWallet                   = "restoreWallet"
	StrRestoreWithHex                  = "restoreWithHex"
	StrResumeAccountDiscoveryTitle     = "resumeAccountDiscoveryTitle"
	StrRetry                           = "retry"
	StrRevocation                      = "revocation"
	StrRevoke                          = "revoke"
	StrRevokeCause                     = "revokeCause"
	StrRevoked                         = "revoked"
	StrRevokeInfo                      = "revokeInfo"
	StrRevokeInfoDisc                  = "revokeInfoDisc"
	StrReward                          = "reward"
	StrRewardsEarned                   = "rewardsEarned"
	StrSave                            = "save"
	StrSearch                          = "search"
	StrSeconds                         = "secs"
	StrSecurity                        = "security"
	StrSecurityTools                   = "securityTools"
	StrSecurityToolsInfo               = "securityToolsInfo"
	StrSeeAll                          = "seeAll"
	StrSeedAlreadyExist                = "seedAlreadyExist"
	StrSeedHex                         = "seedHex"
	StrSeedPhraseToRestore             = "seedPhraseToRestore"
	StrSeedPhraseVerified              = "seedPhraseVerified"
	StrSeedValidationFailed            = "seedValidationFailed"
	StrSelectAcc                       = "selectAcc"
	StrSelectAServer                   = "selectAServer"
	StrSelectChangeAcc                 = "selectChangeAcc"
	StrSelectDexServerToOpen           = "selectDexServerToOpen"
	StrSelectedAccount                 = "selectedAcct"
	StrSelectMixedAcc                  = "selectMixedAcc"
	StrSelectOption                    = "selectOption"
	StrSelectPhrasesToVerify           = "selectPhrasesToVerify"
	StrSelectServerTitle               = "selectServerTitle"
	StrSelectTicket                    = "selectTicket"
	StrSelectVSP                       = "selectVSP"
	StrSelectWallet                    = "selectWallet"
	StrSelectWalletToOpen              = "selectWalletToOpen"
	StrSelectWalletType                = "selectWalletType"
	StrSend                            = "send"
	StrSendConfModalTitle              = "sendConfModalTitle"
	StrSendInfo                        = "sendInfo"
	StrSending                         = "sending"
	StrSendingAcct                     = "sendingAcct"
	StrSendingFrom                     = "sendingFrom"
	StrSendWarning                     = "sendWarning"
	StrSent                            = "sent"
	StrServer                          = "server"
	StrSetChoice                       = "setchoice"
	StrSetGapLimit                     = "setGapLimit"
	StrSetGapLimitInfo                 = "setGapLimitInfo"
	StrSettings                        = "settings"
	StrSetTreasuryPolicy               = "setTreasuryPolicy"
	StrSetUp                           = "setUp"
	StrSetupMixerInfo                  = "setupMixerInfo"
	StrSetUpNeededAccs                 = "setUpNeededAccs"
	StrSetUpPrivacy                    = "setUpPrivacy"
	StrSetupStakeShuffle               = "setUpStakeShuffle"
	StrSetupStartupPassword            = "setupStartupPassword"
	StrSignature                       = "signature"
	StrSignCopied                      = "signCopied"
	StrSignMessage                     = "signMessage"
	StrSignMessageInfo                 = "signMessageInfo"
	StrSource                          = "source"
	StrSourceModalInfo                 = "sourceModalInfo"
	StrSourceWalletNotSynced           = "sourceWalletNotSynced"
	StrSpanish                         = "spanish"
	StrSpendableIn                     = "spendableIn"
	StrSpendingPassword                = "spendingPassword"
	StrSpendingPasswordInfo            = "spendingPasswordInfo"
	StrSpendingPasswordInfo2           = "spendingPasswordInfo2"
	StrSpendingPasswordUpdated         = "spendingPasswordUpdated"
	StrStake                           = "stake"
	StrStakeAge                        = "stakeAge"
	StrStaked                          = "staked"
	StrStakeShuffle                    = "stakeShuffle"
	StrStaking                         = "staking"
	StrStakingActivity                 = "stakingActivity"
	StrStartupPassConfirm              = "startupPassConfirm"
	StrStartupPassword                 = "startupPassword"
	StrStartupPasswordEnabled          = "startupPasswordEnabled"
	StrStartupPasswordInfo             = "startupPasswordInfo"
	StrStatistics                      = "statistics"
	StrStatus                          = "status"
	StrStep1                           = "step1"
	StrStep2of2                        = "step2of2"
	StrSubmit                          = "submit"
	StrSureToCancelMixer               = "sureToCancelMixer"
	StrSureToExitBackup                = "sureToExitBackup"
	StrSureToSafeStoreSeed             = "sureToSafeStoreSeed"
	StrSync                            = "sync"
	StrSyncCompTime                    = "syncCompTime"
	StrSynced                          = "synced"
	StrSyncingProgress                 = "syncingProgress"
	StrSyncingProgressStat             = "syncingProgressStat"
	StrSyncingState                    = "syncingState"
	StrSyncSteps                       = "syncSteps"
	StrTakenAccount                    = "takenAccount"
	StrTapToCopy                       = "tapToCopy"
	StrTicektVoted                     = "ticektVoted"
	StrTicket                          = "ticket"
	StrTicketConfirmed                 = "ticketConfirmed"
	StrTicketDetails                   = "ticketDetails"
	StrTicketError                     = "ticketError"
	StrTicketPrice                     = "ticketPrice"
	StrTicketRecord                    = "ticketRecord"
	StrTicketRevoked                   = "ticketRevoked"
	StrTicketRevokedTitle              = "ticketRevokedTitle"
	StrTickets                         = "tickets"
	StrTicketSettingSaved              = "ticketSettingSaved"
	StrTicketVotedTitle                = "ticketVotedTitle"
	StrTimeLeft                        = "timeLeft"
	StrTo                              = "to"
	StrToken                           = "token"
	StrTotal                           = "total"
	StrTotalBalance                    = "totalBalance"
	StrTotalCost                       = "totalCost"
	StrTotalVotes                      = "totalVotes"
	StrTotalVotesReverse               = "totalVotesReverse"
	StrTransactionDetails              = "transactionDetails"
	StrTransactionID                   = "transactionId"
	StrTransactions                    = "transactions"
	StrTransferred                     = "transferred"
	StrTreasurySpending                = "treasurySpending"
	StrTreasurySpendingInfo            = "treasurySpendingInfo"
	StrTxConfModalInfoTxt              = "txConfModalInfoTxt"
	StrTxdetailsInfo                   = "txDetailsInfo"
	StrTxEstimateErr                   = "txEstimateErr"
	StrTxFee                           = "txFee"
	StrTxHashCopied                    = "txHashCopied"
	StrTxNotification                  = "txNotification"
	StrTxOverview                      = "txOverview"
	StrTxSent                          = "txSent"
	StrTxSize                          = "txSize"
	StrTxStatusPending                 = "txStatusPending"
	StrType                            = "type"
	StrUmined                          = "unmined"
	StrUnconfirmedFunds                = "unconfirmedFunds"
	StrUnconfirmedTx                   = "unconfirmedTx"
	StrUnderReview                     = "underReview"
	StrUnknown                         = "unknown"
	StrUnlock                          = "unlock"
	StrUnlockWithPassword              = "unlockWithPassword"
	StrUnminedInfo                     = "unminedInfo"
	StrUnmixed                         = "unmixed"
	StrUnmixedAccount                  = "unmixedAccount"
	StrUnmixedBalance                  = "unmixedBalance"
	StrUpcoming                        = "upcomming"
	StrUpdated                         = "updated"
	StrUpdatePreference                = "updatePreference"
	StrUpdatevotePref                  = "updateVotePref"
	StrUptime                          = "uptime"
	StrUsdBinance                      = "usdBinance"
	StrUsdBittrex                      = "usdBittrex"
	StrUseMixer                        = "useMixer"
	StrUserAgent                       = "userAgent"
	StrUserAgentDialogTitle            = "userAgentDialogTitle"
	StrUserAgentSummary                = "userAgentSummary"
	StrValidAddress                    = "validAddress"
	StrValidate                        = "validate"
	StrValidateAddr                    = "validateAddr"
	StrValidateHostErr                 = "validateHostErr"
	StrValidateMsg                     = "validateMsg"
	StrValidateNote                    = "validateNote"
	StrValidateWalSeed                 = "validateWalSeed"
	StrValidSignature                  = "validSignature"
	StrVerify                          = "verify"
	StrVerifyGovernanceKeys            = "verifyGovernanceKeys"
	StrVerifyMessage                   = "verifyMessage"
	StrVerifyMessageInfo               = "verifyMessageInfo"
	StrVerifyMsgError                  = "verifyMsgError"
	StrVerifyMsgNote                   = "verifyMsgNote"
	StrVerifySeed                      = "verifySeed"
	StrVerifySeedInfo                  = "verifySeedInfo"
	StrVersion                         = "version"
	StrViewDetails                     = "viewDetails"
	StrViewOnExplorer                  = "viewOnExplorer"
	StrViewOnPoliteia                  = "viewOnPoliteia"
	StrViewProperty                    = "viewProperty"
	StrViewSeedPhrase                  = "viewSeedPhrase"
	StrViewTicket                      = "viewTicket"
	StrVote                            = "vote"
	StrVoteChoice                      = "votechoice"
	StrVoteConfirm                     = "voteConfirm"
	StrVoted                           = "voted"
	StrVotedInfo                       = "votedInfo"
	StrVotedInfoDisc                   = "votedInfoDisc"
	StrVotedOn                         = "votedOn"
	StrVoteEndedNotif                  = "voteEndedNotif"
	StrVoteSent                        = "voteSent"
	StrVoteStartedNotif                = "voteStartedNotif"
	StrVoteTooltip                     = "voteTooltip"
	StrVoteUpdated                     = "voteUpdated"
	StrVoting                          = "voting"
	StrVotingAuthority                 = "votingAuthority"
	StrVotingDashboard                 = "votingDashboard"
	StrVotingInProgress                = "votingInProgress"
	StrVotingPreference                = "votingPreference"
	StrVotingServiceProvider           = "votingServiceProvider"
	StrVotingWallet                    = "votingWallet"
	StrVsp                             = "vsp"
	StrVspFee                          = "vspFee"
	StrWaitingAuthor                   = "waitingForAuthor"
	StrWaitingForAdmin                 = "waitingForAdmin"
	StrWaitingState                    = "waitingState"
	StrWalletCreated                   = "walletCreated"
	StrWalletDirectory                 = "walletDirectory"
	StrWalletExist                     = "walletExist"
	StrWalletLog                       = "walletLog"
	StrWalletName                      = "walletName"
	StrWalletNameLengthError           = "walletLengthError"
	StrWalletNameMismatch              = "walletNameMismatch"
	StrWalletNotExist                  = "walletNotExist"
	StrWalletNotSynced                 = "walletNotSynced"
	StrWalletRemoved                   = "walletRemoved"
	StrWalletRemoveInfo                = "walletRemoveInfo"
	StrWalletRenamed                   = "walletRenamed"
	StrWalletRestored                  = "walletRestored"
	StrWalletRestoreMsg                = "walletRestoreMsg"
	StrWallets                         = "wallets"
	StrWalletsEnabledPrivacy           = "walletsEnabledPrivacy"
	StrWalletSettings                  = "walletSettings"
	StrWalletStatus                    = "walletStatus"
	StrWalletSyncing                   = "walletSyncing"
	StrWalletToPurchaseFrom            = "walletToPurchaseFrom"
	StrWarningVote                     = "warningVote"
	StrWarningWatchWallet              = "warningWatchWallet"
	StrWatchOnly                       = "watchOnly"
	StrWatchOnlyWalletImported         = "watchOnlyWalletImported"
	StrWatchOnlyWalletRemoveInfo       = "watchOnlyWalletRemoveInfo"
	StrWatchOnlyWallets                = "watchOnlyWallets"
	StrWebURL                          = "webURL"
	StrWeekAgo                         = "weekAgo"
	StrWeeksAgo                        = "weeksAgo"
	StrWelcomeNote                     = "welcomeNote"
	StrWhatToCallWallet                = "whatToCallWallet"
	StrWord                            = "word"
	StrWriteDownAll33Words             = "writeDownAll33Words"
	StrWriteDownSeed                   = "writeDownSeed"
	StrWroteAllWords                   = "wroteAllWords"
	StrXInputsConsumed                 = "xInputsConsumed"
	StrXOutputCreated                  = "xOutputCreated"
	StrXpubKeyErr                      = "xpubKeyErr"
	StrXpubWalletExist                 = "xpubWalletExist"
	StrYearAgo                         = "yearAgo"
	StrYearsAgo                        = "yearsAgo"
	StrYes                             = "yes"
	StrYesterday                       = "yesterday"
	StrYourAddress                     = "yourAddress"
	StrYourSeedWords                   = "yourSeedWord"
	StrYourself                        = "yourself"
	StrOrderSubmitted                  = "orderSubmitted"
)
