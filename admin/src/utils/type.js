let isError = function(e){
    return e && e.stack && e.message && typeof e.stack === 'string' 
        && typeof e.message === 'string';
}

const tp = {
    isError
}

export default tp