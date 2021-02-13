class Packet {

    price = 0
    quantity = 0
    amount = 0
    object = 0
    method = 0

    generatePacket() {
        this.price = this.getRandomInt()
        this.quantity = this.getRandomInt()
        this.amount = this.getRandomInt()
        this.object = this.getRandomInt()
        this.method = this.getRandomInt()
    }

    getRandomInt() {
        return Math.floor(Math.random() * 1000)
    }
}

function sendPacket() {
    console.log("sending...")
}

frequency = 2
packetAmount = Math.floor(Math.random() * 10)
ms = Math.floor(1000 / frequency)

//timerID = setInterval(sendPacket, ms)

packetArray = []
for (let i = 0; i < packetAmount; i++) {
    p = new Packet()
    p.generatePacket()
    packetArray.push(p)
}
//p = new Packet()
//p.generatePacket()
console.log(JSON.stringify(packetArray))