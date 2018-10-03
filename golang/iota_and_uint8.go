package main

// Snippets from https://github.com/tuna/tunasync

// Use uint8 as small capacity int.
// Use iota to define some same variable type.
type SyncStatus uint8

const (
	None SyncStatus = iota
	Failed
	Success
	Syncing
	PreSyncing
	Paused
	Disabled
)

func (s SyncStatus) String() string {
	switch s {
	case None:
		return "none"
	case Failed:
		return "failed"
	case Success:
		return "success"
	case Syncing:
		return "syncing"
	case PreSyncing:
		return "pre-syncing"
	case Paused:
		return "paused"
	case Disabled:
		return "disabled"
	default:
		return ""
	}
}
