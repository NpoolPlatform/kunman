package github

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"

	"github.com/NpoolPlatform/kunman/framework/logger"
	constant "github.com/NpoolPlatform/kunman/mal/third/const"
	npool "github.com/NpoolPlatform/kunman/message/third/middleware/v1/oauth"
	"github.com/go-resty/resty/v2"
)

func GetAccessToken(clientID, clientSecret, code string) (*npool.AccessTokenInfo, error) {
	var err error
	reqURL := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s",
		constant.GithubGetAccessTokenURL, clientID, clientSecret, code,
	)
	socksProxy := os.Getenv("ENV_REQUEST_PROXY")

	cli := resty.New()

	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	var resp *resty.Response
	resp, err = cli.R().
		Get(reqURL)
	if err != nil {
		logger.Sugar().Error("req err:", err)
		return nil, err
	}

	if !resp.IsSuccess() {
		logger.Sugar().Error("resp err: ", resp.StatusCode())
		return nil, err
	}

	queryParams, err := url.ParseQuery(resp.String())
	if err != nil {
		logger.Sugar().Error("decode param err: ", err)
		return nil, err
	}

	errStr := queryParams.Get("error_description")
	if errStr != "" {
		return nil, fmt.Errorf("%s", errStr)
	}

	return &npool.AccessTokenInfo{
		AccessToken:        queryParams.Get("access_token"),
		RefreshAccessToken: queryParams.Get("refresh_token"),
		Scope:              queryParams.Get("scope"),
		TokenType:          queryParams.Get("token_type"),
	}, nil
}

type UserInfo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Message   string `json:"message"`
}

func GetUserInfo(accessToken string) (*npool.ThirdUserInfo, error) {
	var err error

	socksProxy := os.Getenv("ENV_REQUEST_PROXY")

	cli := resty.New()

	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	var resp *resty.Response
	resp, err = cli.R().
		SetHeader("accept", "application/json").
		SetHeader("Authorization", fmt.Sprintf("token %s", accessToken)).
		Get(constant.GithubGetUserInfoURL)
	if err != nil {
		logger.Sugar().Error("req err:", err)
		return nil, err
	}

	if !resp.IsSuccess() {
		logger.Sugar().Error("resp error: ", resp.StatusCode())
		return nil, fmt.Errorf("resp error: %v", resp.StatusCode())
	}

	var userInfoMap = make(map[string]interface{})
	if err = json.Unmarshal(resp.Body(), &userInfoMap); err != nil {
		logger.Sugar().Error("decode param err: ", err)
		return nil, err
	}

	jsonData, err := json.Marshal(userInfoMap)
	if err != nil {
		logger.Sugar().Error("decode json err：", err)
		return nil, err
	}

	var userInfo UserInfo
	err = json.Unmarshal(jsonData, &userInfo)
	if err != nil {
		logger.Sugar().Error("decode info err：", err)
		return nil, err
	}

	if userInfo.Message != "" {
		logger.Sugar().Error("get usetinfo err: ", userInfo.Message)
		return nil, fmt.Errorf("%s", userInfo.Message)
	}
	id := fmt.Sprintf("%d", userInfo.ID)

	info := &npool.ThirdUserInfo{
		ID:        id,
		Name:      userInfo.Name,
		Login:     userInfo.Login,
		AvatarURL: userInfo.AvatarURL,
	}

	return info, nil
}
