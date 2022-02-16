package service

import (
	"context"
	"fmt"
	"sync"

	pb "demo-fieldmask/api/helloworld/v1"
	"demo-fieldmask/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// UserInfoService is a greeter service.
type UserInfoService struct {
	pb.UnimplementedUserInfoServer

	uc              *biz.GreeterUsecase
	userProfileRepo *UserProfileRepository
	log             *log.Helper
}

// NewGreeterService new a greeter service.
func NewUserInfoService(uc *biz.GreeterUsecase, logger log.Logger) *UserInfoService {
	return &UserInfoService{uc: uc, userProfileRepo: &UserProfileRepository{}, log: log.NewHelper(logger)}
}

func (s *UserInfoService) GetUserInfo(ctx context.Context, request *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	fmt.Printf("GetUserInfo is called: %s\n", request.String())
	fm := request.FieldMask_Filter()
	resp := &pb.GetUserInfoResponse{}

	wg := &sync.WaitGroup{}

	// todo get UserProfile from local database
	resp.ProfileInfo = &pb.UserProfile{
		UserId: 9999,
		Name:   "lining",
		Email:  "lining.guan@grabtaxi.com",
	}

	//if fm.MaskedOut_ProfileInfo() {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//
	//		// todo get UserProfile from local database
	//		resp.ProfileInfo = &pb.UserProfile{
	//			UserId: "69781",
	//			Name:   "lining",
	//			Email:  "lining.guan@grabtaxi.com",
	//		}
	//	}()
	//}

	if fm.MaskedOut_SocialInfo() {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// todo RPC call UserSocial service
			fmt.Println("get user social info from UserSocial service")
			resp.SocialInfo = &pb.UserSocial{
				Following: 100,
				Follower:  100,
			}
		}()
	}

	if fm.MaskedOut_VipInfo() {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// todo RPC call UserVip service
			fmt.Println("get user vip info from UserVip service")
			resp.VipInfo = &pb.UserVip{
				Level: 6,
				Exp:   50,
			}
		}()
	}

	wg.Wait()

	fm.Mask(resp)

	return resp, nil
}

func (s *UserInfoService) UpdateUserInfo(ctx context.Context, request *pb.UpdateUserInfoRequest) (*pb.UpdateUserInfoResponse, error) {
	fmt.Printf("UpdateUserInfo is called: %s\n", request.String())
	fm := request.FieldMask_Filter()

	profile := &pb.UserProfile{
		UserId: 9999,
		Name:   "lining.guan",
		Email:  "lining.guan@grabtaxi.com",
	}

	profileUpdateMap := map[string]interface{}{}
	if fm.MaskedIn_Name() {
		profileUpdateMap["name"] = request.Name
		profile.Name = request.Name
	}
	if fm.MaskedIn_Email() {
		profileUpdateMap["email"] = request.Email
		profile.Email = request.Email
	}

	_ = s.userProfileRepo.UpdateUserProfile(ctx, request.Id, profileUpdateMap)
	fmt.Printf("update user profile by %+v success\n", profileUpdateMap)

	return &pb.UpdateUserInfoResponse{
		ProfileInfo: profile,
	}, nil
}

type UserProfileRepository struct{}

func (u *UserProfileRepository) UpdateUserProfile(ctx context.Context, id int32, updateMap map[string]interface{}) error {
	return nil
}
