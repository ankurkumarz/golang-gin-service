package main

import (
	Service "entitlement/entitlementsvc"
)

func main() {
	router := Service.Routes()
	router.Run("localhost:8081")
}
