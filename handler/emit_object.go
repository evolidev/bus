package handler

import "time"

type EmitObject interface {
	Time() time.Time
}
