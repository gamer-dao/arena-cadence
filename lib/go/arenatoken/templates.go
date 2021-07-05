package arenatoken

import (
	"log"

	arenacadence "github.com/arena/arena-cadence"
)

var (
	contractTemplate              string
	setupAccountTemplate          string
	mintArenaTemplate             string
	balanceTemplate               string
	transferTemplate              string
	transferAdministratorTemplate string
	destroyAdministratorTemplate  string
)

// read templates from embedded fs
func init() {
	contractTemplate = readContractTemplate("contracts/arenatoken.cdc")
	setupAccountTemplate = readTxTemplate("transactions/arenaToken/setup_account.cdc")
	mintArenaTemplate = readTxTemplate("transactions/arenaToken/mint_arena.cdc")
	balanceTemplate = readScriptTemplate("scripts/arenaToken/balance.cdc")
	transferTemplate = readTxTemplate("transactions/arenaToken/transfer.cdc")
	transferAdministratorTemplate = readTxTemplate("transactions/arenaToken/transfer_admin.cdc")
	destroyAdministratorTemplate = readTxTemplate("transactions/arenaToken/destroy_admin.cdc")
}

func readContractTemplate(path string) string {
	tpl, err := arenacadence.Contracts.ReadFile(path)
	if err != nil {
		log.Fatalf("Missing embedded template: %v", err)
	}
	return string(tpl)
}

func readTxTemplate(path string) string {
	tpl, err := arenacadence.Transactions.ReadFile(path)
	if err != nil {
		log.Fatalf("Missing embedded template: %v", err)
	}
	return string(tpl)
}

func readScriptTemplate(path string) string {
	tpl, err := arenacadence.Scripts.ReadFile(path)
	if err != nil {
		log.Fatalf("Missing embedded template: %v", err)
	}
	return string(tpl)
}
