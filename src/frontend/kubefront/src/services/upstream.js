import bus from "../bus"
import api from "../api"
import router from '../router'

let service = {

    //Properties
    socket: null,
    subscriptions: [],

    //Broadcasts
    CONNECTED: "CONNECTED",
    MESSAGE: "MESSAGE",

    //Methods
    connect(){
        this.socket = new WebSocket("ws://localhost:8081/upstream?token="+encodeURIComponent(api.defaults.headers.common['token']));
        this.socket.onopen = (event) => {
            bus.$emit(this.CONNECTED)
            this.subscriptions.forEach(x => {
                this.socket.send(JSON.stringify({
                    action: "SUBSCRIBE",
                    entity: x,
                }))
            })
        };
        this.socket.onmessage = (event) => {
            let data = JSON.parse(event.data)
            bus.$emit(this.MESSAGE, data)
            bus.$emit(data.action, data.entity, data.data)
        };
    },
    subscribe(topic){
        if(this.socket != null){
            this.socket.send(JSON.stringify({
                action: "SUBSCRIBE",
                entity: topic,
            }))
        }
        this.subscriptions.push(topic);
    },
    unsubscribe(topic){
        if(this.socket != null){
            this.socket.send(JSON.stringify({
                action: "UNSUBSCRIBE",
                entity: topic,
            }))
        }
        this.subscriptions = this.subscriptions.filter(x => x != topic)
    }

}

//Export
export default service