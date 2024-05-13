package virtualization

import (
	"encoding/json"
	"net/url"
)

type VDisk struct {
	ID         string `url:"-" json:"vdisk_id,omitempty"`
	ImageName  string `url:"image_name" json:"image_name,omitempty"`
	ImageID    string `url:"image_id" json:"image_id,omitempty"`
	CreateType int64  `url:"create_type" json:"create_type"`
	Size       int64  `url:"vdisk_size" json:"vdisk_size,omitempty"`
	Controller int64  `url:"-" json:"controller,omitempty"`
	Unmap      bool   `url:"-" json:"unmap,omitempty"`
}

type VNIC struct {
	ID     string `url:"network_id" json:"network_id,omitempty"`
	Name   string `url:"network_name" json:"network_name,omitempty"`
	Mac    string `url:"mac" json:"mac,omitempty"`
	Model  int64  `url:"-" json:"model,omitempty"`
	VnicID string `url:"-" json:"vnic_id,omitempty"`
}

type VDisks []VDisk
type VNICs []VNIC

func (s VDisks) EncodeValues(k string, v *url.Values) error {

	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}
	v.Set(k, string(encoded))

	return nil
}

func (s VNICs) EncodeValues(k string, v *url.Values) error {

	encoded, err := json.Marshal(s)
	if err != nil {
		return err
	}
	v.Set(k, string(encoded))

	return nil
}

type Guest struct {
	ID          string `url:"guest_id,omitempty" json:"guest_id"`
	Name        string `url:"guest_name,omitempty" json:"guest_name"`
	Description string `url:"description,omitempty" json:"description"`
	Status      string `url:"status,omitempty" json:"status"`
	StorageID   string `url:"storage_id,omitempty" json:"storage_id"`
	StorageName string `url:"storage_name,omitempty" json:"storage_name"`
	AutoRun     int64  `url:"autorun,omitempty" json:"autorun"`
	VcpuNum     int64  `url:"vcpu_num,omitempty" json:"vcpu_num"`
	VramSize    int64  `url:"vram_size,omitempty" json:"vram_size"`
	Disks       VDisks `url:"vdisks,omitempty" json:"vdisks"`
	Networks    VNICs  `url:"vnics,omitempty" json:"vnics"`
}

type GuestList struct {
	Guests []Guest `url:"guests" json:"guests"`
}

type GetGuest struct {
	Name string `form:"guest_name" url:"guest_name"`
}
