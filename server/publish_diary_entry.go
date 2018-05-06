package server

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/suusan2go/familog"
	"github.com/suusan2go/familog/app/domain"
	"github.com/suusan2go/familog/app/usecase"
	context "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s FamilogServer) PublishDiaryEntry(ctx context.Context, r *familog.PublishDiaryEntryRequest) (*familog.PublishDiaryEntryResponse, error) {
	us := usecase.NewPublishDiaryEntryUsecase(s.Registry.DiaryEntryRepository())
	output, err := us.PublishDiaryEntry(usecase.PublishDiaryEntryInput{
		Body:     r.DiaryEntryForm.Body,
		Emoji:    r.DiaryEntryForm.Emoji,
		AuthorID: domain.UserID(0), // TODO
	})
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to publish Diary")
	}
	return &familog.PublishDiaryEntryResponse{
		DiaryEntry: &familog.DiaryEntry{
			Id:            int32(output.DiaryEntry.ID),
			Body:          output.DiaryEntry.Body,
			Emoji:         output.DiaryEntry.Emoji,
			CoverImageUrl: output.DiaryEntry.PrimaryImage(),
			ImageUrls:     output.DiaryEntry.Images,
			PublishedAt: &timestamp.Timestamp{
				Seconds: int64(output.DiaryEntry.CreatedAt.Second()),
				Nanos:   int32(output.DiaryEntry.CreatedAt.Nanosecond()),
			},
		},
	}, nil
}
