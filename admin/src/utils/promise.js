

const wrapPromise = (promise)=> {
    return new Promise((resolve, reject) => {
      promise
      .then((info) => resolve({ isok: true, info }))
      .catch((err) => resolve({ isok: false, err }))
    })
  }

function promiseX(promiseLst){
    let pLst = promiseLst.map(r=>{
        return wrapPromise(r)
    })
    return Promise.all(pLst).then(_e=>{
        return _e.map(e=>{
            if (e.isok) {
                return e.info
            } else {
                return e.err
            }
        })
    })
}

export default promiseX