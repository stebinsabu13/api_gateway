package note

import (
	"fmt"

	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/config"
	"github.com/stebinsabu13/note_taking_microservice/api_gateway/pkg/note/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.NoteServiceClient
}

func InitServCli(c *config.Config) pb.NoteServiceClient {
	cc, err := grpc.Dial(c.NoteSuvUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewNoteServiceClient(cc)
}
