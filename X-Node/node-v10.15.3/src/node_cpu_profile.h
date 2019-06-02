#ifndef SRC_NODE_CPU_PROFILE_H_
#define SRC_NODE_CPU_PROFILE_H_

#include "v8.h"
#include "v8-profiler.h"
#include "node.h"

namespace node{


    namespace profileSpace{
        class CpuProfiler{
        public:

            CpuProfiler();

            ~CpuProfiler();

            static void InitializeCpuProfiler();



            void StartProfiling();

            std::string StopProfiling();

            void Delete();

            void SetSamplingInterval();


        };
    }

}


#endif