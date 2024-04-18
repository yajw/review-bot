package submit_txn_review

type Request struct {
	UID           int64                  `json:"uid,omitempty"`
	SceneKey      string                 `json:"scene_key,omitempty" json:"scene_key,omitempty"`
	SceneID       int64                  `json:"scene_id,omitempty" json:"scene_id,omitempty"`
	ReviewContent string                 `json:"review_content,omitempty" json:"review_content,omitempty"`
	ExtraAttrs    map[string]interface{} `json:"extra_attrs,omitempty" json:"extra_attrs,omitempty"`
}
