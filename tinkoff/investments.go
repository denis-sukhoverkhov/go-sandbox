package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var token = flag.String("token", os.Getenv("TINKOFF_TOKEN_API"), "your token")

//var isSandbox = flag.Bool("is_sandbox", true, "is sandbox env")

func main() {
	rand.Seed(time.Now().UnixNano()) // инициируем Seed рандома для функции requestID
	client := sdk.NewRestClient(*token)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Получение всех брокерских счетов")
	accounts, err := client.Accounts(ctx)
	if err != nil {
		log.Fatalln(errorHandle(err))
	}
	log.Printf("%+v\n", accounts)

	log.Println("Получение списка валютных и НЕ валютных активов портфеля для счета по-умолчанию")
	// Метод является совмещеним PositionsPortfolio и CurrenciesPortfolio
	portfolio, err := client.Portfolio(ctx, sdk.DefaultAccount)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", portfolio)
}

func errorHandle(err error) error {
	if err == nil {
		return nil
	}

	if tradingErr, ok := err.(sdk.TradingError); ok {
		if tradingErr.InvalidTokenSpace() {
			tradingErr.Hint = "Do you use sandbox token in production environment or vise verse?"
			return tradingErr
		}
	}

	return err
}
