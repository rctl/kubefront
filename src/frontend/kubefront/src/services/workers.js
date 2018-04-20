import bus from "../bus"
import api from "../api"

let service = {

    //Properties
    jobs: [],

    //Broadcasts
    STARTED: "CONNECTED",
    COMPLETED: "MESSAGE",
    FAILED: "MESSAGE",

    get(){
        return new Promise((resolve, reject) => {
            resolve(this.jobs)
        })
    },
    getByEntity(id){
        return new Promise((resolve, reject) => {
            resolve(this.jobs.filter(x => x.entity == id))
        })
    },
    update(){
        return api.get("/workers/")
        .then(r => {
            this.jobs = r.data
        });
    }

}

bus.$on("JOB_STARTED", (id, data) => {
    service.jobs.push({
        id: id,
        entity: data,
    })
})

bus.$on("JOB_COMPLETED", (id, data) => {
    service.jobs = service.jobs.filter(x => x.id != id)
})

bus.$on("JOB_FAILED", (id, data) => {
    service.jobs = service.jobs.filter(x => x.id != id)
})

//Export
export default service