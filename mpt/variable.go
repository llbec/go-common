package mpt

//tmlog "github.com/tendermint/tendermint/libs/log"

var (
//log tmlog.Logger
)

// IdealBatchSize Code using batches should try to add this much data to the batch.
// The value was determined empirically.
const IdealBatchSize = 100 * 1024
