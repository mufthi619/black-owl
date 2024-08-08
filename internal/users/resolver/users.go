package resolver

import (
	"black-owl/internal/users/proto"
	"context"
)

type UserResolver struct {
}

func (s *UserResolver) Login(ctx context.Context, req *proto.LoginDTO) (finalResponse *proto.LoginResponseDTO, err error) {
	finalResponse.Message = "OK"
	finalResponse.Status = true
	finalResponse.Data = nil
	return
}
