package main

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

const GANACHE_URL = "http://192.168.100.73:10545"
const DEPLOYER = "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823"
const USER = "0xd5a38dD251Aa8493C03954268FF851290051E634"

func main() {
	// client, _ := demo.NewClient(GANACHE_URL)
	// admin := demo.NewAccount(client, "0x872d3d0d6C5c1C0f5E8f9EEc2c4236cc9b5AB823", "0x6ff38a6fcde856869ddba8a1e0058a02cf81742f150607507d5245da607ba48f")
	// erc := demo.NewERC20Controller(client, demo.LOCAL_POINT_ADDR, admin)
	// amount := "0.001"
	// dec, _ := erc.Decimals()
	// amountInteger, _ := demo.ParseDecimalStringToIntegerString(amount, dec)
	// fmt.Println(erc.Approve(USER, amountInteger))
}
