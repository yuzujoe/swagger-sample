/*
 * inazma-sample
 *
 * this api practice inazma
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

type Reservation struct {
	Id int64 `json:"id,omitempty"`

	ReservationDate time.Time `json:"reservationDate,omitempty"`

	ReservationTime time.Time `json:"reservationTime,omitempty"`

	ReservationType int8 `json:"reservationType,omitempty"`

	Purpose string `json:"purpose,omitempty"`
}