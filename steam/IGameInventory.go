package steam

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"

	"github.com/Jleagle/unmarshal-go/ctypes"
)

func (s Steam) GetItemDefArchive(appID int, digest string) (archives []ItemDefArchive, bytes []byte, err error) {

	options := url.Values{}
	options.Set("appid", strconv.Itoa(appID))
	options.Set("digest", digest)

	bytes, err = s.getFromAPI("IGameInventory/GetItemDefArchive/v1", options)
	if err != nil {
		return archives, bytes, err
	}

	err = json.Unmarshal(bytes, &archives)
	return archives, bytes, err
}

type ItemDefArchive struct {
	AppID            ctypes.CInt   `json:"appid"`
	Bundle           string        `json:"bundle"`
	Commodity        bool          `json:"commodity"`
	DateCreated      time.Time     `json:"date_created"`
	Description      string        `json:"description"`
	DisplayType      string        `json:"display_type"`
	DropInterval     int           `json:"drop_interval"`
	DropMaxPerWindow int           `json:"drop_max_per_window"`
	Exchange         string        `json:"exchange"`
	Hash             string        `json:"hash"`
	IconURL          string        `json:"icon_url"`
	IconURLLarge     string        `json:"icon_url_large"`
	ItemdefID        ctypes.CInt   `json:"itemdefid"`
	ItemQuality      string        `json:"item_quality"`
	Marketable       bool          `json:"marketable"`
	Modified         time.Time     `json:"modified"`
	Name             string        `json:"name"`
	Price            string        `json:"price"`
	Promo            string        `json:"promo"`
	Quantity         int           `json:"quantity"`
	Tags             string        `json:"tags"`
	Timestamp        time.Time     `json:"Timestamp"`
	Tradable         bool          `json:"tradable"`
	Type             string        `json:"type"`
	WorkshopID       ctypes.CInt64 `json:"workshopid"`
}
