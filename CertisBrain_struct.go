package CertisBrain_struct

type CertisBrain_struct struct {
	Event_id	string
	Sensor_id	string
	Timestamp	int
	DateTime	string // 存 2020-05-06T04:33:35.739Z 格式的, 要用自己轉

	Global_id_Index		map[string]int	// $$ map[global_id]index <-- ReID 可以順便建, 方便之後worker 可以直接用
	Objects_id_Index	map[int]int	// $$ map[object_id]index <--方便之後worker 可以直接用
	Objects		[]Object_struct // $$ map[object_id]
}

type Object_struct struct {
	Global_id	string
	Object_id	int
	Object_name	string
	Probability			float64

	Pixel_Coordinate	Bounding_box_struct
	Geo_Coordinate		Geo_struct	 

	All_Tags				Tag_struct
	
}



type Geo_struct struct {
	Latitude	float64	// $$ 緯度
	Longitude	 float64 // $$ 經度
	Geohash		string 
}

type Tag_struct struct {
	Basic_Poses_tags		[]string
	
	// $$ 為什麼要map? 因為互相比較 feature的時候, 就不需要把 slice從頭到尾掃一遍才知道他有哪些feature
	Basic_Features			map[string]*Feature_struct

//	Relation_tags		map[string]interface{} // $$ 目前有 Wearable_tags 和 POI_tags

	WearableTag			[]Wearable_struct
	POI_tags 			[]string

	Compound_tags		[]string
	Action_tags			[]string
}

type Feature_struct struct {
	Data_type	string // $$ 這個欄位決定用什麼方式 parse 底下的 value
	Value	interface{}
}

/*
type Wearable_struct struct {
	Object_id	int
	Object_name	string
	Data_type	string 	// $$ 這個欄位決定用什麼方式 parse 底下的 OtherAttributes
	OtherAttributes	interface{}
	// $$ 看未來還需要擴充什麼
}
*/

type Bounding_box_struct struct {
	X		float64	`json:"x"`
	Y		float64	`json:"y"`
	Width	float64	`json:"width"`
	Height	float64	`json:"height"`
}


type Wearable_struct struct {
	Object_id	int
	Object_name	string
	RelationType string
    // $$ 從 Wearable 搬過來
}