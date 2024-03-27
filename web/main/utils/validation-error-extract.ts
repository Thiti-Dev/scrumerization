export default (e:any,state?: globalThis.Ref<any>) => {
    if(Array.isArray(e.graphQLErrors)){
        const ve = e.graphQLErrors[0].extensions.validation_errors as [string]
        const builtUpErr: Record<string,string> = {}
        for(const errorMsg of ve){
            const regex = /\/(\w+)\s->\s(.+)/;
            const match = errorMsg.match(regex);
            if (match) {
                const field = match[1];
                const message = match[2];
                builtUpErr[field] = message
            }
        }
        if(state){
            for (const key in state.value){
                if(builtUpErr.hasOwnProperty(key)){
                    state.value[key] = builtUpErr[key]
                }
            }
        }
        return builtUpErr
    }
}