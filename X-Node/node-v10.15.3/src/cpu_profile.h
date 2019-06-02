#ifndef SRC_CPU_PROFILE_H_
#define SRC_CPU_PROFILE_H_

#include "v8.h"

#include "v8-profiler.h"


namespace node {

    namespace profileSpace {


        class Profile {

        public:
            Profile();

            ~Profile();

        public:
            std::string New(const v8::CpuProfile *node);


        };
    }


}


#endif