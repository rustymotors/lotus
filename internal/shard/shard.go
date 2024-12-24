package shard

import (
	"fmt"
	"net/http"
)

type AddressPair struct {
	Ip string `json:"ip"`
	Port int `json:"port"`
}

type ShardStatus struct {
	Id string `json:"id"`
	Reason string `json:"reason"`
}

type Shard struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	LoginServer AddressPair `json:"loginServer"`
	LobbyServer AddressPair `json:"lobbyServer"`
	DatabaseServer AddressPair `json:"databaseServer"`
	DiagnosticsServer AddressPair `json:"diagnosticsServer"`
	Status ShardStatus `json:"status"`
	Group string `json:"group"`
	Population int `json:"population"`
	MaxProfiles int `json:"maxProfilesAllowed"`
}

func (s Shard) String() string {
	return "[" + s.Name + "]\n" +
	"Description=" + s.Description + "\n" +
	"ShardId=" + s.Id + "\n" +
	"LoginServerIP=" + s.LoginServer.Ip + "\n" +
	"LoginServerPort=" + fmt.Sprint(s.LoginServer.Port) + "\n" +
	"LobbyServerIP=" + s.LobbyServer.Ip + "\n" +
	"LobbyServerPort=" + fmt.Sprint(s.LobbyServer.Port) + "\n" +
	"MCOTSServerIP=" + s.DatabaseServer.Ip + "\n" +
	"StatusId=" + s.Status.Id + "\n" +
	"Status_Reason=" + s.Status.Reason + "\n" +
	"ServerGroup_Name=" + s.Group + "\n" +
	"Population=" + fmt.Sprint(s.Population) + "\n" +
	"MaxPersonasPerUser=" + fmt.Sprint(s.MaxProfiles) + "\n" +
	"DiagnosticServerHost=" + s.DiagnosticsServer.Ip + "\n" +
	"DiagnosticServerPort=" + fmt.Sprint(s.DiagnosticsServer.Port) + "\n"
}

type ShardRepository struct {
	shards []Shard
}

func (r *ShardRepository) String() string {
	var result string
	for _, shard := range r.shards {
		result += shard.String()
	}
	return result
}

func (r *ShardRepository) GetShard(id string) (*Shard, error) {
	for _, shard := range r.shards {
		if shard.Id == id {
			return &shard, nil
		}
	}
	return nil, fmt.Errorf("shard not found")
}

func (r *ShardRepository) GetAllShards() []Shard {
	return r.shards
}

func (r *ShardRepository) AddShard(shard Shard) {
	r.shards = append(r.shards, shard)
}

func (r *ShardRepository) init() {
	selfServerHost := "10.10.5.20"

	r.AddShard(Shard{
			Id: "1",
			Name: "Shard 1",
			Description: "The first shard",
			LoginServer: AddressPair{
				Ip: selfServerHost,
				Port: 8226,
			},
			LobbyServer: AddressPair{
				Ip: selfServerHost,
				Port: 7003,
			},
			DatabaseServer: AddressPair{
				Ip: selfServerHost,
				Port: 43300,
			},
			DiagnosticsServer: AddressPair{
				Ip: selfServerHost,
				Port: 80,
			},
			Status: ShardStatus{
				Id: "0",
				Reason: "",
			},
			Group: "Group-1",
			Population: 0,
			MaxProfiles: 1,
		},
	)
}

var (
	instance *ShardRepository
)

func FetchShardRepository() *ShardRepository {
	if instance == nil {
		instance = &ShardRepository{}
		instance.init()
	}
	return instance
}

func HandleShardList(r *http.Request, w http.ResponseWriter) {
	fmt.Println("ShardList")

	response := FetchShardRepository().GetAllShards()

	w.Header().Set("Content-Type", "text/plain")
	for _, shard := range response {
		fmt.Fprintf(w, "%v\n", shard)
	}

}