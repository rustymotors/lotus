package shard

import (
	"testing"
)

func TestAddShard(t *testing.T) {
	repo := &ShardRepository{}

	shard := Shard{
		Id:          "2",
		Name:        "Shard 2",
		Description: "The second shard",
		LoginServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 8227,
		},
		LobbyServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 7004,
		},
		DatabaseServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 43301,
		},
		DiagnosticsServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 81,
		},
		Status: ShardStatus{
			Id:     "1",
			Reason: "Active",
		},
		Group:       "Group-2",
		Population:  10,
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
		Id:          "2",
		Name:        "Shard 2",
		Description: "The second shard",
		LoginServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 8227,
		},
		LobbyServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 7004,
		},
		DatabaseServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 43301,
		},
		DiagnosticsServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 81,
		},
		Status: ShardStatus{
			Id:     "1",
			Reason: "Active",
		},
		Group:       "Group-2",
		Population:  10,
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
		Id:          "1",
		Name:        "Shard 1",
		Description: "The first shard",
		LoginServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 8226,
		},
		LobbyServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 7003,
		},
		DatabaseServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 43300,
		},
		DiagnosticsServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 80,
		},
		Status: ShardStatus{
			Id:     "0",
			Reason: "",
		},
		Group:       "Group-1",
		Population:  0,
		MaxProfiles: 1,
	}

	shard2 := Shard{
		Id:          "2",
		Name:        "Shard 2",
		Description: "The second shard",
		LoginServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 8227,
		},
		LobbyServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 7004,
		},
		DatabaseServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 43301,
		},
		DiagnosticsServer: AddressPair{
			Ip:   "10.10.5.21",
			Port: 81,
		},
		Status: ShardStatus{
			Id:     "1",
			Reason: "Active",
		},
		Group:       "Group-2",
		Population:  10,
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
func TestShardString(t *testing.T) {
	shard := Shard{
		Id:          "1",
		Name:        "Shard 1",
		Description: "The first shard",
		LoginServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 8226,
		},
		LobbyServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 7003,
		},
		DatabaseServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 43300,
		},
		DiagnosticsServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 80,
		},
		Status: ShardStatus{
			Id:     "0",
			Reason: "",
		},
		Group:       "Group-1",
		Population:  0,
		MaxProfiles: 1,
	}

	expected := `[Shard 1]
Description=The first shard
ShardId=1
LoginServerIP=10.10.5.20
LoginServerPort=8226
LobbyServerIP=10.10.5.20
LobbyServerPort=7003
MCOTSServerIP=10.10.5.20
StatusId=0
Status_Reason=
ServerGroup_Name=Group-1
Population=0
MaxPersonasPerUser=1
DiagnosticServerHost=10.10.5.20
DiagnosticServerPort=80
`

	if shard.String() != expected {
		t.Errorf("expected %v, got %v", expected, shard.String())
	}
}
func TestFetchShardRepository(t *testing.T) {
	repo := FetchShardRepository()

	if repo == nil {
		t.Errorf("expected non-nil repository, got nil")
	}

	if len(repo.shards) != 1 {
		t.Errorf("expected 1 shard, got %d", len(repo.shards))
	}

	expectedShard := Shard{
		Id:          "1",
		Name:        "Shard 1",
		Description: "The first shard",
		LoginServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 8226,
		},
		LobbyServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 7003,
		},
		DatabaseServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 43300,
		},
		DiagnosticsServer: AddressPair{
			Ip:   "10.10.5.20",
			Port: 80,
		},
		Status: ShardStatus{
			Id:     "0",
			Reason: "",
		},
		Group:       "Group-1",
		Population:  0,
		MaxProfiles: 1,
	}

	if repo.shards[0] != expectedShard {
		t.Errorf("expected shard %v, got %v", expectedShard, repo.shards[0])
	}

	// Test singleton behavior
	repo2 := FetchShardRepository()
	if repo != repo2 {
		t.Errorf("expected same repository instance, got different instances")
	}
}
