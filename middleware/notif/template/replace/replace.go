package replace

import (
	"strings"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	npool "github.com/NpoolPlatform/kunman/message/notif/middleware/v1/template"
)

const (
	NotifTemplateVarName      = "{{ NAME }}"
	NotifTemplateVarMessage   = "{{ MESSAGE }}"
	NotifTemplateVarAmount    = "{{ AMOUNT }}"
	NotifTemplateVarCoinUnit  = "{{ COIN_UNIT }}"
	NotifTemplateVarDate      = "{{ DATE }}"
	NotifTemplateVarTime      = "{{ TIME }}"
	NotifTemplateVarAddress   = "{{ ADDRESS }}"
	NotifTemplateVarCode      = "{{ CODE }}"
	NotifTemplateVarIP        = "{{ IP }}"
	NotifTemplateVarLocation  = "{{ LOCATION }}"
	NotifTemplateVarUserAgent = "{{ USERAGENT }}"
)

func formatAmount(amount string) string {
	index := strings.Index(amount, ".")
	if index < 0 {
		index = len(amount)
	}
	_amount := []byte(amount[index:])
	count := 0
	for i := index; i > 0; i-- {
		_amount = append([]byte(amount[i-1:i]), _amount[0:]...)
		count++
		if count == 3 && i > 1 {
			_amount = append([]byte(","), _amount[0:]...)
			count = 0
		}
	}
	return string(_amount)
}

func ReplaceAll(pattern string, vars *npool.TemplateVars) string {
	if vars == nil {
		return pattern
	}

	if vars.Username != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarName, *vars.Username)
	}
	if vars.Message != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarMessage, *vars.Message)
	}
	if vars.Amount != nil {
		p := message.NewPrinter(language.English)
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarAmount, p.Sprintf("%s", formatAmount(*vars.Amount)))
	}
	if vars.CoinUnit != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarCoinUnit, *vars.CoinUnit)
	}
	if vars.Timestamp != nil {
		datetime := time.Unix(int64(*vars.Timestamp), 0).UTC()
		date := datetime.Format("2006-01-02")
		time1 := datetime.Format("15:04:05 UTC")

		pattern = strings.ReplaceAll(pattern, NotifTemplateVarDate, date)
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarTime, time1)
	}
	if vars.Address != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarAddress, *vars.Address)
	}
	if vars.Code != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarCode, *vars.Code)
	}
	if vars.IP != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarIP, *vars.IP)
	}
	if vars.Location != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarLocation, *vars.Location)
	}
	if vars.UserAgent != nil {
		pattern = strings.ReplaceAll(pattern, NotifTemplateVarUserAgent, *vars.UserAgent)
	}

	return pattern
}
