package vast

// Ad is Top-level element, wraps each ad in the response
type Ad struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Identifies the sequence of multiple Ads and defines an Ad Pod.
	Sequence int `xml:"sequence,attr,omitempty" json:"sequence,omitempty"`

	// Second-level element surrounding complete ad data for a single ad
	InLine *InLine `xml:"InLine,omitempty" json:"in_line,omitempty"`

	// Second-level element surrounding wrapper ad pointing to Secondary ad server.
	Wrapper *Wrapper `xml:"Wrapper,omitempty" json:"wrapper,omitempty"`
}

//AdParameters ...
type AdParameters struct {
	Value string `xml:",chardata" json:"value"`

	// Specifies whether the parameters are XML-encoded
	XMLEncoded bool `xml:"xmlEncoded,attr,omitempty" json:"xml_encoded,omitempty"`
}

//AdSystem ...
type AdSystem struct {
	Value string `xml:",chardata" json:"value"`

	// Internal version used by ad system
	Version string `xml:"version,attr,omitempty" json:"version,omitempty"`
}

// ClickThrough is URL to open as destination page when user clicks on the video
type ClickThrough struct {
	Value AnyURI `xml:",chardata" json:"value"`

	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// ClickTracking is URL to request for tracking purposes when user clicks on the video
type ClickTracking struct {
	Value AnyURI `xml:",chardata" json:"value"`

	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

//Companion ...
type Companion struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Pixel dimensions of companion slot
	Width int `xml:"width,attr" json:"width"`

	// Pixel dimensions of companion slot
	Height int `xml:"height,attr" json:"height"`

	// Pixel dimensions of the companion asset
	AssetWidth int `xml:"assetWidth,attr,omitempty" json:"asset_width,omitempty"`

	// Pixel dimensions of the companion asset
	AssetHeight int `xml:"assetHeight,attr,omitempty" json:"asset_height,omitempty"`

	// Pixel dimensions of expanding companion ad when in expanded state
	ExpandedWidth int `xml:"expandedWidth,attr,omitempty" json:"expanded_width,omitempty"`

	// Pixel dimensions of expanding companion ad when in expanded state
	ExpandedHeight int `xml:"expandedHeight,attr,omitempty" json:"expanded_height,omitempty"`

	// The apiFramework defines the method to use for communication with the companion
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// Used to match companion creative to publisher placement areas on the page.
	AdSlotID string `xml:"adSlotId,attr,omitempty" json:"ad_slot_id,omitempty"`

	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:"StaticResource,omitempty" json:"static_resource,omitempty"`

	// URL source for an IFrame to display the companion element
	IFrameResource AnyURI `xml:"IFrameResource,omitempty" json:"i_frame_resource,omitempty"`

	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:"HTMLResource,omitempty" json:"html_resource,omitempty"`

	// Any valid XML may be included in the Extensions node
	CreativeExtensions []CreativeExtension `xml:"CreativeExtensions>CreativeExtension,omitempty" json:"creative_extensions,omitempty"`

	// The name of the event to track for the element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// URL to open as destination page when user clicks on the the companion banner ad.
	CompanionClickThrough AnyURI `xml:"CompanionClickThrough,omitempty" json:"companion_click_through,omitempty"`

	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:"AltText,omitempty" json:"alt_text,omitempty"`

	// Data to be passed into the companion ads. The apiFramework defines the method to use for communication (e.g. “FlashVar”)
	AdParameters *AdParameters `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`
}

//CompanionAds ...
type CompanionAds struct {
	// Provides information about which companion creative to display. All means that the player must attempt to display all. Any means the player must attempt to play at least one. None means all companions are optional.
	Required *CompanionAdsRequired `xml:"required,attr,omitempty" json:"required,omitempty"`

	// Any number of companions in any desired pixel dimensions.
	Companions []Companion `xml:"Companion,omitempty" json:"companions,omitempty"`
}

//CompanionAdsRequired Provides information about which companion creative to display. All means that the player must attempt to display all. Any means the player must attempt to play at least one. None means all companions are optional.
type CompanionAdsRequired string

const (
	//All ...
	All CompanionAdsRequired = "all"
	//Any ...
	Any CompanionAdsRequired = "any"
	//None ...
	None CompanionAdsRequired = "none"
)

//CompanionWrapper ...
type CompanionWrapper struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Pixel dimensions of companion slot
	Width int `xml:"width,attr" json:"width"`

	// Pixel dimensions of companion slot
	Height int `xml:"height,attr" json:"height"`

	// Pixel dimensions of the companion asset
	AssetWidth int `xml:"assetWidth,attr,omitempty" json:"asset_width,omitempty"`

	// Pixel dimensions of the companion asset
	AssetHeight int `xml:"assetHeight,attr,omitempty" json:"asset_height,omitempty"`

	// Pixel dimensions of expanding companion ad when in expanded state
	ExpandedWidth int `xml:"expandedWidth,attr,omitempty" json:"expanded_width,omitempty"`

	// Pixel dimensions of expanding companion ad when in expanded state
	ExpandedHeight int `xml:"expandedHeight,attr,omitempty" json:"expanded_height,omitempty"`

	// The apiFramework defines the method to use for communication with the companion
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// Used to match companion creative to publisher placement areas on the page.
	AdSlotID string `xml:"adSlotId,attr,omitempty" json:"ad_slot_id,omitempty"`

	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:"StaticResource,omitempty" json:"static_resource,omitempty"`

	// URL source for an IFrame to display the companion element
	IFrameResource AnyURI `xml:"IFrameResource,omitempty" json:"i_frame_resource,omitempty"`

	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:"HTMLResource,omitempty" json:"html_resource,omitempty"`

	// Any valid XML may be included in the Extensions node
	CreativeExtensions []CreativeExtension `xml:"CreativeExtensions>CreativeExtension,omitempty" json:"creative_extensions,omitempty"`

	// The name of the event to track for the element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// URL to open as destination page when user clicks on the the companion banner ad.
	CompanionClickThrough AnyURI `xml:"CompanionClickThrough,omitempty" json:"companion_click_through,omitempty"`

	// URLs to ping when user clicks on the the companion banner ad.
	CompanionClickTrackings []AnyURI `xml:"CompanionClickTracking,omitempty" json:"companion_click_trackings,omitempty"`

	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:"AltText,omitempty" json:"alt_text,omitempty"`

	// Data to be passed into the companion ads. The apiFramework defines the method to use for communication (e.g. “FlashVar”)
	AdParameters *AdParameters `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`
}

// Creative is Wraps each creative element within an InLine or Wrapper Ad
type Creative struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// The preferred order in which multiple Creatives should be displayed
	Sequence int `xml:"sequence,attr,omitempty" json:"sequence,omitempty"`

	// Ad-ID for the creative (formerly ISCI)
	AdID string `xml:"AdID,attr,omitempty" json:"ad_id,omitempty"`

	Linear *Linear `xml:"Linear,omitempty" json:"linear,omitempty"`

	CompanionAds *CompanionAds `xml:"CompanionAds,omitempty" json:"companion_ads,omitempty"`

	NonLinearAds *NonLinearAds `xml:"NonLinearAds,omitempty" json:"non_linear_ads,omitempty"`
}

// CreativeExtension is Any valid XML may be included in the Extensions node
type CreativeExtension struct {
}

// CustomClick is URLs to request on custom events such as hotspotted video
type CustomClick struct {
	Value AnyURI `xml:",chardata" json:"value"`

	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// Extension is Any valid XML may be included in the Extensions node
type Extension struct {
}

//HTMLResource ...
type HTMLResource struct {
	Value string `xml:",chardata" json:"value"`

	// Specifies whether the HTML is XML-encoded
	XMLEncoded bool `xml:"xmlEncoded,attr,omitempty" json:"xml_encoded,omitempty"`
}

//Icon ...
type Icon struct {
	// Identifies the industry initiative that the icon supports.
	Program string `xml:"program,attr" json:"program"`

	// Pixel dimensions of icon.
	Width int `xml:"width,attr" json:"width"`

	// Pixel dimensions of icon.
	Height int `xml:"height,attr" json:"height"`

	// The horizontal alignment location (in pixels) or a specific alignment.
	XPosition string `xml:"xPosition,attr" json:"x_position"`

	// The vertical alignment location (in pixels) or a specific alignment.
	YPosition string `xml:"yPosition,attr" json:"y_position"`

	// Start time at which the player should display the icon. Expressed in standard time format hh:mm:ss.
	Offset *Time `xml:"offset,attr,omitempty" json:"offset,omitempty"`

	// The duration for which the player must display the icon. Expressed in standard time format hh:mm:ss.
	Duration *Time `xml:"duration,attr,omitempty" json:"duration,omitempty"`

	// The apiFramework defines the method to use for communication with the icon element
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:"StaticResource,omitempty" json:"static_resource,omitempty"`

	// URL source for an IFrame to display the companion element
	IFrameResource AnyURI `xml:"IFrameResource,omitempty" json:"i_frame_resource,omitempty"`

	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:"HTMLResource,omitempty" json:"html_resource,omitempty"`

	IconClicks *IconClicks `xml:"IconClicks,omitempty" json:"icon_clicks,omitempty"`

	// URLs to ping when icon is shown.
	IconViewTrackings []AnyURI `xml:"IconViewTracking,omitempty" json:"icon_view_trackings,omitempty"`
}

//IconClicks ...
type IconClicks struct {
	// URLs to ping when user clicks on the the icon.
	IconClickTrackings []AnyURI `xml:"IconClickTracking,omitempty" json:"icon_click_trackings,omitempty"`

	// URL to open as destination page when user clicks on the icon.
	IconClickThrough AnyURI `xml:"IconClickThrough,omitempty" json:"icon_click_through,omitempty"`
}

//Impression ...
type Impression struct {
	Value AnyURI `xml:",chardata" json:"value"`

	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

//InLine is Second-level element surrounding complete ad data for a single ad
type InLine struct {
	// Indicates source ad server
	AdSystem AdSystem `xml:"AdSystem" json:"ad_system"`

	// Common name of ad
	AdTitle string `xml:"AdTitle" json:"ad_title"`

	// Longer description of ad
	Description string `xml:"Description,omitempty" json:"description,omitempty"`

	// Common name of advertiser
	Advertiser string `xml:"Advertiser,omitempty" json:"advertiser,omitempty"`

	// The price of the ad.
	Pricing *Pricing `xml:"Pricing,omitempty" json:"pricing,omitempty"`

	// URL of request to survey vendor
	Survey AnyURI `xml:"Survey,omitempty" json:"survey,omitempty"`

	// URL to request if ad does not play due to error
	Error AnyURI `xml:"Error,omitempty" json:"error,omitempty"`

	Impressions []Impression `xml:"Impression" json:"impressions"`

	// Wraps each creative element within an InLine or Wrapper Ad
	Creatives []Creative `xml:"Creatives>Creative" json:"creatives"`

	// Any valid XML may be included in the Extensions node
	Extensions []Extension `xml:"Extensions>Extension,omitempty" json:"extensions,omitempty"`
}

//Linear ...
type Linear struct {
	// The time at which the ad becomes skippable, if absent, the ad is not skippable.
	Skipoffset string `xml:"skipoffset,attr,omitempty" json:"skipoffset,omitempty"`

	// Any number of icons representing advertising industry initiatives.
	Icons []Icon `xml:"Icons>Icon,omitempty" json:"icons,omitempty"`

	// Any valid XML may be included in the Extensions node
	CreativeExtensions []CreativeExtension `xml:"CreativeExtensions>CreativeExtension,omitempty" json:"creative_extensions,omitempty"`

	// Duration in standard time format, hh:mm:ss
	Duration Time `xml:"Duration" json:"duration"`

	// The name of the event to track for the element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// Data to be passed into the video ad
	AdParameters *AdParameters `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`

	VideoClicks *VideoClicks `xml:"VideoClicks,omitempty" json:"video_clicks,omitempty"`

	// Location of linear file
	MediaFiles []MediaFile `xml:"MediaFiles>MediaFile,omitempty" json:"media_files,omitempty"`
}

// MediaFile is Location of linear file
type MediaFile struct {
	Value AnyURI `xml:",chardata" json:"value"`

	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Method of delivery of ad
	Delivery MediaFileDelivery `xml:"delivery,attr" json:"delivery"`

	// MIME type. Popular MIME types include, but are not limited to “video/x-ms-wmv” for Windows Media, and “video/x-flv” for Flash Video. Image ads or interactive ads can be included in the MediaFiles section with appropriate Mime types
	Type string `xml:"type,attr" json:"type"`

	// Bitrate of encoded video in Kbps. If bitrate is supplied, minBitrate and maxBitrate should not be supplied.
	Bitrate int `xml:"bitrate,attr,omitempty" json:"bitrate,omitempty"`

	// Minimum bitrate of an adaptive stream in Kbps. If minBitrate is supplied, maxBitrate must be supplied and bitrate should not be supplied.
	MinBitrate int `xml:"minBitrate,attr,omitempty" json:"min_bitrate,omitempty"`

	// Maximum bitrate of an adaptive stream in Kbps. If maxBitrate is supplied, minBitrate must be supplied and bitrate should not be supplied.
	MaxBitrate int `xml:"maxBitrate,attr,omitempty" json:"max_bitrate,omitempty"`

	// Pixel dimensions of video
	Width int `xml:"width,attr" json:"width"`

	// Pixel dimensions of video
	Height int `xml:"height,attr" json:"height"`

	// Whether it is acceptable to scale the image.
	Scalable bool `xml:"scalable,attr,omitempty" json:"scalable,omitempty"`

	// Whether the ad must have its aspect ratio maintained when scales
	MaintainAspectRatio bool `xml:"maintainAspectRatio,attr,omitempty" json:"maintain_aspect_ratio,omitempty"`

	// The apiFramework defines the method to use for communication if the MediaFile is interactive. Suggested values for this element are “VPAID”, “FlashVars” (for Flash/Flex), “initParams” (for Silverlight) and “GetVariables” (variables placed in key/value pairs on the asset request).
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// The codec used to produce the media file.
	Codec string `xml:"codec,attr,omitempty" json:"codec,omitempty"`
}

//MediaFileDelivery Method of delivery of ad
type MediaFileDelivery string

const (
	//Streaming ...
	Streaming MediaFileDelivery = "streaming"
	//Progressive ...
	Progressive MediaFileDelivery = "progressive"
)

//NonLinear ...
type NonLinear struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Pixel dimensions of companion
	Width int `xml:"width,attr" json:"width"`

	// Pixel dimensions of companion
	Height int `xml:"height,attr" json:"height"`

	// Pixel dimensions of expanding nonlinear ad when in expanded state
	ExpandedWidth int `xml:"expandedWidth,attr,omitempty" json:"expanded_width,omitempty"`

	// Pixel dimensions of expanding nonlinear ad when in expanded state
	ExpandedHeight int `xml:"expandedHeight,attr,omitempty" json:"expanded_height,omitempty"`

	// Whether it is acceptable to scale the image.
	Scalable bool `xml:"scalable,attr,omitempty" json:"scalable,omitempty"`

	// Whether the ad must have its aspect ratio maintained when scales
	MaintainAspectRatio bool `xml:"maintainAspectRatio,attr,omitempty" json:"maintain_aspect_ratio,omitempty"`

	// Suggested duration to display non-linear ad, typically for animation to complete. Expressed in standard time format hh:mm:ss
	MinSuggestedDuration *Time `xml:"minSuggestedDuration,attr,omitempty" json:"min_suggested_duration,omitempty"`

	// The apiFramework defines the method to use for communication with the nonlinear element
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:"StaticResource,omitempty" json:"static_resource,omitempty"`

	// URL source for an IFrame to display the companion element
	IFrameResource AnyURI `xml:"IFrameResource,omitempty" json:"i_frame_resource,omitempty"`

	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:"HTMLResource,omitempty" json:"html_resource,omitempty"`

	// Any valid XML may be included in the Extensions node
	CreativeExtensions []CreativeExtension `xml:"CreativeExtensions>CreativeExtension,omitempty" json:"creative_extensions,omitempty"`

	// URLs to ping when user clicks on the the non-linear ad.
	NonLinearClickTrackings []AnyURI `xml:"NonLinearClickTracking,omitempty" json:"non_linear_click_trackings,omitempty"`

	// URL to open as destination page when user clicks on the non-linear ad unit.
	NonLinearClickThrough AnyURI `xml:"NonLinearClickThrough,omitempty" json:"non_linear_click_through,omitempty"`

	// Data to be passed into the video ad.
	AdParameters *AdParameters `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`
}

//NonLinearAds ...
type NonLinearAds struct {
	// The name of the event to track for the element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// Any number of companions in any desired pixel dimensions.
	NonLinears []NonLinear `xml:"NonLinear" json:"non_linears"`
}

//NonLinearWrapper ...
type NonLinearWrapper struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// Pixel dimensions of companion
	Width int `xml:"width,attr,omitempty" json:"width,omitempty"`

	// Pixel dimensions of companion
	Height int `xml:"height,attr,omitempty" json:"height,omitempty"`

	// Pixel dimensions of expanding nonlinear ad when in expanded state
	ExpandedWidth int `xml:"expandedWidth,attr,omitempty" json:"expanded_width,omitempty"`

	// Pixel dimensions of expanding nonlinear ad when in expanded state
	ExpandedHeight int `xml:"expandedHeight,attr,omitempty" json:"expanded_height,omitempty"`

	// Whether it is acceptable to scale the image.
	Scalable bool `xml:"scalable,attr,omitempty" json:"scalable,omitempty"`

	// Whether the ad must have its aspect ratio maintained when scales
	MaintainAspectRatio bool `xml:"maintainAspectRatio,attr,omitempty" json:"maintain_aspect_ratio,omitempty"`

	// Suggested duration to display non-linear ad, typically for animation to complete. Expressed in standard time format hh:mm:ss
	MinSuggestedDuration *Time `xml:"minSuggestedDuration,attr,omitempty" json:"min_suggested_duration,omitempty"`

	// The apiFramework defines the method to use for communication with the nonlinear element
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// Any valid XML may be included in the Extensions node
	CreativeExtensions []CreativeExtension `xml:"CreativeExtensions>CreativeExtension,omitempty" json:"creative_extensions,omitempty"`

	// URLs to ping when user clicks on the the non-linear ad unit.
	NonLinearClickTrackings []AnyURI `xml:"NonLinearClickTracking,omitempty" json:"non_linear_click_trackings,omitempty"`
}

// Pricing is The price of the ad.
type Pricing struct {
	Value float32 `xml:",chardata" json:"value"`

	// The pricing model used.
	Model PricingModel `xml:"model,attr" json:"model"`

	// The currency of the pricing.
	Currency string `xml:"currency,attr" json:"currency"`
}

//PricingModel The pricing model used.
type PricingModel string

const (
	//CPC ...
	CPC PricingModel = "cpc"
	//CPM ...
	CPM PricingModel = "cpm"
	//CPE ...
	CPE PricingModel = "cpe"
	//CPV ...
	CPV PricingModel = "cpv"
)

// StaticResource is URL to a static file, such as an image or SWF file
type StaticResource struct {
	Value AnyURI `xml:",chardata" json:"value"`

	// Mime type of static resource
	CreativeType string `xml:"creativeType,attr" json:"creative_type"`
}

// Tracking is The name of the event to track for the element. The creativeView should always be requested when present.
type Tracking struct {
	Value AnyURI `xml:",chardata" json:"value"`

	// The name of the event to track. For nonlinear ads these events should be recorded on the video within the ad.
	Event TrackingEvent `xml:"event,attr" json:"event"`

	// The time during the video at which this url should be pinged. Must be present for progress event.
	Offset string `xml:"offset,attr,omitempty" json:"offset,omitempty"`
}

//TrackingEvent The name of the event to track. For nonlinear ads these events should be recorded on the video within the ad.
type TrackingEvent string

const (
	//CreativeView ...
	CreativeView TrackingEvent = "creativeView"
	//Start ...
	Start TrackingEvent = "start"
	//FirstQuartile ...
	FirstQuartile TrackingEvent = "firstQuartile"
	//Midpoint ...
	Midpoint TrackingEvent = "midpoint"
	//ThirdQuartile ...
	ThirdQuartile TrackingEvent = "thirdQuartile"
	//Complete ...
	Complete TrackingEvent = "complete"
	//Mute ...
	Mute TrackingEvent = "mute"
	//Unmute ...
	Unmute TrackingEvent = "unmute"
	//Pause ...
	Pause TrackingEvent = "pause"
	//Rewind ...
	Rewind TrackingEvent = "rewind"
	//Resume ...
	Resume TrackingEvent = "resume"
	//Fullscreen ...
	Fullscreen TrackingEvent = "fullscreen"
	//ExitFullscreen ...
	ExitFullscreen TrackingEvent = "exitFullscreen"
	//Expand ...
	Expand TrackingEvent = "expand"
	//Collapse ...
	Collapse TrackingEvent = "collapse"
	//AcceptInvitation ...
	AcceptInvitation TrackingEvent = "acceptInvitation"
	//Close ...
	Close TrackingEvent = "close"
	//Skip ...
	Skip TrackingEvent = "skip"
	//Progress ...
	Progress TrackingEvent = "progress"
)

// VAST is IAB VAST, Video Ad Serving Template, video xml ad response, Version 3.0.0, xml schema prepared by Google
type VAST struct {
	// Current version is 3.0.
	Version string `xml:"version,attr" json:"version"`

	// Top-level element, wraps each ad in the response
	Ads []Ad `xml:"Ad,omitempty" json:"ads,omitempty"`
}

//VideoClicks ...
type VideoClicks struct {
	// URL to open as destination page when user clicks on the video
	ClickThrough *ClickThrough `xml:"ClickThrough,omitempty" json:"click_through,omitempty"`

	// URL to request for tracking purposes when user clicks on the video
	ClickTrackings []ClickTracking `xml:"ClickTracking,omitempty" json:"click_trackings,omitempty"`

	// URLs to request on custom events such as hotspotted video
	CustomClicks []CustomClick `xml:"CustomClick,omitempty" json:"custom_clicks,omitempty"`
}

// Wrapper is Second-level element surrounding wrapper ad pointing to Secondary ad server.
type Wrapper struct {
	// Indicates source ad server
	AdSystem AdSystem `xml:"AdSystem" json:"ad_system"`

	// URL of ad tag of downstream Secondary Ad Server
	VASTAdTagURI AnyURI `xml:"VASTAdTagURI" json:"vast_ad_tag_uri"`

	// URL to request if ad does not play due to error
	Error AnyURI `xml:"Error,omitempty" json:"error,omitempty"`

	// URL to request to track an impression
	Impressions []AnyURI `xml:"Impression" json:"impressions"`

	// Wraps each creative element within an InLine or Wrapper Ad
	Creatives []Creative `xml:"Creatives>Creative" json:"creatives"`

	// Any valid XML may be included in the Extensions node
	Extensions []Extension `xml:"Extensions>Extension,omitempty" json:"extensions,omitempty"`
}
