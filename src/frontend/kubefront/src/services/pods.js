import bus from "../bus"
import api from "../api"
import router from '../router'
let service = {

    //Properties
    lastUpdate: new Date(),
    pods: [],

    //Broadcasts
    broadcasts: {
        UPDATED: "UPDATED",
        ADDED: "ADDED"
    },

    //Methods
    refresh() {
        return new Promise((resolve, reject) => {
            api.get("/pods/")
            .then(r => {
                this.lastUpdate = new Date()
                this.pods = r.data
                bus.$emit(this.broadcasts.ADDED)
                bus.$emit(this.broadcasts.UPDATED)
            }).catch(reject);
        })
    },
    count() {
        return new Promise((resolve, reject) => {
           resolve(this.pods.length)
        })
    },
    list() {
        return new Promise((resolve, reject) => {
           resolve(this.pods)
        })
    }
}

//Init and listeners
bus.$on("POD_CHANGED", (entityID, data) => {
    let i = service.pods.findIndex(x => x.metadata.name == entityID)
    data.lastUpdate = new Date()
    if(i != -1){
        service.pods[i] = data
        bus.$emit(service.broadcasts.UPDATED)
    }else{
        service.pods.push(data)
        bus.$emit(service.broadcasts.ADDED)
    }
})

//Export
export default service