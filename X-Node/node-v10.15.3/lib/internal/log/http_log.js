

const { HTTPLOGGER }=internalBinding('spdlog');


class HttpLogger{

    constructor(){
        this.httpLogger=new HTTPLOGGER();
    }


    static getInstance(){
        if(!this.instance){
            this.instance=new HttpLogger();
        }
        return this.instance;
    }

    info(type,content){
        this.httpLogger.info(type,content);
    }
    error(type,content){
        this.httpLogger.error(type,content);
    }
}

exports.info=function info(type,content) {
    return HttpLogger.getInstance().info(type,content);
}

exports.error=function error(type,content) {
    return HttpLogger.getInstance().error(type,content);
}