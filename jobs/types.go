package jobs

type GetWeiboFeedResponse struct {
	Ok   int `json:"ok"`
	Data struct {
		CardlistInfo struct {
			Containerid string `json:"containerid"`
			VP          int    `json:"v_p"`
			ShowStyle   int    `json:"show_style"`
			Total       int    `json:"total"`
			SinceID     int64  `json:"since_id"`
		} `json:"cardlistInfo"`
		Cards []struct {
			CardType int    `json:"card_type"`
			Itemid   string `json:"itemid"`
			Scheme   string `json:"scheme"`
			Mblog    struct {
				Visible struct {
					Type   int `json:"type"`
					ListID int `json:"list_id"`
				} `json:"visible"`
				CreatedAt                string `json:"created_at"`
				ID                       string `json:"id"`
				Idstr                    string `json:"idstr"`
				Mid                      string `json:"mid"`
				CanEdit                  bool   `json:"can_edit"`
				ShowAdditionalIndication int    `json:"show_additional_indication"`
				Text                     string `json:"text"`
				TextLength               int    `json:"textLength"`
				Source                   string `json:"source"`
				Favorited                bool   `json:"favorited"`
				PicTypes                 string `json:"pic_types"`
				ThumbnailPic             string `json:"thumbnail_pic"`
				BmiddlePic               string `json:"bmiddle_pic"`
				OriginalPic              string `json:"original_pic"`
				IsPaid                   bool   `json:"is_paid"`
				MblogVipType             int    `json:"mblog_vip_type"`
				User                     struct {
					ID              int64  `json:"id"`
					ScreenName      string `json:"screen_name"`
					ProfileImageURL string `json:"profile_image_url"`
					ProfileURL      string `json:"profile_url"`
					StatusesCount   int    `json:"statuses_count"`
					Verified        bool   `json:"verified"`
					VerifiedType    int    `json:"verified_type"`
					CloseBlueV      bool   `json:"close_blue_v"`
					Description     string `json:"description"`
					Gender          string `json:"gender"`
					Mbtype          int    `json:"mbtype"`
					Urank           int    `json:"urank"`
					Mbrank          int    `json:"mbrank"`
					FollowMe        bool   `json:"follow_me"`
					Following       bool   `json:"following"`
					FollowersCount  int    `json:"followers_count"`
					FollowCount     int    `json:"follow_count"`
					CoverImagePhone string `json:"cover_image_phone"`
					AvatarHd        string `json:"avatar_hd"`
					Like            bool   `json:"like"`
					LikeMe          bool   `json:"like_me"`
					Badge           struct {
						BindTaobao          int `json:"bind_taobao"`
						UnreadPool          int `json:"unread_pool"`
						UnreadPoolExt       int `json:"unread_pool_ext"`
						CzWed2017           int `json:"cz_wed_2017"`
						Panda               int `json:"panda"`
						UserNameCertificate int `json:"user_name_certificate"`
						Hongbao2020         int `json:"hongbao_2020"`
					} `json:"badge"`
				} `json:"user"`
				PicStatus             string `json:"picStatus"`
				RepostsCount          int    `json:"reposts_count"`
				CommentsCount         int    `json:"comments_count"`
				AttitudesCount        int    `json:"attitudes_count"`
				PendingApprovalCount  int    `json:"pending_approval_count"`
				IsLongText            bool   `json:"isLongText"`
				RewardExhibitionType  int    `json:"reward_exhibition_type"`
				HideFlag              int    `json:"hide_flag"`
				Mlevel                int    `json:"mlevel"`
				Mblogtype             int    `json:"mblogtype"`
				MoreInfoType          int    `json:"more_info_type"`
				ExternSafe            int    `json:"extern_safe"`
				NumberDisplayStrategy struct {
					ApplyScenarioFlag    int    `json:"apply_scenario_flag"`
					DisplayTextMinNumber int    `json:"display_text_min_number"`
					DisplayText          string `json:"display_text"`
				} `json:"number_display_strategy"`
				ContentAuth       int `json:"content_auth"`
				PicNum            int `json:"pic_num"`
				MblogMenuNewStyle int `json:"mblog_menu_new_style"`
				EditConfig        struct {
					Edited bool `json:"edited"`
				} `json:"edit_config"`
				WeiboPosition   int    `json:"weibo_position"`
				ShowAttitudeBar int    `json:"show_attitude_bar"`
				Bid             string `json:"bid"`
				Pics            []struct {
					Pid  string `json:"pid"`
					URL  string `json:"url"`
					Size string `json:"size"`
					Geo  struct {
						Width  int  `json:"width"`
						Height int  `json:"height"`
						Croped bool `json:"croped"`
					} `json:"geo"`
					Large struct {
						Size string `json:"size"`
						URL  string `json:"url"`
						Geo  struct {
							Width  interface{} `json:"width"`
							Height interface{} `json:"height"`
							Croped bool        `json:"croped"`
						} `json:"geo"`
					} `json:"large"`
				} `json:"pics"`
			} `json:"mblog,omitempty"`
			ShowType int `json:"show_type"`
		} `json:"cards"`
		Banners interface{} `json:"banners"`
		Scheme  string      `json:"scheme"`
	} `json:"data"`
}
