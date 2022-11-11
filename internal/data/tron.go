package data

import (
	"fmt"
	"github.com/Rhymond/go-money"
	"github.com/f-sev/cryptowatcher/internal/utils"
	"github.com/getlantern/systray"
	"strconv"
	"strings"
)

type TronJson struct {
	Tokens []TronTokenJson `json:"tokens"`
}

type TronTokenJson struct {
	Balance   string `json:"balance"`
	TokenAbbr string `json:"tokenAbbr"`
}

var TronWallet = TronDataSource{
	CryptoDataSource{
		Name: "Tron Wallet",
	},
	TronCredentials{
		apiKey: "Qwerty",
	},
}

type TronCredentials struct {
	apiKey string
}

type TronDataSource struct {
	CryptoDataSource
	TronCredentials
}

func (t *TronDataSource) Collect() {
	t.Balance = make(BalanceType)
	address := "TGYfsCa9ymzW5hZDYt6sv8ubC2YqsxXsXd"

	var tronJson TronJson
	err := utils.GetJson(fmt.Sprintf("https://apilist.tronscan.org/api/account?address=%s", address), &tronJson)
	if err != nil {
		fmt.Printf("Error getting trone data	%s\n", err.Error())
	}

	//answer, _ := json.Marshal(tronJson)
	//fmt.Printf("tron data is: %s", answer)

	for _, token := range tronJson.Tokens {
		value, _ := strconv.Atoi(token.Balance)
		t.Balance[strings.ToUpper(token.TokenAbbr)] = float64(value) / 1000000
	}

}

func (t *TronDataSource) TotalFiat() *money.Money {
	return t.Balance.TotalFiat()
}

func (t *TronDataSource) Display() {
	item := systray.AddMenuItem(t.Name+"("+t.Balance.TotalFiat().Display()+")", "")
	t.Balance.Display(item)
}
