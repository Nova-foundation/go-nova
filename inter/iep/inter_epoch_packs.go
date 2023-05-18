package iep

import (
	"github.com/Nova-foundation/go-nova/inter"
	"github.com/Nova-foundation/go-nova/inter/ier"
)

type LlrEpochPack struct {
	Votes  []inter.LlrSignedEpochVote
	Record ier.LlrIdxFullEpochRecord
}
