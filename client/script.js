class Message {

    price = 0
    quantity = 0
    amount = 0
    object = 0
    method = 0

    generateMessage() {
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

function generateSendPacket() {
    //console.log("sending....")
    packetAmount = Math.floor(Math.random() * 10) + 1
    packetArray = []
    for (let i = 0; i < packetAmount; i++) {
        p = new Message()
        p.generateMessage()
        packetArray.push(p)
    }   
    body = JSON.stringify(packetArray)
    //console.log(body);
    var request = new XMLHttpRequest();   
    request.open("POST", '/upload', true);
    request.setRequestHeader('Content-Type', 'application/json');
    request.send(body);
}

function onload(frequency) {
        console.log("onload...")
        console.log(frequency);
        ms = Math.floor(1000 / frequency)
        timerID = setInterval(generateSendPacket, ms)
};

function loadConfig(url, callback) {
    var request = new XMLHttpRequest();
    request.open('GET', url);
    request.onload = function() {
        if (this.status == 200) {
            obj = JSON.parse(this.responseText);
            frequency = obj.frequency   
            console.log(this.responseText); 
            callback(frequency)
        }
    };

    request.send();
}

loadConfig("/client/config.json", onload)