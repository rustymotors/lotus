package shard

import (
	"testing"
)

func TestAddShard(t *testing.T) {
	repo := &ShardRepository{}

	shard := Shard{
		Id: "2",
		Name: "Shard 2",
		Description: "The second shard",
		LoginServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 8227,
		},
		LobbyServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 7004,
		},
		DatabaseServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 43301,
		},
		DiagnosticsServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 81,
		},
		Status: ShardStatus{
			Id: "1",
			Reason: "Active",
		},
		Group: "Group-2",
		Population: 10,
		MaxProfiles: 2,
	}

	repo.AddShard(shard)

	if len(repo.shards) != 1 {
		t.Errorf("expected 1 shard, got %d", len(repo.shards))
	}

	if repo.shards[0] != shard {
		t.Errorf("expected shard %v, got %v", shard, repo.shards[0])
	}
}

func TestGetShard(t *testing.T) {
	repo := &ShardRepository{}

	shard := Shard{
		Id: "2",
		Name: "Shard 2",
		Description: "The second shard",
		LoginServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 8227,
		},
		LobbyServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 7004,
		},
		DatabaseServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 43301,
		},
		DiagnosticsServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 81,
		},
		Status: ShardStatus{
			Id: "1",
			Reason: "Active",
		},
		Group: "Group-2",
		Population: 10,
		MaxProfiles: 2,
	}

	repo.AddShard(shard)

	retrievedShard, err := repo.GetShard("2")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if retrievedShard == nil || *retrievedShard != shard {
		t.Errorf("expected shard %v, got %v", shard, retrievedShard)
	}

	_, err = repo.GetShard("nonexistent")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGetAllShards(t *testing.T) {
	repo := &ShardRepository{}

	shard1 := Shard{
		Id: "1",
		Name: "Shard 1",
		Description: "The first shard",
		LoginServer: AddressPair{
			Ip: "10.10.5.20",
			Port: 8226,
		},
		LobbyServer: AddressPair{
			Ip: "10.10.5.20",
			Port: 7003,
		},
		DatabaseServer: AddressPair{
			Ip: "10.10.5.20",
			Port: 43300,
		},
		DiagnosticsServer: AddressPair{
			Ip: "10.10.5.20",
			Port: 80,
		},
		Status: ShardStatus{
			Id: "0",
			Reason: "",
		},
		Group: "Group-1",
		Population: 0,
		MaxProfiles: 1,
	}

	shard2 := Shard{
		Id: "2",
		Name: "Shard 2",
		Description: "The second shard",
		LoginServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 8227,
		},
		LobbyServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 7004,
		},
		DatabaseServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 43301,
		},
		DiagnosticsServer: AddressPair{
			Ip: "10.10.5.21",
			Port: 81,
		},
		Status: ShardStatus{
			Id: "1",
			Reason: "Active",
		},
		Group: "Group-2",
		Population: 10,
		MaxProfiles: 2,
	}

	repo.AddShard(shard1)
	repo.AddShard(shard2)

	shards := repo.GetAllShards()
	if len(shards) != 2 {
		t.Errorf("expected 2 shards, got %d", len(shards))
	}

	if shards[0] != shard1 || shards[1] != shard2 {
		t.Errorf("expected shards %v and %v, got %v and %v", shard1, shard2, shards[0], shards[1])
	}
}