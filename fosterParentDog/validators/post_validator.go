package validators

type PostValidator struct {
	Publishing           int         `json:"publishing"`      //公開設定
	DogName              string      `json:"dog_name"`        //犬の名前
	Breed                string      `json:"breed"`           //犬種
	Gender               int         `json:"gender"`          //性別
	Spay                 int         `json:"spay"`            //去勢/避妊手術
	Old                  string      `json:"old"`             //年齢
	SinglePerson         int         `json:"single_person"`   //単身者への譲渡
	SeniorPerson         int         `json:"senior_person"`   //高齢者への譲渡
	TransferStatus       int         `json:"transfer_status"` //譲渡ステータス
	Introduction         string      `json:"introduction"`    //犬の自己紹介
	AppealPoint          string      `json:"appeal_point"`    //性格アピールポイント
	OtherMessage         string      `json:"other_message"`   //健康状態や譲渡条件などの特記事項
	UserId               uint        `json:"user_id"`         //ユーザーID
	ImagePathList        []ImagePath `json:"image_path_list"`
	PostPrefectureIdList []int       `json:"post_prefecture_id_list"`
}
type ImagePath struct {
	Position  int    `json:"position"`
	ImagePath string `json:"image_path"`
}
