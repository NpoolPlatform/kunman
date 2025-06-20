package facebook

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NpoolPlatform/kunman/framework/logger"
	constant "github.com/NpoolPlatform/kunman/mal/third/const"
	npool "github.com/NpoolPlatform/kunman/message/third/middleware/v1/oauth"
	"github.com/go-resty/resty/v2"
)

type Token struct {
	AccessToken      string        `json:"access_token"`
	TokenType        string        `json:"token_type"`
	ExpiresIn        int           `json:"expires_in"`
	ErrorDescription *MessageError `json:"error"`
}

func GetAccessToken(clientID, clientSecret, code, redirectURI string) (*npool.AccessTokenInfo, error) {
	var err error
	reqURL := fmt.Sprintf(
		"%s?client_id=%s&client_secret=%s&code=%s&redirect_uri=%s",
		constant.FacebookGetAccessTokenURL, clientID, clientSecret, code, redirectURI,
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
	var token Token
	err = json.Unmarshal([]byte(resp.String()), &token)
	if err != nil {
		logger.Sugar().Error("decode json err", err)
	}

	if token.ErrorDescription != nil {
		logger.Sugar().Error("resp err: ", token.ErrorDescription.Message)
		return nil, fmt.Errorf("resp err %v", token.ErrorDescription.Message)
	}

	return &npool.AccessTokenInfo{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
	}, nil
}

type PictureData struct {
	AvatarURL string `json:"url"`
}

type Picture struct {
	Data PictureData `json:"data"`
}

type MessageError struct {
	Message      string `json:"message"`
	Type         string `json:"type"`
	Code         int    `json:"code"`
	ErrorSubcode int    `json:"error_subcode"`
	FbtraceID    string `json:"fbtrace_id"`
}

type UserInfo struct {
	ID               string        `json:"id"`
	Name             string        `json:"name"`
	Email            string        `json:"email"`
	Avatar           Picture       `json:"picture"`
	ErrorDescription *MessageError `json:"error"`
}

func GetUserInfo(accessToken string) (*npool.ThirdUserInfo, error) {
	var err error

	socksProxy := os.Getenv("ENV_REQUEST_PROXY")

	cli := resty.New()

	if socksProxy != "" {
		cli = cli.SetProxy(socksProxy)
	}

	reqURL := constant.FacebookGetUserInfoURL + "?fields=id,name,email,picture&access_token=" + accessToken
	var resp *resty.Response
	resp, err = cli.R().
		SetHeader("accept", "application/json").
		Get(reqURL)
	if err != nil {
		logger.Sugar().Error("req err:", err)
		return nil, err
	}

	if !resp.IsSuccess() {
		logger.Sugar().Error("resp error: ", resp.StatusCode())
		return nil, fmt.Errorf("resp error: %v", resp.StatusCode())
	}

	userInfoMap := make(map[string]interface{})
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

	if userInfo.ErrorDescription != nil {
		logger.Sugar().Error("get usetinfo err: ", userInfo.ErrorDescription.Message)
		return nil, fmt.Errorf("%s", userInfo.ErrorDescription.Message)
	}

	info := &npool.ThirdUserInfo{
		ID:        userInfo.ID,
		Name:      userInfo.Name,
		Login:     userInfo.Name,
		AvatarURL: userInfo.Avatar.Data.AvatarURL,
	}

	return info, nil
}
