package rs

const {
	DATA_SHARDS = 4
	PARITY_SHARDS = 2
	ALL_SHARDS = DATA_SHARDS + PARITY_SHARDS
	BLOCK_PRE_SHARD = 8000
	BLOCK_SIZE = BLOCK_PRE_SHARD + DATA_SHARDS
}