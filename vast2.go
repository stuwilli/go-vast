package vast

// Ad is Top-level element, wraps each ad in the response
type Ad struct {
	ID string `xml:"id,attr" json:"id"`

	// Second-level element surrounding complete ad data for a single ad
	InLine *InLine `xml:"InLine,omitempty" json:"in_line,omitempty"`

	// Second-level element surrounding wrapper ad pointing to Secondary ad server.
	Wrapper *Wrapper `xml:"Wrapper,omitempty" json:"wrapper,omitempty"`
}

//AdSystem ...
type AdSystem struct {
	Value string `xml:",chardata" json:"value"`

	// Internal version used by ad system
	Version string `xml:"version,attr,omitempty" json:"version,omitempty"`
}

//AnyURI ...
type AnyURI string

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

	// Pixel dimensions of companion
	Width int `xml:"width,attr" json:"width"`

	// Pixel dimensions of companion
	Height int `xml:"height,attr" json:"height"`

	// Pixel dimensions of expanding companion ad when in expanded state
	ExpandedWidth int `xml:"expandedWidth,attr,omitempty" json:"expanded_width,omitempty"`

	// Pixel dimensions of expanding companion ad when in expanded state
	ExpandedHeight int `xml:"expandedHeight,attr,omitempty" json:"expanded_height,omitempty"`

	// The apiFramework defines the method to use for communication with the companion
	APIFramework string `xml:"apiFramework,attr,omitempty" json:"api_framework,omitempty"`

	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:"StaticResource,omitempty" json:"static_resource,omitempty"`

	// URL source for an IFrame to display the companion element
	IFrameResource AnyURI `xml:"IFrameResource,omitempty" json:"i_frame_resource,omitempty"`

	// HTML to display the companion element
	HTMLResource string `xml:"HTMLResource,omitempty" json:"html_resource,omitempty"`

	// The name of the event to track for the Linear element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// URL to open as destination page when user clicks on the the companion banner ad.
	CompanionClickThrough AnyURI `xml:"CompanionClickThrough,omitempty" json:"companion_click_through,omitempty"`

	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:"AltText,omitempty" json:"alt_text,omitempty"`

	// Data to be passed into the companion ads. The apiFramework defines the method to use for communication (e.g. “FlashVar”)
	AdParameters string `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`
}

// Creative is Wraps each creative element within an InLine or Wrapper Ad
type Creative struct {
	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`

	// The preferred order in which multiple Creatives should be displayed
	Sequence int `xml:"sequence,attr,omitempty" json:"sequence,omitempty"`

	// Ad-ID for the creative (formerly ISCI)
	AdID string `xml:"AdID,attr,omitempty" json:"ad_id,omitempty"`

	Linear *Linear `xml:"Linear,omitempty" json:"linear,omitempty"`

	// Any number of companions in any desired pixel dimensions.
	Companions []Companion `xml:"CompanionAds>Companion,omitempty" json:"companions,omitempty"`

	NonLinearAds *NonLinearAds `xml:"NonLinearAds,omitempty" json:"non_linear_ads,omitempty"`
}

// CustomClick is URLs to request on custom events such as hotspotted video
type CustomClick struct {
	Value AnyURI `xml:",chardata" json:"value"`

	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// Extension represent aribtrary XML provided by the platform to extend the VAST response
type Extension struct {
	Data []byte `xml:",innerxml"`
}

//Impression ...
type Impression struct {
	Value AnyURI `xml:",chardata" json:"value"`

	ID string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// InLine is Second-level element surrounding complete ad data for a single ad
type InLine struct {
	// Indicates source ad server
	AdSystem AdSystem `xml:"AdSystem" json:"ad_system"`

	// Common name of ad
	AdTitle string `xml:"AdTitle" json:"ad_title"`

	// Longer description of ad
	Description string `xml:"Description,omitempty" json:"description,omitempty"`

	// URL of request to survey vendor
	Survey AnyURI `xml:"Survey,omitempty" json:"survey,omitempty"`

	// URL to request if ad does not play due to error
	Error AnyURI `xml:"Error,omitempty" json:"error,omitempty"`

	// URL to track impression
	Impressions []Impression `xml:"Impression" json:"impressions"`

	// Wraps each creative element within an InLine or Wrapper Ad
	Creatives []Creative `xml:"Creatives>Creative" json:"creatives"`

	// Any valid XML may be included in the Extensions node
	Extensions []Extension `xml:"Extensions>Extension,omitempty" json:"extensions,omitempty"`
}

//Linear ...
type Linear struct {
	// Duration in standard time format, hh:mm:ss
	Duration Time `xml:"Duration" json:"duration"`

	// The name of the event to track for the Linear element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// Data to be passed into the video ad
	AdParameters string `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`

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

	// Bitrate of encoded video in Kbps
	Bitrate int `xml:"bitrate,attr,omitempty" json:"bitrate,omitempty"`

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
}

//MediaFileDelivery Method of delivery of ad
type MediaFileDelivery string

const (
	//Streaming ..
	Streaming MediaFileDelivery = "streaming"
	//Progressive ..
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
	HTMLResource string `xml:"HTMLResource,omitempty" json:"html_resource,omitempty"`

	// URL to open as destination page when user clicks on the the non-linear ad unit.
	NonLinearClickThrough AnyURI `xml:"NonLinearClickThrough,omitempty" json:"non_linear_click_through,omitempty"`

	// Data to be passed into the video ad.
	AdParameters string `xml:"AdParameters,omitempty" json:"ad_parameters,omitempty"`
}

//NonLinearAds ...
type NonLinearAds struct {
	// The name of the event to track for the Linear element. The creativeView should always be requested when present.
	Trackings []Tracking `xml:"TrackingEvents>Tracking,omitempty" json:"trackings,omitempty"`

	// Any number of companions in any desired pixel dimensions.
	NonLinears []NonLinear `xml:"NonLinear" json:"non_linears"`
}

// StaticResource is URL to a static file, such as an image or SWF file
type StaticResource struct {
	Value AnyURI `xml:",chardata" json:"value"`

	// Mime type of static resource
	CreativeType string `xml:"creativeType,attr" json:"creative_type"`
}

//Time ...
type Time string

// Tracking is The name of the event to track for the Linear element. The creativeView should always be requested when present.
type Tracking struct {
	Value AnyURI `xml:",chardata" json:"value"`

	// The name of the event to track. For nonlinear ads these events should be recorded on the video within the ad.
	Event TrackingEvent `xml:"event,attr" json:"event"`
}

//TrackingEvent The name of the event to track. For nonlinear ads these events should be recorded on the video within the ad.
type TrackingEvent string

const (
	//CreativeView ...
	CreativeView TrackingEvent = "creativeView"
	//Start ...
	Start TrackingEvent = "start"
	//Midpoint ...
	Midpoint TrackingEvent = "midpoint"
	//FirstQuartile ...
	FirstQuartile TrackingEvent = "firstQuartile"
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
	//Expand ...
	Expand TrackingEvent = "expand"
	//Collapse ...
	Collapse TrackingEvent = "collapse"
	//AcceptInvitation ...
	AcceptInvitation TrackingEvent = "acceptInvitation"
	//Close ...
	Close TrackingEvent = "close"
)

// VAST is IAB VAST, Video Ad Serving Template, video xml ad response, Version 2.0.1, xml schema prepared by Tremor Media
type VAST struct {
	// Current version is 2.0. The 2.0.1 version corrects an error in the Wrapper section related the Linear node's TrackingEvents and VideoCLicks elements.
	Version float32 `xml:"version,attr" json:"version"`

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
