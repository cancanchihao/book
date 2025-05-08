package server

import (
	"context"

	"book/library/rpc/internal/logic"
	"book/library/rpc/internal/svc"
	"book/library/rpc/library"
)

type LibraryServer struct {
	svcCtx *svc.ServiceContext
	library.UnimplementedLibraryServer
}

func NewLibraryServer(svcCtx *svc.ServiceContext) *LibraryServer {
	return &LibraryServer{
		svcCtx: svcCtx,
	}
}

// FindBookByName 通过书籍名称查找书籍
func (s *LibraryServer) FindBookByName(ctx context.Context, in *library.FindBookReq) (*library.FindBookReply, error) {
	l := logic.NewFindBookByNameLogic(ctx, s.svcCtx)
	return l.FindBookByName(in)
}
