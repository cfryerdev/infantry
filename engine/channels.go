package engine

import (
	"infantry/bindings"
)

type DataChannel chan bindings.DataEvent

type DataChannelSlice []DataChannel
