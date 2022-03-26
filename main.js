const axios = require("axios")
const PrivateKeyProvider = require("truffle-privatekey-provider")
const { OpenSeaPort, Network } = require("opensea-js")
const fs = require("fs")

const buyer_provider = new PrivateKeyProvider("")
const buyer_seaport = new OpenSeaPort(buyer_provider, {
    networkName: Network.Main
})

const configData = fs.readFileSync("./config.json")
const config = JSON.parse(configData)



const startup = async() => {
    const totalSupply = (await axios.get(`https://api.opensea.io/api/v1/collection/boredapeyachtclub/stats`)).data.stats.total_supply

    for(i = 0; i < totalSupply+1; i++) {
        const offer = await buyer_seaport.createBuyOrder({
            asset: {
                tokenId: i,
                tokenAddress: config.address,
            },
            accountAddress: config.accountAddress,
            startAmount: config.bidAmount,
            expirationTime: Math.round(Date.now() / 1000 + 60 * 60 * config.expirationTime)
        })
    }
}

startup()