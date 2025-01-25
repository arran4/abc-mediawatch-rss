package abcrss

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"time"
)

// RSS defines the structure of the RSS feed.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

// Channel represents the RSS channel.
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item represents an RSS feed item.
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Thumbnail   string `xml:"thumbnail"`
}

// ABCJSON represents the JSON structure extracted from the page.
type ABCJSON struct {
	Props struct {
		PageProps struct {
			Gtm                 bool `json:"gtm"`
			HeadTagsGTMPrepared struct {
				GtmContainerID string `json:"gtmContainerId"`
				Application    struct {
					Environment      string `json:"environment"`
					GeneratorName    string `json:"generatorName"`
					GeneratorVersion string `json:"generatorVersion"`
					Product          string `json:"product"`
				} `json:"application"`
			} `json:"headTagsGTMPrepared"`
			SnowplowPrepared struct {
				SnowplowCollectorURL string `json:"snowplowCollectorURL"`
				AppVersionID         string `json:"appVersionId"`
				AppID                string `json:"appId"`
				Environment          string `json:"environment"`
			} `json:"snowplowPrepared"`
			DeveloperFlags []any `json:"developerFlags"`
			ConnectData    struct {
				ConnectSection struct {
					ConnectSocialPrepared struct {
						PhoneServices  any `json:"phoneServices"`
						SocialServices []struct {
							CanonicalURL string `json:"canonicalURL"`
						} `json:"socialServices"`
						EmailServices any `json:"emailServices"`
					} `json:"connectSocialPrepared"`
					ConnectSubscribePrepared any `json:"connectSubscribePrepared"`
				} `json:"connectSection"`
				Footer struct {
					SocialServices []struct {
						CanonicalURL string `json:"canonicalURL"`
					} `json:"socialServices"`
					PhoneServices any `json:"phoneServices"`
					EmailServices any `json:"emailServices"`
					Newsletter    any `json:"newsletter"`
				} `json:"footer"`
			} `json:"connectData"`
			DynamicConfigController struct {
				PreparedProps struct {
					PreparedFeatureFlagsProps struct {
						ServerFlags struct {
							ShowItemsInTestHarness struct {
								Enabled bool `json:"enabled"`
							} `json:"showItemsInTestHarness"`
							RolloutEducationNewsNeighbourhood struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutEducationNewsNeighbourhood"`
							RolloutGardeningP2DetailsPage struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutGardeningP2DetailsPage"`
							RolloutHomepageSurveyBanner struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutHomepageSurveyBanner"`
							RolloutHomepageSurveyPanel struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutHomepageSurveyPanel"`
							RolloutPacificMoreLikeThisRecs struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutPacificMoreLikeThisRecs"`
							RolloutTriplejMoreLikeThisRecs struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutTriplejMoreLikeThisRecs"`
							ConfigureRUM struct {
								Enabled    bool `json:"enabled"`
								ReplayRate int  `json:"replayRate"`
								SampleRate int  `json:"sampleRate"`
							} `json:"configureRUM"`
							RolloutCareersEntryLevelProgram struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutCareersEntryLevelProgram"`
							RolloutGivesNewHomepageCollections struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutGivesNewHomepageCollections"`
							RolloutIncludeSearchPageFilters struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutIncludeSearchPageFilters"`
							RolloutListenMoreLikeThisRecs struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutListenMoreLikeThisRecs"`
							RolloutNewsLivePlayerLogo struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutNewsLivePlayerLogo"`
							RolloutRageNewsCarousel struct {
								Enabled bool `json:"enabled"`
							} `json:"rolloutRageNewsCarousel"`
							ShowABCmeSunsetMessaging struct {
								Enabled bool `json:"enabled"`
							} `json:"showABCmeSunsetMessaging"`
							ShowEverydaySunsetMessaging struct {
								Enabled bool `json:"enabled"`
							} `json:"showEverydaySunsetMessaging"`
						} `json:"serverFlags"`
					} `json:"preparedFeatureFlagsProps"`
					PreparedExperimentsProps struct {
						ExperimentsSingletonProviderData struct {
							FallbackEvaluations struct {
							} `json:"fallbackEvaluations"`
							CacheKey string `json:"cacheKey"`
							Remote   struct {
								CredentialRenewURL string `json:"credentialRenewURL"`
								Region             string `json:"region"`
								Project            string `json:"project"`
							} `json:"remote"`
						} `json:"experimentsSingletonProviderData"`
						ExperimentEvaluationsGroupData struct {
							Features []any  `json:"features"`
							AsPath   string `json:"asPath"`
						} `json:"experimentEvaluationsGroupData"`
					} `json:"preparedExperimentsProps"`
				} `json:"preparedProps"`
				DynamicPrepairController struct {
				} `json:"dynamicPrepairController"`
				DynamicPageConfigController struct {
				} `json:"dynamicPageConfigController"`
			} `json:"dynamicConfigController"`
			ExperimentsSingletonProviderData struct {
				FallbackEvaluations struct {
				} `json:"fallbackEvaluations"`
				CacheKey string `json:"cacheKey"`
				Remote   struct {
					CredentialRenewURL string `json:"credentialRenewURL"`
					Region             string `json:"region"`
					Project            string `json:"project"`
				} `json:"remote"`
			} `json:"experimentsSingletonProviderData"`
			ExperimentEvaluationsGroupData struct {
				Features []any  `json:"features"`
				AsPath   string `json:"asPath"`
			} `json:"experimentEvaluationsGroupData"`
			PreparedFeatureFlagsData struct {
				ServerFlags struct {
					ShowItemsInTestHarness struct {
						Enabled bool `json:"enabled"`
					} `json:"showItemsInTestHarness"`
					RolloutEducationNewsNeighbourhood struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutEducationNewsNeighbourhood"`
					RolloutGardeningP2DetailsPage struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutGardeningP2DetailsPage"`
					RolloutHomepageSurveyBanner struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutHomepageSurveyBanner"`
					RolloutHomepageSurveyPanel struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutHomepageSurveyPanel"`
					RolloutPacificMoreLikeThisRecs struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutPacificMoreLikeThisRecs"`
					RolloutTriplejMoreLikeThisRecs struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutTriplejMoreLikeThisRecs"`
					ConfigureRUM struct {
						Enabled    bool `json:"enabled"`
						ReplayRate int  `json:"replayRate"`
						SampleRate int  `json:"sampleRate"`
					} `json:"configureRUM"`
					RolloutCareersEntryLevelProgram struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutCareersEntryLevelProgram"`
					RolloutGivesNewHomepageCollections struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutGivesNewHomepageCollections"`
					RolloutIncludeSearchPageFilters struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutIncludeSearchPageFilters"`
					RolloutListenMoreLikeThisRecs struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutListenMoreLikeThisRecs"`
					RolloutNewsLivePlayerLogo struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutNewsLivePlayerLogo"`
					RolloutRageNewsCarousel struct {
						Enabled bool `json:"enabled"`
					} `json:"rolloutRageNewsCarousel"`
					ShowABCmeSunsetMessaging struct {
						Enabled bool `json:"enabled"`
					} `json:"showABCmeSunsetMessaging"`
					ShowEverydaySunsetMessaging struct {
						Enabled bool `json:"enabled"`
					} `json:"showEverydaySunsetMessaging"`
				} `json:"serverFlags"`
			} `json:"preparedFeatureFlagsData"`
			AppEnvironment     string `json:"appEnvironment"`
			DatadogRUMPrepared struct {
				ApplicationID string `json:"applicationId"`
				ClientToken   string `json:"clientToken"`
				Service       string `json:"service"`
				Env           string `json:"env"`
				Version       string `json:"version"`
			} `json:"datadogRUMPrepared"`
			HeadTagsRUMPrepared struct {
				TraceID   string `json:"traceId"`
				TraceTime int64  `json:"traceTime"`
			} `json:"headTagsRUMPrepared"`
			Analytics struct {
				Debug struct {
					SchemaVersion string `json:"schemaVersion"`
				} `json:"debug"`
				Application struct {
					Environment      string `json:"environment"`
					GeneratorName    string `json:"generatorName"`
					GeneratorVersion string `json:"generatorVersion"`
					Platform         string `json:"platform"`
					Product          string `json:"product"`
				} `json:"application"`
			} `json:"analytics"`
			TemplatePrepared struct {
				MastheadPrepared struct {
					NotFound              bool `json:"notFound"`
					IsProgramPage         bool `json:"isProgramPage"`
					MastheadBrandPrepared struct {
						NavigationData []struct {
							LinkTo   string `json:"linkTo"`
							Children string `json:"children"`
							Active   bool   `json:"active"`
						} `json:"navigationData"`
						BrandLockupPrepared struct {
							LogoHref         string `json:"logoHref"`
							LogoType         string `json:"logoType"`
							ScreenReaderText string `json:"screenReaderText"`
						} `json:"brandLockupPrepared"`
						MastheadCTA struct {
							IconScreenReaderText string `json:"iconScreenReaderText"`
							IconType             string `json:"iconType"`
							LinkTo               string `json:"linkTo"`
							Text                 string `json:"text"`
						} `json:"mastheadCTA"`
						IsCTAOutsideNav bool `json:"isCTAOutsideNav"`
					} `json:"mastheadBrandPrepared"`
					MastheadParentPrepared struct {
						SearchURL              string `json:"searchURL"`
						ShowSearch             bool   `json:"showSearch"`
						HideLogin              bool   `json:"hideLogin"`
						MastheadSearchPrepared struct {
							PluginOptions []struct {
								Type        string `json:"type"`
								HitsPerPage int    `json:"hitsPerPage"`
								Client      struct {
									AppID  string `json:"appId"`
									APIKey string `json:"apiKey"`
								} `json:"client,omitempty"`
								IndexName string `json:"indexName,omitempty"`
							} `json:"pluginOptions"`
						} `json:"mastheadSearchPrepared"`
					} `json:"mastheadParentPrepared"`
					Source string `json:"source"`
				} `json:"mastheadPrepared"`
				ProfilesProviderPrepared struct {
					Environment string `json:"environment"`
					APIKey      string `json:"APIKey"`
					AnalyticsID string `json:"analyticsID"`
				} `json:"profilesProviderPrepared"`
				NewsletterToasterPrepared any `json:"newsletterToasterPrepared"`
				SiteFooterPrepared        struct {
					NotFound        bool   `json:"notFound"`
					FullProductName string `json:"fullProductName"`
					Columns         []struct {
						Component      string `json:"component"`
						ComponentProps struct {
							Acknowledgement    string `json:"acknowledgement"`
							FooterLogoPrepared struct {
								LogoHref         string `json:"logoHref"`
								LogoType         string `json:"logoType"`
								ScreenReaderText string `json:"screenReaderText"`
							} `json:"footerLogoPrepared"`
							List []struct {
								LinkHref  string `json:"linkHref"`
								LinkTitle string `json:"linkTitle"`
							} `json:"list"`
							LinkType        string `json:"linkType"`
							FooterClassName string `json:"footerClassName"`
							SocialServices  []struct {
								CanonicalURL string `json:"canonicalURL"`
							} `json:"socialServices"`
						} `json:"componentProps,omitempty"`
						FooterColumnHeading string `json:"footerColumnHeading,omitempty"`
						LoaderParams        struct {
							SiteFooterLinksVariation string `json:"siteFooterLinksVariation"`
						} `json:"loaderParams,omitempty"`
					} `json:"columns"`
				} `json:"siteFooterPrepared"`
				NotificationMessagePrepared any    `json:"notificationMessagePrepared"`
				Title                       string `json:"title"`
				ShowDate                    bool   `json:"showDate"`
				ShowLocation                bool   `json:"showLocation"`
				Chromeless                  bool   `json:"chromeless"`
			} `json:"templatePrepared"`
			Data struct {
				ComponentsContent []struct {
					Key            string `json:"key"`
					Component      string `json:"component"`
					ComponentProps struct {
						ID              string `json:"id"`
						HeadingPrepared string `json:"headingPrepared"`
						Items           []struct {
							ArticleLink             string `json:"articleLink"`
							CardAttributionPrepared struct {
								PublishedDate       time.Time `json:"publishedDate"`
								PublishedDateFormat bool      `json:"publishedDateFormat"`
							} `json:"cardAttributionPrepared"`
							CardImagePrepared struct {
								Alt    string `json:"alt"`
								Ratio  string `json:"ratio"`
								SrcSet []any  `json:"srcSet"`
							} `json:"cardImagePrepared"`
							CardMediaIndicatorPrepared struct {
								Icon     string `json:"icon"`
								Duration bool   `json:"duration"`
							} `json:"cardMediaIndicatorPrepared"`
							ContentLabelPrepared  any    `json:"contentLabelPrepared"`
							ContentURI            string `json:"contentUri"`
							Description           string `json:"description"`
							CardID                string `json:"cardId"`
							CardTitle             string `json:"cardTitle"`
							NoBorders             bool   `json:"noBorders"`
							PresentersPrepared    any    `json:"presentersPrepared"`
							ImagePositionPrepared struct {
								Mobile  string `json:"mobile"`
								Tablet  string `json:"tablet"`
								Desktop string `json:"desktop"`
							} `json:"imagePositionPrepared"`
							CardContentPositionPrepared struct {
							} `json:"cardContentPositionPrepared"`
							HasMobileFeatured bool   `json:"hasMobileFeatured"`
							DocType           string `json:"docType"`
							Segments          []struct {
								ArticleLink             string `json:"articleLink"`
								CardAttributionPrepared struct {
									PublishedDate       time.Time `json:"publishedDate"`
									PublishedDateFormat bool      `json:"publishedDateFormat"`
								} `json:"cardAttributionPrepared"`
								CardImagePrepared struct {
									Alt    string   `json:"alt"`
									ImgSrc string   `json:"imgSrc"`
									Ratio  string   `json:"ratio"`
									SrcSet []string `json:"srcSet"`
									Width  string   `json:"width"`
									Height string   `json:"height"`
								} `json:"cardImagePrepared"`
								CardMediaIndicatorPrepared struct {
									Icon     string `json:"icon"`
									Duration bool   `json:"duration"`
								} `json:"cardMediaIndicatorPrepared"`
								ContentLabelPrepared struct {
									LabelText string `json:"labelText"`
								} `json:"contentLabelPrepared"`
								ContentURI            string `json:"contentUri"`
								Description           string `json:"description"`
								CardID                string `json:"cardId"`
								CardTitle             string `json:"cardTitle"`
								NoBorders             bool   `json:"noBorders"`
								PresentersPrepared    any    `json:"presentersPrepared"`
								ImagePositionPrepared struct {
									Mobile  string `json:"mobile"`
									Tablet  string `json:"tablet"`
									Desktop string `json:"desktop"`
								} `json:"imagePositionPrepared"`
								CardContentPositionPrepared struct {
								} `json:"cardContentPositionPrepared"`
								HasMobileFeatured bool `json:"hasMobileFeatured"`
							} `json:"segments"`
							ProgramTemplate string `json:"programTemplate"`
							Expanded        bool   `json:"expanded"`
						} `json:"items"`
						Analytics struct {
							URI           string `json:"uri"`
							ModuleURI     string `json:"moduleUri"`
							ContentSource string `json:"contentSource"`
							ContentType   string `json:"contentType"`
							ID            string `json:"id"`
							Title         struct {
								Title string `json:"title"`
							} `json:"title"`
							Items []struct {
								URI string `json:"uri"`
							} `json:"items"`
						} `json:"analytics"`
						Pagination struct {
							CollectionLoaderLimit int `json:"collectionLoaderLimit"`
							Offset                int `json:"offset"`
							Size                  int `json:"size"`
							Total                 int `json:"total"`
						} `json:"pagination"`
						ProgramID       any    `json:"programId"`
						LoadMoreURL     string `json:"loadMoreUrl"`
						ProgramTemplate string `json:"programTemplate"`
						HeadingMessage  string `json:"headingMessage"`
						Variant         string `json:"variant"`
						Label           string `json:"label"`
						Description     string `json:"description"`
						URL             string `json:"url"`
					} `json:"componentProps,omitempty"`
					CollectionPreparerParams struct {
						ShowDate bool `json:"showDate"`
					} `json:"collectionPreparerParams,omitempty"`
					Optional bool `json:"optional,omitempty"`
				} `json:"componentsContent"`
				Title       string `json:"title"`
				Description string `json:"description"`
				Analytics   struct {
					Document struct {
						ContentSource string `json:"contentSource"`
						ContentType   string `json:"contentType"`
						ID            string `json:"id"`
						URI           string `json:"uri"`
						CanonicalURL  string `json:"canonicalUrl"`
						Title         struct {
							Title string `json:"title"`
						} `json:"title"`
						Program struct {
							Name string `json:"name"`
							ID   string `json:"id"`
						} `json:"program"`
						SiteRoot struct {
							Segment string `json:"segment"`
							Title   string `json:"title"`
						} `json:"siteRoot"`
						Language string `json:"language"`
					} `json:"document"`
				} `json:"analytics"`
				ShowDate        bool `json:"showDate"`
				ShowLocation    bool `json:"showLocation"`
				ShowSmartBanner bool `json:"showSmartBanner"`
				Chromeless      bool `json:"chromeless"`
				HeadTagsSocial  struct {
				} `json:"headTagsSocial"`
			} `json:"data"`
			HeadTagsCMSPrepared struct {
				Generator string `json:"generator"`
				Site      string `json:"site"`
			} `json:"headTagsCMSPrepared"`
			HeadTagsPagePrepared struct {
				BasePath          string `json:"basePath"`
				FaviconPath       string `json:"faviconPath"`
				CanonicalURL      string `json:"canonicalURL"`
				Description       string `json:"description"`
				Keywords          []any  `json:"keywords"`
				Lang              string `json:"lang"`
				Robots            any    `json:"robots"`
				Title             string `json:"title"`
				TouchIconFilename string `json:"touchIconFilename"`
			} `json:"headTagsPagePrepared"`
			HeadTagsSocialPrepared struct {
				CanonicalURL string `json:"canonicalURL"`
				Description  string `json:"description"`
				Image        string `json:"image"`
				OgType       string `json:"ogType"`
				Title        string `json:"title"`
				Twitter      string `json:"twitter"`
				Site         string `json:"site"`
			} `json:"headTagsSocialPrepared"`
			AddAScreenReaderOnlyH1   bool `json:"addAScreenReaderOnlyH1"`
			ProfilesProviderPrepared struct {
				Environment string `json:"environment"`
				APIKey      string `json:"APIKey"`
				AnalyticsID string `json:"analyticsID"`
			} `json:"profilesProviderPrepared"`
		} `json:"pageProps"`
		NSsp bool `json:"__N_SSP"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		ProductSlug     string `json:"productSlug"`
		ProductPageSlug string `json:"productPageSlug"`
	} `json:"query"`
	BuildID               string `json:"buildId"`
	IsFallback            bool   `json:"isFallback"`
	IsExperimentalCompile bool   `json:"isExperimentalCompile"`
	Gssp                  bool   `json:"gssp"`
	AppGip                bool   `json:"appGip"`
	ScriptLoader          []any  `json:"scriptLoader"`
}

const BaseURL = "https://www.abc.net.au"

func FetchAndParseToRSS() (error, RSS) {
	resp, err := http.Get(BaseURL + "/mediawatch/episodes")
	if err != nil {
		return fmt.Errorf("fetching news to rss: %v", err), RSS{}
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Printf("Failed to close body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %v", resp.Status), RSS{}
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing news to rss: %v", err), RSS{}
	}

	rss := RSS{
		Version: "2.0",
		Channel: Channel{},
	}

	// Extract JSON data from a <script> tag.
	jsonData := ""
	doc.Find("script#__NEXT_DATA__").Each(func(i int, s *goquery.Selection) {
		jsonData = s.Text()
	})

	if jsonData == "" {
		return fmt.Errorf("no JSON data found"), RSS{}
	}

	var abcData ABCJSON
	if err := json.Unmarshal([]byte(jsonData), &abcData); err != nil {
		return fmt.Errorf("parsing JSON data: %v", err), RSS{}
	}

	// Extract feed header information
	rss.Channel.Title = abcData.Props.PageProps.HeadTagsSocialPrepared.Site
	rss.Channel.Link = abcData.Props.PageProps.HeadTagsSocialPrepared.CanonicalURL
	rss.Channel.Description = abcData.Props.PageProps.HeadTagsSocialPrepared.Description

	seenGUIDs := map[string]bool{}

	// Process components for items
	for _, component := range abcData.Props.PageProps.Data.ComponentsContent {
		if component.Component == "EpisodeCollection" {

			for _, listItem := range component.ComponentProps.Items {
				guid := BaseURL + listItem.ArticleLink
				if seenGUIDs[guid] {
					continue
				}
				seenGUIDs[guid] = true

				rss.Channel.Items = append(rss.Channel.Items, Item{
					Title:       listItem.CardTitle,
					Link:        guid,
					Description: listItem.Description,
					PubDate:     listItem.CardAttributionPrepared.PublishedDate.Format(time.RFC1123),
					GUID:        guid,
				})
			}
		}
	}

	return nil, rss
}
