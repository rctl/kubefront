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
    get(kind, namespace, name){
        return new Promise((resolve, reject) => {
            resolve(this.jobs.filter(x => x.kind == kind && x.namespace == namespace && x.name == name))
        })
    },
    update(){
        return api.get("/workers/")
        .then(r => {
            r.data.forEach(x => {
                if(x.entity.includes("/")){
                    this.jobs.push({
                        id: x.id,
                        kind: x.entity.split("/")[0],
                        namespace: x.entity.split("/")[1],
                        name: x.entity.split("/")[2],
                        entity: x.entity,
                    })
                }else{
                    service.jobs.push({
                        id: x.id,
                        entity: x.entity,
                    })
                }
            });
        });
    }

}

bus.$on("JOB_STARTED", (id, data) => {
    if(data.includes("/")){
        //V2 (Fully Qualified ID)
        service.jobs.push({
            id: id,
            kind: data.split("/")[0],
            namespace: data.split("/")[1],
            name: data.split("/")[2],
            entity: data,
        })
    }else{
        service.jobs.push({
            id: id,
            entity: data,
        })
    }
})

bus.$on("JOB_COMPLETED", (id, data) => {
    service.jobs = service.jobs.filter(x => x.id != id)
})

bus.$on("JOB_FAILED", (id, data) => {
    service.jobs = service.jobs.filter(x => x.id != id)
})

//Export
export default service