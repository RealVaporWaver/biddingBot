const axios = require("axios")
const PrivateKeyProvider = require("truffle-privatekey-provider")
const { OpenSeaPort, Network } = require("opensea-js")
 
const startup = async() => {
    const totalSupply = (await axios.get(`https://api.opensea.io/api/v1/collection/boredapeyachtclub/stats`)).data.stats.total_supply

    for(i = 0; i < totalSupply; i++) {

    }
}

startup()