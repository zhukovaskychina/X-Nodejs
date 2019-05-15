#include "node.h"
#include "node_log.h"
#include "env.h"
#include "debug_utils.h"
#include "util-inl.h"
#include "node_internals.h"
#include "v8.h"
#include "spdlog/common.h"
#include "spdlog/spdlog.h"
#include "spdlog/sinks/basic_file_sink.h"
#include "spdlog/sinks/daily_file_sink.h"
#include "spdlog/logger.h"
#include "spdlog/async.h"
#include "spdlog/sinks/basic_file_sink.h"
namespace node{
    namespace loggerSpace{
        using namespace node;
        using v8::Context;
        using v8::Local;
        using v8::MaybeLocal;
        using v8::Null;
        using v8::Number;
        using v8::Object;
        using v8::String;
        using v8::Value;
        using v8::FunctionCallbackInfo;
        using v8::FunctionTemplate;

        std::shared_ptr<spdlog::async_logger> nml_logger;
        std::shared_ptr<spdlog::async_logger> http_logger;
        std::shared_ptr<spdlog::async_logger> error_logger;
        Logger::Logger() {
            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {
                int pid = getpid();
                string filename = "Node-"+std::to_string(getpid())+"" + ".log";
                string filePath = std::string(logDir) + "/" + filename;
                auto daily_sink = std::make_shared<spdlog::sinks::daily_file_sink_mt>(filePath, 23, 59);

                spdlog::init_thread_pool(8192 , 1);
                nml_logger =std::make_shared<spdlog::async_logger>(std::to_string(pid), daily_sink, spdlog::thread_pool(), async_overflow_policy::block);
            }


        }

        Logger::Logger(const node::loggerSpace::Logger &) {

        }

        Logger& Logger::operator=(const node::loggerSpace::Logger &) {

        }

        Logger::~Logger() {
            spdlog::drop(nml_logger->name());
        }

        Logger* Logger::instance;
        Logger* Logger::getInstance() {
            if (NULL==instance){
                instance = new Logger();
            }
            return instance;
        }

        void Logger::info(const std::string &level, const std::string &filename) {
            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {
                nml_logger->info(level.c_str(),filename);
            }

        }

        void Logger::error( const std::string &filename) {
            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {
                nml_logger->error(filename);
            }
        }


        HttpLogger::HttpLogger() {
            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {
                int pid = getpid();
                string filename = "Node-"+std::to_string(getpid())+"-http" + ".log";
                string filePath = std::string(logDir) + "/" + filename;

                string errorFilename = "Node-error.log";
                string errorFilePath = std::string(logDir) + "/" + errorFilename;
                spdlog::init_thread_pool(8192, 1);
                auto daily_http_sink = std::make_shared<spdlog::sinks::daily_file_sink_mt>(filePath, 23, 59);
                auto daily_error_sink = std::make_shared<spdlog::sinks::daily_file_sink_mt>(errorFilePath, 23, 59);
                http_logger = std::make_shared<spdlog::async_logger>(std::to_string(pid), daily_http_sink, spdlog::thread_pool(), async_overflow_policy::block);

                error_logger=std::make_shared<spdlog::async_logger>(std::to_string(pid), daily_error_sink, spdlog::thread_pool(), async_overflow_policy::block);
            }


        }

        HttpLogger::HttpLogger(const node::loggerSpace::HttpLogger &) {}

        HttpLogger& HttpLogger::operator=(const node::loggerSpace::HttpLogger &) {}

        HttpLogger::HttpLogger(node::Environment *env, v8::Local <Object> object) {
          //  CHECK(args.IsConstructCall());

             new HttpLogger();
        }
        HttpLogger::~HttpLogger() {
            spdlog::drop(http_logger->name());
        }

        HttpLogger* HttpLogger::instance;
        HttpLogger* HttpLogger::getInstance() {
            if (NULL==instance){
                instance = new HttpLogger();
            }
            return instance;
        }

        void HttpLogger::New(const v8::FunctionCallbackInfo <Value> &args) {
            Environment* env = Environment::GetCurrent(args);
            CHECK(args.IsConstructCall());

            if (NULL==instance){
                instance = new HttpLogger(env,args.This());
            }

        }

        void HttpLogger::info( const FunctionCallbackInfo<Value>& args ){

            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {
                std::string fileName;
                std::string fileType;
                Environment* env = Environment::GetCurrent(args);

                if (args.Length() == 2 && !args[0]->IsNullOrUndefined() && !args[1]->IsNullOrUndefined()) {
                    Utf8Value typeValue(
                            args.GetIsolate(),
                            args[0]->ToString(env->context()).FromMaybe(v8::Local<v8::String>()));
                    fileType.append(typeValue.out(),typeValue.length());
                    Utf8Value value(
                            args.GetIsolate(),
                            args[1]->ToString(env->context()).FromMaybe(v8::Local<v8::String>()));
                    fileName.append(value.out(), value.length());
                }
                http_logger->info(fileType.c_str(),fileName);
            }

        }

        void HttpLogger::errorInner(  const std::string &level, const std::string &filename) {
            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {

                error_logger->error(level.c_str(),filename);
            }
        }
        void HttpLogger::error(  const FunctionCallbackInfo<Value>& args ) {
            char *logDir = nullptr;
            logDir = getenv("NODEJS_LOG_DIR");
            if (logDir != NULL) {
                std::string fileName;
                std::string fileType;
                Environment* env = Environment::GetCurrent(args);
                if (args.Length() == 2 && !args[0]->IsNullOrUndefined() && !args[1]->IsNullOrUndefined()) {
                    Utf8Value typeValue(
                            args.GetIsolate(),
                            args[0]->ToString(env->context()).FromMaybe(v8::Local<v8::String>()));
                    fileType.append(typeValue.out(),typeValue.length());
                    Utf8Value value(
                            args.GetIsolate(),
                            args[1]->ToString(env->context()).FromMaybe(v8::Local<v8::String>()));
                    fileName.append(value.out(), value.length());
                }

                error_logger->error(fileType.c_str(),fileName);
            }
        }

        static std::string & trimstr(std::string &s) {
            if(s.empty()) {
                return s;
            }
            string character = "";
            for(int i = 0; i < 33; i++) {
                character += char(i);
            }
            character += char(127);
            s.erase(0,s.find_first_not_of(character));
            s.erase(s.find_last_not_of(character) + 1);
            return s;
        }

        static int _vscprintfanother(const char *format, va_list pargs) {
            int retval;
            va_list argcopy;
            va_copy(argcopy, pargs);
            retval = vsnprintf(NULL, 0, format, argcopy);
            va_end(argcopy);
            return retval;
        }


        static std::string _getAppendString(const char *format, ...) {
            std::string result;
            va_list ap;
            int n=0;
            va_start(ap,format);
            result.resize(loggerSpace::_vscprintfanother(format,ap)+1,0);
            n=vsnprintf(const_cast<char *>(result.c_str()),result.length(),format,ap);
            va_end(ap);
            return result;
        }

     //   namespace {

            void InitializeLog ( Local<Object> target,
                              Local<Value> unused,
                              Local<Context> context){

                Environment* env = Environment::GetCurrent(context);
                Local<String> workerString =
                        FIXED_ONE_BYTE_STRING(env->isolate(), "HTTPLOGGER");



                {
                       Local<FunctionTemplate> w = env->NewFunctionTemplate(HttpLogger::New);
                    w->InstanceTemplate()->SetInternalFieldCount(1);
                    env->SetProtoMethod(w,"info",HttpLogger::info);
                    env->SetProtoMethod(w,"error",HttpLogger::error);
                    w->SetClassName(workerString);
                    target->Set(workerString,w->GetFunction(env->context()).ToLocalChecked());
                }
                // env->SetMethod(target,"getInstance",HttpLogger::operatorHttpLogger);
                //     env->SetMethod(target,"info",HttpLogger::info);


            }

       // }
    }



}


NODE_MODULE_CONTEXT_AWARE_INTERNAL(spdlog, node::loggerSpace::InitializeLog)