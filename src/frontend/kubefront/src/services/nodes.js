import bus from "../bus"
import api from "../api"
import router from '../router'
let service = {

    //Properties
    lastUpdate: new Date(),
    nodes: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "UPDATED"
    },

    //Methods
    refresh() {
        return new Promise((resolve, reject) => {
            api.get("/nodes/")
            .then(r => {
                this.lastUpdate = new Date()
                this.nodes = r.data
                bus.$emit(this.broadcasts.UPDATED)
            }).catch(reject);
        })
    },
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.nodes.length)
        })
    },
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.nodes)
        })
    },
}

//Export
export default service