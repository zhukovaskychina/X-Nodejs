#ifndef SRC_NODE_LOG_H_
#define SRC_NODE_LOG_H_
#include <pthread.h>

#include <iostream>
#include "node.h"

#include "env.h"
#include "spdlog/spdlog.h"
#include "spdlog/logger.h"
#include "v8.h"
#include "util-inl.h"

using namespace std;

using namespace spdlog;

namespace node {

    namespace loggerSpace{

        class Logger{

        public:
            static Logger* getInstance();

            void info(const std::string &level,const std::string &filename);

            void error( const std::string &filename);


        private:



            Logger();

            Logger(const Logger&);

            ~Logger();

            Logger&operator=(const Logger&);

            static Logger* instance;
        };

        class HttpLogger{

        public:
            static HttpLogger* getInstance();

            static void info(const v8::FunctionCallbackInfo<v8::Value>& args);

            static void errorInner(const std::string &level,const std::string &filename);
            static void error( const v8::FunctionCallbackInfo<v8::Value>& args);
            static void New(const v8::FunctionCallbackInfo<v8::Value>& args);
          //  static HttpLogger* operatorHttpLogger(const v8::FunctionCallbackInfo<v8::Value>& args);
        private:



            HttpLogger();

            HttpLogger(Environment* env,
                       v8::Local<v8::Object> object);

            HttpLogger(const HttpLogger&);

            ~HttpLogger();

            HttpLogger&operator=(const HttpLogger&);

            static HttpLogger* instance;
        };



        static std::string& trimstr(std::string &s);

        static int _vscprintfanother(const char * format, va_list pargs);

        static std::string _getAppendString(const char* format,...);



    }
}
#endif