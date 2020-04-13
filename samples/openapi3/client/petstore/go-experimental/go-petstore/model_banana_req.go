/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
)

// BananaReq struct for BananaReq
type BananaReq struct {
	LengthCm float32 `json:"lengthCm"`
	Sweet *bool `json:"sweet,omitempty"`
}

// NewBananaReq instantiates a new BananaReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBananaReq(lengthCm float32, ) *BananaReq {
	this := BananaReq{}
	this.LengthCm = lengthCm
	return &this
}

// NewBananaReqWithDefaults instantiates a new BananaReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBananaReqWithDefaults() *BananaReq {
	this := BananaReq{}
	return &this
}

// GetLengthCm returns the LengthCm field value
func (o *BananaReq) GetLengthCm() float32 {
	if o == nil  {
		var ret float32
		return ret
	}

	return o.LengthCm
}

// GetLengthCmOk returns a tuple with the LengthCm field value
// and a boolean to check if the value has been set.
func (o *BananaReq) GetLengthCmOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LengthCm, true
}

// SetLengthCm sets field value
func (o *BananaReq) SetLengthCm(v float32) {
	o.LengthCm = v
}

// GetSweet returns the Sweet field value if set, zero value otherwise.
func (o *BananaReq) GetSweet() bool {
	if o == nil || o.Sweet == nil {
		var ret bool
		return ret
	}
	return *o.Sweet
}

// GetSweetOk returns a tuple with the Sweet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BananaReq) GetSweetOk() (*bool, bool) {
	if o == nil || o.Sweet == nil {
		return nil, false
	}
	return o.Sweet, true
}

// HasSweet returns a boolean if a field has been set.
func (o *BananaReq) HasSweet() bool {
	if o != nil && o.Sweet != nil {
		return true
	}

	return false
}

// SetSweet gets a reference to the given bool and assigns it to the Sweet field.
func (o *BananaReq) SetSweet(v bool) {
	o.Sweet = &v
}

func (o BananaReq) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["lengthCm"] = o.LengthCm
	}
	if o.Sweet != nil {
		toSerialize["sweet"] = o.Sweet
	}
	return json.Marshal(toSerialize)
}

// AsFruitReq wraps this instance of BananaReq in FruitReq
func (s *BananaReq) AsFruitReq() FruitReq {
	return FruitReq{ FruitReqInterface: s }
}
type NullableBananaReq struct {
	value *BananaReq
	isSet bool
}

func (v NullableBananaReq) Get() *BananaReq {
	return v.value
}

func (v *NullableBananaReq) Set(val *BananaReq) {
	v.value = val
	v.isSet = true
}

func (v NullableBananaReq) IsSet() bool {
	return v.isSet
}

func (v *NullableBananaReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBananaReq(val *BananaReq) *NullableBananaReq {
	return &NullableBananaReq{value: val, isSet: true}
}

func (v NullableBananaReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBananaReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}